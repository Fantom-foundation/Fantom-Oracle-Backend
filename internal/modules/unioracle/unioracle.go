// Package unioracle handles automated updates of the Uniswap Price Oracle
// update procedure.
package unioracle

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"oracle-watchdog/internal/config"
	"oracle-watchdog/internal/modules/utils"
	"oracle-watchdog/internal/supervisor"
	"time"
)

//go:generate abigen --abi ./abi/UniswapOracle.abi --pkg unioracle --type UniswapOracleContract --out ./oracle_bridge.go
//go:generate abigen --abi ./abi/UniswapOracleFactory.abi --pkg unioracle --type UniswapOracleFactoryContract --out ./factory_bridge.go

// updateTimeoutSlip is the time added to the minimal last Uniswap pair update period
// so the update period has some slack to accommodate imprecise block chain time.
const updateTimeoutSlip int64 = 600

// UniswapOracle defines a module for managing
// Uniswap Oracle instance on chain.
type UniswapOracle struct {
	cfg      *UniswapOracleConfig
	sup      supervisor.Supervisor
	sigClose chan bool
}

// New creates a new instance of the price feed oracle module.
func New(cfg *config.ModuleConfig, sup supervisor.Supervisor) (supervisor.Oracle, error) {
	// read the module configuration
	cf, err := configuration(cfg.CfgPath)
	if err != nil {
		sup.Log().Criticalf("can not load Uniswap oracle configuration %s; %s", cfg.CfgPath, err.Error())
		return nil, err
	}

	// make the oracle
	pf := UniswapOracle{
		cfg:      cf,
		sup:      sup,
		sigClose: make(chan bool, 1),
	}

	// make sure to add this oracle to the supervisor before returning
	sup.AddOracle(&pf)
	return &pf, nil
}

// Terminate signals the price feed oracle to stop the process and close.
func (uo *UniswapOracle) Terminate() {
	// signal the oracle main thread to terminate
	uo.sigClose <- true
}

// Run starts the price feed oracle business.
func (uo *UniswapOracle) Run() {
	// log we are done
	uo.sup.Log().Noticef("starting Uniswap oracle for factory %s", uo.cfg.Factory.String())

	// signal supervisor we are good to go
	uo.sup.OracleStarted()
	go uo.monitorOracle()
}

// monitorOracle does the main oracle job by monitoring the factory
// contract and calling update function as needed.
func (uo *UniswapOracle) monitorOracle() {
	// signal the oracle has ended
	defer func() {
		// log we are done
		uo.sup.Log().Noticef("uniswap oracle terminated")

		// signal supervisor we are done here
		uo.sup.OracleDone()
	}()

	// delay represents the interval we scan the oracle for update need
	var delay = time.Duration(uo.cfg.PullDelay) * time.Second

	// loop the function until terminated
	for {
		// update the price
		uo.checkStatus()

		// wait for termination or delay
		select {
		case <-uo.sigClose:
			// stop signal received
			return
		case <-time.After(delay):
			// we repeat the function
		}
	}
}

// checkStatus pulls a list of oracles managed by the factory
// and checks the status of each oracle so it can be triggered
// as needed.
func (uo *UniswapOracle) checkStatus() {
	// log action
	uo.sup.Log().Debugf("checking Uniswap Oracles status on factory %s", uo.cfg.Factory.String())

	// pull list of all oracles managed by the factory
	list, err := uo.pullOracles()
	if err != nil {
		uo.sup.Log().Errorf("list of oracles not available; %s", err.Error())
		return
	}

	// log info
	uo.sup.Log().Debugf("found %d oracle(s) on factory %s", len(list), uo.cfg.Factory.String())

	// loop the list and check each oracle in it
	for _, ora := range list {
		// check this oracle
		err = uo.checkOracle(ora)
		if err != nil {
			uo.sup.Log().Errorf("oracles check failed; %s", err.Error())
		}
	}
}

// pullOracles pulls list of oracles managed by the
func (uo *UniswapOracle) pullOracles() ([]common.Address, error) {
	// log action
	uo.sup.Log().Debugf("loading oracles from %s", uo.cfg.Factory.String())

	// prep Eth client
	eth := ethclient.NewClient(uo.sup.Lachesis())

	// connect the contract
	contract, err := NewUniswapOracleFactoryContract(uo.cfg.Factory, eth)
	if err != nil {
		uo.sup.Log().Errorf("can not access oracle factory contract %s; %s",
			uo.cfg.Factory.String(), err.Error())
		return nil, err
	}

	return contract.Oracles(nil)
}

// checkOracle checks the oracle update timer
func (uo *UniswapOracle) checkOracle(ora common.Address) error {
	// log action
	uo.sup.Log().Debugf("checking oracle %s", ora.String())

	// prep Eth client
	eth := ethclient.NewClient(uo.sup.Lachesis())

	// connect the contract
	contract, err := NewUniswapOracleContract(ora, eth)
	if err != nil {
		uo.sup.Log().Errorf("can not access oracle contract %s; %s", ora.String(), err.Error())
		return err
	}

	// when was the last time oracle was updated
	lastUpdate, err := contract.BlockTimestampLast(uo.sup.ContractCallOpts(uo.cfg.Manager))
	if err != nil {
		uo.sup.Log().Errorf("last update check failed on %s; %s", ora.String(), err.Error())
		return err
	}

	// request period
	period, err := contract.PERIOD(nil)
	if err != nil {
		uo.sup.Log().Errorf("period request failed on %s; %s", ora.String(), err.Error())
		return err
	}

	// log info
	uo.sup.Log().Debugf("oracle %s updated at %s (%d)", ora.String(), time.Unix(int64(lastUpdate), 0).String(), lastUpdate)

	// do we need to update the oracle?
	if time.Now().UTC().Unix()-int64(lastUpdate) >= (period.Int64() + updateTimeoutSlip) {
		// inform about updating
		uo.sup.Log().Debugf("oracle %s is being updated now", ora.String())

		// do the update
		return uo.updateOracle(contract)
	}

	return nil
}

// updateOracle calls the update on the given Uniswap Oracle contract to actualize
// the Oracle price after the given period expired.
func (uo *UniswapOracle) updateOracle(contract *UniswapOracleContract) error {
	// get the transactor
	sig, err := utils.Transactor(uo.sup.Log(), &uo.cfg.KeyStore, &uo.cfg.KeySecret)
	if err != nil {
		uo.sup.Log().Errorf("can not interact with the Uniswap Oracle contract; %s", err.Error())
		return err
	}

	// call the update on contract
	tx, err := contract.Update(sig)
	if err != nil {
		uo.sup.Log().Errorf("Uniswap Oracle update failed; %s", err.Error())
		return err
	}

	// log success and return
	uo.sup.Log().Infof("Uniswap Oracle updated by %s", tx.Hash().String())
	return nil
}
