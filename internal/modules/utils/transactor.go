// Package utils implements utility functions for the executive modules
// of the oracle service.
package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"oracle-watchdog/internal/logger"
	"os"
)

// transactionsGasLimit is the limit of gas applied to a transaction call by default.
const transactionsGasLimit = 4000000

// transactor creates a new Transactor object for communicating
// with on-chain smart transactions via signed transactions.
func Transactor(log logger.Logger, keyStore *string, keySecret *string) (*bind.TransactOpts, error) {
	// read the key store
	f, err := os.Open(*keyStore)
	if err != nil {
		log.Errorf("can not open key store; %s", err.Error())
		return nil, err
	}

	// ensure proper cleanup
	defer func() {
		// make sure to close the opened key store file
		if err := f.Close(); err != nil {
			log.Errorf("error closing key store; %s", err.Error())
		}
	}()

	// create the transactor
	tr, err := bind.NewTransactor(f, *keySecret)
	if err != nil {
		log.Errorf("can not create authorized signing transactor; %s", err.Error())
		return nil, err
	}

	// assign some predefined constants
	tr.GasLimit = transactionsGasLimit

	// log transaction signer
	log.Debugf("using %s signing key", tr.From.String())

	// return the transactor
	return tr, nil
}
