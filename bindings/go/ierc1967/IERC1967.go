// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc1967

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

// IERC1967MetaData contains all meta data concerning the IERC1967 contract.
var IERC1967MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// IERC1967ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1967MetaData.ABI instead.
var IERC1967ABI = IERC1967MetaData.ABI

// IERC1967 is an auto generated Go binding around an Ethereum contract.
type IERC1967 struct {
	IERC1967Caller     // Read-only binding to the contract
	IERC1967Transactor // Write-only binding to the contract
	IERC1967Filterer   // Log filterer for contract events
}

// IERC1967Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1967Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1967Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1967Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1967Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1967Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1967Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1967Session struct {
	Contract     *IERC1967         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC1967CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1967CallerSession struct {
	Contract *IERC1967Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IERC1967TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1967TransactorSession struct {
	Contract     *IERC1967Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IERC1967Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1967Raw struct {
	Contract *IERC1967 // Generic contract binding to access the raw methods on
}

// IERC1967CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1967CallerRaw struct {
	Contract *IERC1967Caller // Generic read-only contract binding to access the raw methods on
}

// IERC1967TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1967TransactorRaw struct {
	Contract *IERC1967Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1967 creates a new instance of IERC1967, bound to a specific deployed contract.
func NewIERC1967(address common.Address, backend bind.ContractBackend) (*IERC1967, error) {
	contract, err := bindIERC1967(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1967{IERC1967Caller: IERC1967Caller{contract: contract}, IERC1967Transactor: IERC1967Transactor{contract: contract}, IERC1967Filterer: IERC1967Filterer{contract: contract}}, nil
}

// NewIERC1967Caller creates a new read-only instance of IERC1967, bound to a specific deployed contract.
func NewIERC1967Caller(address common.Address, caller bind.ContractCaller) (*IERC1967Caller, error) {
	contract, err := bindIERC1967(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1967Caller{contract: contract}, nil
}

// NewIERC1967Transactor creates a new write-only instance of IERC1967, bound to a specific deployed contract.
func NewIERC1967Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC1967Transactor, error) {
	contract, err := bindIERC1967(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1967Transactor{contract: contract}, nil
}

// NewIERC1967Filterer creates a new log filterer instance of IERC1967, bound to a specific deployed contract.
func NewIERC1967Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC1967Filterer, error) {
	contract, err := bindIERC1967(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1967Filterer{contract: contract}, nil
}

// bindIERC1967 binds a generic wrapper to an already deployed contract.
func bindIERC1967(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC1967MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1967 *IERC1967Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1967.Contract.IERC1967Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1967 *IERC1967Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1967.Contract.IERC1967Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1967 *IERC1967Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1967.Contract.IERC1967Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1967 *IERC1967CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1967.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1967 *IERC1967TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1967.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1967 *IERC1967TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1967.Contract.contract.Transact(opts, method, params...)
}

// IERC1967AdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the IERC1967 contract.
type IERC1967AdminChangedIterator struct {
	Event *IERC1967AdminChanged // Event containing the contract specifics and raw log

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
func (it *IERC1967AdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1967AdminChanged)
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
		it.Event = new(IERC1967AdminChanged)
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
func (it *IERC1967AdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1967AdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1967AdminChanged represents a AdminChanged event raised by the IERC1967 contract.
type IERC1967AdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_IERC1967 *IERC1967Filterer) FilterAdminChanged(opts *bind.FilterOpts) (*IERC1967AdminChangedIterator, error) {

	logs, sub, err := _IERC1967.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &IERC1967AdminChangedIterator{contract: _IERC1967.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_IERC1967 *IERC1967Filterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *IERC1967AdminChanged) (event.Subscription, error) {

	logs, sub, err := _IERC1967.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1967AdminChanged)
				if err := _IERC1967.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_IERC1967 *IERC1967Filterer) ParseAdminChanged(log types.Log) (*IERC1967AdminChanged, error) {
	event := new(IERC1967AdminChanged)
	if err := _IERC1967.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1967BeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the IERC1967 contract.
type IERC1967BeaconUpgradedIterator struct {
	Event *IERC1967BeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *IERC1967BeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1967BeaconUpgraded)
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
		it.Event = new(IERC1967BeaconUpgraded)
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
func (it *IERC1967BeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1967BeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1967BeaconUpgraded represents a BeaconUpgraded event raised by the IERC1967 contract.
type IERC1967BeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_IERC1967 *IERC1967Filterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*IERC1967BeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _IERC1967.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &IERC1967BeaconUpgradedIterator{contract: _IERC1967.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_IERC1967 *IERC1967Filterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *IERC1967BeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _IERC1967.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1967BeaconUpgraded)
				if err := _IERC1967.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_IERC1967 *IERC1967Filterer) ParseBeaconUpgraded(log types.Log) (*IERC1967BeaconUpgraded, error) {
	event := new(IERC1967BeaconUpgraded)
	if err := _IERC1967.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1967UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the IERC1967 contract.
type IERC1967UpgradedIterator struct {
	Event *IERC1967Upgraded // Event containing the contract specifics and raw log

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
func (it *IERC1967UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1967Upgraded)
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
		it.Event = new(IERC1967Upgraded)
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
func (it *IERC1967UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1967UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1967Upgraded represents a Upgraded event raised by the IERC1967 contract.
type IERC1967Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IERC1967 *IERC1967Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*IERC1967UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IERC1967.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &IERC1967UpgradedIterator{contract: _IERC1967.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IERC1967 *IERC1967Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *IERC1967Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IERC1967.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1967Upgraded)
				if err := _IERC1967.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_IERC1967 *IERC1967Filterer) ParseUpgraded(log types.Log) (*IERC1967Upgraded, error) {
	event := new(IERC1967Upgraded)
	if err := _IERC1967.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
