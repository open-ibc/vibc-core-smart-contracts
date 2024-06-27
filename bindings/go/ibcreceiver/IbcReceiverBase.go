// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcreceiver

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

// IbcReceiverBaseMetaData contains all meta data concerning the IbcReceiverBase contract.
var IbcReceiverBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_dispatcher\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"dispatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ChannelNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedVersion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"notIbcDispatcher\",\"inputs\":[]}]",
}

// IbcReceiverBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcReceiverBaseMetaData.ABI instead.
var IbcReceiverBaseABI = IbcReceiverBaseMetaData.ABI

// IbcReceiverBase is an auto generated Go binding around an Ethereum contract.
type IbcReceiverBase struct {
	IbcReceiverBaseCaller     // Read-only binding to the contract
	IbcReceiverBaseTransactor // Write-only binding to the contract
	IbcReceiverBaseFilterer   // Log filterer for contract events
}

// IbcReceiverBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcReceiverBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcReceiverBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcReceiverBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcReceiverBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcReceiverBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcReceiverBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcReceiverBaseSession struct {
	Contract     *IbcReceiverBase  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcReceiverBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcReceiverBaseCallerSession struct {
	Contract *IbcReceiverBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IbcReceiverBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcReceiverBaseTransactorSession struct {
	Contract     *IbcReceiverBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IbcReceiverBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcReceiverBaseRaw struct {
	Contract *IbcReceiverBase // Generic contract binding to access the raw methods on
}

// IbcReceiverBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcReceiverBaseCallerRaw struct {
	Contract *IbcReceiverBaseCaller // Generic read-only contract binding to access the raw methods on
}

// IbcReceiverBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcReceiverBaseTransactorRaw struct {
	Contract *IbcReceiverBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcReceiverBase creates a new instance of IbcReceiverBase, bound to a specific deployed contract.
func NewIbcReceiverBase(address common.Address, backend bind.ContractBackend) (*IbcReceiverBase, error) {
	contract, err := bindIbcReceiverBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcReceiverBase{IbcReceiverBaseCaller: IbcReceiverBaseCaller{contract: contract}, IbcReceiverBaseTransactor: IbcReceiverBaseTransactor{contract: contract}, IbcReceiverBaseFilterer: IbcReceiverBaseFilterer{contract: contract}}, nil
}

// NewIbcReceiverBaseCaller creates a new read-only instance of IbcReceiverBase, bound to a specific deployed contract.
func NewIbcReceiverBaseCaller(address common.Address, caller bind.ContractCaller) (*IbcReceiverBaseCaller, error) {
	contract, err := bindIbcReceiverBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcReceiverBaseCaller{contract: contract}, nil
}

// NewIbcReceiverBaseTransactor creates a new write-only instance of IbcReceiverBase, bound to a specific deployed contract.
func NewIbcReceiverBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcReceiverBaseTransactor, error) {
	contract, err := bindIbcReceiverBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcReceiverBaseTransactor{contract: contract}, nil
}

// NewIbcReceiverBaseFilterer creates a new log filterer instance of IbcReceiverBase, bound to a specific deployed contract.
func NewIbcReceiverBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcReceiverBaseFilterer, error) {
	contract, err := bindIbcReceiverBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcReceiverBaseFilterer{contract: contract}, nil
}

// bindIbcReceiverBase binds a generic wrapper to an already deployed contract.
func bindIbcReceiverBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcReceiverBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcReceiverBase *IbcReceiverBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcReceiverBase.Contract.IbcReceiverBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcReceiverBase *IbcReceiverBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcReceiverBase.Contract.IbcReceiverBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcReceiverBase *IbcReceiverBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcReceiverBase.Contract.IbcReceiverBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcReceiverBase *IbcReceiverBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcReceiverBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcReceiverBase *IbcReceiverBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcReceiverBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcReceiverBase *IbcReceiverBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcReceiverBase.Contract.contract.Transact(opts, method, params...)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_IbcReceiverBase *IbcReceiverBaseCaller) Dispatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IbcReceiverBase.contract.Call(opts, &out, "dispatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_IbcReceiverBase *IbcReceiverBaseSession) Dispatcher() (common.Address, error) {
	return _IbcReceiverBase.Contract.Dispatcher(&_IbcReceiverBase.CallOpts)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_IbcReceiverBase *IbcReceiverBaseCallerSession) Dispatcher() (common.Address, error) {
	return _IbcReceiverBase.Contract.Dispatcher(&_IbcReceiverBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IbcReceiverBase *IbcReceiverBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IbcReceiverBase.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IbcReceiverBase *IbcReceiverBaseSession) Owner() (common.Address, error) {
	return _IbcReceiverBase.Contract.Owner(&_IbcReceiverBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IbcReceiverBase *IbcReceiverBaseCallerSession) Owner() (common.Address, error) {
	return _IbcReceiverBase.Contract.Owner(&_IbcReceiverBase.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IbcReceiverBase *IbcReceiverBaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcReceiverBase.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IbcReceiverBase *IbcReceiverBaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _IbcReceiverBase.Contract.RenounceOwnership(&_IbcReceiverBase.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IbcReceiverBase *IbcReceiverBaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IbcReceiverBase.Contract.RenounceOwnership(&_IbcReceiverBase.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IbcReceiverBase *IbcReceiverBaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IbcReceiverBase.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IbcReceiverBase *IbcReceiverBaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IbcReceiverBase.Contract.TransferOwnership(&_IbcReceiverBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IbcReceiverBase *IbcReceiverBaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IbcReceiverBase.Contract.TransferOwnership(&_IbcReceiverBase.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IbcReceiverBase *IbcReceiverBaseTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcReceiverBase.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IbcReceiverBase *IbcReceiverBaseSession) Receive() (*types.Transaction, error) {
	return _IbcReceiverBase.Contract.Receive(&_IbcReceiverBase.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IbcReceiverBase *IbcReceiverBaseTransactorSession) Receive() (*types.Transaction, error) {
	return _IbcReceiverBase.Contract.Receive(&_IbcReceiverBase.TransactOpts)
}

// IbcReceiverBaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IbcReceiverBase contract.
type IbcReceiverBaseOwnershipTransferredIterator struct {
	Event *IbcReceiverBaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IbcReceiverBaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcReceiverBaseOwnershipTransferred)
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
		it.Event = new(IbcReceiverBaseOwnershipTransferred)
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
func (it *IbcReceiverBaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcReceiverBaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcReceiverBaseOwnershipTransferred represents a OwnershipTransferred event raised by the IbcReceiverBase contract.
type IbcReceiverBaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IbcReceiverBase *IbcReceiverBaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IbcReceiverBaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IbcReceiverBase.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IbcReceiverBaseOwnershipTransferredIterator{contract: _IbcReceiverBase.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IbcReceiverBase *IbcReceiverBaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IbcReceiverBaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IbcReceiverBase.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcReceiverBaseOwnershipTransferred)
				if err := _IbcReceiverBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IbcReceiverBase *IbcReceiverBaseFilterer) ParseOwnershipTransferred(log types.Log) (*IbcReceiverBaseOwnershipTransferred, error) {
	event := new(IbcReceiverBaseOwnershipTransferred)
	if err := _IbcReceiverBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
