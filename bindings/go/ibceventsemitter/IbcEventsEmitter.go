// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibceventsemitter

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

// Height is an auto generated low-level Go binding around an user-defined struct.
type Height struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IbcEventsEmitterMetaData contains all meta data concerning the IbcEventsEmitter contract.
var IbcEventsEmitterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"Acknowledgement\",\"inputs\":[{\"name\":\"sourcePortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sourceChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AcknowledgementError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelCloseConfirm\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelCloseConfirmError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelCloseInit\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelCloseInitError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenAck\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenAckError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenConfirm\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenConfirmError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenInit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenInitError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenTry\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenTryError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecvPacket\",\"inputs\":[{\"name\":\"destPortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendPacket\",\"inputs\":[{\"name\":\"sourcePortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sourceChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Timeout\",\"inputs\":[{\"name\":\"sourcePortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sourceChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimeoutError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WriteAckPacket\",\"inputs\":[{\"name\":\"writerPortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"writerChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"ackPacket\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WriteTimeoutPacket\",\"inputs\":[{\"name\":\"writerPortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"writerChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false}]",
}

// IbcEventsEmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcEventsEmitterMetaData.ABI instead.
var IbcEventsEmitterABI = IbcEventsEmitterMetaData.ABI

// IbcEventsEmitter is an auto generated Go binding around an Ethereum contract.
type IbcEventsEmitter struct {
	IbcEventsEmitterCaller     // Read-only binding to the contract
	IbcEventsEmitterTransactor // Write-only binding to the contract
	IbcEventsEmitterFilterer   // Log filterer for contract events
}

// IbcEventsEmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcEventsEmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcEventsEmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcEventsEmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcEventsEmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcEventsEmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcEventsEmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcEventsEmitterSession struct {
	Contract     *IbcEventsEmitter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcEventsEmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcEventsEmitterCallerSession struct {
	Contract *IbcEventsEmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IbcEventsEmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcEventsEmitterTransactorSession struct {
	Contract     *IbcEventsEmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IbcEventsEmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcEventsEmitterRaw struct {
	Contract *IbcEventsEmitter // Generic contract binding to access the raw methods on
}

// IbcEventsEmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcEventsEmitterCallerRaw struct {
	Contract *IbcEventsEmitterCaller // Generic read-only contract binding to access the raw methods on
}

// IbcEventsEmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcEventsEmitterTransactorRaw struct {
	Contract *IbcEventsEmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcEventsEmitter creates a new instance of IbcEventsEmitter, bound to a specific deployed contract.
func NewIbcEventsEmitter(address common.Address, backend bind.ContractBackend) (*IbcEventsEmitter, error) {
	contract, err := bindIbcEventsEmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitter{IbcEventsEmitterCaller: IbcEventsEmitterCaller{contract: contract}, IbcEventsEmitterTransactor: IbcEventsEmitterTransactor{contract: contract}, IbcEventsEmitterFilterer: IbcEventsEmitterFilterer{contract: contract}}, nil
}

// NewIbcEventsEmitterCaller creates a new read-only instance of IbcEventsEmitter, bound to a specific deployed contract.
func NewIbcEventsEmitterCaller(address common.Address, caller bind.ContractCaller) (*IbcEventsEmitterCaller, error) {
	contract, err := bindIbcEventsEmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterCaller{contract: contract}, nil
}

// NewIbcEventsEmitterTransactor creates a new write-only instance of IbcEventsEmitter, bound to a specific deployed contract.
func NewIbcEventsEmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcEventsEmitterTransactor, error) {
	contract, err := bindIbcEventsEmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterTransactor{contract: contract}, nil
}

// NewIbcEventsEmitterFilterer creates a new log filterer instance of IbcEventsEmitter, bound to a specific deployed contract.
func NewIbcEventsEmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcEventsEmitterFilterer, error) {
	contract, err := bindIbcEventsEmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterFilterer{contract: contract}, nil
}

// bindIbcEventsEmitter binds a generic wrapper to an already deployed contract.
func bindIbcEventsEmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcEventsEmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcEventsEmitter *IbcEventsEmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcEventsEmitter.Contract.IbcEventsEmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcEventsEmitter *IbcEventsEmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcEventsEmitter.Contract.IbcEventsEmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcEventsEmitter *IbcEventsEmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcEventsEmitter.Contract.IbcEventsEmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcEventsEmitter *IbcEventsEmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcEventsEmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcEventsEmitter *IbcEventsEmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcEventsEmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcEventsEmitter *IbcEventsEmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcEventsEmitter.Contract.contract.Transact(opts, method, params...)
}

// IbcEventsEmitterAcknowledgementIterator is returned from FilterAcknowledgement and is used to iterate over the raw logs and unpacked data for Acknowledgement events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterAcknowledgementIterator struct {
	Event *IbcEventsEmitterAcknowledgement // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterAcknowledgementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterAcknowledgement)
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
		it.Event = new(IbcEventsEmitterAcknowledgement)
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
func (it *IbcEventsEmitterAcknowledgementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterAcknowledgementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterAcknowledgement represents a Acknowledgement event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterAcknowledgement struct {
	SourcePortAddress common.Address
	SourceChannelId   [32]byte
	Sequence          uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAcknowledgement is a free log retrieval operation binding the contract event 0xe46f6591236abe528fe47a3b281fb002524dadd3e62b1f317ed285d07273c3b1.
//
// Solidity: event Acknowledgement(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 sequence)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterAcknowledgement(opts *bind.FilterOpts, sourcePortAddress []common.Address, sourceChannelId [][32]byte) (*IbcEventsEmitterAcknowledgementIterator, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "Acknowledgement", sourcePortAddressRule, sourceChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterAcknowledgementIterator{contract: _IbcEventsEmitter.contract, event: "Acknowledgement", logs: logs, sub: sub}, nil
}

// WatchAcknowledgement is a free log subscription operation binding the contract event 0xe46f6591236abe528fe47a3b281fb002524dadd3e62b1f317ed285d07273c3b1.
//
// Solidity: event Acknowledgement(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 sequence)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchAcknowledgement(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterAcknowledgement, sourcePortAddress []common.Address, sourceChannelId [][32]byte) (event.Subscription, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "Acknowledgement", sourcePortAddressRule, sourceChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterAcknowledgement)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "Acknowledgement", log); err != nil {
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

// ParseAcknowledgement is a log parse operation binding the contract event 0xe46f6591236abe528fe47a3b281fb002524dadd3e62b1f317ed285d07273c3b1.
//
// Solidity: event Acknowledgement(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 sequence)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseAcknowledgement(log types.Log) (*IbcEventsEmitterAcknowledgement, error) {
	event := new(IbcEventsEmitterAcknowledgement)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "Acknowledgement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterAcknowledgementErrorIterator is returned from FilterAcknowledgementError and is used to iterate over the raw logs and unpacked data for AcknowledgementError events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterAcknowledgementErrorIterator struct {
	Event *IbcEventsEmitterAcknowledgementError // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterAcknowledgementErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterAcknowledgementError)
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
		it.Event = new(IbcEventsEmitterAcknowledgementError)
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
func (it *IbcEventsEmitterAcknowledgementErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterAcknowledgementErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterAcknowledgementError represents a AcknowledgementError event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterAcknowledgementError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcknowledgementError is a free log retrieval operation binding the contract event 0x625eea143c9dae6915c809da47016c22d9cd006c3ace7c345c5cbcf57d3aefbc.
//
// Solidity: event AcknowledgementError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterAcknowledgementError(opts *bind.FilterOpts, receiver []common.Address) (*IbcEventsEmitterAcknowledgementErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "AcknowledgementError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterAcknowledgementErrorIterator{contract: _IbcEventsEmitter.contract, event: "AcknowledgementError", logs: logs, sub: sub}, nil
}

// WatchAcknowledgementError is a free log subscription operation binding the contract event 0x625eea143c9dae6915c809da47016c22d9cd006c3ace7c345c5cbcf57d3aefbc.
//
// Solidity: event AcknowledgementError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchAcknowledgementError(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterAcknowledgementError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "AcknowledgementError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterAcknowledgementError)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "AcknowledgementError", log); err != nil {
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

// ParseAcknowledgementError is a log parse operation binding the contract event 0x625eea143c9dae6915c809da47016c22d9cd006c3ace7c345c5cbcf57d3aefbc.
//
// Solidity: event AcknowledgementError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseAcknowledgementError(log types.Log) (*IbcEventsEmitterAcknowledgementError, error) {
	event := new(IbcEventsEmitterAcknowledgementError)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "AcknowledgementError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterChannelCloseConfirmIterator is returned from FilterChannelCloseConfirm and is used to iterate over the raw logs and unpacked data for ChannelCloseConfirm events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelCloseConfirmIterator struct {
	Event *IbcEventsEmitterChannelCloseConfirm // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterChannelCloseConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterChannelCloseConfirm)
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
		it.Event = new(IbcEventsEmitterChannelCloseConfirm)
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
func (it *IbcEventsEmitterChannelCloseConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterChannelCloseConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterChannelCloseConfirm represents a ChannelCloseConfirm event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelCloseConfirm struct {
	PortAddress common.Address
	ChannelId   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseConfirm is a free log retrieval operation binding the contract event 0x5f010dbbd6bf46aec8131c5456301a75cd688d3cca9bc8380c9e26301be20729.
//
// Solidity: event ChannelCloseConfirm(address indexed portAddress, bytes32 indexed channelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterChannelCloseConfirm(opts *bind.FilterOpts, portAddress []common.Address, channelId [][32]byte) (*IbcEventsEmitterChannelCloseConfirmIterator, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "ChannelCloseConfirm", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterChannelCloseConfirmIterator{contract: _IbcEventsEmitter.contract, event: "ChannelCloseConfirm", logs: logs, sub: sub}, nil
}

// WatchChannelCloseConfirm is a free log subscription operation binding the contract event 0x5f010dbbd6bf46aec8131c5456301a75cd688d3cca9bc8380c9e26301be20729.
//
// Solidity: event ChannelCloseConfirm(address indexed portAddress, bytes32 indexed channelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchChannelCloseConfirm(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterChannelCloseConfirm, portAddress []common.Address, channelId [][32]byte) (event.Subscription, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "ChannelCloseConfirm", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterChannelCloseConfirm)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelCloseConfirm", log); err != nil {
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

// ParseChannelCloseConfirm is a log parse operation binding the contract event 0x5f010dbbd6bf46aec8131c5456301a75cd688d3cca9bc8380c9e26301be20729.
//
// Solidity: event ChannelCloseConfirm(address indexed portAddress, bytes32 indexed channelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseChannelCloseConfirm(log types.Log) (*IbcEventsEmitterChannelCloseConfirm, error) {
	event := new(IbcEventsEmitterChannelCloseConfirm)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelCloseConfirm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterChannelCloseConfirmErrorIterator is returned from FilterChannelCloseConfirmError and is used to iterate over the raw logs and unpacked data for ChannelCloseConfirmError events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelCloseConfirmErrorIterator struct {
	Event *IbcEventsEmitterChannelCloseConfirmError // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterChannelCloseConfirmErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterChannelCloseConfirmError)
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
		it.Event = new(IbcEventsEmitterChannelCloseConfirmError)
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
func (it *IbcEventsEmitterChannelCloseConfirmErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterChannelCloseConfirmErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterChannelCloseConfirmError represents a ChannelCloseConfirmError event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelCloseConfirmError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseConfirmError is a free log retrieval operation binding the contract event 0xc9d36d7a317cb116925d5cb66f0069fe825822fe23e9cf3f421c38cf444caa30.
//
// Solidity: event ChannelCloseConfirmError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterChannelCloseConfirmError(opts *bind.FilterOpts, receiver []common.Address) (*IbcEventsEmitterChannelCloseConfirmErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "ChannelCloseConfirmError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterChannelCloseConfirmErrorIterator{contract: _IbcEventsEmitter.contract, event: "ChannelCloseConfirmError", logs: logs, sub: sub}, nil
}

// WatchChannelCloseConfirmError is a free log subscription operation binding the contract event 0xc9d36d7a317cb116925d5cb66f0069fe825822fe23e9cf3f421c38cf444caa30.
//
// Solidity: event ChannelCloseConfirmError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchChannelCloseConfirmError(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterChannelCloseConfirmError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "ChannelCloseConfirmError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterChannelCloseConfirmError)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelCloseConfirmError", log); err != nil {
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

// ParseChannelCloseConfirmError is a log parse operation binding the contract event 0xc9d36d7a317cb116925d5cb66f0069fe825822fe23e9cf3f421c38cf444caa30.
//
// Solidity: event ChannelCloseConfirmError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseChannelCloseConfirmError(log types.Log) (*IbcEventsEmitterChannelCloseConfirmError, error) {
	event := new(IbcEventsEmitterChannelCloseConfirmError)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelCloseConfirmError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterChannelCloseInitIterator is returned from FilterChannelCloseInit and is used to iterate over the raw logs and unpacked data for ChannelCloseInit events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelCloseInitIterator struct {
	Event *IbcEventsEmitterChannelCloseInit // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterChannelCloseInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterChannelCloseInit)
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
		it.Event = new(IbcEventsEmitterChannelCloseInit)
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
func (it *IbcEventsEmitterChannelCloseInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterChannelCloseInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterChannelCloseInit represents a ChannelCloseInit event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelCloseInit struct {
	PortAddress common.Address
	ChannelId   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseInit is a free log retrieval operation binding the contract event 0x21372e37743553ba8e5f61239803174a827c7732d53ec8adcb76c6b3bb2c13f1.
//
// Solidity: event ChannelCloseInit(address indexed portAddress, bytes32 indexed channelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterChannelCloseInit(opts *bind.FilterOpts, portAddress []common.Address, channelId [][32]byte) (*IbcEventsEmitterChannelCloseInitIterator, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "ChannelCloseInit", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterChannelCloseInitIterator{contract: _IbcEventsEmitter.contract, event: "ChannelCloseInit", logs: logs, sub: sub}, nil
}

// WatchChannelCloseInit is a free log subscription operation binding the contract event 0x21372e37743553ba8e5f61239803174a827c7732d53ec8adcb76c6b3bb2c13f1.
//
// Solidity: event ChannelCloseInit(address indexed portAddress, bytes32 indexed channelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchChannelCloseInit(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterChannelCloseInit, portAddress []common.Address, channelId [][32]byte) (event.Subscription, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "ChannelCloseInit", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterChannelCloseInit)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelCloseInit", log); err != nil {
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

// ParseChannelCloseInit is a log parse operation binding the contract event 0x21372e37743553ba8e5f61239803174a827c7732d53ec8adcb76c6b3bb2c13f1.
//
// Solidity: event ChannelCloseInit(address indexed portAddress, bytes32 indexed channelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseChannelCloseInit(log types.Log) (*IbcEventsEmitterChannelCloseInit, error) {
	event := new(IbcEventsEmitterChannelCloseInit)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelCloseInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterChannelCloseInitErrorIterator is returned from FilterChannelCloseInitError and is used to iterate over the raw logs and unpacked data for ChannelCloseInitError events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelCloseInitErrorIterator struct {
	Event *IbcEventsEmitterChannelCloseInitError // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterChannelCloseInitErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterChannelCloseInitError)
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
		it.Event = new(IbcEventsEmitterChannelCloseInitError)
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
func (it *IbcEventsEmitterChannelCloseInitErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterChannelCloseInitErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterChannelCloseInitError represents a ChannelCloseInitError event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelCloseInitError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseInitError is a free log retrieval operation binding the contract event 0xb1be59c1bcd39c54c7132a8e0d321af5db427575ddb3265560d8862804f4381b.
//
// Solidity: event ChannelCloseInitError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterChannelCloseInitError(opts *bind.FilterOpts, receiver []common.Address) (*IbcEventsEmitterChannelCloseInitErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "ChannelCloseInitError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterChannelCloseInitErrorIterator{contract: _IbcEventsEmitter.contract, event: "ChannelCloseInitError", logs: logs, sub: sub}, nil
}

// WatchChannelCloseInitError is a free log subscription operation binding the contract event 0xb1be59c1bcd39c54c7132a8e0d321af5db427575ddb3265560d8862804f4381b.
//
// Solidity: event ChannelCloseInitError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchChannelCloseInitError(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterChannelCloseInitError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "ChannelCloseInitError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterChannelCloseInitError)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelCloseInitError", log); err != nil {
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

// ParseChannelCloseInitError is a log parse operation binding the contract event 0xb1be59c1bcd39c54c7132a8e0d321af5db427575ddb3265560d8862804f4381b.
//
// Solidity: event ChannelCloseInitError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseChannelCloseInitError(log types.Log) (*IbcEventsEmitterChannelCloseInitError, error) {
	event := new(IbcEventsEmitterChannelCloseInitError)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelCloseInitError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterChannelOpenAckIterator is returned from FilterChannelOpenAck and is used to iterate over the raw logs and unpacked data for ChannelOpenAck events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenAckIterator struct {
	Event *IbcEventsEmitterChannelOpenAck // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterChannelOpenAckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterChannelOpenAck)
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
		it.Event = new(IbcEventsEmitterChannelOpenAck)
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
func (it *IbcEventsEmitterChannelOpenAckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterChannelOpenAckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterChannelOpenAck represents a ChannelOpenAck event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenAck struct {
	Receiver  common.Address
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenAck is a free log retrieval operation binding the contract event 0xcf8be9ab2b5edf8beb2c45abe8e0cc7646318ac19f6c3164ba2e19e93a8a32af.
//
// Solidity: event ChannelOpenAck(address indexed receiver, bytes32 channelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterChannelOpenAck(opts *bind.FilterOpts, receiver []common.Address) (*IbcEventsEmitterChannelOpenAckIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "ChannelOpenAck", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterChannelOpenAckIterator{contract: _IbcEventsEmitter.contract, event: "ChannelOpenAck", logs: logs, sub: sub}, nil
}

// WatchChannelOpenAck is a free log subscription operation binding the contract event 0xcf8be9ab2b5edf8beb2c45abe8e0cc7646318ac19f6c3164ba2e19e93a8a32af.
//
// Solidity: event ChannelOpenAck(address indexed receiver, bytes32 channelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchChannelOpenAck(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterChannelOpenAck, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "ChannelOpenAck", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterChannelOpenAck)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenAck", log); err != nil {
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

// ParseChannelOpenAck is a log parse operation binding the contract event 0xcf8be9ab2b5edf8beb2c45abe8e0cc7646318ac19f6c3164ba2e19e93a8a32af.
//
// Solidity: event ChannelOpenAck(address indexed receiver, bytes32 channelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseChannelOpenAck(log types.Log) (*IbcEventsEmitterChannelOpenAck, error) {
	event := new(IbcEventsEmitterChannelOpenAck)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenAck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterChannelOpenAckErrorIterator is returned from FilterChannelOpenAckError and is used to iterate over the raw logs and unpacked data for ChannelOpenAckError events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenAckErrorIterator struct {
	Event *IbcEventsEmitterChannelOpenAckError // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterChannelOpenAckErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterChannelOpenAckError)
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
		it.Event = new(IbcEventsEmitterChannelOpenAckError)
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
func (it *IbcEventsEmitterChannelOpenAckErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterChannelOpenAckErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterChannelOpenAckError represents a ChannelOpenAckError event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenAckError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenAckError is a free log retrieval operation binding the contract event 0x971a4433f5bff5f011728a4123aeeca4b5275ac20b013cf276e65510491ac26f.
//
// Solidity: event ChannelOpenAckError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterChannelOpenAckError(opts *bind.FilterOpts, receiver []common.Address) (*IbcEventsEmitterChannelOpenAckErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "ChannelOpenAckError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterChannelOpenAckErrorIterator{contract: _IbcEventsEmitter.contract, event: "ChannelOpenAckError", logs: logs, sub: sub}, nil
}

// WatchChannelOpenAckError is a free log subscription operation binding the contract event 0x971a4433f5bff5f011728a4123aeeca4b5275ac20b013cf276e65510491ac26f.
//
// Solidity: event ChannelOpenAckError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchChannelOpenAckError(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterChannelOpenAckError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "ChannelOpenAckError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterChannelOpenAckError)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenAckError", log); err != nil {
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

// ParseChannelOpenAckError is a log parse operation binding the contract event 0x971a4433f5bff5f011728a4123aeeca4b5275ac20b013cf276e65510491ac26f.
//
// Solidity: event ChannelOpenAckError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseChannelOpenAckError(log types.Log) (*IbcEventsEmitterChannelOpenAckError, error) {
	event := new(IbcEventsEmitterChannelOpenAckError)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenAckError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterChannelOpenConfirmIterator is returned from FilterChannelOpenConfirm and is used to iterate over the raw logs and unpacked data for ChannelOpenConfirm events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenConfirmIterator struct {
	Event *IbcEventsEmitterChannelOpenConfirm // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterChannelOpenConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterChannelOpenConfirm)
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
		it.Event = new(IbcEventsEmitterChannelOpenConfirm)
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
func (it *IbcEventsEmitterChannelOpenConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterChannelOpenConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterChannelOpenConfirm represents a ChannelOpenConfirm event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenConfirm struct {
	Receiver  common.Address
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenConfirm is a free log retrieval operation binding the contract event 0xe80f571f70f7cabf9d7ac60ece08421be374117776c311c327a083ca398f802f.
//
// Solidity: event ChannelOpenConfirm(address indexed receiver, bytes32 channelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterChannelOpenConfirm(opts *bind.FilterOpts, receiver []common.Address) (*IbcEventsEmitterChannelOpenConfirmIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "ChannelOpenConfirm", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterChannelOpenConfirmIterator{contract: _IbcEventsEmitter.contract, event: "ChannelOpenConfirm", logs: logs, sub: sub}, nil
}

// WatchChannelOpenConfirm is a free log subscription operation binding the contract event 0xe80f571f70f7cabf9d7ac60ece08421be374117776c311c327a083ca398f802f.
//
// Solidity: event ChannelOpenConfirm(address indexed receiver, bytes32 channelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchChannelOpenConfirm(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterChannelOpenConfirm, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "ChannelOpenConfirm", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterChannelOpenConfirm)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenConfirm", log); err != nil {
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

// ParseChannelOpenConfirm is a log parse operation binding the contract event 0xe80f571f70f7cabf9d7ac60ece08421be374117776c311c327a083ca398f802f.
//
// Solidity: event ChannelOpenConfirm(address indexed receiver, bytes32 channelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseChannelOpenConfirm(log types.Log) (*IbcEventsEmitterChannelOpenConfirm, error) {
	event := new(IbcEventsEmitterChannelOpenConfirm)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenConfirm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterChannelOpenConfirmErrorIterator is returned from FilterChannelOpenConfirmError and is used to iterate over the raw logs and unpacked data for ChannelOpenConfirmError events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenConfirmErrorIterator struct {
	Event *IbcEventsEmitterChannelOpenConfirmError // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterChannelOpenConfirmErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterChannelOpenConfirmError)
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
		it.Event = new(IbcEventsEmitterChannelOpenConfirmError)
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
func (it *IbcEventsEmitterChannelOpenConfirmErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterChannelOpenConfirmErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterChannelOpenConfirmError represents a ChannelOpenConfirmError event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenConfirmError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenConfirmError is a free log retrieval operation binding the contract event 0xf6a58ef30f66943749e8c29c661c84da143a1c8ed017f5faa92b509e0000875a.
//
// Solidity: event ChannelOpenConfirmError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterChannelOpenConfirmError(opts *bind.FilterOpts, receiver []common.Address) (*IbcEventsEmitterChannelOpenConfirmErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "ChannelOpenConfirmError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterChannelOpenConfirmErrorIterator{contract: _IbcEventsEmitter.contract, event: "ChannelOpenConfirmError", logs: logs, sub: sub}, nil
}

// WatchChannelOpenConfirmError is a free log subscription operation binding the contract event 0xf6a58ef30f66943749e8c29c661c84da143a1c8ed017f5faa92b509e0000875a.
//
// Solidity: event ChannelOpenConfirmError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchChannelOpenConfirmError(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterChannelOpenConfirmError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "ChannelOpenConfirmError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterChannelOpenConfirmError)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenConfirmError", log); err != nil {
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

// ParseChannelOpenConfirmError is a log parse operation binding the contract event 0xf6a58ef30f66943749e8c29c661c84da143a1c8ed017f5faa92b509e0000875a.
//
// Solidity: event ChannelOpenConfirmError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseChannelOpenConfirmError(log types.Log) (*IbcEventsEmitterChannelOpenConfirmError, error) {
	event := new(IbcEventsEmitterChannelOpenConfirmError)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenConfirmError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterChannelOpenInitIterator is returned from FilterChannelOpenInit and is used to iterate over the raw logs and unpacked data for ChannelOpenInit events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenInitIterator struct {
	Event *IbcEventsEmitterChannelOpenInit // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterChannelOpenInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterChannelOpenInit)
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
		it.Event = new(IbcEventsEmitterChannelOpenInit)
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
func (it *IbcEventsEmitterChannelOpenInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterChannelOpenInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterChannelOpenInit represents a ChannelOpenInit event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenInit struct {
	Receiver           common.Address
	Version            string
	Ordering           uint8
	FeeEnabled         bool
	ConnectionHops     []string
	CounterpartyPortId string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenInit is a free log retrieval operation binding the contract event 0x20fd8a5856711b18d00def4aa6abafbe00ce6d60795e015cc1cad35eb9b46359.
//
// Solidity: event ChannelOpenInit(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterChannelOpenInit(opts *bind.FilterOpts, receiver []common.Address) (*IbcEventsEmitterChannelOpenInitIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "ChannelOpenInit", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterChannelOpenInitIterator{contract: _IbcEventsEmitter.contract, event: "ChannelOpenInit", logs: logs, sub: sub}, nil
}

// WatchChannelOpenInit is a free log subscription operation binding the contract event 0x20fd8a5856711b18d00def4aa6abafbe00ce6d60795e015cc1cad35eb9b46359.
//
// Solidity: event ChannelOpenInit(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchChannelOpenInit(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterChannelOpenInit, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "ChannelOpenInit", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterChannelOpenInit)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenInit", log); err != nil {
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

// ParseChannelOpenInit is a log parse operation binding the contract event 0x20fd8a5856711b18d00def4aa6abafbe00ce6d60795e015cc1cad35eb9b46359.
//
// Solidity: event ChannelOpenInit(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseChannelOpenInit(log types.Log) (*IbcEventsEmitterChannelOpenInit, error) {
	event := new(IbcEventsEmitterChannelOpenInit)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterChannelOpenInitErrorIterator is returned from FilterChannelOpenInitError and is used to iterate over the raw logs and unpacked data for ChannelOpenInitError events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenInitErrorIterator struct {
	Event *IbcEventsEmitterChannelOpenInitError // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterChannelOpenInitErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterChannelOpenInitError)
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
		it.Event = new(IbcEventsEmitterChannelOpenInitError)
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
func (it *IbcEventsEmitterChannelOpenInitErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterChannelOpenInitErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterChannelOpenInitError represents a ChannelOpenInitError event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenInitError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenInitError is a free log retrieval operation binding the contract event 0x69c1283cce89382f0f9ddf19b7c4f05b4d9b3c30c84fc148b1ec800284be58d5.
//
// Solidity: event ChannelOpenInitError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterChannelOpenInitError(opts *bind.FilterOpts, receiver []common.Address) (*IbcEventsEmitterChannelOpenInitErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "ChannelOpenInitError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterChannelOpenInitErrorIterator{contract: _IbcEventsEmitter.contract, event: "ChannelOpenInitError", logs: logs, sub: sub}, nil
}

// WatchChannelOpenInitError is a free log subscription operation binding the contract event 0x69c1283cce89382f0f9ddf19b7c4f05b4d9b3c30c84fc148b1ec800284be58d5.
//
// Solidity: event ChannelOpenInitError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchChannelOpenInitError(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterChannelOpenInitError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "ChannelOpenInitError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterChannelOpenInitError)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenInitError", log); err != nil {
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

// ParseChannelOpenInitError is a log parse operation binding the contract event 0x69c1283cce89382f0f9ddf19b7c4f05b4d9b3c30c84fc148b1ec800284be58d5.
//
// Solidity: event ChannelOpenInitError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseChannelOpenInitError(log types.Log) (*IbcEventsEmitterChannelOpenInitError, error) {
	event := new(IbcEventsEmitterChannelOpenInitError)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenInitError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterChannelOpenTryIterator is returned from FilterChannelOpenTry and is used to iterate over the raw logs and unpacked data for ChannelOpenTry events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenTryIterator struct {
	Event *IbcEventsEmitterChannelOpenTry // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterChannelOpenTryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterChannelOpenTry)
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
		it.Event = new(IbcEventsEmitterChannelOpenTry)
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
func (it *IbcEventsEmitterChannelOpenTryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterChannelOpenTryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterChannelOpenTry represents a ChannelOpenTry event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenTry struct {
	Receiver              common.Address
	Version               string
	Ordering              uint8
	FeeEnabled            bool
	ConnectionHops        []string
	CounterpartyPortId    string
	CounterpartyChannelId [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenTry is a free log retrieval operation binding the contract event 0xf910705a7a768eb5958f281a5f84cae8bffc5dd811ca5cd303dda140a423698c.
//
// Solidity: event ChannelOpenTry(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId, bytes32 counterpartyChannelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterChannelOpenTry(opts *bind.FilterOpts, receiver []common.Address) (*IbcEventsEmitterChannelOpenTryIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "ChannelOpenTry", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterChannelOpenTryIterator{contract: _IbcEventsEmitter.contract, event: "ChannelOpenTry", logs: logs, sub: sub}, nil
}

// WatchChannelOpenTry is a free log subscription operation binding the contract event 0xf910705a7a768eb5958f281a5f84cae8bffc5dd811ca5cd303dda140a423698c.
//
// Solidity: event ChannelOpenTry(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId, bytes32 counterpartyChannelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchChannelOpenTry(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterChannelOpenTry, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "ChannelOpenTry", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterChannelOpenTry)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenTry", log); err != nil {
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

// ParseChannelOpenTry is a log parse operation binding the contract event 0xf910705a7a768eb5958f281a5f84cae8bffc5dd811ca5cd303dda140a423698c.
//
// Solidity: event ChannelOpenTry(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId, bytes32 counterpartyChannelId)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseChannelOpenTry(log types.Log) (*IbcEventsEmitterChannelOpenTry, error) {
	event := new(IbcEventsEmitterChannelOpenTry)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenTry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterChannelOpenTryErrorIterator is returned from FilterChannelOpenTryError and is used to iterate over the raw logs and unpacked data for ChannelOpenTryError events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenTryErrorIterator struct {
	Event *IbcEventsEmitterChannelOpenTryError // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterChannelOpenTryErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterChannelOpenTryError)
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
		it.Event = new(IbcEventsEmitterChannelOpenTryError)
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
func (it *IbcEventsEmitterChannelOpenTryErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterChannelOpenTryErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterChannelOpenTryError represents a ChannelOpenTryError event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterChannelOpenTryError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenTryError is a free log retrieval operation binding the contract event 0x9e2fe55a3b54b57f82334c273f8d048cd7f05ad19c16cf334276a8c1fec4b6fd.
//
// Solidity: event ChannelOpenTryError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterChannelOpenTryError(opts *bind.FilterOpts, receiver []common.Address) (*IbcEventsEmitterChannelOpenTryErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "ChannelOpenTryError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterChannelOpenTryErrorIterator{contract: _IbcEventsEmitter.contract, event: "ChannelOpenTryError", logs: logs, sub: sub}, nil
}

// WatchChannelOpenTryError is a free log subscription operation binding the contract event 0x9e2fe55a3b54b57f82334c273f8d048cd7f05ad19c16cf334276a8c1fec4b6fd.
//
// Solidity: event ChannelOpenTryError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchChannelOpenTryError(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterChannelOpenTryError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "ChannelOpenTryError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterChannelOpenTryError)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenTryError", log); err != nil {
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

// ParseChannelOpenTryError is a log parse operation binding the contract event 0x9e2fe55a3b54b57f82334c273f8d048cd7f05ad19c16cf334276a8c1fec4b6fd.
//
// Solidity: event ChannelOpenTryError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseChannelOpenTryError(log types.Log) (*IbcEventsEmitterChannelOpenTryError, error) {
	event := new(IbcEventsEmitterChannelOpenTryError)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "ChannelOpenTryError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterRecvPacketIterator is returned from FilterRecvPacket and is used to iterate over the raw logs and unpacked data for RecvPacket events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterRecvPacketIterator struct {
	Event *IbcEventsEmitterRecvPacket // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterRecvPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterRecvPacket)
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
		it.Event = new(IbcEventsEmitterRecvPacket)
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
func (it *IbcEventsEmitterRecvPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterRecvPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterRecvPacket represents a RecvPacket event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterRecvPacket struct {
	DestPortAddress common.Address
	DestChannelId   [32]byte
	Sequence        uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRecvPacket is a free log retrieval operation binding the contract event 0xde5b57e6566d68a30b0979431df3d5df6db3b9aa89f8820f595b9315bf86d067.
//
// Solidity: event RecvPacket(address indexed destPortAddress, bytes32 indexed destChannelId, uint64 sequence)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterRecvPacket(opts *bind.FilterOpts, destPortAddress []common.Address, destChannelId [][32]byte) (*IbcEventsEmitterRecvPacketIterator, error) {

	var destPortAddressRule []interface{}
	for _, destPortAddressItem := range destPortAddress {
		destPortAddressRule = append(destPortAddressRule, destPortAddressItem)
	}
	var destChannelIdRule []interface{}
	for _, destChannelIdItem := range destChannelId {
		destChannelIdRule = append(destChannelIdRule, destChannelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "RecvPacket", destPortAddressRule, destChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterRecvPacketIterator{contract: _IbcEventsEmitter.contract, event: "RecvPacket", logs: logs, sub: sub}, nil
}

// WatchRecvPacket is a free log subscription operation binding the contract event 0xde5b57e6566d68a30b0979431df3d5df6db3b9aa89f8820f595b9315bf86d067.
//
// Solidity: event RecvPacket(address indexed destPortAddress, bytes32 indexed destChannelId, uint64 sequence)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchRecvPacket(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterRecvPacket, destPortAddress []common.Address, destChannelId [][32]byte) (event.Subscription, error) {

	var destPortAddressRule []interface{}
	for _, destPortAddressItem := range destPortAddress {
		destPortAddressRule = append(destPortAddressRule, destPortAddressItem)
	}
	var destChannelIdRule []interface{}
	for _, destChannelIdItem := range destChannelId {
		destChannelIdRule = append(destChannelIdRule, destChannelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "RecvPacket", destPortAddressRule, destChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterRecvPacket)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "RecvPacket", log); err != nil {
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

// ParseRecvPacket is a log parse operation binding the contract event 0xde5b57e6566d68a30b0979431df3d5df6db3b9aa89f8820f595b9315bf86d067.
//
// Solidity: event RecvPacket(address indexed destPortAddress, bytes32 indexed destChannelId, uint64 sequence)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseRecvPacket(log types.Log) (*IbcEventsEmitterRecvPacket, error) {
	event := new(IbcEventsEmitterRecvPacket)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "RecvPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterSendPacketIterator is returned from FilterSendPacket and is used to iterate over the raw logs and unpacked data for SendPacket events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterSendPacketIterator struct {
	Event *IbcEventsEmitterSendPacket // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterSendPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterSendPacket)
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
		it.Event = new(IbcEventsEmitterSendPacket)
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
func (it *IbcEventsEmitterSendPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterSendPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterSendPacket represents a SendPacket event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterSendPacket struct {
	SourcePortAddress common.Address
	SourceChannelId   [32]byte
	Packet            []byte
	Sequence          uint64
	TimeoutTimestamp  uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSendPacket is a free log retrieval operation binding the contract event 0xb5bff96e18da044e4e34510d16df9053b9f1920f6a960732e5aaf22fe9b80136.
//
// Solidity: event SendPacket(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, bytes packet, uint64 sequence, uint64 timeoutTimestamp)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterSendPacket(opts *bind.FilterOpts, sourcePortAddress []common.Address, sourceChannelId [][32]byte) (*IbcEventsEmitterSendPacketIterator, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "SendPacket", sourcePortAddressRule, sourceChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterSendPacketIterator{contract: _IbcEventsEmitter.contract, event: "SendPacket", logs: logs, sub: sub}, nil
}

// WatchSendPacket is a free log subscription operation binding the contract event 0xb5bff96e18da044e4e34510d16df9053b9f1920f6a960732e5aaf22fe9b80136.
//
// Solidity: event SendPacket(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, bytes packet, uint64 sequence, uint64 timeoutTimestamp)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchSendPacket(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterSendPacket, sourcePortAddress []common.Address, sourceChannelId [][32]byte) (event.Subscription, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "SendPacket", sourcePortAddressRule, sourceChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterSendPacket)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "SendPacket", log); err != nil {
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

// ParseSendPacket is a log parse operation binding the contract event 0xb5bff96e18da044e4e34510d16df9053b9f1920f6a960732e5aaf22fe9b80136.
//
// Solidity: event SendPacket(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, bytes packet, uint64 sequence, uint64 timeoutTimestamp)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseSendPacket(log types.Log) (*IbcEventsEmitterSendPacket, error) {
	event := new(IbcEventsEmitterSendPacket)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "SendPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterTimeoutIterator is returned from FilterTimeout and is used to iterate over the raw logs and unpacked data for Timeout events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterTimeoutIterator struct {
	Event *IbcEventsEmitterTimeout // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterTimeout)
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
		it.Event = new(IbcEventsEmitterTimeout)
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
func (it *IbcEventsEmitterTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterTimeout represents a Timeout event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterTimeout struct {
	SourcePortAddress common.Address
	SourceChannelId   [32]byte
	Sequence          uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTimeout is a free log retrieval operation binding the contract event 0x19ac40c4084d9bfb5b43f819a94bf01c70789b0d579871f59e4f86def04d9344.
//
// Solidity: event Timeout(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 indexed sequence)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterTimeout(opts *bind.FilterOpts, sourcePortAddress []common.Address, sourceChannelId [][32]byte, sequence []uint64) (*IbcEventsEmitterTimeoutIterator, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "Timeout", sourcePortAddressRule, sourceChannelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterTimeoutIterator{contract: _IbcEventsEmitter.contract, event: "Timeout", logs: logs, sub: sub}, nil
}

// WatchTimeout is a free log subscription operation binding the contract event 0x19ac40c4084d9bfb5b43f819a94bf01c70789b0d579871f59e4f86def04d9344.
//
// Solidity: event Timeout(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 indexed sequence)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchTimeout(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterTimeout, sourcePortAddress []common.Address, sourceChannelId [][32]byte, sequence []uint64) (event.Subscription, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "Timeout", sourcePortAddressRule, sourceChannelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterTimeout)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "Timeout", log); err != nil {
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

// ParseTimeout is a log parse operation binding the contract event 0x19ac40c4084d9bfb5b43f819a94bf01c70789b0d579871f59e4f86def04d9344.
//
// Solidity: event Timeout(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 indexed sequence)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseTimeout(log types.Log) (*IbcEventsEmitterTimeout, error) {
	event := new(IbcEventsEmitterTimeout)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "Timeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterTimeoutErrorIterator is returned from FilterTimeoutError and is used to iterate over the raw logs and unpacked data for TimeoutError events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterTimeoutErrorIterator struct {
	Event *IbcEventsEmitterTimeoutError // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterTimeoutErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterTimeoutError)
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
		it.Event = new(IbcEventsEmitterTimeoutError)
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
func (it *IbcEventsEmitterTimeoutErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterTimeoutErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterTimeoutError represents a TimeoutError event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterTimeoutError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTimeoutError is a free log retrieval operation binding the contract event 0x83adb31803bee4e18cda1d04a781d77f6f271718a61b25e3a06f319b5103a330.
//
// Solidity: event TimeoutError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterTimeoutError(opts *bind.FilterOpts, receiver []common.Address) (*IbcEventsEmitterTimeoutErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "TimeoutError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterTimeoutErrorIterator{contract: _IbcEventsEmitter.contract, event: "TimeoutError", logs: logs, sub: sub}, nil
}

// WatchTimeoutError is a free log subscription operation binding the contract event 0x83adb31803bee4e18cda1d04a781d77f6f271718a61b25e3a06f319b5103a330.
//
// Solidity: event TimeoutError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchTimeoutError(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterTimeoutError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "TimeoutError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterTimeoutError)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "TimeoutError", log); err != nil {
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

// ParseTimeoutError is a log parse operation binding the contract event 0x83adb31803bee4e18cda1d04a781d77f6f271718a61b25e3a06f319b5103a330.
//
// Solidity: event TimeoutError(address indexed receiver, bytes error)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseTimeoutError(log types.Log) (*IbcEventsEmitterTimeoutError, error) {
	event := new(IbcEventsEmitterTimeoutError)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "TimeoutError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterWriteAckPacketIterator is returned from FilterWriteAckPacket and is used to iterate over the raw logs and unpacked data for WriteAckPacket events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterWriteAckPacketIterator struct {
	Event *IbcEventsEmitterWriteAckPacket // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterWriteAckPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterWriteAckPacket)
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
		it.Event = new(IbcEventsEmitterWriteAckPacket)
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
func (it *IbcEventsEmitterWriteAckPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterWriteAckPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterWriteAckPacket represents a WriteAckPacket event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterWriteAckPacket struct {
	WriterPortAddress common.Address
	WriterChannelId   [32]byte
	Sequence          uint64
	AckPacket         AckPacket
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWriteAckPacket is a free log retrieval operation binding the contract event 0xa32e6f42b1d63fb83ad73b009a6dbb9413d1da02e95b1bb08f081815eea8db20.
//
// Solidity: event WriteAckPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (bool,bytes) ackPacket)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterWriteAckPacket(opts *bind.FilterOpts, writerPortAddress []common.Address, writerChannelId [][32]byte) (*IbcEventsEmitterWriteAckPacketIterator, error) {

	var writerPortAddressRule []interface{}
	for _, writerPortAddressItem := range writerPortAddress {
		writerPortAddressRule = append(writerPortAddressRule, writerPortAddressItem)
	}
	var writerChannelIdRule []interface{}
	for _, writerChannelIdItem := range writerChannelId {
		writerChannelIdRule = append(writerChannelIdRule, writerChannelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "WriteAckPacket", writerPortAddressRule, writerChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterWriteAckPacketIterator{contract: _IbcEventsEmitter.contract, event: "WriteAckPacket", logs: logs, sub: sub}, nil
}

// WatchWriteAckPacket is a free log subscription operation binding the contract event 0xa32e6f42b1d63fb83ad73b009a6dbb9413d1da02e95b1bb08f081815eea8db20.
//
// Solidity: event WriteAckPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (bool,bytes) ackPacket)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchWriteAckPacket(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterWriteAckPacket, writerPortAddress []common.Address, writerChannelId [][32]byte) (event.Subscription, error) {

	var writerPortAddressRule []interface{}
	for _, writerPortAddressItem := range writerPortAddress {
		writerPortAddressRule = append(writerPortAddressRule, writerPortAddressItem)
	}
	var writerChannelIdRule []interface{}
	for _, writerChannelIdItem := range writerChannelId {
		writerChannelIdRule = append(writerChannelIdRule, writerChannelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "WriteAckPacket", writerPortAddressRule, writerChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterWriteAckPacket)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "WriteAckPacket", log); err != nil {
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

// ParseWriteAckPacket is a log parse operation binding the contract event 0xa32e6f42b1d63fb83ad73b009a6dbb9413d1da02e95b1bb08f081815eea8db20.
//
// Solidity: event WriteAckPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (bool,bytes) ackPacket)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseWriteAckPacket(log types.Log) (*IbcEventsEmitterWriteAckPacket, error) {
	event := new(IbcEventsEmitterWriteAckPacket)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "WriteAckPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IbcEventsEmitterWriteTimeoutPacketIterator is returned from FilterWriteTimeoutPacket and is used to iterate over the raw logs and unpacked data for WriteTimeoutPacket events raised by the IbcEventsEmitter contract.
type IbcEventsEmitterWriteTimeoutPacketIterator struct {
	Event *IbcEventsEmitterWriteTimeoutPacket // Event containing the contract specifics and raw log

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
func (it *IbcEventsEmitterWriteTimeoutPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcEventsEmitterWriteTimeoutPacket)
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
		it.Event = new(IbcEventsEmitterWriteTimeoutPacket)
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
func (it *IbcEventsEmitterWriteTimeoutPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcEventsEmitterWriteTimeoutPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcEventsEmitterWriteTimeoutPacket represents a WriteTimeoutPacket event raised by the IbcEventsEmitter contract.
type IbcEventsEmitterWriteTimeoutPacket struct {
	WriterPortAddress common.Address
	WriterChannelId   [32]byte
	Sequence          uint64
	TimeoutHeight     Height
	TimeoutTimestamp  uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWriteTimeoutPacket is a free log retrieval operation binding the contract event 0xedbcd9eeb09d85c3ea1b5bf002c04478059cb261cab82c903885cefccae374bc.
//
// Solidity: event WriteTimeoutPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) FilterWriteTimeoutPacket(opts *bind.FilterOpts, writerPortAddress []common.Address, writerChannelId [][32]byte) (*IbcEventsEmitterWriteTimeoutPacketIterator, error) {

	var writerPortAddressRule []interface{}
	for _, writerPortAddressItem := range writerPortAddress {
		writerPortAddressRule = append(writerPortAddressRule, writerPortAddressItem)
	}
	var writerChannelIdRule []interface{}
	for _, writerChannelIdItem := range writerChannelId {
		writerChannelIdRule = append(writerChannelIdRule, writerChannelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.FilterLogs(opts, "WriteTimeoutPacket", writerPortAddressRule, writerChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &IbcEventsEmitterWriteTimeoutPacketIterator{contract: _IbcEventsEmitter.contract, event: "WriteTimeoutPacket", logs: logs, sub: sub}, nil
}

// WatchWriteTimeoutPacket is a free log subscription operation binding the contract event 0xedbcd9eeb09d85c3ea1b5bf002c04478059cb261cab82c903885cefccae374bc.
//
// Solidity: event WriteTimeoutPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) WatchWriteTimeoutPacket(opts *bind.WatchOpts, sink chan<- *IbcEventsEmitterWriteTimeoutPacket, writerPortAddress []common.Address, writerChannelId [][32]byte) (event.Subscription, error) {

	var writerPortAddressRule []interface{}
	for _, writerPortAddressItem := range writerPortAddress {
		writerPortAddressRule = append(writerPortAddressRule, writerPortAddressItem)
	}
	var writerChannelIdRule []interface{}
	for _, writerChannelIdItem := range writerChannelId {
		writerChannelIdRule = append(writerChannelIdRule, writerChannelIdItem)
	}

	logs, sub, err := _IbcEventsEmitter.contract.WatchLogs(opts, "WriteTimeoutPacket", writerPortAddressRule, writerChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcEventsEmitterWriteTimeoutPacket)
				if err := _IbcEventsEmitter.contract.UnpackLog(event, "WriteTimeoutPacket", log); err != nil {
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

// ParseWriteTimeoutPacket is a log parse operation binding the contract event 0xedbcd9eeb09d85c3ea1b5bf002c04478059cb261cab82c903885cefccae374bc.
//
// Solidity: event WriteTimeoutPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_IbcEventsEmitter *IbcEventsEmitterFilterer) ParseWriteTimeoutPacket(log types.Log) (*IbcEventsEmitterWriteTimeoutPacket, error) {
	event := new(IbcEventsEmitterWriteTimeoutPacket)
	if err := _IbcEventsEmitter.contract.UnpackLog(event, "WriteTimeoutPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
