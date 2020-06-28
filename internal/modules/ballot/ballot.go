// Package ballot handles closed ballot finalization.
// It implements an oracle which uses Fantom API to collect
// ballot participants' account total and feeds it into
// the smart contract to allow ballot winner calculation.
package ballot

import (
	"oracle-watchdog/internal/config"
	"oracle-watchdog/internal/supervisor"
)

// FinalizingOracle defines an oracle module for feeding
// ballot participants' account totals and finalizing
// closed ballots to allow winner calculations.
type FinalizingOracle struct {
	cfg *config.ModuleConfig
	sup supervisor.Supervisor
}

// New creates a new instance of the ballot oracle module.
func New(cfg *config.ModuleConfig, sup supervisor.Supervisor) (supervisor.Oracle, error) {
	// make the ballot finalizing oracle
	bo := FinalizingOracle{
		cfg: cfg,
		sup: sup,
	}

	// make sure to add this oracle to the supervisor
	sup.AddOracle(&bo)
	return &bo, nil
}
