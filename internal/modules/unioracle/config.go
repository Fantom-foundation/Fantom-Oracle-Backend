// Package unioracle handles automated updates of the Uniswap Price Oracle
// update procedure.
package unioracle

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
)

// PriceOracleConfig represents a configuration
// of the Binance price feed oracle module.
type UniswapOracleConfig struct {
	Factory   common.Address `json:"factory"`
	Manager   common.Address `json:"manager"`
	KeyStore  string         `json:"keystore"`
	KeySecret string         `json:"key_secret"`
	PullDelay int64          `json:"pull_delay"`
}

// configuration loads and parses the Uniswap oracle automated updater.
func configuration(path string) (*UniswapOracleConfig, error) {
	// read the config file
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// parse the modules json structure
	cfg := UniswapOracleConfig{}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
