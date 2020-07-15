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
	"oracle-watchdog/internal/modules/utils"
	"oracle-watchdog/internal/supervisor"
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
	decimals     *big.Float
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

	// get the current price and round so we know how to proceed
	pro.init()

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
	price, stamp, err := pro.pullPrice()
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
		go pro.writePrice(price, stamp)
	}
}

// pullPrice pulls a new price from the Binance server.
func (pro *PriceOracle) pullPrice() (float64, int64, error) {
	// prep a http request to the Binance API
	req, err := http.NewRequest(http.MethodGet, pro.cfg.ApiUrl, nil)
	if err != nil {
		pro.sup.Log().Criticalf("can not create http API request; %s", err.Error())
		return 0, 0, err
	}

	// set headers
	req.Header.Set("User-Agent", "Fantom-Backend-Server v1.0")

	// send request
	res, err := pro.http.Do(req)
	if err != nil {
		pro.sup.Log().Errorf("http API request failed; %s", err.Error())
		return 0, 0, err
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
		return 0, 0, err
	}

	return pro.decodePrice(body)
}

// decodePrice decodes remote price API response to get the actual price.
func (pro *PriceOracle) decodePrice(res []byte) (float64, int64, error) {
	// prep price response and error containers
	var err error
	var data struct {
		Symbol   string  `json:"symbol"`
		PriceStr string  `json:"price"`
		Price    float64 `json:"-"`
	}

	// try to decode the data
	if err := json.Unmarshal(res, &data); err != nil {
		pro.sup.Log().Errorf("can not decode API call response; %s", err.Error())
		return 0, 0, err
	}

	// decode price from string
	data.Price, err = strconv.ParseFloat(data.PriceStr, 64)
	if err != nil {
		pro.sup.Log().Errorf("can not parse price from API call; %s", err.Error())
		return 0, 0, err
	}

	// log and return the price
	return data.Price, time.Now().UTC().Unix(), nil
}

// writePrice sends the new price into the on-chain Oracle smart contract.
func (pro *PriceOracle) writePrice(price float64, stamp int64) {
	// log action
	pro.sup.Log().Debugf("updating %s on-chain price to %0.8f", pro.cfg.Symbol, price)

	// prep Eth client
	eth := ethclient.NewClient(pro.sup.Lachesis())

	// connect the contract
	contract, err := NewPriceFeedContract(pro.cfg.PriceAggregate, eth)
	if err != nil {
		pro.sup.Log().Errorf("can not access pricefeed contract; %s", err.Error())
		return
	}

	// prep the transactor
	sig, err := utils.Transactor(pro.sup.Log(), &pro.cfg.KeyStore, &pro.cfg.KeySecret)
	if err != nil {
		pro.sup.Log().Errorf("can not interact with the pricefeed contract; %s", err.Error())
		return
	}

	// send the price multiplied to the destination decimals
	val, _ := new(big.Float).Mul(big.NewFloat(price), pro.decimals).Int(nil)
	err = pro.updateAnswer(contract, sig, val, new(big.Int).SetInt64(stamp))
	if err != nil {
		pro.sup.Log().Errorf("can not push new price to the contract; %s", err.Error())
		return
	}
}

// updateAnswer updates the answer inside the price feed aggregate contract.
func (pro *PriceOracle) updateAnswer(contract *PriceFeedContract, sig *bind.TransactOpts, val *big.Int, stamp *big.Int) error {
	// info about the price update TX
	pro.sup.Log().Debugf("updating price feed %s from address %s", pro.cfg.Symbol, sig.From.String())

	// pull the current round from the contract
	round, err := contract.LatestRound(nil)
	if err != nil {
		pro.sup.Log().Errorf("can not check the latest round; %s", err.Error())
		return err
	}

	// step up to the next round
	round = new(big.Int).Add(round, big.NewInt(1))

	// do the actual update call
	tx, err := contract.UpdateAnswer(sig, round, val, stamp)
	if err != nil {
		pro.sup.Log().Errorf("price update transaction failed; %s", err.Error())
		return err
	}

	// info about the price update TX
	pro.sup.Log().Infof("price of %s updated by trx %s", pro.cfg.Symbol, tx.Hash().String())
	return nil
}

// init initializes the price feed internal status to enable it's operations.
// it loads current price value from the contract to know where to start.
func (pro *PriceOracle) init() {
	// prep Eth client
	eth := ethclient.NewClient(pro.sup.Lachesis())

	// connect the contract
	contract, err := NewPriceFeedContract(pro.cfg.PriceAggregate, eth)
	if err != nil {
		pro.sup.Log().Errorf("can not access pricefeed contract; %s", err.Error())
		return
	}

	// get the latest answer
	cp, err := contract.LatestAnswer(nil)
	if err != nil {
		pro.sup.Log().Errorf("can not load the latest price; %s", err.Error())
		return
	}

	// prep the calculation elements
	latest := new(big.Float).SetInt(cp)
	pro.decimals = new(big.Float).SetFloat64(math.Pow10(int(pro.cfg.Decimals)))

	// make sure we have some recovery if the price calculation panics
	defer func() {
		if r := recover(); r != nil {
			pro.sup.Log().Errorf("invalid data received for price calculation")
			pro.currentPrice = 0.0
		}
	}()

	// calculate the current price
	pr := new(big.Float).Quo(latest, pro.decimals)
	pro.currentPrice, _ = pr.Float64()

	// log the status
	pro.sup.Log().Debugf("current price for %s is: %0.8f", pro.cfg.Symbol, pro.currentPrice)
}
