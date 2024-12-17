/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Contract,
  ContractFactory,
  ContractTransactionResponse,
  Interface,
} from "ethers";
import type { Signer, ContractDeployTransaction, ContractRunner } from "ethers";
import type { NonPayableOverrides } from "../common";
import type { Ibc, IbcInterface } from "../Ibc";

const _abi = [
  {
    type: "function",
    name: "ackProofKey",
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
    outputs: [
      {
        name: "proofKey",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "ackProofValue",
    inputs: [
      {
        name: "ack",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [
      {
        name: "proofValue",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "channelProofKey",
    inputs: [
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
    outputs: [
      {
        name: "proofKey",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "channelProofKeyMemory",
    inputs: [
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
    outputs: [
      {
        name: "proofKey",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "channelProofValue",
    inputs: [
      {
        name: "state",
        type: "ChannelState",
        internalType: "enum ChannelState",
      },
      {
        name: "ordering",
        type: "ChannelOrder",
        internalType: "enum ChannelOrder",
      },
      {
        name: "version",
        type: "string",
        internalType: "string",
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
      {
        name: "counterpartyChannelId",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    outputs: [
      {
        name: "proofValue",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "channelProofValueMemory",
    inputs: [
      {
        name: "state",
        type: "ChannelState",
        internalType: "enum ChannelState",
      },
      {
        name: "ordering",
        type: "ChannelOrder",
        internalType: "enum ChannelOrder",
      },
      {
        name: "version",
        type: "string",
        internalType: "string",
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
      {
        name: "counterpartyChannelId",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    outputs: [
      {
        name: "proofValue",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "packetCommitmentProofKey",
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
    outputs: [
      {
        name: "proofKey",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "packetCommitmentProofValue",
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
    outputs: [
      {
        name: "proofValue",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "parseAckData",
    inputs: [
      {
        name: "ack",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [
      {
        name: "ackData",
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
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "toStr",
    inputs: [
      {
        name: "b",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    outputs: [
      {
        name: "outStr",
        type: "string",
        internalType: "string",
      },
    ],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "toStr",
    inputs: [
      {
        name: "_number",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "outStr",
        type: "string",
        internalType: "string",
      },
    ],
    stateMutability: "pure",
  },
] as const;

const _bytecode =
  "0x611bb961003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100b35760003560e01c80634f9b0fb31161007b5780634f9b0fb31461013a5780636f6547261461014d57806374e970451461016057806391d6df7d14610181578063fb10f0a814610194578063fc8f29da146101a757600080fd5b806311a7a373146100b85780631dcd0305146100e157806325e0dd0e146100f4578063360b8cd7146101075780634b5728d114610127575b600080fd5b6100cb6100c6366004611126565b6101ba565b6040516100d891906111bc565b60405180910390f35b6100cb6100ef3660046111cf565b610236565b6100cb61010236600461124e565b610340565b61011a610115366004611341565b610435565b6040516100d89190611382565b6100cb610135366004611126565b610572565b6100cb6101483660046113a9565b6105bb565b6100cb61015b3660046114a9565b6105f2565b61017361016e366004611341565b610626565b6040519081526020016100d8565b6100cb61018f3660046111cf565b61067a565b6100cb6101a2366004611571565b61078e565b6101736101b5366004611126565b61080d565b60606101c96020830183611639565b6101d39080611659565b6101ec6101e36020860186611639565b60200135610236565b61020d6101ff606087016040880161169f565b6001600160401b031661067a565b60405160200161022094939291906116c8565b6040516020818303038152906040529050919050565b606060005b60208160ff1610801561026f5750828160ff166020811061025e5761025e611746565b1a60f81b6001600160f81b03191615155b15610286578061027e81611772565b91505061023b565b60008160ff166001600160401b038111156102a3576102a36113f4565b6040519080825280601f01601f1916602001820160405280156102cd576020820181803683370190505b50905060005b8260ff168160ff16101561033857848160ff16602081106102f6576102f6611746565b1a60f81b828260ff168151811061030f5761030f611746565b60200101906001600160f81b031916908160001a9053508061033081611772565b9150506102d3565b509392505050565b60606104276040518060a001604052808c600981111561036257610362611791565b60030b81526020018b600281111561037c5761037c611791565b60030b8152602001604051806040016040528088888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506020016103d487610236565b905281526020016103e5888a6117a7565b81526020018a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050915250610937565b9a9950505050505050505050565b6040805180820190915260008152606060208201527fcf118b5b37063214cf5ee4e00a21cbc1f63c9adff4e41aef620d6c96005c7a256104796009600185876117b4565b6040516104879291906117de565b6040518091039020146104fd5760408051808201909152600081526020810184600a856104b56002826117ee565b926104c2939291906117b4565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505091525061056b565b60408051808201909152600181526020810161056885600b866105216002826117ee565b9261052e939291906117b4565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061099f92505050565b90525b9392505050565b606061057e8280611639565b6105889080611659565b6105956101e38580611639565b6105a86101ff606087016040880161169f565b6040516020016102209493929190611805565b606083836105c884610236565b6040516020016105da93929190611889565b60405160208183030381529060405290509392505050565b6060826105fe83610236565b60405160200161060f9291906118e0565b604051602081830303815290604052905092915050565b60006002838360405161063a9291906117de565b602060405180830381855afa158015610657573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061056b9190611943565b6060816000036106a15750506040805180820190915260018152600360fc1b602082015290565b6000825b80156106cb57816106b58161195c565b92506106c49050600a8261198b565b90506106a5565b6000826001600160401b038111156106e5576106e56113f4565b6040519080825280601f01601f19166020018201604052801561070f576020820181803683370190505b509050825b801561078557610725600a8761199f565b6107309060306119b3565b60f81b8261073f6001846117ee565b8151811061074f5761074f611746565b60200101906001600160f81b031916908160001a905350610771600a8761198b565b95508061077d816119cb565b915050610714565b50949350505050565b60606108026040518060a001604052808960098111156107b0576107b0611791565b60030b81526020018860028111156107ca576107ca611791565b60030b815260200160405180604001604052808781526020016107ec87610236565b9052815260208101879052604001879052610937565b979650505050505050565b6000600261082160e0840160c0850161169f565b61083160a085016080860161169f565b61084160c0860160a0870161169f565b60026108506060880188611659565b60405161085e9291906117de565b602060405180830381855afa15801561087b573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061089e9190611943565b6040516001600160c01b031960c095861b8116602083015293851b841660288201529190931b9091166030820152603881019190915260580160408051601f19818403018152908290526108f1916119e2565b602060405180830381855afa15801561090e573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906109319190611943565b92915050565b6060600061094483610b54565b6001600160401b0381111561095b5761095b6113f4565b6040519080825280601f01601f191660200182016040528015610985576020820181803683370190505b509050600061099684602084610c5a565b82525092915050565b805160609082906000036109c3576040805160008082526020820190925290610338565b600481516109d1919061199f565b15610a225760405162461bcd60e51b815260206004820152601c60248201527f696e76616c696420626173653634206465636f64657220696e70757400000000604482015260640160405180910390fd5b60006040518060a0016040528060808152602001611b04608091399050600060048351610a4f919061198b565b610a5a9060036119f4565b90506000610a698260206119b3565b6001600160401b03811115610a8057610a806113f4565b6040519080825280601f01601f191660200182016040528015610aaa576020820181803683370190505b5090508351840151603d60ff821603610ad757600183039250613d3d61ffff821603610ad7576001830392505b50818152600183018485518101602084015b81831015610b4657600483019250825160ff8082168601511660ff808360081c168701511660061b0160ff808360101c1687015116600c1b60ff808460181c168801511660121b010190508060e81b825250600381019050610ae9565b509298975050505050505050565b6000806000610b668460000151610de4565b610b719060016119b3565b610b7b90836119b3565b9150610b8a8460200151610de4565b610b959060016119b3565b610b9f90836119b3565b9150610bb6610bb18560400151610e0d565b610e58565b610bc19060016119b3565b610bcb90836119b3565b9150600090505b836060015151811015610c2f57610c0684606001518281518110610bf857610bf8611746565b602002602001015151610e58565b610c119060016119b3565b610c1b90836119b3565b915080610c278161195c565b915050610bd2565b610c3d846080015151610e58565b610c489060016119b3565b610c5290836119b3565b949350505050565b825160009083908190839060030b15610ca357610c7b600160008488610e6d565b610c8590836119b3565b9150610c9687600001518387610e8d565b610ca090836119b3565b91505b602087015160030b15610ce657610cbe600260008488610e6d565b610cc890836119b3565b9150610cd987602001518387610e8d565b610ce390836119b3565b91505b610cf4600360028488610e6d565b610cfe90836119b3565b9150610d0f87604001518387610ead565b610d1990836119b3565b9150866060015151600014610d99575060005b866060015151811015610d9957610d47600460028488610e6d565b610d5190836119b3565b9150610d7b87606001518281518110610d6c57610d6c611746565b60200260200101518387610f60565b610d8590836119b3565b915080610d918161195c565b915050610d2c565b60808701515115610dda57610db2600560028488610e6d565b610dbc90836119b3565b9150610dcd87608001518387610f60565b610dd790836119b3565b91505b61080283836117ee565b6000808260030b1215610df95750600a919050565b6109318263ffffffff16610f6d565b919050565b600080610e1e836000015151610e58565b610e299060016119b3565b610e3390826119b3565b9050610e43836020015151610e58565b610e4e9060016119b3565b61056b90826119b3565b6000610e6382610f6d565b61093190836119b3565b6000600885026007851617610e83818585610f8a565b9695505050505050565b600083610ea46001600160401b0382168585610f8a565b95945050505050565b6000828082610ebb87610e0d565b6001600160401b03811115610ed257610ed26113f4565b6040519080825280601f01601f191660200182016040528015610efc576020820181803683370190505b50905080856000610f0f8a602085610fcd565b9050610f1c81868a610f8a565b610f2690866119b3565b9450610f46610f368460206119b3565b610f4087856119b3565b8361105c565b610f5081866119b3565b94506060935061042786866117ee565b6000610c528484846110db565b60071c600060015b82156109315760079290921c91600101610f75565b600080828401607f86165b600787901c15610fbd578060801782535060079590951c9460019182019101607f8616610f95565b8082535050600101949350505050565b825151600090839081901561101257610fea600160028387610e6d565b610ff490826119b3565b905061100586600001518286610f60565b61100f90826119b3565b90505b602086015151156110525761102a6002808387610e6d565b61103490826119b3565b905061104586602001518286610f60565b61104f90826119b3565b90505b610e8382826117ee565b8060000361106957505050565b60208111156110a257825182526110816020836119b3565b915061108e6020846119b3565b925061109b6020826117ee565b9050611069565b600060016110b18360206117ee565b6110bd90610100611af7565b6110c791906117ee565b935183518516941916939093179091525050565b8251600090816110ec828686610f8a565b905060008186018501602088015b8483101561111957805160001a825360019283019291820191016110fa565b50610802905081836119b3565b60006020828403121561113857600080fd5b81356001600160401b0381111561114e57600080fd5b820160e0818503121561056b57600080fd5b60005b8381101561117b578181015183820152602001611163565b8381111561118a576000848401525b50505050565b600081518084526111a8816020860160208601611160565b601f01601f19169290920160200192915050565b60208152600061056b6020830184611190565b6000602082840312156111e157600080fd5b5035919050565b8035600a8110610e0857600080fd5b803560038110610e0857600080fd5b60008083601f84011261121857600080fd5b5081356001600160401b0381111561122f57600080fd5b60208301915083602082850101111561124757600080fd5b9250929050565b600080600080600080600080600060c08a8c03121561126c57600080fd5b6112758a6111e8565b985061128360208b016111f7565b975060408a01356001600160401b038082111561129f57600080fd5b6112ab8d838e01611206565b909950975060608c01359150808211156112c457600080fd5b818c0191508c601f8301126112d857600080fd5b8135818111156112e757600080fd5b8d60208260051b85010111156112fc57600080fd5b6020830197508096505060808c013591508082111561131a57600080fd5b506113278c828d01611206565b9a9d999c50979a9699959894979660a00135949350505050565b6000806020838503121561135457600080fd5b82356001600160401b0381111561136a57600080fd5b61137685828601611206565b90969095509350505050565b6020815281511515602082015260006020830151604080840152610c526060840182611190565b6000806000604084860312156113be57600080fd5b83356001600160401b038111156113d457600080fd5b6113e086828701611206565b909790965060209590950135949350505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715611432576114326113f4565b604052919050565b600082601f83011261144b57600080fd5b81356001600160401b03811115611464576114646113f4565b611477601f8201601f191660200161140a565b81815284602083860101111561148c57600080fd5b816020850160208301376000918101602001919091529392505050565b600080604083850312156114bc57600080fd5b82356001600160401b038111156114d257600080fd5b6114de8582860161143a565b95602094909401359450505050565b60006001600160401b0380841115611507576115076113f4565b8360051b602061151881830161140a565b8681529350908401908084018783111561153157600080fd5b855b838110156115655780358581111561154b5760008081fd5b6115578a828a0161143a565b835250908201908201611533565b50505050509392505050565b60008060008060008060c0878903121561158a57600080fd5b611593876111e8565b95506115a1602088016111f7565b945060408701356001600160401b03808211156115bd57600080fd5b6115c98a838b0161143a565b955060608901359150808211156115df57600080fd5b818901915089601f8301126115f357600080fd5b6116028a8335602085016114ed565b9450608089013591508082111561161857600080fd5b5061162589828a0161143a565b92505060a087013590509295509295509295565b60008235603e1983360301811261164f57600080fd5b9190910192915050565b6000808335601e1984360301811261167057600080fd5b8301803591506001600160401b0382111561168a57600080fd5b60200191503681900382131561124757600080fd5b6000602082840312156116b157600080fd5b81356001600160401b038116811461056b57600080fd5b6a61636b732f706f7274732f60a81b81528385600b8301376000848201692f6368616e6e656c732f60b01b600b820152845161170b816015840160208901611160565b8082019150506a2f73657175656e6365732f60a81b60158201528351611738816020840160208801611160565b016020019695505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff81036117885761178861175c565b60010192915050565b634e487b7160e01b600052602160045260246000fd5b600061056b3684846114ed565b600080858511156117c457600080fd5b838611156117d157600080fd5b5050820193919092039150565b8183823760009101908152919050565b6000828210156118005761180061175c565b500390565b71636f6d6d69746d656e74732f706f7274732f60701b8152838560128301376000848201692f6368616e6e656c732f60b01b6012820152845161184f81601c840160208901611160565b6a2f73657175656e6365732f60a81b601c9290910191820152835161187b816027840160208801611160565b016027019695505050505050565b716368616e6e656c456e64732f706f7274732f60701b8152828460128301376000838201692f6368616e6e656c732f60b01b601282015283516118d381601c840160208801611160565b01601c0195945050505050565b716368616e6e656c456e64732f706f7274732f60701b81526000835161190d816012850160208801611160565b692f6368616e6e656c732f60b01b601291840191820152835161193781601c840160208801611160565b01601c01949350505050565b60006020828403121561195557600080fd5b5051919050565b60006001820161196e5761196e61175c565b5060010190565b634e487b7160e01b600052601260045260246000fd5b60008261199a5761199a611975565b500490565b6000826119ae576119ae611975565b500690565b600082198211156119c6576119c661175c565b500190565b6000816119da576119da61175c565b506000190190565b6000825161164f818460208701611160565b6000816000190483118215151615611a0e57611a0e61175c565b500290565b600181815b80851115611a4e578160001904821115611a3457611a3461175c565b80851615611a4157918102915b93841c9390800290611a18565b509250929050565b600082611a6557506001610931565b81611a7257506000610931565b8160018114611a885760028114611a9257611aae565b6001915050610931565b60ff841115611aa357611aa361175c565b50506001821b610931565b5060208310610133831016604e8410600b8410161715611ad1575081810a610931565b611adb8383611a13565b8060001904821115611aef57611aef61175c565b029392505050565b600061056b8383611a5656fe000000000000000000000000000000000000000000000000000000000000000000000000000000000000003e0000003f3435363738393a3b3c3d00000000000000000102030405060708090a0b0c0d0e0f101112131415161718190000000000001a1b1c1d1e1f202122232425262728292a2b2c2d2e2f303132330000000000a264697066735822122049ce89fe6a9fa569687af5e0327afd786f3148fad195a5cead501417bc839a7764736f6c634300080f0033";

type IbcConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: IbcConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class Ibc__factory extends ContractFactory {
  constructor(...args: IbcConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(overrides || {});
  }
  override deploy(overrides?: NonPayableOverrides & { from?: string }) {
    return super.deploy(overrides || {}) as Promise<
      Ibc & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): Ibc__factory {
    return super.connect(runner) as Ibc__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): IbcInterface {
    return new Interface(_abi) as IbcInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): Ibc {
    return new Contract(address, _abi, runner) as unknown as Ibc;
  }
}
