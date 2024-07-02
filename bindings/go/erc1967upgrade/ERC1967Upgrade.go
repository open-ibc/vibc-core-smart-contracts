// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc1967upgrade

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
	_ = abi.ConvertType
)

// ERC1967UpgradeMetaData contains all meta data concerning the ERC1967Upgrade contract.
var ERC1967UpgradeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// ERC1967UpgradeABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1967UpgradeMetaData.ABI instead.
var ERC1967UpgradeABI = ERC1967UpgradeMetaData.ABI

// ERC1967Upgrade is an auto generated Go binding around an Ethereum contract.
type ERC1967Upgrade struct {
	ERC1967UpgradeCaller     // Read-only binding to the contract
	ERC1967UpgradeTransactor // Write-only binding to the contract
	ERC1967UpgradeFilterer   // Log filterer for contract events
}

// ERC1967UpgradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1967UpgradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1967UpgradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1967UpgradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1967UpgradeSession struct {
	Contract     *ERC1967Upgrade   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC1967UpgradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1967UpgradeCallerSession struct {
	Contract *ERC1967UpgradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC1967UpgradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1967UpgradeTransactorSession struct {
	Contract     *ERC1967UpgradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC1967UpgradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1967UpgradeRaw struct {
	Contract *ERC1967Upgrade // Generic contract binding to access the raw methods on
}

// ERC1967UpgradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1967UpgradeCallerRaw struct {
	Contract *ERC1967UpgradeCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1967UpgradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1967UpgradeTransactorRaw struct {
	Contract *ERC1967UpgradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1967Upgrade creates a new instance of ERC1967Upgrade, bound to a specific deployed contract.
func NewERC1967Upgrade(address common.Address, backend bind.ContractBackend) (*ERC1967Upgrade, error) {
	contract, err := bindERC1967Upgrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1967Upgrade{ERC1967UpgradeCaller: ERC1967UpgradeCaller{contract: contract}, ERC1967UpgradeTransactor: ERC1967UpgradeTransactor{contract: contract}, ERC1967UpgradeFilterer: ERC1967UpgradeFilterer{contract: contract}}, nil
}

// NewERC1967UpgradeCaller creates a new read-only instance of ERC1967Upgrade, bound to a specific deployed contract.
func NewERC1967UpgradeCaller(address common.Address, caller bind.ContractCaller) (*ERC1967UpgradeCaller, error) {
	contract, err := bindERC1967Upgrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeCaller{contract: contract}, nil
}

// NewERC1967UpgradeTransactor creates a new write-only instance of ERC1967Upgrade, bound to a specific deployed contract.
func NewERC1967UpgradeTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1967UpgradeTransactor, error) {
	contract, err := bindERC1967Upgrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeTransactor{contract: contract}, nil
}

// NewERC1967UpgradeFilterer creates a new log filterer instance of ERC1967Upgrade, bound to a specific deployed contract.
func NewERC1967UpgradeFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1967UpgradeFilterer, error) {
	contract, err := bindERC1967Upgrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeFilterer{contract: contract}, nil
}

// bindERC1967Upgrade binds a generic wrapper to an already deployed contract.
func bindERC1967Upgrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1967UpgradeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967Upgrade *ERC1967UpgradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967Upgrade.Contract.ERC1967UpgradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967Upgrade *ERC1967UpgradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967Upgrade.Contract.ERC1967UpgradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967Upgrade *ERC1967UpgradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967Upgrade.Contract.ERC1967UpgradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967Upgrade *ERC1967UpgradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967Upgrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967Upgrade *ERC1967UpgradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967Upgrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967Upgrade *ERC1967UpgradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967Upgrade.Contract.contract.Transact(opts, method, params...)
}

// ERC1967UpgradeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ERC1967Upgrade contract.
type ERC1967UpgradeAdminChangedIterator struct {
	Event *ERC1967UpgradeAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeAdminChanged)
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
		it.Event = new(ERC1967UpgradeAdminChanged)
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
func (it *ERC1967UpgradeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeAdminChanged represents a AdminChanged event raised by the ERC1967Upgrade contract.
type ERC1967UpgradeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ERC1967UpgradeAdminChangedIterator, error) {

	logs, sub, err := _ERC1967Upgrade.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeAdminChangedIterator{contract: _ERC1967Upgrade.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ERC1967Upgrade.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeAdminChanged)
				if err := _ERC1967Upgrade.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) ParseAdminChanged(log types.Log) (*ERC1967UpgradeAdminChanged, error) {
	event := new(ERC1967UpgradeAdminChanged)
	if err := _ERC1967Upgrade.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ERC1967Upgrade contract.
type ERC1967UpgradeBeaconUpgradedIterator struct {
	Event *ERC1967UpgradeBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeBeaconUpgraded)
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
		it.Event = new(ERC1967UpgradeBeaconUpgraded)
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
func (it *ERC1967UpgradeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeBeaconUpgraded represents a BeaconUpgraded event raised by the ERC1967Upgrade contract.
type ERC1967UpgradeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ERC1967UpgradeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967Upgrade.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeBeaconUpgradedIterator{contract: _ERC1967Upgrade.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967Upgrade.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeBeaconUpgraded)
				if err := _ERC1967Upgrade.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) ParseBeaconUpgraded(log types.Log) (*ERC1967UpgradeBeaconUpgraded, error) {
	event := new(ERC1967UpgradeBeaconUpgraded)
	if err := _ERC1967Upgrade.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC1967Upgrade contract.
type ERC1967UpgradeUpgradedIterator struct {
	Event *ERC1967UpgradeUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgraded)
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
		it.Event = new(ERC1967UpgradeUpgraded)
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
func (it *ERC1967UpgradeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgraded represents a Upgraded event raised by the ERC1967Upgrade contract.
type ERC1967UpgradeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC1967UpgradeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967Upgrade.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradedIterator{contract: _ERC1967Upgrade.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967Upgrade.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgraded)
				if err := _ERC1967Upgrade.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967Upgrade *ERC1967UpgradeFilterer) ParseUpgraded(log types.Log) (*ERC1967UpgradeUpgraded, error) {
	event := new(ERC1967UpgradeUpgraded)
	if err := _ERC1967Upgrade.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
