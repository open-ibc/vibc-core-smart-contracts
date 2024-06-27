// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcreceiverupgradeable

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

// IbcReceiverBaseUpgradeableMetaData contains all meta data concerning the IbcReceiverBaseUpgradeable contract.
var IbcReceiverBaseUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"dispatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ChannelNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedVersion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"notIbcDispatcher\",\"inputs\":[]}]",
}

// IbcReceiverBaseUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcReceiverBaseUpgradeableMetaData.ABI instead.
var IbcReceiverBaseUpgradeableABI = IbcReceiverBaseUpgradeableMetaData.ABI

// IbcReceiverBaseUpgradeable is an auto generated Go binding around an Ethereum contract.
type IbcReceiverBaseUpgradeable struct {
	IbcReceiverBaseUpgradeableCaller     // Read-only binding to the contract
	IbcReceiverBaseUpgradeableTransactor // Write-only binding to the contract
	IbcReceiverBaseUpgradeableFilterer   // Log filterer for contract events
}

// IbcReceiverBaseUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcReceiverBaseUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcReceiverBaseUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcReceiverBaseUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcReceiverBaseUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcReceiverBaseUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcReceiverBaseUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcReceiverBaseUpgradeableSession struct {
	Contract     *IbcReceiverBaseUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IbcReceiverBaseUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcReceiverBaseUpgradeableCallerSession struct {
	Contract *IbcReceiverBaseUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// IbcReceiverBaseUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcReceiverBaseUpgradeableTransactorSession struct {
	Contract     *IbcReceiverBaseUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// IbcReceiverBaseUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcReceiverBaseUpgradeableRaw struct {
	Contract *IbcReceiverBaseUpgradeable // Generic contract binding to access the raw methods on
}

// IbcReceiverBaseUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcReceiverBaseUpgradeableCallerRaw struct {
	Contract *IbcReceiverBaseUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IbcReceiverBaseUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcReceiverBaseUpgradeableTransactorRaw struct {
	Contract *IbcReceiverBaseUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcReceiverBaseUpgradeable creates a new instance of IbcReceiverBaseUpgradeable, bound to a specific deployed contract.
func NewIbcReceiverBaseUpgradeable(address common.Address, backend bind.ContractBackend) (*IbcReceiverBaseUpgradeable, error) {
	contract, err := bindIbcReceiverBaseUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcReceiverBaseUpgradeable{IbcReceiverBaseUpgradeableCaller: IbcReceiverBaseUpgradeableCaller{contract: contract}, IbcReceiverBaseUpgradeableTransactor: IbcReceiverBaseUpgradeableTransactor{contract: contract}, IbcReceiverBaseUpgradeableFilterer: IbcReceiverBaseUpgradeableFilterer{contract: contract}}, nil
}

// NewIbcReceiverBaseUpgradeableCaller creates a new read-only instance of IbcReceiverBaseUpgradeable, bound to a specific deployed contract.
func NewIbcReceiverBaseUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IbcReceiverBaseUpgradeableCaller, error) {
	contract, err := bindIbcReceiverBaseUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcReceiverBaseUpgradeableCaller{contract: contract}, nil
}

// NewIbcReceiverBaseUpgradeableTransactor creates a new write-only instance of IbcReceiverBaseUpgradeable, bound to a specific deployed contract.
func NewIbcReceiverBaseUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcReceiverBaseUpgradeableTransactor, error) {
	contract, err := bindIbcReceiverBaseUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcReceiverBaseUpgradeableTransactor{contract: contract}, nil
}

// NewIbcReceiverBaseUpgradeableFilterer creates a new log filterer instance of IbcReceiverBaseUpgradeable, bound to a specific deployed contract.
func NewIbcReceiverBaseUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcReceiverBaseUpgradeableFilterer, error) {
	contract, err := bindIbcReceiverBaseUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcReceiverBaseUpgradeableFilterer{contract: contract}, nil
}

// bindIbcReceiverBaseUpgradeable binds a generic wrapper to an already deployed contract.
func bindIbcReceiverBaseUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcReceiverBaseUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcReceiverBaseUpgradeable.Contract.IbcReceiverBaseUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcReceiverBaseUpgradeable.Contract.IbcReceiverBaseUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcReceiverBaseUpgradeable.Contract.IbcReceiverBaseUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcReceiverBaseUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcReceiverBaseUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcReceiverBaseUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableCaller) Dispatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IbcReceiverBaseUpgradeable.contract.Call(opts, &out, "dispatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableSession) Dispatcher() (common.Address, error) {
	return _IbcReceiverBaseUpgradeable.Contract.Dispatcher(&_IbcReceiverBaseUpgradeable.CallOpts)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableCallerSession) Dispatcher() (common.Address, error) {
	return _IbcReceiverBaseUpgradeable.Contract.Dispatcher(&_IbcReceiverBaseUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IbcReceiverBaseUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableSession) Owner() (common.Address, error) {
	return _IbcReceiverBaseUpgradeable.Contract.Owner(&_IbcReceiverBaseUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableCallerSession) Owner() (common.Address, error) {
	return _IbcReceiverBaseUpgradeable.Contract.Owner(&_IbcReceiverBaseUpgradeable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcReceiverBaseUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _IbcReceiverBaseUpgradeable.Contract.RenounceOwnership(&_IbcReceiverBaseUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IbcReceiverBaseUpgradeable.Contract.RenounceOwnership(&_IbcReceiverBaseUpgradeable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IbcReceiverBaseUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IbcReceiverBaseUpgradeable.Contract.TransferOwnership(&_IbcReceiverBaseUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IbcReceiverBaseUpgradeable.Contract.TransferOwnership(&_IbcReceiverBaseUpgradeable.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcReceiverBaseUpgradeable.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableSession) Receive() (*types.Transaction, error) {
	return _IbcReceiverBaseUpgradeable.Contract.Receive(&_IbcReceiverBaseUpgradeable.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableTransactorSession) Receive() (*types.Transaction, error) {
	return _IbcReceiverBaseUpgradeable.Contract.Receive(&_IbcReceiverBaseUpgradeable.TransactOpts)
}

// IbcReceiverBaseUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IbcReceiverBaseUpgradeable contract.
type IbcReceiverBaseUpgradeableInitializedIterator struct {
	Event *IbcReceiverBaseUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *IbcReceiverBaseUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcReceiverBaseUpgradeableInitialized)
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
		it.Event = new(IbcReceiverBaseUpgradeableInitialized)
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
func (it *IbcReceiverBaseUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcReceiverBaseUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcReceiverBaseUpgradeableInitialized represents a Initialized event raised by the IbcReceiverBaseUpgradeable contract.
type IbcReceiverBaseUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*IbcReceiverBaseUpgradeableInitializedIterator, error) {

	logs, sub, err := _IbcReceiverBaseUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IbcReceiverBaseUpgradeableInitializedIterator{contract: _IbcReceiverBaseUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IbcReceiverBaseUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _IbcReceiverBaseUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcReceiverBaseUpgradeableInitialized)
				if err := _IbcReceiverBaseUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableFilterer) ParseInitialized(log types.Log) (*IbcReceiverBaseUpgradeableInitialized, error) {
	event := new(IbcReceiverBaseUpgradeableInitialized)
	if err := _IbcReceiverBaseUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcReceiverBaseUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IbcReceiverBaseUpgradeable contract.
type IbcReceiverBaseUpgradeableOwnershipTransferredIterator struct {
	Event *IbcReceiverBaseUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IbcReceiverBaseUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcReceiverBaseUpgradeableOwnershipTransferred)
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
		it.Event = new(IbcReceiverBaseUpgradeableOwnershipTransferred)
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
func (it *IbcReceiverBaseUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcReceiverBaseUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcReceiverBaseUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the IbcReceiverBaseUpgradeable contract.
type IbcReceiverBaseUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IbcReceiverBaseUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IbcReceiverBaseUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IbcReceiverBaseUpgradeableOwnershipTransferredIterator{contract: _IbcReceiverBaseUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IbcReceiverBaseUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IbcReceiverBaseUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcReceiverBaseUpgradeableOwnershipTransferred)
				if err := _IbcReceiverBaseUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IbcReceiverBaseUpgradeable *IbcReceiverBaseUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*IbcReceiverBaseUpgradeableOwnershipTransferred, error) {
	event := new(IbcReceiverBaseUpgradeableOwnershipTransferred)
	if err := _IbcReceiverBaseUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
