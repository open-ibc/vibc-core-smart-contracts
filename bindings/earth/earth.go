// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package earth

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

// EarthMetaData contains all meta data concerning the Earth contract.
var EarthMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_middleware\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ackPackets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authorizeChannel\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authorizeMiddleware\",\"inputs\":[{\"name\":\"middleware\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authorizedChannelIds\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authorizedMws\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"generateAckPacket\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"srcPortAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"greet\",\"inputs\":[{\"name\":\"destPortAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"greetWithFee\",\"inputs\":[{\"name\":\"destPortAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"mw\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onRecvUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onUniversalAcknowledgement\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recvedPackets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultMw\",\"inputs\":[{\"name\":\"_middleware\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"timeoutPackets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"UnauthorizedIbcMiddleware\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ackAddressMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ackDataTooShort\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidChannelId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"unauthorizedChannel\",\"inputs\":[]}]",
}

// EarthABI is the input ABI used to generate the binding from.
// Deprecated: Use EarthMetaData.ABI instead.
var EarthABI = EarthMetaData.ABI

// Earth is an auto generated Go binding around an Ethereum contract.
type Earth struct {
	EarthCaller     // Read-only binding to the contract
	EarthTransactor // Write-only binding to the contract
	EarthFilterer   // Log filterer for contract events
}

// EarthCaller is an auto generated read-only Go binding around an Ethereum contract.
type EarthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EarthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EarthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EarthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EarthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EarthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EarthSession struct {
	Contract     *Earth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EarthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EarthCallerSession struct {
	Contract *EarthCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EarthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EarthTransactorSession struct {
	Contract     *EarthTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EarthRaw is an auto generated low-level Go binding around an Ethereum contract.
type EarthRaw struct {
	Contract *Earth // Generic contract binding to access the raw methods on
}

// EarthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EarthCallerRaw struct {
	Contract *EarthCaller // Generic read-only contract binding to access the raw methods on
}

// EarthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EarthTransactorRaw struct {
	Contract *EarthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEarth creates a new instance of Earth, bound to a specific deployed contract.
func NewEarth(address common.Address, backend bind.ContractBackend) (*Earth, error) {
	contract, err := bindEarth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Earth{EarthCaller: EarthCaller{contract: contract}, EarthTransactor: EarthTransactor{contract: contract}, EarthFilterer: EarthFilterer{contract: contract}}, nil
}

// NewEarthCaller creates a new read-only instance of Earth, bound to a specific deployed contract.
func NewEarthCaller(address common.Address, caller bind.ContractCaller) (*EarthCaller, error) {
	contract, err := bindEarth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EarthCaller{contract: contract}, nil
}

// NewEarthTransactor creates a new write-only instance of Earth, bound to a specific deployed contract.
func NewEarthTransactor(address common.Address, transactor bind.ContractTransactor) (*EarthTransactor, error) {
	contract, err := bindEarth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EarthTransactor{contract: contract}, nil
}

// NewEarthFilterer creates a new log filterer instance of Earth, bound to a specific deployed contract.
func NewEarthFilterer(address common.Address, filterer bind.ContractFilterer) (*EarthFilterer, error) {
	contract, err := bindEarth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EarthFilterer{contract: contract}, nil
}

// bindEarth binds a generic wrapper to an already deployed contract.
func bindEarth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EarthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Earth *EarthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Earth.Contract.EarthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Earth *EarthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Earth.Contract.EarthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Earth *EarthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Earth.Contract.EarthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Earth *EarthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Earth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Earth *EarthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Earth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Earth *EarthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Earth.Contract.contract.Transact(opts, method, params...)
}

// AckPackets is a free data retrieval call binding the contract method 0x4252ae9b.
//
// Solidity: function ackPackets(uint256 ) view returns(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack)
func (_Earth *EarthCaller) AckPackets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChannelId [32]byte
	Packet    UniversalPacket
	Ack       AckPacket
}, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "ackPackets", arg0)

	outstruct := new(struct {
		ChannelId [32]byte
		Packet    UniversalPacket
		Ack       AckPacket
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChannelId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Packet = *abi.ConvertType(out[1], new(UniversalPacket)).(*UniversalPacket)
	outstruct.Ack = *abi.ConvertType(out[2], new(AckPacket)).(*AckPacket)

	return *outstruct, err

}

// AckPackets is a free data retrieval call binding the contract method 0x4252ae9b.
//
// Solidity: function ackPackets(uint256 ) view returns(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack)
func (_Earth *EarthSession) AckPackets(arg0 *big.Int) (struct {
	ChannelId [32]byte
	Packet    UniversalPacket
	Ack       AckPacket
}, error) {
	return _Earth.Contract.AckPackets(&_Earth.CallOpts, arg0)
}

// AckPackets is a free data retrieval call binding the contract method 0x4252ae9b.
//
// Solidity: function ackPackets(uint256 ) view returns(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack)
func (_Earth *EarthCallerSession) AckPackets(arg0 *big.Int) (struct {
	ChannelId [32]byte
	Packet    UniversalPacket
	Ack       AckPacket
}, error) {
	return _Earth.Contract.AckPackets(&_Earth.CallOpts, arg0)
}

// AuthorizedChannelIds is a free data retrieval call binding the contract method 0x2eed7c70.
//
// Solidity: function authorizedChannelIds(bytes32 ) view returns(bool)
func (_Earth *EarthCaller) AuthorizedChannelIds(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "authorizedChannelIds", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedChannelIds is a free data retrieval call binding the contract method 0x2eed7c70.
//
// Solidity: function authorizedChannelIds(bytes32 ) view returns(bool)
func (_Earth *EarthSession) AuthorizedChannelIds(arg0 [32]byte) (bool, error) {
	return _Earth.Contract.AuthorizedChannelIds(&_Earth.CallOpts, arg0)
}

// AuthorizedChannelIds is a free data retrieval call binding the contract method 0x2eed7c70.
//
// Solidity: function authorizedChannelIds(bytes32 ) view returns(bool)
func (_Earth *EarthCallerSession) AuthorizedChannelIds(arg0 [32]byte) (bool, error) {
	return _Earth.Contract.AuthorizedChannelIds(&_Earth.CallOpts, arg0)
}

// AuthorizedMws is a free data retrieval call binding the contract method 0x3a7fbcbd.
//
// Solidity: function authorizedMws(address ) view returns(bool)
func (_Earth *EarthCaller) AuthorizedMws(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "authorizedMws", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedMws is a free data retrieval call binding the contract method 0x3a7fbcbd.
//
// Solidity: function authorizedMws(address ) view returns(bool)
func (_Earth *EarthSession) AuthorizedMws(arg0 common.Address) (bool, error) {
	return _Earth.Contract.AuthorizedMws(&_Earth.CallOpts, arg0)
}

// AuthorizedMws is a free data retrieval call binding the contract method 0x3a7fbcbd.
//
// Solidity: function authorizedMws(address ) view returns(bool)
func (_Earth *EarthCallerSession) AuthorizedMws(arg0 common.Address) (bool, error) {
	return _Earth.Contract.AuthorizedMws(&_Earth.CallOpts, arg0)
}

// GenerateAckPacket is a free data retrieval call binding the contract method 0x866f3f97.
//
// Solidity: function generateAckPacket(bytes32 , address srcPortAddr, bytes appData) view returns((bool,bytes) ackPacket)
func (_Earth *EarthCaller) GenerateAckPacket(opts *bind.CallOpts, arg0 [32]byte, srcPortAddr common.Address, appData []byte) (AckPacket, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "generateAckPacket", arg0, srcPortAddr, appData)

	if err != nil {
		return *new(AckPacket), err
	}

	out0 := *abi.ConvertType(out[0], new(AckPacket)).(*AckPacket)

	return out0, err

}

// GenerateAckPacket is a free data retrieval call binding the contract method 0x866f3f97.
//
// Solidity: function generateAckPacket(bytes32 , address srcPortAddr, bytes appData) view returns((bool,bytes) ackPacket)
func (_Earth *EarthSession) GenerateAckPacket(arg0 [32]byte, srcPortAddr common.Address, appData []byte) (AckPacket, error) {
	return _Earth.Contract.GenerateAckPacket(&_Earth.CallOpts, arg0, srcPortAddr, appData)
}

// GenerateAckPacket is a free data retrieval call binding the contract method 0x866f3f97.
//
// Solidity: function generateAckPacket(bytes32 , address srcPortAddr, bytes appData) view returns((bool,bytes) ackPacket)
func (_Earth *EarthCallerSession) GenerateAckPacket(arg0 [32]byte, srcPortAddr common.Address, appData []byte) (AckPacket, error) {
	return _Earth.Contract.GenerateAckPacket(&_Earth.CallOpts, arg0, srcPortAddr, appData)
}

// Mw is a free data retrieval call binding the contract method 0xa742d78c.
//
// Solidity: function mw() view returns(address)
func (_Earth *EarthCaller) Mw(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "mw")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mw is a free data retrieval call binding the contract method 0xa742d78c.
//
// Solidity: function mw() view returns(address)
func (_Earth *EarthSession) Mw() (common.Address, error) {
	return _Earth.Contract.Mw(&_Earth.CallOpts)
}

// Mw is a free data retrieval call binding the contract method 0xa742d78c.
//
// Solidity: function mw() view returns(address)
func (_Earth *EarthCallerSession) Mw() (common.Address, error) {
	return _Earth.Contract.Mw(&_Earth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Earth *EarthCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Earth *EarthSession) Owner() (common.Address, error) {
	return _Earth.Contract.Owner(&_Earth.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Earth *EarthCallerSession) Owner() (common.Address, error) {
	return _Earth.Contract.Owner(&_Earth.CallOpts)
}

// RecvedPackets is a free data retrieval call binding the contract method 0xf12b758a.
//
// Solidity: function recvedPackets(uint256 ) view returns(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet)
func (_Earth *EarthCaller) RecvedPackets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChannelId [32]byte
	Packet    UniversalPacket
}, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "recvedPackets", arg0)

	outstruct := new(struct {
		ChannelId [32]byte
		Packet    UniversalPacket
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChannelId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Packet = *abi.ConvertType(out[1], new(UniversalPacket)).(*UniversalPacket)

	return *outstruct, err

}

// RecvedPackets is a free data retrieval call binding the contract method 0xf12b758a.
//
// Solidity: function recvedPackets(uint256 ) view returns(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet)
func (_Earth *EarthSession) RecvedPackets(arg0 *big.Int) (struct {
	ChannelId [32]byte
	Packet    UniversalPacket
}, error) {
	return _Earth.Contract.RecvedPackets(&_Earth.CallOpts, arg0)
}

// RecvedPackets is a free data retrieval call binding the contract method 0xf12b758a.
//
// Solidity: function recvedPackets(uint256 ) view returns(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet)
func (_Earth *EarthCallerSession) RecvedPackets(arg0 *big.Int) (struct {
	ChannelId [32]byte
	Packet    UniversalPacket
}, error) {
	return _Earth.Contract.RecvedPackets(&_Earth.CallOpts, arg0)
}

// TimeoutPackets is a free data retrieval call binding the contract method 0x4eeb7391.
//
// Solidity: function timeoutPackets(uint256 ) view returns(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet)
func (_Earth *EarthCaller) TimeoutPackets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ChannelId [32]byte
	Packet    UniversalPacket
}, error) {
	var out []interface{}
	err := _Earth.contract.Call(opts, &out, "timeoutPackets", arg0)

	outstruct := new(struct {
		ChannelId [32]byte
		Packet    UniversalPacket
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChannelId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Packet = *abi.ConvertType(out[1], new(UniversalPacket)).(*UniversalPacket)

	return *outstruct, err

}

// TimeoutPackets is a free data retrieval call binding the contract method 0x4eeb7391.
//
// Solidity: function timeoutPackets(uint256 ) view returns(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet)
func (_Earth *EarthSession) TimeoutPackets(arg0 *big.Int) (struct {
	ChannelId [32]byte
	Packet    UniversalPacket
}, error) {
	return _Earth.Contract.TimeoutPackets(&_Earth.CallOpts, arg0)
}

// TimeoutPackets is a free data retrieval call binding the contract method 0x4eeb7391.
//
// Solidity: function timeoutPackets(uint256 ) view returns(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet)
func (_Earth *EarthCallerSession) TimeoutPackets(arg0 *big.Int) (struct {
	ChannelId [32]byte
	Packet    UniversalPacket
}, error) {
	return _Earth.Contract.TimeoutPackets(&_Earth.CallOpts, arg0)
}

// AuthorizeChannel is a paid mutator transaction binding the contract method 0x92dfa392.
//
// Solidity: function authorizeChannel(bytes32 channelId) returns()
func (_Earth *EarthTransactor) AuthorizeChannel(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "authorizeChannel", channelId)
}

// AuthorizeChannel is a paid mutator transaction binding the contract method 0x92dfa392.
//
// Solidity: function authorizeChannel(bytes32 channelId) returns()
func (_Earth *EarthSession) AuthorizeChannel(channelId [32]byte) (*types.Transaction, error) {
	return _Earth.Contract.AuthorizeChannel(&_Earth.TransactOpts, channelId)
}

// AuthorizeChannel is a paid mutator transaction binding the contract method 0x92dfa392.
//
// Solidity: function authorizeChannel(bytes32 channelId) returns()
func (_Earth *EarthTransactorSession) AuthorizeChannel(channelId [32]byte) (*types.Transaction, error) {
	return _Earth.Contract.AuthorizeChannel(&_Earth.TransactOpts, channelId)
}

// AuthorizeMiddleware is a paid mutator transaction binding the contract method 0x3b90b042.
//
// Solidity: function authorizeMiddleware(address middleware) returns()
func (_Earth *EarthTransactor) AuthorizeMiddleware(opts *bind.TransactOpts, middleware common.Address) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "authorizeMiddleware", middleware)
}

// AuthorizeMiddleware is a paid mutator transaction binding the contract method 0x3b90b042.
//
// Solidity: function authorizeMiddleware(address middleware) returns()
func (_Earth *EarthSession) AuthorizeMiddleware(middleware common.Address) (*types.Transaction, error) {
	return _Earth.Contract.AuthorizeMiddleware(&_Earth.TransactOpts, middleware)
}

// AuthorizeMiddleware is a paid mutator transaction binding the contract method 0x3b90b042.
//
// Solidity: function authorizeMiddleware(address middleware) returns()
func (_Earth *EarthTransactorSession) AuthorizeMiddleware(middleware common.Address) (*types.Transaction, error) {
	return _Earth.Contract.AuthorizeMiddleware(&_Earth.TransactOpts, middleware)
}

// Greet is a paid mutator transaction binding the contract method 0xd24ba024.
//
// Solidity: function greet(address destPortAddr, bytes32 channelId, bytes message, uint64 timeoutTimestamp) returns()
func (_Earth *EarthTransactor) Greet(opts *bind.TransactOpts, destPortAddr common.Address, channelId [32]byte, message []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "greet", destPortAddr, channelId, message, timeoutTimestamp)
}

// Greet is a paid mutator transaction binding the contract method 0xd24ba024.
//
// Solidity: function greet(address destPortAddr, bytes32 channelId, bytes message, uint64 timeoutTimestamp) returns()
func (_Earth *EarthSession) Greet(destPortAddr common.Address, channelId [32]byte, message []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Earth.Contract.Greet(&_Earth.TransactOpts, destPortAddr, channelId, message, timeoutTimestamp)
}

// Greet is a paid mutator transaction binding the contract method 0xd24ba024.
//
// Solidity: function greet(address destPortAddr, bytes32 channelId, bytes message, uint64 timeoutTimestamp) returns()
func (_Earth *EarthTransactorSession) Greet(destPortAddr common.Address, channelId [32]byte, message []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Earth.Contract.Greet(&_Earth.TransactOpts, destPortAddr, channelId, message, timeoutTimestamp)
}

// GreetWithFee is a paid mutator transaction binding the contract method 0x2466911c.
//
// Solidity: function greetWithFee(address destPortAddr, bytes32 channelId, bytes message, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_Earth *EarthTransactor) GreetWithFee(opts *bind.TransactOpts, destPortAddr common.Address, channelId [32]byte, message []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "greetWithFee", destPortAddr, channelId, message, timeoutTimestamp, gasLimits, gasPrices)
}

// GreetWithFee is a paid mutator transaction binding the contract method 0x2466911c.
//
// Solidity: function greetWithFee(address destPortAddr, bytes32 channelId, bytes message, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_Earth *EarthSession) GreetWithFee(destPortAddr common.Address, channelId [32]byte, message []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _Earth.Contract.GreetWithFee(&_Earth.TransactOpts, destPortAddr, channelId, message, timeoutTimestamp, gasLimits, gasPrices)
}

// GreetWithFee is a paid mutator transaction binding the contract method 0x2466911c.
//
// Solidity: function greetWithFee(address destPortAddr, bytes32 channelId, bytes message, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_Earth *EarthTransactorSession) GreetWithFee(destPortAddr common.Address, channelId [32]byte, message []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _Earth.Contract.GreetWithFee(&_Earth.TransactOpts, destPortAddr, channelId, message, timeoutTimestamp, gasLimits, gasPrices)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns((bool,bytes) ackPacket)
func (_Earth *EarthTransactor) OnRecvUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "onRecvUniversalPacket", channelId, packet)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns((bool,bytes) ackPacket)
func (_Earth *EarthSession) OnRecvUniversalPacket(channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _Earth.Contract.OnRecvUniversalPacket(&_Earth.TransactOpts, channelId, packet)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns((bool,bytes) ackPacket)
func (_Earth *EarthTransactorSession) OnRecvUniversalPacket(channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _Earth.Contract.OnRecvUniversalPacket(&_Earth.TransactOpts, channelId, packet)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_Earth *EarthTransactor) OnTimeoutUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "onTimeoutUniversalPacket", channelId, packet)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_Earth *EarthSession) OnTimeoutUniversalPacket(channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _Earth.Contract.OnTimeoutUniversalPacket(&_Earth.TransactOpts, channelId, packet)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_Earth *EarthTransactorSession) OnTimeoutUniversalPacket(channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _Earth.Contract.OnTimeoutUniversalPacket(&_Earth.TransactOpts, channelId, packet)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_Earth *EarthTransactor) OnUniversalAcknowledgement(opts *bind.TransactOpts, channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "onUniversalAcknowledgement", channelId, packet, ack)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_Earth *EarthSession) OnUniversalAcknowledgement(channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _Earth.Contract.OnUniversalAcknowledgement(&_Earth.TransactOpts, channelId, packet, ack)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_Earth *EarthTransactorSession) OnUniversalAcknowledgement(channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _Earth.Contract.OnUniversalAcknowledgement(&_Earth.TransactOpts, channelId, packet, ack)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Earth *EarthTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Earth *EarthSession) RenounceOwnership() (*types.Transaction, error) {
	return _Earth.Contract.RenounceOwnership(&_Earth.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Earth *EarthTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Earth.Contract.RenounceOwnership(&_Earth.TransactOpts)
}

// SetDefaultMw is a paid mutator transaction binding the contract method 0x00e82cef.
//
// Solidity: function setDefaultMw(address _middleware) returns()
func (_Earth *EarthTransactor) SetDefaultMw(opts *bind.TransactOpts, _middleware common.Address) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "setDefaultMw", _middleware)
}

// SetDefaultMw is a paid mutator transaction binding the contract method 0x00e82cef.
//
// Solidity: function setDefaultMw(address _middleware) returns()
func (_Earth *EarthSession) SetDefaultMw(_middleware common.Address) (*types.Transaction, error) {
	return _Earth.Contract.SetDefaultMw(&_Earth.TransactOpts, _middleware)
}

// SetDefaultMw is a paid mutator transaction binding the contract method 0x00e82cef.
//
// Solidity: function setDefaultMw(address _middleware) returns()
func (_Earth *EarthTransactorSession) SetDefaultMw(_middleware common.Address) (*types.Transaction, error) {
	return _Earth.Contract.SetDefaultMw(&_Earth.TransactOpts, _middleware)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Earth *EarthTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Earth.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Earth *EarthSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Earth.Contract.TransferOwnership(&_Earth.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Earth *EarthTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Earth.Contract.TransferOwnership(&_Earth.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Earth *EarthTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Earth.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Earth *EarthSession) Receive() (*types.Transaction, error) {
	return _Earth.Contract.Receive(&_Earth.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Earth *EarthTransactorSession) Receive() (*types.Transaction, error) {
	return _Earth.Contract.Receive(&_Earth.TransactOpts)
}

// EarthOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Earth contract.
type EarthOwnershipTransferredIterator struct {
	Event *EarthOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EarthOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EarthOwnershipTransferred)
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
		it.Event = new(EarthOwnershipTransferred)
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
func (it *EarthOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EarthOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EarthOwnershipTransferred represents a OwnershipTransferred event raised by the Earth contract.
type EarthOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Earth *EarthFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EarthOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Earth.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EarthOwnershipTransferredIterator{contract: _Earth.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Earth *EarthFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EarthOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Earth.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EarthOwnershipTransferred)
				if err := _Earth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Earth *EarthFilterer) ParseOwnershipTransferred(log types.Log) (*EarthOwnershipTransferred, error) {
	event := new(EarthOwnershipTransferred)
	if err := _Earth.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
