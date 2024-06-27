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

// AckPacket is an auto generated low-level Go binding around an user-defined struct.
type AckPacket struct {
	Success bool
	Data    []byte
}

// IbcMwEventsEmitterMetaData contains all meta data concerning the IbcMwEventsEmitter contract.
var IbcMwEventsEmitterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"RecvMWAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"mwId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"appData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mwData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"ack\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecvMWPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"mwId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"appData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mwData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecvMWTimeout\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"mwId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"appData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mwData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendMWPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"mwId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"appData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"mwData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
}

// IbcMwEventsEmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcMwEventsEmitterMetaData.ABI instead.
var IbcMwEventsEmitterABI = IbcMwEventsEmitterMetaData.ABI

// IbcMwEventsEmitter is an auto generated Go binding around an Ethereum contract.
type IbcMwEventsEmitter struct {
	IbcMwEventsEmitterCaller     // Read-only binding to the contract
	IbcMwEventsEmitterTransactor // Write-only binding to the contract
	IbcMwEventsEmitterFilterer   // Log filterer for contract events
}

// IbcMwEventsEmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcMwEventsEmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMwEventsEmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcMwEventsEmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMwEventsEmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcMwEventsEmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMwEventsEmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcMwEventsEmitterSession struct {
	Contract     *IbcMwEventsEmitter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IbcMwEventsEmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcMwEventsEmitterCallerSession struct {
	Contract *IbcMwEventsEmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IbcMwEventsEmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcMwEventsEmitterTransactorSession struct {
	Contract     *IbcMwEventsEmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IbcMwEventsEmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcMwEventsEmitterRaw struct {
	Contract *IbcMwEventsEmitter // Generic contract binding to access the raw methods on
}

// IbcMwEventsEmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcMwEventsEmitterCallerRaw struct {
	Contract *IbcMwEventsEmitterCaller // Generic read-only contract binding to access the raw methods on
}

// IbcMwEventsEmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcMwEventsEmitterTransactorRaw struct {
	Contract *IbcMwEventsEmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcMwEventsEmitter creates a new instance of IbcMwEventsEmitter, bound to a specific deployed contract.
func NewIbcMwEventsEmitter(address common.Address, backend bind.ContractBackend) (*IbcMwEventsEmitter, error) {
	contract, err := bindIbcMwEventsEmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcMwEventsEmitter{IbcMwEventsEmitterCaller: IbcMwEventsEmitterCaller{contract: contract}, IbcMwEventsEmitterTransactor: IbcMwEventsEmitterTransactor{contract: contract}, IbcMwEventsEmitterFilterer: IbcMwEventsEmitterFilterer{contract: contract}}, nil
}

// NewIbcMwEventsEmitterCaller creates a new read-only instance of IbcMwEventsEmitter, bound to a specific deployed contract.
func NewIbcMwEventsEmitterCaller(address common.Address, caller bind.ContractCaller) (*IbcMwEventsEmitterCaller, error) {
	contract, err := bindIbcMwEventsEmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcMwEventsEmitterCaller{contract: contract}, nil
}

// NewIbcMwEventsEmitterTransactor creates a new write-only instance of IbcMwEventsEmitter, bound to a specific deployed contract.
func NewIbcMwEventsEmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcMwEventsEmitterTransactor, error) {
	contract, err := bindIbcMwEventsEmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcMwEventsEmitterTransactor{contract: contract}, nil
}

// NewIbcMwEventsEmitterFilterer creates a new log filterer instance of IbcMwEventsEmitter, bound to a specific deployed contract.
func NewIbcMwEventsEmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcMwEventsEmitterFilterer, error) {
	contract, err := bindIbcMwEventsEmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcMwEventsEmitterFilterer{contract: contract}, nil
}

// bindIbcMwEventsEmitter binds a generic wrapper to an already deployed contract.
func bindIbcMwEventsEmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcMwEventsEmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcMwEventsEmitter *IbcMwEventsEmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcMwEventsEmitter.Contract.IbcMwEventsEmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcMwEventsEmitter *IbcMwEventsEmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcMwEventsEmitter.Contract.IbcMwEventsEmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcMwEventsEmitter *IbcMwEventsEmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcMwEventsEmitter.Contract.IbcMwEventsEmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcMwEventsEmitter *IbcMwEventsEmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcMwEventsEmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcMwEventsEmitter *IbcMwEventsEmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcMwEventsEmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcMwEventsEmitter *IbcMwEventsEmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcMwEventsEmitter.Contract.contract.Transact(opts, method, params...)
}

// IbcMwEventsEmitterRecvMWAckIterator is returned from FilterRecvMWAck and is used to iterate over the raw logs and unpacked data for RecvMWAck events raised by the IbcMwEventsEmitter contract.
type IbcMwEventsEmitterRecvMWAckIterator struct {
	Event *IbcMwEventsEmitterRecvMWAck // Event containing the contract specifics and raw log

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
func (it *IbcMwEventsEmitterRecvMWAckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcMwEventsEmitterRecvMWAck)
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
		it.Event = new(IbcMwEventsEmitterRecvMWAck)
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
func (it *IbcMwEventsEmitterRecvMWAckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcMwEventsEmitterRecvMWAckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcMwEventsEmitterRecvMWAck represents a RecvMWAck event raised by the IbcMwEventsEmitter contract.
type IbcMwEventsEmitterRecvMWAck struct {
	ChannelId    [32]byte
	SrcPortAddr  [32]byte
	DestPortAddr [32]byte
	MwId         *big.Int
	AppData      []byte
	MwData       []byte
	Ack          AckPacket
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRecvMWAck is a free log retrieval operation binding the contract event 0x0ad1351cd77bd217ef00c2ab94f17283f9a51d4ebd8d189c151760d8075f01c2.
//
// Solidity: event RecvMWAck(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, bytes mwData, (bool,bytes) ack)
func (_IbcMwEventsEmitter *IbcMwEventsEmitterFilterer) FilterRecvMWAck(opts *bind.FilterOpts, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (*IbcMwEventsEmitterRecvMWAckIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var srcPortAddrRule []interface{}
	for _, srcPortAddrItem := range srcPortAddr {
		srcPortAddrRule = append(srcPortAddrRule, srcPortAddrItem)
	}
	var destPortAddrRule []interface{}
	for _, destPortAddrItem := range destPortAddr {
		destPortAddrRule = append(destPortAddrRule, destPortAddrItem)
	}

	logs, sub, err := _IbcMwEventsEmitter.contract.FilterLogs(opts, "RecvMWAck", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return &IbcMwEventsEmitterRecvMWAckIterator{contract: _IbcMwEventsEmitter.contract, event: "RecvMWAck", logs: logs, sub: sub}, nil
}

// WatchRecvMWAck is a free log subscription operation binding the contract event 0x0ad1351cd77bd217ef00c2ab94f17283f9a51d4ebd8d189c151760d8075f01c2.
//
// Solidity: event RecvMWAck(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, bytes mwData, (bool,bytes) ack)
func (_IbcMwEventsEmitter *IbcMwEventsEmitterFilterer) WatchRecvMWAck(opts *bind.WatchOpts, sink chan<- *IbcMwEventsEmitterRecvMWAck, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var srcPortAddrRule []interface{}
	for _, srcPortAddrItem := range srcPortAddr {
		srcPortAddrRule = append(srcPortAddrRule, srcPortAddrItem)
	}
	var destPortAddrRule []interface{}
	for _, destPortAddrItem := range destPortAddr {
		destPortAddrRule = append(destPortAddrRule, destPortAddrItem)
	}

	logs, sub, err := _IbcMwEventsEmitter.contract.WatchLogs(opts, "RecvMWAck", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcMwEventsEmitterRecvMWAck)
				if err := _IbcMwEventsEmitter.contract.UnpackLog(event, "RecvMWAck", log); err != nil {
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

// ParseRecvMWAck is a log parse operation binding the contract event 0x0ad1351cd77bd217ef00c2ab94f17283f9a51d4ebd8d189c151760d8075f01c2.
//
// Solidity: event RecvMWAck(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, bytes mwData, (bool,bytes) ack)
func (_IbcMwEventsEmitter *IbcMwEventsEmitterFilterer) ParseRecvMWAck(log types.Log) (*IbcMwEventsEmitterRecvMWAck, error) {
	event := new(IbcMwEventsEmitterRecvMWAck)
	if err := _IbcMwEventsEmitter.contract.UnpackLog(event, "RecvMWAck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcMwEventsEmitterRecvMWPacketIterator is returned from FilterRecvMWPacket and is used to iterate over the raw logs and unpacked data for RecvMWPacket events raised by the IbcMwEventsEmitter contract.
type IbcMwEventsEmitterRecvMWPacketIterator struct {
	Event *IbcMwEventsEmitterRecvMWPacket // Event containing the contract specifics and raw log

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
func (it *IbcMwEventsEmitterRecvMWPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcMwEventsEmitterRecvMWPacket)
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
		it.Event = new(IbcMwEventsEmitterRecvMWPacket)
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
func (it *IbcMwEventsEmitterRecvMWPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcMwEventsEmitterRecvMWPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcMwEventsEmitterRecvMWPacket represents a RecvMWPacket event raised by the IbcMwEventsEmitter contract.
type IbcMwEventsEmitterRecvMWPacket struct {
	ChannelId    [32]byte
	SrcPortAddr  [32]byte
	DestPortAddr [32]byte
	MwId         *big.Int
	AppData      []byte
	MwData       []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRecvMWPacket is a free log retrieval operation binding the contract event 0xdbf4e1c85f2bbbba15b8c4ca75bb0ef2a3e496e986f4fde306aa5942a48ad180.
//
// Solidity: event RecvMWPacket(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, bytes mwData)
func (_IbcMwEventsEmitter *IbcMwEventsEmitterFilterer) FilterRecvMWPacket(opts *bind.FilterOpts, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (*IbcMwEventsEmitterRecvMWPacketIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var srcPortAddrRule []interface{}
	for _, srcPortAddrItem := range srcPortAddr {
		srcPortAddrRule = append(srcPortAddrRule, srcPortAddrItem)
	}
	var destPortAddrRule []interface{}
	for _, destPortAddrItem := range destPortAddr {
		destPortAddrRule = append(destPortAddrRule, destPortAddrItem)
	}

	logs, sub, err := _IbcMwEventsEmitter.contract.FilterLogs(opts, "RecvMWPacket", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return &IbcMwEventsEmitterRecvMWPacketIterator{contract: _IbcMwEventsEmitter.contract, event: "RecvMWPacket", logs: logs, sub: sub}, nil
}

// WatchRecvMWPacket is a free log subscription operation binding the contract event 0xdbf4e1c85f2bbbba15b8c4ca75bb0ef2a3e496e986f4fde306aa5942a48ad180.
//
// Solidity: event RecvMWPacket(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, bytes mwData)
func (_IbcMwEventsEmitter *IbcMwEventsEmitterFilterer) WatchRecvMWPacket(opts *bind.WatchOpts, sink chan<- *IbcMwEventsEmitterRecvMWPacket, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var srcPortAddrRule []interface{}
	for _, srcPortAddrItem := range srcPortAddr {
		srcPortAddrRule = append(srcPortAddrRule, srcPortAddrItem)
	}
	var destPortAddrRule []interface{}
	for _, destPortAddrItem := range destPortAddr {
		destPortAddrRule = append(destPortAddrRule, destPortAddrItem)
	}

	logs, sub, err := _IbcMwEventsEmitter.contract.WatchLogs(opts, "RecvMWPacket", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcMwEventsEmitterRecvMWPacket)
				if err := _IbcMwEventsEmitter.contract.UnpackLog(event, "RecvMWPacket", log); err != nil {
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

// ParseRecvMWPacket is a log parse operation binding the contract event 0xdbf4e1c85f2bbbba15b8c4ca75bb0ef2a3e496e986f4fde306aa5942a48ad180.
//
// Solidity: event RecvMWPacket(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, bytes mwData)
func (_IbcMwEventsEmitter *IbcMwEventsEmitterFilterer) ParseRecvMWPacket(log types.Log) (*IbcMwEventsEmitterRecvMWPacket, error) {
	event := new(IbcMwEventsEmitterRecvMWPacket)
	if err := _IbcMwEventsEmitter.contract.UnpackLog(event, "RecvMWPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcMwEventsEmitterRecvMWTimeoutIterator is returned from FilterRecvMWTimeout and is used to iterate over the raw logs and unpacked data for RecvMWTimeout events raised by the IbcMwEventsEmitter contract.
type IbcMwEventsEmitterRecvMWTimeoutIterator struct {
	Event *IbcMwEventsEmitterRecvMWTimeout // Event containing the contract specifics and raw log

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
func (it *IbcMwEventsEmitterRecvMWTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcMwEventsEmitterRecvMWTimeout)
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
		it.Event = new(IbcMwEventsEmitterRecvMWTimeout)
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
func (it *IbcMwEventsEmitterRecvMWTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcMwEventsEmitterRecvMWTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcMwEventsEmitterRecvMWTimeout represents a RecvMWTimeout event raised by the IbcMwEventsEmitter contract.
type IbcMwEventsEmitterRecvMWTimeout struct {
	ChannelId    [32]byte
	SrcPortAddr  [32]byte
	DestPortAddr [32]byte
	MwId         *big.Int
	AppData      []byte
	MwData       []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRecvMWTimeout is a free log retrieval operation binding the contract event 0xc2c6119a1960aa25d517e5173005d0d39b4f58db74eda3eb48acda26cac7ca62.
//
// Solidity: event RecvMWTimeout(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, bytes mwData)
func (_IbcMwEventsEmitter *IbcMwEventsEmitterFilterer) FilterRecvMWTimeout(opts *bind.FilterOpts, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (*IbcMwEventsEmitterRecvMWTimeoutIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var srcPortAddrRule []interface{}
	for _, srcPortAddrItem := range srcPortAddr {
		srcPortAddrRule = append(srcPortAddrRule, srcPortAddrItem)
	}
	var destPortAddrRule []interface{}
	for _, destPortAddrItem := range destPortAddr {
		destPortAddrRule = append(destPortAddrRule, destPortAddrItem)
	}

	logs, sub, err := _IbcMwEventsEmitter.contract.FilterLogs(opts, "RecvMWTimeout", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return &IbcMwEventsEmitterRecvMWTimeoutIterator{contract: _IbcMwEventsEmitter.contract, event: "RecvMWTimeout", logs: logs, sub: sub}, nil
}

// WatchRecvMWTimeout is a free log subscription operation binding the contract event 0xc2c6119a1960aa25d517e5173005d0d39b4f58db74eda3eb48acda26cac7ca62.
//
// Solidity: event RecvMWTimeout(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, bytes mwData)
func (_IbcMwEventsEmitter *IbcMwEventsEmitterFilterer) WatchRecvMWTimeout(opts *bind.WatchOpts, sink chan<- *IbcMwEventsEmitterRecvMWTimeout, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var srcPortAddrRule []interface{}
	for _, srcPortAddrItem := range srcPortAddr {
		srcPortAddrRule = append(srcPortAddrRule, srcPortAddrItem)
	}
	var destPortAddrRule []interface{}
	for _, destPortAddrItem := range destPortAddr {
		destPortAddrRule = append(destPortAddrRule, destPortAddrItem)
	}

	logs, sub, err := _IbcMwEventsEmitter.contract.WatchLogs(opts, "RecvMWTimeout", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcMwEventsEmitterRecvMWTimeout)
				if err := _IbcMwEventsEmitter.contract.UnpackLog(event, "RecvMWTimeout", log); err != nil {
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

// ParseRecvMWTimeout is a log parse operation binding the contract event 0xc2c6119a1960aa25d517e5173005d0d39b4f58db74eda3eb48acda26cac7ca62.
//
// Solidity: event RecvMWTimeout(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, bytes mwData)
func (_IbcMwEventsEmitter *IbcMwEventsEmitterFilterer) ParseRecvMWTimeout(log types.Log) (*IbcMwEventsEmitterRecvMWTimeout, error) {
	event := new(IbcMwEventsEmitterRecvMWTimeout)
	if err := _IbcMwEventsEmitter.contract.UnpackLog(event, "RecvMWTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcMwEventsEmitterSendMWPacketIterator is returned from FilterSendMWPacket and is used to iterate over the raw logs and unpacked data for SendMWPacket events raised by the IbcMwEventsEmitter contract.
type IbcMwEventsEmitterSendMWPacketIterator struct {
	Event *IbcMwEventsEmitterSendMWPacket // Event containing the contract specifics and raw log

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
func (it *IbcMwEventsEmitterSendMWPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcMwEventsEmitterSendMWPacket)
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
		it.Event = new(IbcMwEventsEmitterSendMWPacket)
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
func (it *IbcMwEventsEmitterSendMWPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcMwEventsEmitterSendMWPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcMwEventsEmitterSendMWPacket represents a SendMWPacket event raised by the IbcMwEventsEmitter contract.
type IbcMwEventsEmitterSendMWPacket struct {
	ChannelId        [32]byte
	SrcPortAddr      [32]byte
	DestPortAddr     [32]byte
	MwId             *big.Int
	AppData          []byte
	TimeoutTimestamp uint64
	MwData           []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSendMWPacket is a free log retrieval operation binding the contract event 0x7c1d77d7984ba9b37263b84eda35413a5b3911d71bd1f24b60bb43cfcaf539f7.
//
// Solidity: event SendMWPacket(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, uint64 timeoutTimestamp, bytes mwData)
func (_IbcMwEventsEmitter *IbcMwEventsEmitterFilterer) FilterSendMWPacket(opts *bind.FilterOpts, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (*IbcMwEventsEmitterSendMWPacketIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var srcPortAddrRule []interface{}
	for _, srcPortAddrItem := range srcPortAddr {
		srcPortAddrRule = append(srcPortAddrRule, srcPortAddrItem)
	}
	var destPortAddrRule []interface{}
	for _, destPortAddrItem := range destPortAddr {
		destPortAddrRule = append(destPortAddrRule, destPortAddrItem)
	}

	logs, sub, err := _IbcMwEventsEmitter.contract.FilterLogs(opts, "SendMWPacket", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return &IbcMwEventsEmitterSendMWPacketIterator{contract: _IbcMwEventsEmitter.contract, event: "SendMWPacket", logs: logs, sub: sub}, nil
}

// WatchSendMWPacket is a free log subscription operation binding the contract event 0x7c1d77d7984ba9b37263b84eda35413a5b3911d71bd1f24b60bb43cfcaf539f7.
//
// Solidity: event SendMWPacket(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, uint64 timeoutTimestamp, bytes mwData)
func (_IbcMwEventsEmitter *IbcMwEventsEmitterFilterer) WatchSendMWPacket(opts *bind.WatchOpts, sink chan<- *IbcMwEventsEmitterSendMWPacket, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var srcPortAddrRule []interface{}
	for _, srcPortAddrItem := range srcPortAddr {
		srcPortAddrRule = append(srcPortAddrRule, srcPortAddrItem)
	}
	var destPortAddrRule []interface{}
	for _, destPortAddrItem := range destPortAddr {
		destPortAddrRule = append(destPortAddrRule, destPortAddrItem)
	}

	logs, sub, err := _IbcMwEventsEmitter.contract.WatchLogs(opts, "SendMWPacket", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcMwEventsEmitterSendMWPacket)
				if err := _IbcMwEventsEmitter.contract.UnpackLog(event, "SendMWPacket", log); err != nil {
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

// ParseSendMWPacket is a log parse operation binding the contract event 0x7c1d77d7984ba9b37263b84eda35413a5b3911d71bd1f24b60bb43cfcaf539f7.
//
// Solidity: event SendMWPacket(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, uint64 timeoutTimestamp, bytes mwData)
func (_IbcMwEventsEmitter *IbcMwEventsEmitterFilterer) ParseSendMWPacket(log types.Log) (*IbcMwEventsEmitterSendMWPacket, error) {
	event := new(IbcMwEventsEmitterSendMWPacket)
	if err := _IbcMwEventsEmitter.contract.UnpackLog(event, "SendMWPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
