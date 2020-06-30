// Package main implements the entry point into the Watchdog app.
package main

import (
	"oracle-watchdog/internal/config"
	"oracle-watchdog/internal/logger"
	"oracle-watchdog/internal/supervisor"
	"os"
	"os/signal"
)

// main initializes the Watchdog server and starts it when ready.
func main() {
	// read the configuration
	cfg := config.New()

	// setup the logger
	log := logger.New(&cfg)

	// setup the supervisor
	sup := supervisor.New(cfg.NodeRpc, log)
	err := loadOracles(&cfg, sup)
	if err != nil {
		log.Critical("can not load oracles into the supervisor")
		os.Exit(1)
	}

	// capture termination signals
	setupSignals(sup, log)

	// run the supervisor
	sup.Run()
}

// setupSignals creates a system signal listener and handles graceful termination upon receiving one.
func setupSignals(sup supervisor.Supervisor, log logger.Logger) {
	// we use a channel to capture interrupt and kill signal from OS
	ts := make(chan os.Signal, 2)
	signal.Notify(ts, os.Interrupt, os.Kill)

	// start monitoring
	go func() {
		// wait for the signal
		<-ts

		// log and signal supervisor to close the business
		log.Info("oracle is terminating")
		sup.Terminate()

		// we are done here
		log.Info("done")
		os.Exit(0)
	}()
}
