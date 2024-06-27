// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generalmiddleware

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

// GeneralMiddlewareMetaData contains all meta data concerning the GeneralMiddleware contract.
var GeneralMiddlewareMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"mwId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_middleware\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MW_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authorizeMiddleware\",\"inputs\":[{\"name\":\"middleware\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authorizedMws\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mw\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onRecvMWAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ucPacket\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"mwIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mwAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvMWPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ucPacket\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"mwIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mwAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvMWTimeout\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ucPacket\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"mwIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mwAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ucPacket\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onUniversalAcknowledgement\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendMWPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"srcMwIds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacketWithFee\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setDefaultMw\",\"inputs\":[{\"name\":\"_middleware\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecvMWAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"mwId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"appData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mwData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"ack\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecvMWPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"mwId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"appData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mwData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecvMWTimeout\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"mwId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"appData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mwData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendMWPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"mwId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"appData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"mwData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UCHPacketSent\",\"inputs\":[{\"name\":\"source\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"UnauthorizedIbcMiddleware\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ackAddressMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ackDataTooShort\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidChannelId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"unauthorizedChannel\",\"inputs\":[]}]",
}

// GeneralMiddlewareABI is the input ABI used to generate the binding from.
// Deprecated: Use GeneralMiddlewareMetaData.ABI instead.
var GeneralMiddlewareABI = GeneralMiddlewareMetaData.ABI

// GeneralMiddleware is an auto generated Go binding around an Ethereum contract.
type GeneralMiddleware struct {
	GeneralMiddlewareCaller     // Read-only binding to the contract
	GeneralMiddlewareTransactor // Write-only binding to the contract
	GeneralMiddlewareFilterer   // Log filterer for contract events
}

// GeneralMiddlewareCaller is an auto generated read-only Go binding around an Ethereum contract.
type GeneralMiddlewareCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneralMiddlewareTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GeneralMiddlewareTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneralMiddlewareFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GeneralMiddlewareFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneralMiddlewareSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GeneralMiddlewareSession struct {
	Contract     *GeneralMiddleware // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GeneralMiddlewareCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GeneralMiddlewareCallerSession struct {
	Contract *GeneralMiddlewareCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GeneralMiddlewareTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GeneralMiddlewareTransactorSession struct {
	Contract     *GeneralMiddlewareTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GeneralMiddlewareRaw is an auto generated low-level Go binding around an Ethereum contract.
type GeneralMiddlewareRaw struct {
	Contract *GeneralMiddleware // Generic contract binding to access the raw methods on
}

// GeneralMiddlewareCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GeneralMiddlewareCallerRaw struct {
	Contract *GeneralMiddlewareCaller // Generic read-only contract binding to access the raw methods on
}

// GeneralMiddlewareTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GeneralMiddlewareTransactorRaw struct {
	Contract *GeneralMiddlewareTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGeneralMiddleware creates a new instance of GeneralMiddleware, bound to a specific deployed contract.
func NewGeneralMiddleware(address common.Address, backend bind.ContractBackend) (*GeneralMiddleware, error) {
	contract, err := bindGeneralMiddleware(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GeneralMiddleware{GeneralMiddlewareCaller: GeneralMiddlewareCaller{contract: contract}, GeneralMiddlewareTransactor: GeneralMiddlewareTransactor{contract: contract}, GeneralMiddlewareFilterer: GeneralMiddlewareFilterer{contract: contract}}, nil
}

// NewGeneralMiddlewareCaller creates a new read-only instance of GeneralMiddleware, bound to a specific deployed contract.
func NewGeneralMiddlewareCaller(address common.Address, caller bind.ContractCaller) (*GeneralMiddlewareCaller, error) {
	contract, err := bindGeneralMiddleware(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GeneralMiddlewareCaller{contract: contract}, nil
}

// NewGeneralMiddlewareTransactor creates a new write-only instance of GeneralMiddleware, bound to a specific deployed contract.
func NewGeneralMiddlewareTransactor(address common.Address, transactor bind.ContractTransactor) (*GeneralMiddlewareTransactor, error) {
	contract, err := bindGeneralMiddleware(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GeneralMiddlewareTransactor{contract: contract}, nil
}

// NewGeneralMiddlewareFilterer creates a new log filterer instance of GeneralMiddleware, bound to a specific deployed contract.
func NewGeneralMiddlewareFilterer(address common.Address, filterer bind.ContractFilterer) (*GeneralMiddlewareFilterer, error) {
	contract, err := bindGeneralMiddleware(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GeneralMiddlewareFilterer{contract: contract}, nil
}

// bindGeneralMiddleware binds a generic wrapper to an already deployed contract.
func bindGeneralMiddleware(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GeneralMiddlewareMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneralMiddleware *GeneralMiddlewareRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GeneralMiddleware.Contract.GeneralMiddlewareCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneralMiddleware *GeneralMiddlewareRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.GeneralMiddlewareTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneralMiddleware *GeneralMiddlewareRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.GeneralMiddlewareTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneralMiddleware *GeneralMiddlewareCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GeneralMiddleware.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneralMiddleware *GeneralMiddlewareTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneralMiddleware *GeneralMiddlewareTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.contract.Transact(opts, method, params...)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256)
func (_GeneralMiddleware *GeneralMiddlewareCaller) MWID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GeneralMiddleware.contract.Call(opts, &out, "MW_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256)
func (_GeneralMiddleware *GeneralMiddlewareSession) MWID() (*big.Int, error) {
	return _GeneralMiddleware.Contract.MWID(&_GeneralMiddleware.CallOpts)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256)
func (_GeneralMiddleware *GeneralMiddlewareCallerSession) MWID() (*big.Int, error) {
	return _GeneralMiddleware.Contract.MWID(&_GeneralMiddleware.CallOpts)
}

// AuthorizedMws is a free data retrieval call binding the contract method 0x3a7fbcbd.
//
// Solidity: function authorizedMws(address ) view returns(bool)
func (_GeneralMiddleware *GeneralMiddlewareCaller) AuthorizedMws(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GeneralMiddleware.contract.Call(opts, &out, "authorizedMws", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedMws is a free data retrieval call binding the contract method 0x3a7fbcbd.
//
// Solidity: function authorizedMws(address ) view returns(bool)
func (_GeneralMiddleware *GeneralMiddlewareSession) AuthorizedMws(arg0 common.Address) (bool, error) {
	return _GeneralMiddleware.Contract.AuthorizedMws(&_GeneralMiddleware.CallOpts, arg0)
}

// AuthorizedMws is a free data retrieval call binding the contract method 0x3a7fbcbd.
//
// Solidity: function authorizedMws(address ) view returns(bool)
func (_GeneralMiddleware *GeneralMiddlewareCallerSession) AuthorizedMws(arg0 common.Address) (bool, error) {
	return _GeneralMiddleware.Contract.AuthorizedMws(&_GeneralMiddleware.CallOpts, arg0)
}

// Mw is a free data retrieval call binding the contract method 0xa742d78c.
//
// Solidity: function mw() view returns(address)
func (_GeneralMiddleware *GeneralMiddlewareCaller) Mw(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GeneralMiddleware.contract.Call(opts, &out, "mw")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mw is a free data retrieval call binding the contract method 0xa742d78c.
//
// Solidity: function mw() view returns(address)
func (_GeneralMiddleware *GeneralMiddlewareSession) Mw() (common.Address, error) {
	return _GeneralMiddleware.Contract.Mw(&_GeneralMiddleware.CallOpts)
}

// Mw is a free data retrieval call binding the contract method 0xa742d78c.
//
// Solidity: function mw() view returns(address)
func (_GeneralMiddleware *GeneralMiddlewareCallerSession) Mw() (common.Address, error) {
	return _GeneralMiddleware.Contract.Mw(&_GeneralMiddleware.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GeneralMiddleware *GeneralMiddlewareCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GeneralMiddleware.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GeneralMiddleware *GeneralMiddlewareSession) Owner() (common.Address, error) {
	return _GeneralMiddleware.Contract.Owner(&_GeneralMiddleware.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GeneralMiddleware *GeneralMiddlewareCallerSession) Owner() (common.Address, error) {
	return _GeneralMiddleware.Contract.Owner(&_GeneralMiddleware.CallOpts)
}

// AuthorizeMiddleware is a paid mutator transaction binding the contract method 0x3b90b042.
//
// Solidity: function authorizeMiddleware(address middleware) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactor) AuthorizeMiddleware(opts *bind.TransactOpts, middleware common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.Transact(opts, "authorizeMiddleware", middleware)
}

// AuthorizeMiddleware is a paid mutator transaction binding the contract method 0x3b90b042.
//
// Solidity: function authorizeMiddleware(address middleware) returns()
func (_GeneralMiddleware *GeneralMiddlewareSession) AuthorizeMiddleware(middleware common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.AuthorizeMiddleware(&_GeneralMiddleware.TransactOpts, middleware)
}

// AuthorizeMiddleware is a paid mutator transaction binding the contract method 0x3b90b042.
//
// Solidity: function authorizeMiddleware(address middleware) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) AuthorizeMiddleware(middleware common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.AuthorizeMiddleware(&_GeneralMiddleware.TransactOpts, middleware)
}

// OnRecvMWAck is a paid mutator transaction binding the contract method 0xa968dc60.
//
// Solidity: function onRecvMWAck(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs, (bool,bytes) ack) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactor) OnRecvMWAck(opts *bind.TransactOpts, channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address, ack AckPacket) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.Transact(opts, "onRecvMWAck", channelId, ucPacket, mwIndex, mwAddrs, ack)
}

// OnRecvMWAck is a paid mutator transaction binding the contract method 0xa968dc60.
//
// Solidity: function onRecvMWAck(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs, (bool,bytes) ack) returns()
func (_GeneralMiddleware *GeneralMiddlewareSession) OnRecvMWAck(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address, ack AckPacket) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.OnRecvMWAck(&_GeneralMiddleware.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs, ack)
}

// OnRecvMWAck is a paid mutator transaction binding the contract method 0xa968dc60.
//
// Solidity: function onRecvMWAck(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs, (bool,bytes) ack) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) OnRecvMWAck(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address, ack AckPacket) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.OnRecvMWAck(&_GeneralMiddleware.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs, ack)
}

// OnRecvMWPacket is a paid mutator transaction binding the contract method 0xce2ce646.
//
// Solidity: function onRecvMWPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns((bool,bytes) ackPacket)
func (_GeneralMiddleware *GeneralMiddlewareTransactor) OnRecvMWPacket(opts *bind.TransactOpts, channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.Transact(opts, "onRecvMWPacket", channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWPacket is a paid mutator transaction binding the contract method 0xce2ce646.
//
// Solidity: function onRecvMWPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns((bool,bytes) ackPacket)
func (_GeneralMiddleware *GeneralMiddlewareSession) OnRecvMWPacket(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.OnRecvMWPacket(&_GeneralMiddleware.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWPacket is a paid mutator transaction binding the contract method 0xce2ce646.
//
// Solidity: function onRecvMWPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns((bool,bytes) ackPacket)
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) OnRecvMWPacket(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.OnRecvMWPacket(&_GeneralMiddleware.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWTimeout is a paid mutator transaction binding the contract method 0xc2cfcd73.
//
// Solidity: function onRecvMWTimeout(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactor) OnRecvMWTimeout(opts *bind.TransactOpts, channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.Transact(opts, "onRecvMWTimeout", channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWTimeout is a paid mutator transaction binding the contract method 0xc2cfcd73.
//
// Solidity: function onRecvMWTimeout(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns()
func (_GeneralMiddleware *GeneralMiddlewareSession) OnRecvMWTimeout(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.OnRecvMWTimeout(&_GeneralMiddleware.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWTimeout is a paid mutator transaction binding the contract method 0xc2cfcd73.
//
// Solidity: function onRecvMWTimeout(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) OnRecvMWTimeout(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.OnRecvMWTimeout(&_GeneralMiddleware.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket) returns((bool,bytes) ackPacket)
func (_GeneralMiddleware *GeneralMiddlewareTransactor) OnRecvUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, ucPacket UniversalPacket) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.Transact(opts, "onRecvUniversalPacket", channelId, ucPacket)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket) returns((bool,bytes) ackPacket)
func (_GeneralMiddleware *GeneralMiddlewareSession) OnRecvUniversalPacket(channelId [32]byte, ucPacket UniversalPacket) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.OnRecvUniversalPacket(&_GeneralMiddleware.TransactOpts, channelId, ucPacket)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket) returns((bool,bytes) ackPacket)
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) OnRecvUniversalPacket(channelId [32]byte, ucPacket UniversalPacket) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.OnRecvUniversalPacket(&_GeneralMiddleware.TransactOpts, channelId, ucPacket)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactor) OnTimeoutUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.Transact(opts, "onTimeoutUniversalPacket", channelId, packet)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_GeneralMiddleware *GeneralMiddlewareSession) OnTimeoutUniversalPacket(channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.OnTimeoutUniversalPacket(&_GeneralMiddleware.TransactOpts, channelId, packet)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) OnTimeoutUniversalPacket(channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.OnTimeoutUniversalPacket(&_GeneralMiddleware.TransactOpts, channelId, packet)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactor) OnUniversalAcknowledgement(opts *bind.TransactOpts, channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.Transact(opts, "onUniversalAcknowledgement", channelId, packet, ack)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_GeneralMiddleware *GeneralMiddlewareSession) OnUniversalAcknowledgement(channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.OnUniversalAcknowledgement(&_GeneralMiddleware.TransactOpts, channelId, packet, ack)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) OnUniversalAcknowledgement(channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.OnUniversalAcknowledgement(&_GeneralMiddleware.TransactOpts, channelId, packet, ack)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GeneralMiddleware *GeneralMiddlewareSession) RenounceOwnership() (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.RenounceOwnership(&_GeneralMiddleware.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.RenounceOwnership(&_GeneralMiddleware.TransactOpts)
}

// SendMWPacket is a paid mutator transaction binding the contract method 0x1b67943d.
//
// Solidity: function sendMWPacket(bytes32 channelId, bytes32 srcPortAddr, bytes32 destPortAddr, uint256 srcMwIds, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_GeneralMiddleware *GeneralMiddlewareTransactor) SendMWPacket(opts *bind.TransactOpts, channelId [32]byte, srcPortAddr [32]byte, destPortAddr [32]byte, srcMwIds *big.Int, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.Transact(opts, "sendMWPacket", channelId, srcPortAddr, destPortAddr, srcMwIds, appData, timeoutTimestamp)
}

// SendMWPacket is a paid mutator transaction binding the contract method 0x1b67943d.
//
// Solidity: function sendMWPacket(bytes32 channelId, bytes32 srcPortAddr, bytes32 destPortAddr, uint256 srcMwIds, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_GeneralMiddleware *GeneralMiddlewareSession) SendMWPacket(channelId [32]byte, srcPortAddr [32]byte, destPortAddr [32]byte, srcMwIds *big.Int, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.SendMWPacket(&_GeneralMiddleware.TransactOpts, channelId, srcPortAddr, destPortAddr, srcMwIds, appData, timeoutTimestamp)
}

// SendMWPacket is a paid mutator transaction binding the contract method 0x1b67943d.
//
// Solidity: function sendMWPacket(bytes32 channelId, bytes32 srcPortAddr, bytes32 destPortAddr, uint256 srcMwIds, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) SendMWPacket(channelId [32]byte, srcPortAddr [32]byte, destPortAddr [32]byte, srcMwIds *big.Int, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.SendMWPacket(&_GeneralMiddleware.TransactOpts, channelId, srcPortAddr, destPortAddr, srcMwIds, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_GeneralMiddleware *GeneralMiddlewareTransactor) SendUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.Transact(opts, "sendUniversalPacket", channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_GeneralMiddleware *GeneralMiddlewareSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.SendUniversalPacket(&_GeneralMiddleware.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.SendUniversalPacket(&_GeneralMiddleware.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_GeneralMiddleware *GeneralMiddlewareTransactor) SendUniversalPacketWithFee(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.Transact(opts, "sendUniversalPacketWithFee", channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_GeneralMiddleware *GeneralMiddlewareSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.SendUniversalPacketWithFee(&_GeneralMiddleware.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.SendUniversalPacketWithFee(&_GeneralMiddleware.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SetDefaultMw is a paid mutator transaction binding the contract method 0x00e82cef.
//
// Solidity: function setDefaultMw(address _middleware) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactor) SetDefaultMw(opts *bind.TransactOpts, _middleware common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.Transact(opts, "setDefaultMw", _middleware)
}

// SetDefaultMw is a paid mutator transaction binding the contract method 0x00e82cef.
//
// Solidity: function setDefaultMw(address _middleware) returns()
func (_GeneralMiddleware *GeneralMiddlewareSession) SetDefaultMw(_middleware common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.SetDefaultMw(&_GeneralMiddleware.TransactOpts, _middleware)
}

// SetDefaultMw is a paid mutator transaction binding the contract method 0x00e82cef.
//
// Solidity: function setDefaultMw(address _middleware) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) SetDefaultMw(_middleware common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.SetDefaultMw(&_GeneralMiddleware.TransactOpts, _middleware)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GeneralMiddleware *GeneralMiddlewareSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.TransferOwnership(&_GeneralMiddleware.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.TransferOwnership(&_GeneralMiddleware.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneralMiddleware.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GeneralMiddleware *GeneralMiddlewareSession) Receive() (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.Receive(&_GeneralMiddleware.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GeneralMiddleware *GeneralMiddlewareTransactorSession) Receive() (*types.Transaction, error) {
	return _GeneralMiddleware.Contract.Receive(&_GeneralMiddleware.TransactOpts)
}

// GeneralMiddlewareOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GeneralMiddleware contract.
type GeneralMiddlewareOwnershipTransferredIterator struct {
	Event *GeneralMiddlewareOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GeneralMiddlewareOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralMiddlewareOwnershipTransferred)
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
		it.Event = new(GeneralMiddlewareOwnershipTransferred)
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
func (it *GeneralMiddlewareOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralMiddlewareOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralMiddlewareOwnershipTransferred represents a OwnershipTransferred event raised by the GeneralMiddleware contract.
type GeneralMiddlewareOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GeneralMiddleware *GeneralMiddlewareFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GeneralMiddlewareOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GeneralMiddleware.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GeneralMiddlewareOwnershipTransferredIterator{contract: _GeneralMiddleware.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GeneralMiddleware *GeneralMiddlewareFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GeneralMiddlewareOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GeneralMiddleware.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralMiddlewareOwnershipTransferred)
				if err := _GeneralMiddleware.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GeneralMiddleware *GeneralMiddlewareFilterer) ParseOwnershipTransferred(log types.Log) (*GeneralMiddlewareOwnershipTransferred, error) {
	event := new(GeneralMiddlewareOwnershipTransferred)
	if err := _GeneralMiddleware.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GeneralMiddlewareRecvMWAckIterator is returned from FilterRecvMWAck and is used to iterate over the raw logs and unpacked data for RecvMWAck events raised by the GeneralMiddleware contract.
type GeneralMiddlewareRecvMWAckIterator struct {
	Event *GeneralMiddlewareRecvMWAck // Event containing the contract specifics and raw log

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
func (it *GeneralMiddlewareRecvMWAckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralMiddlewareRecvMWAck)
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
		it.Event = new(GeneralMiddlewareRecvMWAck)
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
func (it *GeneralMiddlewareRecvMWAckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralMiddlewareRecvMWAckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralMiddlewareRecvMWAck represents a RecvMWAck event raised by the GeneralMiddleware contract.
type GeneralMiddlewareRecvMWAck struct {
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
func (_GeneralMiddleware *GeneralMiddlewareFilterer) FilterRecvMWAck(opts *bind.FilterOpts, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (*GeneralMiddlewareRecvMWAckIterator, error) {

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

	logs, sub, err := _GeneralMiddleware.contract.FilterLogs(opts, "RecvMWAck", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return &GeneralMiddlewareRecvMWAckIterator{contract: _GeneralMiddleware.contract, event: "RecvMWAck", logs: logs, sub: sub}, nil
}

// WatchRecvMWAck is a free log subscription operation binding the contract event 0x0ad1351cd77bd217ef00c2ab94f17283f9a51d4ebd8d189c151760d8075f01c2.
//
// Solidity: event RecvMWAck(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, bytes mwData, (bool,bytes) ack)
func (_GeneralMiddleware *GeneralMiddlewareFilterer) WatchRecvMWAck(opts *bind.WatchOpts, sink chan<- *GeneralMiddlewareRecvMWAck, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GeneralMiddleware.contract.WatchLogs(opts, "RecvMWAck", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralMiddlewareRecvMWAck)
				if err := _GeneralMiddleware.contract.UnpackLog(event, "RecvMWAck", log); err != nil {
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
func (_GeneralMiddleware *GeneralMiddlewareFilterer) ParseRecvMWAck(log types.Log) (*GeneralMiddlewareRecvMWAck, error) {
	event := new(GeneralMiddlewareRecvMWAck)
	if err := _GeneralMiddleware.contract.UnpackLog(event, "RecvMWAck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GeneralMiddlewareRecvMWPacketIterator is returned from FilterRecvMWPacket and is used to iterate over the raw logs and unpacked data for RecvMWPacket events raised by the GeneralMiddleware contract.
type GeneralMiddlewareRecvMWPacketIterator struct {
	Event *GeneralMiddlewareRecvMWPacket // Event containing the contract specifics and raw log

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
func (it *GeneralMiddlewareRecvMWPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralMiddlewareRecvMWPacket)
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
		it.Event = new(GeneralMiddlewareRecvMWPacket)
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
func (it *GeneralMiddlewareRecvMWPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralMiddlewareRecvMWPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralMiddlewareRecvMWPacket represents a RecvMWPacket event raised by the GeneralMiddleware contract.
type GeneralMiddlewareRecvMWPacket struct {
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
func (_GeneralMiddleware *GeneralMiddlewareFilterer) FilterRecvMWPacket(opts *bind.FilterOpts, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (*GeneralMiddlewareRecvMWPacketIterator, error) {

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

	logs, sub, err := _GeneralMiddleware.contract.FilterLogs(opts, "RecvMWPacket", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return &GeneralMiddlewareRecvMWPacketIterator{contract: _GeneralMiddleware.contract, event: "RecvMWPacket", logs: logs, sub: sub}, nil
}

// WatchRecvMWPacket is a free log subscription operation binding the contract event 0xdbf4e1c85f2bbbba15b8c4ca75bb0ef2a3e496e986f4fde306aa5942a48ad180.
//
// Solidity: event RecvMWPacket(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, bytes mwData)
func (_GeneralMiddleware *GeneralMiddlewareFilterer) WatchRecvMWPacket(opts *bind.WatchOpts, sink chan<- *GeneralMiddlewareRecvMWPacket, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GeneralMiddleware.contract.WatchLogs(opts, "RecvMWPacket", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralMiddlewareRecvMWPacket)
				if err := _GeneralMiddleware.contract.UnpackLog(event, "RecvMWPacket", log); err != nil {
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
func (_GeneralMiddleware *GeneralMiddlewareFilterer) ParseRecvMWPacket(log types.Log) (*GeneralMiddlewareRecvMWPacket, error) {
	event := new(GeneralMiddlewareRecvMWPacket)
	if err := _GeneralMiddleware.contract.UnpackLog(event, "RecvMWPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GeneralMiddlewareRecvMWTimeoutIterator is returned from FilterRecvMWTimeout and is used to iterate over the raw logs and unpacked data for RecvMWTimeout events raised by the GeneralMiddleware contract.
type GeneralMiddlewareRecvMWTimeoutIterator struct {
	Event *GeneralMiddlewareRecvMWTimeout // Event containing the contract specifics and raw log

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
func (it *GeneralMiddlewareRecvMWTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralMiddlewareRecvMWTimeout)
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
		it.Event = new(GeneralMiddlewareRecvMWTimeout)
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
func (it *GeneralMiddlewareRecvMWTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralMiddlewareRecvMWTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralMiddlewareRecvMWTimeout represents a RecvMWTimeout event raised by the GeneralMiddleware contract.
type GeneralMiddlewareRecvMWTimeout struct {
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
func (_GeneralMiddleware *GeneralMiddlewareFilterer) FilterRecvMWTimeout(opts *bind.FilterOpts, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (*GeneralMiddlewareRecvMWTimeoutIterator, error) {

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

	logs, sub, err := _GeneralMiddleware.contract.FilterLogs(opts, "RecvMWTimeout", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return &GeneralMiddlewareRecvMWTimeoutIterator{contract: _GeneralMiddleware.contract, event: "RecvMWTimeout", logs: logs, sub: sub}, nil
}

// WatchRecvMWTimeout is a free log subscription operation binding the contract event 0xc2c6119a1960aa25d517e5173005d0d39b4f58db74eda3eb48acda26cac7ca62.
//
// Solidity: event RecvMWTimeout(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, bytes mwData)
func (_GeneralMiddleware *GeneralMiddlewareFilterer) WatchRecvMWTimeout(opts *bind.WatchOpts, sink chan<- *GeneralMiddlewareRecvMWTimeout, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GeneralMiddleware.contract.WatchLogs(opts, "RecvMWTimeout", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralMiddlewareRecvMWTimeout)
				if err := _GeneralMiddleware.contract.UnpackLog(event, "RecvMWTimeout", log); err != nil {
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
func (_GeneralMiddleware *GeneralMiddlewareFilterer) ParseRecvMWTimeout(log types.Log) (*GeneralMiddlewareRecvMWTimeout, error) {
	event := new(GeneralMiddlewareRecvMWTimeout)
	if err := _GeneralMiddleware.contract.UnpackLog(event, "RecvMWTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GeneralMiddlewareSendMWPacketIterator is returned from FilterSendMWPacket and is used to iterate over the raw logs and unpacked data for SendMWPacket events raised by the GeneralMiddleware contract.
type GeneralMiddlewareSendMWPacketIterator struct {
	Event *GeneralMiddlewareSendMWPacket // Event containing the contract specifics and raw log

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
func (it *GeneralMiddlewareSendMWPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralMiddlewareSendMWPacket)
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
		it.Event = new(GeneralMiddlewareSendMWPacket)
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
func (it *GeneralMiddlewareSendMWPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralMiddlewareSendMWPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralMiddlewareSendMWPacket represents a SendMWPacket event raised by the GeneralMiddleware contract.
type GeneralMiddlewareSendMWPacket struct {
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
func (_GeneralMiddleware *GeneralMiddlewareFilterer) FilterSendMWPacket(opts *bind.FilterOpts, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (*GeneralMiddlewareSendMWPacketIterator, error) {

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

	logs, sub, err := _GeneralMiddleware.contract.FilterLogs(opts, "SendMWPacket", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return &GeneralMiddlewareSendMWPacketIterator{contract: _GeneralMiddleware.contract, event: "SendMWPacket", logs: logs, sub: sub}, nil
}

// WatchSendMWPacket is a free log subscription operation binding the contract event 0x7c1d77d7984ba9b37263b84eda35413a5b3911d71bd1f24b60bb43cfcaf539f7.
//
// Solidity: event SendMWPacket(bytes32 indexed channelId, bytes32 indexed srcPortAddr, bytes32 indexed destPortAddr, uint256 mwId, bytes appData, uint64 timeoutTimestamp, bytes mwData)
func (_GeneralMiddleware *GeneralMiddlewareFilterer) WatchSendMWPacket(opts *bind.WatchOpts, sink chan<- *GeneralMiddlewareSendMWPacket, channelId [][32]byte, srcPortAddr [][32]byte, destPortAddr [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GeneralMiddleware.contract.WatchLogs(opts, "SendMWPacket", channelIdRule, srcPortAddrRule, destPortAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralMiddlewareSendMWPacket)
				if err := _GeneralMiddleware.contract.UnpackLog(event, "SendMWPacket", log); err != nil {
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
func (_GeneralMiddleware *GeneralMiddlewareFilterer) ParseSendMWPacket(log types.Log) (*GeneralMiddlewareSendMWPacket, error) {
	event := new(GeneralMiddlewareSendMWPacket)
	if err := _GeneralMiddleware.contract.UnpackLog(event, "SendMWPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GeneralMiddlewareUCHPacketSentIterator is returned from FilterUCHPacketSent and is used to iterate over the raw logs and unpacked data for UCHPacketSent events raised by the GeneralMiddleware contract.
type GeneralMiddlewareUCHPacketSentIterator struct {
	Event *GeneralMiddlewareUCHPacketSent // Event containing the contract specifics and raw log

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
func (it *GeneralMiddlewareUCHPacketSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GeneralMiddlewareUCHPacketSent)
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
		it.Event = new(GeneralMiddlewareUCHPacketSent)
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
func (it *GeneralMiddlewareUCHPacketSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GeneralMiddlewareUCHPacketSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GeneralMiddlewareUCHPacketSent represents a UCHPacketSent event raised by the GeneralMiddleware contract.
type GeneralMiddlewareUCHPacketSent struct {
	Source      common.Address
	Destination [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUCHPacketSent is a free log retrieval operation binding the contract event 0x9831d8c66285bfd33de069ced58ad437d6bf08f63446bf06c3713e40b4b7e873.
//
// Solidity: event UCHPacketSent(address source, bytes32 destination)
func (_GeneralMiddleware *GeneralMiddlewareFilterer) FilterUCHPacketSent(opts *bind.FilterOpts) (*GeneralMiddlewareUCHPacketSentIterator, error) {

	logs, sub, err := _GeneralMiddleware.contract.FilterLogs(opts, "UCHPacketSent")
	if err != nil {
		return nil, err
	}
	return &GeneralMiddlewareUCHPacketSentIterator{contract: _GeneralMiddleware.contract, event: "UCHPacketSent", logs: logs, sub: sub}, nil
}

// WatchUCHPacketSent is a free log subscription operation binding the contract event 0x9831d8c66285bfd33de069ced58ad437d6bf08f63446bf06c3713e40b4b7e873.
//
// Solidity: event UCHPacketSent(address source, bytes32 destination)
func (_GeneralMiddleware *GeneralMiddlewareFilterer) WatchUCHPacketSent(opts *bind.WatchOpts, sink chan<- *GeneralMiddlewareUCHPacketSent) (event.Subscription, error) {

	logs, sub, err := _GeneralMiddleware.contract.WatchLogs(opts, "UCHPacketSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GeneralMiddlewareUCHPacketSent)
				if err := _GeneralMiddleware.contract.UnpackLog(event, "UCHPacketSent", log); err != nil {
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

// ParseUCHPacketSent is a log parse operation binding the contract event 0x9831d8c66285bfd33de069ced58ad437d6bf08f63446bf06c3713e40b4b7e873.
//
// Solidity: event UCHPacketSent(address source, bytes32 destination)
func (_GeneralMiddleware *GeneralMiddlewareFilterer) ParseUCHPacketSent(log types.Log) (*GeneralMiddlewareUCHPacketSent, error) {
	event := new(GeneralMiddlewareUCHPacketSent)
	if err := _GeneralMiddleware.contract.UnpackLog(event, "UCHPacketSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
