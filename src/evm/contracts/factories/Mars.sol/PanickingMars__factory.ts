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
  PanickingMars,
  PanickingMarsInterface,
} from "../../Mars.sol/PanickingMars";

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
    outputs: [],
    stateMutability: "nonpayable",
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
    name: "onChanOpenInit",
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
        name: "",
        type: "string",
        internalType: "string",
      },
      {
        name: "version",
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
    stateMutability: "view",
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
  "0x600360c0818152620312e360ec1b60e0526080908152610140604052610100918252620322e360ec1b6101205260a09190915262000042906006906002620000f9565b503480156200005057600080fd5b5060405162002419380380620024198339810160408190526200007391620001d0565b80806200008033620000a9565b600180546001600160a01b0319166001600160a01b039290921691909117905550620003739050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b82805482825590600052602060002090810192821562000144579160200282015b82811115620001445782518290620001339082620002a7565b50916020019190600101906200011a565b506200015292915062000156565b5090565b80821115620001525760006200016d828262000177565b5060010162000156565b508054620001859062000218565b6000825580601f1062000196575050565b601f016020900490600052602060002090810190620001b69190620001b9565b50565b5b80821115620001525760008155600101620001ba565b600060208284031215620001e357600080fd5b81516001600160a01b0381168114620001fb57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200022d57607f821691505b6020821081036200024e57634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620002a257600081815260208120601f850160051c810160208610156200027d5750805b601f850160051c820191505b818110156200029e5782815560010162000289565b5050505b505050565b81516001600160401b03811115620002c357620002c362000202565b620002db81620002d4845462000218565b8462000254565b602080601f831160018114620003135760008415620002fa5750858301515b600019600386901b1c1916600185901b1785556200029e565b600085815260208120601f198616915b82811015620003445788860151825594840194600190910190840162000323565b5085821015620003635787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61209680620003836000396000f3fe60806040526004361061012e5760003560e01c80637a805598116100ab578063bb3f9f8d1161006f578063bb3f9f8d14610366578063cb7e905714610394578063e847e280146103b4578063f12b758a146103d4578063f2fde38b146103f4578063fad28a241461041457600080fd5b80637a805598146102b45780637a9ccc4b146102d45780637d622184146102f45780637e1d42b5146103145780638da5cb5b1461033457600080fd5b80634eeb7391116100f25780634eeb73911461020d578063558850ac1461023f5780635bfd12b81461025f578063602f98341461027f578063715018a61461029f57600080fd5b80631eb7dd5e1461013a5780633f9fdbe41461015c5780634252ae9b1461017c5780634bdb5597146101b35780634dcc0aa6146101e057600080fd5b3661013557005b600080fd5b34801561014657600080fd5b5061015a610155366004611015565b610434565b005b34801561016857600080fd5b5061015a610177366004611015565b610465565b34801561018857600080fd5b5061019c610197366004611067565b610521565b6040516101aa9291906110cd565b60405180910390f35b3480156101bf57600080fd5b506101d36101ce366004611203565b6105dd565b6040516101aa919061133c565b3480156101ec57600080fd5b506102006101fb366004611402565b610622565b6040516101aa91906114ea565b34801561021957600080fd5b5061022d610228366004611067565b610691565b6040516101aa9695949392919061153a565b34801561024b57600080fd5b5061015a61025a366004611067565b6108ee565b34801561026b57600080fd5b5061015a61027a3660046115af565b610950565b34801561028b57600080fd5b5061015a61029a366004611625565b6109be565b3480156102ab57600080fd5b5061015a610a2f565b3480156102c057600080fd5b5061015a6102cf3660046116ab565b610a43565b3480156102e057600080fd5b506101d36102ef36600461176c565b610ac5565b34801561030057600080fd5b506101d361030f366004611067565b610afd565b34801561032057600080fd5b5061015a61032f3660046117f6565b610ba9565b34801561034057600080fd5b506000546001600160a01b03165b6040516001600160a01b0390911681526020016101aa565b34801561037257600080fd5b50610386610381366004611067565b610c15565b6040519081526020016101aa565b3480156103a057600080fd5b5060015461034e906001600160a01b031681565b3480156103c057600080fd5b5061015a6103cf366004611860565b610c36565b3480156103e057600080fd5b5061022d6103ef366004611067565b610c6c565b34801561040057600080fd5b5061015a61040f3660046118b2565b610c7c565b34801561042057600080fd5b5061015a61042f366004611067565b610cfa565b6001546001600160a01b0316331461045f576040516321bf7f4960e01b815260040160405180910390fd5b50505050565b6001546001600160a01b03163314610490576040516321bf7f4960e01b815260040160405180910390fd5b6000805b6005548110156104fb5785600582815481106104b2576104b26118db565b9060005260206000200154036104e957600581815481106104d5576104d56118db565b6000918252602082200155600191506104fb565b806104f3816118f1565b915050610494565b508061051a57604051630781f76560e21b815260040160405180910390fd5b5050505050565b6003818154811061053157600080fd5b60009182526020909120600290910201805460018201805460ff90921693509061055a90611918565b80601f016020809104026020016040519081016040528092919081815260200182805461058690611918565b80156105d35780601f106105a8576101008083540402835291602001916105d3565b820191906000526020600020905b8154815290600101906020018083116105b657829003601f168201915b5050505050905082565b6001546060906001600160a01b0316331461060b576040516321bf7f4960e01b815260040160405180910390fd5b610616868484610d25565b98975050505050505050565b6040805180820190915260008152606060208201526001546001600160a01b03163314610662576040516321bf7f4960e01b815260040160405180910390fd5b61066a61194c565b50506040805180820182526000808252825160208181019094529081529181019190915290565b600481815481106106a157600080fd5b9060005260206000209060080201600091509050806000016040518060400160405290816000820180546106d490611918565b80601f016020809104026020016040519081016040528092919081815260200182805461070090611918565b801561074d5780601f106107225761010080835404028352916020019161074d565b820191906000526020600020905b81548152906001019060200180831161073057829003601f168201915b50505050508152602001600182015481525050908060020160405180604001604052908160008201805461078090611918565b80601f01602080910402602001604051908101604052809291908181526020018280546107ac90611918565b80156107f95780601f106107ce576101008083540402835291602001916107f9565b820191906000526020600020905b8154815290600101906020018083116107dc57829003601f168201915b505050918352505060019190910154602090910152600482015460058301805492936001600160401b039092169261083090611918565b80601f016020809104026020016040519081016040528092919081815260200182805461085c90611918565b80156108a95780601f1061087e576101008083540402835291602001916108a9565b820191906000526020600020905b81548152906001019060200180831161088c57829003601f168201915b50506040805180820190915260068601546001600160401b03808216835268010000000000000000909104811660208301526007909601549495909416925088915050565b6108f6610e4b565b6001546040516381bc079b60e01b8152600481018390526001600160a01b03909116906381bc079b90602401600060405180830381600087803b15801561093c57600080fd5b505af115801561051a573d6000803e3d6000fd5b6001546040516330f8455760e21b81526001600160a01b039091169063c3e1155c9061098690859088908890879060040161198b565b600060405180830381600087803b1580156109a057600080fd5b505af11580156109b4573d6000803e3d6000fd5b5050505050505050565b6001546001600160a01b031633146109e9576040516321bf7f4960e01b815260040160405180910390fd5b6004805460018101825560009190915281906008027f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b01610a2a8282611c84565b505050565b610a37610e4b565b610a416000610ea5565b565b610a4b610e4b565b60015460405163418925b760e01b81526001600160a01b039091169063418925b790610a89908b908b908b908b908b908b908b908b90600401611e04565b600060405180830381600087803b158015610aa357600080fd5b505af1158015610ab7573d6000803e3d6000fd5b505050505050505050505050565b6001546060906001600160a01b03163314610af3576040516321bf7f4960e01b815260040160405180910390fd5b6106168383610ef5565b60068181548110610b0d57600080fd5b906000526020600020016000915090508054610b2890611918565b80601f0160208091040260200160405190810160405280929190818152602001828054610b5490611918565b8015610ba15780601f10610b7657610100808354040283529160200191610ba1565b820191906000526020600020905b815481529060010190602001808311610b8457829003601f168201915b505050505081565b6001546001600160a01b03163314610bd4576040516321bf7f4960e01b815260040160405180910390fd5b6003805460018101825560009190915281906002027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0161045f8282611ef7565b60058181548110610c2557600080fd5b600091825260209091200154905081565b6001546001600160a01b03163314610c61576040516321bf7f4960e01b815260040160405180910390fd5b61051a848383610d25565b600281815481106106a157600080fd5b610c84610e4b565b6001600160a01b038116610cee5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b610cf781610ea5565b50565b6001546001600160a01b03163314610cf7576040516321bf7f4960e01b815260040160405180910390fd5b606060005b600654811015610e2a5760068181548110610d4757610d476118db565b90600052602060002001604051602001610d619190611fda565b604051602081830303815290604052805190602001208484604051602001610d8a929190612050565b6040516020818303038152906040528051906020012003610e1857600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001859055604080516020601f8601819004810282018101909252848152908590859081908401838280828437600092019190915250929450610e449350505050565b80610e22816118f1565b915050610d2a565b5060405163b01318a560e01b815260040160405180910390fd5b9392505050565b6000546001600160a01b03163314610a415760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610ce5565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b606060005b600654811015610e2a5760068181548110610f1757610f176118db565b90600052602060002001604051602001610f319190611fda565b604051602081830303815290604052805190602001208484604051602001610f5a929190612050565b6040516020818303038152906040528051906020012003610fb55783838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929450610fc79350505050565b80610fbf816118f1565b915050610efa565b92915050565b60008083601f840112610fdf57600080fd5b5081356001600160401b03811115610ff657600080fd5b60208301915083602082850101111561100e57600080fd5b9250929050565b6000806000806060858703121561102b57600080fd5b8435935060208501356001600160401b0381111561104857600080fd5b61105487828801610fcd565b9598909750949560400135949350505050565b60006020828403121561107957600080fd5b5035919050565b6000815180845260005b818110156110a65760208185018101518683018201520161108a565b818111156110b8576000602083870101525b50601f01601f19169290920160200192915050565b82151581526040602082015260006110e86040830184611080565b949350505050565b8035600381106110ff57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561113c5761113c611104565b60405290565b60405160c081016001600160401b038111828210171561113c5761113c611104565b604051601f8201601f191681016001600160401b038111828210171561118c5761118c611104565b604052919050565b600082601f8301126111a557600080fd5b81356001600160401b038111156111be576111be611104565b6111d1601f8201601f1916602001611164565b8181528460208386010111156111e657600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600080600060c0888a03121561121e57600080fd5b611227886110f0565b965060208801356001600160401b038082111561124357600080fd5b818a0191508a601f83011261125757600080fd5b81358181111561126957611269611104565b8060051b61127960208201611164565b9182526020818501810192908101908e84111561129557600080fd5b6020860192505b838310156112d35784833511156112b257600080fd5b6112c28f60208535890101611194565b82526020928301929091019061129c565b9a5050505060408a0135965060608a01359150808211156112f357600080fd5b6112ff8b838c01611194565b955060808a0135945060a08a013591508082111561131c57600080fd5b506113298a828b01610fcd565b989b979a50959850939692959293505050565b602081526000610e446020830184611080565b60006040828403121561136157600080fd5b61136961111a565b905081356001600160401b0381111561138157600080fd5b61138d84828501611194565b8252506020820135602082015292915050565b6001600160401b0381168114610cf757600080fd5b80356110ff816113a0565b6000604082840312156113d257600080fd5b6113da61111a565b905081356113e7816113a0565b815260208201356113f7816113a0565b602082015292915050565b60006020828403121561141457600080fd5b81356001600160401b038082111561142b57600080fd5b9083019060e0828603121561143f57600080fd5b611447611142565b82358281111561145657600080fd5b6114628782860161134f565b82525060208301358281111561147757600080fd5b6114838782860161134f565b602083015250611495604084016113b5565b60408201526060830135828111156114ac57600080fd5b6114b887828601611194565b6060830152506114cb86608085016113c0565b60808201526114dc60c084016113b5565b60a082015295945050505050565b60208152815115156020820152600060208301516040808401526110e86060840182611080565b60008151604084526115266040850182611080565b602093840151949093019390935250919050565b60e08152600061154d60e0830189611511565b828103602084015261155f8189611511565b90506001600160401b03808816604085015283820360608501526115838288611080565b92508086511660808501528060208701511660a085015280851660c08501525050979650505050505050565b600080600080606085870312156115c557600080fd5b84356001600160401b038111156115db57600080fd5b6115e787828801610fcd565b909550935050602085013591506040850135611602816113a0565b939692955090935050565b600060e0828403121561161f57600080fd5b50919050565b60006020828403121561163757600080fd5b81356001600160401b0381111561164d57600080fd5b6110e88482850161160d565b8015158114610cf757600080fd5b60008083601f84011261167957600080fd5b5081356001600160401b0381111561169057600080fd5b6020830191508360208260051b850101111561100e57600080fd5b60008060008060008060008060a0898b0312156116c757600080fd5b88356001600160401b03808211156116de57600080fd5b6116ea8c838d01610fcd565b909a5098508891506116fe60208c016110f0565b975060408b0135915061171082611659565b90955060608a0135908082111561172657600080fd5b6117328c838d01611667565b909650945060808b013591508082111561174b57600080fd5b506117588b828c01610fcd565b999c989b5096995094979396929594505050565b60008060008060008060006080888a03121561178757600080fd5b611790886110f0565b965060208801356001600160401b03808211156117ac57600080fd5b6117b88b838c01611667565b909850965060408a01359150808211156117d157600080fd5b6117dd8b838c01610fcd565b909650945060608a013591508082111561131c57600080fd5b6000806040838503121561180957600080fd5b82356001600160401b038082111561182057600080fd5b61182c8683870161160d565b9350602085013591508082111561184257600080fd5b5083016040818603121561185557600080fd5b809150509250929050565b6000806000806060858703121561187657600080fd5b843593506020850135925060408501356001600160401b0381111561189a57600080fd5b6118a687828801610fcd565b95989497509550505050565b6000602082840312156118c457600080fd5b81356001600160a01b0381168114610e4457600080fd5b634e487b7160e01b600052603260045260246000fd5b60006001820161191157634e487b7160e01b600052601160045260246000fd5b5060010190565b600181811c9082168061192c57607f821691505b60208210810361161f57634e487b7160e01b600052602260045260246000fd5b634e487b7160e01b600052600160045260246000fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8481526060602082015260006119a5606083018587611962565b90506001600160401b038316604083015295945050505050565b60008235603e198336030181126119d557600080fd5b9190910192915050565b6000808335601e198436030181126119f657600080fd5b8301803591506001600160401b03821115611a1057600080fd5b60200191503681900382131561100e57600080fd5b601f821115610a2a57600081815260208120601f850160051c81016020861015611a4c5750805b601f850160051c820191505b81811015611a6b57828155600101611a58565b505050505050565b600019600383901b1c191660019190911b1790565b611a9282836119df565b6001600160401b03811115611aa957611aa9611104565b611abd81611ab78554611918565b85611a25565b6000601f821160018114611aeb5760008315611ad95750838201355b611ae38482611a73565b865550611b45565b600085815260209020601f19841690835b82811015611b1c5786850135825560209485019460019092019101611afc565b5084821015611b395760001960f88660031b161c19848701351681555b505060018360011b0185555b50505050602082013560018201555050565b60008135610fc7816113a0565b6001600160401b03831115611b7b57611b7b611104565b611b8f83611b898354611918565b83611a25565b6000601f841160018114611bbd5760008515611bab5750838201355b611bb58682611a73565b84555061051a565b600083815260209020601f19861690835b82811015611bee5786850135825560209485019460019092019101611bce565b5086821015611c0b5760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b8135611c28816113a0565b815467ffffffffffffffff19166001600160401b038216178255506020820135611c51816113a0565b81546fffffffffffffffff0000000000000000191660409190911b6fffffffffffffffff00000000000000001617905550565b611c8e82836119bf565b611c9881826119df565b6001600160401b03811115611caf57611caf611104565b611cc381611cbd8654611918565b86611a25565b6000601f821160018114611cf15760008315611cdf5750838201355b611ce98482611a73565b875550611d4b565b600086815260209020601f19841690835b82811015611d225786850135825560209485019460019092019101611d02565b5084821015611d3f5760001960f88660031b161c19848701351681555b505060018360011b0186555b505050506020810135600183015550611d73611d6a60208401846119bf565b60028301611a88565b611da3611d8260408401611b57565b600483016001600160401b0382166001600160401b03198254161781555050565b611db060608301836119df565b611dbe818360058601611b64565b5050611dd06080830160068301611c1d565b611e00611ddf60c08401611b57565b600783016001600160401b0382166001600160401b03198254161781555050565b5050565b60a081526000611e1860a083018a8c611962565b602060038a10611e3857634e487b7160e01b600052602160045260246000fd5b8381018a905288151560408501528382036060850152868252818101600588901b830182018960005b8a811015611ed057858303601f190184528135368d9003601e19018112611e8757600080fd5b8c0185810190356001600160401b03811115611ea257600080fd5b803603821315611eb157600080fd5b611ebc858284611962565b958701959450505090840190600101611e61565b50508581036080870152611ee581888a611962565b9e9d5050505050505050505050505050565b8135611f0281611659565b815490151560ff1660ff199190911617815560018082016020611f27858201866119df565b6001600160401b03811115611f3e57611f3e611104565b611f4c81611cbd8654611918565b6000601f821160018114611f7a5760008315611f685750838201355b611f728482611a73565b875550611fcf565b600086815260209020601f19841690835b82811015611fa85786850135825593870193908901908701611f8b565b5084821015611fc55760001960f88660031b161c19848701351681555b50508683881b0186555b505050505050505050565b6000808354611fe881611918565b60018281168015612000576001811461201557612044565b60ff1984168752821515830287019450612044565b8760005260208060002060005b8581101561203b5781548a820152908401908201612022565b50505082870194505b50929695505050505050565b818382376000910190815291905056fea26469706673582212208d52024926d12549d0086ff90bfe76b7655864f006263e8c517cc276562e0bdc64736f6c634300080f0033";

type PanickingMarsConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: PanickingMarsConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class PanickingMars__factory extends ContractFactory {
  constructor(...args: PanickingMarsConstructorParams) {
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
      PanickingMars & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): PanickingMars__factory {
    return super.connect(runner) as PanickingMars__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): PanickingMarsInterface {
    return new Interface(_abi) as PanickingMarsInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): PanickingMars {
    return new Contract(address, _abi, runner) as unknown as PanickingMars;
  }
}
