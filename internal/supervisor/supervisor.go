// Package supervisor manages running modules and takes care
// of the inter-modules communication, if required.
package supervisor

import (
	"oracle-watchdog/internal/logger"
)

// WatchdogSupervisor implements the
type WatchdogSupervisor struct {
	log     logger.Logger
	oracles []Oracle
}

// New creates new instance of the Supervisor.
func New(log logger.Logger) Supervisor {
	// make the watchdog
	ws := WatchdogSupervisor{
		log:     log,
		oracles: make([]Oracle, 0),
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

// Run inits the supervisor modules and starts them to do their job.
func (ws *WatchdogSupervisor) Log() logger.Logger {
	return ws.log
}

// Run inits the supervisor modules and starts them to do their job.
func (ws *WatchdogSupervisor) Run() {

}

// Terminate signals watchdog supervisor to stop all running modules
// and finish their job.
func (ws *WatchdogSupervisor) Terminate() {

}
