/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Contract,
  ContractFactory,
  ContractTransactionResponse,
  Interface,
} from "ethers";
import type {
  Signer,
  AddressLike,
  ContractDeployTransaction,
  ContractRunner,
} from "ethers";
import type { NonPayableOverrides } from "../../common";
import type {
  RevertingBytesMars,
  RevertingBytesMarsInterface,
} from "../../Mars.sol/RevertingBytesMars";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "_dispatcher",
        type: "address",
        internalType: "contract IbcDispatcher",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "receive",
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "ackPackets",
    inputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "success",
        type: "bool",
        internalType: "bool",
      },
      {
        name: "data",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "connectedChannels",
    inputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "dispatcher",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "contract IbcDispatcher",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "greet",
    inputs: [
      {
        name: "message",
        type: "string",
        internalType: "string",
      },
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "timeoutTimestamp",
        type: "uint64",
        internalType: "uint64",
      },
    ],
    outputs: [
      {
        name: "sequence",
        type: "uint64",
        internalType: "uint64",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "greetWithFee",
    inputs: [
      {
        name: "message",
        type: "string",
        internalType: "string",
      },
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "timeoutTimestamp",
        type: "uint64",
        internalType: "uint64",
      },
      {
        name: "gasLimits",
        type: "uint256[2]",
        internalType: "uint256[2]",
      },
      {
        name: "gasPrices",
        type: "uint256[2]",
        internalType: "uint256[2]",
      },
    ],
    outputs: [
      {
        name: "sequence",
        type: "uint64",
        internalType: "uint64",
      },
    ],
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "onAcknowledgementPacket",
    inputs: [
      {
        name: "",
        type: "tuple",
        internalType: "struct IbcPacket",
        components: [
          {
            name: "src",
            type: "tuple",
            internalType: "struct IbcEndpoint",
            components: [
              {
                name: "portId",
                type: "string",
                internalType: "string",
              },
              {
                name: "channelId",
                type: "bytes32",
                internalType: "bytes32",
              },
            ],
          },
          {
            name: "dest",
            type: "tuple",
            internalType: "struct IbcEndpoint",
            components: [
              {
                name: "portId",
                type: "string",
                internalType: "string",
              },
              {
                name: "channelId",
                type: "bytes32",
                internalType: "bytes32",
              },
            ],
          },
          {
            name: "sequence",
            type: "uint64",
            internalType: "uint64",
          },
          {
            name: "data",
            type: "bytes",
            internalType: "bytes",
          },
          {
            name: "timeoutHeight",
            type: "tuple",
            internalType: "struct Height",
            components: [
              {
                name: "revision_number",
                type: "uint64",
                internalType: "uint64",
              },
              {
                name: "revision_height",
                type: "uint64",
                internalType: "uint64",
              },
            ],
          },
          {
            name: "timeoutTimestamp",
            type: "uint64",
            internalType: "uint64",
          },
        ],
      },
      {
        name: "ack",
        type: "tuple",
        internalType: "struct AckPacket",
        components: [
          {
            name: "success",
            type: "bool",
            internalType: "bool",
          },
          {
            name: "data",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "onChanCloseConfirm",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "",
        type: "string",
        internalType: "string",
      },
      {
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "onChanCloseInit",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "",
        type: "string",
        internalType: "string",
      },
      {
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "onChanOpenAck",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "counterpartyVersion",
        type: "string",
        internalType: "string",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "onChanOpenConfirm",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "onChanOpenTry",
    inputs: [
      {
        name: "",
        type: "uint8",
        internalType: "enum ChannelOrder",
      },
      {
        name: "",
        type: "string[]",
        internalType: "string[]",
      },
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "",
        type: "string",
        internalType: "string",
      },
      {
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "counterpartyVersion",
        type: "string",
        internalType: "string",
      },
    ],
    outputs: [
      {
        name: "selectedVersion",
        type: "string",
        internalType: "string",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "onRecvPacket",
    inputs: [
      {
        name: "",
        type: "tuple",
        internalType: "struct IbcPacket",
        components: [
          {
            name: "src",
            type: "tuple",
            internalType: "struct IbcEndpoint",
            components: [
              {
                name: "portId",
                type: "string",
                internalType: "string",
              },
              {
                name: "channelId",
                type: "bytes32",
                internalType: "bytes32",
              },
            ],
          },
          {
            name: "dest",
            type: "tuple",
            internalType: "struct IbcEndpoint",
            components: [
              {
                name: "portId",
                type: "string",
                internalType: "string",
              },
              {
                name: "channelId",
                type: "bytes32",
                internalType: "bytes32",
              },
            ],
          },
          {
            name: "sequence",
            type: "uint64",
            internalType: "uint64",
          },
          {
            name: "data",
            type: "bytes",
            internalType: "bytes",
          },
          {
            name: "timeoutHeight",
            type: "tuple",
            internalType: "struct Height",
            components: [
              {
                name: "revision_number",
                type: "uint64",
                internalType: "uint64",
              },
              {
                name: "revision_height",
                type: "uint64",
                internalType: "uint64",
              },
            ],
          },
          {
            name: "timeoutTimestamp",
            type: "uint64",
            internalType: "uint64",
          },
        ],
      },
    ],
    outputs: [
      {
        name: "ack",
        type: "tuple",
        internalType: "struct AckPacket",
        components: [
          {
            name: "success",
            type: "bool",
            internalType: "bool",
          },
          {
            name: "data",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
      {
        name: "skipAck",
        type: "bool",
        internalType: "bool",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "onTimeoutPacket",
    inputs: [
      {
        name: "",
        type: "tuple",
        internalType: "struct IbcPacket",
        components: [
          {
            name: "src",
            type: "tuple",
            internalType: "struct IbcEndpoint",
            components: [
              {
                name: "portId",
                type: "string",
                internalType: "string",
              },
              {
                name: "channelId",
                type: "bytes32",
                internalType: "bytes32",
              },
            ],
          },
          {
            name: "dest",
            type: "tuple",
            internalType: "struct IbcEndpoint",
            components: [
              {
                name: "portId",
                type: "string",
                internalType: "string",
              },
              {
                name: "channelId",
                type: "bytes32",
                internalType: "bytes32",
              },
            ],
          },
          {
            name: "sequence",
            type: "uint64",
            internalType: "uint64",
          },
          {
            name: "data",
            type: "bytes",
            internalType: "bytes",
          },
          {
            name: "timeoutHeight",
            type: "tuple",
            internalType: "struct Height",
            components: [
              {
                name: "revision_number",
                type: "uint64",
                internalType: "uint64",
              },
              {
                name: "revision_height",
                type: "uint64",
                internalType: "uint64",
              },
            ],
          },
          {
            name: "timeoutTimestamp",
            type: "uint64",
            internalType: "uint64",
          },
        ],
      },
    ],
    outputs: [],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "owner",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "recvedPackets",
    inputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "src",
        type: "tuple",
        internalType: "struct IbcEndpoint",
        components: [
          {
            name: "portId",
            type: "string",
            internalType: "string",
          },
          {
            name: "channelId",
            type: "bytes32",
            internalType: "bytes32",
          },
        ],
      },
      {
        name: "dest",
        type: "tuple",
        internalType: "struct IbcEndpoint",
        components: [
          {
            name: "portId",
            type: "string",
            internalType: "string",
          },
          {
            name: "channelId",
            type: "bytes32",
            internalType: "bytes32",
          },
        ],
      },
      {
        name: "sequence",
        type: "uint64",
        internalType: "uint64",
      },
      {
        name: "data",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "timeoutHeight",
        type: "tuple",
        internalType: "struct Height",
        components: [
          {
            name: "revision_number",
            type: "uint64",
            internalType: "uint64",
          },
          {
            name: "revision_height",
            type: "uint64",
            internalType: "uint64",
          },
        ],
      },
      {
        name: "timeoutTimestamp",
        type: "uint64",
        internalType: "uint64",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "renounceOwnership",
    inputs: [],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "supportedVersions",
    inputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "",
        type: "string",
        internalType: "string",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "timeoutPackets",
    inputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "src",
        type: "tuple",
        internalType: "struct IbcEndpoint",
        components: [
          {
            name: "portId",
            type: "string",
            internalType: "string",
          },
          {
            name: "channelId",
            type: "bytes32",
            internalType: "bytes32",
          },
        ],
      },
      {
        name: "dest",
        type: "tuple",
        internalType: "struct IbcEndpoint",
        components: [
          {
            name: "portId",
            type: "string",
            internalType: "string",
          },
          {
            name: "channelId",
            type: "bytes32",
            internalType: "bytes32",
          },
        ],
      },
      {
        name: "sequence",
        type: "uint64",
        internalType: "uint64",
      },
      {
        name: "data",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "timeoutHeight",
        type: "tuple",
        internalType: "struct Height",
        components: [
          {
            name: "revision_number",
            type: "uint64",
            internalType: "uint64",
          },
          {
            name: "revision_height",
            type: "uint64",
            internalType: "uint64",
          },
        ],
      },
      {
        name: "timeoutTimestamp",
        type: "uint64",
        internalType: "uint64",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "transferOwnership",
    inputs: [
      {
        name: "newOwner",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "triggerChannelClose",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "triggerChannelInit",
    inputs: [
      {
        name: "version",
        type: "string",
        internalType: "string",
      },
      {
        name: "ordering",
        type: "uint8",
        internalType: "enum ChannelOrder",
      },
      {
        name: "feeEnabled",
        type: "bool",
        internalType: "bool",
      },
      {
        name: "connectionHops",
        type: "string[]",
        internalType: "string[]",
      },
      {
        name: "counterpartyPortId",
        type: "string",
        internalType: "string",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "triggerChannelInitWithFee",
    inputs: [
      {
        name: "version",
        type: "string",
        internalType: "string",
      },
      {
        name: "ordering",
        type: "uint8",
        internalType: "enum ChannelOrder",
      },
      {
        name: "feeEnabled",
        type: "bool",
        internalType: "bool",
      },
      {
        name: "connectionHops",
        type: "string[]",
        internalType: "string[]",
      },
      {
        name: "counterpartyPortId",
        type: "string",
        internalType: "string",
      },
    ],
    outputs: [],
    stateMutability: "payable",
  },
  {
    type: "event",
    name: "OwnershipTransferred",
    inputs: [
      {
        name: "previousOwner",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "newOwner",
        type: "address",
        indexed: true,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "error",
    name: "ChannelNotFound",
    inputs: [],
  },
  {
    type: "error",
    name: "OnRecvPacketRevert",
    inputs: [],
  },
  {
    type: "error",
    name: "OnTimeoutPacket",
    inputs: [],
  },
  {
    type: "error",
    name: "UnsupportedVersion",
    inputs: [],
  },
  {
    type: "error",
    name: "notIbcDispatcher",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x600360c0818152620312e360ec1b60e0526080908152610140604052610100918252620322e360ec1b6101205260a09190915262000042906006906002620000f9565b503480156200005057600080fd5b50604051620023b8380380620023b88339810160408190526200007391620001d0565b80806200008033620000a9565b600180546001600160a01b0319166001600160a01b039290921691909117905550620003739050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b82805482825590600052602060002090810192821562000144579160200282015b82811115620001445782518290620001339082620002a7565b50916020019190600101906200011a565b506200015292915062000156565b5090565b80821115620001525760006200016d828262000177565b5060010162000156565b508054620001859062000218565b6000825580601f1062000196575050565b601f016020900490600052602060002090810190620001b69190620001b9565b50565b5b80821115620001525760008155600101620001ba565b600060208284031215620001e357600080fd5b81516001600160a01b0381168114620001fb57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200022d57607f821691505b6020821081036200024e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620002a257600081815260208120601f850160051c810160208610156200027d5750805b601f850160051c820191505b818110156200029e5782815560010162000289565b5050505b505050565b81516001600160401b03811115620002c357620002c362000202565b620002db81620002d4845462000218565b8462000254565b602080601f831160018114620003135760008415620002fa5750858301515b600019600386901b1c1916600185901b1785556200029e565b600085815260208120601f198616915b82811015620003445788860151825594840194600190910190840162000323565b5085821015620003635787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61203580620003836000396000f3fe6080604052600436106101395760003560e01c8063715018a6116100ab578063bb3f9f8d1161006f578063bb3f9f8d1461038c578063cb7e9057146103ba578063e847e280146103da578063f12b758a146103fa578063f2fde38b1461041a578063fad28a241461043a57600080fd5b8063715018a6146102e55780637a805598146102fa5780637d6221841461031a5780637e1d42b51461033a5780638da5cb5b1461035a57600080fd5b80634dcc0aa6116100fd5780634dcc0aa6146102125780634eeb739114610240578063558850ac146102725780635bfd12b814610292578063602f9834146102b257806361995001146102d257600080fd5b80631eb7dd5e146101455780633513a995146101675780633f9fdbe4146101975780634252ae9b146101b75780634bdb5597146101e557600080fd5b3661014057005b600080fd5b34801561015157600080fd5b50610165610160366004611212565b61045a565b005b61017a610175366004611388565b61048b565b6040516001600160401b0390911681526020015b60405180910390f35b3480156101a357600080fd5b506101656101b2366004611212565b61052b565b3480156101c357600080fd5b506101d76101d2366004611408565b6105e7565b60405161018e92919061146e565b3480156101f157600080fd5b5061020561020036600461150f565b6106a3565b60405161018e9190611648565b34801561021e57600080fd5b5061023261022d3660046116ee565b6106e8565b60405161018e9291906117d6565b34801561024c57600080fd5b5061026061025b366004611408565b610766565b60405161018e96959493929190611839565b34801561027e57600080fd5b5061016561028d366004611408565b6109c3565b34801561029e57600080fd5b5061017a6102ad3660046118ae565b610a25565b3480156102be57600080fd5b506101656102cd366004611924565b610aa8565b6101656102e0366004611966565b610aec565b3480156102f157600080fd5b50610165610bb8565b34801561030657600080fd5b50610165610315366004611966565b610bcc565b34801561032657600080fd5b50610205610335366004611408565b610c4e565b34801561034657600080fd5b50610165610355366004611a58565b610cfa565b34801561036657600080fd5b506000546001600160a01b03165b6040516001600160a01b03909116815260200161018e565b34801561039857600080fd5b506103ac6103a7366004611408565b610d66565b60405190815260200161018e565b3480156103c657600080fd5b50600154610374906001600160a01b031681565b3480156103e657600080fd5b506101656103f5366004611ac2565b610d87565b34801561040657600080fd5b50610260610415366004611408565b610dbd565b34801561042657600080fd5b50610165610435366004611b29565b610dcd565b34801561044657600080fd5b50610165610455366004611408565b610e4b565b6001546001600160a01b03163314610485576040516321bf7f4960e01b815260040160405180910390fd5b50505050565b6001546040516330f8455760e21b81526000916001600160a01b03169063c3e1155c906104c29088908b908b908a90600401611b6f565b6020604051808303816000875af11580156104e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105059190611ba3565b600154909150610521906001600160a01b031686838686610e76565b9695505050505050565b6001546001600160a01b03163314610556576040516321bf7f4960e01b815260040160405180910390fd5b6000805b6005548110156105c157856005828154811061057857610578611bc0565b9060005260206000200154036105af576005818154811061059b5761059b611bc0565b6000918252602082200155600191506105c1565b806105b981611bd6565b91505061055a565b50806105e057604051630781f76560e21b815260040160405180910390fd5b5050505050565b600381815481106105f757600080fd5b60009182526020909120600290910201805460018201805460ff90921693509061062090611bfd565b80601f016020809104026020016040519081016040528092919081815260200182805461064c90611bfd565b80156106995780601f1061066e57610100808354040283529160200191610699565b820191906000526020600020905b81548152906001019060200180831161067c57829003601f168201915b5050505050905082565b6001546060906001600160a01b031633146106d1576040516321bf7f4960e01b815260040160405180910390fd5b6106dc868484610f45565b98975050505050505050565b6040805180820190915260008152606060208201526001546000906001600160a01b0316331461072b576040516321bf7f4960e01b815260040160405180910390fd5b505060408051808201825260008082528251602081810185528282528301528251639889d82160e01b81529251919290919081900360040190fd5b6004818154811061077657600080fd5b9060005260206000209060080201600091509050806000016040518060400160405290816000820180546107a990611bfd565b80601f01602080910402602001604051908101604052809291908181526020018280546107d590611bfd565b80156108225780601f106107f757610100808354040283529160200191610822565b820191906000526020600020905b81548152906001019060200180831161080557829003601f168201915b50505050508152602001600182015481525050908060020160405180604001604052908160008201805461085590611bfd565b80601f016020809104026020016040519081016040528092919081815260200182805461088190611bfd565b80156108ce5780601f106108a3576101008083540402835291602001916108ce565b820191906000526020600020905b8154815290600101906020018083116108b157829003601f168201915b505050918352505060019190910154602090910152600482015460058301805492936001600160401b039092169261090590611bfd565b80601f016020809104026020016040519081016040528092919081815260200182805461093190611bfd565b801561097e5780601f106109535761010080835404028352916020019161097e565b820191906000526020600020905b81548152906001019060200180831161096157829003601f168201915b50506040805180820190915260068601546001600160401b03808216835268010000000000000000909104811660208301526007909601549495909416925088915050565b6109cb61106b565b6001546040516381bc079b60e01b8152600481018390526001600160a01b03909116906381bc079b90602401600060405180830381600087803b158015610a1157600080fd5b505af11580156105e0573d6000803e3d6000fd5b6001546040516330f8455760e21b81526000916001600160a01b03169063c3e1155c90610a5c908690899089908890600401611b6f565b6020604051808303816000875af1158015610a7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a9f9190611ba3565b95945050505050565b6001546001600160a01b03163314610ad3576040516321bf7f4960e01b815260040160405180910390fd5b604051631021bb3b60e31b815260040160405180910390fd5b610af461106b565b60015460405163418925b760e01b81526001600160a01b0390911690819063418925b790610b34908c908c908c908c908c908c908c908c90600401611ce4565b600060405180830381600087803b158015610b4e57600080fd5b505af1158015610b62573d6000803e3d6000fd5b50505050610bad818a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c92508a915089905088886110c5565b505050505050505050565b610bc061106b565b610bca600061117a565b565b610bd461106b565b60015460405163418925b760e01b81526001600160a01b039091169063418925b790610c12908b908b908b908b908b908b908b908b90600401611ce4565b600060405180830381600087803b158015610c2c57600080fd5b505af1158015610c40573d6000803e3d6000fd5b505050505050505050505050565b60068181548110610c5e57600080fd5b906000526020600020016000915090508054610c7990611bfd565b80601f0160208091040260200160405190810160405280929190818152602001828054610ca590611bfd565b8015610cf25780601f10610cc757610100808354040283529160200191610cf2565b820191906000526020600020905b815481529060010190602001808311610cd557829003601f168201915b505050505081565b6001546001600160a01b03163314610d25576040516321bf7f4960e01b815260040160405180910390fd5b6003805460018101825560009190915281906002027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b016104858282611d93565b60058181548110610d7657600080fd5b600091825260209091200154905081565b6001546001600160a01b03163314610db2576040516321bf7f4960e01b815260040160405180910390fd5b6105e0848383610f45565b6002818154811061077657600080fd5b610dd561106b565b6001600160a01b038116610e3f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b610e488161117a565b50565b6001546001600160a01b03163314610e48576040516321bf7f4960e01b815260040160405180910390fd5b846001600160a01b031663478222c26040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610eb6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eda9190611ea2565b6001600160a01b03166318e3404b34868686866040518663ffffffff1660e01b8152600401610f0c9493929190611ee2565b6000604051808303818588803b158015610f2557600080fd5b505af1158015610f39573d6000803e3d6000fd5b50505050505050505050565b606060005b60065481101561104a5760068181548110610f6757610f67611bc0565b90600052602060002001604051602001610f819190611f12565b604051602081830303815290604052805190602001208484604051602001610faa929190611f88565b604051602081830303815290604052805190602001200361103857600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001859055604080516020601f86018190048102820181019092528481529085908590819084018382808284376000920191909152509294506110649350505050565b8061104281611bd6565b915050610f4a565b5060405163b01318a560e01b815260040160405180910390fd5b9392505050565b6000546001600160a01b03163314610bca5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610e36565b866001600160a01b031663478222c26040518163ffffffff1660e01b81526004016020604051808303816000875af1158015611105573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111299190611ea2565b6001600160a01b031663fce34e4034308989898989896040518963ffffffff1660e01b81526004016111619796959493929190611f98565b6000604051808303818588803b158015610c2c57600080fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60008083601f8401126111dc57600080fd5b5081356001600160401b038111156111f357600080fd5b60208301915083602082850101111561120b57600080fd5b9250929050565b6000806000806060858703121561122857600080fd5b8435935060208501356001600160401b0381111561124557600080fd5b611251878288016111ca565b9598909750949560400135949350505050565b6001600160401b0381168114610e4857600080fd5b803561128481611264565b919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156112c1576112c1611289565b60405290565b60405160c081016001600160401b03811182821017156112c1576112c1611289565b604051601f8201601f191681016001600160401b038111828210171561131157611311611289565b604052919050565b600082601f83011261132a57600080fd5b604051604081018181106001600160401b038211171561134c5761134c611289565b806040525080604084018581111561136357600080fd5b845b8181101561137d578035835260209283019201611365565b509195945050505050565b60008060008060008060e087890312156113a157600080fd5b86356001600160401b038111156113b757600080fd5b6113c389828a016111ca565b9097509550506020870135935060408701356113de81611264565b92506113ed8860608901611319565b91506113fc8860a08901611319565b90509295509295509295565b60006020828403121561141a57600080fd5b5035919050565b6000815180845260005b818110156114475760208185018101518683018201520161142b565b81811115611459576000602083870101525b50601f01601f19169290920160200192915050565b82151581526040602082015260006114896040830184611421565b949350505050565b80356003811061128457600080fd5b600082601f8301126114b157600080fd5b81356001600160401b038111156114ca576114ca611289565b6114dd601f8201601f19166020016112e9565b8181528460208386010111156114f257600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600080600060c0888a03121561152a57600080fd5b61153388611491565b965060208801356001600160401b038082111561154f57600080fd5b818a0191508a601f83011261156357600080fd5b81358181111561157557611575611289565b8060051b611585602082016112e9565b9182526020818501810192908101908e8411156115a157600080fd5b6020860192505b838310156115df5784833511156115be57600080fd5b6115ce8f602085358901016114a0565b8252602092830192909101906115a8565b9a5050505060408a0135965060608a01359150808211156115ff57600080fd5b61160b8b838c016114a0565b955060808a0135945060a08a013591508082111561162857600080fd5b506116358a828b016111ca565b989b979a50959850939692959293505050565b6020815260006110646020830184611421565b60006040828403121561166d57600080fd5b61167561129f565b905081356001600160401b0381111561168d57600080fd5b611699848285016114a0565b8252506020820135602082015292915050565b6000604082840312156116be57600080fd5b6116c661129f565b905081356116d381611264565b815260208201356116e381611264565b602082015292915050565b60006020828403121561170057600080fd5b81356001600160401b038082111561171757600080fd5b9083019060e0828603121561172b57600080fd5b6117336112c7565b82358281111561174257600080fd5b61174e8782860161165b565b82525060208301358281111561176357600080fd5b61176f8782860161165b565b60208301525061178160408401611279565b604082015260608301358281111561179857600080fd5b6117a4878286016114a0565b6060830152506117b786608085016116ac565b60808201526117c860c08401611279565b60a082015295945050505050565b6040815282511515604082015260006020840151604060608401526117fe6080840182611421565b91505082151560208301529392505050565b60008151604084526118256040850182611421565b602093840151949093019390935250919050565b60e08152600061184c60e0830189611810565b828103602084015261185e8189611810565b90506001600160401b03808816604085015283820360608501526118828288611421565b92508086511660808501528060208701511660a085015280851660c08501525050979650505050505050565b600080600080606085870312156118c457600080fd5b84356001600160401b038111156118da57600080fd5b6118e6878288016111ca565b90955093505060208501359150604085013561190181611264565b939692955090935050565b600060e0828403121561191e57600080fd5b50919050565b60006020828403121561193657600080fd5b81356001600160401b0381111561194c57600080fd5b6114898482850161190c565b8015158114610e4857600080fd5b60008060008060008060008060a0898b03121561198257600080fd5b88356001600160401b038082111561199957600080fd5b6119a58c838d016111ca565b909a5098508891506119b960208c01611491565b975060408b013591506119cb82611958565b90955060608a013590808211156119e157600080fd5b818b0191508b601f8301126119f557600080fd5b813581811115611a0457600080fd5b8c60208260051b8501011115611a1957600080fd5b6020830196508095505060808b0135915080821115611a3757600080fd5b50611a448b828c016111ca565b999c989b5096995094979396929594505050565b60008060408385031215611a6b57600080fd5b82356001600160401b0380821115611a8257600080fd5b611a8e8683870161190c565b93506020850135915080821115611aa457600080fd5b50830160408186031215611ab757600080fd5b809150509250929050565b60008060008060608587031215611ad857600080fd5b843593506020850135925060408501356001600160401b03811115611afc57600080fd5b611b08878288016111ca565b95989497509550505050565b6001600160a01b0381168114610e4857600080fd5b600060208284031215611b3b57600080fd5b813561106481611b14565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b848152606060208201526000611b89606083018587611b46565b90506001600160401b038316604083015295945050505050565b600060208284031215611bb557600080fd5b815161106481611264565b634e487b7160e01b600052603260045260246000fd5b600060018201611bf657634e487b7160e01b600052601160045260246000fd5b5060010190565b600181811c90821680611c1157607f821691505b60208210810361191e57634e487b7160e01b600052602260045260246000fd5b60038110611c4f57634e487b7160e01b600052602160045260246000fd5b9052565b81835260006020808501808196508560051b810191508460005b87811015611cd75782840389528135601e19883603018112611c8e57600080fd5b870185810190356001600160401b03811115611ca957600080fd5b803603821315611cb857600080fd5b611cc3868284611b46565b9a87019a9550505090840190600101611c6d565b5091979650505050505050565b60a081526000611cf860a083018a8c611b46565b611d05602084018a611c31565b87151560408401528281036060840152611d20818789611c53565b90508281036080840152611d35818587611b46565b9b9a5050505050505050505050565b601f821115611d8e57600081815260208120601f850160051c81016020861015611d6b5750805b601f850160051c820191505b81811015611d8a57828155600101611d77565b5050505b505050565b8135611d9e81611958565b815490151560ff1660ff1991909116178155600180820160208481013536869003601e19018112611dce57600080fd5b850180356001600160401b03811115611de657600080fd5b8036038383011315611df757600080fd5b611e0b81611e058654611bfd565b86611d44565b6000601f821160018114611e415760008315611e2957508382018501355b600019600385901b1c1916600184901b178655610bad565b600086815260209020601f19841690835b82811015611e7157868501880135825593870193908901908701611e52565b5084821015611e905760001960f88660031b161c198785880101351681555b50505050841b90930190915550505050565b600060208284031215611eb457600080fd5b815161106481611b14565b8060005b6002811015610485578151845260209384019390910190600101611ec3565b8481526001600160401b038416602082015260c08101611f056040830185611ebf565b610a9f6080830184611ebf565b6000808354611f2081611bfd565b60018281168015611f385760018114611f4d57611f7c565b60ff1984168752821515830287019450611f7c565b8760005260208060002060005b85811015611f735781548a820152908401908201611f5a565b50505082870194505b50929695505050505050565b8183823760009101908152919050565b6001600160a01b038816815260a060208201819052600090611fbc90830189611421565b611fc96040840189611c31565b8281036060840152611fdc818789611c53565b90508281036080840152611ff1818587611b46565b9a995050505050505050505056fea26469706673582212206b86326d1552a06317cfb2778fdd8b35dadf516778d5911b614aeb920992290564736f6c634300080f0033";

type RevertingBytesMarsConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: RevertingBytesMarsConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class RevertingBytesMars__factory extends ContractFactory {
  constructor(...args: RevertingBytesMarsConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    _dispatcher: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(_dispatcher, overrides || {});
  }
  override deploy(
    _dispatcher: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(_dispatcher, overrides || {}) as Promise<
      RevertingBytesMars & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): RevertingBytesMars__factory {
    return super.connect(runner) as RevertingBytesMars__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): RevertingBytesMarsInterface {
    return new Interface(_abi) as RevertingBytesMarsInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): RevertingBytesMars {
    return new Contract(address, _abi, runner) as unknown as RevertingBytesMars;
  }
}
