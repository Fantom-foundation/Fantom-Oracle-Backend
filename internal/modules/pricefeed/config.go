// Package pricefeed implements Supervisor oracle module for feeding conversion rate
// into a price oracle in the blockchain.
package pricefeed

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
)

// PriceOracleConfig represents a configuration
// of the Binance price feed oracle module.
type PriceOracleConfig struct {
	Name            string         `json:"name"`
	Symbol          string         `json:"symbol"`
	Decimals        uint32         `json:"decimals"`
	ApiUrl          string         `json:"api_point"`
	ApiKey          string         `json:"api_key"`
	ApiSecret       string         `json:"api_secret"`
	Token           common.Address `json:"token"`
	PriceAggregate  common.Address `json:"aggregate"`
	KeyStore        string         `json:"keystore"`
	KeySecret       string         `json:"key_secret"`
	PullDelayMs     int64          `json:"pull_delay_milliseconds"`
	WriteBarrierPct float64        `json:"write_barrier_pct"`
}

// configuration loads and parses the Binance price feed oracle module configuration
// from a config JSON file.
func configuration(path string) (*PriceOracleConfig, error) {
	// read the config file
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// parse the modules json structure
	cfg := PriceOracleConfig{}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
