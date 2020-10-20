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
const PriceFeedContractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"SourceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"SourceDropped\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addSource\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"dropSource\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sources\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_value\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_stamp\",\"type\":\"uint256\"}],\"name\":\"updateAnswer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
func (_PriceFeedContract *PriceFeedContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_PriceFeedContract *PriceFeedContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 roundId) view returns(int256)
func (_PriceFeedContract *PriceFeedContractCaller) GetAnswer(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "getAnswer", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 roundId) view returns(int256)
func (_PriceFeedContract *PriceFeedContractSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _PriceFeedContract.Contract.GetAnswer(&_PriceFeedContract.CallOpts, roundId)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 roundId) view returns(int256)
func (_PriceFeedContract *PriceFeedContractCallerSession) GetAnswer(roundId *big.Int) (*big.Int, error) {
	return _PriceFeedContract.Contract.GetAnswer(&_PriceFeedContract.CallOpts, roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 roundId) view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCaller) GetTimestamp(opts *bind.CallOpts, roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "getTimestamp", roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 roundId) view returns(uint256)
func (_PriceFeedContract *PriceFeedContractSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _PriceFeedContract.Contract.GetTimestamp(&_PriceFeedContract.CallOpts, roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 roundId) view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCallerSession) GetTimestamp(roundId *big.Int) (*big.Int, error) {
	return _PriceFeedContract.Contract.GetTimestamp(&_PriceFeedContract.CallOpts, roundId)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_PriceFeedContract *PriceFeedContractCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_PriceFeedContract *PriceFeedContractSession) LatestAnswer() (*big.Int, error) {
	return _PriceFeedContract.Contract.LatestAnswer(&_PriceFeedContract.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_PriceFeedContract *PriceFeedContractCallerSession) LatestAnswer() (*big.Int, error) {
	return _PriceFeedContract.Contract.LatestAnswer(&_PriceFeedContract.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractSession) LatestRound() (*big.Int, error) {
	return _PriceFeedContract.Contract.LatestRound(&_PriceFeedContract.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCallerSession) LatestRound() (*big.Int, error) {
	return _PriceFeedContract.Contract.LatestRound(&_PriceFeedContract.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractSession) LatestTimestamp() (*big.Int, error) {
	return _PriceFeedContract.Contract.LatestTimestamp(&_PriceFeedContract.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_PriceFeedContract *PriceFeedContractCallerSession) LatestTimestamp() (*big.Int, error) {
	return _PriceFeedContract.Contract.LatestTimestamp(&_PriceFeedContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PriceFeedContract *PriceFeedContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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

// Sources is a free data retrieval call binding the contract method 0xb750bdde.
//
// Solidity: function sources(address ) view returns(bool)
func (_PriceFeedContract *PriceFeedContractCaller) Sources(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PriceFeedContract.contract.Call(opts, &out, "sources", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

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

// UpdateAnswer is a paid mutator transaction binding the contract method 0x0f7dbf0f.
//
// Solidity: function updateAnswer(uint256 _round, int256 _value, uint256 _stamp) returns()
func (_PriceFeedContract *PriceFeedContractTransactor) UpdateAnswer(opts *bind.TransactOpts, _round *big.Int, _value *big.Int, _stamp *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.contract.Transact(opts, "updateAnswer", _round, _value, _stamp)
}

// UpdateAnswer is a paid mutator transaction binding the contract method 0x0f7dbf0f.
//
// Solidity: function updateAnswer(uint256 _round, int256 _value, uint256 _stamp) returns()
func (_PriceFeedContract *PriceFeedContractSession) UpdateAnswer(_round *big.Int, _value *big.Int, _stamp *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpdateAnswer(&_PriceFeedContract.TransactOpts, _round, _value, _stamp)
}

// UpdateAnswer is a paid mutator transaction binding the contract method 0x0f7dbf0f.
//
// Solidity: function updateAnswer(uint256 _round, int256 _value, uint256 _stamp) returns()
func (_PriceFeedContract *PriceFeedContractTransactorSession) UpdateAnswer(_round *big.Int, _value *big.Int, _stamp *big.Int) (*types.Transaction, error) {
	return _PriceFeedContract.Contract.UpdateAnswer(&_PriceFeedContract.TransactOpts, _round, _value, _stamp)
}

// PriceFeedContractAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the PriceFeedContract contract.
type PriceFeedContractAnswerUpdatedIterator struct {
	Event *PriceFeedContractAnswerUpdated // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractAnswerUpdated)
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
		it.Event = new(PriceFeedContractAnswerUpdated)
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
func (it *PriceFeedContractAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractAnswerUpdated represents a AnswerUpdated event raised by the PriceFeedContract contract.
type PriceFeedContractAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 timestamp)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*PriceFeedContractAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractAnswerUpdatedIterator{contract: _PriceFeedContract.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 timestamp)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *PriceFeedContractAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractAnswerUpdated)
				if err := _PriceFeedContract.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 timestamp)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseAnswerUpdated(log types.Log) (*PriceFeedContractAnswerUpdated, error) {
	event := new(PriceFeedContractAnswerUpdated)
	if err := _PriceFeedContract.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceFeedContractNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the PriceFeedContract contract.
type PriceFeedContractNewRoundIterator struct {
	Event *PriceFeedContractNewRound // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractNewRound)
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
		it.Event = new(PriceFeedContractNewRound)
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
func (it *PriceFeedContractNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractNewRound represents a NewRound event raised by the PriceFeedContract contract.
type PriceFeedContractNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*PriceFeedContractNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractNewRoundIterator{contract: _PriceFeedContract.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *PriceFeedContractNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractNewRound)
				if err := _PriceFeedContract.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseNewRound(log types.Log) (*PriceFeedContractNewRound, error) {
	event := new(PriceFeedContractNewRound)
	if err := _PriceFeedContract.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceFeedContractSourceAddedIterator is returned from FilterSourceAdded and is used to iterate over the raw logs and unpacked data for SourceAdded events raised by the PriceFeedContract contract.
type PriceFeedContractSourceAddedIterator struct {
	Event *PriceFeedContractSourceAdded // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractSourceAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractSourceAdded)
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
		it.Event = new(PriceFeedContractSourceAdded)
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
func (it *PriceFeedContractSourceAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractSourceAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractSourceAdded represents a SourceAdded event raised by the PriceFeedContract contract.
type PriceFeedContractSourceAdded struct {
	Source    common.Address
	TimeStamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSourceAdded is a free log retrieval operation binding the contract event 0x4599483543f98164ba8860ed8115742ca9d0b29473d7c8f55c4e367a028d94a4.
//
// Solidity: event SourceAdded(address source, uint256 timeStamp)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterSourceAdded(opts *bind.FilterOpts) (*PriceFeedContractSourceAddedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "SourceAdded")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractSourceAddedIterator{contract: _PriceFeedContract.contract, event: "SourceAdded", logs: logs, sub: sub}, nil
}

// WatchSourceAdded is a free log subscription operation binding the contract event 0x4599483543f98164ba8860ed8115742ca9d0b29473d7c8f55c4e367a028d94a4.
//
// Solidity: event SourceAdded(address source, uint256 timeStamp)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchSourceAdded(opts *bind.WatchOpts, sink chan<- *PriceFeedContractSourceAdded) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "SourceAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractSourceAdded)
				if err := _PriceFeedContract.contract.UnpackLog(event, "SourceAdded", log); err != nil {
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

// ParseSourceAdded is a log parse operation binding the contract event 0x4599483543f98164ba8860ed8115742ca9d0b29473d7c8f55c4e367a028d94a4.
//
// Solidity: event SourceAdded(address source, uint256 timeStamp)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseSourceAdded(log types.Log) (*PriceFeedContractSourceAdded, error) {
	event := new(PriceFeedContractSourceAdded)
	if err := _PriceFeedContract.contract.UnpackLog(event, "SourceAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PriceFeedContractSourceDroppedIterator is returned from FilterSourceDropped and is used to iterate over the raw logs and unpacked data for SourceDropped events raised by the PriceFeedContract contract.
type PriceFeedContractSourceDroppedIterator struct {
	Event *PriceFeedContractSourceDropped // Event containing the contract specifics and raw log

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
func (it *PriceFeedContractSourceDroppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PriceFeedContractSourceDropped)
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
		it.Event = new(PriceFeedContractSourceDropped)
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
func (it *PriceFeedContractSourceDroppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PriceFeedContractSourceDroppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PriceFeedContractSourceDropped represents a SourceDropped event raised by the PriceFeedContract contract.
type PriceFeedContractSourceDropped struct {
	Source    common.Address
	TimeStamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSourceDropped is a free log retrieval operation binding the contract event 0x311f8a5f5a633747dc3f0b1f43ae5bdd2d3ba73f0faaab6a8a0b9fd63a784152.
//
// Solidity: event SourceDropped(address source, uint256 timeStamp)
func (_PriceFeedContract *PriceFeedContractFilterer) FilterSourceDropped(opts *bind.FilterOpts) (*PriceFeedContractSourceDroppedIterator, error) {

	logs, sub, err := _PriceFeedContract.contract.FilterLogs(opts, "SourceDropped")
	if err != nil {
		return nil, err
	}
	return &PriceFeedContractSourceDroppedIterator{contract: _PriceFeedContract.contract, event: "SourceDropped", logs: logs, sub: sub}, nil
}

// WatchSourceDropped is a free log subscription operation binding the contract event 0x311f8a5f5a633747dc3f0b1f43ae5bdd2d3ba73f0faaab6a8a0b9fd63a784152.
//
// Solidity: event SourceDropped(address source, uint256 timeStamp)
func (_PriceFeedContract *PriceFeedContractFilterer) WatchSourceDropped(opts *bind.WatchOpts, sink chan<- *PriceFeedContractSourceDropped) (event.Subscription, error) {

	logs, sub, err := _PriceFeedContract.contract.WatchLogs(opts, "SourceDropped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PriceFeedContractSourceDropped)
				if err := _PriceFeedContract.contract.UnpackLog(event, "SourceDropped", log); err != nil {
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

// ParseSourceDropped is a log parse operation binding the contract event 0x311f8a5f5a633747dc3f0b1f43ae5bdd2d3ba73f0faaab6a8a0b9fd63a784152.
//
// Solidity: event SourceDropped(address source, uint256 timeStamp)
func (_PriceFeedContract *PriceFeedContractFilterer) ParseSourceDropped(log types.Log) (*PriceFeedContractSourceDropped, error) {
	event := new(PriceFeedContractSourceDropped)
	if err := _PriceFeedContract.contract.UnpackLog(event, "SourceDropped", log); err != nil {
		return nil, err
	}
	return event, nil
}
