// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package venus

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

// VenusMetaData contains all meta data concerning the Venus contract.
var VenusMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_prover\",\"type\":\"address\",\"internalType\":\"contractICrossL2Prover\"},{\"name\":\"_counterParty\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"counterParty\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastReceivedTransmission\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"prover\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICrossL2Prover\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveEvent\",\"inputs\":[{\"name\":\"receiptIndex\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiptRLPEncodedBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"logBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receiveReceipt\",\"inputs\":[{\"name\":\"receiptIndex\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiptRLPEncodedBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"receiveTransmissionEvent\",\"inputs\":[{\"name\":\"receiptIndex\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiptRLPEncodedBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"SuccessfulEvent\",\"inputs\":[{\"name\":\"eventIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SuccessfulReceipt\",\"inputs\":[{\"name\":\"receiptIndex\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"receiptRLP\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransmissionReceived\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransmitToHouston\",\"inputs\":[{\"name\":\"message\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"invalidCounterpartyEvent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidEventProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidEventSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidProverAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidReceiptProof\",\"inputs\":[]}]",
}

// VenusABI is the input ABI used to generate the binding from.
// Deprecated: Use VenusMetaData.ABI instead.
var VenusABI = VenusMetaData.ABI

// Venus is an auto generated Go binding around an Ethereum contract.
type Venus struct {
	VenusCaller     // Read-only binding to the contract
	VenusTransactor // Write-only binding to the contract
	VenusFilterer   // Log filterer for contract events
}

// VenusCaller is an auto generated read-only Go binding around an Ethereum contract.
type VenusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VenusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VenusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VenusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VenusSession struct {
	Contract     *Venus            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VenusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VenusCallerSession struct {
	Contract *VenusCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VenusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VenusTransactorSession struct {
	Contract     *VenusTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VenusRaw is an auto generated low-level Go binding around an Ethereum contract.
type VenusRaw struct {
	Contract *Venus // Generic contract binding to access the raw methods on
}

// VenusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VenusCallerRaw struct {
	Contract *VenusCaller // Generic read-only contract binding to access the raw methods on
}

// VenusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VenusTransactorRaw struct {
	Contract *VenusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVenus creates a new instance of Venus, bound to a specific deployed contract.
func NewVenus(address common.Address, backend bind.ContractBackend) (*Venus, error) {
	contract, err := bindVenus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Venus{VenusCaller: VenusCaller{contract: contract}, VenusTransactor: VenusTransactor{contract: contract}, VenusFilterer: VenusFilterer{contract: contract}}, nil
}

// NewVenusCaller creates a new read-only instance of Venus, bound to a specific deployed contract.
func NewVenusCaller(address common.Address, caller bind.ContractCaller) (*VenusCaller, error) {
	contract, err := bindVenus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VenusCaller{contract: contract}, nil
}

// NewVenusTransactor creates a new write-only instance of Venus, bound to a specific deployed contract.
func NewVenusTransactor(address common.Address, transactor bind.ContractTransactor) (*VenusTransactor, error) {
	contract, err := bindVenus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VenusTransactor{contract: contract}, nil
}

// NewVenusFilterer creates a new log filterer instance of Venus, bound to a specific deployed contract.
func NewVenusFilterer(address common.Address, filterer bind.ContractFilterer) (*VenusFilterer, error) {
	contract, err := bindVenus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VenusFilterer{contract: contract}, nil
}

// bindVenus binds a generic wrapper to an already deployed contract.
func bindVenus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VenusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Venus *VenusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Venus.Contract.VenusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Venus *VenusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Venus.Contract.VenusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Venus *VenusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Venus.Contract.VenusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Venus *VenusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Venus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Venus *VenusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Venus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Venus *VenusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Venus.Contract.contract.Transact(opts, method, params...)
}

// CounterParty is a free data retrieval call binding the contract method 0x2c453bdf.
//
// Solidity: function counterParty() view returns(address)
func (_Venus *VenusCaller) CounterParty(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venus.contract.Call(opts, &out, "counterParty")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CounterParty is a free data retrieval call binding the contract method 0x2c453bdf.
//
// Solidity: function counterParty() view returns(address)
func (_Venus *VenusSession) CounterParty() (common.Address, error) {
	return _Venus.Contract.CounterParty(&_Venus.CallOpts)
}

// CounterParty is a free data retrieval call binding the contract method 0x2c453bdf.
//
// Solidity: function counterParty() view returns(address)
func (_Venus *VenusCallerSession) CounterParty() (common.Address, error) {
	return _Venus.Contract.CounterParty(&_Venus.CallOpts)
}

// LastReceivedTransmission is a free data retrieval call binding the contract method 0xe03d0aac.
//
// Solidity: function lastReceivedTransmission() view returns(bytes32)
func (_Venus *VenusCaller) LastReceivedTransmission(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Venus.contract.Call(opts, &out, "lastReceivedTransmission")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastReceivedTransmission is a free data retrieval call binding the contract method 0xe03d0aac.
//
// Solidity: function lastReceivedTransmission() view returns(bytes32)
func (_Venus *VenusSession) LastReceivedTransmission() ([32]byte, error) {
	return _Venus.Contract.LastReceivedTransmission(&_Venus.CallOpts)
}

// LastReceivedTransmission is a free data retrieval call binding the contract method 0xe03d0aac.
//
// Solidity: function lastReceivedTransmission() view returns(bytes32)
func (_Venus *VenusCallerSession) LastReceivedTransmission() ([32]byte, error) {
	return _Venus.Contract.LastReceivedTransmission(&_Venus.CallOpts)
}

// Prover is a free data retrieval call binding the contract method 0x32a8f30f.
//
// Solidity: function prover() view returns(address)
func (_Venus *VenusCaller) Prover(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Venus.contract.Call(opts, &out, "prover")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prover is a free data retrieval call binding the contract method 0x32a8f30f.
//
// Solidity: function prover() view returns(address)
func (_Venus *VenusSession) Prover() (common.Address, error) {
	return _Venus.Contract.Prover(&_Venus.CallOpts)
}

// Prover is a free data retrieval call binding the contract method 0x32a8f30f.
//
// Solidity: function prover() view returns(address)
func (_Venus *VenusCallerSession) Prover() (common.Address, error) {
	return _Venus.Contract.Prover(&_Venus.CallOpts)
}

// ReceiveEvent is a paid mutator transaction binding the contract method 0x163be0d8.
//
// Solidity: function receiveEvent(bytes receiptIndex, bytes receiptRLPEncodedBytes, uint256 logIndex, bytes logBytes, bytes proof) returns()
func (_Venus *VenusTransactor) ReceiveEvent(opts *bind.TransactOpts, receiptIndex []byte, receiptRLPEncodedBytes []byte, logIndex *big.Int, logBytes []byte, proof []byte) (*types.Transaction, error) {
	return _Venus.contract.Transact(opts, "receiveEvent", receiptIndex, receiptRLPEncodedBytes, logIndex, logBytes, proof)
}

// ReceiveEvent is a paid mutator transaction binding the contract method 0x163be0d8.
//
// Solidity: function receiveEvent(bytes receiptIndex, bytes receiptRLPEncodedBytes, uint256 logIndex, bytes logBytes, bytes proof) returns()
func (_Venus *VenusSession) ReceiveEvent(receiptIndex []byte, receiptRLPEncodedBytes []byte, logIndex *big.Int, logBytes []byte, proof []byte) (*types.Transaction, error) {
	return _Venus.Contract.ReceiveEvent(&_Venus.TransactOpts, receiptIndex, receiptRLPEncodedBytes, logIndex, logBytes, proof)
}

// ReceiveEvent is a paid mutator transaction binding the contract method 0x163be0d8.
//
// Solidity: function receiveEvent(bytes receiptIndex, bytes receiptRLPEncodedBytes, uint256 logIndex, bytes logBytes, bytes proof) returns()
func (_Venus *VenusTransactorSession) ReceiveEvent(receiptIndex []byte, receiptRLPEncodedBytes []byte, logIndex *big.Int, logBytes []byte, proof []byte) (*types.Transaction, error) {
	return _Venus.Contract.ReceiveEvent(&_Venus.TransactOpts, receiptIndex, receiptRLPEncodedBytes, logIndex, logBytes, proof)
}

// ReceiveReceipt is a paid mutator transaction binding the contract method 0x2004bbf0.
//
// Solidity: function receiveReceipt(bytes receiptIndex, bytes receiptRLPEncodedBytes, bytes proof) returns()
func (_Venus *VenusTransactor) ReceiveReceipt(opts *bind.TransactOpts, receiptIndex []byte, receiptRLPEncodedBytes []byte, proof []byte) (*types.Transaction, error) {
	return _Venus.contract.Transact(opts, "receiveReceipt", receiptIndex, receiptRLPEncodedBytes, proof)
}

// ReceiveReceipt is a paid mutator transaction binding the contract method 0x2004bbf0.
//
// Solidity: function receiveReceipt(bytes receiptIndex, bytes receiptRLPEncodedBytes, bytes proof) returns()
func (_Venus *VenusSession) ReceiveReceipt(receiptIndex []byte, receiptRLPEncodedBytes []byte, proof []byte) (*types.Transaction, error) {
	return _Venus.Contract.ReceiveReceipt(&_Venus.TransactOpts, receiptIndex, receiptRLPEncodedBytes, proof)
}

// ReceiveReceipt is a paid mutator transaction binding the contract method 0x2004bbf0.
//
// Solidity: function receiveReceipt(bytes receiptIndex, bytes receiptRLPEncodedBytes, bytes proof) returns()
func (_Venus *VenusTransactorSession) ReceiveReceipt(receiptIndex []byte, receiptRLPEncodedBytes []byte, proof []byte) (*types.Transaction, error) {
	return _Venus.Contract.ReceiveReceipt(&_Venus.TransactOpts, receiptIndex, receiptRLPEncodedBytes, proof)
}

// ReceiveTransmissionEvent is a paid mutator transaction binding the contract method 0x10f7b4c1.
//
// Solidity: function receiveTransmissionEvent(bytes receiptIndex, bytes receiptRLPEncodedBytes, uint256 logIndex, bytes proof) returns()
func (_Venus *VenusTransactor) ReceiveTransmissionEvent(opts *bind.TransactOpts, receiptIndex []byte, receiptRLPEncodedBytes []byte, logIndex *big.Int, proof []byte) (*types.Transaction, error) {
	return _Venus.contract.Transact(opts, "receiveTransmissionEvent", receiptIndex, receiptRLPEncodedBytes, logIndex, proof)
}

// ReceiveTransmissionEvent is a paid mutator transaction binding the contract method 0x10f7b4c1.
//
// Solidity: function receiveTransmissionEvent(bytes receiptIndex, bytes receiptRLPEncodedBytes, uint256 logIndex, bytes proof) returns()
func (_Venus *VenusSession) ReceiveTransmissionEvent(receiptIndex []byte, receiptRLPEncodedBytes []byte, logIndex *big.Int, proof []byte) (*types.Transaction, error) {
	return _Venus.Contract.ReceiveTransmissionEvent(&_Venus.TransactOpts, receiptIndex, receiptRLPEncodedBytes, logIndex, proof)
}

// ReceiveTransmissionEvent is a paid mutator transaction binding the contract method 0x10f7b4c1.
//
// Solidity: function receiveTransmissionEvent(bytes receiptIndex, bytes receiptRLPEncodedBytes, uint256 logIndex, bytes proof) returns()
func (_Venus *VenusTransactorSession) ReceiveTransmissionEvent(receiptIndex []byte, receiptRLPEncodedBytes []byte, logIndex *big.Int, proof []byte) (*types.Transaction, error) {
	return _Venus.Contract.ReceiveTransmissionEvent(&_Venus.TransactOpts, receiptIndex, receiptRLPEncodedBytes, logIndex, proof)
}

// VenusSuccessfulEventIterator is returned from FilterSuccessfulEvent and is used to iterate over the raw logs and unpacked data for SuccessfulEvent events raised by the Venus contract.
type VenusSuccessfulEventIterator struct {
	Event *VenusSuccessfulEvent // Event containing the contract specifics and raw log

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
func (it *VenusSuccessfulEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusSuccessfulEvent)
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
		it.Event = new(VenusSuccessfulEvent)
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
func (it *VenusSuccessfulEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusSuccessfulEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusSuccessfulEvent represents a SuccessfulEvent event raised by the Venus contract.
type VenusSuccessfulEvent struct {
	EventIndex *big.Int
	Sender     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSuccessfulEvent is a free log retrieval operation binding the contract event 0x49a61a3517534657df66eaaaef62f55a830c07d22ca1760e0eff4a2c823e0bc9.
//
// Solidity: event SuccessfulEvent(uint256 eventIndex, address sender)
func (_Venus *VenusFilterer) FilterSuccessfulEvent(opts *bind.FilterOpts) (*VenusSuccessfulEventIterator, error) {

	logs, sub, err := _Venus.contract.FilterLogs(opts, "SuccessfulEvent")
	if err != nil {
		return nil, err
	}
	return &VenusSuccessfulEventIterator{contract: _Venus.contract, event: "SuccessfulEvent", logs: logs, sub: sub}, nil
}

// WatchSuccessfulEvent is a free log subscription operation binding the contract event 0x49a61a3517534657df66eaaaef62f55a830c07d22ca1760e0eff4a2c823e0bc9.
//
// Solidity: event SuccessfulEvent(uint256 eventIndex, address sender)
func (_Venus *VenusFilterer) WatchSuccessfulEvent(opts *bind.WatchOpts, sink chan<- *VenusSuccessfulEvent) (event.Subscription, error) {

	logs, sub, err := _Venus.contract.WatchLogs(opts, "SuccessfulEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusSuccessfulEvent)
				if err := _Venus.contract.UnpackLog(event, "SuccessfulEvent", log); err != nil {
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

// ParseSuccessfulEvent is a log parse operation binding the contract event 0x49a61a3517534657df66eaaaef62f55a830c07d22ca1760e0eff4a2c823e0bc9.
//
// Solidity: event SuccessfulEvent(uint256 eventIndex, address sender)
func (_Venus *VenusFilterer) ParseSuccessfulEvent(log types.Log) (*VenusSuccessfulEvent, error) {
	event := new(VenusSuccessfulEvent)
	if err := _Venus.contract.UnpackLog(event, "SuccessfulEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusSuccessfulReceiptIterator is returned from FilterSuccessfulReceipt and is used to iterate over the raw logs and unpacked data for SuccessfulReceipt events raised by the Venus contract.
type VenusSuccessfulReceiptIterator struct {
	Event *VenusSuccessfulReceipt // Event containing the contract specifics and raw log

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
func (it *VenusSuccessfulReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusSuccessfulReceipt)
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
		it.Event = new(VenusSuccessfulReceipt)
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
func (it *VenusSuccessfulReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusSuccessfulReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusSuccessfulReceipt represents a SuccessfulReceipt event raised by the Venus contract.
type VenusSuccessfulReceipt struct {
	ReceiptIndex []byte
	ReceiptRLP   []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSuccessfulReceipt is a free log retrieval operation binding the contract event 0xd6728eaf25dd1431eb8afabc6f371b3379a6c7beb972e468f5bf992fc6e822d5.
//
// Solidity: event SuccessfulReceipt(bytes receiptIndex, bytes receiptRLP)
func (_Venus *VenusFilterer) FilterSuccessfulReceipt(opts *bind.FilterOpts) (*VenusSuccessfulReceiptIterator, error) {

	logs, sub, err := _Venus.contract.FilterLogs(opts, "SuccessfulReceipt")
	if err != nil {
		return nil, err
	}
	return &VenusSuccessfulReceiptIterator{contract: _Venus.contract, event: "SuccessfulReceipt", logs: logs, sub: sub}, nil
}

// WatchSuccessfulReceipt is a free log subscription operation binding the contract event 0xd6728eaf25dd1431eb8afabc6f371b3379a6c7beb972e468f5bf992fc6e822d5.
//
// Solidity: event SuccessfulReceipt(bytes receiptIndex, bytes receiptRLP)
func (_Venus *VenusFilterer) WatchSuccessfulReceipt(opts *bind.WatchOpts, sink chan<- *VenusSuccessfulReceipt) (event.Subscription, error) {

	logs, sub, err := _Venus.contract.WatchLogs(opts, "SuccessfulReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusSuccessfulReceipt)
				if err := _Venus.contract.UnpackLog(event, "SuccessfulReceipt", log); err != nil {
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

// ParseSuccessfulReceipt is a log parse operation binding the contract event 0xd6728eaf25dd1431eb8afabc6f371b3379a6c7beb972e468f5bf992fc6e822d5.
//
// Solidity: event SuccessfulReceipt(bytes receiptIndex, bytes receiptRLP)
func (_Venus *VenusFilterer) ParseSuccessfulReceipt(log types.Log) (*VenusSuccessfulReceipt, error) {
	event := new(VenusSuccessfulReceipt)
	if err := _Venus.contract.UnpackLog(event, "SuccessfulReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusTransmissionReceivedIterator is returned from FilterTransmissionReceived and is used to iterate over the raw logs and unpacked data for TransmissionReceived events raised by the Venus contract.
type VenusTransmissionReceivedIterator struct {
	Event *VenusTransmissionReceived // Event containing the contract specifics and raw log

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
func (it *VenusTransmissionReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusTransmissionReceived)
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
		it.Event = new(VenusTransmissionReceived)
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
func (it *VenusTransmissionReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusTransmissionReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusTransmissionReceived represents a TransmissionReceived event raised by the Venus contract.
type VenusTransmissionReceived struct {
	Message   [32]byte
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransmissionReceived is a free log retrieval operation binding the contract event 0x777932fb4709c8bfb29f9190e22e536eaf00fd97a76a15a41b54e1b619fe5863.
//
// Solidity: event TransmissionReceived(bytes32 message, uint64 timestamp)
func (_Venus *VenusFilterer) FilterTransmissionReceived(opts *bind.FilterOpts) (*VenusTransmissionReceivedIterator, error) {

	logs, sub, err := _Venus.contract.FilterLogs(opts, "TransmissionReceived")
	if err != nil {
		return nil, err
	}
	return &VenusTransmissionReceivedIterator{contract: _Venus.contract, event: "TransmissionReceived", logs: logs, sub: sub}, nil
}

// WatchTransmissionReceived is a free log subscription operation binding the contract event 0x777932fb4709c8bfb29f9190e22e536eaf00fd97a76a15a41b54e1b619fe5863.
//
// Solidity: event TransmissionReceived(bytes32 message, uint64 timestamp)
func (_Venus *VenusFilterer) WatchTransmissionReceived(opts *bind.WatchOpts, sink chan<- *VenusTransmissionReceived) (event.Subscription, error) {

	logs, sub, err := _Venus.contract.WatchLogs(opts, "TransmissionReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusTransmissionReceived)
				if err := _Venus.contract.UnpackLog(event, "TransmissionReceived", log); err != nil {
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

// ParseTransmissionReceived is a log parse operation binding the contract event 0x777932fb4709c8bfb29f9190e22e536eaf00fd97a76a15a41b54e1b619fe5863.
//
// Solidity: event TransmissionReceived(bytes32 message, uint64 timestamp)
func (_Venus *VenusFilterer) ParseTransmissionReceived(log types.Log) (*VenusTransmissionReceived, error) {
	event := new(VenusTransmissionReceived)
	if err := _Venus.contract.UnpackLog(event, "TransmissionReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VenusTransmitToHoustonIterator is returned from FilterTransmitToHouston and is used to iterate over the raw logs and unpacked data for TransmitToHouston events raised by the Venus contract.
type VenusTransmitToHoustonIterator struct {
	Event *VenusTransmitToHouston // Event containing the contract specifics and raw log

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
func (it *VenusTransmitToHoustonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VenusTransmitToHouston)
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
		it.Event = new(VenusTransmitToHouston)
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
func (it *VenusTransmitToHoustonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VenusTransmitToHoustonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VenusTransmitToHouston represents a TransmitToHouston event raised by the Venus contract.
type VenusTransmitToHouston struct {
	Message   [32]byte
	Timestamp uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransmitToHouston is a free log retrieval operation binding the contract event 0xc134cee5dddcc539de40a587edc2d37fb04e8ec090cb8dcf8ef0c9b23d41908d.
//
// Solidity: event TransmitToHouston(bytes32 message, uint64 timestamp)
func (_Venus *VenusFilterer) FilterTransmitToHouston(opts *bind.FilterOpts) (*VenusTransmitToHoustonIterator, error) {

	logs, sub, err := _Venus.contract.FilterLogs(opts, "TransmitToHouston")
	if err != nil {
		return nil, err
	}
	return &VenusTransmitToHoustonIterator{contract: _Venus.contract, event: "TransmitToHouston", logs: logs, sub: sub}, nil
}

// WatchTransmitToHouston is a free log subscription operation binding the contract event 0xc134cee5dddcc539de40a587edc2d37fb04e8ec090cb8dcf8ef0c9b23d41908d.
//
// Solidity: event TransmitToHouston(bytes32 message, uint64 timestamp)
func (_Venus *VenusFilterer) WatchTransmitToHouston(opts *bind.WatchOpts, sink chan<- *VenusTransmitToHouston) (event.Subscription, error) {

	logs, sub, err := _Venus.contract.WatchLogs(opts, "TransmitToHouston")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VenusTransmitToHouston)
				if err := _Venus.contract.UnpackLog(event, "TransmitToHouston", log); err != nil {
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

// ParseTransmitToHouston is a log parse operation binding the contract event 0xc134cee5dddcc539de40a587edc2d37fb04e8ec090cb8dcf8ef0c9b23d41908d.
//
// Solidity: event TransmitToHouston(bytes32 message, uint64 timestamp)
func (_Venus *VenusFilterer) ParseTransmitToHouston(log types.Log) (*VenusTransmitToHouston, error) {
	event := new(VenusTransmitToHouston)
	if err := _Venus.contract.UnpackLog(event, "TransmitToHouston", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
