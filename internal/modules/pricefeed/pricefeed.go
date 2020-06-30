// Package pricefeed implements Supervisor oracle module for feeding conversion rate
// into a price oracle in the blockchain.
package pricefeed

//go:generate abigen --abi ./contract/oracle.abi --pkg pricefeed --type PriceFeedContract --out ./bridge.go

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"oracle-watchdog/internal/config"
	"oracle-watchdog/internal/supervisor"
	"os"
	"strconv"
	"time"
)

// What the Binance price feed oracle does:
//  1. downloads current symbol price from Binance API ticker
//  2. compares the new price with the previous value
//  3. writes the new value to PriceOracle contract,
// 	   if the new value differs from the previous for more than configured %
//  4. waits for certain amount of milliseconds (Binance API has call limits)
//  5. repeats the process from #1

// httpClientTimeout represents a max time we tolerate for Binance API request.
const httpClientTimeout = time.Second * 5

// PriceOracle defines an oracle module for feeding
// conversion rate between defined symbol pairs into
// a pre-configured price oracle.
type PriceOracle struct {
	cfg          *PriceOracleConfig
	sup          supervisor.Supervisor
	http         http.Client
	sigClose     chan bool
	symbol       [32]byte
	currentPrice float64
}

// New creates a new instance of the price feed oracle module.
func New(cfg *config.ModuleConfig, sup supervisor.Supervisor) (supervisor.Oracle, error) {
	// read the module configuration
	cf, err := configuration(cfg.CfgPath)
	if err != nil {
		sup.Log().Criticalf("can not load oracle configuration %s; %s", cfg.CfgPath, err.Error())
		return nil, err
	}

	// prep fixed size symbol name for contract calls
	var symbol [32]byte
	copy(symbol[:], cf.Symbol)

	// make the ballot oracle
	pf := PriceOracle{
		cfg:      cf,
		sup:      sup,
		symbol:   symbol,
		http:     http.Client{Timeout: httpClientTimeout},
		sigClose: make(chan bool, 1),
	}

	// make sure to add this oracle to the supervisor before returning
	sup.AddOracle(&pf)
	return &pf, nil
}

// Terminate signals the price feed oracle to stop the process and close.
func (pro *PriceOracle) Terminate() {
	// signal the oracle main thread to terminate
	pro.sigClose <- true
}

// Run starts the price feed oracle business.
func (pro *PriceOracle) Run() {
	// log we are done
	pro.sup.Log().Noticef("starting %s oracle", pro.cfg.Name)

	// signal supervisor we are good to go
	pro.sup.OracleStarted()
	go pro.feedPrice()
}

// feedPrice does the main oracle job by pulling the current symbol price
// from Binance API and feeding it into the on-chain Oracle contract
// on configured criteria.
func (pro *PriceOracle) feedPrice() {
	// signal the oracle has ended
	defer func() {
		// log we are done
		pro.sup.Log().Noticef("oracle %s terminated", pro.cfg.Name)

		// signal supervisor we are done here
		pro.sup.OracleDone()
	}()

	// delay represents the Binance API price pulling delay; check Binance API limits
	// @see https://api.binance.com/api/v1/exchangeInfo
	var delay = time.Duration(pro.cfg.PullDelayMs) * time.Millisecond

	// loop the function until terminated
	for {
		// update the price
		pro.checkPrice()

		// wait for termination or delay
		select {
		case <-pro.sigClose:
			// stop signal received
			return
		case <-time.After(delay):
			// we repeat the function
		}
	}
}

// checkPrice pulls a new price for the target symbol in relation
// to the source symbol using Binance API and based on configured
// write wall criteria.
func (pro *PriceOracle) checkPrice() {
	// pull the new price from Binance API
	price, err := pro.pullPrice()
	if err != nil {
		pro.sup.Log().Errorf("can not pull a new price for %s; %s", pro.cfg.Name, err.Error())
		return
	}

	// compare the price with the previous one by calculating delta percentage
	// we always use positive delta percentage since we only need an absolute value
	var pct = 100.0
	if pro.currentPrice != 0 {
		pct = math.Abs(price-pro.currentPrice) / pro.currentPrice
	}

	// is the delta over the barrier
	if pct >= pro.cfg.WriteBarrierPct {
		// we are on a different price from now on
		pro.currentPrice = price

		// write the price to the backend contract, use a separate thread to do so
		go pro.writePrice(price)
	}
}

// pullPrice pulls a new price from the Binance server.
func (pro *PriceOracle) pullPrice() (float64, error) {
	// prep a http request to the Binance API
	req, err := http.NewRequest(http.MethodGet, pro.cfg.ApiUrl, nil)
	if err != nil {
		pro.sup.Log().Criticalf("can not create http API request; %s", err.Error())
		return 0, err
	}

	// set headers
	req.Header.Set("User-Agent", "Fantom-Backend-Server v1.0")

	// send request
	res, err := pro.http.Do(req)
	if err != nil {
		pro.sup.Log().Errorf("http API request failed; %s", err.Error())
		return 0, err
	}

	// make sure to close the body reader when we are done
	defer func() {
		// no body reader to close
		if res.Body == nil {
			return
		}

		// try to cloe the body reader; handle error gracefully
		err := res.Body.Close()
		if err != nil {
			pro.sup.Log().Errorf("error closing http request body reader; %s", err.Error())
		}
	}()

	// read the request content
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		pro.sup.Log().Errorf("can not read http response body; %s", err.Error())
		return 0, err
	}

	// prep price response container
	var data struct {
		Symbol   string  `json:"symbol"`
		PriceStr string  `json:"price"`
		Price    float64 `json:"-"`
	}

	// try to decode the data
	if err := json.Unmarshal(body, &data); err != nil {
		pro.sup.Log().Errorf("can not decode API call response; %s", err.Error())
		return 0, err
	}

	// decode price from string
	data.Price, err = strconv.ParseFloat(data.PriceStr, 64)
	if err != nil {
		pro.sup.Log().Errorf("can not parse price from API call; %s", err.Error())
		return 0, err
	}

	// log and return the price
	pro.sup.Log().Debugf("current %s price is %0.18f", data.Symbol, data.Price)
	return data.Price, nil
}

// transactor creates an authorized transactor for signed contract interactions.
// The transactor uses local key storage and configured key secret for unlocking the storage.
func (pro *PriceOracle) transactor() (*bind.TransactOpts, error) {
	// read the key store
	f, err := os.Open(pro.cfg.KeyStore)
	if err != nil {
		pro.sup.Log().Errorf("can not open key store; %s", err.Error())
		return nil, err
	}

	// ensure proper cleanup
	defer func() {
		// make sure to close the opened key store file
		if err := f.Close(); err != nil {
			pro.sup.Log().Errorf("error closing key store; %s", err.Error())
		}
	}()

	// create the transactor
	tr, err := bind.NewTransactor(f, pro.cfg.KeySecret)
	if err != nil {
		pro.sup.Log().Errorf("can not create authorized signing transactor; %s", err.Error())
		return nil, err
	}

	return tr, nil
}

// writePrice sends the new price into the on-chain Oracle smart contract.
func (pro *PriceOracle) writePrice(price float64) {
	// log action
	pro.sup.Log().Debugf("updating %s on-chain price to %0.18f", pro.cfg.Symbol, price)

	// calculate the target value
	val, _ := new(big.Float).Mul(big.NewFloat(price), big.NewFloat(math.Pow10(18))).Int(nil)
	if val.IsUint64() {
		pro.sup.Log().Debugf("using value %d", val.Uint64())
	}

	// prep Eth client
	eth := ethclient.NewClient(pro.sup.Lachesis())

	// connect the contract
	contract, err := NewPriceFeedContract(pro.cfg.Contract, eth)
	if err != nil {
		pro.sup.Log().Errorf("can not access pricefeed contract; %s", err.Error())
		return
	}

	// prep the transactor
	sig, err := pro.transactor()
	if err != nil {
		pro.sup.Log().Errorf("can not interact with the pricefeed contract; %s", err.Error())
		return
	}

	// send the price
	tx, err := contract.SetPrice(sig, pro.symbol, val)
	if err != nil {
		pro.sup.Log().Errorf("can not push new price to the contract; %s", err.Error())
		return
	}

	// info about the price update TX
	pro.sup.Log().Errorf("price of %s updated by tx %s", pro.cfg.Symbol, tx.Hash().String())
}
