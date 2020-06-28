// Package pricefeed implements Supervisor oracle module for feeding conversion rate
// into a price oracle in the blockchain.
package pricefeed

import (
	"oracle-watchdog/internal/config"
	"oracle-watchdog/internal/supervisor"
)

// PriceOracle defines an oracle module for feeding
// conversion rate between defined symbol pairs into
// a pre-configured price oracle.
type PriceOracle struct {
	cfg *config.ModuleConfig
	sup supervisor.Supervisor
}

// New creates a new instance of the price feed oracle module.
func New(cfg *config.ModuleConfig, sup supervisor.Supervisor) (supervisor.Oracle, error) {
	// make the ballot oracle
	pf := PriceOracle{
		cfg: cfg,
		sup: sup,
	}

	// make sure to add this oracle to the supervisor before returning
	sup.AddOracle(&pf)
	return &pf, nil
}
