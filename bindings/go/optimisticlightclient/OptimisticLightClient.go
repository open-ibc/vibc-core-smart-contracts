// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimisticlightclient

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

// Ics23Proof is an auto generated low-level Go binding around an user-defined struct.
type Ics23Proof struct {
	Proof  []OpIcs23Proof
	Height *big.Int
}

// L1Header is an auto generated low-level Go binding around an user-defined struct.
type L1Header struct {
	Header    [][]byte
	StateRoot [32]byte
	Number    uint64
}

// OpIcs23Proof is an auto generated low-level Go binding around an user-defined struct.
type OpIcs23Proof struct {
	Path   []OpIcs23ProofPath
	Key    []byte
	Value  []byte
	Prefix []byte
}

// OpIcs23ProofPath is an auto generated low-level Go binding around an user-defined struct.
type OpIcs23ProofPath struct {
	Prefix []byte
	Suffix []byte
}

// OpL2StateProof is an auto generated low-level Go binding around an user-defined struct.
type OpL2StateProof struct {
	AccountProof        [][]byte
	OutputRootProof     [][]byte
	L2OutputProposalKey [32]byte
	L2BlockHash         [32]byte
}

// OptimisticLightClientMetaData contains all meta data concerning the OptimisticLightClient contract.
var OptimisticLightClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"fraudProofWindowSeconds_\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"verifier_\",\"type\":\"address\",\"internalType\":\"contractIProofVerifier\"},{\"name\":\"_l1BlockProvider\",\"type\":\"address\",\"internalType\":\"contractL1Block\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOpConsensusState\",\"inputs\":[{\"name\":\"l1header\",\"type\":\"tuple\",\"internalType\":\"structL1Header\",\"components\":[{\"name\":\"header\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"number\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structOpL2StateProof\",\"components\":[{\"name\":\"accountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"outputRootProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"l2OutputProposalKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"appHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fraudProofEndTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ended\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"consensusStates\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fraudProofEndtime\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fraudProofWindowSeconds\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFraudProofEndtime\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fraudProofEndTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInternalState\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"appHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fraudProofEndTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ended\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getState\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"appHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fraudProofEndTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ended\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1BlockProvider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractL1Block\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIProofVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expectedValue\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"AppHashHasNotPassedFraudProofWindow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUpdatePendingOptimisticConsensusState\",\"inputs\":[]}]",
}

// OptimisticLightClientABI is the input ABI used to generate the binding from.
// Deprecated: Use OptimisticLightClientMetaData.ABI instead.
var OptimisticLightClientABI = OptimisticLightClientMetaData.ABI

// OptimisticLightClient is an auto generated Go binding around an Ethereum contract.
type OptimisticLightClient struct {
	OptimisticLightClientCaller     // Read-only binding to the contract
	OptimisticLightClientTransactor // Write-only binding to the contract
	OptimisticLightClientFilterer   // Log filterer for contract events
}

// OptimisticLightClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptimisticLightClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimisticLightClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptimisticLightClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimisticLightClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptimisticLightClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimisticLightClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptimisticLightClientSession struct {
	Contract     *OptimisticLightClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OptimisticLightClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptimisticLightClientCallerSession struct {
	Contract *OptimisticLightClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// OptimisticLightClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptimisticLightClientTransactorSession struct {
	Contract     *OptimisticLightClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// OptimisticLightClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptimisticLightClientRaw struct {
	Contract *OptimisticLightClient // Generic contract binding to access the raw methods on
}

// OptimisticLightClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptimisticLightClientCallerRaw struct {
	Contract *OptimisticLightClientCaller // Generic read-only contract binding to access the raw methods on
}

// OptimisticLightClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptimisticLightClientTransactorRaw struct {
	Contract *OptimisticLightClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptimisticLightClient creates a new instance of OptimisticLightClient, bound to a specific deployed contract.
func NewOptimisticLightClient(address common.Address, backend bind.ContractBackend) (*OptimisticLightClient, error) {
	contract, err := bindOptimisticLightClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimisticLightClient{OptimisticLightClientCaller: OptimisticLightClientCaller{contract: contract}, OptimisticLightClientTransactor: OptimisticLightClientTransactor{contract: contract}, OptimisticLightClientFilterer: OptimisticLightClientFilterer{contract: contract}}, nil
}

// NewOptimisticLightClientCaller creates a new read-only instance of OptimisticLightClient, bound to a specific deployed contract.
func NewOptimisticLightClientCaller(address common.Address, caller bind.ContractCaller) (*OptimisticLightClientCaller, error) {
	contract, err := bindOptimisticLightClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimisticLightClientCaller{contract: contract}, nil
}

// NewOptimisticLightClientTransactor creates a new write-only instance of OptimisticLightClient, bound to a specific deployed contract.
func NewOptimisticLightClientTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimisticLightClientTransactor, error) {
	contract, err := bindOptimisticLightClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimisticLightClientTransactor{contract: contract}, nil
}

// NewOptimisticLightClientFilterer creates a new log filterer instance of OptimisticLightClient, bound to a specific deployed contract.
func NewOptimisticLightClientFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimisticLightClientFilterer, error) {
	contract, err := bindOptimisticLightClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimisticLightClientFilterer{contract: contract}, nil
}

// bindOptimisticLightClient binds a generic wrapper to an already deployed contract.
func bindOptimisticLightClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimisticLightClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimisticLightClient *OptimisticLightClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimisticLightClient.Contract.OptimisticLightClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimisticLightClient *OptimisticLightClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimisticLightClient.Contract.OptimisticLightClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimisticLightClient *OptimisticLightClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimisticLightClient.Contract.OptimisticLightClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimisticLightClient *OptimisticLightClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimisticLightClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimisticLightClient *OptimisticLightClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimisticLightClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimisticLightClient *OptimisticLightClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimisticLightClient.Contract.contract.Transact(opts, method, params...)
}

// ConsensusStates is a free data retrieval call binding the contract method 0x1b738a22.
//
// Solidity: function consensusStates(uint256 ) view returns(uint256)
func (_OptimisticLightClient *OptimisticLightClientCaller) ConsensusStates(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OptimisticLightClient.contract.Call(opts, &out, "consensusStates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConsensusStates is a free data retrieval call binding the contract method 0x1b738a22.
//
// Solidity: function consensusStates(uint256 ) view returns(uint256)
func (_OptimisticLightClient *OptimisticLightClientSession) ConsensusStates(arg0 *big.Int) (*big.Int, error) {
	return _OptimisticLightClient.Contract.ConsensusStates(&_OptimisticLightClient.CallOpts, arg0)
}

// ConsensusStates is a free data retrieval call binding the contract method 0x1b738a22.
//
// Solidity: function consensusStates(uint256 ) view returns(uint256)
func (_OptimisticLightClient *OptimisticLightClientCallerSession) ConsensusStates(arg0 *big.Int) (*big.Int, error) {
	return _OptimisticLightClient.Contract.ConsensusStates(&_OptimisticLightClient.CallOpts, arg0)
}

// FraudProofEndtime is a free data retrieval call binding the contract method 0x34b80a41.
//
// Solidity: function fraudProofEndtime(uint256 ) view returns(uint256)
func (_OptimisticLightClient *OptimisticLightClientCaller) FraudProofEndtime(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OptimisticLightClient.contract.Call(opts, &out, "fraudProofEndtime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FraudProofEndtime is a free data retrieval call binding the contract method 0x34b80a41.
//
// Solidity: function fraudProofEndtime(uint256 ) view returns(uint256)
func (_OptimisticLightClient *OptimisticLightClientSession) FraudProofEndtime(arg0 *big.Int) (*big.Int, error) {
	return _OptimisticLightClient.Contract.FraudProofEndtime(&_OptimisticLightClient.CallOpts, arg0)
}

// FraudProofEndtime is a free data retrieval call binding the contract method 0x34b80a41.
//
// Solidity: function fraudProofEndtime(uint256 ) view returns(uint256)
func (_OptimisticLightClient *OptimisticLightClientCallerSession) FraudProofEndtime(arg0 *big.Int) (*big.Int, error) {
	return _OptimisticLightClient.Contract.FraudProofEndtime(&_OptimisticLightClient.CallOpts, arg0)
}

// FraudProofWindowSeconds is a free data retrieval call binding the contract method 0x63042720.
//
// Solidity: function fraudProofWindowSeconds() view returns(uint256)
func (_OptimisticLightClient *OptimisticLightClientCaller) FraudProofWindowSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OptimisticLightClient.contract.Call(opts, &out, "fraudProofWindowSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FraudProofWindowSeconds is a free data retrieval call binding the contract method 0x63042720.
//
// Solidity: function fraudProofWindowSeconds() view returns(uint256)
func (_OptimisticLightClient *OptimisticLightClientSession) FraudProofWindowSeconds() (*big.Int, error) {
	return _OptimisticLightClient.Contract.FraudProofWindowSeconds(&_OptimisticLightClient.CallOpts)
}

// FraudProofWindowSeconds is a free data retrieval call binding the contract method 0x63042720.
//
// Solidity: function fraudProofWindowSeconds() view returns(uint256)
func (_OptimisticLightClient *OptimisticLightClientCallerSession) FraudProofWindowSeconds() (*big.Int, error) {
	return _OptimisticLightClient.Contract.FraudProofWindowSeconds(&_OptimisticLightClient.CallOpts)
}

// GetFraudProofEndtime is a free data retrieval call binding the contract method 0xd56ff842.
//
// Solidity: function getFraudProofEndtime(uint256 height) view returns(uint256 fraudProofEndTime)
func (_OptimisticLightClient *OptimisticLightClientCaller) GetFraudProofEndtime(opts *bind.CallOpts, height *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OptimisticLightClient.contract.Call(opts, &out, "getFraudProofEndtime", height)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFraudProofEndtime is a free data retrieval call binding the contract method 0xd56ff842.
//
// Solidity: function getFraudProofEndtime(uint256 height) view returns(uint256 fraudProofEndTime)
func (_OptimisticLightClient *OptimisticLightClientSession) GetFraudProofEndtime(height *big.Int) (*big.Int, error) {
	return _OptimisticLightClient.Contract.GetFraudProofEndtime(&_OptimisticLightClient.CallOpts, height)
}

// GetFraudProofEndtime is a free data retrieval call binding the contract method 0xd56ff842.
//
// Solidity: function getFraudProofEndtime(uint256 height) view returns(uint256 fraudProofEndTime)
func (_OptimisticLightClient *OptimisticLightClientCallerSession) GetFraudProofEndtime(height *big.Int) (*big.Int, error) {
	return _OptimisticLightClient.Contract.GetFraudProofEndtime(&_OptimisticLightClient.CallOpts, height)
}

// GetInternalState is a free data retrieval call binding the contract method 0x1bc97a78.
//
// Solidity: function getInternalState(uint256 height) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_OptimisticLightClient *OptimisticLightClientCaller) GetInternalState(opts *bind.CallOpts, height *big.Int) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	var out []interface{}
	err := _OptimisticLightClient.contract.Call(opts, &out, "getInternalState", height)

	outstruct := new(struct {
		AppHash           *big.Int
		FraudProofEndTime *big.Int
		Ended             bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AppHash = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FraudProofEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ended = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetInternalState is a free data retrieval call binding the contract method 0x1bc97a78.
//
// Solidity: function getInternalState(uint256 height) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_OptimisticLightClient *OptimisticLightClientSession) GetInternalState(height *big.Int) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	return _OptimisticLightClient.Contract.GetInternalState(&_OptimisticLightClient.CallOpts, height)
}

// GetInternalState is a free data retrieval call binding the contract method 0x1bc97a78.
//
// Solidity: function getInternalState(uint256 height) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_OptimisticLightClient *OptimisticLightClientCallerSession) GetInternalState(height *big.Int) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	return _OptimisticLightClient.Contract.GetInternalState(&_OptimisticLightClient.CallOpts, height)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_OptimisticLightClient *OptimisticLightClientCaller) GetState(opts *bind.CallOpts, height *big.Int) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	var out []interface{}
	err := _OptimisticLightClient.contract.Call(opts, &out, "getState", height)

	outstruct := new(struct {
		AppHash           *big.Int
		FraudProofEndTime *big.Int
		Ended             bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AppHash = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FraudProofEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ended = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_OptimisticLightClient *OptimisticLightClientSession) GetState(height *big.Int) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	return _OptimisticLightClient.Contract.GetState(&_OptimisticLightClient.CallOpts, height)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_OptimisticLightClient *OptimisticLightClientCallerSession) GetState(height *big.Int) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	return _OptimisticLightClient.Contract.GetState(&_OptimisticLightClient.CallOpts, height)
}

// L1BlockProvider is a free data retrieval call binding the contract method 0xeb772058.
//
// Solidity: function l1BlockProvider() view returns(address)
func (_OptimisticLightClient *OptimisticLightClientCaller) L1BlockProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimisticLightClient.contract.Call(opts, &out, "l1BlockProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1BlockProvider is a free data retrieval call binding the contract method 0xeb772058.
//
// Solidity: function l1BlockProvider() view returns(address)
func (_OptimisticLightClient *OptimisticLightClientSession) L1BlockProvider() (common.Address, error) {
	return _OptimisticLightClient.Contract.L1BlockProvider(&_OptimisticLightClient.CallOpts)
}

// L1BlockProvider is a free data retrieval call binding the contract method 0xeb772058.
//
// Solidity: function l1BlockProvider() view returns(address)
func (_OptimisticLightClient *OptimisticLightClientCallerSession) L1BlockProvider() (common.Address, error) {
	return _OptimisticLightClient.Contract.L1BlockProvider(&_OptimisticLightClient.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_OptimisticLightClient *OptimisticLightClientCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimisticLightClient.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_OptimisticLightClient *OptimisticLightClientSession) Verifier() (common.Address, error) {
	return _OptimisticLightClient.Contract.Verifier(&_OptimisticLightClient.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_OptimisticLightClient *OptimisticLightClientCallerSession) Verifier() (common.Address, error) {
	return _OptimisticLightClient.Contract.Verifier(&_OptimisticLightClient.CallOpts)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key, bytes expectedValue) view returns()
func (_OptimisticLightClient *OptimisticLightClientCaller) VerifyMembership(opts *bind.CallOpts, proof Ics23Proof, key []byte, expectedValue []byte) error {
	var out []interface{}
	err := _OptimisticLightClient.contract.Call(opts, &out, "verifyMembership", proof, key, expectedValue)

	if err != nil {
		return err
	}

	return err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key, bytes expectedValue) view returns()
func (_OptimisticLightClient *OptimisticLightClientSession) VerifyMembership(proof Ics23Proof, key []byte, expectedValue []byte) error {
	return _OptimisticLightClient.Contract.VerifyMembership(&_OptimisticLightClient.CallOpts, proof, key, expectedValue)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key, bytes expectedValue) view returns()
func (_OptimisticLightClient *OptimisticLightClientCallerSession) VerifyMembership(proof Ics23Proof, key []byte, expectedValue []byte) error {
	return _OptimisticLightClient.Contract.VerifyMembership(&_OptimisticLightClient.CallOpts, proof, key, expectedValue)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key) view returns()
func (_OptimisticLightClient *OptimisticLightClientCaller) VerifyNonMembership(opts *bind.CallOpts, proof Ics23Proof, key []byte) error {
	var out []interface{}
	err := _OptimisticLightClient.contract.Call(opts, &out, "verifyNonMembership", proof, key)

	if err != nil {
		return err
	}

	return err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key) view returns()
func (_OptimisticLightClient *OptimisticLightClientSession) VerifyNonMembership(proof Ics23Proof, key []byte) error {
	return _OptimisticLightClient.Contract.VerifyNonMembership(&_OptimisticLightClient.CallOpts, proof, key)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key) view returns()
func (_OptimisticLightClient *OptimisticLightClientCallerSession) VerifyNonMembership(proof Ics23Proof, key []byte) error {
	return _OptimisticLightClient.Contract.VerifyNonMembership(&_OptimisticLightClient.CallOpts, proof, key)
}

// AddOpConsensusState is a paid mutator transaction binding the contract method 0x5922f420.
//
// Solidity: function addOpConsensusState((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, uint256 height, uint256 appHash) returns(uint256 fraudProofEndTime, bool ended)
func (_OptimisticLightClient *OptimisticLightClientTransactor) AddOpConsensusState(opts *bind.TransactOpts, l1header L1Header, proof OpL2StateProof, height *big.Int, appHash *big.Int) (*types.Transaction, error) {
	return _OptimisticLightClient.contract.Transact(opts, "addOpConsensusState", l1header, proof, height, appHash)
}

// AddOpConsensusState is a paid mutator transaction binding the contract method 0x5922f420.
//
// Solidity: function addOpConsensusState((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, uint256 height, uint256 appHash) returns(uint256 fraudProofEndTime, bool ended)
func (_OptimisticLightClient *OptimisticLightClientSession) AddOpConsensusState(l1header L1Header, proof OpL2StateProof, height *big.Int, appHash *big.Int) (*types.Transaction, error) {
	return _OptimisticLightClient.Contract.AddOpConsensusState(&_OptimisticLightClient.TransactOpts, l1header, proof, height, appHash)
}

// AddOpConsensusState is a paid mutator transaction binding the contract method 0x5922f420.
//
// Solidity: function addOpConsensusState((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, uint256 height, uint256 appHash) returns(uint256 fraudProofEndTime, bool ended)
func (_OptimisticLightClient *OptimisticLightClientTransactorSession) AddOpConsensusState(l1header L1Header, proof OpL2StateProof, height *big.Int, appHash *big.Int) (*types.Transaction, error) {
	return _OptimisticLightClient.Contract.AddOpConsensusState(&_OptimisticLightClient.TransactOpts, l1header, proof, height, appHash)
}
