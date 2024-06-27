// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package idispatcher

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

// Channel is an auto generated low-level Go binding around an user-defined struct.
type Channel struct {
	Version               string
	Ordering              uint8
	FeeEnabled            bool
	ConnectionHops        []string
	CounterpartyPortId    string
	CounterpartyChannelId [32]byte
	PortId                string
}

// ChannelEnd is an auto generated low-level Go binding around an user-defined struct.
type ChannelEnd struct {
	PortId    string
	ChannelId [32]byte
	Version   string
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

// IDispatcherMetaData contains all meta data concerning the IDispatcher contract.
var IDispatcherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acknowledgement\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"ack\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelCloseConfirm\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelCloseInit\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelOpenAck\",\"inputs\":[{\"name\":\"local\",\"type\":\"tuple\",\"internalType\":\"structChannelEnd\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structChannelEnd\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelOpenConfirm\",\"inputs\":[{\"name\":\"local\",\"type\":\"tuple\",\"internalType\":\"structChannelEnd\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structChannelEnd\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelOpenInit\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelOpenTry\",\"inputs\":[{\"name\":\"local\",\"type\":\"tuple\",\"internalType\":\"structChannelEnd\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structChannelEnd\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeVault\",\"inputs\":[],\"outputs\":[{\"name\":\"feeVault\",\"type\":\"address\",\"internalType\":\"contractIFeeVault\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getChannel\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"channel\",\"type\":\"tuple\",\"internalType\":\"structChannel\",\"components\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOptimisticConsensusState\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"connection\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"appHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fraudProofEndTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ended\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"portPrefix\",\"inputs\":[],\"outputs\":[{\"name\":\"portPrefix\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeConnection\",\"inputs\":[{\"name\":\"connection\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClientForConnection\",\"inputs\":[{\"name\":\"connection\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"lightClient\",\"type\":\"address\",\"internalType\":\"contractILightClient\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPortPrefix\",\"inputs\":[{\"name\":\"_portPrefix\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"timeout\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateClientWithOptimisticConsensusState\",\"inputs\":[{\"name\":\"l1header\",\"type\":\"tuple\",\"internalType\":\"structL1Header\",\"components\":[{\"name\":\"header\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"number\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structOpL2StateProof\",\"components\":[{\"name\":\"accountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"outputRootProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"l2OutputProposalKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"appHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"connection\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"fraudProofEndTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ended\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeTimeoutPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Acknowledgement\",\"inputs\":[{\"name\":\"sourcePortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sourceChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AcknowledgementError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelCloseConfirm\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelCloseConfirmError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelCloseInit\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelCloseInitError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenAck\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenAckError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenConfirm\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenConfirmError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenInit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenInitError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenTry\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenTryError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecvPacket\",\"inputs\":[{\"name\":\"destPortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendPacket\",\"inputs\":[{\"name\":\"sourcePortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sourceChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Timeout\",\"inputs\":[{\"name\":\"sourcePortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sourceChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimeoutError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WriteAckPacket\",\"inputs\":[{\"name\":\"writerPortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"writerChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"ackPacket\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WriteTimeoutPacket\",\"inputs\":[{\"name\":\"writerPortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"writerChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false}]",
}

// IDispatcherABI is the input ABI used to generate the binding from.
// Deprecated: Use IDispatcherMetaData.ABI instead.
var IDispatcherABI = IDispatcherMetaData.ABI

// IDispatcher is an auto generated Go binding around an Ethereum contract.
type IDispatcher struct {
	IDispatcherCaller     // Read-only binding to the contract
	IDispatcherTransactor // Write-only binding to the contract
	IDispatcherFilterer   // Log filterer for contract events
}

// IDispatcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDispatcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDispatcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDispatcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDispatcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDispatcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDispatcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDispatcherSession struct {
	Contract     *IDispatcher      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDispatcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDispatcherCallerSession struct {
	Contract *IDispatcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IDispatcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDispatcherTransactorSession struct {
	Contract     *IDispatcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IDispatcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDispatcherRaw struct {
	Contract *IDispatcher // Generic contract binding to access the raw methods on
}

// IDispatcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDispatcherCallerRaw struct {
	Contract *IDispatcherCaller // Generic read-only contract binding to access the raw methods on
}

// IDispatcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDispatcherTransactorRaw struct {
	Contract *IDispatcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDispatcher creates a new instance of IDispatcher, bound to a specific deployed contract.
func NewIDispatcher(address common.Address, backend bind.ContractBackend) (*IDispatcher, error) {
	contract, err := bindIDispatcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDispatcher{IDispatcherCaller: IDispatcherCaller{contract: contract}, IDispatcherTransactor: IDispatcherTransactor{contract: contract}, IDispatcherFilterer: IDispatcherFilterer{contract: contract}}, nil
}

// NewIDispatcherCaller creates a new read-only instance of IDispatcher, bound to a specific deployed contract.
func NewIDispatcherCaller(address common.Address, caller bind.ContractCaller) (*IDispatcherCaller, error) {
	contract, err := bindIDispatcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDispatcherCaller{contract: contract}, nil
}

// NewIDispatcherTransactor creates a new write-only instance of IDispatcher, bound to a specific deployed contract.
func NewIDispatcherTransactor(address common.Address, transactor bind.ContractTransactor) (*IDispatcherTransactor, error) {
	contract, err := bindIDispatcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDispatcherTransactor{contract: contract}, nil
}

// NewIDispatcherFilterer creates a new log filterer instance of IDispatcher, bound to a specific deployed contract.
func NewIDispatcherFilterer(address common.Address, filterer bind.ContractFilterer) (*IDispatcherFilterer, error) {
	contract, err := bindIDispatcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDispatcherFilterer{contract: contract}, nil
}

// bindIDispatcher binds a generic wrapper to an already deployed contract.
func bindIDispatcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IDispatcherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDispatcher *IDispatcherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDispatcher.Contract.IDispatcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDispatcher *IDispatcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDispatcher.Contract.IDispatcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDispatcher *IDispatcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDispatcher.Contract.IDispatcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDispatcher *IDispatcherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDispatcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDispatcher *IDispatcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDispatcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDispatcher *IDispatcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDispatcher.Contract.contract.Transact(opts, method, params...)
}

// GetChannel is a free data retrieval call binding the contract method 0x42852d24.
//
// Solidity: function getChannel(address portAddress, bytes32 channelId) view returns((string,uint8,bool,string[],string,bytes32,string) channel)
func (_IDispatcher *IDispatcherCaller) GetChannel(opts *bind.CallOpts, portAddress common.Address, channelId [32]byte) (Channel, error) {
	var out []interface{}
	err := _IDispatcher.contract.Call(opts, &out, "getChannel", portAddress, channelId)

	if err != nil {
		return *new(Channel), err
	}

	out0 := *abi.ConvertType(out[0], new(Channel)).(*Channel)

	return out0, err

}

// GetChannel is a free data retrieval call binding the contract method 0x42852d24.
//
// Solidity: function getChannel(address portAddress, bytes32 channelId) view returns((string,uint8,bool,string[],string,bytes32,string) channel)
func (_IDispatcher *IDispatcherSession) GetChannel(portAddress common.Address, channelId [32]byte) (Channel, error) {
	return _IDispatcher.Contract.GetChannel(&_IDispatcher.CallOpts, portAddress, channelId)
}

// GetChannel is a free data retrieval call binding the contract method 0x42852d24.
//
// Solidity: function getChannel(address portAddress, bytes32 channelId) view returns((string,uint8,bool,string[],string,bytes32,string) channel)
func (_IDispatcher *IDispatcherCallerSession) GetChannel(portAddress common.Address, channelId [32]byte) (Channel, error) {
	return _IDispatcher.Contract.GetChannel(&_IDispatcher.CallOpts, portAddress, channelId)
}

// GetOptimisticConsensusState is a free data retrieval call binding the contract method 0x8dd34bb4.
//
// Solidity: function getOptimisticConsensusState(uint256 height, string connection) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_IDispatcher *IDispatcherCaller) GetOptimisticConsensusState(opts *bind.CallOpts, height *big.Int, connection string) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	var out []interface{}
	err := _IDispatcher.contract.Call(opts, &out, "getOptimisticConsensusState", height, connection)

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

// GetOptimisticConsensusState is a free data retrieval call binding the contract method 0x8dd34bb4.
//
// Solidity: function getOptimisticConsensusState(uint256 height, string connection) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_IDispatcher *IDispatcherSession) GetOptimisticConsensusState(height *big.Int, connection string) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	return _IDispatcher.Contract.GetOptimisticConsensusState(&_IDispatcher.CallOpts, height, connection)
}

// GetOptimisticConsensusState is a free data retrieval call binding the contract method 0x8dd34bb4.
//
// Solidity: function getOptimisticConsensusState(uint256 height, string connection) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_IDispatcher *IDispatcherCallerSession) GetOptimisticConsensusState(height *big.Int, connection string) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	return _IDispatcher.Contract.GetOptimisticConsensusState(&_IDispatcher.CallOpts, height, connection)
}

// PortPrefix is a free data retrieval call binding the contract method 0x7774a6d3.
//
// Solidity: function portPrefix() view returns(string portPrefix)
func (_IDispatcher *IDispatcherCaller) PortPrefix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IDispatcher.contract.Call(opts, &out, "portPrefix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PortPrefix is a free data retrieval call binding the contract method 0x7774a6d3.
//
// Solidity: function portPrefix() view returns(string portPrefix)
func (_IDispatcher *IDispatcherSession) PortPrefix() (string, error) {
	return _IDispatcher.Contract.PortPrefix(&_IDispatcher.CallOpts)
}

// PortPrefix is a free data retrieval call binding the contract method 0x7774a6d3.
//
// Solidity: function portPrefix() view returns(string portPrefix)
func (_IDispatcher *IDispatcherCallerSession) PortPrefix() (string, error) {
	return _IDispatcher.Contract.PortPrefix(&_IDispatcher.CallOpts)
}

// Acknowledgement is a paid mutator transaction binding the contract method 0xba5a4d25.
//
// Solidity: function acknowledgement(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, bytes ack, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactor) Acknowledgement(opts *bind.TransactOpts, packet IbcPacket, ack []byte, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "acknowledgement", packet, ack, proof)
}

// Acknowledgement is a paid mutator transaction binding the contract method 0xba5a4d25.
//
// Solidity: function acknowledgement(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, bytes ack, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherSession) Acknowledgement(packet IbcPacket, ack []byte, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.Acknowledgement(&_IDispatcher.TransactOpts, packet, ack, proof)
}

// Acknowledgement is a paid mutator transaction binding the contract method 0xba5a4d25.
//
// Solidity: function acknowledgement(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, bytes ack, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactorSession) Acknowledgement(packet IbcPacket, ack []byte, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.Acknowledgement(&_IDispatcher.TransactOpts, packet, ack, proof)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0xf90b8e96.
//
// Solidity: function channelCloseConfirm(address portAddress, bytes32 channelId, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactor) ChannelCloseConfirm(opts *bind.TransactOpts, portAddress common.Address, channelId [32]byte, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "channelCloseConfirm", portAddress, channelId, proof)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0xf90b8e96.
//
// Solidity: function channelCloseConfirm(address portAddress, bytes32 channelId, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherSession) ChannelCloseConfirm(portAddress common.Address, channelId [32]byte, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.ChannelCloseConfirm(&_IDispatcher.TransactOpts, portAddress, channelId, proof)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0xf90b8e96.
//
// Solidity: function channelCloseConfirm(address portAddress, bytes32 channelId, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactorSession) ChannelCloseConfirm(portAddress common.Address, channelId [32]byte, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.ChannelCloseConfirm(&_IDispatcher.TransactOpts, portAddress, channelId, proof)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0x81bc079b.
//
// Solidity: function channelCloseInit(bytes32 channelId) returns()
func (_IDispatcher *IDispatcherTransactor) ChannelCloseInit(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "channelCloseInit", channelId)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0x81bc079b.
//
// Solidity: function channelCloseInit(bytes32 channelId) returns()
func (_IDispatcher *IDispatcherSession) ChannelCloseInit(channelId [32]byte) (*types.Transaction, error) {
	return _IDispatcher.Contract.ChannelCloseInit(&_IDispatcher.TransactOpts, channelId)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0x81bc079b.
//
// Solidity: function channelCloseInit(bytes32 channelId) returns()
func (_IDispatcher *IDispatcherTransactorSession) ChannelCloseInit(channelId [32]byte) (*types.Transaction, error) {
	return _IDispatcher.Contract.ChannelCloseInit(&_IDispatcher.TransactOpts, channelId)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0x1eb9fc86.
//
// Solidity: function channelOpenAck((string,bytes32,string) local, string[] connectionHops, uint8 ordering, bool feeEnabled, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactor) ChannelOpenAck(opts *bind.TransactOpts, local ChannelEnd, connectionHops []string, ordering uint8, feeEnabled bool, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "channelOpenAck", local, connectionHops, ordering, feeEnabled, counterparty, proof)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0x1eb9fc86.
//
// Solidity: function channelOpenAck((string,bytes32,string) local, string[] connectionHops, uint8 ordering, bool feeEnabled, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherSession) ChannelOpenAck(local ChannelEnd, connectionHops []string, ordering uint8, feeEnabled bool, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.ChannelOpenAck(&_IDispatcher.TransactOpts, local, connectionHops, ordering, feeEnabled, counterparty, proof)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0x1eb9fc86.
//
// Solidity: function channelOpenAck((string,bytes32,string) local, string[] connectionHops, uint8 ordering, bool feeEnabled, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactorSession) ChannelOpenAck(local ChannelEnd, connectionHops []string, ordering uint8, feeEnabled bool, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.ChannelOpenAck(&_IDispatcher.TransactOpts, local, connectionHops, ordering, feeEnabled, counterparty, proof)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x429446b6.
//
// Solidity: function channelOpenConfirm((string,bytes32,string) local, string[] connectionHops, uint8 ordering, bool feeEnabled, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactor) ChannelOpenConfirm(opts *bind.TransactOpts, local ChannelEnd, connectionHops []string, ordering uint8, feeEnabled bool, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "channelOpenConfirm", local, connectionHops, ordering, feeEnabled, counterparty, proof)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x429446b6.
//
// Solidity: function channelOpenConfirm((string,bytes32,string) local, string[] connectionHops, uint8 ordering, bool feeEnabled, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherSession) ChannelOpenConfirm(local ChannelEnd, connectionHops []string, ordering uint8, feeEnabled bool, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.ChannelOpenConfirm(&_IDispatcher.TransactOpts, local, connectionHops, ordering, feeEnabled, counterparty, proof)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x429446b6.
//
// Solidity: function channelOpenConfirm((string,bytes32,string) local, string[] connectionHops, uint8 ordering, bool feeEnabled, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactorSession) ChannelOpenConfirm(local ChannelEnd, connectionHops []string, ordering uint8, feeEnabled bool, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.ChannelOpenConfirm(&_IDispatcher.TransactOpts, local, connectionHops, ordering, feeEnabled, counterparty, proof)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0x418925b7.
//
// Solidity: function channelOpenInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_IDispatcher *IDispatcherTransactor) ChannelOpenInit(opts *bind.TransactOpts, version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "channelOpenInit", version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0x418925b7.
//
// Solidity: function channelOpenInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_IDispatcher *IDispatcherSession) ChannelOpenInit(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _IDispatcher.Contract.ChannelOpenInit(&_IDispatcher.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0x418925b7.
//
// Solidity: function channelOpenInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_IDispatcher *IDispatcherTransactorSession) ChannelOpenInit(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _IDispatcher.Contract.ChannelOpenInit(&_IDispatcher.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x2bf5d19d.
//
// Solidity: function channelOpenTry((string,bytes32,string) local, uint8 ordering, bool feeEnabled, string[] connectionHops, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactor) ChannelOpenTry(opts *bind.TransactOpts, local ChannelEnd, ordering uint8, feeEnabled bool, connectionHops []string, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "channelOpenTry", local, ordering, feeEnabled, connectionHops, counterparty, proof)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x2bf5d19d.
//
// Solidity: function channelOpenTry((string,bytes32,string) local, uint8 ordering, bool feeEnabled, string[] connectionHops, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherSession) ChannelOpenTry(local ChannelEnd, ordering uint8, feeEnabled bool, connectionHops []string, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.ChannelOpenTry(&_IDispatcher.TransactOpts, local, ordering, feeEnabled, connectionHops, counterparty, proof)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x2bf5d19d.
//
// Solidity: function channelOpenTry((string,bytes32,string) local, uint8 ordering, bool feeEnabled, string[] connectionHops, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactorSession) ChannelOpenTry(local ChannelEnd, ordering uint8, feeEnabled bool, connectionHops []string, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.ChannelOpenTry(&_IDispatcher.TransactOpts, local, ordering, feeEnabled, connectionHops, counterparty, proof)
}

// FeeVault is a paid mutator transaction binding the contract method 0x478222c2.
//
// Solidity: function feeVault() returns(address feeVault)
func (_IDispatcher *IDispatcherTransactor) FeeVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "feeVault")
}

// FeeVault is a paid mutator transaction binding the contract method 0x478222c2.
//
// Solidity: function feeVault() returns(address feeVault)
func (_IDispatcher *IDispatcherSession) FeeVault() (*types.Transaction, error) {
	return _IDispatcher.Contract.FeeVault(&_IDispatcher.TransactOpts)
}

// FeeVault is a paid mutator transaction binding the contract method 0x478222c2.
//
// Solidity: function feeVault() returns(address feeVault)
func (_IDispatcher *IDispatcherTransactorSession) FeeVault() (*types.Transaction, error) {
	return _IDispatcher.Contract.FeeVault(&_IDispatcher.TransactOpts)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x6b67055e.
//
// Solidity: function recvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactor) RecvPacket(opts *bind.TransactOpts, packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "recvPacket", packet, proof)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x6b67055e.
//
// Solidity: function recvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherSession) RecvPacket(packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.RecvPacket(&_IDispatcher.TransactOpts, packet, proof)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x6b67055e.
//
// Solidity: function recvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactorSession) RecvPacket(packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.RecvPacket(&_IDispatcher.TransactOpts, packet, proof)
}

// RemoveConnection is a paid mutator transaction binding the contract method 0xc00fa7c0.
//
// Solidity: function removeConnection(string connection) returns()
func (_IDispatcher *IDispatcherTransactor) RemoveConnection(opts *bind.TransactOpts, connection string) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "removeConnection", connection)
}

// RemoveConnection is a paid mutator transaction binding the contract method 0xc00fa7c0.
//
// Solidity: function removeConnection(string connection) returns()
func (_IDispatcher *IDispatcherSession) RemoveConnection(connection string) (*types.Transaction, error) {
	return _IDispatcher.Contract.RemoveConnection(&_IDispatcher.TransactOpts, connection)
}

// RemoveConnection is a paid mutator transaction binding the contract method 0xc00fa7c0.
//
// Solidity: function removeConnection(string connection) returns()
func (_IDispatcher *IDispatcherTransactorSession) RemoveConnection(connection string) (*types.Transaction, error) {
	return _IDispatcher.Contract.RemoveConnection(&_IDispatcher.TransactOpts, connection)
}

// SendPacket is a paid mutator transaction binding the contract method 0xc3e1155c.
//
// Solidity: function sendPacket(bytes32 channelId, bytes packet, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IDispatcher *IDispatcherTransactor) SendPacket(opts *bind.TransactOpts, channelId [32]byte, packet []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "sendPacket", channelId, packet, timeoutTimestamp)
}

// SendPacket is a paid mutator transaction binding the contract method 0xc3e1155c.
//
// Solidity: function sendPacket(bytes32 channelId, bytes packet, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IDispatcher *IDispatcherSession) SendPacket(channelId [32]byte, packet []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IDispatcher.Contract.SendPacket(&_IDispatcher.TransactOpts, channelId, packet, timeoutTimestamp)
}

// SendPacket is a paid mutator transaction binding the contract method 0xc3e1155c.
//
// Solidity: function sendPacket(bytes32 channelId, bytes packet, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IDispatcher *IDispatcherTransactorSession) SendPacket(channelId [32]byte, packet []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IDispatcher.Contract.SendPacket(&_IDispatcher.TransactOpts, channelId, packet, timeoutTimestamp)
}

// SetClientForConnection is a paid mutator transaction binding the contract method 0x556d5178.
//
// Solidity: function setClientForConnection(string connection, address lightClient) returns()
func (_IDispatcher *IDispatcherTransactor) SetClientForConnection(opts *bind.TransactOpts, connection string, lightClient common.Address) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "setClientForConnection", connection, lightClient)
}

// SetClientForConnection is a paid mutator transaction binding the contract method 0x556d5178.
//
// Solidity: function setClientForConnection(string connection, address lightClient) returns()
func (_IDispatcher *IDispatcherSession) SetClientForConnection(connection string, lightClient common.Address) (*types.Transaction, error) {
	return _IDispatcher.Contract.SetClientForConnection(&_IDispatcher.TransactOpts, connection, lightClient)
}

// SetClientForConnection is a paid mutator transaction binding the contract method 0x556d5178.
//
// Solidity: function setClientForConnection(string connection, address lightClient) returns()
func (_IDispatcher *IDispatcherTransactorSession) SetClientForConnection(connection string, lightClient common.Address) (*types.Transaction, error) {
	return _IDispatcher.Contract.SetClientForConnection(&_IDispatcher.TransactOpts, connection, lightClient)
}

// SetPortPrefix is a paid mutator transaction binding the contract method 0x9f59ae71.
//
// Solidity: function setPortPrefix(string _portPrefix) returns()
func (_IDispatcher *IDispatcherTransactor) SetPortPrefix(opts *bind.TransactOpts, _portPrefix string) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "setPortPrefix", _portPrefix)
}

// SetPortPrefix is a paid mutator transaction binding the contract method 0x9f59ae71.
//
// Solidity: function setPortPrefix(string _portPrefix) returns()
func (_IDispatcher *IDispatcherSession) SetPortPrefix(_portPrefix string) (*types.Transaction, error) {
	return _IDispatcher.Contract.SetPortPrefix(&_IDispatcher.TransactOpts, _portPrefix)
}

// SetPortPrefix is a paid mutator transaction binding the contract method 0x9f59ae71.
//
// Solidity: function setPortPrefix(string _portPrefix) returns()
func (_IDispatcher *IDispatcherTransactorSession) SetPortPrefix(_portPrefix string) (*types.Transaction, error) {
	return _IDispatcher.Contract.SetPortPrefix(&_IDispatcher.TransactOpts, _portPrefix)
}

// Timeout is a paid mutator transaction binding the contract method 0x6050b5f3.
//
// Solidity: function timeout(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactor) Timeout(opts *bind.TransactOpts, packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "timeout", packet, proof)
}

// Timeout is a paid mutator transaction binding the contract method 0x6050b5f3.
//
// Solidity: function timeout(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherSession) Timeout(packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.Timeout(&_IDispatcher.TransactOpts, packet, proof)
}

// Timeout is a paid mutator transaction binding the contract method 0x6050b5f3.
//
// Solidity: function timeout(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactorSession) Timeout(packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.Timeout(&_IDispatcher.TransactOpts, packet, proof)
}

// UpdateClientWithOptimisticConsensusState is a paid mutator transaction binding the contract method 0x940265cb.
//
// Solidity: function updateClientWithOptimisticConsensusState((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, uint256 height, uint256 appHash, string connection) returns(uint256 fraudProofEndTime, bool ended)
func (_IDispatcher *IDispatcherTransactor) UpdateClientWithOptimisticConsensusState(opts *bind.TransactOpts, l1header L1Header, proof OpL2StateProof, height *big.Int, appHash *big.Int, connection string) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "updateClientWithOptimisticConsensusState", l1header, proof, height, appHash, connection)
}

// UpdateClientWithOptimisticConsensusState is a paid mutator transaction binding the contract method 0x940265cb.
//
// Solidity: function updateClientWithOptimisticConsensusState((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, uint256 height, uint256 appHash, string connection) returns(uint256 fraudProofEndTime, bool ended)
func (_IDispatcher *IDispatcherSession) UpdateClientWithOptimisticConsensusState(l1header L1Header, proof OpL2StateProof, height *big.Int, appHash *big.Int, connection string) (*types.Transaction, error) {
	return _IDispatcher.Contract.UpdateClientWithOptimisticConsensusState(&_IDispatcher.TransactOpts, l1header, proof, height, appHash, connection)
}

// UpdateClientWithOptimisticConsensusState is a paid mutator transaction binding the contract method 0x940265cb.
//
// Solidity: function updateClientWithOptimisticConsensusState((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, uint256 height, uint256 appHash, string connection) returns(uint256 fraudProofEndTime, bool ended)
func (_IDispatcher *IDispatcherTransactorSession) UpdateClientWithOptimisticConsensusState(l1header L1Header, proof OpL2StateProof, height *big.Int, appHash *big.Int, connection string) (*types.Transaction, error) {
	return _IDispatcher.Contract.UpdateClientWithOptimisticConsensusState(&_IDispatcher.TransactOpts, l1header, proof, height, appHash, connection)
}

// WriteTimeoutPacket is a paid mutator transaction binding the contract method 0x5d7adf96.
//
// Solidity: function writeTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactor) WriteTimeoutPacket(opts *bind.TransactOpts, packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.contract.Transact(opts, "writeTimeoutPacket", packet, proof)
}

// WriteTimeoutPacket is a paid mutator transaction binding the contract method 0x5d7adf96.
//
// Solidity: function writeTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherSession) WriteTimeoutPacket(packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.WriteTimeoutPacket(&_IDispatcher.TransactOpts, packet, proof)
}

// WriteTimeoutPacket is a paid mutator transaction binding the contract method 0x5d7adf96.
//
// Solidity: function writeTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IDispatcher *IDispatcherTransactorSession) WriteTimeoutPacket(packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _IDispatcher.Contract.WriteTimeoutPacket(&_IDispatcher.TransactOpts, packet, proof)
}

// IDispatcherAcknowledgementIterator is returned from FilterAcknowledgement and is used to iterate over the raw logs and unpacked data for Acknowledgement events raised by the IDispatcher contract.
type IDispatcherAcknowledgementIterator struct {
	Event *IDispatcherAcknowledgement // Event containing the contract specifics and raw log

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
func (it *IDispatcherAcknowledgementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherAcknowledgement)
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
		it.Event = new(IDispatcherAcknowledgement)
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
func (it *IDispatcherAcknowledgementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherAcknowledgementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherAcknowledgement represents a Acknowledgement event raised by the IDispatcher contract.
type IDispatcherAcknowledgement struct {
	SourcePortAddress common.Address
	SourceChannelId   [32]byte
	Sequence          uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAcknowledgement is a free log retrieval operation binding the contract event 0xe46f6591236abe528fe47a3b281fb002524dadd3e62b1f317ed285d07273c3b1.
//
// Solidity: event Acknowledgement(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 sequence)
func (_IDispatcher *IDispatcherFilterer) FilterAcknowledgement(opts *bind.FilterOpts, sourcePortAddress []common.Address, sourceChannelId [][32]byte) (*IDispatcherAcknowledgementIterator, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "Acknowledgement", sourcePortAddressRule, sourceChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherAcknowledgementIterator{contract: _IDispatcher.contract, event: "Acknowledgement", logs: logs, sub: sub}, nil
}

// WatchAcknowledgement is a free log subscription operation binding the contract event 0xe46f6591236abe528fe47a3b281fb002524dadd3e62b1f317ed285d07273c3b1.
//
// Solidity: event Acknowledgement(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 sequence)
func (_IDispatcher *IDispatcherFilterer) WatchAcknowledgement(opts *bind.WatchOpts, sink chan<- *IDispatcherAcknowledgement, sourcePortAddress []common.Address, sourceChannelId [][32]byte) (event.Subscription, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "Acknowledgement", sourcePortAddressRule, sourceChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherAcknowledgement)
				if err := _IDispatcher.contract.UnpackLog(event, "Acknowledgement", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseAcknowledgement(log types.Log) (*IDispatcherAcknowledgement, error) {
	event := new(IDispatcherAcknowledgement)
	if err := _IDispatcher.contract.UnpackLog(event, "Acknowledgement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherAcknowledgementErrorIterator is returned from FilterAcknowledgementError and is used to iterate over the raw logs and unpacked data for AcknowledgementError events raised by the IDispatcher contract.
type IDispatcherAcknowledgementErrorIterator struct {
	Event *IDispatcherAcknowledgementError // Event containing the contract specifics and raw log

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
func (it *IDispatcherAcknowledgementErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherAcknowledgementError)
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
		it.Event = new(IDispatcherAcknowledgementError)
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
func (it *IDispatcherAcknowledgementErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherAcknowledgementErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherAcknowledgementError represents a AcknowledgementError event raised by the IDispatcher contract.
type IDispatcherAcknowledgementError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcknowledgementError is a free log retrieval operation binding the contract event 0x625eea143c9dae6915c809da47016c22d9cd006c3ace7c345c5cbcf57d3aefbc.
//
// Solidity: event AcknowledgementError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) FilterAcknowledgementError(opts *bind.FilterOpts, receiver []common.Address) (*IDispatcherAcknowledgementErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "AcknowledgementError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherAcknowledgementErrorIterator{contract: _IDispatcher.contract, event: "AcknowledgementError", logs: logs, sub: sub}, nil
}

// WatchAcknowledgementError is a free log subscription operation binding the contract event 0x625eea143c9dae6915c809da47016c22d9cd006c3ace7c345c5cbcf57d3aefbc.
//
// Solidity: event AcknowledgementError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) WatchAcknowledgementError(opts *bind.WatchOpts, sink chan<- *IDispatcherAcknowledgementError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "AcknowledgementError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherAcknowledgementError)
				if err := _IDispatcher.contract.UnpackLog(event, "AcknowledgementError", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseAcknowledgementError(log types.Log) (*IDispatcherAcknowledgementError, error) {
	event := new(IDispatcherAcknowledgementError)
	if err := _IDispatcher.contract.UnpackLog(event, "AcknowledgementError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherChannelCloseConfirmIterator is returned from FilterChannelCloseConfirm and is used to iterate over the raw logs and unpacked data for ChannelCloseConfirm events raised by the IDispatcher contract.
type IDispatcherChannelCloseConfirmIterator struct {
	Event *IDispatcherChannelCloseConfirm // Event containing the contract specifics and raw log

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
func (it *IDispatcherChannelCloseConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherChannelCloseConfirm)
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
		it.Event = new(IDispatcherChannelCloseConfirm)
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
func (it *IDispatcherChannelCloseConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherChannelCloseConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherChannelCloseConfirm represents a ChannelCloseConfirm event raised by the IDispatcher contract.
type IDispatcherChannelCloseConfirm struct {
	PortAddress common.Address
	ChannelId   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseConfirm is a free log retrieval operation binding the contract event 0x5f010dbbd6bf46aec8131c5456301a75cd688d3cca9bc8380c9e26301be20729.
//
// Solidity: event ChannelCloseConfirm(address indexed portAddress, bytes32 indexed channelId)
func (_IDispatcher *IDispatcherFilterer) FilterChannelCloseConfirm(opts *bind.FilterOpts, portAddress []common.Address, channelId [][32]byte) (*IDispatcherChannelCloseConfirmIterator, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "ChannelCloseConfirm", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherChannelCloseConfirmIterator{contract: _IDispatcher.contract, event: "ChannelCloseConfirm", logs: logs, sub: sub}, nil
}

// WatchChannelCloseConfirm is a free log subscription operation binding the contract event 0x5f010dbbd6bf46aec8131c5456301a75cd688d3cca9bc8380c9e26301be20729.
//
// Solidity: event ChannelCloseConfirm(address indexed portAddress, bytes32 indexed channelId)
func (_IDispatcher *IDispatcherFilterer) WatchChannelCloseConfirm(opts *bind.WatchOpts, sink chan<- *IDispatcherChannelCloseConfirm, portAddress []common.Address, channelId [][32]byte) (event.Subscription, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "ChannelCloseConfirm", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherChannelCloseConfirm)
				if err := _IDispatcher.contract.UnpackLog(event, "ChannelCloseConfirm", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseChannelCloseConfirm(log types.Log) (*IDispatcherChannelCloseConfirm, error) {
	event := new(IDispatcherChannelCloseConfirm)
	if err := _IDispatcher.contract.UnpackLog(event, "ChannelCloseConfirm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherChannelCloseConfirmErrorIterator is returned from FilterChannelCloseConfirmError and is used to iterate over the raw logs and unpacked data for ChannelCloseConfirmError events raised by the IDispatcher contract.
type IDispatcherChannelCloseConfirmErrorIterator struct {
	Event *IDispatcherChannelCloseConfirmError // Event containing the contract specifics and raw log

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
func (it *IDispatcherChannelCloseConfirmErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherChannelCloseConfirmError)
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
		it.Event = new(IDispatcherChannelCloseConfirmError)
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
func (it *IDispatcherChannelCloseConfirmErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherChannelCloseConfirmErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherChannelCloseConfirmError represents a ChannelCloseConfirmError event raised by the IDispatcher contract.
type IDispatcherChannelCloseConfirmError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseConfirmError is a free log retrieval operation binding the contract event 0xc9d36d7a317cb116925d5cb66f0069fe825822fe23e9cf3f421c38cf444caa30.
//
// Solidity: event ChannelCloseConfirmError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) FilterChannelCloseConfirmError(opts *bind.FilterOpts, receiver []common.Address) (*IDispatcherChannelCloseConfirmErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "ChannelCloseConfirmError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherChannelCloseConfirmErrorIterator{contract: _IDispatcher.contract, event: "ChannelCloseConfirmError", logs: logs, sub: sub}, nil
}

// WatchChannelCloseConfirmError is a free log subscription operation binding the contract event 0xc9d36d7a317cb116925d5cb66f0069fe825822fe23e9cf3f421c38cf444caa30.
//
// Solidity: event ChannelCloseConfirmError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) WatchChannelCloseConfirmError(opts *bind.WatchOpts, sink chan<- *IDispatcherChannelCloseConfirmError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "ChannelCloseConfirmError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherChannelCloseConfirmError)
				if err := _IDispatcher.contract.UnpackLog(event, "ChannelCloseConfirmError", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseChannelCloseConfirmError(log types.Log) (*IDispatcherChannelCloseConfirmError, error) {
	event := new(IDispatcherChannelCloseConfirmError)
	if err := _IDispatcher.contract.UnpackLog(event, "ChannelCloseConfirmError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherChannelCloseInitIterator is returned from FilterChannelCloseInit and is used to iterate over the raw logs and unpacked data for ChannelCloseInit events raised by the IDispatcher contract.
type IDispatcherChannelCloseInitIterator struct {
	Event *IDispatcherChannelCloseInit // Event containing the contract specifics and raw log

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
func (it *IDispatcherChannelCloseInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherChannelCloseInit)
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
		it.Event = new(IDispatcherChannelCloseInit)
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
func (it *IDispatcherChannelCloseInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherChannelCloseInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherChannelCloseInit represents a ChannelCloseInit event raised by the IDispatcher contract.
type IDispatcherChannelCloseInit struct {
	PortAddress common.Address
	ChannelId   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseInit is a free log retrieval operation binding the contract event 0x21372e37743553ba8e5f61239803174a827c7732d53ec8adcb76c6b3bb2c13f1.
//
// Solidity: event ChannelCloseInit(address indexed portAddress, bytes32 indexed channelId)
func (_IDispatcher *IDispatcherFilterer) FilterChannelCloseInit(opts *bind.FilterOpts, portAddress []common.Address, channelId [][32]byte) (*IDispatcherChannelCloseInitIterator, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "ChannelCloseInit", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherChannelCloseInitIterator{contract: _IDispatcher.contract, event: "ChannelCloseInit", logs: logs, sub: sub}, nil
}

// WatchChannelCloseInit is a free log subscription operation binding the contract event 0x21372e37743553ba8e5f61239803174a827c7732d53ec8adcb76c6b3bb2c13f1.
//
// Solidity: event ChannelCloseInit(address indexed portAddress, bytes32 indexed channelId)
func (_IDispatcher *IDispatcherFilterer) WatchChannelCloseInit(opts *bind.WatchOpts, sink chan<- *IDispatcherChannelCloseInit, portAddress []common.Address, channelId [][32]byte) (event.Subscription, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "ChannelCloseInit", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherChannelCloseInit)
				if err := _IDispatcher.contract.UnpackLog(event, "ChannelCloseInit", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseChannelCloseInit(log types.Log) (*IDispatcherChannelCloseInit, error) {
	event := new(IDispatcherChannelCloseInit)
	if err := _IDispatcher.contract.UnpackLog(event, "ChannelCloseInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherChannelCloseInitErrorIterator is returned from FilterChannelCloseInitError and is used to iterate over the raw logs and unpacked data for ChannelCloseInitError events raised by the IDispatcher contract.
type IDispatcherChannelCloseInitErrorIterator struct {
	Event *IDispatcherChannelCloseInitError // Event containing the contract specifics and raw log

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
func (it *IDispatcherChannelCloseInitErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherChannelCloseInitError)
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
		it.Event = new(IDispatcherChannelCloseInitError)
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
func (it *IDispatcherChannelCloseInitErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherChannelCloseInitErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherChannelCloseInitError represents a ChannelCloseInitError event raised by the IDispatcher contract.
type IDispatcherChannelCloseInitError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseInitError is a free log retrieval operation binding the contract event 0xb1be59c1bcd39c54c7132a8e0d321af5db427575ddb3265560d8862804f4381b.
//
// Solidity: event ChannelCloseInitError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) FilterChannelCloseInitError(opts *bind.FilterOpts, receiver []common.Address) (*IDispatcherChannelCloseInitErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "ChannelCloseInitError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherChannelCloseInitErrorIterator{contract: _IDispatcher.contract, event: "ChannelCloseInitError", logs: logs, sub: sub}, nil
}

// WatchChannelCloseInitError is a free log subscription operation binding the contract event 0xb1be59c1bcd39c54c7132a8e0d321af5db427575ddb3265560d8862804f4381b.
//
// Solidity: event ChannelCloseInitError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) WatchChannelCloseInitError(opts *bind.WatchOpts, sink chan<- *IDispatcherChannelCloseInitError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "ChannelCloseInitError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherChannelCloseInitError)
				if err := _IDispatcher.contract.UnpackLog(event, "ChannelCloseInitError", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseChannelCloseInitError(log types.Log) (*IDispatcherChannelCloseInitError, error) {
	event := new(IDispatcherChannelCloseInitError)
	if err := _IDispatcher.contract.UnpackLog(event, "ChannelCloseInitError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherChannelOpenAckIterator is returned from FilterChannelOpenAck and is used to iterate over the raw logs and unpacked data for ChannelOpenAck events raised by the IDispatcher contract.
type IDispatcherChannelOpenAckIterator struct {
	Event *IDispatcherChannelOpenAck // Event containing the contract specifics and raw log

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
func (it *IDispatcherChannelOpenAckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherChannelOpenAck)
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
		it.Event = new(IDispatcherChannelOpenAck)
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
func (it *IDispatcherChannelOpenAckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherChannelOpenAckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherChannelOpenAck represents a ChannelOpenAck event raised by the IDispatcher contract.
type IDispatcherChannelOpenAck struct {
	Receiver  common.Address
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenAck is a free log retrieval operation binding the contract event 0xcf8be9ab2b5edf8beb2c45abe8e0cc7646318ac19f6c3164ba2e19e93a8a32af.
//
// Solidity: event ChannelOpenAck(address indexed receiver, bytes32 channelId)
func (_IDispatcher *IDispatcherFilterer) FilterChannelOpenAck(opts *bind.FilterOpts, receiver []common.Address) (*IDispatcherChannelOpenAckIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "ChannelOpenAck", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherChannelOpenAckIterator{contract: _IDispatcher.contract, event: "ChannelOpenAck", logs: logs, sub: sub}, nil
}

// WatchChannelOpenAck is a free log subscription operation binding the contract event 0xcf8be9ab2b5edf8beb2c45abe8e0cc7646318ac19f6c3164ba2e19e93a8a32af.
//
// Solidity: event ChannelOpenAck(address indexed receiver, bytes32 channelId)
func (_IDispatcher *IDispatcherFilterer) WatchChannelOpenAck(opts *bind.WatchOpts, sink chan<- *IDispatcherChannelOpenAck, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "ChannelOpenAck", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherChannelOpenAck)
				if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenAck", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseChannelOpenAck(log types.Log) (*IDispatcherChannelOpenAck, error) {
	event := new(IDispatcherChannelOpenAck)
	if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenAck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherChannelOpenAckErrorIterator is returned from FilterChannelOpenAckError and is used to iterate over the raw logs and unpacked data for ChannelOpenAckError events raised by the IDispatcher contract.
type IDispatcherChannelOpenAckErrorIterator struct {
	Event *IDispatcherChannelOpenAckError // Event containing the contract specifics and raw log

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
func (it *IDispatcherChannelOpenAckErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherChannelOpenAckError)
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
		it.Event = new(IDispatcherChannelOpenAckError)
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
func (it *IDispatcherChannelOpenAckErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherChannelOpenAckErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherChannelOpenAckError represents a ChannelOpenAckError event raised by the IDispatcher contract.
type IDispatcherChannelOpenAckError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenAckError is a free log retrieval operation binding the contract event 0x971a4433f5bff5f011728a4123aeeca4b5275ac20b013cf276e65510491ac26f.
//
// Solidity: event ChannelOpenAckError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) FilterChannelOpenAckError(opts *bind.FilterOpts, receiver []common.Address) (*IDispatcherChannelOpenAckErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "ChannelOpenAckError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherChannelOpenAckErrorIterator{contract: _IDispatcher.contract, event: "ChannelOpenAckError", logs: logs, sub: sub}, nil
}

// WatchChannelOpenAckError is a free log subscription operation binding the contract event 0x971a4433f5bff5f011728a4123aeeca4b5275ac20b013cf276e65510491ac26f.
//
// Solidity: event ChannelOpenAckError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) WatchChannelOpenAckError(opts *bind.WatchOpts, sink chan<- *IDispatcherChannelOpenAckError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "ChannelOpenAckError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherChannelOpenAckError)
				if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenAckError", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseChannelOpenAckError(log types.Log) (*IDispatcherChannelOpenAckError, error) {
	event := new(IDispatcherChannelOpenAckError)
	if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenAckError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherChannelOpenConfirmIterator is returned from FilterChannelOpenConfirm and is used to iterate over the raw logs and unpacked data for ChannelOpenConfirm events raised by the IDispatcher contract.
type IDispatcherChannelOpenConfirmIterator struct {
	Event *IDispatcherChannelOpenConfirm // Event containing the contract specifics and raw log

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
func (it *IDispatcherChannelOpenConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherChannelOpenConfirm)
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
		it.Event = new(IDispatcherChannelOpenConfirm)
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
func (it *IDispatcherChannelOpenConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherChannelOpenConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherChannelOpenConfirm represents a ChannelOpenConfirm event raised by the IDispatcher contract.
type IDispatcherChannelOpenConfirm struct {
	Receiver  common.Address
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenConfirm is a free log retrieval operation binding the contract event 0xe80f571f70f7cabf9d7ac60ece08421be374117776c311c327a083ca398f802f.
//
// Solidity: event ChannelOpenConfirm(address indexed receiver, bytes32 channelId)
func (_IDispatcher *IDispatcherFilterer) FilterChannelOpenConfirm(opts *bind.FilterOpts, receiver []common.Address) (*IDispatcherChannelOpenConfirmIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "ChannelOpenConfirm", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherChannelOpenConfirmIterator{contract: _IDispatcher.contract, event: "ChannelOpenConfirm", logs: logs, sub: sub}, nil
}

// WatchChannelOpenConfirm is a free log subscription operation binding the contract event 0xe80f571f70f7cabf9d7ac60ece08421be374117776c311c327a083ca398f802f.
//
// Solidity: event ChannelOpenConfirm(address indexed receiver, bytes32 channelId)
func (_IDispatcher *IDispatcherFilterer) WatchChannelOpenConfirm(opts *bind.WatchOpts, sink chan<- *IDispatcherChannelOpenConfirm, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "ChannelOpenConfirm", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherChannelOpenConfirm)
				if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenConfirm", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseChannelOpenConfirm(log types.Log) (*IDispatcherChannelOpenConfirm, error) {
	event := new(IDispatcherChannelOpenConfirm)
	if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenConfirm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherChannelOpenConfirmErrorIterator is returned from FilterChannelOpenConfirmError and is used to iterate over the raw logs and unpacked data for ChannelOpenConfirmError events raised by the IDispatcher contract.
type IDispatcherChannelOpenConfirmErrorIterator struct {
	Event *IDispatcherChannelOpenConfirmError // Event containing the contract specifics and raw log

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
func (it *IDispatcherChannelOpenConfirmErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherChannelOpenConfirmError)
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
		it.Event = new(IDispatcherChannelOpenConfirmError)
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
func (it *IDispatcherChannelOpenConfirmErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherChannelOpenConfirmErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherChannelOpenConfirmError represents a ChannelOpenConfirmError event raised by the IDispatcher contract.
type IDispatcherChannelOpenConfirmError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenConfirmError is a free log retrieval operation binding the contract event 0xf6a58ef30f66943749e8c29c661c84da143a1c8ed017f5faa92b509e0000875a.
//
// Solidity: event ChannelOpenConfirmError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) FilterChannelOpenConfirmError(opts *bind.FilterOpts, receiver []common.Address) (*IDispatcherChannelOpenConfirmErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "ChannelOpenConfirmError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherChannelOpenConfirmErrorIterator{contract: _IDispatcher.contract, event: "ChannelOpenConfirmError", logs: logs, sub: sub}, nil
}

// WatchChannelOpenConfirmError is a free log subscription operation binding the contract event 0xf6a58ef30f66943749e8c29c661c84da143a1c8ed017f5faa92b509e0000875a.
//
// Solidity: event ChannelOpenConfirmError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) WatchChannelOpenConfirmError(opts *bind.WatchOpts, sink chan<- *IDispatcherChannelOpenConfirmError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "ChannelOpenConfirmError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherChannelOpenConfirmError)
				if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenConfirmError", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseChannelOpenConfirmError(log types.Log) (*IDispatcherChannelOpenConfirmError, error) {
	event := new(IDispatcherChannelOpenConfirmError)
	if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenConfirmError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherChannelOpenInitIterator is returned from FilterChannelOpenInit and is used to iterate over the raw logs and unpacked data for ChannelOpenInit events raised by the IDispatcher contract.
type IDispatcherChannelOpenInitIterator struct {
	Event *IDispatcherChannelOpenInit // Event containing the contract specifics and raw log

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
func (it *IDispatcherChannelOpenInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherChannelOpenInit)
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
		it.Event = new(IDispatcherChannelOpenInit)
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
func (it *IDispatcherChannelOpenInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherChannelOpenInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherChannelOpenInit represents a ChannelOpenInit event raised by the IDispatcher contract.
type IDispatcherChannelOpenInit struct {
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
func (_IDispatcher *IDispatcherFilterer) FilterChannelOpenInit(opts *bind.FilterOpts, receiver []common.Address) (*IDispatcherChannelOpenInitIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "ChannelOpenInit", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherChannelOpenInitIterator{contract: _IDispatcher.contract, event: "ChannelOpenInit", logs: logs, sub: sub}, nil
}

// WatchChannelOpenInit is a free log subscription operation binding the contract event 0x20fd8a5856711b18d00def4aa6abafbe00ce6d60795e015cc1cad35eb9b46359.
//
// Solidity: event ChannelOpenInit(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId)
func (_IDispatcher *IDispatcherFilterer) WatchChannelOpenInit(opts *bind.WatchOpts, sink chan<- *IDispatcherChannelOpenInit, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "ChannelOpenInit", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherChannelOpenInit)
				if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenInit", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseChannelOpenInit(log types.Log) (*IDispatcherChannelOpenInit, error) {
	event := new(IDispatcherChannelOpenInit)
	if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherChannelOpenInitErrorIterator is returned from FilterChannelOpenInitError and is used to iterate over the raw logs and unpacked data for ChannelOpenInitError events raised by the IDispatcher contract.
type IDispatcherChannelOpenInitErrorIterator struct {
	Event *IDispatcherChannelOpenInitError // Event containing the contract specifics and raw log

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
func (it *IDispatcherChannelOpenInitErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherChannelOpenInitError)
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
		it.Event = new(IDispatcherChannelOpenInitError)
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
func (it *IDispatcherChannelOpenInitErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherChannelOpenInitErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherChannelOpenInitError represents a ChannelOpenInitError event raised by the IDispatcher contract.
type IDispatcherChannelOpenInitError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenInitError is a free log retrieval operation binding the contract event 0x69c1283cce89382f0f9ddf19b7c4f05b4d9b3c30c84fc148b1ec800284be58d5.
//
// Solidity: event ChannelOpenInitError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) FilterChannelOpenInitError(opts *bind.FilterOpts, receiver []common.Address) (*IDispatcherChannelOpenInitErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "ChannelOpenInitError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherChannelOpenInitErrorIterator{contract: _IDispatcher.contract, event: "ChannelOpenInitError", logs: logs, sub: sub}, nil
}

// WatchChannelOpenInitError is a free log subscription operation binding the contract event 0x69c1283cce89382f0f9ddf19b7c4f05b4d9b3c30c84fc148b1ec800284be58d5.
//
// Solidity: event ChannelOpenInitError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) WatchChannelOpenInitError(opts *bind.WatchOpts, sink chan<- *IDispatcherChannelOpenInitError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "ChannelOpenInitError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherChannelOpenInitError)
				if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenInitError", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseChannelOpenInitError(log types.Log) (*IDispatcherChannelOpenInitError, error) {
	event := new(IDispatcherChannelOpenInitError)
	if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenInitError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherChannelOpenTryIterator is returned from FilterChannelOpenTry and is used to iterate over the raw logs and unpacked data for ChannelOpenTry events raised by the IDispatcher contract.
type IDispatcherChannelOpenTryIterator struct {
	Event *IDispatcherChannelOpenTry // Event containing the contract specifics and raw log

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
func (it *IDispatcherChannelOpenTryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherChannelOpenTry)
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
		it.Event = new(IDispatcherChannelOpenTry)
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
func (it *IDispatcherChannelOpenTryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherChannelOpenTryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherChannelOpenTry represents a ChannelOpenTry event raised by the IDispatcher contract.
type IDispatcherChannelOpenTry struct {
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
func (_IDispatcher *IDispatcherFilterer) FilterChannelOpenTry(opts *bind.FilterOpts, receiver []common.Address) (*IDispatcherChannelOpenTryIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "ChannelOpenTry", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherChannelOpenTryIterator{contract: _IDispatcher.contract, event: "ChannelOpenTry", logs: logs, sub: sub}, nil
}

// WatchChannelOpenTry is a free log subscription operation binding the contract event 0xf910705a7a768eb5958f281a5f84cae8bffc5dd811ca5cd303dda140a423698c.
//
// Solidity: event ChannelOpenTry(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId, bytes32 counterpartyChannelId)
func (_IDispatcher *IDispatcherFilterer) WatchChannelOpenTry(opts *bind.WatchOpts, sink chan<- *IDispatcherChannelOpenTry, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "ChannelOpenTry", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherChannelOpenTry)
				if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenTry", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseChannelOpenTry(log types.Log) (*IDispatcherChannelOpenTry, error) {
	event := new(IDispatcherChannelOpenTry)
	if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenTry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherChannelOpenTryErrorIterator is returned from FilterChannelOpenTryError and is used to iterate over the raw logs and unpacked data for ChannelOpenTryError events raised by the IDispatcher contract.
type IDispatcherChannelOpenTryErrorIterator struct {
	Event *IDispatcherChannelOpenTryError // Event containing the contract specifics and raw log

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
func (it *IDispatcherChannelOpenTryErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherChannelOpenTryError)
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
		it.Event = new(IDispatcherChannelOpenTryError)
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
func (it *IDispatcherChannelOpenTryErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherChannelOpenTryErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherChannelOpenTryError represents a ChannelOpenTryError event raised by the IDispatcher contract.
type IDispatcherChannelOpenTryError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenTryError is a free log retrieval operation binding the contract event 0x9e2fe55a3b54b57f82334c273f8d048cd7f05ad19c16cf334276a8c1fec4b6fd.
//
// Solidity: event ChannelOpenTryError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) FilterChannelOpenTryError(opts *bind.FilterOpts, receiver []common.Address) (*IDispatcherChannelOpenTryErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "ChannelOpenTryError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherChannelOpenTryErrorIterator{contract: _IDispatcher.contract, event: "ChannelOpenTryError", logs: logs, sub: sub}, nil
}

// WatchChannelOpenTryError is a free log subscription operation binding the contract event 0x9e2fe55a3b54b57f82334c273f8d048cd7f05ad19c16cf334276a8c1fec4b6fd.
//
// Solidity: event ChannelOpenTryError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) WatchChannelOpenTryError(opts *bind.WatchOpts, sink chan<- *IDispatcherChannelOpenTryError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "ChannelOpenTryError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherChannelOpenTryError)
				if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenTryError", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseChannelOpenTryError(log types.Log) (*IDispatcherChannelOpenTryError, error) {
	event := new(IDispatcherChannelOpenTryError)
	if err := _IDispatcher.contract.UnpackLog(event, "ChannelOpenTryError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherRecvPacketIterator is returned from FilterRecvPacket and is used to iterate over the raw logs and unpacked data for RecvPacket events raised by the IDispatcher contract.
type IDispatcherRecvPacketIterator struct {
	Event *IDispatcherRecvPacket // Event containing the contract specifics and raw log

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
func (it *IDispatcherRecvPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherRecvPacket)
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
		it.Event = new(IDispatcherRecvPacket)
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
func (it *IDispatcherRecvPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherRecvPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherRecvPacket represents a RecvPacket event raised by the IDispatcher contract.
type IDispatcherRecvPacket struct {
	DestPortAddress common.Address
	DestChannelId   [32]byte
	Sequence        uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRecvPacket is a free log retrieval operation binding the contract event 0xde5b57e6566d68a30b0979431df3d5df6db3b9aa89f8820f595b9315bf86d067.
//
// Solidity: event RecvPacket(address indexed destPortAddress, bytes32 indexed destChannelId, uint64 sequence)
func (_IDispatcher *IDispatcherFilterer) FilterRecvPacket(opts *bind.FilterOpts, destPortAddress []common.Address, destChannelId [][32]byte) (*IDispatcherRecvPacketIterator, error) {

	var destPortAddressRule []interface{}
	for _, destPortAddressItem := range destPortAddress {
		destPortAddressRule = append(destPortAddressRule, destPortAddressItem)
	}
	var destChannelIdRule []interface{}
	for _, destChannelIdItem := range destChannelId {
		destChannelIdRule = append(destChannelIdRule, destChannelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "RecvPacket", destPortAddressRule, destChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherRecvPacketIterator{contract: _IDispatcher.contract, event: "RecvPacket", logs: logs, sub: sub}, nil
}

// WatchRecvPacket is a free log subscription operation binding the contract event 0xde5b57e6566d68a30b0979431df3d5df6db3b9aa89f8820f595b9315bf86d067.
//
// Solidity: event RecvPacket(address indexed destPortAddress, bytes32 indexed destChannelId, uint64 sequence)
func (_IDispatcher *IDispatcherFilterer) WatchRecvPacket(opts *bind.WatchOpts, sink chan<- *IDispatcherRecvPacket, destPortAddress []common.Address, destChannelId [][32]byte) (event.Subscription, error) {

	var destPortAddressRule []interface{}
	for _, destPortAddressItem := range destPortAddress {
		destPortAddressRule = append(destPortAddressRule, destPortAddressItem)
	}
	var destChannelIdRule []interface{}
	for _, destChannelIdItem := range destChannelId {
		destChannelIdRule = append(destChannelIdRule, destChannelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "RecvPacket", destPortAddressRule, destChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherRecvPacket)
				if err := _IDispatcher.contract.UnpackLog(event, "RecvPacket", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseRecvPacket(log types.Log) (*IDispatcherRecvPacket, error) {
	event := new(IDispatcherRecvPacket)
	if err := _IDispatcher.contract.UnpackLog(event, "RecvPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherSendPacketIterator is returned from FilterSendPacket and is used to iterate over the raw logs and unpacked data for SendPacket events raised by the IDispatcher contract.
type IDispatcherSendPacketIterator struct {
	Event *IDispatcherSendPacket // Event containing the contract specifics and raw log

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
func (it *IDispatcherSendPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherSendPacket)
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
		it.Event = new(IDispatcherSendPacket)
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
func (it *IDispatcherSendPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherSendPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherSendPacket represents a SendPacket event raised by the IDispatcher contract.
type IDispatcherSendPacket struct {
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
func (_IDispatcher *IDispatcherFilterer) FilterSendPacket(opts *bind.FilterOpts, sourcePortAddress []common.Address, sourceChannelId [][32]byte) (*IDispatcherSendPacketIterator, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "SendPacket", sourcePortAddressRule, sourceChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherSendPacketIterator{contract: _IDispatcher.contract, event: "SendPacket", logs: logs, sub: sub}, nil
}

// WatchSendPacket is a free log subscription operation binding the contract event 0xb5bff96e18da044e4e34510d16df9053b9f1920f6a960732e5aaf22fe9b80136.
//
// Solidity: event SendPacket(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, bytes packet, uint64 sequence, uint64 timeoutTimestamp)
func (_IDispatcher *IDispatcherFilterer) WatchSendPacket(opts *bind.WatchOpts, sink chan<- *IDispatcherSendPacket, sourcePortAddress []common.Address, sourceChannelId [][32]byte) (event.Subscription, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "SendPacket", sourcePortAddressRule, sourceChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherSendPacket)
				if err := _IDispatcher.contract.UnpackLog(event, "SendPacket", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseSendPacket(log types.Log) (*IDispatcherSendPacket, error) {
	event := new(IDispatcherSendPacket)
	if err := _IDispatcher.contract.UnpackLog(event, "SendPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherTimeoutIterator is returned from FilterTimeout and is used to iterate over the raw logs and unpacked data for Timeout events raised by the IDispatcher contract.
type IDispatcherTimeoutIterator struct {
	Event *IDispatcherTimeout // Event containing the contract specifics and raw log

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
func (it *IDispatcherTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherTimeout)
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
		it.Event = new(IDispatcherTimeout)
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
func (it *IDispatcherTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherTimeout represents a Timeout event raised by the IDispatcher contract.
type IDispatcherTimeout struct {
	SourcePortAddress common.Address
	SourceChannelId   [32]byte
	Sequence          uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTimeout is a free log retrieval operation binding the contract event 0x19ac40c4084d9bfb5b43f819a94bf01c70789b0d579871f59e4f86def04d9344.
//
// Solidity: event Timeout(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 indexed sequence)
func (_IDispatcher *IDispatcherFilterer) FilterTimeout(opts *bind.FilterOpts, sourcePortAddress []common.Address, sourceChannelId [][32]byte, sequence []uint64) (*IDispatcherTimeoutIterator, error) {

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

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "Timeout", sourcePortAddressRule, sourceChannelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherTimeoutIterator{contract: _IDispatcher.contract, event: "Timeout", logs: logs, sub: sub}, nil
}

// WatchTimeout is a free log subscription operation binding the contract event 0x19ac40c4084d9bfb5b43f819a94bf01c70789b0d579871f59e4f86def04d9344.
//
// Solidity: event Timeout(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 indexed sequence)
func (_IDispatcher *IDispatcherFilterer) WatchTimeout(opts *bind.WatchOpts, sink chan<- *IDispatcherTimeout, sourcePortAddress []common.Address, sourceChannelId [][32]byte, sequence []uint64) (event.Subscription, error) {

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

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "Timeout", sourcePortAddressRule, sourceChannelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherTimeout)
				if err := _IDispatcher.contract.UnpackLog(event, "Timeout", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseTimeout(log types.Log) (*IDispatcherTimeout, error) {
	event := new(IDispatcherTimeout)
	if err := _IDispatcher.contract.UnpackLog(event, "Timeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherTimeoutErrorIterator is returned from FilterTimeoutError and is used to iterate over the raw logs and unpacked data for TimeoutError events raised by the IDispatcher contract.
type IDispatcherTimeoutErrorIterator struct {
	Event *IDispatcherTimeoutError // Event containing the contract specifics and raw log

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
func (it *IDispatcherTimeoutErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherTimeoutError)
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
		it.Event = new(IDispatcherTimeoutError)
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
func (it *IDispatcherTimeoutErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherTimeoutErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherTimeoutError represents a TimeoutError event raised by the IDispatcher contract.
type IDispatcherTimeoutError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTimeoutError is a free log retrieval operation binding the contract event 0x83adb31803bee4e18cda1d04a781d77f6f271718a61b25e3a06f319b5103a330.
//
// Solidity: event TimeoutError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) FilterTimeoutError(opts *bind.FilterOpts, receiver []common.Address) (*IDispatcherTimeoutErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "TimeoutError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherTimeoutErrorIterator{contract: _IDispatcher.contract, event: "TimeoutError", logs: logs, sub: sub}, nil
}

// WatchTimeoutError is a free log subscription operation binding the contract event 0x83adb31803bee4e18cda1d04a781d77f6f271718a61b25e3a06f319b5103a330.
//
// Solidity: event TimeoutError(address indexed receiver, bytes error)
func (_IDispatcher *IDispatcherFilterer) WatchTimeoutError(opts *bind.WatchOpts, sink chan<- *IDispatcherTimeoutError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "TimeoutError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherTimeoutError)
				if err := _IDispatcher.contract.UnpackLog(event, "TimeoutError", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseTimeoutError(log types.Log) (*IDispatcherTimeoutError, error) {
	event := new(IDispatcherTimeoutError)
	if err := _IDispatcher.contract.UnpackLog(event, "TimeoutError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherWriteAckPacketIterator is returned from FilterWriteAckPacket and is used to iterate over the raw logs and unpacked data for WriteAckPacket events raised by the IDispatcher contract.
type IDispatcherWriteAckPacketIterator struct {
	Event *IDispatcherWriteAckPacket // Event containing the contract specifics and raw log

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
func (it *IDispatcherWriteAckPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherWriteAckPacket)
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
		it.Event = new(IDispatcherWriteAckPacket)
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
func (it *IDispatcherWriteAckPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherWriteAckPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherWriteAckPacket represents a WriteAckPacket event raised by the IDispatcher contract.
type IDispatcherWriteAckPacket struct {
	WriterPortAddress common.Address
	WriterChannelId   [32]byte
	Sequence          uint64
	AckPacket         AckPacket
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWriteAckPacket is a free log retrieval operation binding the contract event 0xa32e6f42b1d63fb83ad73b009a6dbb9413d1da02e95b1bb08f081815eea8db20.
//
// Solidity: event WriteAckPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (bool,bytes) ackPacket)
func (_IDispatcher *IDispatcherFilterer) FilterWriteAckPacket(opts *bind.FilterOpts, writerPortAddress []common.Address, writerChannelId [][32]byte) (*IDispatcherWriteAckPacketIterator, error) {

	var writerPortAddressRule []interface{}
	for _, writerPortAddressItem := range writerPortAddress {
		writerPortAddressRule = append(writerPortAddressRule, writerPortAddressItem)
	}
	var writerChannelIdRule []interface{}
	for _, writerChannelIdItem := range writerChannelId {
		writerChannelIdRule = append(writerChannelIdRule, writerChannelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "WriteAckPacket", writerPortAddressRule, writerChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherWriteAckPacketIterator{contract: _IDispatcher.contract, event: "WriteAckPacket", logs: logs, sub: sub}, nil
}

// WatchWriteAckPacket is a free log subscription operation binding the contract event 0xa32e6f42b1d63fb83ad73b009a6dbb9413d1da02e95b1bb08f081815eea8db20.
//
// Solidity: event WriteAckPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (bool,bytes) ackPacket)
func (_IDispatcher *IDispatcherFilterer) WatchWriteAckPacket(opts *bind.WatchOpts, sink chan<- *IDispatcherWriteAckPacket, writerPortAddress []common.Address, writerChannelId [][32]byte) (event.Subscription, error) {

	var writerPortAddressRule []interface{}
	for _, writerPortAddressItem := range writerPortAddress {
		writerPortAddressRule = append(writerPortAddressRule, writerPortAddressItem)
	}
	var writerChannelIdRule []interface{}
	for _, writerChannelIdItem := range writerChannelId {
		writerChannelIdRule = append(writerChannelIdRule, writerChannelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "WriteAckPacket", writerPortAddressRule, writerChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherWriteAckPacket)
				if err := _IDispatcher.contract.UnpackLog(event, "WriteAckPacket", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseWriteAckPacket(log types.Log) (*IDispatcherWriteAckPacket, error) {
	event := new(IDispatcherWriteAckPacket)
	if err := _IDispatcher.contract.UnpackLog(event, "WriteAckPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDispatcherWriteTimeoutPacketIterator is returned from FilterWriteTimeoutPacket and is used to iterate over the raw logs and unpacked data for WriteTimeoutPacket events raised by the IDispatcher contract.
type IDispatcherWriteTimeoutPacketIterator struct {
	Event *IDispatcherWriteTimeoutPacket // Event containing the contract specifics and raw log

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
func (it *IDispatcherWriteTimeoutPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDispatcherWriteTimeoutPacket)
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
		it.Event = new(IDispatcherWriteTimeoutPacket)
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
func (it *IDispatcherWriteTimeoutPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDispatcherWriteTimeoutPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDispatcherWriteTimeoutPacket represents a WriteTimeoutPacket event raised by the IDispatcher contract.
type IDispatcherWriteTimeoutPacket struct {
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
func (_IDispatcher *IDispatcherFilterer) FilterWriteTimeoutPacket(opts *bind.FilterOpts, writerPortAddress []common.Address, writerChannelId [][32]byte) (*IDispatcherWriteTimeoutPacketIterator, error) {

	var writerPortAddressRule []interface{}
	for _, writerPortAddressItem := range writerPortAddress {
		writerPortAddressRule = append(writerPortAddressRule, writerPortAddressItem)
	}
	var writerChannelIdRule []interface{}
	for _, writerChannelIdItem := range writerChannelId {
		writerChannelIdRule = append(writerChannelIdRule, writerChannelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.FilterLogs(opts, "WriteTimeoutPacket", writerPortAddressRule, writerChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &IDispatcherWriteTimeoutPacketIterator{contract: _IDispatcher.contract, event: "WriteTimeoutPacket", logs: logs, sub: sub}, nil
}

// WatchWriteTimeoutPacket is a free log subscription operation binding the contract event 0xedbcd9eeb09d85c3ea1b5bf002c04478059cb261cab82c903885cefccae374bc.
//
// Solidity: event WriteTimeoutPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_IDispatcher *IDispatcherFilterer) WatchWriteTimeoutPacket(opts *bind.WatchOpts, sink chan<- *IDispatcherWriteTimeoutPacket, writerPortAddress []common.Address, writerChannelId [][32]byte) (event.Subscription, error) {

	var writerPortAddressRule []interface{}
	for _, writerPortAddressItem := range writerPortAddress {
		writerPortAddressRule = append(writerPortAddressRule, writerPortAddressItem)
	}
	var writerChannelIdRule []interface{}
	for _, writerChannelIdItem := range writerChannelId {
		writerChannelIdRule = append(writerChannelIdRule, writerChannelIdItem)
	}

	logs, sub, err := _IDispatcher.contract.WatchLogs(opts, "WriteTimeoutPacket", writerPortAddressRule, writerChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDispatcherWriteTimeoutPacket)
				if err := _IDispatcher.contract.UnpackLog(event, "WriteTimeoutPacket", log); err != nil {
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
func (_IDispatcher *IDispatcherFilterer) ParseWriteTimeoutPacket(log types.Log) (*IDispatcherWriteTimeoutPacket, error) {
	event := new(IDispatcherWriteTimeoutPacket)
	if err := _IDispatcher.contract.UnpackLog(event, "WriteTimeoutPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
