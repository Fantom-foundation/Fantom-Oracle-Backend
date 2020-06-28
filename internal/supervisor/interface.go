// Package supervisor manages running modules and takes care
// of the inter-modules communication, if required.
package supervisor

import "oracle-watchdog/internal/logger"

// Supervisor defines the interface for a process managing
// set of oracle modules with specific functionality.
type Supervisor interface {
	// AddModule adds new module to the supervisor.
	AddOracle(Oracle)

	// Run starts the supervisor functions.
	Run()

	// Terminate signals supervisor to stop all running modules
	// and finish the job.
	Terminate()

	// Log returns an instance of a logger to be used for log messages
	// by the oracles.
	Log() logger.Logger
}

// Oracle defines the interface for an oracle module supervised
// by the supervisor.
type Oracle interface {
}
