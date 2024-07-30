// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feevault

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

// FeeVaultMetaData contains all meta data concerning the FeeVault contract.
var FeeVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositOpenChannelFee\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositSendPacketFee\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFeesToOwner\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OpenChannelFeeDeposited\",\"inputs\":[{\"name\":\"sourceAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumChannelOrder\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"feeAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendPacketFeeDeposited\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"indexed\":false,\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"indexed\":false,\"internalType\":\"uint256[2]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"IncorrectFeeSent\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NoFeeSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SenderNotDispatcher\",\"inputs\":[]}]",
}

// FeeVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeVaultMetaData.ABI instead.
var FeeVaultABI = FeeVaultMetaData.ABI

// FeeVault is an auto generated Go binding around an Ethereum contract.
type FeeVault struct {
	FeeVaultCaller     // Read-only binding to the contract
	FeeVaultTransactor // Write-only binding to the contract
	FeeVaultFilterer   // Log filterer for contract events
}

// FeeVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeVaultSession struct {
	Contract     *FeeVault         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeVaultCallerSession struct {
	Contract *FeeVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FeeVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeVaultTransactorSession struct {
	Contract     *FeeVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FeeVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeVaultRaw struct {
	Contract *FeeVault // Generic contract binding to access the raw methods on
}

// FeeVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeVaultCallerRaw struct {
	Contract *FeeVaultCaller // Generic read-only contract binding to access the raw methods on
}

// FeeVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeVaultTransactorRaw struct {
	Contract *FeeVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeVault creates a new instance of FeeVault, bound to a specific deployed contract.
func NewFeeVault(address common.Address, backend bind.ContractBackend) (*FeeVault, error) {
	contract, err := bindFeeVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeVault{FeeVaultCaller: FeeVaultCaller{contract: contract}, FeeVaultTransactor: FeeVaultTransactor{contract: contract}, FeeVaultFilterer: FeeVaultFilterer{contract: contract}}, nil
}

// NewFeeVaultCaller creates a new read-only instance of FeeVault, bound to a specific deployed contract.
func NewFeeVaultCaller(address common.Address, caller bind.ContractCaller) (*FeeVaultCaller, error) {
	contract, err := bindFeeVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeVaultCaller{contract: contract}, nil
}

// NewFeeVaultTransactor creates a new write-only instance of FeeVault, bound to a specific deployed contract.
func NewFeeVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeVaultTransactor, error) {
	contract, err := bindFeeVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeVaultTransactor{contract: contract}, nil
}

// NewFeeVaultFilterer creates a new log filterer instance of FeeVault, bound to a specific deployed contract.
func NewFeeVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeVaultFilterer, error) {
	contract, err := bindFeeVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeVaultFilterer{contract: contract}, nil
}

// bindFeeVault binds a generic wrapper to an already deployed contract.
func bindFeeVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeVault *FeeVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeVault.Contract.FeeVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeVault *FeeVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeVault.Contract.FeeVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeVault *FeeVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeVault.Contract.FeeVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeVault *FeeVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeVault *FeeVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeVault *FeeVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeVault.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeVault *FeeVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeVault *FeeVaultSession) Owner() (common.Address, error) {
	return _FeeVault.Contract.Owner(&_FeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeVault *FeeVaultCallerSession) Owner() (common.Address, error) {
	return _FeeVault.Contract.Owner(&_FeeVault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FeeVault *FeeVaultCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeVault.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FeeVault *FeeVaultSession) PendingOwner() (common.Address, error) {
	return _FeeVault.Contract.PendingOwner(&_FeeVault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FeeVault *FeeVaultCallerSession) PendingOwner() (common.Address, error) {
	return _FeeVault.Contract.PendingOwner(&_FeeVault.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FeeVault *FeeVaultTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeVault.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FeeVault *FeeVaultSession) AcceptOwnership() (*types.Transaction, error) {
	return _FeeVault.Contract.AcceptOwnership(&_FeeVault.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FeeVault *FeeVaultTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FeeVault.Contract.AcceptOwnership(&_FeeVault.TransactOpts)
}

// DepositOpenChannelFee is a paid mutator transaction binding the contract method 0xfce34e40.
//
// Solidity: function depositOpenChannelFee(address src, string version, uint8 ordering, string[] connectionHops, string counterpartyPortId) payable returns()
func (_FeeVault *FeeVaultTransactor) DepositOpenChannelFee(opts *bind.TransactOpts, src common.Address, version string, ordering uint8, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _FeeVault.contract.Transact(opts, "depositOpenChannelFee", src, version, ordering, connectionHops, counterpartyPortId)
}

// DepositOpenChannelFee is a paid mutator transaction binding the contract method 0xfce34e40.
//
// Solidity: function depositOpenChannelFee(address src, string version, uint8 ordering, string[] connectionHops, string counterpartyPortId) payable returns()
func (_FeeVault *FeeVaultSession) DepositOpenChannelFee(src common.Address, version string, ordering uint8, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _FeeVault.Contract.DepositOpenChannelFee(&_FeeVault.TransactOpts, src, version, ordering, connectionHops, counterpartyPortId)
}

// DepositOpenChannelFee is a paid mutator transaction binding the contract method 0xfce34e40.
//
// Solidity: function depositOpenChannelFee(address src, string version, uint8 ordering, string[] connectionHops, string counterpartyPortId) payable returns()
func (_FeeVault *FeeVaultTransactorSession) DepositOpenChannelFee(src common.Address, version string, ordering uint8, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _FeeVault.Contract.DepositOpenChannelFee(&_FeeVault.TransactOpts, src, version, ordering, connectionHops, counterpartyPortId)
}

// DepositSendPacketFee is a paid mutator transaction binding the contract method 0x18e3404b.
//
// Solidity: function depositSendPacketFee(bytes32 channelId, uint64 sequence, uint256[2] gasLimits, uint256[2] gasPrices) payable returns()
func (_FeeVault *FeeVaultTransactor) DepositSendPacketFee(opts *bind.TransactOpts, channelId [32]byte, sequence uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _FeeVault.contract.Transact(opts, "depositSendPacketFee", channelId, sequence, gasLimits, gasPrices)
}

// DepositSendPacketFee is a paid mutator transaction binding the contract method 0x18e3404b.
//
// Solidity: function depositSendPacketFee(bytes32 channelId, uint64 sequence, uint256[2] gasLimits, uint256[2] gasPrices) payable returns()
func (_FeeVault *FeeVaultSession) DepositSendPacketFee(channelId [32]byte, sequence uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _FeeVault.Contract.DepositSendPacketFee(&_FeeVault.TransactOpts, channelId, sequence, gasLimits, gasPrices)
}

// DepositSendPacketFee is a paid mutator transaction binding the contract method 0x18e3404b.
//
// Solidity: function depositSendPacketFee(bytes32 channelId, uint64 sequence, uint256[2] gasLimits, uint256[2] gasPrices) payable returns()
func (_FeeVault *FeeVaultTransactorSession) DepositSendPacketFee(channelId [32]byte, sequence uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _FeeVault.Contract.DepositSendPacketFee(&_FeeVault.TransactOpts, channelId, sequence, gasLimits, gasPrices)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeVault *FeeVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeVault *FeeVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeVault.Contract.RenounceOwnership(&_FeeVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeVault *FeeVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeVault.Contract.RenounceOwnership(&_FeeVault.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeVault *FeeVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeeVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeVault *FeeVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeVault.Contract.TransferOwnership(&_FeeVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeVault *FeeVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeVault.Contract.TransferOwnership(&_FeeVault.TransactOpts, newOwner)
}

// WithdrawFeesToOwner is a paid mutator transaction binding the contract method 0x0be6a22d.
//
// Solidity: function withdrawFeesToOwner() returns()
func (_FeeVault *FeeVaultTransactor) WithdrawFeesToOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeVault.contract.Transact(opts, "withdrawFeesToOwner")
}

// WithdrawFeesToOwner is a paid mutator transaction binding the contract method 0x0be6a22d.
//
// Solidity: function withdrawFeesToOwner() returns()
func (_FeeVault *FeeVaultSession) WithdrawFeesToOwner() (*types.Transaction, error) {
	return _FeeVault.Contract.WithdrawFeesToOwner(&_FeeVault.TransactOpts)
}

// WithdrawFeesToOwner is a paid mutator transaction binding the contract method 0x0be6a22d.
//
// Solidity: function withdrawFeesToOwner() returns()
func (_FeeVault *FeeVaultTransactorSession) WithdrawFeesToOwner() (*types.Transaction, error) {
	return _FeeVault.Contract.WithdrawFeesToOwner(&_FeeVault.TransactOpts)
}

// FeeVaultOpenChannelFeeDepositedIterator is returned from FilterOpenChannelFeeDeposited and is used to iterate over the raw logs and unpacked data for OpenChannelFeeDeposited events raised by the FeeVault contract.
type FeeVaultOpenChannelFeeDepositedIterator struct {
	Event *FeeVaultOpenChannelFeeDeposited // Event containing the contract specifics and raw log

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
func (it *FeeVaultOpenChannelFeeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeVaultOpenChannelFeeDeposited)
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
		it.Event = new(FeeVaultOpenChannelFeeDeposited)
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
func (it *FeeVaultOpenChannelFeeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeVaultOpenChannelFeeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeVaultOpenChannelFeeDeposited represents a OpenChannelFeeDeposited event raised by the FeeVault contract.
type FeeVaultOpenChannelFeeDeposited struct {
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
func (_FeeVault *FeeVaultFilterer) FilterOpenChannelFeeDeposited(opts *bind.FilterOpts) (*FeeVaultOpenChannelFeeDepositedIterator, error) {

	logs, sub, err := _FeeVault.contract.FilterLogs(opts, "OpenChannelFeeDeposited")
	if err != nil {
		return nil, err
	}
	return &FeeVaultOpenChannelFeeDepositedIterator{contract: _FeeVault.contract, event: "OpenChannelFeeDeposited", logs: logs, sub: sub}, nil
}

// WatchOpenChannelFeeDeposited is a free log subscription operation binding the contract event 0x8ab5595b5ac9231b64513ba86f6bd9fb73c51cae40c36083f7dfc2298e4429e6.
//
// Solidity: event OpenChannelFeeDeposited(address sourceAddress, string version, uint8 ordering, string[] connectionHops, string counterpartyPortId, uint256 feeAmount)
func (_FeeVault *FeeVaultFilterer) WatchOpenChannelFeeDeposited(opts *bind.WatchOpts, sink chan<- *FeeVaultOpenChannelFeeDeposited) (event.Subscription, error) {

	logs, sub, err := _FeeVault.contract.WatchLogs(opts, "OpenChannelFeeDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeVaultOpenChannelFeeDeposited)
				if err := _FeeVault.contract.UnpackLog(event, "OpenChannelFeeDeposited", log); err != nil {
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
func (_FeeVault *FeeVaultFilterer) ParseOpenChannelFeeDeposited(log types.Log) (*FeeVaultOpenChannelFeeDeposited, error) {
	event := new(FeeVaultOpenChannelFeeDeposited)
	if err := _FeeVault.contract.UnpackLog(event, "OpenChannelFeeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeVaultOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the FeeVault contract.
type FeeVaultOwnershipTransferStartedIterator struct {
	Event *FeeVaultOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *FeeVaultOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeVaultOwnershipTransferStarted)
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
		it.Event = new(FeeVaultOwnershipTransferStarted)
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
func (it *FeeVaultOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeVaultOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeVaultOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the FeeVault contract.
type FeeVaultOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FeeVault *FeeVaultFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeeVaultOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeVault.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeVaultOwnershipTransferStartedIterator{contract: _FeeVault.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FeeVault *FeeVaultFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *FeeVaultOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeVault.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeVaultOwnershipTransferStarted)
				if err := _FeeVault.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FeeVault *FeeVaultFilterer) ParseOwnershipTransferStarted(log types.Log) (*FeeVaultOwnershipTransferStarted, error) {
	event := new(FeeVaultOwnershipTransferStarted)
	if err := _FeeVault.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeeVault contract.
type FeeVaultOwnershipTransferredIterator struct {
	Event *FeeVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeeVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeVaultOwnershipTransferred)
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
		it.Event = new(FeeVaultOwnershipTransferred)
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
func (it *FeeVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeVaultOwnershipTransferred represents a OwnershipTransferred event raised by the FeeVault contract.
type FeeVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeVault *FeeVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeeVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeVaultOwnershipTransferredIterator{contract: _FeeVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeVault *FeeVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeVaultOwnershipTransferred)
				if err := _FeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FeeVault *FeeVaultFilterer) ParseOwnershipTransferred(log types.Log) (*FeeVaultOwnershipTransferred, error) {
	event := new(FeeVaultOwnershipTransferred)
	if err := _FeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeVaultSendPacketFeeDepositedIterator is returned from FilterSendPacketFeeDeposited and is used to iterate over the raw logs and unpacked data for SendPacketFeeDeposited events raised by the FeeVault contract.
type FeeVaultSendPacketFeeDepositedIterator struct {
	Event *FeeVaultSendPacketFeeDeposited // Event containing the contract specifics and raw log

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
func (it *FeeVaultSendPacketFeeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeVaultSendPacketFeeDeposited)
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
		it.Event = new(FeeVaultSendPacketFeeDeposited)
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
func (it *FeeVaultSendPacketFeeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeVaultSendPacketFeeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeVaultSendPacketFeeDeposited represents a SendPacketFeeDeposited event raised by the FeeVault contract.
type FeeVaultSendPacketFeeDeposited struct {
	ChannelId [32]byte
	Sequence  uint64
	GasLimits [2]*big.Int
	GasPrices [2]*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSendPacketFeeDeposited is a free log retrieval operation binding the contract event 0x0733dc80f277e205edf5d913fa5d91fa0c4cc2635db600b365471c688356c034.
//
// Solidity: event SendPacketFeeDeposited(bytes32 indexed channelId, uint64 indexed sequence, uint256[2] gasLimits, uint256[2] gasPrices)
func (_FeeVault *FeeVaultFilterer) FilterSendPacketFeeDeposited(opts *bind.FilterOpts, channelId [][32]byte, sequence []uint64) (*FeeVaultSendPacketFeeDepositedIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _FeeVault.contract.FilterLogs(opts, "SendPacketFeeDeposited", channelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &FeeVaultSendPacketFeeDepositedIterator{contract: _FeeVault.contract, event: "SendPacketFeeDeposited", logs: logs, sub: sub}, nil
}

// WatchSendPacketFeeDeposited is a free log subscription operation binding the contract event 0x0733dc80f277e205edf5d913fa5d91fa0c4cc2635db600b365471c688356c034.
//
// Solidity: event SendPacketFeeDeposited(bytes32 indexed channelId, uint64 indexed sequence, uint256[2] gasLimits, uint256[2] gasPrices)
func (_FeeVault *FeeVaultFilterer) WatchSendPacketFeeDeposited(opts *bind.WatchOpts, sink chan<- *FeeVaultSendPacketFeeDeposited, channelId [][32]byte, sequence []uint64) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _FeeVault.contract.WatchLogs(opts, "SendPacketFeeDeposited", channelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeVaultSendPacketFeeDeposited)
				if err := _FeeVault.contract.UnpackLog(event, "SendPacketFeeDeposited", log); err != nil {
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
func (_FeeVault *FeeVaultFilterer) ParseSendPacketFeeDeposited(log types.Log) (*FeeVaultSendPacketFeeDeposited, error) {
	event := new(FeeVaultSendPacketFeeDeposited)
	if err := _FeeVault.contract.UnpackLog(event, "SendPacketFeeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
