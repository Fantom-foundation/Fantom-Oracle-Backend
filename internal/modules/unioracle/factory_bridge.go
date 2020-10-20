// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unioracle

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

// UniswapOracleFactoryContractABI is the input ABI used to generate the binding from.
const UniswapOracleFactoryContractABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"deploy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deployed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"oraclesArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UniswapOracleFactoryContract is an auto generated Go binding around an Ethereum contract.
type UniswapOracleFactoryContract struct {
	UniswapOracleFactoryContractCaller     // Read-only binding to the contract
	UniswapOracleFactoryContractTransactor // Write-only binding to the contract
	UniswapOracleFactoryContractFilterer   // Log filterer for contract events
}

// UniswapOracleFactoryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapOracleFactoryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapOracleFactoryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapOracleFactoryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapOracleFactoryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapOracleFactoryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapOracleFactoryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapOracleFactoryContractSession struct {
	Contract     *UniswapOracleFactoryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// UniswapOracleFactoryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapOracleFactoryContractCallerSession struct {
	Contract *UniswapOracleFactoryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// UniswapOracleFactoryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapOracleFactoryContractTransactorSession struct {
	Contract     *UniswapOracleFactoryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// UniswapOracleFactoryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapOracleFactoryContractRaw struct {
	Contract *UniswapOracleFactoryContract // Generic contract binding to access the raw methods on
}

// UniswapOracleFactoryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapOracleFactoryContractCallerRaw struct {
	Contract *UniswapOracleFactoryContractCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapOracleFactoryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapOracleFactoryContractTransactorRaw struct {
	Contract *UniswapOracleFactoryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapOracleFactoryContract creates a new instance of UniswapOracleFactoryContract, bound to a specific deployed contract.
func NewUniswapOracleFactoryContract(address common.Address, backend bind.ContractBackend) (*UniswapOracleFactoryContract, error) {
	contract, err := bindUniswapOracleFactoryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapOracleFactoryContract{UniswapOracleFactoryContractCaller: UniswapOracleFactoryContractCaller{contract: contract}, UniswapOracleFactoryContractTransactor: UniswapOracleFactoryContractTransactor{contract: contract}, UniswapOracleFactoryContractFilterer: UniswapOracleFactoryContractFilterer{contract: contract}}, nil
}

// NewUniswapOracleFactoryContractCaller creates a new read-only instance of UniswapOracleFactoryContract, bound to a specific deployed contract.
func NewUniswapOracleFactoryContractCaller(address common.Address, caller bind.ContractCaller) (*UniswapOracleFactoryContractCaller, error) {
	contract, err := bindUniswapOracleFactoryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapOracleFactoryContractCaller{contract: contract}, nil
}

// NewUniswapOracleFactoryContractTransactor creates a new write-only instance of UniswapOracleFactoryContract, bound to a specific deployed contract.
func NewUniswapOracleFactoryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapOracleFactoryContractTransactor, error) {
	contract, err := bindUniswapOracleFactoryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapOracleFactoryContractTransactor{contract: contract}, nil
}

// NewUniswapOracleFactoryContractFilterer creates a new log filterer instance of UniswapOracleFactoryContract, bound to a specific deployed contract.
func NewUniswapOracleFactoryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapOracleFactoryContractFilterer, error) {
	contract, err := bindUniswapOracleFactoryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapOracleFactoryContractFilterer{contract: contract}, nil
}

// bindUniswapOracleFactoryContract binds a generic wrapper to an already deployed contract.
func bindUniswapOracleFactoryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapOracleFactoryContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapOracleFactoryContract.Contract.UniswapOracleFactoryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapOracleFactoryContract.Contract.UniswapOracleFactoryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapOracleFactoryContract.Contract.UniswapOracleFactoryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapOracleFactoryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapOracleFactoryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapOracleFactoryContract.Contract.contract.Transact(opts, method, params...)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCaller) FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapOracleFactoryContract.contract.Call(opts, &out, "FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractSession) FACTORY() (common.Address, error) {
	return _UniswapOracleFactoryContract.Contract.FACTORY(&_UniswapOracleFactoryContract.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCallerSession) FACTORY() (common.Address, error) {
	return _UniswapOracleFactoryContract.Contract.FACTORY(&_UniswapOracleFactoryContract.CallOpts)
}

// Deployed is a free data retrieval call binding the contract method 0x85bb3923.
//
// Solidity: function deployed(address ) view returns(bool)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCaller) Deployed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _UniswapOracleFactoryContract.contract.Call(opts, &out, "deployed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Deployed is a free data retrieval call binding the contract method 0x85bb3923.
//
// Solidity: function deployed(address ) view returns(bool)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractSession) Deployed(arg0 common.Address) (bool, error) {
	return _UniswapOracleFactoryContract.Contract.Deployed(&_UniswapOracleFactoryContract.CallOpts, arg0)
}

// Deployed is a free data retrieval call binding the contract method 0x85bb3923.
//
// Solidity: function deployed(address ) view returns(bool)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCallerSession) Deployed(arg0 common.Address) (bool, error) {
	return _UniswapOracleFactoryContract.Contract.Deployed(&_UniswapOracleFactoryContract.CallOpts, arg0)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) pure returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _UniswapOracleFactoryContract.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) pure returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _UniswapOracleFactoryContract.Contract.GetPair(&_UniswapOracleFactoryContract.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) pure returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _UniswapOracleFactoryContract.Contract.GetPair(&_UniswapOracleFactoryContract.CallOpts, tokenA, tokenB)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapOracleFactoryContract.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractSession) Governance() (common.Address, error) {
	return _UniswapOracleFactoryContract.Contract.Governance(&_UniswapOracleFactoryContract.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCallerSession) Governance() (common.Address, error) {
	return _UniswapOracleFactoryContract.Contract.Governance(&_UniswapOracleFactoryContract.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[])
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCaller) Oracles(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _UniswapOracleFactoryContract.contract.Call(opts, &out, "oracles")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[])
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractSession) Oracles() ([]common.Address, error) {
	return _UniswapOracleFactoryContract.Contract.Oracles(&_UniswapOracleFactoryContract.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[])
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCallerSession) Oracles() ([]common.Address, error) {
	return _UniswapOracleFactoryContract.Contract.Oracles(&_UniswapOracleFactoryContract.CallOpts)
}

// OraclesArray is a free data retrieval call binding the contract method 0x52ee56cd.
//
// Solidity: function oraclesArray(uint256 ) view returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCaller) OraclesArray(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UniswapOracleFactoryContract.contract.Call(opts, &out, "oraclesArray", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OraclesArray is a free data retrieval call binding the contract method 0x52ee56cd.
//
// Solidity: function oraclesArray(uint256 ) view returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractSession) OraclesArray(arg0 *big.Int) (common.Address, error) {
	return _UniswapOracleFactoryContract.Contract.OraclesArray(&_UniswapOracleFactoryContract.CallOpts, arg0)
}

// OraclesArray is a free data retrieval call binding the contract method 0x52ee56cd.
//
// Solidity: function oraclesArray(uint256 ) view returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCallerSession) OraclesArray(arg0 *big.Int) (common.Address, error) {
	return _UniswapOracleFactoryContract.Contract.OraclesArray(&_UniswapOracleFactoryContract.CallOpts, arg0)
}

// Quote is a free data retrieval call binding the contract method 0xb6466384.
//
// Solidity: function quote(address tokenIn, address tokenOut, uint256 amountIn) view returns(uint256 amountOut)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCaller) Quote(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapOracleFactoryContract.contract.Call(opts, &out, "quote", tokenIn, tokenOut, amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xb6466384.
//
// Solidity: function quote(address tokenIn, address tokenOut, uint256 amountIn) view returns(uint256 amountOut)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractSession) Quote(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int) (*big.Int, error) {
	return _UniswapOracleFactoryContract.Contract.Quote(&_UniswapOracleFactoryContract.CallOpts, tokenIn, tokenOut, amountIn)
}

// Quote is a free data retrieval call binding the contract method 0xb6466384.
//
// Solidity: function quote(address tokenIn, address tokenOut, uint256 amountIn) view returns(uint256 amountOut)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCallerSession) Quote(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int) (*big.Int, error) {
	return _UniswapOracleFactoryContract.Contract.Quote(&_UniswapOracleFactoryContract.CallOpts, tokenIn, tokenOut, amountIn)
}

// Registry is a free data retrieval call binding the contract method 0x038defd7.
//
// Solidity: function registry(address ) view returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCaller) Registry(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _UniswapOracleFactoryContract.contract.Call(opts, &out, "registry", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x038defd7.
//
// Solidity: function registry(address ) view returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractSession) Registry(arg0 common.Address) (common.Address, error) {
	return _UniswapOracleFactoryContract.Contract.Registry(&_UniswapOracleFactoryContract.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x038defd7.
//
// Solidity: function registry(address ) view returns(address)
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractCallerSession) Registry(arg0 common.Address) (common.Address, error) {
	return _UniswapOracleFactoryContract.Contract.Registry(&_UniswapOracleFactoryContract.CallOpts, arg0)
}

// Deploy is a paid mutator transaction binding the contract method 0x545e7c61.
//
// Solidity: function deploy(address tokenA, address tokenB) returns()
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractTransactor) Deploy(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _UniswapOracleFactoryContract.contract.Transact(opts, "deploy", tokenA, tokenB)
}

// Deploy is a paid mutator transaction binding the contract method 0x545e7c61.
//
// Solidity: function deploy(address tokenA, address tokenB) returns()
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractSession) Deploy(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _UniswapOracleFactoryContract.Contract.Deploy(&_UniswapOracleFactoryContract.TransactOpts, tokenA, tokenB)
}

// Deploy is a paid mutator transaction binding the contract method 0x545e7c61.
//
// Solidity: function deploy(address tokenA, address tokenB) returns()
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractTransactorSession) Deploy(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _UniswapOracleFactoryContract.Contract.Deploy(&_UniswapOracleFactoryContract.TransactOpts, tokenA, tokenB)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractTransactor) SetGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _UniswapOracleFactoryContract.contract.Transact(opts, "setGovernance", _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _UniswapOracleFactoryContract.Contract.SetGovernance(&_UniswapOracleFactoryContract.TransactOpts, _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractTransactorSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _UniswapOracleFactoryContract.Contract.SetGovernance(&_UniswapOracleFactoryContract.TransactOpts, _governance)
}

// Update is a paid mutator transaction binding the contract method 0xc640752d.
//
// Solidity: function update(address tokenA, address tokenB) returns()
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractTransactor) Update(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _UniswapOracleFactoryContract.contract.Transact(opts, "update", tokenA, tokenB)
}

// Update is a paid mutator transaction binding the contract method 0xc640752d.
//
// Solidity: function update(address tokenA, address tokenB) returns()
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractSession) Update(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _UniswapOracleFactoryContract.Contract.Update(&_UniswapOracleFactoryContract.TransactOpts, tokenA, tokenB)
}

// Update is a paid mutator transaction binding the contract method 0xc640752d.
//
// Solidity: function update(address tokenA, address tokenB) returns()
func (_UniswapOracleFactoryContract *UniswapOracleFactoryContractTransactorSession) Update(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _UniswapOracleFactoryContract.Contract.Update(&_UniswapOracleFactoryContract.TransactOpts, tokenA, tokenB)
}
