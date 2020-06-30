// Package supervisor manages running modules and takes care
// of the inter-modules communication, if required.
package supervisor

import (
	"github.com/ethereum/go-ethereum/rpc"
	"oracle-watchdog/internal/logger"
	"sync"
)

// WatchdogSupervisor implements the
type WatchdogSupervisor struct {
	nodeRpcPath string
	log         logger.Logger
	client      *rpc.Client
	oracles     []Oracle
	sigClose    chan bool
	waitGroup   sync.WaitGroup
}

// New creates new instance of the Supervisor.
func New(rpcPath string, log logger.Logger) Supervisor {
	// make the oracle
	ws := WatchdogSupervisor{
		log:         log,
		nodeRpcPath: rpcPath,
		oracles:     make([]Oracle, 0),
		sigClose:    make(chan bool, 1),
	}

	return &ws
}

// AddModule adds new module to the supervisor.
func (ws *WatchdogSupervisor) AddOracle(ora Oracle) {
	// supervisor must not be running already
	// adding can be done in the init phase only

	// add the module
	ws.oracles = append(ws.oracles, ora)
}

// OracleStarted signals to supervisor about new running oracle.
func (ws *WatchdogSupervisor) OracleStarted() {
	ws.waitGroup.Add(1)
}

// OracleDone signals to supervisor about finished oracle.
func (ws *WatchdogSupervisor) OracleDone() {
	ws.waitGroup.Done()
}

// Run inits the supervisor modules and starts them to do their job.
func (ws *WatchdogSupervisor) Log() logger.Logger {
	return ws.log
}

// Terminate signals oracle supervisor to stop all running modules
// and finish their job.
func (ws *WatchdogSupervisor) Terminate() {
	// log the process
	ws.log.Notice("supervisor is closing")

	// signal all the oracles to close
	for _, ora := range ws.oracles {
		ora.Terminate()
	}

	// wait for all the oracles to signal termination
	ws.waitGroup.Wait()

	// signal myself to terminate
	ws.sigClose <- true
}

// Run inits the supervisor modules and starts them to do their job.
func (ws *WatchdogSupervisor) Run() {
	// any modules to run?
	if 0 == len(ws.oracles) {
		ws.log.Error("no oracles defined, nothing to do")
		return
	}

	defer func() {
		// log the process
		ws.log.Notice("supervisor closed")
	}()

	// start all the registered oracles
	for _, ora := range ws.oracles {
		ora.Run()
	}

	// wait patiently for the terminate signal
	<-ws.sigClose
}
