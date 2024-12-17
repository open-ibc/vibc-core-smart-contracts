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
import type { NonPayableOverrides } from "../common";
import type { Earth, EarthInterface } from "../Earth";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "_middleware",
        type: "address",
        internalType: "address",
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
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "packet",
        type: "tuple",
        internalType: "struct UniversalPacket",
        components: [
          {
            name: "srcPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "mwBitmap",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "destPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "appData",
            type: "bytes",
            internalType: "bytes",
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
    stateMutability: "view",
  },
  {
    type: "function",
    name: "authorizeChannel",
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
    name: "authorizedChannelIds",
    inputs: [
      {
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    outputs: [
      {
        name: "",
        type: "bool",
        internalType: "bool",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "generateAckPacket",
    inputs: [
      {
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "srcPortAddr",
        type: "address",
        internalType: "address",
      },
      {
        name: "appData",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [
      {
        name: "ackPacket",
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
    name: "greet",
    inputs: [
      {
        name: "destPortAddr",
        type: "address",
        internalType: "address",
      },
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "message",
        type: "bytes",
        internalType: "bytes",
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
    name: "greetWithFee",
    inputs: [
      {
        name: "destPortAddr",
        type: "address",
        internalType: "address",
      },
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "message",
        type: "bytes",
        internalType: "bytes",
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
    name: "onRecvUniversalPacket",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "packet",
        type: "tuple",
        internalType: "struct UniversalPacket",
        components: [
          {
            name: "srcPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "mwBitmap",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "destPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "appData",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
    ],
    outputs: [
      {
        name: "ackPacket",
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
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "onTimeoutUniversalPacket",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "packet",
        type: "tuple",
        internalType: "struct UniversalPacket",
        components: [
          {
            name: "srcPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "mwBitmap",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "destPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "appData",
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
    name: "onUniversalAcknowledgement",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "packet",
        type: "tuple",
        internalType: "struct UniversalPacket",
        components: [
          {
            name: "srcPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "mwBitmap",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "destPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "appData",
            type: "bytes",
            internalType: "bytes",
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
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "packet",
        type: "tuple",
        internalType: "struct UniversalPacket",
        components: [
          {
            name: "srcPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "mwBitmap",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "destPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "appData",
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
    name: "renounceOwnership",
    inputs: [],
    outputs: [],
    stateMutability: "nonpayable",
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
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "packet",
        type: "tuple",
        internalType: "struct UniversalPacket",
        components: [
          {
            name: "srcPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "mwBitmap",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "destPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "appData",
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
    name: "uch",
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
    name: "updateUch",
    inputs: [
      {
        name: "_newUch",
        type: "address",
        internalType: "address",
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
    name: "CallerNotUCH",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidUCHAddress",
    inputs: [],
  },
  {
    type: "error",
    name: "ackAddressMismatch",
    inputs: [],
  },
  {
    type: "error",
    name: "ackDataTooShort",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidChannelId",
    inputs: [],
  },
  {
    type: "error",
    name: "unauthorizedChannel",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x608060405234801561001057600080fd5b5060405161199638038061199683398101604081905261002f916100b2565b808061003a33610062565b600180546001600160a01b0319166001600160a01b0392909216919091179055506100e29050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100c457600080fd5b81516001600160a01b03811681146100db57600080fd5b9392505050565b6118a5806100f16000396000f3fe6080604052600436106100f75760003560e01c80635b7615851161008a57806392dfa3921161005957806392dfa392146102f8578063d24ba02414610318578063f12b758a14610338578063f2fde38b1461035857600080fd5b80635b7615851461026a578063715018a614610298578063866f3f97146102ad5780638da5cb5b146102da57600080fd5b8063400d9f5d116100c6578063400d9f5d146101cd5780634252ae9b146101ed5780634eeb73911461021c578063588152ca1461024a57600080fd5b80631b2f37d0146101035780632466911c146101255780632cc422d9146101555780632eed7c701461018d57600080fd5b366100fe57005b600080fd5b34801561010f57600080fd5b5061012361011e366004610e23565b610378565b005b610138610133366004610f7f565b6103c9565b6040516001600160401b0390911681526020015b60405180910390f35b34801561016157600080fd5b50600154610175906001600160a01b031681565b6040516001600160a01b03909116815260200161014c565b34801561019957600080fd5b506101bd6101a8366004611012565b60026020526000908152604090205460ff1681565b604051901515815260200161014c565b3480156101d957600080fd5b506101236101e836600461102b565b610465565b3480156101f957600080fd5b5061020d610208366004611012565b610540565b60405161014c93929190611130565b34801561022857600080fd5b5061023c610237366004611012565b6106e4565b60405161014c929190611165565b34801561025657600080fd5b5061012361026536600461127d565b6107d2565b34801561027657600080fd5b5061028a61028536600461102b565b610980565b60405161014c9291906112f0565b3480156102a457600080fd5b50610123610b08565b3480156102b957600080fd5b506102cd6102c8366004611314565b610b1c565b60405161014c919061136d565b3480156102e657600080fd5b506000546001600160a01b0316610175565b34801561030457600080fd5b50610123610313366004611012565b610b77565b34801561032457600080fd5b50610123610333366004611380565b610b8b565b34801561034457600080fd5b5061023c610353366004611012565b610c1b565b34801561036457600080fd5b50610123610373366004610e23565b610c2b565b610380610ca6565b6001600160a01b0381166103a75760405163a944796960e01b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6001546000906001600160a01b031663462fdf8334896103ef8c6001600160a01b031690565b8a8a8a8a8a6040518963ffffffff1660e01b8152600401610416979695949392919061143d565b60206040518083038185885af1158015610434573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250810190610459919061148a565b98975050505050505050565b6001546001600160a01b0316331461049057604051631323efc560e01b815260040160405180910390fd5b600082815260026020526040902054829060ff166104c157604051630100e70560e51b815260040160405180910390fd5b60056040518060400160405280858152602001846104de906114a7565b9052815460018181018455600093845260209384902083516005909302019182558383015180519183019182559384015160028301556040840151600383015560608401519293919260048401906105369082611538565b5050505050505050565b6004818154811061055057600080fd5b9060005260206000209060070201600091509050806000015490806001016040518060800160405290816000820154815260200160018201548152602001600282015481526020016003820180546105a7906114b3565b80601f01602080910402602001604051908101604052809291908181526020018280546105d3906114b3565b80156106205780601f106105f557610100808354040283529160200191610620565b820191906000526020600020905b81548152906001019060200180831161060357829003601f168201915b5050509190925250506040805180820190915260058401805460ff1615158252600685018054949594929350909160208401919061065d906114b3565b80601f0160208091040260200160405190810160405280929190818152602001828054610689906114b3565b80156106d65780601f106106ab576101008083540402835291602001916106d6565b820191906000526020600020905b8154815290600101906020018083116106b957829003601f168201915b505050505081525050905083565b600581815481106106f457600080fd5b90600052602060002090600502016000915090508060000154908060010160405180608001604052908160008201548152602001600182015481526020016002820154815260200160038201805461074b906114b3565b80601f0160208091040260200160405190810160405280929190818152602001828054610777906114b3565b80156107c45780601f10610799576101008083540402835291602001916107c4565b820191906000526020600020905b8154815290600101906020018083116107a757829003601f168201915b505050505081525050905082565b6001546001600160a01b031633146107fd57604051631323efc560e01b815260040160405180910390fd5b600083815260026020526040902054839060ff1661082e57604051630100e70560e51b815260040160405180910390fd5b601461083d60208401846115f7565b9050101561085e5760405163503b43db60e01b815260040160405180910390fd5b600061086d60208401846115f7565b61087c9160149160009161163d565b61088591611667565b60601c9050806001600160a01b03166108a18560400151610d00565b6001600160a01b0316146108c857604051631863a42d60e31b815260040160405180910390fd5b60046040518060600160405280878152602001868152602001856108eb906116aa565b9052815460018181018455600093845260209384902083516007909302019182558383015180519183019182559384015160028301556040840151600383015560608401519293919260048401906109439082611538565b5050506040820151805160058301805460ff1916911515919091178155602082015160068401906109749082611538565b50505050505050505050565b6040805180820190915260008152606060208201526001546000906001600160a01b031633146109c357604051631323efc560e01b815260040160405180910390fd5b600084815260026020526040902054849060ff166109f457604051630100e70560e51b815260040160405180910390fd5b6003604051806040016040528087815260200186610a11906114a7565b905281546001818101845560009384526020938490208351600590930201918255838301518051918301918255938401516002830155604084015160038301556060840151929391926004840190610a699082611538565b5050505050306001600160a01b031663866f3f9786610a8b8760000135610d00565b610a9860608901896115f7565b6040518563ffffffff1660e01b8152600401610ab79493929190611702565b600060405180830381865afa158015610ad4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610afc919081019061172d565b95600095509350505050565b610b10610ca6565b610b1a6000610d11565b565b604080518082019091526000815260606020820152604051806040016040528060011515815260200130868686604051602001610b5c94939291906117ea565b60408051601f19818403018152919052905295945050505050565b610b7f610ca6565b610b8881610d61565b50565b6001546001600160a01b0316631f3a583085610bad886001600160a01b031690565b8686866040518663ffffffff1660e01b8152600401610bd0959493929190611834565b6020604051808303816000875af1158015610bef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c13919061148a565b505050505050565b600381815481106106f457600080fd5b610c33610ca6565b6001600160a01b038116610c9d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b610b8881610d11565b6000546001600160a01b03163314610b1a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610c94565b6000610d0b82610d9a565b92915050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80610d7f57604051631d6f5e8b60e01b815260040160405180910390fd5b6000908152600260205260409020805460ff19166001179055565b60006001600160a01b03821115610e035760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e20316044820152663630206269747360c81b6064820152608401610c94565b5090565b80356001600160a01b0381168114610e1e57600080fd5b919050565b600060208284031215610e3557600080fd5b610e3e82610e07565b9392505050565b60008083601f840112610e5757600080fd5b5081356001600160401b03811115610e6e57600080fd5b602083019150836020828501011115610e8657600080fd5b9250929050565b6001600160401b0381168114610b8857600080fd5b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715610eda57610eda610ea2565b60405290565b604051601f8201601f191681016001600160401b0381118282101715610f0857610f08610ea2565b604052919050565b600082601f830112610f2157600080fd5b604051604081018181106001600160401b0382111715610f4357610f43610ea2565b8060405250806040840185811115610f5a57600080fd5b845b81811015610f74578035835260209283019201610f5c565b509195945050505050565b6000806000806000806000610100888a031215610f9b57600080fd5b610fa488610e07565b96506020880135955060408801356001600160401b03811115610fc657600080fd5b610fd28a828b01610e45565b9096509450506060880135610fe681610e8d565b9250610ff58960808a01610f10565b91506110048960c08a01610f10565b905092959891949750929550565b60006020828403121561102457600080fd5b5035919050565b6000806040838503121561103e57600080fd5b8235915060208301356001600160401b0381111561105b57600080fd5b83016080818603121561106d57600080fd5b809150509250929050565b60005b8381101561109357818101518382015260200161107b565b838111156110a2576000848401525b50505050565b600081518084526110c0816020860160208601611078565b601f01601f19169290920160200192915050565b805182526020810151602083015260408101516040830152600060608201516080606085015261110760808501826110a8565b949350505050565b805115158252600060208201516040602085015261110760408501826110a8565b83815260606020820152600061114960608301856110d4565b828103604084015261115b818561110f565b9695505050505050565b82815260406020820152600061110760408301846110d4565b60006001600160401b0382111561119757611197610ea2565b50601f01601f191660200190565b600082601f8301126111b657600080fd5b81356111c96111c48261117e565b610ee0565b8181528460208386010111156111de57600080fd5b816020850160208301376000918101602001919091529392505050565b60006080828403121561120d57600080fd5b604051608081016001600160401b03828210818311171561123057611230610ea2565b81604052829350843583526020850135602084015260408501356040840152606085013591508082111561126357600080fd5b50611270858286016111a5565b6060830152505092915050565b60008060006060848603121561129257600080fd5b8335925060208401356001600160401b03808211156112b057600080fd5b6112bc878388016111fb565b935060408601359150808211156112d257600080fd5b508401604081870312156112e557600080fd5b809150509250925092565b604081526000611303604083018561110f565b905082151560208301529392505050565b6000806000806060858703121561132a57600080fd5b8435935061133a60208601610e07565b925060408501356001600160401b0381111561135557600080fd5b61136187828801610e45565b95989497509550505050565b602081526000610e3e602083018461110f565b60008060008060006080868803121561139857600080fd5b6113a186610e07565b94506020860135935060408601356001600160401b038111156113c357600080fd5b6113cf88828901610e45565b90945092505060608601356113e381610e8d565b809150509295509295909350565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8060005b60028110156110a257815184526020938401939091019060010161141e565b600061010089835288602084015280604084015261145e818401888a6113f1565b9150506001600160401b038516606083015261147d608083018561141a565b61045960c083018461141a565b60006020828403121561149c57600080fd5b8151610e3e81610e8d565b6000610d0b36836111fb565b600181811c908216806114c757607f821691505b6020821081036114e757634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561153357600081815260208120601f850160051c810160208610156115145750805b601f850160051c820191505b81811015610c1357828155600101611520565b505050565b81516001600160401b0381111561155157611551610ea2565b6115658161155f84546114b3565b846114ed565b602080601f83116001811461159a57600084156115825750858301515b600019600386901b1c1916600185901b178555610c13565b600085815260208120601f198616915b828110156115c9578886015182559484019460019091019084016115aa565b50858210156115e75787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000808335601e1984360301811261160e57600080fd5b8301803591506001600160401b0382111561162857600080fd5b602001915036819003821315610e8657600080fd5b6000808585111561164d57600080fd5b8386111561165a57600080fd5b5050820193919092039150565b6bffffffffffffffffffffffff1981358181169160148510156116945780818660140360031b1b83161692505b505092915050565b8015158114610b8857600080fd5b6000604082360312156116bc57600080fd5b6116c4610eb8565b82356116cf8161169c565b815260208301356001600160401b038111156116ea57600080fd5b6116f6368286016111a5565b60208301525092915050565b8481526001600160a01b038416602082015260606040820181905260009061115b90830184866113f1565b6000602080838503121561174057600080fd5b82516001600160401b038082111561175757600080fd5b908401906040828703121561176b57600080fd5b611773610eb8565b825161177e8161169c565b8152828401518281111561179157600080fd5b80840193505086601f8401126117a657600080fd5b825191506117b66111c48361117e565b82815287858486010111156117ca57600080fd5b6117d983868301878701611078565b938101939093525090949350505050565b60006bffffffffffffffffffffffff19808760601b168352808660601b166014840152506361636b2d60e01b60288301528284602c8401375060009101602c019081529392505050565b8581528460208201526080604082015260006118546080830185876113f1565b90506001600160401b0383166060830152969550505050505056fea26469706673582212207d54b34347e3befb611a010b63ecb658fe938bd48fd38f4a617aa5af84d3252764736f6c634300080f0033";

type EarthConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: EarthConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class Earth__factory extends ContractFactory {
  constructor(...args: EarthConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    _middleware: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(_middleware, overrides || {});
  }
  override deploy(
    _middleware: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(_middleware, overrides || {}) as Promise<
      Earth & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): Earth__factory {
    return super.connect(runner) as Earth__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): EarthInterface {
    return new Interface(_abi) as EarthInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): Earth {
    return new Contract(address, _abi, runner) as unknown as Earth;
  }
}
