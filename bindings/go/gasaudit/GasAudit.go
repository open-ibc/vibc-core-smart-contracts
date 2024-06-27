// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gasaudit

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

// GasAuditMetaData contains all meta data concerning the GasAudit contract.
var GasAuditMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"callWithBytes32\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"coutnerpartyChannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"callWithString\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"coutnerpartyChannelId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelIds1\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"channelIds2\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OpenIbcChannel1\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"coutnerpartyChannelId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OpenIbcChannel2\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"coutnerpartyChannelId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]",
}

// GasAuditABI is the input ABI used to generate the binding from.
// Deprecated: Use GasAuditMetaData.ABI instead.
var GasAuditABI = GasAuditMetaData.ABI

// GasAudit is an auto generated Go binding around an Ethereum contract.
type GasAudit struct {
	GasAuditCaller     // Read-only binding to the contract
	GasAuditTransactor // Write-only binding to the contract
	GasAuditFilterer   // Log filterer for contract events
}

// GasAuditCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasAuditCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasAuditTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasAuditTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasAuditFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasAuditFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasAuditSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasAuditSession struct {
	Contract     *GasAudit         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasAuditCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasAuditCallerSession struct {
	Contract *GasAuditCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GasAuditTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasAuditTransactorSession struct {
	Contract     *GasAuditTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GasAuditRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasAuditRaw struct {
	Contract *GasAudit // Generic contract binding to access the raw methods on
}

// GasAuditCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasAuditCallerRaw struct {
	Contract *GasAuditCaller // Generic read-only contract binding to access the raw methods on
}

// GasAuditTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasAuditTransactorRaw struct {
	Contract *GasAuditTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasAudit creates a new instance of GasAudit, bound to a specific deployed contract.
func NewGasAudit(address common.Address, backend bind.ContractBackend) (*GasAudit, error) {
	contract, err := bindGasAudit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasAudit{GasAuditCaller: GasAuditCaller{contract: contract}, GasAuditTransactor: GasAuditTransactor{contract: contract}, GasAuditFilterer: GasAuditFilterer{contract: contract}}, nil
}

// NewGasAuditCaller creates a new read-only instance of GasAudit, bound to a specific deployed contract.
func NewGasAuditCaller(address common.Address, caller bind.ContractCaller) (*GasAuditCaller, error) {
	contract, err := bindGasAudit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasAuditCaller{contract: contract}, nil
}

// NewGasAuditTransactor creates a new write-only instance of GasAudit, bound to a specific deployed contract.
func NewGasAuditTransactor(address common.Address, transactor bind.ContractTransactor) (*GasAuditTransactor, error) {
	contract, err := bindGasAudit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasAuditTransactor{contract: contract}, nil
}

// NewGasAuditFilterer creates a new log filterer instance of GasAudit, bound to a specific deployed contract.
func NewGasAuditFilterer(address common.Address, filterer bind.ContractFilterer) (*GasAuditFilterer, error) {
	contract, err := bindGasAudit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasAuditFilterer{contract: contract}, nil
}

// bindGasAudit binds a generic wrapper to an already deployed contract.
func bindGasAudit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GasAuditMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasAudit *GasAuditRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasAudit.Contract.GasAuditCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasAudit *GasAuditRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasAudit.Contract.GasAuditTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasAudit *GasAuditRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasAudit.Contract.GasAuditTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasAudit *GasAuditCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasAudit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasAudit *GasAuditTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasAudit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasAudit *GasAuditTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasAudit.Contract.contract.Transact(opts, method, params...)
}

// ChannelIds1 is a free data retrieval call binding the contract method 0x5160f9e7.
//
// Solidity: function channelIds1(bytes32 ) view returns(bool)
func (_GasAudit *GasAuditCaller) ChannelIds1(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _GasAudit.contract.Call(opts, &out, "channelIds1", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChannelIds1 is a free data retrieval call binding the contract method 0x5160f9e7.
//
// Solidity: function channelIds1(bytes32 ) view returns(bool)
func (_GasAudit *GasAuditSession) ChannelIds1(arg0 [32]byte) (bool, error) {
	return _GasAudit.Contract.ChannelIds1(&_GasAudit.CallOpts, arg0)
}

// ChannelIds1 is a free data retrieval call binding the contract method 0x5160f9e7.
//
// Solidity: function channelIds1(bytes32 ) view returns(bool)
func (_GasAudit *GasAuditCallerSession) ChannelIds1(arg0 [32]byte) (bool, error) {
	return _GasAudit.Contract.ChannelIds1(&_GasAudit.CallOpts, arg0)
}

// ChannelIds2 is a free data retrieval call binding the contract method 0x187c35e5.
//
// Solidity: function channelIds2(string ) view returns(bool)
func (_GasAudit *GasAuditCaller) ChannelIds2(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _GasAudit.contract.Call(opts, &out, "channelIds2", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChannelIds2 is a free data retrieval call binding the contract method 0x187c35e5.
//
// Solidity: function channelIds2(string ) view returns(bool)
func (_GasAudit *GasAuditSession) ChannelIds2(arg0 string) (bool, error) {
	return _GasAudit.Contract.ChannelIds2(&_GasAudit.CallOpts, arg0)
}

// ChannelIds2 is a free data retrieval call binding the contract method 0x187c35e5.
//
// Solidity: function channelIds2(string ) view returns(bool)
func (_GasAudit *GasAuditCallerSession) ChannelIds2(arg0 string) (bool, error) {
	return _GasAudit.Contract.ChannelIds2(&_GasAudit.CallOpts, arg0)
}

// CallWithBytes32 is a paid mutator transaction binding the contract method 0xfca6eb29.
//
// Solidity: function callWithBytes32(address portAddress, bytes32 channelId, string counterpartyPortId, bytes32 coutnerpartyChannelId) returns()
func (_GasAudit *GasAuditTransactor) CallWithBytes32(opts *bind.TransactOpts, portAddress common.Address, channelId [32]byte, counterpartyPortId string, coutnerpartyChannelId [32]byte) (*types.Transaction, error) {
	return _GasAudit.contract.Transact(opts, "callWithBytes32", portAddress, channelId, counterpartyPortId, coutnerpartyChannelId)
}

// CallWithBytes32 is a paid mutator transaction binding the contract method 0xfca6eb29.
//
// Solidity: function callWithBytes32(address portAddress, bytes32 channelId, string counterpartyPortId, bytes32 coutnerpartyChannelId) returns()
func (_GasAudit *GasAuditSession) CallWithBytes32(portAddress common.Address, channelId [32]byte, counterpartyPortId string, coutnerpartyChannelId [32]byte) (*types.Transaction, error) {
	return _GasAudit.Contract.CallWithBytes32(&_GasAudit.TransactOpts, portAddress, channelId, counterpartyPortId, coutnerpartyChannelId)
}

// CallWithBytes32 is a paid mutator transaction binding the contract method 0xfca6eb29.
//
// Solidity: function callWithBytes32(address portAddress, bytes32 channelId, string counterpartyPortId, bytes32 coutnerpartyChannelId) returns()
func (_GasAudit *GasAuditTransactorSession) CallWithBytes32(portAddress common.Address, channelId [32]byte, counterpartyPortId string, coutnerpartyChannelId [32]byte) (*types.Transaction, error) {
	return _GasAudit.Contract.CallWithBytes32(&_GasAudit.TransactOpts, portAddress, channelId, counterpartyPortId, coutnerpartyChannelId)
}

// CallWithString is a paid mutator transaction binding the contract method 0x485b39b5.
//
// Solidity: function callWithString(address portAddress, string channelId, string counterpartyPortId, string coutnerpartyChannelId) returns()
func (_GasAudit *GasAuditTransactor) CallWithString(opts *bind.TransactOpts, portAddress common.Address, channelId string, counterpartyPortId string, coutnerpartyChannelId string) (*types.Transaction, error) {
	return _GasAudit.contract.Transact(opts, "callWithString", portAddress, channelId, counterpartyPortId, coutnerpartyChannelId)
}

// CallWithString is a paid mutator transaction binding the contract method 0x485b39b5.
//
// Solidity: function callWithString(address portAddress, string channelId, string counterpartyPortId, string coutnerpartyChannelId) returns()
func (_GasAudit *GasAuditSession) CallWithString(portAddress common.Address, channelId string, counterpartyPortId string, coutnerpartyChannelId string) (*types.Transaction, error) {
	return _GasAudit.Contract.CallWithString(&_GasAudit.TransactOpts, portAddress, channelId, counterpartyPortId, coutnerpartyChannelId)
}

// CallWithString is a paid mutator transaction binding the contract method 0x485b39b5.
//
// Solidity: function callWithString(address portAddress, string channelId, string counterpartyPortId, string coutnerpartyChannelId) returns()
func (_GasAudit *GasAuditTransactorSession) CallWithString(portAddress common.Address, channelId string, counterpartyPortId string, coutnerpartyChannelId string) (*types.Transaction, error) {
	return _GasAudit.Contract.CallWithString(&_GasAudit.TransactOpts, portAddress, channelId, counterpartyPortId, coutnerpartyChannelId)
}

// GasAuditOpenIbcChannel1Iterator is returned from FilterOpenIbcChannel1 and is used to iterate over the raw logs and unpacked data for OpenIbcChannel1 events raised by the GasAudit contract.
type GasAuditOpenIbcChannel1Iterator struct {
	Event *GasAuditOpenIbcChannel1 // Event containing the contract specifics and raw log

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
func (it *GasAuditOpenIbcChannel1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasAuditOpenIbcChannel1)
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
		it.Event = new(GasAuditOpenIbcChannel1)
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
func (it *GasAuditOpenIbcChannel1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasAuditOpenIbcChannel1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasAuditOpenIbcChannel1 represents a OpenIbcChannel1 event raised by the GasAudit contract.
type GasAuditOpenIbcChannel1 struct {
	PortAddress           common.Address
	ChannelId             [32]byte
	CounterpartyPortId    string
	CoutnerpartyChannelId [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterOpenIbcChannel1 is a free log retrieval operation binding the contract event 0x9bfa7e679cc290246349f80646c1c53a881353ed5b17ad82b5eb6ba283718155.
//
// Solidity: event OpenIbcChannel1(address indexed portAddress, bytes32 indexed channelId, string counterpartyPortId, bytes32 coutnerpartyChannelId)
func (_GasAudit *GasAuditFilterer) FilterOpenIbcChannel1(opts *bind.FilterOpts, portAddress []common.Address, channelId [][32]byte) (*GasAuditOpenIbcChannel1Iterator, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _GasAudit.contract.FilterLogs(opts, "OpenIbcChannel1", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &GasAuditOpenIbcChannel1Iterator{contract: _GasAudit.contract, event: "OpenIbcChannel1", logs: logs, sub: sub}, nil
}

// WatchOpenIbcChannel1 is a free log subscription operation binding the contract event 0x9bfa7e679cc290246349f80646c1c53a881353ed5b17ad82b5eb6ba283718155.
//
// Solidity: event OpenIbcChannel1(address indexed portAddress, bytes32 indexed channelId, string counterpartyPortId, bytes32 coutnerpartyChannelId)
func (_GasAudit *GasAuditFilterer) WatchOpenIbcChannel1(opts *bind.WatchOpts, sink chan<- *GasAuditOpenIbcChannel1, portAddress []common.Address, channelId [][32]byte) (event.Subscription, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _GasAudit.contract.WatchLogs(opts, "OpenIbcChannel1", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasAuditOpenIbcChannel1)
				if err := _GasAudit.contract.UnpackLog(event, "OpenIbcChannel1", log); err != nil {
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

// ParseOpenIbcChannel1 is a log parse operation binding the contract event 0x9bfa7e679cc290246349f80646c1c53a881353ed5b17ad82b5eb6ba283718155.
//
// Solidity: event OpenIbcChannel1(address indexed portAddress, bytes32 indexed channelId, string counterpartyPortId, bytes32 coutnerpartyChannelId)
func (_GasAudit *GasAuditFilterer) ParseOpenIbcChannel1(log types.Log) (*GasAuditOpenIbcChannel1, error) {
	event := new(GasAuditOpenIbcChannel1)
	if err := _GasAudit.contract.UnpackLog(event, "OpenIbcChannel1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasAuditOpenIbcChannel2Iterator is returned from FilterOpenIbcChannel2 and is used to iterate over the raw logs and unpacked data for OpenIbcChannel2 events raised by the GasAudit contract.
type GasAuditOpenIbcChannel2Iterator struct {
	Event *GasAuditOpenIbcChannel2 // Event containing the contract specifics and raw log

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
func (it *GasAuditOpenIbcChannel2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasAuditOpenIbcChannel2)
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
		it.Event = new(GasAuditOpenIbcChannel2)
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
func (it *GasAuditOpenIbcChannel2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasAuditOpenIbcChannel2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasAuditOpenIbcChannel2 represents a OpenIbcChannel2 event raised by the GasAudit contract.
type GasAuditOpenIbcChannel2 struct {
	PortAddress           common.Address
	ChannelId             string
	CounterpartyPortId    string
	CoutnerpartyChannelId string
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterOpenIbcChannel2 is a free log retrieval operation binding the contract event 0x57ddcf2e4fd363065bccadcdb56928237fb413c06e42c021b0dca4a3b8e58051.
//
// Solidity: event OpenIbcChannel2(address indexed portAddress, string channelId, string counterpartyPortId, string coutnerpartyChannelId)
func (_GasAudit *GasAuditFilterer) FilterOpenIbcChannel2(opts *bind.FilterOpts, portAddress []common.Address) (*GasAuditOpenIbcChannel2Iterator, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}

	logs, sub, err := _GasAudit.contract.FilterLogs(opts, "OpenIbcChannel2", portAddressRule)
	if err != nil {
		return nil, err
	}
	return &GasAuditOpenIbcChannel2Iterator{contract: _GasAudit.contract, event: "OpenIbcChannel2", logs: logs, sub: sub}, nil
}

// WatchOpenIbcChannel2 is a free log subscription operation binding the contract event 0x57ddcf2e4fd363065bccadcdb56928237fb413c06e42c021b0dca4a3b8e58051.
//
// Solidity: event OpenIbcChannel2(address indexed portAddress, string channelId, string counterpartyPortId, string coutnerpartyChannelId)
func (_GasAudit *GasAuditFilterer) WatchOpenIbcChannel2(opts *bind.WatchOpts, sink chan<- *GasAuditOpenIbcChannel2, portAddress []common.Address) (event.Subscription, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}

	logs, sub, err := _GasAudit.contract.WatchLogs(opts, "OpenIbcChannel2", portAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasAuditOpenIbcChannel2)
				if err := _GasAudit.contract.UnpackLog(event, "OpenIbcChannel2", log); err != nil {
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

// ParseOpenIbcChannel2 is a log parse operation binding the contract event 0x57ddcf2e4fd363065bccadcdb56928237fb413c06e42c021b0dca4a3b8e58051.
//
// Solidity: event OpenIbcChannel2(address indexed portAddress, string channelId, string counterpartyPortId, string coutnerpartyChannelId)
func (_GasAudit *GasAuditFilterer) ParseOpenIbcChannel2(log types.Log) (*GasAuditOpenIbcChannel2, error) {
	event := new(GasAuditOpenIbcChannel2)
	if err := _GasAudit.contract.UnpackLog(event, "OpenIbcChannel2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
