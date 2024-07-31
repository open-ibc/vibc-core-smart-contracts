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
  RevertingStringMars,
  RevertingStringMarsInterface,
} from "../../Mars.sol/RevertingStringMars";

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
        name: "",
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
    stateMutability: "view",
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
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "",
        type: "string",
        internalType: "string",
      },
    ],
    outputs: [],
    stateMutability: "view",
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
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "onTimeoutPacket",
    inputs: [
      {
        name: "packet",
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
    stateMutability: "nonpayable",
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
  "0x600360c0818152620312e360ec1b60e0526080908152610140604052610100918252620322e360ec1b6101205260a09190915262000042906006906002620000f9565b503480156200005057600080fd5b506040516200270f3803806200270f8339810160408190526200007391620001d0565b80806200008033620000a9565b600180546001600160a01b0319166001600160a01b039290921691909117905550620003739050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b82805482825590600052602060002090810192821562000144579160200282015b82811115620001445782518290620001339082620002a7565b50916020019190600101906200011a565b506200015292915062000156565b5090565b80821115620001525760006200016d828262000177565b5060010162000156565b508054620001859062000218565b6000825580601f1062000196575050565b601f016020900490600052602060002090810190620001b69190620001b9565b50565b5b80821115620001525760008155600101620001ba565b600060208284031215620001e357600080fd5b81516001600160a01b0381168114620001fb57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200022d57607f821691505b6020821081036200024e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620002a257600081815260208120601f850160051c810160208610156200027d5750805b601f850160051c820191505b818110156200029e5782815560010162000289565b5050505b505050565b81516001600160401b03811115620002c357620002c362000202565b620002db81620002d4845462000218565b8462000254565b602080601f831160018114620003135760008415620002fa5750858301515b600019600386901b1c1916600185901b1785556200029e565b600085815260208120601f198616915b82811015620003445788860151825594840194600190910190840162000323565b5085821015620003635787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61238c80620003836000396000f3fe6080604052600436106101395760003560e01c8063715018a6116100ab578063bb3f9f8d1161006f578063bb3f9f8d1461038b578063cb7e9057146103b9578063e847e280146103d9578063f12b758a146103f9578063f2fde38b14610419578063fad28a241461043957600080fd5b8063715018a6146102e45780637a805598146102f95780637d622184146103195780637e1d42b5146103395780638da5cb5b1461035957600080fd5b80634dcc0aa6116100fd5780634dcc0aa6146102125780634eeb73911461023f578063558850ac146102715780635bfd12b814610291578063602f9834146102b157806361995001146102d157600080fd5b80631eb7dd5e146101455780633513a995146101675780633f9fdbe4146101975780634252ae9b146101b75780634bdb5597146101e557600080fd5b3661014057005b600080fd5b34801561015157600080fd5b5061016561016036600461129c565b610459565b005b61017a610175366004611412565b61048a565b6040516001600160401b0390911681526020015b60405180910390f35b3480156101a357600080fd5b506101656101b236600461129c565b61052a565b3480156101c357600080fd5b506101d76101d2366004611492565b6105e6565b60405161018e9291906114f8565b3480156101f157600080fd5b50610205610200366004611599565b6106a2565b60405161018e91906116d2565b34801561021e57600080fd5b5061023261022d366004611778565b6106e7565b60405161018e9190611860565b34801561024b57600080fd5b5061025f61025a366004611492565b610774565b60405161018e969594939291906118b0565b34801561027d57600080fd5b5061016561028c366004611492565b6109d1565b34801561029d57600080fd5b5061017a6102ac366004611925565b610a33565b3480156102bd57600080fd5b506101656102cc36600461199b565b610ab6565b6101656102df3660046119cf565b610b27565b3480156102f057600080fd5b50610165610bf3565b34801561030557600080fd5b506101656103143660046119cf565b610c07565b34801561032557600080fd5b50610205610334366004611492565b610c89565b34801561034557600080fd5b50610165610354366004611ac6565b610d35565b34801561036557600080fd5b506000546001600160a01b03165b6040516001600160a01b03909116815260200161018e565b34801561039757600080fd5b506103ab6103a6366004611492565b610db8565b60405190815260200161018e565b3480156103c557600080fd5b50600154610373906001600160a01b031681565b3480156103e557600080fd5b506101656103f4366004611b30565b610dd9565b34801561040557600080fd5b5061025f610414366004611492565b610e4c565b34801561042557600080fd5b50610165610434366004611b97565b610e5c565b34801561044557600080fd5b50610165610454366004611492565b610ed5565b6001546001600160a01b03163314610484576040516321bf7f4960e01b815260040160405180910390fd5b50505050565b6001546040516330f8455760e21b81526000916001600160a01b03169063c3e1155c906104c19088908b908b908a90600401611bdd565b6020604051808303816000875af11580156104e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105049190611c11565b600154909150610520906001600160a01b031686838686610f00565b9695505050505050565b6001546001600160a01b03163314610555576040516321bf7f4960e01b815260040160405180910390fd5b6000805b6005548110156105c057856005828154811061057757610577611c2e565b9060005260206000200154036105ae576005818154811061059a5761059a611c2e565b6000918252602082200155600191506105c0565b806105b881611c44565b915050610559565b50806105df57604051630781f76560e21b815260040160405180910390fd5b5050505050565b600381815481106105f657600080fd5b60009182526020909120600290910201805460018201805460ff90921693509061061f90611c6b565b80601f016020809104026020016040519081016040528092919081815260200182805461064b90611c6b565b80156106985780601f1061066d57610100808354040283529160200191610698565b820191906000526020600020905b81548152906001019060200180831161067b57829003601f168201915b5050505050905082565b6001546060906001600160a01b031633146106d0576040516321bf7f4960e01b815260040160405180910390fd5b6106db868484610fcf565b98975050505050505050565b6040805180820190915260008152606060208201526001546001600160a01b03163314610727576040516321bf7f4960e01b815260040160405180910390fd5b60405162461bcd60e51b815260206004820152601b60248201527f6f6e2072656376207061636b657420697320726576657274696e67000000000060448201526064015b60405180910390fd5b6004818154811061078457600080fd5b9060005260206000209060080201600091509050806000016040518060400160405290816000820180546107b790611c6b565b80601f01602080910402602001604051908101604052809291908181526020018280546107e390611c6b565b80156108305780601f1061080557610100808354040283529160200191610830565b820191906000526020600020905b81548152906001019060200180831161081357829003601f168201915b50505050508152602001600182015481525050908060020160405180604001604052908160008201805461086390611c6b565b80601f016020809104026020016040519081016040528092919081815260200182805461088f90611c6b565b80156108dc5780601f106108b1576101008083540402835291602001916108dc565b820191906000526020600020905b8154815290600101906020018083116108bf57829003601f168201915b505050918352505060019190910154602090910152600482015460058301805492936001600160401b039092169261091390611c6b565b80601f016020809104026020016040519081016040528092919081815260200182805461093f90611c6b565b801561098c5780601f106109615761010080835404028352916020019161098c565b820191906000526020600020905b81548152906001019060200180831161096f57829003601f168201915b50506040805180820190915260068601546001600160401b03808216835268010000000000000000909104811660208301526007909601549495909416925088915050565b6109d96110f5565b6001546040516381bc079b60e01b8152600481018390526001600160a01b03909116906381bc079b90602401600060405180830381600087803b158015610a1f57600080fd5b505af11580156105df573d6000803e3d6000fd5b6001546040516330f8455760e21b81526000916001600160a01b03169063c3e1155c90610a6a908690899089908890600401611bdd565b6020604051808303816000875af1158015610a89573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aad9190611c11565b95945050505050565b6001546001600160a01b03163314610ae1576040516321bf7f4960e01b815260040160405180910390fd5b6004805460018101825560009190915281906008027f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01610b228282611f6a565b505050565b610b2f6110f5565b60015460405163418925b760e01b81526001600160a01b0390911690819063418925b790610b6f908c908c908c908c908c908c908c908c90600401612199565b600060405180830381600087803b158015610b8957600080fd5b505af1158015610b9d573d6000803e3d6000fd5b50505050610be8818a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c92508a9150899050888861114f565b505050505050505050565b610bfb6110f5565b610c056000611204565b565b610c0f6110f5565b60015460405163418925b760e01b81526001600160a01b039091169063418925b790610c4d908b908b908b908b908b908b908b908b90600401612199565b600060405180830381600087803b158015610c6757600080fd5b505af1158015610c7b573d6000803e3d6000fd5b505050505050505050505050565b60068181548110610c9957600080fd5b906000526020600020016000915090508054610cb490611c6b565b80601f0160208091040260200160405190810160405280929190818152602001828054610ce090611c6b565b8015610d2d5780601f10610d0257610100808354040283529160200191610d2d565b820191906000526020600020905b815481529060010190602001808311610d1057829003601f168201915b505050505081565b6001546001600160a01b03163314610d60576040516321bf7f4960e01b815260040160405180910390fd5b60405162461bcd60e51b815260206004820152602360248201527f61636b6e6f776c656467656d656e74207061636b657420697320726576657274604482015262696e6760e81b606482015260840161076b565b5050565b60058181548110610dc857600080fd5b600091825260209091200154905081565b6001546001600160a01b03163314610e04576040516321bf7f4960e01b815260040160405180910390fd5b60405162461bcd60e51b815260206004820181905260248201527f636f6e6e65637420696263206368616e6e656c20697320726576657274696e67604482015260640161076b565b6002818154811061078457600080fd5b610e646110f5565b6001600160a01b038116610ec95760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161076b565b610ed281611204565b50565b6001546001600160a01b03163314610ed2576040516321bf7f4960e01b815260040160405180910390fd5b846001600160a01b031663478222c26040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610f40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f6491906121f9565b6001600160a01b03166318e3404b34868686866040518663ffffffff1660e01b8152600401610f969493929190612239565b6000604051808303818588803b158015610faf57600080fd5b505af1158015610fc3573d6000803e3d6000fd5b50505050505050505050565b606060005b6006548110156110d45760068181548110610ff157610ff1611c2e565b9060005260206000200160405160200161100b9190612269565b6040516020818303038152906040528051906020012084846040516020016110349291906122df565b60405160208183030381529060405280519060200120036110c257600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001859055604080516020601f86018190048102820181019092528481529085908590819084018382808284376000920191909152509294506110ee9350505050565b806110cc81611c44565b915050610fd4565b5060405163b01318a560e01b815260040160405180910390fd5b9392505050565b6000546001600160a01b03163314610c055760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161076b565b866001600160a01b031663478222c26040518163ffffffff1660e01b81526004016020604051808303816000875af115801561118f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111b391906121f9565b6001600160a01b031663fce34e4034308989898989896040518963ffffffff1660e01b81526004016111eb97969594939291906122ef565b6000604051808303818588803b158015610c6757600080fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60008083601f84011261126657600080fd5b5081356001600160401b0381111561127d57600080fd5b60208301915083602082850101111561129557600080fd5b9250929050565b600080600080606085870312156112b257600080fd5b8435935060208501356001600160401b038111156112cf57600080fd5b6112db87828801611254565b9598909750949560400135949350505050565b6001600160401b0381168114610ed257600080fd5b803561130e816112ee565b919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561134b5761134b611313565b60405290565b60405160c081016001600160401b038111828210171561134b5761134b611313565b604051601f8201601f191681016001600160401b038111828210171561139b5761139b611313565b604052919050565b600082601f8301126113b457600080fd5b604051604081018181106001600160401b03821117156113d6576113d6611313565b80604052508060408401858111156113ed57600080fd5b845b818110156114075780358352602092830192016113ef565b509195945050505050565b60008060008060008060e0878903121561142b57600080fd5b86356001600160401b0381111561144157600080fd5b61144d89828a01611254565b909750955050602087013593506040870135611468816112ee565b925061147788606089016113a3565b91506114868860a089016113a3565b90509295509295509295565b6000602082840312156114a457600080fd5b5035919050565b6000815180845260005b818110156114d1576020818501810151868301820152016114b5565b818111156114e3576000602083870101525b50601f01601f19169290920160200192915050565b821515815260406020820152600061151360408301846114ab565b949350505050565b80356003811061130e57600080fd5b600082601f83011261153b57600080fd5b81356001600160401b0381111561155457611554611313565b611567601f8201601f1916602001611373565b81815284602083860101111561157c57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600080600060c0888a0312156115b457600080fd5b6115bd8861151b565b965060208801356001600160401b03808211156115d957600080fd5b818a0191508a601f8301126115ed57600080fd5b8135818111156115ff576115ff611313565b8060051b61160f60208201611373565b9182526020818501810192908101908e84111561162b57600080fd5b6020860192505b8383101561166957848335111561164857600080fd5b6116588f6020853589010161152a565b825260209283019290910190611632565b9a5050505060408a0135965060608a013591508082111561168957600080fd5b6116958b838c0161152a565b955060808a0135945060a08a01359150808211156116b257600080fd5b506116bf8a828b01611254565b989b979a50959850939692959293505050565b6020815260006110ee60208301846114ab565b6000604082840312156116f757600080fd5b6116ff611329565b905081356001600160401b0381111561171757600080fd5b6117238482850161152a565b8252506020820135602082015292915050565b60006040828403121561174857600080fd5b611750611329565b9050813561175d816112ee565b8152602082013561176d816112ee565b602082015292915050565b60006020828403121561178a57600080fd5b81356001600160401b03808211156117a157600080fd5b9083019060e082860312156117b557600080fd5b6117bd611351565b8235828111156117cc57600080fd5b6117d8878286016116e5565b8252506020830135828111156117ed57600080fd5b6117f9878286016116e5565b60208301525061180b60408401611303565b604082015260608301358281111561182257600080fd5b61182e8782860161152a565b6060830152506118418660808501611736565b608082015261185260c08401611303565b60a082015295945050505050565b602081528151151560208201526000602083015160408084015261151360608401826114ab565b600081516040845261189c60408501826114ab565b602093840151949093019390935250919050565b60e0815260006118c360e0830189611887565b82810360208401526118d58189611887565b90506001600160401b03808816604085015283820360608501526118f982886114ab565b92508086511660808501528060208701511660a085015280851660c08501525050979650505050505050565b6000806000806060858703121561193b57600080fd5b84356001600160401b0381111561195157600080fd5b61195d87828801611254565b909550935050602085013591506040850135611978816112ee565b939692955090935050565b600060e0828403121561199557600080fd5b50919050565b6000602082840312156119ad57600080fd5b81356001600160401b038111156119c357600080fd5b61151384828501611983565b60008060008060008060008060a0898b0312156119eb57600080fd5b88356001600160401b0380821115611a0257600080fd5b611a0e8c838d01611254565b909a509850889150611a2260208c0161151b565b975060408b013591508115158214611a3957600080fd5b90955060608a01359080821115611a4f57600080fd5b818b0191508b601f830112611a6357600080fd5b813581811115611a7257600080fd5b8c60208260051b8501011115611a8757600080fd5b6020830196508095505060808b0135915080821115611aa557600080fd5b50611ab28b828c01611254565b999c989b5096995094979396929594505050565b60008060408385031215611ad957600080fd5b82356001600160401b0380821115611af057600080fd5b611afc86838701611983565b93506020850135915080821115611b1257600080fd5b50830160408186031215611b2557600080fd5b809150509250929050565b60008060008060608587031215611b4657600080fd5b843593506020850135925060408501356001600160401b03811115611b6a57600080fd5b611b7687828801611254565b95989497509550505050565b6001600160a01b0381168114610ed257600080fd5b600060208284031215611ba957600080fd5b81356110ee81611b82565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b848152606060208201526000611bf7606083018587611bb4565b90506001600160401b038316604083015295945050505050565b600060208284031215611c2357600080fd5b81516110ee816112ee565b634e487b7160e01b600052603260045260246000fd5b600060018201611c6457634e487b7160e01b600052601160045260246000fd5b5060010190565b600181811c90821680611c7f57607f821691505b60208210810361199557634e487b7160e01b600052602260045260246000fd5b60008235603e19833603018112611cb557600080fd5b9190910192915050565b6000808335601e19843603018112611cd657600080fd5b8301803591506001600160401b03821115611cf057600080fd5b60200191503681900382131561129557600080fd5b601f821115610b2257600081815260208120601f850160051c81016020861015611d2c5750805b601f850160051c820191505b81811015611d4b57828155600101611d38565b505050505050565b600019600383901b1c191660019190911b1790565b611d728283611cbf565b6001600160401b03811115611d8957611d89611313565b611d9d81611d978554611c6b565b85611d05565b6000601f821160018114611dcb5760008315611db95750838201355b611dc38482611d53565b865550611e25565b600085815260209020601f19841690835b82811015611dfc5786850135825560209485019460019092019101611ddc565b5084821015611e195760001960f88660031b161c19848701351681555b505060018360011b0185555b50505050602082013560018201555050565b60008135611e44816112ee565b92915050565b6001600160401b03831115611e6157611e61611313565b611e7583611e6f8354611c6b565b83611d05565b6000601f841160018114611ea35760008515611e915750838201355b611e9b8682611d53565b8455506105df565b600083815260209020601f19861690835b82811015611ed45786850135825560209485019460019092019101611eb4565b5086821015611ef15760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b8135611f0e816112ee565b815467ffffffffffffffff19166001600160401b038216178255506020820135611f37816112ee565b81546fffffffffffffffff0000000000000000191660409190911b6fffffffffffffffff00000000000000001617905550565b611f748283611c9f565b611f7e8182611cbf565b6001600160401b03811115611f9557611f95611313565b611fa981611fa38654611c6b565b86611d05565b6000601f821160018114611fd75760008315611fc55750838201355b611fcf8482611d53565b875550612031565b600086815260209020601f19841690835b828110156120085786850135825560209485019460019092019101611fe8565b50848210156120255760001960f88660031b161c19848701351681555b505060018360011b0186555b5050505060208101356001830155506120596120506020840184611c9f565b60028301611d68565b61208961206860408401611e37565b600483016001600160401b0382166001600160401b03198254161781555050565b6120966060830183611cbf565b6120a4818360058601611e4a565b50506120b66080830160068301611f03565b610db46120c560c08401611e37565b600783016001600160401b0382166001600160401b03198254161781555050565b6003811061210457634e487b7160e01b600052602160045260246000fd5b9052565b81835260006020808501808196508560051b810191508460005b8781101561218c5782840389528135601e1988360301811261214357600080fd5b870185810190356001600160401b0381111561215e57600080fd5b80360382131561216d57600080fd5b612178868284611bb4565b9a87019a9550505090840190600101612122565b5091979650505050505050565b60a0815260006121ad60a083018a8c611bb4565b6121ba602084018a6120e6565b871515604084015282810360608401526121d5818789612108565b905082810360808401526121ea818587611bb4565b9b9a5050505050505050505050565b60006020828403121561220b57600080fd5b81516110ee81611b82565b8060005b600281101561048457815184526020938401939091019060010161221a565b8481526001600160401b038416602082015260c0810161225c6040830185612216565b610aad6080830184612216565b600080835461227781611c6b565b6001828116801561228f57600181146122a4576122d3565b60ff19841687528215158302870194506122d3565b8760005260208060002060005b858110156122ca5781548a8201529084019082016122b1565b50505082870194505b50929695505050505050565b8183823760009101908152919050565b6001600160a01b038816815260a060208201819052600090612313908301896114ab565b61232060408401896120e6565b8281036060840152612333818789612108565b90508281036080840152612348818587611bb4565b9a995050505050505050505056fea264697066735822122028de94cdc5ddc73807a3341483c8bf5fec3ed6a9359496fbbd9f2098a090544064736f6c634300080f0033";

type RevertingStringMarsConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: RevertingStringMarsConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class RevertingStringMars__factory extends ContractFactory {
  constructor(...args: RevertingStringMarsConstructorParams) {
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
      RevertingStringMars & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(
    runner: ContractRunner | null
  ): RevertingStringMars__factory {
    return super.connect(runner) as RevertingStringMars__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): RevertingStringMarsInterface {
    return new Interface(_abi) as RevertingStringMarsInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): RevertingStringMars {
    return new Contract(
      address,
      _abi,
      runner
    ) as unknown as RevertingStringMars;
  }
}
