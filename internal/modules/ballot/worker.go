// Package ballot handles closed ballot finalization.
// It implements an oracle which uses Fantom API to collect
// ballot participants' account total and feeds it into
// the smart contract to allow ballot winner calculation.
package ballot

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/machinebox/graphql"
	"math"
	"math/big"
	"oracle-watchdog/internal/logger"
	"oracle-watchdog/internal/modules/utils"
	"sync"
	"time"
)

const graphQLTotalValueQuery = `
query ($addr: Address!) {
    account(address: $addr) {
        totalValue
    }
}`

// participantsWeightPushPackSize represents the number of participants
// pushed into the voting contract at once.
const participantsWeightPushPackSize = 20

// FinalizingWorker represents a worker structure and processor for single
// closed ballot processing.
type FinalizingWorker struct {
	cfg          *FinalizingOracleConfig
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
	Name      string         `json:"name"`
	Address   common.Address `json:"address"`
	Proposals []Proposal     `json:"proposals"`
}

type Proposal struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

// Participant represents a single address involved in the ballot voting.
type Participant struct {
	Address   common.Address
	Total     *big.Int
	Vote      int
	TimeStamp int64
}

// NewWorker creates a new ballot oracle worker structure
// ready to run ballots finalization.
func NewWorker(
	rpc *ethclient.Client,
	api *graphql.Client,
	wg *sync.WaitGroup,
	queue chan ClosedBallot,
	log logger.Logger,
	cfg *FinalizingOracleConfig,
) *FinalizingWorker {
	// make the worker
	w := FinalizingWorker{
		cfg:          cfg,
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
	go fw.waitForBallot()
}

// waitForBallot is the main worker routine where the actual worker
// business happens.
func (fw *FinalizingWorker) waitForBallot() {
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

		// process the ballot
		if err := fw.processBallot(); err != nil {
			fw.log.Errorf("can not process ballot %s; %s", fw.ballot.Address.String(), err.Error())
		}
	}
}

// processBallot processes the received ballot and updates
// the ballot data and status on chain.
func (fw *FinalizingWorker) processBallot() error {
	// connect the contract
	contract, err := NewBallotContract(fw.ballot.Address, fw.eth)
	if err != nil {
		fw.log.Errorf("can not access ballot contract; %s", err.Error())
		return err
	}

	// collect list of all participants of the ballot
	party, err := fw.participants(contract)
	if err != nil {
		fw.log.Debugf("ballot processing error; %s", err.Error())
		return err
	}

	// do we have any participants at all?
	if party == nil {
		fw.log.Noticef("nil voters received for %s", fw.ballot.Address.String())
		return nil
	}

	// get the signing transactor
	sig, err := utils.Transactor(fw.log, &fw.cfg.KeyStore, &fw.cfg.KeySecret)
	if err != nil {
		fw.log.Errorf("can not get push transactor; %s", err.Error())
		return err
	}

	// push participants data to the ballot contract
	if err := fw.pushVoters(contract, sig, party); err != nil {
		fw.log.Debugf("can not push voters; %s", err.Error())
		return err
	}

	// finalize the ballot
	if err := fw.finalize(contract, sig); err != nil {
		fw.log.Debugf("can not finalize ballot; %s", err.Error())
		return err
	}

	// calculate and log the winner of the ballot
	fw.winner(party)
	return nil
}

// participants collects all the participants of the ballot with their votes
// so we can push the weights/totals into the contract and finalize it.
func (fw *FinalizingWorker) participants(contract *BallotContract) ([]Participant, error) {
	// inform
	fw.log.Debugf("loading participants of %s", fw.ballot.Address.String())

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

		// inform
		fw.log.Debugf("participants of %s loaded", fw.ballot.Address.String())
	}()

	// prep the list
	list := make([]Participant, 0)

	// loop the filter
	for it.Next() {
		// make the participant
		party := Participant{
			Address:   it.Event.Voter,
			Vote:      int(it.Event.Vote.Uint64()),
			Total:     fw.accountTotal(it.Event.Voter),
			TimeStamp: time.Now().UTC().Unix(),
		}

		// push to the list
		if party.Total != nil {
			list = append(list, party)
		}

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
func (fw *FinalizingWorker) accountTotal(addr common.Address) *big.Int {
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
		return nil
	}

	// log the value
	fw.log.Debugf("account total for [%s] is %s", addr.String(), res.Account.TotalValue.String())
	return res.Account.TotalValue.ToInt()
}

// pushVoters updates the ballot voters information inside the contract
// so the winner can be decided.
func (fw *FinalizingWorker) pushVoters(contract *BallotContract, sig *bind.TransactOpts, party []Participant) error {
	// inform
	fw.log.Debugf("pushing %d voters stats into %s", len(party), fw.ballot.Address.String())

	// prep the pack containers
	voters := make([]common.Address, 0)
	totals := make([]*big.Int, 0)
	stamps := make([]*big.Int, 0)

	// loop the whole party and each time build a pack to be pushed
	var index int
	for index < len(party) {
		// add the participant to the pack
		voters = append(voters, party[index].Address)
		totals = append(totals, party[index].Total)
		stamps = append(stamps, big.NewInt(party[index].TimeStamp))
		index++

		// if we reached the pack boundary, push the pack to contract
		if len(voters) >= participantsWeightPushPackSize {
			// make the push
			if err := fw.pushPack(contract, sig, voters, totals, stamps); err != nil {
				fw.log.Errorf("can not push a voters pack; %s", err.Error())
				return err
			}

			// clear the pack so we can start again
			voters = voters[:0]
			totals = totals[:0]
			stamps = stamps[:0]
		}
	}

	// any remaining pack members left to process?
	if 0 < len(voters) {
		// make the last pack push
		if err := fw.pushPack(contract, sig, voters, totals, stamps); err != nil {
			fw.log.Errorf("can not push a voters pack; %s", err.Error())
			return err
		}
	}

	return nil
}

// pushPack uses the contract transaction call to make a push to the contract
func (fw *FinalizingWorker) pushPack(contract *BallotContract, sig *bind.TransactOpts, voters []common.Address, totals []*big.Int, stamps []*big.Int) error {
	// make a transaction
	tx, err := contract.FeedWeights(sig, voters, totals, stamps)
	if err != nil {
		fw.log.Errorf("voters transaction failed; %s", err.Error())
		return err
	}

	// inform
	fw.log.Infof("voters pack for %s sent as %s", fw.ballot.Address.String(), tx.Hash().String())
	return nil
}

// finalize locks the ballot in finished state and allows winner calculation.
func (fw *FinalizingWorker) finalize(contract *BallotContract, sig *bind.TransactOpts) error {
	// inform
	fw.log.Debugf("finalizing ballot %s", fw.ballot.Address.String())

	// invoke the finalize
	tx, err := contract.Finalize(sig)
	if err != nil {
		fw.log.Errorf("can not finalize ballot %s; %s", fw.ballot.Address.String(), err.Error())
		return err
	}

	// inform about success
	fw.log.Infof("ballot %s has been finalized by %s", fw.ballot.Address.String(), tx.Hash().String())
	return nil
}

// winner calculates the winning proposal of the ballot.
func (fw *FinalizingWorker) winner(party []Participant) {
	// inform
	fw.log.Debugf("calculating ballot %s results", fw.ballot.Address.String())

	// container for votes
	votes := make([]uint64, len(fw.ballot.Proposals)+1)

	// container for weights
	weights := make([]*big.Int, len(fw.ballot.Proposals)+1)

	// loop all voters
	for _, voter := range party {
		// advance votes counter
		votes[voter.Vote]++

		// make sure to list the weight
		if weights[voter.Vote] == nil {
			weights[voter.Vote] = new(big.Int)
		}

		// advance weight
		weights[voter.Vote] = new(big.Int).Add(weights[voter.Vote], voter.Total)
	}

	// log results
	for _, prop := range fw.ballot.Proposals {
		// do we have it?
		if weights[prop.Id] != nil {
			w := new(big.Int).Div(weights[prop.Id], big.NewInt(int64(math.Pow10(18)))).Uint64()
			fw.log.Noticef("%s: Proposal #%d %s, votes %d, weight %d FTM", fw.ballot.Name, prop.Id, prop.Name, votes[prop.Id], w)
		}
	}
}
