// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package moon

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

// IbcEndpoint is an auto generated low-level Go binding around an user-defined struct.
type IbcEndpoint struct {
	PortId    string
	ChannelId [32]byte
}

// IbcPacket is an auto generated low-level Go binding around an user-defined struct.
type IbcPacket struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}

// MoonMetaData contains all meta data concerning the Moon contract.
var MoonMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_dispatcher\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ackPackets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"connectedChannels\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dispatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"greet\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"greetWithFee\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"onAcknowledgementPacket\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseInit\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenInit\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onChanOpenTry\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recvedPackets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportedVersions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"timeoutPackets\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerChannelClose\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerChannelInit\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"triggerChannelInitWithFee\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ChannelNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedVersion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"notIbcDispatcher\",\"inputs\":[]}]",
}

// MoonABI is the input ABI used to generate the binding from.
// Deprecated: Use MoonMetaData.ABI instead.
var MoonABI = MoonMetaData.ABI

// Moon is an auto generated Go binding around an Ethereum contract.
type Moon struct {
	MoonCaller     // Read-only binding to the contract
	MoonTransactor // Write-only binding to the contract
	MoonFilterer   // Log filterer for contract events
}

// MoonCaller is an auto generated read-only Go binding around an Ethereum contract.
type MoonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MoonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MoonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MoonSession struct {
	Contract     *Moon             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MoonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MoonCallerSession struct {
	Contract *MoonCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MoonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MoonTransactorSession struct {
	Contract     *MoonTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MoonRaw is an auto generated low-level Go binding around an Ethereum contract.
type MoonRaw struct {
	Contract *Moon // Generic contract binding to access the raw methods on
}

// MoonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MoonCallerRaw struct {
	Contract *MoonCaller // Generic read-only contract binding to access the raw methods on
}

// MoonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MoonTransactorRaw struct {
	Contract *MoonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMoon creates a new instance of Moon, bound to a specific deployed contract.
func NewMoon(address common.Address, backend bind.ContractBackend) (*Moon, error) {
	contract, err := bindMoon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Moon{MoonCaller: MoonCaller{contract: contract}, MoonTransactor: MoonTransactor{contract: contract}, MoonFilterer: MoonFilterer{contract: contract}}, nil
}

// NewMoonCaller creates a new read-only instance of Moon, bound to a specific deployed contract.
func NewMoonCaller(address common.Address, caller bind.ContractCaller) (*MoonCaller, error) {
	contract, err := bindMoon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MoonCaller{contract: contract}, nil
}

// NewMoonTransactor creates a new write-only instance of Moon, bound to a specific deployed contract.
func NewMoonTransactor(address common.Address, transactor bind.ContractTransactor) (*MoonTransactor, error) {
	contract, err := bindMoon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MoonTransactor{contract: contract}, nil
}

// NewMoonFilterer creates a new log filterer instance of Moon, bound to a specific deployed contract.
func NewMoonFilterer(address common.Address, filterer bind.ContractFilterer) (*MoonFilterer, error) {
	contract, err := bindMoon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MoonFilterer{contract: contract}, nil
}

// bindMoon binds a generic wrapper to an already deployed contract.
func bindMoon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MoonMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Moon *MoonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Moon.Contract.MoonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Moon *MoonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Moon.Contract.MoonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Moon *MoonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Moon.Contract.MoonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Moon *MoonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Moon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Moon *MoonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Moon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Moon *MoonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Moon.Contract.contract.Transact(opts, method, params...)
}

// AckPackets is a free data retrieval call binding the contract method 0x4252ae9b.
//
// Solidity: function ackPackets(uint256 ) view returns(bool success, bytes data)
func (_Moon *MoonCaller) AckPackets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Success bool
	Data    []byte
}, error) {
	var out []interface{}
	err := _Moon.contract.Call(opts, &out, "ackPackets", arg0)

	outstruct := new(struct {
		Success bool
		Data    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Success = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Data = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// AckPackets is a free data retrieval call binding the contract method 0x4252ae9b.
//
// Solidity: function ackPackets(uint256 ) view returns(bool success, bytes data)
func (_Moon *MoonSession) AckPackets(arg0 *big.Int) (struct {
	Success bool
	Data    []byte
}, error) {
	return _Moon.Contract.AckPackets(&_Moon.CallOpts, arg0)
}

// AckPackets is a free data retrieval call binding the contract method 0x4252ae9b.
//
// Solidity: function ackPackets(uint256 ) view returns(bool success, bytes data)
func (_Moon *MoonCallerSession) AckPackets(arg0 *big.Int) (struct {
	Success bool
	Data    []byte
}, error) {
	return _Moon.Contract.AckPackets(&_Moon.CallOpts, arg0)
}

// ConnectedChannels is a free data retrieval call binding the contract method 0xbb3f9f8d.
//
// Solidity: function connectedChannels(uint256 ) view returns(bytes32)
func (_Moon *MoonCaller) ConnectedChannels(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Moon.contract.Call(opts, &out, "connectedChannels", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConnectedChannels is a free data retrieval call binding the contract method 0xbb3f9f8d.
//
// Solidity: function connectedChannels(uint256 ) view returns(bytes32)
func (_Moon *MoonSession) ConnectedChannels(arg0 *big.Int) ([32]byte, error) {
	return _Moon.Contract.ConnectedChannels(&_Moon.CallOpts, arg0)
}

// ConnectedChannels is a free data retrieval call binding the contract method 0xbb3f9f8d.
//
// Solidity: function connectedChannels(uint256 ) view returns(bytes32)
func (_Moon *MoonCallerSession) ConnectedChannels(arg0 *big.Int) ([32]byte, error) {
	return _Moon.Contract.ConnectedChannels(&_Moon.CallOpts, arg0)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_Moon *MoonCaller) Dispatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Moon.contract.Call(opts, &out, "dispatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_Moon *MoonSession) Dispatcher() (common.Address, error) {
	return _Moon.Contract.Dispatcher(&_Moon.CallOpts)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_Moon *MoonCallerSession) Dispatcher() (common.Address, error) {
	return _Moon.Contract.Dispatcher(&_Moon.CallOpts)
}

// OnChanOpenInit is a free data retrieval call binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 , string[] , string , string version) view returns(string selectedVersion)
func (_Moon *MoonCaller) OnChanOpenInit(opts *bind.CallOpts, arg0 uint8, arg1 []string, arg2 string, version string) (string, error) {
	var out []interface{}
	err := _Moon.contract.Call(opts, &out, "onChanOpenInit", arg0, arg1, arg2, version)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OnChanOpenInit is a free data retrieval call binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 , string[] , string , string version) view returns(string selectedVersion)
func (_Moon *MoonSession) OnChanOpenInit(arg0 uint8, arg1 []string, arg2 string, version string) (string, error) {
	return _Moon.Contract.OnChanOpenInit(&_Moon.CallOpts, arg0, arg1, arg2, version)
}

// OnChanOpenInit is a free data retrieval call binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 , string[] , string , string version) view returns(string selectedVersion)
func (_Moon *MoonCallerSession) OnChanOpenInit(arg0 uint8, arg1 []string, arg2 string, version string) (string, error) {
	return _Moon.Contract.OnChanOpenInit(&_Moon.CallOpts, arg0, arg1, arg2, version)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Moon *MoonCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Moon.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Moon *MoonSession) Owner() (common.Address, error) {
	return _Moon.Contract.Owner(&_Moon.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Moon *MoonCallerSession) Owner() (common.Address, error) {
	return _Moon.Contract.Owner(&_Moon.CallOpts)
}

// RecvedPackets is a free data retrieval call binding the contract method 0xf12b758a.
//
// Solidity: function recvedPackets(uint256 ) view returns((string,bytes32) src, (string,bytes32) dest, uint64 sequence, bytes data, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_Moon *MoonCaller) RecvedPackets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}, error) {
	var out []interface{}
	err := _Moon.contract.Call(opts, &out, "recvedPackets", arg0)

	outstruct := new(struct {
		Src              IbcEndpoint
		Dest             IbcEndpoint
		Sequence         uint64
		Data             []byte
		TimeoutHeight    Height
		TimeoutTimestamp uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Src = *abi.ConvertType(out[0], new(IbcEndpoint)).(*IbcEndpoint)
	outstruct.Dest = *abi.ConvertType(out[1], new(IbcEndpoint)).(*IbcEndpoint)
	outstruct.Sequence = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Data = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.TimeoutHeight = *abi.ConvertType(out[4], new(Height)).(*Height)
	outstruct.TimeoutTimestamp = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// RecvedPackets is a free data retrieval call binding the contract method 0xf12b758a.
//
// Solidity: function recvedPackets(uint256 ) view returns((string,bytes32) src, (string,bytes32) dest, uint64 sequence, bytes data, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_Moon *MoonSession) RecvedPackets(arg0 *big.Int) (struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}, error) {
	return _Moon.Contract.RecvedPackets(&_Moon.CallOpts, arg0)
}

// RecvedPackets is a free data retrieval call binding the contract method 0xf12b758a.
//
// Solidity: function recvedPackets(uint256 ) view returns((string,bytes32) src, (string,bytes32) dest, uint64 sequence, bytes data, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_Moon *MoonCallerSession) RecvedPackets(arg0 *big.Int) (struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}, error) {
	return _Moon.Contract.RecvedPackets(&_Moon.CallOpts, arg0)
}

// SupportedVersions is a free data retrieval call binding the contract method 0x7d622184.
//
// Solidity: function supportedVersions(uint256 ) view returns(string)
func (_Moon *MoonCaller) SupportedVersions(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Moon.contract.Call(opts, &out, "supportedVersions", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SupportedVersions is a free data retrieval call binding the contract method 0x7d622184.
//
// Solidity: function supportedVersions(uint256 ) view returns(string)
func (_Moon *MoonSession) SupportedVersions(arg0 *big.Int) (string, error) {
	return _Moon.Contract.SupportedVersions(&_Moon.CallOpts, arg0)
}

// SupportedVersions is a free data retrieval call binding the contract method 0x7d622184.
//
// Solidity: function supportedVersions(uint256 ) view returns(string)
func (_Moon *MoonCallerSession) SupportedVersions(arg0 *big.Int) (string, error) {
	return _Moon.Contract.SupportedVersions(&_Moon.CallOpts, arg0)
}

// TimeoutPackets is a free data retrieval call binding the contract method 0x4eeb7391.
//
// Solidity: function timeoutPackets(uint256 ) view returns((string,bytes32) src, (string,bytes32) dest, uint64 sequence, bytes data, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_Moon *MoonCaller) TimeoutPackets(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}, error) {
	var out []interface{}
	err := _Moon.contract.Call(opts, &out, "timeoutPackets", arg0)

	outstruct := new(struct {
		Src              IbcEndpoint
		Dest             IbcEndpoint
		Sequence         uint64
		Data             []byte
		TimeoutHeight    Height
		TimeoutTimestamp uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Src = *abi.ConvertType(out[0], new(IbcEndpoint)).(*IbcEndpoint)
	outstruct.Dest = *abi.ConvertType(out[1], new(IbcEndpoint)).(*IbcEndpoint)
	outstruct.Sequence = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Data = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.TimeoutHeight = *abi.ConvertType(out[4], new(Height)).(*Height)
	outstruct.TimeoutTimestamp = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// TimeoutPackets is a free data retrieval call binding the contract method 0x4eeb7391.
//
// Solidity: function timeoutPackets(uint256 ) view returns((string,bytes32) src, (string,bytes32) dest, uint64 sequence, bytes data, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_Moon *MoonSession) TimeoutPackets(arg0 *big.Int) (struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}, error) {
	return _Moon.Contract.TimeoutPackets(&_Moon.CallOpts, arg0)
}

// TimeoutPackets is a free data retrieval call binding the contract method 0x4eeb7391.
//
// Solidity: function timeoutPackets(uint256 ) view returns((string,bytes32) src, (string,bytes32) dest, uint64 sequence, bytes data, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_Moon *MoonCallerSession) TimeoutPackets(arg0 *big.Int) (struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}, error) {
	return _Moon.Contract.TimeoutPackets(&_Moon.CallOpts, arg0)
}

// Greet is a paid mutator transaction binding the contract method 0x5bfd12b8.
//
// Solidity: function greet(string message, bytes32 channelId, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_Moon *MoonTransactor) Greet(opts *bind.TransactOpts, message string, channelId [32]byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "greet", message, channelId, timeoutTimestamp)
}

// Greet is a paid mutator transaction binding the contract method 0x5bfd12b8.
//
// Solidity: function greet(string message, bytes32 channelId, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_Moon *MoonSession) Greet(message string, channelId [32]byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Moon.Contract.Greet(&_Moon.TransactOpts, message, channelId, timeoutTimestamp)
}

// Greet is a paid mutator transaction binding the contract method 0x5bfd12b8.
//
// Solidity: function greet(string message, bytes32 channelId, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_Moon *MoonTransactorSession) Greet(message string, channelId [32]byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Moon.Contract.Greet(&_Moon.TransactOpts, message, channelId, timeoutTimestamp)
}

// GreetWithFee is a paid mutator transaction binding the contract method 0x3513a995.
//
// Solidity: function greetWithFee(string message, bytes32 channelId, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_Moon *MoonTransactor) GreetWithFee(opts *bind.TransactOpts, message string, channelId [32]byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "greetWithFee", message, channelId, timeoutTimestamp, gasLimits, gasPrices)
}

// GreetWithFee is a paid mutator transaction binding the contract method 0x3513a995.
//
// Solidity: function greetWithFee(string message, bytes32 channelId, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_Moon *MoonSession) GreetWithFee(message string, channelId [32]byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _Moon.Contract.GreetWithFee(&_Moon.TransactOpts, message, channelId, timeoutTimestamp, gasLimits, gasPrices)
}

// GreetWithFee is a paid mutator transaction binding the contract method 0x3513a995.
//
// Solidity: function greetWithFee(string message, bytes32 channelId, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_Moon *MoonTransactorSession) GreetWithFee(message string, channelId [32]byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _Moon.Contract.GreetWithFee(&_Moon.TransactOpts, message, channelId, timeoutTimestamp, gasLimits, gasPrices)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) , (bool,bytes) ack) returns()
func (_Moon *MoonTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, arg0 IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "onAcknowledgementPacket", arg0, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) , (bool,bytes) ack) returns()
func (_Moon *MoonSession) OnAcknowledgementPacket(arg0 IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _Moon.Contract.OnAcknowledgementPacket(&_Moon.TransactOpts, arg0, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) , (bool,bytes) ack) returns()
func (_Moon *MoonTransactorSession) OnAcknowledgementPacket(arg0 IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _Moon.Contract.OnAcknowledgementPacket(&_Moon.TransactOpts, arg0, ack)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string , bytes32 ) returns()
func (_Moon *MoonTransactor) OnChanCloseConfirm(opts *bind.TransactOpts, channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "onChanCloseConfirm", channelId, arg1, arg2)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string , bytes32 ) returns()
func (_Moon *MoonSession) OnChanCloseConfirm(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _Moon.Contract.OnChanCloseConfirm(&_Moon.TransactOpts, channelId, arg1, arg2)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string , bytes32 ) returns()
func (_Moon *MoonTransactorSession) OnChanCloseConfirm(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _Moon.Contract.OnChanCloseConfirm(&_Moon.TransactOpts, channelId, arg1, arg2)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string , bytes32 ) returns()
func (_Moon *MoonTransactor) OnChanCloseInit(opts *bind.TransactOpts, channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "onChanCloseInit", channelId, arg1, arg2)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string , bytes32 ) returns()
func (_Moon *MoonSession) OnChanCloseInit(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _Moon.Contract.OnChanCloseInit(&_Moon.TransactOpts, channelId, arg1, arg2)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string , bytes32 ) returns()
func (_Moon *MoonTransactorSession) OnChanCloseInit(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _Moon.Contract.OnChanCloseInit(&_Moon.TransactOpts, channelId, arg1, arg2)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) returns()
func (_Moon *MoonTransactor) OnChanOpenAck(opts *bind.TransactOpts, channelId [32]byte, arg1 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "onChanOpenAck", channelId, arg1, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) returns()
func (_Moon *MoonSession) OnChanOpenAck(channelId [32]byte, arg1 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _Moon.Contract.OnChanOpenAck(&_Moon.TransactOpts, channelId, arg1, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) returns()
func (_Moon *MoonTransactorSession) OnChanOpenAck(channelId [32]byte, arg1 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _Moon.Contract.OnChanOpenAck(&_Moon.TransactOpts, channelId, arg1, counterpartyVersion)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_Moon *MoonTransactor) OnChanOpenConfirm(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "onChanOpenConfirm", channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_Moon *MoonSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _Moon.Contract.OnChanOpenConfirm(&_Moon.TransactOpts, channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_Moon *MoonTransactorSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _Moon.Contract.OnChanOpenConfirm(&_Moon.TransactOpts, channelId)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 , string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) returns(string selectedVersion)
func (_Moon *MoonTransactor) OnChanOpenTry(opts *bind.TransactOpts, arg0 uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "onChanOpenTry", arg0, arg1, channelId, arg3, arg4, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 , string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) returns(string selectedVersion)
func (_Moon *MoonSession) OnChanOpenTry(arg0 uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _Moon.Contract.OnChanOpenTry(&_Moon.TransactOpts, arg0, arg1, channelId, arg3, arg4, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 , string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) returns(string selectedVersion)
func (_Moon *MoonTransactorSession) OnChanOpenTry(arg0 uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _Moon.Contract.OnChanOpenTry(&_Moon.TransactOpts, arg0, arg1, channelId, arg3, arg4, counterpartyVersion)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_Moon *MoonTransactor) OnRecvPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "onRecvPacket", packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_Moon *MoonSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _Moon.Contract.OnRecvPacket(&_Moon.TransactOpts, packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_Moon *MoonTransactorSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _Moon.Contract.OnRecvPacket(&_Moon.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_Moon *MoonTransactor) OnTimeoutPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "onTimeoutPacket", packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_Moon *MoonSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _Moon.Contract.OnTimeoutPacket(&_Moon.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_Moon *MoonTransactorSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _Moon.Contract.OnTimeoutPacket(&_Moon.TransactOpts, packet)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Moon *MoonTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Moon *MoonSession) RenounceOwnership() (*types.Transaction, error) {
	return _Moon.Contract.RenounceOwnership(&_Moon.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Moon *MoonTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Moon.Contract.RenounceOwnership(&_Moon.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Moon *MoonTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Moon *MoonSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Moon.Contract.TransferOwnership(&_Moon.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Moon *MoonTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Moon.Contract.TransferOwnership(&_Moon.TransactOpts, newOwner)
}

// TriggerChannelClose is a paid mutator transaction binding the contract method 0x558850ac.
//
// Solidity: function triggerChannelClose(bytes32 channelId) returns()
func (_Moon *MoonTransactor) TriggerChannelClose(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "triggerChannelClose", channelId)
}

// TriggerChannelClose is a paid mutator transaction binding the contract method 0x558850ac.
//
// Solidity: function triggerChannelClose(bytes32 channelId) returns()
func (_Moon *MoonSession) TriggerChannelClose(channelId [32]byte) (*types.Transaction, error) {
	return _Moon.Contract.TriggerChannelClose(&_Moon.TransactOpts, channelId)
}

// TriggerChannelClose is a paid mutator transaction binding the contract method 0x558850ac.
//
// Solidity: function triggerChannelClose(bytes32 channelId) returns()
func (_Moon *MoonTransactorSession) TriggerChannelClose(channelId [32]byte) (*types.Transaction, error) {
	return _Moon.Contract.TriggerChannelClose(&_Moon.TransactOpts, channelId)
}

// TriggerChannelInit is a paid mutator transaction binding the contract method 0x7a805598.
//
// Solidity: function triggerChannelInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_Moon *MoonTransactor) TriggerChannelInit(opts *bind.TransactOpts, version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "triggerChannelInit", version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// TriggerChannelInit is a paid mutator transaction binding the contract method 0x7a805598.
//
// Solidity: function triggerChannelInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_Moon *MoonSession) TriggerChannelInit(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _Moon.Contract.TriggerChannelInit(&_Moon.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// TriggerChannelInit is a paid mutator transaction binding the contract method 0x7a805598.
//
// Solidity: function triggerChannelInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_Moon *MoonTransactorSession) TriggerChannelInit(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _Moon.Contract.TriggerChannelInit(&_Moon.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// TriggerChannelInitWithFee is a paid mutator transaction binding the contract method 0x61995001.
//
// Solidity: function triggerChannelInitWithFee(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) payable returns()
func (_Moon *MoonTransactor) TriggerChannelInitWithFee(opts *bind.TransactOpts, version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _Moon.contract.Transact(opts, "triggerChannelInitWithFee", version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// TriggerChannelInitWithFee is a paid mutator transaction binding the contract method 0x61995001.
//
// Solidity: function triggerChannelInitWithFee(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) payable returns()
func (_Moon *MoonSession) TriggerChannelInitWithFee(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _Moon.Contract.TriggerChannelInitWithFee(&_Moon.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// TriggerChannelInitWithFee is a paid mutator transaction binding the contract method 0x61995001.
//
// Solidity: function triggerChannelInitWithFee(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) payable returns()
func (_Moon *MoonTransactorSession) TriggerChannelInitWithFee(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _Moon.Contract.TriggerChannelInitWithFee(&_Moon.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Moon *MoonTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Moon.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Moon *MoonSession) Receive() (*types.Transaction, error) {
	return _Moon.Contract.Receive(&_Moon.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Moon *MoonTransactorSession) Receive() (*types.Transaction, error) {
	return _Moon.Contract.Receive(&_Moon.TransactOpts)
}

// MoonOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Moon contract.
type MoonOwnershipTransferredIterator struct {
	Event *MoonOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MoonOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonOwnershipTransferred)
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
		it.Event = new(MoonOwnershipTransferred)
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
func (it *MoonOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonOwnershipTransferred represents a OwnershipTransferred event raised by the Moon contract.
type MoonOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Moon *MoonFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MoonOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Moon.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MoonOwnershipTransferredIterator{contract: _Moon.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Moon *MoonFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MoonOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Moon.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonOwnershipTransferred)
				if err := _Moon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Moon *MoonFilterer) ParseOwnershipTransferred(log types.Log) (*MoonOwnershipTransferred, error) {
	event := new(MoonOwnershipTransferred)
	if err := _Moon.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
