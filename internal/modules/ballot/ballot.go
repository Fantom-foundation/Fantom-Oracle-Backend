// Package ballot handles closed ballot finalization.
// It implements an oracle which uses Fantom API to collect
// ballot participants' account total and feeds it into
// the smart contract to allow ballot winner calculation.
package ballot

//go:generate abigen --abi ./contract/ballot.abi --pkg ballot --type BallotContract --out ./bridge.go

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinebox/graphql"
	"oracle-watchdog/internal/config"
	"oracle-watchdog/internal/supervisor"
	"sync"
	"time"
)

// What the ballot finalizing oracle does:
//  1. downloads list of closed and unfinished ballots from the API server
//  2. it collects total value of participants' accounts for each ballot
//  3. uses TRX to populate ballot contracts with the accounts' totals
//  4. uses TRX to finalize ballots
//  5. waits certain amount of time to repeat the process from #1

// graphQLBallotsListQuery represents a GraphQL query to API server for
// loading list of closed and unfinished ballots.
const graphQLBallotsListQuery = `
query {
    ballotsClosed(finalized: false) {
        name
        address
    }
}
`

// workersPoolSize represents the number of workers
// waiting for the closed ballots to process them.
const workersPoolSize = 2

// jobChannelBuffer represents the size of workers job channel
// buffer.
const jobChannelBuffer = 10

// FinalizingOracle defines an oracle module for feeding
// ballot participants' account totals and finalizing
// closed ballots to allow winner calculations.
type FinalizingOracle struct {
	cfg          *FinalizingOracleConfig
	sup          supervisor.Supervisor
	apiClient    *graphql.Client
	workers      []*FinalizingWorker
	workersGroup *sync.WaitGroup
	sigClose     chan bool
	jobQueue     chan ClosedBallot
}

// New creates a new instance of the ballot oracle module.
func New(cfg *config.ModuleConfig, sup supervisor.Supervisor) (supervisor.Oracle, error) {
	// read the module configuration
	cf, err := configuration(cfg.CfgPath)
	if err != nil {
		sup.Log().Criticalf("can not load oracle configuration %s; %s", cfg.CfgPath, err.Error())
		return nil, err
	}

	// make the ballot finalizing oracle
	bo := FinalizingOracle{
		cfg:          cf,
		sup:          sup,
		apiClient:    graphql.NewClient(cf.ApiUrl),
		workersGroup: new(sync.WaitGroup),
		sigClose:     make(chan bool, 1),
		jobQueue:     make(chan ClosedBallot, jobChannelBuffer),
	}

	// make sure to add this oracle to the supervisor
	sup.AddOracle(&bo)
	return &bo, nil
}

// Terminate signals the finalizing oracle to stop the process and close.
func (fo *FinalizingOracle) Terminate() {
	// signal workers to terminate
	if fo.workers != nil {
		for _, w := range fo.workers {
			w.Terminate()
		}
	}

	// signal the oracle itself to terminate
	fo.sigClose <- true
}

// Run starts the ballot finalization oracle business.
func (fo *FinalizingOracle) Run() {
	// log we are done
	fo.sup.Log().Noticef("starting %s oracle", fo.cfg.Name)

	// prep Eth client
	eth := ethclient.NewClient(fo.sup.Lachesis())

	// make workers
	fo.workers = make([]*FinalizingWorker, workersPoolSize)
	for i := 0; i < workersPoolSize; i++ {
		fo.workers[i] = NewWorker(eth, fo.apiClient, fo.workersGroup, fo.jobQueue, fo.sup.Log())
		fo.workers[i].Run()
	}

	// signal supervisor we are good to go
	fo.sup.OracleStarted()
	go fo.watchBallots()
}

// watchBallots performs the finalizing oracle main function
// by loading the ballots list and processing each of them
// in a managed sub-routine.
func (fo *FinalizingOracle) watchBallots() {
	// signal the oracle has ended
	defer func() {
		// wait runners to finish
		fo.workersGroup.Wait()

		// log we are done
		fo.sup.Log().Noticef("oracle %s terminated", fo.cfg.Name)

		// signal supervisor we are done
		fo.sup.OracleDone()
	}()

	// what delay do we wait before re-starting the scanner again
	var delay = time.Duration(fo.cfg.ScanDelay) * time.Second

	// loop the function
	for {
		// check ballots and process them as needed
		fo.checkBallots()

		// wait for either termination signal, or timeout
		select {
		case <-fo.sigClose:
			// stop signal received?
			return
		case <-time.After(delay):
			// we go around
		}
	}
}

// listBallots loads a list of closed and unfinished ballots from the API server
// so they can be processed by this oracle.
func (fo *FinalizingOracle) checkBallots() {
	// log what we do
	fo.sup.Log().Debugf("oracle %s starts ballots check", fo.cfg.Name)

	// download closed ballots list
	list, err := fo.listBallots()
	if err != nil {
		fo.sup.Log().Criticalf("can not get the list of closed ballots; %s", err.Error())
		return
	}

	// finalize found ballots by pushing them into the work queue
	// waiting workers will pull them and process them
	for _, ba := range list {
		fo.jobQueue <- ba
	}
}

// listBallots loads a list of closed and unfinished ballots from the API server
// so they can be processed by this oracle.
func (fo *FinalizingOracle) listBallots() ([]ClosedBallot, error) {
	// prep new ballots list query
	req := graphql.NewRequest(graphQLBallotsListQuery)

	// prep response container
	var res struct {
		BallotsClosed []ClosedBallot
	}

	// execute the query and get the result
	if err := fo.apiClient.Run(context.Background(), req, &res); err != nil {
		fo.sup.Log().Errorf("closed ballots list API query failed; %s", err.Error())
		return nil, err
	}

	return res.BallotsClosed, nil
}
