// Package supervisor manages running modules and takes care
// of the inter-modules communication, if required.
package supervisor

import (
	"github.com/ethereum/go-ethereum/rpc"
	"oracle-watchdog/internal/logger"
)

// Supervisor defines the interface for a process managing
// set of oracle modules with specific functionality.
type Supervisor interface {
	// AddModule adds new module to the supervisor.
	AddOracle(Oracle)

	// OracleStarted signals to supervisor about new running oracle.
	OracleStarted()

	// OracleDone signals to supervisor about finished oracle.
	OracleDone()

	// Run starts the supervisor functions.
	Run()

	// Terminate signals supervisor to stop all running modules
	// and finish the job.
	Terminate()

	// Log returns an instance of a logger to be used for log messages
	// by the oracles.
	Log() logger.Logger

	// Lachesis returns an established and valid lachesis node connection
	// if possible, or nil if the connection is not available.
	Lachesis() *rpc.Client
}

// Oracle defines the interface for an oracle module supervised
// by the supervisor.
type Oracle interface {
	// Run starts the oracle functions.
	Run()

	// Terminate signals oracle to terminate it's process.
	Terminate()
}
