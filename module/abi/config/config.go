// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package config

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

// ConfigMetaData contains all meta data concerning the Config contract.
var ConfigMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocationProportion\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_graphStakingAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_grtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_thawingTime\",\"type\":\"uint256\"}],\"name\":\"setConfigInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGraphStakingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getGrtTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getAllocationProportion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getThawingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]",
}

// ConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use ConfigMetaData.ABI instead.
var ConfigABI = ConfigMetaData.ABI

// Config is an auto generated Go binding around an Ethereum contract.
type Config struct {
	ConfigCaller     // Read-only binding to the contract
	ConfigTransactor // Write-only binding to the contract
	ConfigFilterer   // Log filterer for contract events
}

// ConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConfigSession struct {
	Contract     *Config           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConfigCallerSession struct {
	Contract *ConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConfigTransactorSession struct {
	Contract     *ConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConfigRaw struct {
	Contract *Config // Generic contract binding to access the raw methods on
}

// ConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConfigCallerRaw struct {
	Contract *ConfigCaller // Generic read-only contract binding to access the raw methods on
}

// ConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConfigTransactorRaw struct {
	Contract *ConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConfig creates a new instance of Config, bound to a specific deployed contract.
func NewConfig(address common.Address, backend bind.ContractBackend) (*Config, error) {
	contract, err := bindConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Config{ConfigCaller: ConfigCaller{contract: contract}, ConfigTransactor: ConfigTransactor{contract: contract}, ConfigFilterer: ConfigFilterer{contract: contract}}, nil
}

// NewConfigCaller creates a new read-only instance of Config, bound to a specific deployed contract.
func NewConfigCaller(address common.Address, caller bind.ContractCaller) (*ConfigCaller, error) {
	contract, err := bindConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigCaller{contract: contract}, nil
}

// NewConfigTransactor creates a new write-only instance of Config, bound to a specific deployed contract.
func NewConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfigTransactor, error) {
	contract, err := bindConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigTransactor{contract: contract}, nil
}

// NewConfigFilterer creates a new log filterer instance of Config, bound to a specific deployed contract.
func NewConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfigFilterer, error) {
	contract, err := bindConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfigFilterer{contract: contract}, nil
}

// bindConfig binds a generic wrapper to an already deployed contract.
func bindConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Config *ConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Config.Contract.ConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Config *ConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.Contract.ConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Config *ConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Config.Contract.ConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Config *ConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Config.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Config *ConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Config *ConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Config.Contract.contract.Transact(opts, method, params...)
}

// GetAllocationProportion is a free data retrieval call binding the contract method 0x16522e2d.
//
// Solidity: function getAllocationProportion() view returns(uint256)
func (_Config *ConfigCaller) GetAllocationProportion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "getAllocationProportion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllocationProportion is a free data retrieval call binding the contract method 0x16522e2d.
//
// Solidity: function getAllocationProportion() view returns(uint256)
func (_Config *ConfigSession) GetAllocationProportion() (*big.Int, error) {
	return _Config.Contract.GetAllocationProportion(&_Config.CallOpts)
}

// GetAllocationProportion is a free data retrieval call binding the contract method 0x16522e2d.
//
// Solidity: function getAllocationProportion() view returns(uint256)
func (_Config *ConfigCallerSession) GetAllocationProportion() (*big.Int, error) {
	return _Config.Contract.GetAllocationProportion(&_Config.CallOpts)
}

// GetGraphStakingAddress is a free data retrieval call binding the contract method 0x4c94f61a.
//
// Solidity: function getGraphStakingAddress() view returns(address)
func (_Config *ConfigCaller) GetGraphStakingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "getGraphStakingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGraphStakingAddress is a free data retrieval call binding the contract method 0x4c94f61a.
//
// Solidity: function getGraphStakingAddress() view returns(address)
func (_Config *ConfigSession) GetGraphStakingAddress() (common.Address, error) {
	return _Config.Contract.GetGraphStakingAddress(&_Config.CallOpts)
}

// GetGraphStakingAddress is a free data retrieval call binding the contract method 0x4c94f61a.
//
// Solidity: function getGraphStakingAddress() view returns(address)
func (_Config *ConfigCallerSession) GetGraphStakingAddress() (common.Address, error) {
	return _Config.Contract.GetGraphStakingAddress(&_Config.CallOpts)
}

// GetGrtTokenAddress is a free data retrieval call binding the contract method 0x2053065f.
//
// Solidity: function getGrtTokenAddress() view returns(address)
func (_Config *ConfigCaller) GetGrtTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "getGrtTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGrtTokenAddress is a free data retrieval call binding the contract method 0x2053065f.
//
// Solidity: function getGrtTokenAddress() view returns(address)
func (_Config *ConfigSession) GetGrtTokenAddress() (common.Address, error) {
	return _Config.Contract.GetGrtTokenAddress(&_Config.CallOpts)
}

// GetGrtTokenAddress is a free data retrieval call binding the contract method 0x2053065f.
//
// Solidity: function getGrtTokenAddress() view returns(address)
func (_Config *ConfigCallerSession) GetGrtTokenAddress() (common.Address, error) {
	return _Config.Contract.GetGrtTokenAddress(&_Config.CallOpts)
}

// GetThawingTime is a free data retrieval call binding the contract method 0x1d244359.
//
// Solidity: function getThawingTime() view returns(uint256)
func (_Config *ConfigCaller) GetThawingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Config.contract.Call(opts, &out, "getThawingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThawingTime is a free data retrieval call binding the contract method 0x1d244359.
//
// Solidity: function getThawingTime() view returns(uint256)
func (_Config *ConfigSession) GetThawingTime() (*big.Int, error) {
	return _Config.Contract.GetThawingTime(&_Config.CallOpts)
}

// GetThawingTime is a free data retrieval call binding the contract method 0x1d244359.
//
// Solidity: function getThawingTime() view returns(uint256)
func (_Config *ConfigCallerSession) GetThawingTime() (*big.Int, error) {
	return _Config.Contract.GetThawingTime(&_Config.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Config *ConfigTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Config *ConfigSession) Initialize() (*types.Transaction, error) {
	return _Config.Contract.Initialize(&_Config.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Config *ConfigTransactorSession) Initialize() (*types.Transaction, error) {
	return _Config.Contract.Initialize(&_Config.TransactOpts)
}

// SetConfigInfo is a paid mutator transaction binding the contract method 0x16db9d84.
//
// Solidity: function setConfigInfo(uint256 _allocationProportion, address _graphStakingAddress, address _grtTokenAddress, uint256 _thawingTime) returns()
func (_Config *ConfigTransactor) SetConfigInfo(opts *bind.TransactOpts, _allocationProportion *big.Int, _graphStakingAddress common.Address, _grtTokenAddress common.Address, _thawingTime *big.Int) (*types.Transaction, error) {
	return _Config.contract.Transact(opts, "setConfigInfo", _allocationProportion, _graphStakingAddress, _grtTokenAddress, _thawingTime)
}

// SetConfigInfo is a paid mutator transaction binding the contract method 0x16db9d84.
//
// Solidity: function setConfigInfo(uint256 _allocationProportion, address _graphStakingAddress, address _grtTokenAddress, uint256 _thawingTime) returns()
func (_Config *ConfigSession) SetConfigInfo(_allocationProportion *big.Int, _graphStakingAddress common.Address, _grtTokenAddress common.Address, _thawingTime *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetConfigInfo(&_Config.TransactOpts, _allocationProportion, _graphStakingAddress, _grtTokenAddress, _thawingTime)
}

// SetConfigInfo is a paid mutator transaction binding the contract method 0x16db9d84.
//
// Solidity: function setConfigInfo(uint256 _allocationProportion, address _graphStakingAddress, address _grtTokenAddress, uint256 _thawingTime) returns()
func (_Config *ConfigTransactorSession) SetConfigInfo(_allocationProportion *big.Int, _graphStakingAddress common.Address, _grtTokenAddress common.Address, _thawingTime *big.Int) (*types.Transaction, error) {
	return _Config.Contract.SetConfigInfo(&_Config.TransactOpts, _allocationProportion, _graphStakingAddress, _grtTokenAddress, _thawingTime)
}

// ConfigInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Config contract.
type ConfigInitializedIterator struct {
	Event *ConfigInitialized // Event containing the contract specifics and raw log

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
func (it *ConfigInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConfigInitialized)
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
		it.Event = new(ConfigInitialized)
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
func (it *ConfigInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConfigInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConfigInitialized represents a Initialized event raised by the Config contract.
type ConfigInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Config *ConfigFilterer) FilterInitialized(opts *bind.FilterOpts) (*ConfigInitializedIterator, error) {

	logs, sub, err := _Config.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ConfigInitializedIterator{contract: _Config.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Config *ConfigFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ConfigInitialized) (event.Subscription, error) {

	logs, sub, err := _Config.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConfigInitialized)
				if err := _Config.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Config *ConfigFilterer) ParseInitialized(log types.Log) (*ConfigInitialized, error) {
	event := new(ConfigInitialized)
	if err := _Config.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
