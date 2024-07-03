// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ifeevault

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

// IFeeVaultMetaData contains all meta data concerning the IFeeVault contract.
var IFeeVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"depositOpenChannelFee\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositSendPacketFee\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawFeesToOwner\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OpenChannelFeeDeposited\",\"inputs\":[{\"name\":\"sourceAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumChannelOrder\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendPacketFeeDeposited\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"indexed\":false,\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"indexed\":false,\"internalType\":\"uint256[2]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"IncorrectFeeSent\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NoFeeSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotDispatcher\",\"inputs\":[]}]",
}

// IFeeVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use IFeeVaultMetaData.ABI instead.
var IFeeVaultABI = IFeeVaultMetaData.ABI

// IFeeVault is an auto generated Go binding around an Ethereum contract.
type IFeeVault struct {
	IFeeVaultCaller     // Read-only binding to the contract
	IFeeVaultTransactor // Write-only binding to the contract
	IFeeVaultFilterer   // Log filterer for contract events
}

// IFeeVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFeeVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFeeVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFeeVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFeeVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFeeVaultSession struct {
	Contract     *IFeeVault        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFeeVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFeeVaultCallerSession struct {
	Contract *IFeeVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IFeeVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFeeVaultTransactorSession struct {
	Contract     *IFeeVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IFeeVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFeeVaultRaw struct {
	Contract *IFeeVault // Generic contract binding to access the raw methods on
}

// IFeeVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFeeVaultCallerRaw struct {
	Contract *IFeeVaultCaller // Generic read-only contract binding to access the raw methods on
}

// IFeeVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFeeVaultTransactorRaw struct {
	Contract *IFeeVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFeeVault creates a new instance of IFeeVault, bound to a specific deployed contract.
func NewIFeeVault(address common.Address, backend bind.ContractBackend) (*IFeeVault, error) {
	contract, err := bindIFeeVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFeeVault{IFeeVaultCaller: IFeeVaultCaller{contract: contract}, IFeeVaultTransactor: IFeeVaultTransactor{contract: contract}, IFeeVaultFilterer: IFeeVaultFilterer{contract: contract}}, nil
}

// NewIFeeVaultCaller creates a new read-only instance of IFeeVault, bound to a specific deployed contract.
func NewIFeeVaultCaller(address common.Address, caller bind.ContractCaller) (*IFeeVaultCaller, error) {
	contract, err := bindIFeeVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeVaultCaller{contract: contract}, nil
}

// NewIFeeVaultTransactor creates a new write-only instance of IFeeVault, bound to a specific deployed contract.
func NewIFeeVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*IFeeVaultTransactor, error) {
	contract, err := bindIFeeVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFeeVaultTransactor{contract: contract}, nil
}

// NewIFeeVaultFilterer creates a new log filterer instance of IFeeVault, bound to a specific deployed contract.
func NewIFeeVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*IFeeVaultFilterer, error) {
	contract, err := bindIFeeVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFeeVaultFilterer{contract: contract}, nil
}

// bindIFeeVault binds a generic wrapper to an already deployed contract.
func bindIFeeVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IFeeVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFeeVault *IFeeVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFeeVault.Contract.IFeeVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFeeVault *IFeeVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeVault.Contract.IFeeVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFeeVault *IFeeVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFeeVault.Contract.IFeeVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFeeVault *IFeeVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFeeVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFeeVault *IFeeVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFeeVault *IFeeVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFeeVault.Contract.contract.Transact(opts, method, params...)
}

// DepositOpenChannelFee is a paid mutator transaction binding the contract method 0xfce34e40.
//
// Solidity: function depositOpenChannelFee(address sender, string version, uint8 ordering, string[] connectionHops, string counterpartyPortId) payable returns()
func (_IFeeVault *IFeeVaultTransactor) DepositOpenChannelFee(opts *bind.TransactOpts, sender common.Address, version string, ordering uint8, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _IFeeVault.contract.Transact(opts, "depositOpenChannelFee", sender, version, ordering, connectionHops, counterpartyPortId)
}

// DepositOpenChannelFee is a paid mutator transaction binding the contract method 0xfce34e40.
//
// Solidity: function depositOpenChannelFee(address sender, string version, uint8 ordering, string[] connectionHops, string counterpartyPortId) payable returns()
func (_IFeeVault *IFeeVaultSession) DepositOpenChannelFee(sender common.Address, version string, ordering uint8, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _IFeeVault.Contract.DepositOpenChannelFee(&_IFeeVault.TransactOpts, sender, version, ordering, connectionHops, counterpartyPortId)
}

// DepositOpenChannelFee is a paid mutator transaction binding the contract method 0xfce34e40.
//
// Solidity: function depositOpenChannelFee(address sender, string version, uint8 ordering, string[] connectionHops, string counterpartyPortId) payable returns()
func (_IFeeVault *IFeeVaultTransactorSession) DepositOpenChannelFee(sender common.Address, version string, ordering uint8, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _IFeeVault.Contract.DepositOpenChannelFee(&_IFeeVault.TransactOpts, sender, version, ordering, connectionHops, counterpartyPortId)
}

// DepositSendPacketFee is a paid mutator transaction binding the contract method 0x18e3404b.
//
// Solidity: function depositSendPacketFee(bytes32 channelId, uint64 sequence, uint256[2] gasLimits, uint256[2] gasPrices) payable returns()
func (_IFeeVault *IFeeVaultTransactor) DepositSendPacketFee(opts *bind.TransactOpts, channelId [32]byte, sequence uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IFeeVault.contract.Transact(opts, "depositSendPacketFee", channelId, sequence, gasLimits, gasPrices)
}

// DepositSendPacketFee is a paid mutator transaction binding the contract method 0x18e3404b.
//
// Solidity: function depositSendPacketFee(bytes32 channelId, uint64 sequence, uint256[2] gasLimits, uint256[2] gasPrices) payable returns()
func (_IFeeVault *IFeeVaultSession) DepositSendPacketFee(channelId [32]byte, sequence uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IFeeVault.Contract.DepositSendPacketFee(&_IFeeVault.TransactOpts, channelId, sequence, gasLimits, gasPrices)
}

// DepositSendPacketFee is a paid mutator transaction binding the contract method 0x18e3404b.
//
// Solidity: function depositSendPacketFee(bytes32 channelId, uint64 sequence, uint256[2] gasLimits, uint256[2] gasPrices) payable returns()
func (_IFeeVault *IFeeVaultTransactorSession) DepositSendPacketFee(channelId [32]byte, sequence uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IFeeVault.Contract.DepositSendPacketFee(&_IFeeVault.TransactOpts, channelId, sequence, gasLimits, gasPrices)
}

// WithdrawFeesToOwner is a paid mutator transaction binding the contract method 0x0be6a22d.
//
// Solidity: function withdrawFeesToOwner() returns()
func (_IFeeVault *IFeeVaultTransactor) WithdrawFeesToOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFeeVault.contract.Transact(opts, "withdrawFeesToOwner")
}

// WithdrawFeesToOwner is a paid mutator transaction binding the contract method 0x0be6a22d.
//
// Solidity: function withdrawFeesToOwner() returns()
func (_IFeeVault *IFeeVaultSession) WithdrawFeesToOwner() (*types.Transaction, error) {
	return _IFeeVault.Contract.WithdrawFeesToOwner(&_IFeeVault.TransactOpts)
}

// WithdrawFeesToOwner is a paid mutator transaction binding the contract method 0x0be6a22d.
//
// Solidity: function withdrawFeesToOwner() returns()
func (_IFeeVault *IFeeVaultTransactorSession) WithdrawFeesToOwner() (*types.Transaction, error) {
	return _IFeeVault.Contract.WithdrawFeesToOwner(&_IFeeVault.TransactOpts)
}

// IFeeVaultOpenChannelFeeDepositedIterator is returned from FilterOpenChannelFeeDeposited and is used to iterate over the raw logs and unpacked data for OpenChannelFeeDeposited events raised by the IFeeVault contract.
type IFeeVaultOpenChannelFeeDepositedIterator struct {
	Event *IFeeVaultOpenChannelFeeDeposited // Event containing the contract specifics and raw log

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
func (it *IFeeVaultOpenChannelFeeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFeeVaultOpenChannelFeeDeposited)
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
		it.Event = new(IFeeVaultOpenChannelFeeDeposited)
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
func (it *IFeeVaultOpenChannelFeeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFeeVaultOpenChannelFeeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFeeVaultOpenChannelFeeDeposited represents a OpenChannelFeeDeposited event raised by the IFeeVault contract.
type IFeeVaultOpenChannelFeeDeposited struct {
	SourceAddress      common.Address
	Version            string
	Ordering           uint8
	ConnectionHops     []string
	CounterpartyPortId string
	FeeAmount          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOpenChannelFeeDeposited is a free log retrieval operation binding the contract event 0x8ab5595b5ac9231b64513ba86f6bd9fb73c51cae40c36083f7dfc2298e4429e6.
//
// Solidity: event OpenChannelFeeDeposited(address sourceAddress, string version, uint8 ordering, string[] connectionHops, string counterpartyPortId, uint256 feeAmount)
func (_IFeeVault *IFeeVaultFilterer) FilterOpenChannelFeeDeposited(opts *bind.FilterOpts) (*IFeeVaultOpenChannelFeeDepositedIterator, error) {

	logs, sub, err := _IFeeVault.contract.FilterLogs(opts, "OpenChannelFeeDeposited")
	if err != nil {
		return nil, err
	}
	return &IFeeVaultOpenChannelFeeDepositedIterator{contract: _IFeeVault.contract, event: "OpenChannelFeeDeposited", logs: logs, sub: sub}, nil
}

// WatchOpenChannelFeeDeposited is a free log subscription operation binding the contract event 0x8ab5595b5ac9231b64513ba86f6bd9fb73c51cae40c36083f7dfc2298e4429e6.
//
// Solidity: event OpenChannelFeeDeposited(address sourceAddress, string version, uint8 ordering, string[] connectionHops, string counterpartyPortId, uint256 feeAmount)
func (_IFeeVault *IFeeVaultFilterer) WatchOpenChannelFeeDeposited(opts *bind.WatchOpts, sink chan<- *IFeeVaultOpenChannelFeeDeposited) (event.Subscription, error) {

	logs, sub, err := _IFeeVault.contract.WatchLogs(opts, "OpenChannelFeeDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFeeVaultOpenChannelFeeDeposited)
				if err := _IFeeVault.contract.UnpackLog(event, "OpenChannelFeeDeposited", log); err != nil {
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

// ParseOpenChannelFeeDeposited is a log parse operation binding the contract event 0x8ab5595b5ac9231b64513ba86f6bd9fb73c51cae40c36083f7dfc2298e4429e6.
//
// Solidity: event OpenChannelFeeDeposited(address sourceAddress, string version, uint8 ordering, string[] connectionHops, string counterpartyPortId, uint256 feeAmount)
func (_IFeeVault *IFeeVaultFilterer) ParseOpenChannelFeeDeposited(log types.Log) (*IFeeVaultOpenChannelFeeDeposited, error) {
	event := new(IFeeVaultOpenChannelFeeDeposited)
	if err := _IFeeVault.contract.UnpackLog(event, "OpenChannelFeeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFeeVaultSendPacketFeeDepositedIterator is returned from FilterSendPacketFeeDeposited and is used to iterate over the raw logs and unpacked data for SendPacketFeeDeposited events raised by the IFeeVault contract.
type IFeeVaultSendPacketFeeDepositedIterator struct {
	Event *IFeeVaultSendPacketFeeDeposited // Event containing the contract specifics and raw log

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
func (it *IFeeVaultSendPacketFeeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFeeVaultSendPacketFeeDeposited)
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
		it.Event = new(IFeeVaultSendPacketFeeDeposited)
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
func (it *IFeeVaultSendPacketFeeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFeeVaultSendPacketFeeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFeeVaultSendPacketFeeDeposited represents a SendPacketFeeDeposited event raised by the IFeeVault contract.
type IFeeVaultSendPacketFeeDeposited struct {
	ChannelId [32]byte
	Sequence  uint64
	GasLimits [2]*big.Int
	GasPrices [2]*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSendPacketFeeDeposited is a free log retrieval operation binding the contract event 0x0733dc80f277e205edf5d913fa5d91fa0c4cc2635db600b365471c688356c034.
//
// Solidity: event SendPacketFeeDeposited(bytes32 indexed channelId, uint64 indexed sequence, uint256[2] gasLimits, uint256[2] gasPrices)
func (_IFeeVault *IFeeVaultFilterer) FilterSendPacketFeeDeposited(opts *bind.FilterOpts, channelId [][32]byte, sequence []uint64) (*IFeeVaultSendPacketFeeDepositedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _IFeeVault.contract.FilterLogs(opts, "SendPacketFeeDeposited", channelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &IFeeVaultSendPacketFeeDepositedIterator{contract: _IFeeVault.contract, event: "SendPacketFeeDeposited", logs: logs, sub: sub}, nil
}

// WatchSendPacketFeeDeposited is a free log subscription operation binding the contract event 0x0733dc80f277e205edf5d913fa5d91fa0c4cc2635db600b365471c688356c034.
//
// Solidity: event SendPacketFeeDeposited(bytes32 indexed channelId, uint64 indexed sequence, uint256[2] gasLimits, uint256[2] gasPrices)
func (_IFeeVault *IFeeVaultFilterer) WatchSendPacketFeeDeposited(opts *bind.WatchOpts, sink chan<- *IFeeVaultSendPacketFeeDeposited, channelId [][32]byte, sequence []uint64) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _IFeeVault.contract.WatchLogs(opts, "SendPacketFeeDeposited", channelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFeeVaultSendPacketFeeDeposited)
				if err := _IFeeVault.contract.UnpackLog(event, "SendPacketFeeDeposited", log); err != nil {
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

// ParseSendPacketFeeDeposited is a log parse operation binding the contract event 0x0733dc80f277e205edf5d913fa5d91fa0c4cc2635db600b365471c688356c034.
//
// Solidity: event SendPacketFeeDeposited(bytes32 indexed channelId, uint64 indexed sequence, uint256[2] gasLimits, uint256[2] gasPrices)
func (_IFeeVault *IFeeVaultFilterer) ParseSendPacketFeeDeposited(log types.Log) (*IFeeVaultSendPacketFeeDeposited, error) {
	event := new(IFeeVaultSendPacketFeeDeposited)
	if err := _IFeeVault.contract.UnpackLog(event, "SendPacketFeeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
