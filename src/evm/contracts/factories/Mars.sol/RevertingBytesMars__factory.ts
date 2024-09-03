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
  "0x600360c0818152620312e360ec1b60e0526080908152610140604052610100918252620322e360ec1b6101205260a09190915262000042906006906002620000f9565b503480156200005057600080fd5b50604051620023a2380380620023a28339810160408190526200007391620001d0565b80806200008033620000a9565b600180546001600160a01b0319166001600160a01b039290921691909117905550620003739050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b82805482825590600052602060002090810192821562000144579160200282015b82811115620001445782518290620001339082620002a7565b50916020019190600101906200011a565b506200015292915062000156565b5090565b80821115620001525760006200016d828262000177565b5060010162000156565b508054620001859062000218565b6000825580601f1062000196575050565b601f016020900490600052602060002090810190620001b69190620001b9565b50565b5b80821115620001525760008155600101620001ba565b600060208284031215620001e357600080fd5b81516001600160a01b0381168114620001fb57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200022d57607f821691505b6020821081036200024e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620002a257600081815260208120601f850160051c810160208610156200027d5750805b601f850160051c820191505b818110156200029e5782815560010162000289565b5050505b505050565b81516001600160401b03811115620002c357620002c362000202565b620002db81620002d4845462000218565b8462000254565b602080601f831160018114620003135760008415620002fa5750858301515b600019600386901b1c1916600185901b1785556200029e565b600085815260208120601f198616915b82811015620003445788860151825594840194600190910190840162000323565b5085821015620003635787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61201f80620003836000396000f3fe6080604052600436106101395760003560e01c8063715018a6116100ab578063bb3f9f8d1161006f578063bb3f9f8d1461038b578063cb7e9057146103b9578063e847e280146103d9578063f12b758a146103f9578063f2fde38b14610419578063fad28a241461043957600080fd5b8063715018a6146102e45780637a805598146102f95780637d622184146103195780637e1d42b5146103395780638da5cb5b1461035957600080fd5b80634dcc0aa6116100fd5780634dcc0aa6146102125780634eeb73911461023f578063558850ac146102715780635bfd12b814610291578063602f9834146102b157806361995001146102d157600080fd5b80631eb7dd5e146101455780633513a995146101675780633f9fdbe4146101975780634252ae9b146101b75780634bdb5597146101e557600080fd5b3661014057005b600080fd5b34801561015157600080fd5b5061016561016036600461120b565b610459565b005b61017a610175366004611381565b61048a565b6040516001600160401b0390911681526020015b60405180910390f35b3480156101a357600080fd5b506101656101b236600461120b565b61052a565b3480156101c357600080fd5b506101d76101d2366004611401565b6105e6565b60405161018e929190611467565b3480156101f157600080fd5b50610205610200366004611508565b6106a2565b60405161018e9190611641565b34801561021e57600080fd5b5061023261022d3660046116e7565b6106e7565b60405161018e91906117cf565b34801561024b57600080fd5b5061025f61025a366004611401565b61075f565b60405161018e9695949392919061181f565b34801561027d57600080fd5b5061016561028c366004611401565b6109bc565b34801561029d57600080fd5b5061017a6102ac366004611894565b610a1e565b3480156102bd57600080fd5b506101656102cc36600461190a565b610aa1565b6101656102df36600461194c565b610ae5565b3480156102f057600080fd5b50610165610bb1565b34801561030557600080fd5b5061016561031436600461194c565b610bc5565b34801561032557600080fd5b50610205610334366004611401565b610c47565b34801561034557600080fd5b50610165610354366004611a3e565b610cf3565b34801561036557600080fd5b506000546001600160a01b03165b6040516001600160a01b03909116815260200161018e565b34801561039757600080fd5b506103ab6103a6366004611401565b610d5f565b60405190815260200161018e565b3480156103c557600080fd5b50600154610373906001600160a01b031681565b3480156103e557600080fd5b506101656103f4366004611aa8565b610d80565b34801561040557600080fd5b5061025f610414366004611401565b610db6565b34801561042557600080fd5b50610165610434366004611b0f565b610dc6565b34801561044557600080fd5b50610165610454366004611401565b610e44565b6001546001600160a01b03163314610484576040516321bf7f4960e01b815260040160405180910390fd5b50505050565b6001546040516330f8455760e21b81526000916001600160a01b03169063c3e1155c906104c19088908b908b908a90600401611b55565b6020604051808303816000875af11580156104e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105049190611b89565b600154909150610520906001600160a01b031686838686610e6f565b9695505050505050565b6001546001600160a01b03163314610555576040516321bf7f4960e01b815260040160405180910390fd5b6000805b6005548110156105c057856005828154811061057757610577611ba6565b9060005260206000200154036105ae576005818154811061059a5761059a611ba6565b6000918252602082200155600191506105c0565b806105b881611bbc565b915050610559565b50806105df57604051630781f76560e21b815260040160405180910390fd5b5050505050565b600381815481106105f657600080fd5b60009182526020909120600290910201805460018201805460ff90921693509061061f90611be3565b80601f016020809104026020016040519081016040528092919081815260200182805461064b90611be3565b80156106985780601f1061066d57610100808354040283529160200191610698565b820191906000526020600020905b81548152906001019060200180831161067b57829003601f168201915b5050505050905082565b6001546060906001600160a01b031633146106d0576040516321bf7f4960e01b815260040160405180910390fd5b6106db868484610f3e565b98975050505050505050565b6040805180820190915260008152606060208201526001546001600160a01b03163314610727576040516321bf7f4960e01b815260040160405180910390fd5b506040805180820182526000808252825160208181018552918152908201528151639889d82160e01b81529151909181900360040190fd5b6004818154811061076f57600080fd5b9060005260206000209060080201600091509050806000016040518060400160405290816000820180546107a290611be3565b80601f01602080910402602001604051908101604052809291908181526020018280546107ce90611be3565b801561081b5780601f106107f05761010080835404028352916020019161081b565b820191906000526020600020905b8154815290600101906020018083116107fe57829003601f168201915b50505050508152602001600182015481525050908060020160405180604001604052908160008201805461084e90611be3565b80601f016020809104026020016040519081016040528092919081815260200182805461087a90611be3565b80156108c75780601f1061089c576101008083540402835291602001916108c7565b820191906000526020600020905b8154815290600101906020018083116108aa57829003601f168201915b505050918352505060019190910154602090910152600482015460058301805492936001600160401b03909216926108fe90611be3565b80601f016020809104026020016040519081016040528092919081815260200182805461092a90611be3565b80156109775780601f1061094c57610100808354040283529160200191610977565b820191906000526020600020905b81548152906001019060200180831161095a57829003601f168201915b50506040805180820190915260068601546001600160401b03808216835268010000000000000000909104811660208301526007909601549495909416925088915050565b6109c4611064565b6001546040516381bc079b60e01b8152600481018390526001600160a01b03909116906381bc079b90602401600060405180830381600087803b158015610a0a57600080fd5b505af11580156105df573d6000803e3d6000fd5b6001546040516330f8455760e21b81526000916001600160a01b03169063c3e1155c90610a55908690899089908890600401611b55565b6020604051808303816000875af1158015610a74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a989190611b89565b95945050505050565b6001546001600160a01b03163314610acc576040516321bf7f4960e01b815260040160405180910390fd5b604051631021bb3b60e31b815260040160405180910390fd5b610aed611064565b60015460405163418925b760e01b81526001600160a01b0390911690819063418925b790610b2d908c908c908c908c908c908c908c908c90600401611cce565b600060405180830381600087803b158015610b4757600080fd5b505af1158015610b5b573d6000803e3d6000fd5b50505050610ba6818a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508c92508a915089905088886110be565b505050505050505050565b610bb9611064565b610bc36000611173565b565b610bcd611064565b60015460405163418925b760e01b81526001600160a01b039091169063418925b790610c0b908b908b908b908b908b908b908b908b90600401611cce565b600060405180830381600087803b158015610c2557600080fd5b505af1158015610c39573d6000803e3d6000fd5b505050505050505050505050565b60068181548110610c5757600080fd5b906000526020600020016000915090508054610c7290611be3565b80601f0160208091040260200160405190810160405280929190818152602001828054610c9e90611be3565b8015610ceb5780601f10610cc057610100808354040283529160200191610ceb565b820191906000526020600020905b815481529060010190602001808311610cce57829003601f168201915b505050505081565b6001546001600160a01b03163314610d1e576040516321bf7f4960e01b815260040160405180910390fd5b6003805460018101825560009190915281906002027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b016104848282611d7d565b60058181548110610d6f57600080fd5b600091825260209091200154905081565b6001546001600160a01b03163314610dab576040516321bf7f4960e01b815260040160405180910390fd5b6105df848383610f3e565b6002818154811061076f57600080fd5b610dce611064565b6001600160a01b038116610e385760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b610e4181611173565b50565b6001546001600160a01b03163314610e41576040516321bf7f4960e01b815260040160405180910390fd5b846001600160a01b031663478222c26040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610eaf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed39190611e8c565b6001600160a01b03166318e3404b34868686866040518663ffffffff1660e01b8152600401610f059493929190611ecc565b6000604051808303818588803b158015610f1e57600080fd5b505af1158015610f32573d6000803e3d6000fd5b50505050505050505050565b606060005b6006548110156110435760068181548110610f6057610f60611ba6565b90600052602060002001604051602001610f7a9190611efc565b604051602081830303815290604052805190602001208484604051602001610fa3929190611f72565b604051602081830303815290604052805190602001200361103157600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001859055604080516020601f860181900481028201810190925284815290859085908190840183828082843760009201919091525092945061105d9350505050565b8061103b81611bbc565b915050610f43565b5060405163b01318a560e01b815260040160405180910390fd5b9392505050565b6000546001600160a01b03163314610bc35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610e2f565b866001600160a01b031663478222c26040518163ffffffff1660e01b81526004016020604051808303816000875af11580156110fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111229190611e8c565b6001600160a01b031663fce34e4034308989898989896040518963ffffffff1660e01b815260040161115a9796959493929190611f82565b6000604051808303818588803b158015610c2557600080fd5b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60008083601f8401126111d557600080fd5b5081356001600160401b038111156111ec57600080fd5b60208301915083602082850101111561120457600080fd5b9250929050565b6000806000806060858703121561122157600080fd5b8435935060208501356001600160401b0381111561123e57600080fd5b61124a878288016111c3565b9598909750949560400135949350505050565b6001600160401b0381168114610e4157600080fd5b803561127d8161125d565b919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156112ba576112ba611282565b60405290565b60405160c081016001600160401b03811182821017156112ba576112ba611282565b604051601f8201601f191681016001600160401b038111828210171561130a5761130a611282565b604052919050565b600082601f83011261132357600080fd5b604051604081018181106001600160401b038211171561134557611345611282565b806040525080604084018581111561135c57600080fd5b845b8181101561137657803583526020928301920161135e565b509195945050505050565b60008060008060008060e0878903121561139a57600080fd5b86356001600160401b038111156113b057600080fd5b6113bc89828a016111c3565b9097509550506020870135935060408701356113d78161125d565b92506113e68860608901611312565b91506113f58860a08901611312565b90509295509295509295565b60006020828403121561141357600080fd5b5035919050565b6000815180845260005b8181101561144057602081850181015186830182015201611424565b81811115611452576000602083870101525b50601f01601f19169290920160200192915050565b8215158152604060208201526000611482604083018461141a565b949350505050565b80356003811061127d57600080fd5b600082601f8301126114aa57600080fd5b81356001600160401b038111156114c3576114c3611282565b6114d6601f8201601f19166020016112e2565b8181528460208386010111156114eb57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600080600060c0888a03121561152357600080fd5b61152c8861148a565b965060208801356001600160401b038082111561154857600080fd5b818a0191508a601f83011261155c57600080fd5b81358181111561156e5761156e611282565b8060051b61157e602082016112e2565b9182526020818501810192908101908e84111561159a57600080fd5b6020860192505b838310156115d85784833511156115b757600080fd5b6115c78f60208535890101611499565b8252602092830192909101906115a1565b9a5050505060408a0135965060608a01359150808211156115f857600080fd5b6116048b838c01611499565b955060808a0135945060a08a013591508082111561162157600080fd5b5061162e8a828b016111c3565b989b979a50959850939692959293505050565b60208152600061105d602083018461141a565b60006040828403121561166657600080fd5b61166e611298565b905081356001600160401b0381111561168657600080fd5b61169284828501611499565b8252506020820135602082015292915050565b6000604082840312156116b757600080fd5b6116bf611298565b905081356116cc8161125d565b815260208201356116dc8161125d565b602082015292915050565b6000602082840312156116f957600080fd5b81356001600160401b038082111561171057600080fd5b9083019060e0828603121561172457600080fd5b61172c6112c0565b82358281111561173b57600080fd5b61174787828601611654565b82525060208301358281111561175c57600080fd5b61176887828601611654565b60208301525061177a60408401611272565b604082015260608301358281111561179157600080fd5b61179d87828601611499565b6060830152506117b086608085016116a5565b60808201526117c160c08401611272565b60a082015295945050505050565b6020815281511515602082015260006020830151604080840152611482606084018261141a565b600081516040845261180b604085018261141a565b602093840151949093019390935250919050565b60e08152600061183260e08301896117f6565b828103602084015261184481896117f6565b90506001600160401b0380881660408501528382036060850152611868828861141a565b92508086511660808501528060208701511660a085015280851660c08501525050979650505050505050565b600080600080606085870312156118aa57600080fd5b84356001600160401b038111156118c057600080fd5b6118cc878288016111c3565b9095509350506020850135915060408501356118e78161125d565b939692955090935050565b600060e0828403121561190457600080fd5b50919050565b60006020828403121561191c57600080fd5b81356001600160401b0381111561193257600080fd5b611482848285016118f2565b8015158114610e4157600080fd5b60008060008060008060008060a0898b03121561196857600080fd5b88356001600160401b038082111561197f57600080fd5b61198b8c838d016111c3565b909a50985088915061199f60208c0161148a565b975060408b013591506119b18261193e565b90955060608a013590808211156119c757600080fd5b818b0191508b601f8301126119db57600080fd5b8135818111156119ea57600080fd5b8c60208260051b85010111156119ff57600080fd5b6020830196508095505060808b0135915080821115611a1d57600080fd5b50611a2a8b828c016111c3565b999c989b5096995094979396929594505050565b60008060408385031215611a5157600080fd5b82356001600160401b0380821115611a6857600080fd5b611a74868387016118f2565b93506020850135915080821115611a8a57600080fd5b50830160408186031215611a9d57600080fd5b809150509250929050565b60008060008060608587031215611abe57600080fd5b843593506020850135925060408501356001600160401b03811115611ae257600080fd5b611aee878288016111c3565b95989497509550505050565b6001600160a01b0381168114610e4157600080fd5b600060208284031215611b2157600080fd5b813561105d81611afa565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b848152606060208201526000611b6f606083018587611b2c565b90506001600160401b038316604083015295945050505050565b600060208284031215611b9b57600080fd5b815161105d8161125d565b634e487b7160e01b600052603260045260246000fd5b600060018201611bdc57634e487b7160e01b600052601160045260246000fd5b5060010190565b600181811c90821680611bf757607f821691505b60208210810361190457634e487b7160e01b600052602260045260246000fd5b60038110611c3557634e487b7160e01b600052602160045260246000fd5b9052565b60008383855260208086019550808560051b8301018460005b87811015611cc157848303601f19018952813536889003601e19018112611c7857600080fd5b870184810190356001600160401b03811115611c9357600080fd5b803603821315611ca257600080fd5b611cad858284611b2c565b9a86019a9450505090830190600101611c52565b5090979650505050505050565b60a081526000611ce260a083018a8c611b2c565b611cef602084018a611c17565b87151560408401528281036060840152611d0a818789611c39565b90508281036080840152611d1f818587611b2c565b9b9a5050505050505050505050565b601f821115611d7857600081815260208120601f850160051c81016020861015611d555750805b601f850160051c820191505b81811015611d7457828155600101611d61565b5050505b505050565b8135611d888161193e565b815490151560ff1660ff1991909116178155600180820160208481013536869003601e19018112611db857600080fd5b850180356001600160401b03811115611dd057600080fd5b8036038383011315611de157600080fd5b611df581611def8654611be3565b86611d2e565b6000601f821160018114611e2b5760008315611e1357508382018501355b600019600385901b1c1916600184901b178655610ba6565b600086815260209020601f19841690835b82811015611e5b57868501880135825593870193908901908701611e3c565b5084821015611e7a5760001960f88660031b161c198785880101351681555b50505050841b90930190915550505050565b600060208284031215611e9e57600080fd5b815161105d81611afa565b8060005b6002811015610484578151845260209384019390910190600101611ead565b8481526001600160401b038416602082015260c08101611eef6040830185611ea9565b610a986080830184611ea9565b6000808354611f0a81611be3565b60018281168015611f225760018114611f3757611f66565b60ff1984168752821515830287019450611f66565b8760005260208060002060005b85811015611f5d5781548a820152908401908201611f44565b50505082870194505b50929695505050505050565b8183823760009101908152919050565b6001600160a01b038816815260a060208201819052600090611fa69083018961141a565b611fb36040840189611c17565b8281036060840152611fc6818789611c39565b90508281036080840152611fdb818587611b2c565b9a995050505050505050505056fea264697066735822122070eabf3bdb9f0f3f87a52efc54c99b258f1feb469d02ec141e1d99997b29f47964736f6c634300080f0033";

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
