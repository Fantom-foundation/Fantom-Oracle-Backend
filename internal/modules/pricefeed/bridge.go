// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pricefeed

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PriceFeedContractABI is the input ABI used to generate the binding from.
const PriceFeedContractABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"feeds\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"PriceExpirationPeriodChanged\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addSource\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"changeExpirationPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"dropSource\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priceExpirationPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updated\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"symbol\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sources\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PriceFeedContract is an auto generated Go binding around an Ethereum contract.
type PriceFeedContract struct {
	PriceFeedContractCaller     // Read-only binding to the contract
	PriceFeedContractTransactor // Write-only binding to the contract
	PriceFeedContractFilterer   // Log filterer for contract events
}

// PriceFeedContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriceFeedContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriceFeedContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriceFeedContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriceFeedContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriceFeedContractSession struct {
	Contract     *PriceFeedContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PriceFeedContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriceFeedContractCallerSession struct {
	Contract *PriceFeedContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PriceFeedContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriceFeedContractTransactorSession struct {
	Contract     *PriceFeedContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PriceFeedContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriceFeedContractRaw struct {
	Contract *PriceFeedContract // Generic contract binding to access the raw methods on
}

// PriceFeedContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriceFeedContractCallerRaw struct {
	Contract *PriceFeedContractCaller // Generic read-only contract binding to access the raw methods on
}

// PriceFeedContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriceFeedContractTransactorRaw struct {
	Contract *PriceFeedContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriceFeedContract creates a new instance of PriceFeedContract, bound to a specific deployed contract.
func NewPriceFeedContract(address common.Address, backend bind.ContractBackend) (*PriceFeedContract, error) {
	contract, err := bindPriceFeedContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContract{PriceFeedContractCaller: PriceFeedContractCaller{contract: contract}, PriceFeedContractTransactor: PriceFeedContractTransactor{contract: contract}, PriceFeedContractFilterer: PriceFeedContractFilterer{contract: contract}}, nil
}

// NewPriceFeedContractCaller creates a new read-only instance of PriceFeedContract, bound to a specific deployed contract.
func NewPriceFeedContractCaller(address common.Address, caller bind.ContractCaller) (*PriceFeedContractCaller, error) {
	contract, err := bindPriceFeedContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractCaller{contract: contract}, nil
}

// NewPriceFeedContractTransactor creates a new write-only instance of PriceFeedContract, bound to a specific deployed contract.
func NewPriceFeedContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PriceFeedContractTransactor, error) {
	contract, err := bindPriceFeedContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractTransactor{contract: contract}, nil
}

// NewPriceFeedContractFilterer creates a new log filterer instance of PriceFeedContract, bound to a specific deployed contract.
func NewPriceFeedContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PriceFeedContractFilterer, error) {
	contract, err := bindPriceFeedContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractFilterer{contract: contract}, nil
}

// bindPriceFeedContract binds a generic wrapper to an already deployed contract.
func bindPriceFeedContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriceFeedContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedContract *PriceFeedContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceFeedContract.Contract.PriceFeedContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedContract *PriceFeedContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.PriceFeedContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedContract *PriceFeedContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.PriceFeedContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriceFeedContract *PriceFeedContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriceFeedContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriceFeedContract *PriceFeedContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriceFeedContract *PriceFeedContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.contract.Transact(opts, method, params...)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 symbol) view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCaller) GetPrice(opts *bind.CallOpts, symbol [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceFeedContract.contract.Call(opts, out, "getPrice", symbol)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 symbol) view returns(uint256)
func (_PriceFeedContract *PriceFeedContractSession) GetPrice(symbol [32]byte) (*big.Int, error) {
	return _PriceFeedContract.Contract.GetPrice(&_PriceFeedContract.CallOpts, symbol)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 symbol) view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCallerSession) GetPrice(symbol [32]byte) (*big.Int, error) {
	return _PriceFeedContract.Contract.GetPrice(&_PriceFeedContract.CallOpts, symbol)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeedContract *PriceFeedContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PriceFeedContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeedContract *PriceFeedContractSession) Owner() (common.Address, error) {
	return _PriceFeedContract.Contract.Owner(&_PriceFeedContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeedContract *PriceFeedContractCallerSession) Owner() (common.Address, error) {
	return _PriceFeedContract.Contract.Owner(&_PriceFeedContract.CallOpts)
}

// PriceExpirationPeriod is a free data retrieval call binding the contract method 0x125d088b.
//
// Solidity: function priceExpirationPeriod() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCaller) PriceExpirationPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriceFeedContract.contract.Call(opts, out, "priceExpirationPeriod")
	return *ret0, err
}

// PriceExpirationPeriod is a free data retrieval call binding the contract method 0x125d088b.
//
// Solidity: function priceExpirationPeriod() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractSession) PriceExpirationPeriod() (*big.Int, error) {
	return _PriceFeedContract.Contract.PriceExpirationPeriod(&_PriceFeedContract.CallOpts)
}

// PriceExpirationPeriod is a free data retrieval call binding the contract method 0x125d088b.
//
// Solidity: function priceExpirationPeriod() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCallerSession) PriceExpirationPeriod() (*big.Int, error) {
	return _PriceFeedContract.Contract.PriceExpirationPeriod(&_PriceFeedContract.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) view returns(uint256 price, uint256 updated)
func (_PriceFeedContract *PriceFeedContractCaller) Prices(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Price   *big.Int
	Updated *big.Int
}, error) {
	ret := new(struct {
		Price   *big.Int
		Updated *big.Int
	})
	out := ret
	err := _PriceFeedContract.contract.Call(opts, out, "prices", arg0)
	return *ret, err
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) view returns(uint256 price, uint256 updated)
func (_PriceFeedContract *PriceFeedContractSession) Prices(arg0 [32]byte) (struct {
	Price   *big.Int
	Updated *big.Int
}, error) {
	return _PriceFeedContract.Contract.Prices(&_PriceFeedContract.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0x60846bc6.
//
// Solidity: function prices(bytes32 ) view returns(uint256 price, uint256 updated)
func (_PriceFeedContract *PriceFeedContractCallerSession) Prices(arg0 [32]byte) (struct {
	Price   *big.Int
	Updated *big.Int
}, error) {
	return _PriceFeedContract.Contract.Prices(&_PriceFeedContract.CallOpts, arg0)
}

// Sources is a free data retrieval call binding the contract method 0xb750bdde.
//
// Solidity: function sources(address ) view returns(bool)
func (_PriceFeedContract *PriceFeedContractCaller) Sources(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PriceFeedContract.contract.Call(opts, out, "sources", arg0)
	return *ret0, err
}

// Sources is a free data retrieval call binding the contract method 0xb750bdde.
//
// Solidity: function sources(address ) view returns(bool)
func (_PriceFeedContract *PriceFeedContractSession) Sources(arg0 common.Address) (bool, error) {
	return _PriceFeedContract.Contract.Sources(&_PriceFeedContract.CallOpts, arg0)
}

// Sources is a free data retrieval call binding the contract method 0xb750bdde.
//
// Solidity: function sources(address ) view returns(bool)
func (_PriceFeedContract *PriceFeedContractCallerSession) Sources(arg0 common.Address) (bool, error) {
	return _PriceFeedContract.Contract.Sources(&_PriceFeedContract.CallOpts, arg0)
}

// AddSource is a paid mutator transaction binding the contract method 0x2a142b0b.
//
// Solidity: function addSource(address addr) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) AddSource(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "addSource", addr)
}

// AddSource is a paid mutator transaction binding the contract method 0x2a142b0b.
//
// Solidity: function addSource(address addr) returns()
func (_PriceFeedContract *PriceFeedContractSession) AddSource(addr common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.AddSource(&_PriceFeedContract.TransactOpts, addr)
}

// AddSource is a paid mutator transaction binding the contract method 0x2a142b0b.
//
// Solidity: function addSource(address addr) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) AddSource(addr common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.AddSource(&_PriceFeedContract.TransactOpts, addr)
}

// ChangeExpirationPeriod is a paid mutator transaction binding the contract method 0x8ef314c6.
//
// Solidity: function changeExpirationPeriod(uint256 expiration) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) ChangeExpirationPeriod(opts *bind.TransactOpts, expiration *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "changeExpirationPeriod", expiration)
}

// ChangeExpirationPeriod is a paid mutator transaction binding the contract method 0x8ef314c6.
//
// Solidity: function changeExpirationPeriod(uint256 expiration) returns()
func (_PriceFeedContract *PriceFeedContractSession) ChangeExpirationPeriod(expiration *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.ChangeExpirationPeriod(&_PriceFeedContract.TransactOpts, expiration)
}

// ChangeExpirationPeriod is a paid mutator transaction binding the contract method 0x8ef314c6.
//
// Solidity: function changeExpirationPeriod(uint256 expiration) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) ChangeExpirationPeriod(expiration *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.ChangeExpirationPeriod(&_PriceFeedContract.TransactOpts, expiration)
}

// DropSource is a paid mutator transaction binding the contract method 0xb98ffeb1.
//
// Solidity: function dropSource(address addr) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) DropSource(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "dropSource", addr)
}

// DropSource is a paid mutator transaction binding the contract method 0xb98ffeb1.
//
// Solidity: function dropSource(address addr) returns()
func (_PriceFeedContract *PriceFeedContractSession) DropSource(addr common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.DropSource(&_PriceFeedContract.TransactOpts, addr)
}

// DropSource is a paid mutator transaction binding the contract method 0xb98ffeb1.
//
// Solidity: function dropSource(address addr) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) DropSource(addr common.Address) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.DropSource(&_PriceFeedContract.TransactOpts, addr)
}

// SetPrice is a paid mutator transaction binding the contract method 0x10d8d74d.
//
// Solidity: function setPrice(bytes32 symbol, uint256 newPrice) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) SetPrice(opts *bind.TransactOpts, symbol [32]byte, newPrice *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "setPrice", symbol, newPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0x10d8d74d.
//
// Solidity: function setPrice(bytes32 symbol, uint256 newPrice) returns()
func (_PriceFeedContract *PriceFeedContractSession) SetPrice(symbol [32]byte, newPrice *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.SetPrice(&_PriceFeedContract.TransactOpts, symbol, newPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0x10d8d74d.
//
// Solidity: function setPrice(bytes32 symbol, uint256 newPrice) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) SetPrice(symbol [32]byte, newPrice *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.SetPrice(&_PriceFeedContract.TransactOpts, symbol, newPrice)
}

// PriceFeedContractPriceChangedIterator is returned from FilterPriceChanged and is used to iterate over the raw logs and unpacked data for PriceChanged events raised by the PriceFeedContract contract.
type PriceFeedContractPriceChangedIterator struct {
	Event *PriceFeedContractPriceChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceFeedContractPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractPriceChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedContractPriceChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedContractPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractPriceChanged represents a PriceChanged event raised by the PriceFeedContract contract.
type PriceFeedContractPriceChanged struct {
	Symbol [32]byte
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPriceChanged is a free log retrieval operation binding the contract event 0x6007aeeeb4a07f99c2208537b6cf8bf3710752ce9e7d85b7ab294174ba801545.
//
// Solidity: event PriceChanged(bytes32 indexed symbol, uint256 price)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterPriceChanged(opts *bind.FilterOpts, symbol [][32]byte) (*PriceFeedContractPriceChangedIterator, error) {

	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "PriceChanged", symbolRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractPriceChangedIterator{contract: _PriceFeedContract.contract, event: "PriceChanged", logs: logs, sub: sub}, nil
}

// WatchPriceChanged is a free log subscription operation binding the contract event 0x6007aeeeb4a07f99c2208537b6cf8bf3710752ce9e7d85b7ab294174ba801545.
//
// Solidity: event PriceChanged(bytes32 indexed symbol, uint256 price)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchPriceChanged(opts *bind.WatchOpts, sink chan<- *PriceFeedContractPriceChanged, symbol [][32]byte) (event.Subscription, error) {

	var symbolRule []interface{}
	for _, symbolItem := range symbol {
		symbolRule = append(symbolRule, symbolItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "PriceChanged", symbolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractPriceChanged)
				if err := _PriceFeedContract.contract.UnpackLog(event, "PriceChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceChanged is a log parse operation binding the contract event 0x6007aeeeb4a07f99c2208537b6cf8bf3710752ce9e7d85b7ab294174ba801545.
//
// Solidity: event PriceChanged(bytes32 indexed symbol, uint256 price)
func (_PriceFeedContract *PriceFeedContractFilterer) ParsePriceChanged(log types.Log) (*PriceFeedContractPriceChanged, error) {
	event := new(PriceFeedContractPriceChanged)
	if err := _PriceFeedContract.contract.UnpackLog(event, "PriceChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceFeedContractPriceExpirationPeriodChangedIterator is returned from FilterPriceExpirationPeriodChanged and is used to iterate over the raw logs and unpacked data for PriceExpirationPeriodChanged events raised by the PriceFeedContract contract.
type PriceFeedContractPriceExpirationPeriodChangedIterator struct {
	Event *PriceFeedContractPriceExpirationPeriodChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PriceFeedContractPriceExpirationPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractPriceExpirationPeriodChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PriceFeedContractPriceExpirationPeriodChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PriceFeedContractPriceExpirationPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractPriceExpirationPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractPriceExpirationPeriodChanged represents a PriceExpirationPeriodChanged event raised by the PriceFeedContract contract.
type PriceFeedContractPriceExpirationPeriodChanged struct {
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceExpirationPeriodChanged is a free log retrieval operation binding the contract event 0x17f1771bfbf65f2f9f6890ae150425ce55886916bdfe0bed0e64ad65296759dd.
//
// Solidity: event PriceExpirationPeriodChanged(uint256 newPeriod)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterPriceExpirationPeriodChanged(opts *bind.FilterOpts) (*PriceFeedContractPriceExpirationPeriodChangedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "PriceExpirationPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractPriceExpirationPeriodChangedIterator{contract: _PriceFeedContract.contract, event: "PriceExpirationPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchPriceExpirationPeriodChanged is a free log subscription operation binding the contract event 0x17f1771bfbf65f2f9f6890ae150425ce55886916bdfe0bed0e64ad65296759dd.
//
// Solidity: event PriceExpirationPeriodChanged(uint256 newPeriod)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchPriceExpirationPeriodChanged(opts *bind.WatchOpts, sink chan<- *PriceFeedContractPriceExpirationPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "PriceExpirationPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractPriceExpirationPeriodChanged)
				if err := _PriceFeedContract.contract.UnpackLog(event, "PriceExpirationPeriodChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceExpirationPeriodChanged is a log parse operation binding the contract event 0x17f1771bfbf65f2f9f6890ae150425ce55886916bdfe0bed0e64ad65296759dd.
//
// Solidity: event PriceExpirationPeriodChanged(uint256 newPeriod)
func (_PriceFeedContract *PriceFeedContractFilterer) ParsePriceExpirationPeriodChanged(log types.Log) (*PriceFeedContractPriceExpirationPeriodChanged, error) {
	event := new(PriceFeedContractPriceExpirationPeriodChanged)
	if err := _PriceFeedContract.contract.UnpackLog(event, "PriceExpirationPeriodChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
