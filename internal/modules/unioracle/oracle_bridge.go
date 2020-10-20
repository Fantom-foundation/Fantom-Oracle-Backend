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

// UniswapOracleContractABI is the input ABI used to generate the binding from.
const UniswapOracleContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockTimestampLast\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"consult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0Average\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"_x\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1Average\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"_x\",\"type\":\"uint224\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UniswapOracleContract is an auto generated Go binding around an Ethereum contract.
type UniswapOracleContract struct {
	UniswapOracleContractCaller     // Read-only binding to the contract
	UniswapOracleContractTransactor // Write-only binding to the contract
	UniswapOracleContractFilterer   // Log filterer for contract events
}

// UniswapOracleContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapOracleContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapOracleContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapOracleContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapOracleContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapOracleContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapOracleContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapOracleContractSession struct {
	Contract     *UniswapOracleContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UniswapOracleContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapOracleContractCallerSession struct {
	Contract *UniswapOracleContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// UniswapOracleContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapOracleContractTransactorSession struct {
	Contract     *UniswapOracleContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// UniswapOracleContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapOracleContractRaw struct {
	Contract *UniswapOracleContract // Generic contract binding to access the raw methods on
}

// UniswapOracleContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapOracleContractCallerRaw struct {
	Contract *UniswapOracleContractCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapOracleContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapOracleContractTransactorRaw struct {
	Contract *UniswapOracleContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapOracleContract creates a new instance of UniswapOracleContract, bound to a specific deployed contract.
func NewUniswapOracleContract(address common.Address, backend bind.ContractBackend) (*UniswapOracleContract, error) {
	contract, err := bindUniswapOracleContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapOracleContract{UniswapOracleContractCaller: UniswapOracleContractCaller{contract: contract}, UniswapOracleContractTransactor: UniswapOracleContractTransactor{contract: contract}, UniswapOracleContractFilterer: UniswapOracleContractFilterer{contract: contract}}, nil
}

// NewUniswapOracleContractCaller creates a new read-only instance of UniswapOracleContract, bound to a specific deployed contract.
func NewUniswapOracleContractCaller(address common.Address, caller bind.ContractCaller) (*UniswapOracleContractCaller, error) {
	contract, err := bindUniswapOracleContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapOracleContractCaller{contract: contract}, nil
}

// NewUniswapOracleContractTransactor creates a new write-only instance of UniswapOracleContract, bound to a specific deployed contract.
func NewUniswapOracleContractTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapOracleContractTransactor, error) {
	contract, err := bindUniswapOracleContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapOracleContractTransactor{contract: contract}, nil
}

// NewUniswapOracleContractFilterer creates a new log filterer instance of UniswapOracleContract, bound to a specific deployed contract.
func NewUniswapOracleContractFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapOracleContractFilterer, error) {
	contract, err := bindUniswapOracleContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapOracleContractFilterer{contract: contract}, nil
}

// bindUniswapOracleContract binds a generic wrapper to an already deployed contract.
func bindUniswapOracleContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapOracleContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapOracleContract *UniswapOracleContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapOracleContract.Contract.UniswapOracleContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapOracleContract *UniswapOracleContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapOracleContract.Contract.UniswapOracleContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapOracleContract *UniswapOracleContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapOracleContract.Contract.UniswapOracleContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapOracleContract *UniswapOracleContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapOracleContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapOracleContract *UniswapOracleContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapOracleContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapOracleContract *UniswapOracleContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapOracleContract.Contract.contract.Transact(opts, method, params...)
}

// PERIOD is a free data retrieval call binding the contract method 0xb4d1d795.
//
// Solidity: function PERIOD() view returns(uint256)
func (_UniswapOracleContract *UniswapOracleContractCaller) PERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapOracleContract.contract.Call(opts, &out, "PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERIOD is a free data retrieval call binding the contract method 0xb4d1d795.
//
// Solidity: function PERIOD() view returns(uint256)
func (_UniswapOracleContract *UniswapOracleContractSession) PERIOD() (*big.Int, error) {
	return _UniswapOracleContract.Contract.PERIOD(&_UniswapOracleContract.CallOpts)
}

// PERIOD is a free data retrieval call binding the contract method 0xb4d1d795.
//
// Solidity: function PERIOD() view returns(uint256)
func (_UniswapOracleContract *UniswapOracleContractCallerSession) PERIOD() (*big.Int, error) {
	return _UniswapOracleContract.Contract.PERIOD(&_UniswapOracleContract.CallOpts)
}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint32)
func (_UniswapOracleContract *UniswapOracleContractCaller) BlockTimestampLast(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _UniswapOracleContract.contract.Call(opts, &out, "blockTimestampLast")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint32)
func (_UniswapOracleContract *UniswapOracleContractSession) BlockTimestampLast() (uint32, error) {
	return _UniswapOracleContract.Contract.BlockTimestampLast(&_UniswapOracleContract.CallOpts)
}

// BlockTimestampLast is a free data retrieval call binding the contract method 0xc5700a02.
//
// Solidity: function blockTimestampLast() view returns(uint32)
func (_UniswapOracleContract *UniswapOracleContractCallerSession) BlockTimestampLast() (uint32, error) {
	return _UniswapOracleContract.Contract.BlockTimestampLast(&_UniswapOracleContract.CallOpts)
}

// Consult is a free data retrieval call binding the contract method 0x3ddac953.
//
// Solidity: function consult(address token, uint256 amountIn) view returns(uint256 amountOut)
func (_UniswapOracleContract *UniswapOracleContractCaller) Consult(opts *bind.CallOpts, token common.Address, amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UniswapOracleContract.contract.Call(opts, &out, "consult", token, amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Consult is a free data retrieval call binding the contract method 0x3ddac953.
//
// Solidity: function consult(address token, uint256 amountIn) view returns(uint256 amountOut)
func (_UniswapOracleContract *UniswapOracleContractSession) Consult(token common.Address, amountIn *big.Int) (*big.Int, error) {
	return _UniswapOracleContract.Contract.Consult(&_UniswapOracleContract.CallOpts, token, amountIn)
}

// Consult is a free data retrieval call binding the contract method 0x3ddac953.
//
// Solidity: function consult(address token, uint256 amountIn) view returns(uint256 amountOut)
func (_UniswapOracleContract *UniswapOracleContractCallerSession) Consult(token common.Address, amountIn *big.Int) (*big.Int, error) {
	return _UniswapOracleContract.Contract.Consult(&_UniswapOracleContract.CallOpts, token, amountIn)
}

// Price0Average is a free data retrieval call binding the contract method 0xa6bb4539.
//
// Solidity: function price0Average() view returns(uint224 _x)
func (_UniswapOracleContract *UniswapOracleContractCaller) Price0Average(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapOracleContract.contract.Call(opts, &out, "price0Average")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0Average is a free data retrieval call binding the contract method 0xa6bb4539.
//
// Solidity: function price0Average() view returns(uint224 _x)
func (_UniswapOracleContract *UniswapOracleContractSession) Price0Average() (*big.Int, error) {
	return _UniswapOracleContract.Contract.Price0Average(&_UniswapOracleContract.CallOpts)
}

// Price0Average is a free data retrieval call binding the contract method 0xa6bb4539.
//
// Solidity: function price0Average() view returns(uint224 _x)
func (_UniswapOracleContract *UniswapOracleContractCallerSession) Price0Average() (*big.Int, error) {
	return _UniswapOracleContract.Contract.Price0Average(&_UniswapOracleContract.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_UniswapOracleContract *UniswapOracleContractCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapOracleContract.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_UniswapOracleContract *UniswapOracleContractSession) Price0CumulativeLast() (*big.Int, error) {
	return _UniswapOracleContract.Contract.Price0CumulativeLast(&_UniswapOracleContract.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_UniswapOracleContract *UniswapOracleContractCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _UniswapOracleContract.Contract.Price0CumulativeLast(&_UniswapOracleContract.CallOpts)
}

// Price1Average is a free data retrieval call binding the contract method 0x5e6aaf2c.
//
// Solidity: function price1Average() view returns(uint224 _x)
func (_UniswapOracleContract *UniswapOracleContractCaller) Price1Average(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapOracleContract.contract.Call(opts, &out, "price1Average")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1Average is a free data retrieval call binding the contract method 0x5e6aaf2c.
//
// Solidity: function price1Average() view returns(uint224 _x)
func (_UniswapOracleContract *UniswapOracleContractSession) Price1Average() (*big.Int, error) {
	return _UniswapOracleContract.Contract.Price1Average(&_UniswapOracleContract.CallOpts)
}

// Price1Average is a free data retrieval call binding the contract method 0x5e6aaf2c.
//
// Solidity: function price1Average() view returns(uint224 _x)
func (_UniswapOracleContract *UniswapOracleContractCallerSession) Price1Average() (*big.Int, error) {
	return _UniswapOracleContract.Contract.Price1Average(&_UniswapOracleContract.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_UniswapOracleContract *UniswapOracleContractCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniswapOracleContract.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_UniswapOracleContract *UniswapOracleContractSession) Price1CumulativeLast() (*big.Int, error) {
	return _UniswapOracleContract.Contract.Price1CumulativeLast(&_UniswapOracleContract.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_UniswapOracleContract *UniswapOracleContractCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _UniswapOracleContract.Contract.Price1CumulativeLast(&_UniswapOracleContract.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_UniswapOracleContract *UniswapOracleContractCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapOracleContract.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_UniswapOracleContract *UniswapOracleContractSession) Token0() (common.Address, error) {
	return _UniswapOracleContract.Contract.Token0(&_UniswapOracleContract.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_UniswapOracleContract *UniswapOracleContractCallerSession) Token0() (common.Address, error) {
	return _UniswapOracleContract.Contract.Token0(&_UniswapOracleContract.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_UniswapOracleContract *UniswapOracleContractCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapOracleContract.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_UniswapOracleContract *UniswapOracleContractSession) Token1() (common.Address, error) {
	return _UniswapOracleContract.Contract.Token1(&_UniswapOracleContract.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_UniswapOracleContract *UniswapOracleContractCallerSession) Token1() (common.Address, error) {
	return _UniswapOracleContract.Contract.Token1(&_UniswapOracleContract.CallOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_UniswapOracleContract *UniswapOracleContractTransactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapOracleContract.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_UniswapOracleContract *UniswapOracleContractSession) Update() (*types.Transaction, error) {
	return _UniswapOracleContract.Contract.Update(&_UniswapOracleContract.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_UniswapOracleContract *UniswapOracleContractTransactorSession) Update() (*types.Transaction, error) {
	return _UniswapOracleContract.Contract.Update(&_UniswapOracleContract.TransactOpts)
}
