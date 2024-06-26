// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcmiddleware

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

// IbcMwUserMetaData contains all meta data concerning the IbcMwUser contract.
var IbcMwUserMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_middleware\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"authorizeMiddleware\",\"inputs\":[{\"name\":\"middleware\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authorizedMws\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mw\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultMw\",\"inputs\":[{\"name\":\"_middleware\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"UnauthorizedIbcMiddleware\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ackAddressMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ackDataTooShort\",\"inputs\":[]}]",
}

// IbcMwUserABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcMwUserMetaData.ABI instead.
var IbcMwUserABI = IbcMwUserMetaData.ABI

// IbcMwUser is an auto generated Go binding around an Ethereum contract.
type IbcMwUser struct {
	IbcMwUserCaller     // Read-only binding to the contract
	IbcMwUserTransactor // Write-only binding to the contract
	IbcMwUserFilterer   // Log filterer for contract events
}

// IbcMwUserCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcMwUserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMwUserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcMwUserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMwUserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcMwUserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMwUserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcMwUserSession struct {
	Contract     *IbcMwUser        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcMwUserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcMwUserCallerSession struct {
	Contract *IbcMwUserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IbcMwUserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcMwUserTransactorSession struct {
	Contract     *IbcMwUserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IbcMwUserRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcMwUserRaw struct {
	Contract *IbcMwUser // Generic contract binding to access the raw methods on
}

// IbcMwUserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcMwUserCallerRaw struct {
	Contract *IbcMwUserCaller // Generic read-only contract binding to access the raw methods on
}

// IbcMwUserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcMwUserTransactorRaw struct {
	Contract *IbcMwUserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcMwUser creates a new instance of IbcMwUser, bound to a specific deployed contract.
func NewIbcMwUser(address common.Address, backend bind.ContractBackend) (*IbcMwUser, error) {
	contract, err := bindIbcMwUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcMwUser{IbcMwUserCaller: IbcMwUserCaller{contract: contract}, IbcMwUserTransactor: IbcMwUserTransactor{contract: contract}, IbcMwUserFilterer: IbcMwUserFilterer{contract: contract}}, nil
}

// NewIbcMwUserCaller creates a new read-only instance of IbcMwUser, bound to a specific deployed contract.
func NewIbcMwUserCaller(address common.Address, caller bind.ContractCaller) (*IbcMwUserCaller, error) {
	contract, err := bindIbcMwUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcMwUserCaller{contract: contract}, nil
}

// NewIbcMwUserTransactor creates a new write-only instance of IbcMwUser, bound to a specific deployed contract.
func NewIbcMwUserTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcMwUserTransactor, error) {
	contract, err := bindIbcMwUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcMwUserTransactor{contract: contract}, nil
}

// NewIbcMwUserFilterer creates a new log filterer instance of IbcMwUser, bound to a specific deployed contract.
func NewIbcMwUserFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcMwUserFilterer, error) {
	contract, err := bindIbcMwUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcMwUserFilterer{contract: contract}, nil
}

// bindIbcMwUser binds a generic wrapper to an already deployed contract.
func bindIbcMwUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcMwUserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcMwUser *IbcMwUserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcMwUser.Contract.IbcMwUserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcMwUser *IbcMwUserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcMwUser.Contract.IbcMwUserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcMwUser *IbcMwUserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcMwUser.Contract.IbcMwUserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcMwUser *IbcMwUserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcMwUser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcMwUser *IbcMwUserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcMwUser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcMwUser *IbcMwUserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcMwUser.Contract.contract.Transact(opts, method, params...)
}

// AuthorizedMws is a free data retrieval call binding the contract method 0x3a7fbcbd.
//
// Solidity: function authorizedMws(address ) view returns(bool)
func (_IbcMwUser *IbcMwUserCaller) AuthorizedMws(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _IbcMwUser.contract.Call(opts, &out, "authorizedMws", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedMws is a free data retrieval call binding the contract method 0x3a7fbcbd.
//
// Solidity: function authorizedMws(address ) view returns(bool)
func (_IbcMwUser *IbcMwUserSession) AuthorizedMws(arg0 common.Address) (bool, error) {
	return _IbcMwUser.Contract.AuthorizedMws(&_IbcMwUser.CallOpts, arg0)
}

// AuthorizedMws is a free data retrieval call binding the contract method 0x3a7fbcbd.
//
// Solidity: function authorizedMws(address ) view returns(bool)
func (_IbcMwUser *IbcMwUserCallerSession) AuthorizedMws(arg0 common.Address) (bool, error) {
	return _IbcMwUser.Contract.AuthorizedMws(&_IbcMwUser.CallOpts, arg0)
}

// Mw is a free data retrieval call binding the contract method 0xa742d78c.
//
// Solidity: function mw() view returns(address)
func (_IbcMwUser *IbcMwUserCaller) Mw(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IbcMwUser.contract.Call(opts, &out, "mw")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mw is a free data retrieval call binding the contract method 0xa742d78c.
//
// Solidity: function mw() view returns(address)
func (_IbcMwUser *IbcMwUserSession) Mw() (common.Address, error) {
	return _IbcMwUser.Contract.Mw(&_IbcMwUser.CallOpts)
}

// Mw is a free data retrieval call binding the contract method 0xa742d78c.
//
// Solidity: function mw() view returns(address)
func (_IbcMwUser *IbcMwUserCallerSession) Mw() (common.Address, error) {
	return _IbcMwUser.Contract.Mw(&_IbcMwUser.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IbcMwUser *IbcMwUserCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IbcMwUser.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IbcMwUser *IbcMwUserSession) Owner() (common.Address, error) {
	return _IbcMwUser.Contract.Owner(&_IbcMwUser.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IbcMwUser *IbcMwUserCallerSession) Owner() (common.Address, error) {
	return _IbcMwUser.Contract.Owner(&_IbcMwUser.CallOpts)
}

// AuthorizeMiddleware is a paid mutator transaction binding the contract method 0x3b90b042.
//
// Solidity: function authorizeMiddleware(address middleware) returns()
func (_IbcMwUser *IbcMwUserTransactor) AuthorizeMiddleware(opts *bind.TransactOpts, middleware common.Address) (*types.Transaction, error) {
	return _IbcMwUser.contract.Transact(opts, "authorizeMiddleware", middleware)
}

// AuthorizeMiddleware is a paid mutator transaction binding the contract method 0x3b90b042.
//
// Solidity: function authorizeMiddleware(address middleware) returns()
func (_IbcMwUser *IbcMwUserSession) AuthorizeMiddleware(middleware common.Address) (*types.Transaction, error) {
	return _IbcMwUser.Contract.AuthorizeMiddleware(&_IbcMwUser.TransactOpts, middleware)
}

// AuthorizeMiddleware is a paid mutator transaction binding the contract method 0x3b90b042.
//
// Solidity: function authorizeMiddleware(address middleware) returns()
func (_IbcMwUser *IbcMwUserTransactorSession) AuthorizeMiddleware(middleware common.Address) (*types.Transaction, error) {
	return _IbcMwUser.Contract.AuthorizeMiddleware(&_IbcMwUser.TransactOpts, middleware)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IbcMwUser *IbcMwUserTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcMwUser.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IbcMwUser *IbcMwUserSession) RenounceOwnership() (*types.Transaction, error) {
	return _IbcMwUser.Contract.RenounceOwnership(&_IbcMwUser.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IbcMwUser *IbcMwUserTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IbcMwUser.Contract.RenounceOwnership(&_IbcMwUser.TransactOpts)
}

// SetDefaultMw is a paid mutator transaction binding the contract method 0x00e82cef.
//
// Solidity: function setDefaultMw(address _middleware) returns()
func (_IbcMwUser *IbcMwUserTransactor) SetDefaultMw(opts *bind.TransactOpts, _middleware common.Address) (*types.Transaction, error) {
	return _IbcMwUser.contract.Transact(opts, "setDefaultMw", _middleware)
}

// SetDefaultMw is a paid mutator transaction binding the contract method 0x00e82cef.
//
// Solidity: function setDefaultMw(address _middleware) returns()
func (_IbcMwUser *IbcMwUserSession) SetDefaultMw(_middleware common.Address) (*types.Transaction, error) {
	return _IbcMwUser.Contract.SetDefaultMw(&_IbcMwUser.TransactOpts, _middleware)
}

// SetDefaultMw is a paid mutator transaction binding the contract method 0x00e82cef.
//
// Solidity: function setDefaultMw(address _middleware) returns()
func (_IbcMwUser *IbcMwUserTransactorSession) SetDefaultMw(_middleware common.Address) (*types.Transaction, error) {
	return _IbcMwUser.Contract.SetDefaultMw(&_IbcMwUser.TransactOpts, _middleware)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IbcMwUser *IbcMwUserTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IbcMwUser.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IbcMwUser *IbcMwUserSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IbcMwUser.Contract.TransferOwnership(&_IbcMwUser.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IbcMwUser *IbcMwUserTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IbcMwUser.Contract.TransferOwnership(&_IbcMwUser.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IbcMwUser *IbcMwUserTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcMwUser.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IbcMwUser *IbcMwUserSession) Receive() (*types.Transaction, error) {
	return _IbcMwUser.Contract.Receive(&_IbcMwUser.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IbcMwUser *IbcMwUserTransactorSession) Receive() (*types.Transaction, error) {
	return _IbcMwUser.Contract.Receive(&_IbcMwUser.TransactOpts)
}

// IbcMwUserOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IbcMwUser contract.
type IbcMwUserOwnershipTransferredIterator struct {
	Event *IbcMwUserOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IbcMwUserOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcMwUserOwnershipTransferred)
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
		it.Event = new(IbcMwUserOwnershipTransferred)
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
func (it *IbcMwUserOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcMwUserOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcMwUserOwnershipTransferred represents a OwnershipTransferred event raised by the IbcMwUser contract.
type IbcMwUserOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IbcMwUser *IbcMwUserFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IbcMwUserOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IbcMwUser.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IbcMwUserOwnershipTransferredIterator{contract: _IbcMwUser.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IbcMwUser *IbcMwUserFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IbcMwUserOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IbcMwUser.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcMwUserOwnershipTransferred)
				if err := _IbcMwUser.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IbcMwUser *IbcMwUserFilterer) ParseOwnershipTransferred(log types.Log) (*IbcMwUserOwnershipTransferred, error) {
	event := new(IbcMwUserOwnershipTransferred)
	if err := _IbcMwUser.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
