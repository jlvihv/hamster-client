// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stake_proxy_factory

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

// StakeProxyFactoryMetaData contains all meta data concerning the StakeProxyFactory contract.
var StakeProxyFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexerWalletAddress\",\"type\":\"address\"}],\"name\":\"createStakingContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_indexerWalletAddress\",\"type\":\"address\"}],\"name\":\"getStakingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakeProxyFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeProxyFactoryMetaData.ABI instead.
var StakeProxyFactoryABI = StakeProxyFactoryMetaData.ABI

// StakeProxyFactory is an auto generated Go binding around an Ethereum contract.
type StakeProxyFactory struct {
	StakeProxyFactoryCaller     // Read-only binding to the contract
	StakeProxyFactoryTransactor // Write-only binding to the contract
	StakeProxyFactoryFilterer   // Log filterer for contract events
}

// StakeProxyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeProxyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeProxyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeProxyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeProxyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeProxyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeProxyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeProxyFactorySession struct {
	Contract     *StakeProxyFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakeProxyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeProxyFactoryCallerSession struct {
	Contract *StakeProxyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// StakeProxyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeProxyFactoryTransactorSession struct {
	Contract     *StakeProxyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// StakeProxyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeProxyFactoryRaw struct {
	Contract *StakeProxyFactory // Generic contract binding to access the raw methods on
}

// StakeProxyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeProxyFactoryCallerRaw struct {
	Contract *StakeProxyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// StakeProxyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeProxyFactoryTransactorRaw struct {
	Contract *StakeProxyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeProxyFactory creates a new instance of StakeProxyFactory, bound to a specific deployed contract.
func NewStakeProxyFactory(address common.Address, backend bind.ContractBackend) (*StakeProxyFactory, error) {
	contract, err := bindStakeProxyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeProxyFactory{StakeProxyFactoryCaller: StakeProxyFactoryCaller{contract: contract}, StakeProxyFactoryTransactor: StakeProxyFactoryTransactor{contract: contract}, StakeProxyFactoryFilterer: StakeProxyFactoryFilterer{contract: contract}}, nil
}

// NewStakeProxyFactoryCaller creates a new read-only instance of StakeProxyFactory, bound to a specific deployed contract.
func NewStakeProxyFactoryCaller(address common.Address, caller bind.ContractCaller) (*StakeProxyFactoryCaller, error) {
	contract, err := bindStakeProxyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeProxyFactoryCaller{contract: contract}, nil
}

// NewStakeProxyFactoryTransactor creates a new write-only instance of StakeProxyFactory, bound to a specific deployed contract.
func NewStakeProxyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeProxyFactoryTransactor, error) {
	contract, err := bindStakeProxyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeProxyFactoryTransactor{contract: contract}, nil
}

// NewStakeProxyFactoryFilterer creates a new log filterer instance of StakeProxyFactory, bound to a specific deployed contract.
func NewStakeProxyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeProxyFactoryFilterer, error) {
	contract, err := bindStakeProxyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeProxyFactoryFilterer{contract: contract}, nil
}

// bindStakeProxyFactory binds a generic wrapper to an already deployed contract.
func bindStakeProxyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeProxyFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeProxyFactory *StakeProxyFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeProxyFactory.Contract.StakeProxyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeProxyFactory *StakeProxyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeProxyFactory.Contract.StakeProxyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeProxyFactory *StakeProxyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeProxyFactory.Contract.StakeProxyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeProxyFactory *StakeProxyFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeProxyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeProxyFactory *StakeProxyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeProxyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeProxyFactory *StakeProxyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeProxyFactory.Contract.contract.Transact(opts, method, params...)
}

// GetStakingAddress is a free data retrieval call binding the contract method 0x7465cf97.
//
// Solidity: function getStakingAddress(address _indexerWalletAddress) view returns(address)
func (_StakeProxyFactory *StakeProxyFactoryCaller) GetStakingAddress(opts *bind.CallOpts, _indexerWalletAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _StakeProxyFactory.contract.Call(opts, &out, "getStakingAddress", _indexerWalletAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakingAddress is a free data retrieval call binding the contract method 0x7465cf97.
//
// Solidity: function getStakingAddress(address _indexerWalletAddress) view returns(address)
func (_StakeProxyFactory *StakeProxyFactorySession) GetStakingAddress(_indexerWalletAddress common.Address) (common.Address, error) {
	return _StakeProxyFactory.Contract.GetStakingAddress(&_StakeProxyFactory.CallOpts, _indexerWalletAddress)
}

// GetStakingAddress is a free data retrieval call binding the contract method 0x7465cf97.
//
// Solidity: function getStakingAddress(address _indexerWalletAddress) view returns(address)
func (_StakeProxyFactory *StakeProxyFactoryCallerSession) GetStakingAddress(_indexerWalletAddress common.Address) (common.Address, error) {
	return _StakeProxyFactory.Contract.GetStakingAddress(&_StakeProxyFactory.CallOpts, _indexerWalletAddress)
}

// CreateStakingContract is a paid mutator transaction binding the contract method 0x4a4cef40.
//
// Solidity: function createStakingContract(address _indexerWalletAddress) returns()
func (_StakeProxyFactory *StakeProxyFactoryTransactor) CreateStakingContract(opts *bind.TransactOpts, _indexerWalletAddress common.Address) (*types.Transaction, error) {
	return _StakeProxyFactory.contract.Transact(opts, "createStakingContract", _indexerWalletAddress)
}

// CreateStakingContract is a paid mutator transaction binding the contract method 0x4a4cef40.
//
// Solidity: function createStakingContract(address _indexerWalletAddress) returns()
func (_StakeProxyFactory *StakeProxyFactorySession) CreateStakingContract(_indexerWalletAddress common.Address) (*types.Transaction, error) {
	return _StakeProxyFactory.Contract.CreateStakingContract(&_StakeProxyFactory.TransactOpts, _indexerWalletAddress)
}

// CreateStakingContract is a paid mutator transaction binding the contract method 0x4a4cef40.
//
// Solidity: function createStakingContract(address _indexerWalletAddress) returns()
func (_StakeProxyFactory *StakeProxyFactoryTransactorSession) CreateStakingContract(_indexerWalletAddress common.Address) (*types.Transaction, error) {
	return _StakeProxyFactory.Contract.CreateStakingContract(&_StakeProxyFactory.TransactOpts, _indexerWalletAddress)
}
