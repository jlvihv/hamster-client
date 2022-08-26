// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake_distribution_proxy

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StakeDistributionProxyMetaData contains all meta data concerning the StakeDistributionProxy contract.
var StakeDistributionProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexerWalletAddress\",\"type\":\"address\"}],\"name\":\"setIndexerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingAmount\",\"type\":\"uint256\"}],\"name\":\"staking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingAmount\",\"type\":\"uint256\"}],\"name\":\"rePledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawIncome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"retrieveStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gainIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnStakingAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakeDistributionProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeDistributionProxyMetaData.ABI instead.
var StakeDistributionProxyABI = StakeDistributionProxyMetaData.ABI

// StakeDistributionProxy is an auto generated Go binding around an Ethereum contract.
type StakeDistributionProxy struct {
	StakeDistributionProxyCaller     // Read-only binding to the contract
	StakeDistributionProxyTransactor // Write-only binding to the contract
	StakeDistributionProxyFilterer   // Log filterer for contract events
}

// StakeDistributionProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeDistributionProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeDistributionProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeDistributionProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeDistributionProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeDistributionProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeDistributionProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeDistributionProxySession struct {
	Contract     *StakeDistributionProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakeDistributionProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeDistributionProxyCallerSession struct {
	Contract *StakeDistributionProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// StakeDistributionProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeDistributionProxyTransactorSession struct {
	Contract     *StakeDistributionProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// StakeDistributionProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeDistributionProxyRaw struct {
	Contract *StakeDistributionProxy // Generic contract binding to access the raw methods on
}

// StakeDistributionProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeDistributionProxyCallerRaw struct {
	Contract *StakeDistributionProxyCaller // Generic read-only contract binding to access the raw methods on
}

// StakeDistributionProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeDistributionProxyTransactorRaw struct {
	Contract *StakeDistributionProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeDistributionProxy creates a new instance of StakeDistributionProxy, bound to a specific deployed contract.
func NewStakeDistributionProxy(address common.Address, backend bind.ContractBackend) (*StakeDistributionProxy, error) {
	contract, err := bindStakeDistributionProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeDistributionProxy{StakeDistributionProxyCaller: StakeDistributionProxyCaller{contract: contract}, StakeDistributionProxyTransactor: StakeDistributionProxyTransactor{contract: contract}, StakeDistributionProxyFilterer: StakeDistributionProxyFilterer{contract: contract}}, nil
}

// NewStakeDistributionProxyCaller creates a new read-only instance of StakeDistributionProxy, bound to a specific deployed contract.
func NewStakeDistributionProxyCaller(address common.Address, caller bind.ContractCaller) (*StakeDistributionProxyCaller, error) {
	contract, err := bindStakeDistributionProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeDistributionProxyCaller{contract: contract}, nil
}

// NewStakeDistributionProxyTransactor creates a new write-only instance of StakeDistributionProxy, bound to a specific deployed contract.
func NewStakeDistributionProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeDistributionProxyTransactor, error) {
	contract, err := bindStakeDistributionProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeDistributionProxyTransactor{contract: contract}, nil
}

// NewStakeDistributionProxyFilterer creates a new log filterer instance of StakeDistributionProxy, bound to a specific deployed contract.
func NewStakeDistributionProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeDistributionProxyFilterer, error) {
	contract, err := bindStakeDistributionProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeDistributionProxyFilterer{contract: contract}, nil
}

// bindStakeDistributionProxy binds a generic wrapper to an already deployed contract.
func bindStakeDistributionProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeDistributionProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeDistributionProxy *StakeDistributionProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeDistributionProxy.Contract.StakeDistributionProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeDistributionProxy *StakeDistributionProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.StakeDistributionProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeDistributionProxy *StakeDistributionProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.StakeDistributionProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeDistributionProxy *StakeDistributionProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeDistributionProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeDistributionProxy *StakeDistributionProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeDistributionProxy *StakeDistributionProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.contract.Transact(opts, method, params...)
}

// GainIncome is a free data retrieval call binding the contract method 0xac06f366.
//
// Solidity: function gainIncome() view returns(uint256)
func (_StakeDistributionProxy *StakeDistributionProxyCaller) GainIncome(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeDistributionProxy.contract.Call(opts, &out, "gainIncome")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GainIncome is a free data retrieval call binding the contract method 0xac06f366.
//
// Solidity: function gainIncome() view returns(uint256)
func (_StakeDistributionProxy *StakeDistributionProxySession) GainIncome() (*big.Int, error) {
	return _StakeDistributionProxy.Contract.GainIncome(&_StakeDistributionProxy.CallOpts)
}

// GainIncome is a free data retrieval call binding the contract method 0xac06f366.
//
// Solidity: function gainIncome() view returns(uint256)
func (_StakeDistributionProxy *StakeDistributionProxyCallerSession) GainIncome() (*big.Int, error) {
	return _StakeDistributionProxy.Contract.GainIncome(&_StakeDistributionProxy.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_StakeDistributionProxy *StakeDistributionProxyCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeDistributionProxy.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_StakeDistributionProxy *StakeDistributionProxySession) GetBalance() (*big.Int, error) {
	return _StakeDistributionProxy.Contract.GetBalance(&_StakeDistributionProxy.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_StakeDistributionProxy *StakeDistributionProxyCallerSession) GetBalance() (*big.Int, error) {
	return _StakeDistributionProxy.Contract.GetBalance(&_StakeDistributionProxy.CallOpts)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() view returns(address)
func (_StakeDistributionProxy *StakeDistributionProxyCaller) GetProxyAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeDistributionProxy.contract.Call(opts, &out, "getProxyAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() view returns(address)
func (_StakeDistributionProxy *StakeDistributionProxySession) GetProxyAddress() (common.Address, error) {
	return _StakeDistributionProxy.Contract.GetProxyAddress(&_StakeDistributionProxy.CallOpts)
}

// GetProxyAddress is a free data retrieval call binding the contract method 0x43a73d9a.
//
// Solidity: function getProxyAddress() view returns(address)
func (_StakeDistributionProxy *StakeDistributionProxyCallerSession) GetProxyAddress() (common.Address, error) {
	return _StakeDistributionProxy.Contract.GetProxyAddress(&_StakeDistributionProxy.CallOpts)
}

// GetStakingAmount is a free data retrieval call binding the contract method 0xced066c9.
//
// Solidity: function getStakingAmount() view returns(uint256)
func (_StakeDistributionProxy *StakeDistributionProxyCaller) GetStakingAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeDistributionProxy.contract.Call(opts, &out, "getStakingAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingAmount is a free data retrieval call binding the contract method 0xced066c9.
//
// Solidity: function getStakingAmount() view returns(uint256)
func (_StakeDistributionProxy *StakeDistributionProxySession) GetStakingAmount() (*big.Int, error) {
	return _StakeDistributionProxy.Contract.GetStakingAmount(&_StakeDistributionProxy.CallOpts)
}

// GetStakingAmount is a free data retrieval call binding the contract method 0xced066c9.
//
// Solidity: function getStakingAmount() view returns(uint256)
func (_StakeDistributionProxy *StakeDistributionProxyCallerSession) GetStakingAmount() (*big.Int, error) {
	return _StakeDistributionProxy.Contract.GetStakingAmount(&_StakeDistributionProxy.CallOpts)
}

// GetUnStakingAmount is a free data retrieval call binding the contract method 0xb9513408.
//
// Solidity: function getUnStakingAmount() view returns(uint256)
func (_StakeDistributionProxy *StakeDistributionProxyCaller) GetUnStakingAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeDistributionProxy.contract.Call(opts, &out, "getUnStakingAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnStakingAmount is a free data retrieval call binding the contract method 0xb9513408.
//
// Solidity: function getUnStakingAmount() view returns(uint256)
func (_StakeDistributionProxy *StakeDistributionProxySession) GetUnStakingAmount() (*big.Int, error) {
	return _StakeDistributionProxy.Contract.GetUnStakingAmount(&_StakeDistributionProxy.CallOpts)
}

// GetUnStakingAmount is a free data retrieval call binding the contract method 0xb9513408.
//
// Solidity: function getUnStakingAmount() view returns(uint256)
func (_StakeDistributionProxy *StakeDistributionProxyCallerSession) GetUnStakingAmount() (*big.Int, error) {
	return _StakeDistributionProxy.Contract.GetUnStakingAmount(&_StakeDistributionProxy.CallOpts)
}

// RePledge is a paid mutator transaction binding the contract method 0x43d342a1.
//
// Solidity: function rePledge(uint256 _stakingAmount) returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactor) RePledge(opts *bind.TransactOpts, _stakingAmount *big.Int) (*types.Transaction, error) {
	return _StakeDistributionProxy.contract.Transact(opts, "rePledge", _stakingAmount)
}

// RePledge is a paid mutator transaction binding the contract method 0x43d342a1.
//
// Solidity: function rePledge(uint256 _stakingAmount) returns()
func (_StakeDistributionProxy *StakeDistributionProxySession) RePledge(_stakingAmount *big.Int) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.RePledge(&_StakeDistributionProxy.TransactOpts, _stakingAmount)
}

// RePledge is a paid mutator transaction binding the contract method 0x43d342a1.
//
// Solidity: function rePledge(uint256 _stakingAmount) returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactorSession) RePledge(_stakingAmount *big.Int) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.RePledge(&_StakeDistributionProxy.TransactOpts, _stakingAmount)
}

// RetrieveStaking is a paid mutator transaction binding the contract method 0xe44ad80c.
//
// Solidity: function retrieveStaking(uint256 _tokens) returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactor) RetrieveStaking(opts *bind.TransactOpts, _tokens *big.Int) (*types.Transaction, error) {
	return _StakeDistributionProxy.contract.Transact(opts, "retrieveStaking", _tokens)
}

// RetrieveStaking is a paid mutator transaction binding the contract method 0xe44ad80c.
//
// Solidity: function retrieveStaking(uint256 _tokens) returns()
func (_StakeDistributionProxy *StakeDistributionProxySession) RetrieveStaking(_tokens *big.Int) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.RetrieveStaking(&_StakeDistributionProxy.TransactOpts, _tokens)
}

// RetrieveStaking is a paid mutator transaction binding the contract method 0xe44ad80c.
//
// Solidity: function retrieveStaking(uint256 _tokens) returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactorSession) RetrieveStaking(_tokens *big.Int) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.RetrieveStaking(&_StakeDistributionProxy.TransactOpts, _tokens)
}

// SetIndexerAddress is a paid mutator transaction binding the contract method 0xb6ba3e13.
//
// Solidity: function setIndexerAddress(address _indexerWalletAddress) returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactor) SetIndexerAddress(opts *bind.TransactOpts, _indexerWalletAddress common.Address) (*types.Transaction, error) {
	return _StakeDistributionProxy.contract.Transact(opts, "setIndexerAddress", _indexerWalletAddress)
}

// SetIndexerAddress is a paid mutator transaction binding the contract method 0xb6ba3e13.
//
// Solidity: function setIndexerAddress(address _indexerWalletAddress) returns()
func (_StakeDistributionProxy *StakeDistributionProxySession) SetIndexerAddress(_indexerWalletAddress common.Address) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.SetIndexerAddress(&_StakeDistributionProxy.TransactOpts, _indexerWalletAddress)
}

// SetIndexerAddress is a paid mutator transaction binding the contract method 0xb6ba3e13.
//
// Solidity: function setIndexerAddress(address _indexerWalletAddress) returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactorSession) SetIndexerAddress(_indexerWalletAddress common.Address) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.SetIndexerAddress(&_StakeDistributionProxy.TransactOpts, _indexerWalletAddress)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool _allowed) returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address, _allowed bool) (*types.Transaction, error) {
	return _StakeDistributionProxy.contract.Transact(opts, "setOperator", _operator, _allowed)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool _allowed) returns()
func (_StakeDistributionProxy *StakeDistributionProxySession) SetOperator(_operator common.Address, _allowed bool) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.SetOperator(&_StakeDistributionProxy.TransactOpts, _operator, _allowed)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address _operator, bool _allowed) returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactorSession) SetOperator(_operator common.Address, _allowed bool) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.SetOperator(&_StakeDistributionProxy.TransactOpts, _operator, _allowed)
}

// Staking is a paid mutator transaction binding the contract method 0x1dbb2a22.
//
// Solidity: function staking(uint256 _stakingAmount) returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactor) Staking(opts *bind.TransactOpts, _stakingAmount *big.Int) (*types.Transaction, error) {
	return _StakeDistributionProxy.contract.Transact(opts, "staking", _stakingAmount)
}

// Staking is a paid mutator transaction binding the contract method 0x1dbb2a22.
//
// Solidity: function staking(uint256 _stakingAmount) returns()
func (_StakeDistributionProxy *StakeDistributionProxySession) Staking(_stakingAmount *big.Int) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.Staking(&_StakeDistributionProxy.TransactOpts, _stakingAmount)
}

// Staking is a paid mutator transaction binding the contract method 0x1dbb2a22.
//
// Solidity: function staking(uint256 _stakingAmount) returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactorSession) Staking(_stakingAmount *big.Int) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.Staking(&_StakeDistributionProxy.TransactOpts, _stakingAmount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _tokens) returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactor) Unstake(opts *bind.TransactOpts, _tokens *big.Int) (*types.Transaction, error) {
	return _StakeDistributionProxy.contract.Transact(opts, "unstake", _tokens)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _tokens) returns()
func (_StakeDistributionProxy *StakeDistributionProxySession) Unstake(_tokens *big.Int) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.Unstake(&_StakeDistributionProxy.TransactOpts, _tokens)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _tokens) returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactorSession) Unstake(_tokens *big.Int) (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.Unstake(&_StakeDistributionProxy.TransactOpts, _tokens)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeDistributionProxy.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StakeDistributionProxy *StakeDistributionProxySession) Withdraw() (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.Withdraw(&_StakeDistributionProxy.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactorSession) Withdraw() (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.Withdraw(&_StakeDistributionProxy.TransactOpts)
}

// WithdrawIncome is a paid mutator transaction binding the contract method 0x1a0fa8c3.
//
// Solidity: function withdrawIncome() returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactor) WithdrawIncome(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeDistributionProxy.contract.Transact(opts, "withdrawIncome")
}

// WithdrawIncome is a paid mutator transaction binding the contract method 0x1a0fa8c3.
//
// Solidity: function withdrawIncome() returns()
func (_StakeDistributionProxy *StakeDistributionProxySession) WithdrawIncome() (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.WithdrawIncome(&_StakeDistributionProxy.TransactOpts)
}

// WithdrawIncome is a paid mutator transaction binding the contract method 0x1a0fa8c3.
//
// Solidity: function withdrawIncome() returns()
func (_StakeDistributionProxy *StakeDistributionProxyTransactorSession) WithdrawIncome() (*types.Transaction, error) {
	return _StakeDistributionProxy.Contract.WithdrawIncome(&_StakeDistributionProxy.TransactOpts)
}
