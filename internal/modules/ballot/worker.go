// Package ballot handles closed ballot finalization.
// It implements an oracle which uses Fantom API to collect
// ballot participants' account total and feeds it into
// the smart contract to allow ballot winner calculation.
package ballot

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/machinebox/graphql"
	"oracle-watchdog/internal/logger"
	"sync"
	"time"
)

const graphQLTotalValueQuery = `
query ($addr: Address!) {
    account(address: $addr) {
        totalValue
    }
}`

// FinalizingWorker represents a worker structure and processor for single
// closed ballot processing.
type FinalizingWorker struct {
	eth          *ethclient.Client
	api          *graphql.Client
	sigTerminate chan bool
	waitGroup    *sync.WaitGroup
	jobQueue     chan ClosedBallot
	log          logger.Logger
	ballot       *ClosedBallot
}

// ClosedBallot represents a ballot record received from the remote API server
// and prepared to be finished by the oracle.
type ClosedBallot struct {
	Name    string         `json:"name"`
	Address common.Address `json:"address"`
}

// Participant represents a single address involved in the ballot voting.
type Participant struct {
	Address   common.Address
	Total     hexutil.Big
	TimeStamp int64
}

// NewWorker creates a new ballot oracle worker structure
// ready to run ballots finalization.
func NewWorker(rpc *ethclient.Client, api *graphql.Client, wg *sync.WaitGroup, queue chan ClosedBallot, log logger.Logger) *FinalizingWorker {
	// make the worker
	w := FinalizingWorker{
		eth:          rpc,
		api:          api,
		waitGroup:    wg,
		jobQueue:     queue,
		log:          log,
		sigTerminate: make(chan bool, 1),
	}

	return &w
}

// Terminate signals the worker to stop and quit.
func (fw *FinalizingWorker) Terminate() {
	// log the termination request and pass it to the actual execution thread
	log.Debug("ballot worker terminates")
	fw.sigTerminate <- true
}

// Run starts the worker process.
func (fw *FinalizingWorker) Run() {
	// log the action
	log.Debug("starting new ballot worker")

	// register myself with the oracle
	fw.waitGroup.Add(1)

	// execute me
	go fw.execute()
}

// execute is the main worker routine where the actual worker
// business happens.
func (fw *FinalizingWorker) execute() {
	// make sure to sign off when done
	defer func() {
		// log the action
		fw.log.Debug("ballot worker ended")

		// sign off the oracle
		fw.waitGroup.Done()
	}()

	// loop inside this forever
	for {
		// wait for a new closed ballot to be processed
		// or a termination signal
		select {
		case ballot := <-fw.jobQueue:
			// process this received ballot
			fw.log.Debugf("new ballot [%s] received by worker thread", ballot.Address.String())
			fw.ballot = &ballot
		case <-fw.sigTerminate:
			// exit the worker main thread
			return
		}

		// connect the contract
		contract, err := NewBallotContract(fw.ballot.Address, fw.eth)
		if err != nil {
			fw.log.Errorf("can not access ballot contract; %s", err.Error())
			continue
		}

		// collect list of all participants of the ballot
		party, err := fw.participants(contract)
		if err != nil {
			fw.log.Debugf("ballot processing error; %s", err.Error())
			continue
		}

		// do we have any participants at all?
		if party == nil || len(party) == 0 {
			continue
		}

		// push participants data to the ballot contract

		// finalize the ballot
	}
}

// participants collects all the participants of the ballot with their votes
// so we can push the weights/totals into the contract and finalize it.
func (fw *FinalizingWorker) participants(contract *BallotContract) ([]Participant, error) {
	// create the filter to iterate over the list
	it, err := contract.FilterVoted(nil, []common.Address{fw.ballot.Address}, nil)
	if err != nil {
		return nil, err
	}

	// make sure to close the iterator on leaving
	defer func() {
		err := it.Close()
		if err != nil {
			fw.log.Errorf("failed to close votes iterator; %s", err.Error())
		}
	}()

	// prep the list
	list := make([]Participant, 0)

	// loop the filter
	for it.Next() {
		// make the participant
		party := Participant{
			Address:   it.Event.Voter,
			Total:     fw.accountTotal(it.Event.Voter),
			TimeStamp: time.Now().UTC().Unix(),
		}

		// push to the list
		list = append(list, party)

		// check for possible termination request
		select {
		case <-fw.sigTerminate:
			// re-enter consumed termination signal and end the loader
			fw.sigTerminate <- true
			return nil, nil
		default:
			// just continue
		}
	}

	// return the list of participants
	return list, nil
}

// accountTotal pulls the current total value of the given account
// from remote API server using GraphQL call.
func (fw *FinalizingWorker) accountTotal(addr common.Address) hexutil.Big {
	// log action
	fw.log.Debugf("loading account total for [%s]", addr.String())

	// prep the graphQL request
	req := graphql.NewRequest(graphQLTotalValueQuery)
	req.Var("addr", addr.String())

	// prep response structure for decoding
	var res struct {
		Account struct {
			TotalValue hexutil.Big
		}
	}

	// execute the request and parse the response
	if err := fw.api.Run(context.Background(), req, &res); err != nil {
		fw.log.Errorf("can not pull account total for [%s]; %s", addr.String(), err.Error())
		return hexutil.Big{}
	}

	// log the value
	fw.log.Debugf("account total for [%s] is %s", addr.String(), res.Account.TotalValue.String())
	return res.Account.TotalValue
}
