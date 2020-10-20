// Package main implements the entry point into the Watchdog app.
package main

import (
	"encoding/json"
	"io/ioutil"
	"oracle-watchdog/internal/config"
	"oracle-watchdog/internal/modules/ballot"
	"oracle-watchdog/internal/modules/pricefeed"
	"oracle-watchdog/internal/modules/unioracle"
	"oracle-watchdog/internal/supervisor"
)

// loadOracles loads configured oracle modules into the supervisor.
func loadOracles(cfg *config.Cfg, sup supervisor.Supervisor) error {
	// load oracle modules catalog
	cat, err := catalog(cfg.ModulesConfigPath)
	if err != nil {
		sup.Log().Error("modules catalog not available")
		return err
	}

	// loop the catalog modules and load their oracles into the supervisor
	for _, mc := range cat.Modules {
		var err error

		// decide what to load based on oracle type
		// specified in the config file
		switch mc.Type {
		case "ballot":
			_, err = ballot.New(&mc, sup)
		case "pricefeed":
			_, err = pricefeed.New(&mc, sup)
		case "unioracle":
			_, err = unioracle.New(&mc, sup)
		}

		// is the module loaded?
		if err != nil {
			return err
		}
	}

	return nil
}

// catalog load the module catalog from expected modules
// catalog file within the configured path.
func catalog(ctPath string) (*config.ModuleCatalog, error) {
	// read the config file
	file, err := ioutil.ReadFile(ctPath)
	if err != nil {
		return nil, err
	}

	// parse the modules json structure
	cat := config.ModuleCatalog{}
	err = json.Unmarshal(file, &cat)
	if err != nil {
		return nil, err
	}

	return &cat, nil
}
