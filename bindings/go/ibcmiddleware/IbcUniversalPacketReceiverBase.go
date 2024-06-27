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

// UniversalPacket is an auto generated low-level Go binding around an user-defined struct.
type UniversalPacket struct {
	SrcPortAddr  [32]byte
	MwBitmap     *big.Int
	DestPortAddr [32]byte
	AppData      []byte
}

// IbcUniversalPacketReceiverBaseMetaData contains all meta data concerning the IbcUniversalPacketReceiverBase contract.
var IbcUniversalPacketReceiverBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"authorizeMiddleware\",\"inputs\":[{\"name\":\"middleware\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authorizedChannelIds\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authorizedMws\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mw\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onRecvUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ucPacket\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onUniversalAcknowledgement\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultMw\",\"inputs\":[{\"name\":\"_middleware\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"UnauthorizedIbcMiddleware\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ackAddressMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ackDataTooShort\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidChannelId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"unauthorizedChannel\",\"inputs\":[]}]",
}

// IbcUniversalPacketReceiverBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcUniversalPacketReceiverBaseMetaData.ABI instead.
var IbcUniversalPacketReceiverBaseABI = IbcUniversalPacketReceiverBaseMetaData.ABI

// IbcUniversalPacketReceiverBase is an auto generated Go binding around an Ethereum contract.
type IbcUniversalPacketReceiverBase struct {
	IbcUniversalPacketReceiverBaseCaller     // Read-only binding to the contract
	IbcUniversalPacketReceiverBaseTransactor // Write-only binding to the contract
	IbcUniversalPacketReceiverBaseFilterer   // Log filterer for contract events
}

// IbcUniversalPacketReceiverBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcUniversalPacketReceiverBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUniversalPacketReceiverBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcUniversalPacketReceiverBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUniversalPacketReceiverBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcUniversalPacketReceiverBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUniversalPacketReceiverBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcUniversalPacketReceiverBaseSession struct {
	Contract     *IbcUniversalPacketReceiverBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IbcUniversalPacketReceiverBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcUniversalPacketReceiverBaseCallerSession struct {
	Contract *IbcUniversalPacketReceiverBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// IbcUniversalPacketReceiverBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcUniversalPacketReceiverBaseTransactorSession struct {
	Contract     *IbcUniversalPacketReceiverBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// IbcUniversalPacketReceiverBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcUniversalPacketReceiverBaseRaw struct {
	Contract *IbcUniversalPacketReceiverBase // Generic contract binding to access the raw methods on
}

// IbcUniversalPacketReceiverBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcUniversalPacketReceiverBaseCallerRaw struct {
	Contract *IbcUniversalPacketReceiverBaseCaller // Generic read-only contract binding to access the raw methods on
}

// IbcUniversalPacketReceiverBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcUniversalPacketReceiverBaseTransactorRaw struct {
	Contract *IbcUniversalPacketReceiverBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcUniversalPacketReceiverBase creates a new instance of IbcUniversalPacketReceiverBase, bound to a specific deployed contract.
func NewIbcUniversalPacketReceiverBase(address common.Address, backend bind.ContractBackend) (*IbcUniversalPacketReceiverBase, error) {
	contract, err := bindIbcUniversalPacketReceiverBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalPacketReceiverBase{IbcUniversalPacketReceiverBaseCaller: IbcUniversalPacketReceiverBaseCaller{contract: contract}, IbcUniversalPacketReceiverBaseTransactor: IbcUniversalPacketReceiverBaseTransactor{contract: contract}, IbcUniversalPacketReceiverBaseFilterer: IbcUniversalPacketReceiverBaseFilterer{contract: contract}}, nil
}

// NewIbcUniversalPacketReceiverBaseCaller creates a new read-only instance of IbcUniversalPacketReceiverBase, bound to a specific deployed contract.
func NewIbcUniversalPacketReceiverBaseCaller(address common.Address, caller bind.ContractCaller) (*IbcUniversalPacketReceiverBaseCaller, error) {
	contract, err := bindIbcUniversalPacketReceiverBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalPacketReceiverBaseCaller{contract: contract}, nil
}

// NewIbcUniversalPacketReceiverBaseTransactor creates a new write-only instance of IbcUniversalPacketReceiverBase, bound to a specific deployed contract.
func NewIbcUniversalPacketReceiverBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcUniversalPacketReceiverBaseTransactor, error) {
	contract, err := bindIbcUniversalPacketReceiverBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalPacketReceiverBaseTransactor{contract: contract}, nil
}

// NewIbcUniversalPacketReceiverBaseFilterer creates a new log filterer instance of IbcUniversalPacketReceiverBase, bound to a specific deployed contract.
func NewIbcUniversalPacketReceiverBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcUniversalPacketReceiverBaseFilterer, error) {
	contract, err := bindIbcUniversalPacketReceiverBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalPacketReceiverBaseFilterer{contract: contract}, nil
}

// bindIbcUniversalPacketReceiverBase binds a generic wrapper to an already deployed contract.
func bindIbcUniversalPacketReceiverBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcUniversalPacketReceiverBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcUniversalPacketReceiverBase.Contract.IbcUniversalPacketReceiverBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.IbcUniversalPacketReceiverBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.IbcUniversalPacketReceiverBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcUniversalPacketReceiverBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.contract.Transact(opts, method, params...)
}

// AuthorizedChannelIds is a free data retrieval call binding the contract method 0x2eed7c70.
//
// Solidity: function authorizedChannelIds(bytes32 ) view returns(bool)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseCaller) AuthorizedChannelIds(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _IbcUniversalPacketReceiverBase.contract.Call(opts, &out, "authorizedChannelIds", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedChannelIds is a free data retrieval call binding the contract method 0x2eed7c70.
//
// Solidity: function authorizedChannelIds(bytes32 ) view returns(bool)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseSession) AuthorizedChannelIds(arg0 [32]byte) (bool, error) {
	return _IbcUniversalPacketReceiverBase.Contract.AuthorizedChannelIds(&_IbcUniversalPacketReceiverBase.CallOpts, arg0)
}

// AuthorizedChannelIds is a free data retrieval call binding the contract method 0x2eed7c70.
//
// Solidity: function authorizedChannelIds(bytes32 ) view returns(bool)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseCallerSession) AuthorizedChannelIds(arg0 [32]byte) (bool, error) {
	return _IbcUniversalPacketReceiverBase.Contract.AuthorizedChannelIds(&_IbcUniversalPacketReceiverBase.CallOpts, arg0)
}

// AuthorizedMws is a free data retrieval call binding the contract method 0x3a7fbcbd.
//
// Solidity: function authorizedMws(address ) view returns(bool)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseCaller) AuthorizedMws(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _IbcUniversalPacketReceiverBase.contract.Call(opts, &out, "authorizedMws", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedMws is a free data retrieval call binding the contract method 0x3a7fbcbd.
//
// Solidity: function authorizedMws(address ) view returns(bool)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseSession) AuthorizedMws(arg0 common.Address) (bool, error) {
	return _IbcUniversalPacketReceiverBase.Contract.AuthorizedMws(&_IbcUniversalPacketReceiverBase.CallOpts, arg0)
}

// AuthorizedMws is a free data retrieval call binding the contract method 0x3a7fbcbd.
//
// Solidity: function authorizedMws(address ) view returns(bool)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseCallerSession) AuthorizedMws(arg0 common.Address) (bool, error) {
	return _IbcUniversalPacketReceiverBase.Contract.AuthorizedMws(&_IbcUniversalPacketReceiverBase.CallOpts, arg0)
}

// Mw is a free data retrieval call binding the contract method 0xa742d78c.
//
// Solidity: function mw() view returns(address)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseCaller) Mw(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IbcUniversalPacketReceiverBase.contract.Call(opts, &out, "mw")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mw is a free data retrieval call binding the contract method 0xa742d78c.
//
// Solidity: function mw() view returns(address)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseSession) Mw() (common.Address, error) {
	return _IbcUniversalPacketReceiverBase.Contract.Mw(&_IbcUniversalPacketReceiverBase.CallOpts)
}

// Mw is a free data retrieval call binding the contract method 0xa742d78c.
//
// Solidity: function mw() view returns(address)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseCallerSession) Mw() (common.Address, error) {
	return _IbcUniversalPacketReceiverBase.Contract.Mw(&_IbcUniversalPacketReceiverBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IbcUniversalPacketReceiverBase.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseSession) Owner() (common.Address, error) {
	return _IbcUniversalPacketReceiverBase.Contract.Owner(&_IbcUniversalPacketReceiverBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseCallerSession) Owner() (common.Address, error) {
	return _IbcUniversalPacketReceiverBase.Contract.Owner(&_IbcUniversalPacketReceiverBase.CallOpts)
}

// AuthorizeMiddleware is a paid mutator transaction binding the contract method 0x3b90b042.
//
// Solidity: function authorizeMiddleware(address middleware) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactor) AuthorizeMiddleware(opts *bind.TransactOpts, middleware common.Address) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.contract.Transact(opts, "authorizeMiddleware", middleware)
}

// AuthorizeMiddleware is a paid mutator transaction binding the contract method 0x3b90b042.
//
// Solidity: function authorizeMiddleware(address middleware) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseSession) AuthorizeMiddleware(middleware common.Address) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.AuthorizeMiddleware(&_IbcUniversalPacketReceiverBase.TransactOpts, middleware)
}

// AuthorizeMiddleware is a paid mutator transaction binding the contract method 0x3b90b042.
//
// Solidity: function authorizeMiddleware(address middleware) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactorSession) AuthorizeMiddleware(middleware common.Address) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.AuthorizeMiddleware(&_IbcUniversalPacketReceiverBase.TransactOpts, middleware)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket) returns((bool,bytes) ackPacket)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactor) OnRecvUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, ucPacket UniversalPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.contract.Transact(opts, "onRecvUniversalPacket", channelId, ucPacket)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket) returns((bool,bytes) ackPacket)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseSession) OnRecvUniversalPacket(channelId [32]byte, ucPacket UniversalPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.OnRecvUniversalPacket(&_IbcUniversalPacketReceiverBase.TransactOpts, channelId, ucPacket)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket) returns((bool,bytes) ackPacket)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactorSession) OnRecvUniversalPacket(channelId [32]byte, ucPacket UniversalPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.OnRecvUniversalPacket(&_IbcUniversalPacketReceiverBase.TransactOpts, channelId, ucPacket)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactor) OnTimeoutUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.contract.Transact(opts, "onTimeoutUniversalPacket", channelId, packet)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseSession) OnTimeoutUniversalPacket(channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.OnTimeoutUniversalPacket(&_IbcUniversalPacketReceiverBase.TransactOpts, channelId, packet)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactorSession) OnTimeoutUniversalPacket(channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.OnTimeoutUniversalPacket(&_IbcUniversalPacketReceiverBase.TransactOpts, channelId, packet)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactor) OnUniversalAcknowledgement(opts *bind.TransactOpts, channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.contract.Transact(opts, "onUniversalAcknowledgement", channelId, packet, ack)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseSession) OnUniversalAcknowledgement(channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.OnUniversalAcknowledgement(&_IbcUniversalPacketReceiverBase.TransactOpts, channelId, packet, ack)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactorSession) OnUniversalAcknowledgement(channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.OnUniversalAcknowledgement(&_IbcUniversalPacketReceiverBase.TransactOpts, channelId, packet, ack)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.RenounceOwnership(&_IbcUniversalPacketReceiverBase.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.RenounceOwnership(&_IbcUniversalPacketReceiverBase.TransactOpts)
}

// SetDefaultMw is a paid mutator transaction binding the contract method 0x00e82cef.
//
// Solidity: function setDefaultMw(address _middleware) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactor) SetDefaultMw(opts *bind.TransactOpts, _middleware common.Address) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.contract.Transact(opts, "setDefaultMw", _middleware)
}

// SetDefaultMw is a paid mutator transaction binding the contract method 0x00e82cef.
//
// Solidity: function setDefaultMw(address _middleware) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseSession) SetDefaultMw(_middleware common.Address) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.SetDefaultMw(&_IbcUniversalPacketReceiverBase.TransactOpts, _middleware)
}

// SetDefaultMw is a paid mutator transaction binding the contract method 0x00e82cef.
//
// Solidity: function setDefaultMw(address _middleware) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactorSession) SetDefaultMw(_middleware common.Address) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.SetDefaultMw(&_IbcUniversalPacketReceiverBase.TransactOpts, _middleware)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.TransferOwnership(&_IbcUniversalPacketReceiverBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.TransferOwnership(&_IbcUniversalPacketReceiverBase.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseSession) Receive() (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.Receive(&_IbcUniversalPacketReceiverBase.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseTransactorSession) Receive() (*types.Transaction, error) {
	return _IbcUniversalPacketReceiverBase.Contract.Receive(&_IbcUniversalPacketReceiverBase.TransactOpts)
}

// IbcUniversalPacketReceiverBaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IbcUniversalPacketReceiverBase contract.
type IbcUniversalPacketReceiverBaseOwnershipTransferredIterator struct {
	Event *IbcUniversalPacketReceiverBaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IbcUniversalPacketReceiverBaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IbcUniversalPacketReceiverBaseOwnershipTransferred)
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
		it.Event = new(IbcUniversalPacketReceiverBaseOwnershipTransferred)
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
func (it *IbcUniversalPacketReceiverBaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IbcUniversalPacketReceiverBaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IbcUniversalPacketReceiverBaseOwnershipTransferred represents a OwnershipTransferred event raised by the IbcUniversalPacketReceiverBase contract.
type IbcUniversalPacketReceiverBaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IbcUniversalPacketReceiverBaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IbcUniversalPacketReceiverBase.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalPacketReceiverBaseOwnershipTransferredIterator{contract: _IbcUniversalPacketReceiverBase.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IbcUniversalPacketReceiverBaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IbcUniversalPacketReceiverBase.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IbcUniversalPacketReceiverBaseOwnershipTransferred)
				if err := _IbcUniversalPacketReceiverBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IbcUniversalPacketReceiverBase *IbcUniversalPacketReceiverBaseFilterer) ParseOwnershipTransferred(log types.Log) (*IbcUniversalPacketReceiverBaseOwnershipTransferred, error) {
	event := new(IbcUniversalPacketReceiverBaseOwnershipTransferred)
	if err := _IbcUniversalPacketReceiverBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
