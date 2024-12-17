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
import type { CrossL2Prover, CrossL2ProverInterface } from "../CrossL2Prover";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "verifier_",
        type: "address",
        internalType: "contract ISignatureVerifier",
      },
      {
        name: "clientType_",
        type: "string",
        internalType: "string",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "LIGHT_CLIENT_TYPE",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "uint8",
        internalType: "enum LightClientType",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "clientType",
    inputs: [],
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
    name: "getState",
    inputs: [
      {
        name: "height",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "peptideAppHashes",
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
        type: "uint256",
        internalType: "uint256",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "updateClient",
    inputs: [
      {
        name: "proof",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "peptideHeight",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "peptideAppHash",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "validateLog",
    inputs: [
      {
        name: "logIndex",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "proof",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [
      {
        name: "chainId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "emittingContract",
        type: "address",
        internalType: "address",
      },
      {
        name: "topics",
        type: "bytes[]",
        internalType: "bytes[]",
      },
      {
        name: "unindexedData",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "validateReceipt",
    inputs: [
      {
        name: "proof",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [
      {
        name: "srcChainID",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "receiptRLP",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "verifier",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "contract ISignatureVerifier",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "verifyMembership",
    inputs: [
      {
        name: "appHash",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "key",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "value",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "proofs",
        type: "tuple",
        internalType: "struct Ics23Proof",
        components: [
          {
            name: "proof",
            type: "tuple[]",
            internalType: "struct OpIcs23Proof[]",
            components: [
              {
                name: "path",
                type: "tuple[]",
                internalType: "struct OpIcs23ProofPath[]",
                components: [
                  {
                    name: "prefix",
                    type: "bytes",
                    internalType: "bytes",
                  },
                  {
                    name: "suffix",
                    type: "bytes",
                    internalType: "bytes",
                  },
                ],
              },
              {
                name: "key",
                type: "bytes",
                internalType: "bytes",
              },
              {
                name: "value",
                type: "bytes",
                internalType: "bytes",
              },
              {
                name: "prefix",
                type: "bytes",
                internalType: "bytes",
              },
            ],
          },
          {
            name: "height",
            type: "uint256",
            internalType: "uint256",
          },
        ],
      },
    ],
    outputs: [],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "verifyNonMembership",
    inputs: [
      {
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "",
        type: "tuple",
        internalType: "struct Ics23Proof",
        components: [
          {
            name: "proof",
            type: "tuple[]",
            internalType: "struct OpIcs23Proof[]",
            components: [
              {
                name: "path",
                type: "tuple[]",
                internalType: "struct OpIcs23ProofPath[]",
                components: [
                  {
                    name: "prefix",
                    type: "bytes",
                    internalType: "bytes",
                  },
                  {
                    name: "suffix",
                    type: "bytes",
                    internalType: "bytes",
                  },
                ],
              },
              {
                name: "key",
                type: "bytes",
                internalType: "bytes",
              },
              {
                name: "value",
                type: "bytes",
                internalType: "bytes",
              },
              {
                name: "prefix",
                type: "bytes",
                internalType: "bytes",
              },
            ],
          },
          {
            name: "height",
            type: "uint256",
            internalType: "uint256",
          },
        ],
      },
    ],
    outputs: [],
    stateMutability: "pure",
  },
  {
    type: "error",
    name: "CannotUpdateClientWithDifferentAppHash",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidAppHash",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidIbcStateProof",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidL1BlockHash",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidL1BlockNumber",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidPacketProof",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidProofKey",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidProofValue",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidRLPEncodedL1BlockNumber",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidRLPEncodedL1StateRoot",
    inputs: [],
  },
  {
    type: "error",
    name: "MethodNotImplemented",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x60a06040523480156200001157600080fd5b50604051620034c7380380620034c783398101604081905262000034916200006e565b6001600160a01b03821660805260006200004f8282620001fd565b505050620002c9565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156200008257600080fd5b82516001600160a01b03811681146200009a57600080fd5b602084810151919350906001600160401b0380821115620000ba57600080fd5b818601915086601f830112620000cf57600080fd5b815181811115620000e457620000e462000058565b604051601f8201601f19908116603f011681019083821181831017156200010f576200010f62000058565b8160405282815289868487010111156200012857600080fd5b600093505b828410156200014c57848401860151818501870152928501926200012d565b828411156200015e5760008684830101525b8096505050505050509250929050565b600181811c908216806200018357607f821691505b602082108103620001a457634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620001f857600081815260208120601f850160051c81016020861015620001d35750805b601f850160051c820191505b81811015620001f457828155600101620001df565b5050505b505050565b81516001600160401b0381111562000219576200021962000058565b62000231816200022a84546200016e565b84620001aa565b602080601f831160018114620002695760008415620002505750858301515b600019600386901b1c1916600185901b178555620001f4565b600085815260208120601f198616915b828110156200029a5788860151825594840194600190910190840162000279565b5085821015620002b95787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6080516131dc620002eb6000396000818160bd015261103b01526131dc6000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806357c1c5f41161006657806357c1c5f414610151578063b3768f0d14610166578063c2f0329f1461017b578063c67e15f71461018e578063fc396ddb146101ae57600080fd5b80632a6ded74146100a35780632b7ac3f3146100b85780632cd78e77146100fc57806344c9af281461011d57806349ff245e1461013e575b600080fd5b6100b66100b1366004612386565b6101d1565b005b6100df7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b61010f61010a3660046123f8565b6101ea565b6040516100f3929190612491565b61013061012b3660046124b2565b61036b565b6040519081526020016100f3565b6100b661014c3660046124cb565b610381565b610159600281565b6040516100f39190612531565b61016e610393565b6040516100f39190612559565b6100b661018936600461266b565b610421565b61013061019c3660046124b2565b60016020526000908152604090205481565b6101c16101bc3660046126e3565b6105e3565b6040516100f3949392919061272e565b604051632974974360e01b815260040160405180910390fd5b60006060818080808080610200898b018b612ad5565b955095509550955095509550306001600160a01b031663c2f0329f6102416001896020015161022f9190612bb0565b60009081526001602052604090205490565b60001b6102e2856000805461025590612bc7565b80601f016020809104026020016040519081016040528092919081815260200182805461028190612bc7565b80156102ce5780601f106102a3576101008083540402835291602001916102ce565b820191906000526020600020905b8154815290600101906020018083116102b157829003601f168201915b5050505050886001600160401b031661071e565b60408051602081018a9052016040516020818303038152906040528a6040518563ffffffff1660e01b815260040161031d9493929190612bfb565b60006040518083038186803b15801561033557600080fd5b505afa158015610349573d6000803e3d6000fd5b5050505087610359828787610755565b975097505050505050505b9250929050565b6000818152600160205260408120545b92915050565b61038d84848484610fea565b50505050565b600080546103a090612bc7565b80601f01602080910402602001604051908101604052809291908181526020018280546103cc90612bc7565b80156104195780601f106103ee57610100808354040283529160200191610419565b820191906000526020600020905b8154815290600101906020018083116103fc57829003601f168201915b505050505081565b61042b8180612d5d565b600081811061043c5761043c612da6565b905060200281019061044e9190612dbc565b61045c906020810190612ddc565b60405161046a929190612e22565b60405180910390208380519060200120146104985760405163026a287560e51b815260040160405180910390fd5b6104a28180612d5d565b60008181106104b3576104b3612da6565b90506020028101906104c59190612dbc565b6104d3906040810190612ddc565b6040516104e1929190612e22565b604051809103902082805190602001201461050f576040516310d9300f60e11b815260040160405180910390fd5b61054461051c8280612d5d565b600081811061052d5761052d612da6565b905060200281019061053f9190612dbc565b6110ec565b61054e8280612d5d565b600181811061055f5761055f612da6565b90506020028101906105719190612dbc565b61057f906040810190612ddc565b61058891612e32565b146105a657604051636589f0e160e11b815260040160405180910390fd5b6105c46105b38280612d5d565b600181811061052d5761052d612da6565b841461038d576040516392cb8fbb60e01b815260040160405180910390fd5b60008060608060606105f587876101ea565b8092508196505050600061061f61061a8360018086516106159190612bb0565b611310565b611452565b9050600061065a6106498360038151811061063c5761063c612da6565b6020026020010151611460565b8b8151811061063c5761063c612da6565b905061067f8160008151811061067257610672612da6565b6020026020010151611683565b61068890612e50565b60001c955060006106a58260018151811061063c5761063c612da6565b90506106bd8260028151811061067257610672612da6565b945060005b8151811015610710576106e082828151811061067257610672612da6565b8782815181106106f2576106f2612da6565b6020026020010181905250808061070890612e74565b9150506106c2565b505050505093509350935093565b6060838361072b846117af565b60405160200161073d93929190612e8d565b60405160208183030381529060405290509392505050565b606060008451116107a55760405162461bcd60e51b81526020600482015260156024820152744d65726b6c65547269653a20656d707479206b657960581b60448201526064015b60405180910390fd5b60006107b0846118ba565b905060006107bd8661199e565b90506000846040516020016107d491815260200190565b60405160208183030381529060405290506000805b8451811015610f8c57600085828151811061080657610806612da6565b6020026020010151905084518311156108785760405162461bcd60e51b815260206004820152602e60248201527f4d65726b6c65547269653a206b657920696e646578206578636565647320746f60448201526d0e8c2d840d6caf240d8cadccee8d60931b606482015260840161079c565b8260000361091757805180516020918201206040516108c6926108a092910190815260200190565b604051602081830303815290604052858051602091820120825192909101919091201490565b6109125760405162461bcd60e51b815260206004820152601d60248201527f4d65726b6c65547269653a20696e76616c696420726f6f742068617368000000604482015260640161079c565b610a0d565b80515160201161099d5780518051602091820120604051610941926108a092910190815260200190565b6109125760405162461bcd60e51b815260206004820152602760248201527f4d65726b6c65547269653a20696e76616c6964206c6172676520696e7465726e6044820152660c2d840d0c2e6d60cb1b606482015260840161079c565b805184516020808701919091208251919092012014610a0d5760405162461bcd60e51b815260206004820152602660248201527f4d65726b6c65547269653a20696e76616c696420696e7465726e616c206e6f646044820152650ca40d0c2e6d60d31b606482015260840161079c565b610a1960106001612f11565b81602001515103610bb45784518303610b4c57610a46816020015160108151811061067257610672612da6565b96506000875111610abf5760405162461bcd60e51b815260206004820152603b60248201527f4d65726b6c65547269653a2076616c7565206c656e677468206d75737420626560448201527f2067726561746572207468616e207a65726f20286272616e6368290000000000606482015260840161079c565b60018651610acd9190612bb0565b8214610b415760405162461bcd60e51b815260206004820152603a60248201527f4d65726b6c65547269653a2076616c7565206e6f6465206d757374206265206c60448201527f617374206e6f646520696e2070726f6f6620286272616e636829000000000000606482015260840161079c565b505050505050610fe3565b6000858481518110610b6057610b60612da6565b602001015160f81c60f81b60f81c9050600082602001518260ff1681518110610b8b57610b8b612da6565b60200260200101519050610b9e81611a01565b9550610bab600186612f11565b94505050610f79565b600281602001515103610f20576000610bcc82611a26565b9050600081600081518110610be357610be3612da6565b016020015160f81c90506000610bfa600283612f3f565b610c05906002612f61565b90506000610c16848360ff16611a4a565b90506000610c248a89611a4a565b90506000610c328383611a7b565b905080835114610caa5760405162461bcd60e51b815260206004820152603a60248201527f4d65726b6c65547269653a20706174682072656d61696e646572206d7573742060448201527f736861726520616c6c206e6962626c65732077697468206b6579000000000000606482015260840161079c565b60ff851660021480610cbf575060ff85166003145b15610e605780825114610d3a5760405162461bcd60e51b815260206004820152603d60248201527f4d65726b6c65547269653a206b65792072656d61696e646572206d757374206260448201527f65206964656e746963616c20746f20706174682072656d61696e646572000000606482015260840161079c565b610d54876020015160018151811061067257610672612da6565b9c5060008d5111610dcd5760405162461bcd60e51b815260206004820152603960248201527f4d65726b6c65547269653a2076616c7565206c656e677468206d75737420626560448201527f2067726561746572207468616e207a65726f20286c6561662900000000000000606482015260840161079c565b60018c51610ddb9190612bb0565b8814610e4f5760405162461bcd60e51b815260206004820152603860248201527f4d65726b6c65547269653a2076616c7565206e6f6465206d757374206265206c60448201527f617374206e6f646520696e2070726f6f6620286c656166290000000000000000606482015260840161079c565b505050505050505050505050610fe3565b60ff85161580610e73575060ff85166001145b15610eb257610e9f8760200151600181518110610e9257610e92612da6565b6020026020010151611a01565b9950610eab818a612f11565b9850610f15565b60405162461bcd60e51b815260206004820152603260248201527f4d65726b6c65547269653a2072656365697665642061206e6f64652077697468604482015271040c2dc40eadcd6dcdeeedc40e0e4caccd2f60731b606482015260840161079c565b505050505050610f79565b60405162461bcd60e51b815260206004820152602860248201527f4d65726b6c65547269653a20726563656976656420616e20756e706172736561604482015267626c65206e6f646560c01b606482015260840161079c565b5080610f8481612e74565b9150506107e9565b5060405162461bcd60e51b815260206004820152602560248201527f4d65726b6c65547269653a2072616e206f7574206f662070726f6f6620656c656044820152646d656e747360d81b606482015260840161079c565b9392505050565b6000828152600160205260409020541561103157600082815260016020526040902054811461102c57604051631549535560e01b815260040160405180910390fd5b61038d565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663e8d90039838361107060206000898b612f84565b61107991612e32565b611086886020818c612f84565b6040518663ffffffff1660e01b81526004016110a6959493929190612fae565b600060405180830381600087803b1580156110c057600080fd5b505af11580156110d4573d6000803e3d6000fd5b50505060009283525060016020526040909120555050565b60008060026110fe6040850185612ddc565b60405161110c929190612e22565b602060405180830381855afa158015611129573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061114c9190612ff2565b9050600261115d6060850185612ddc565b61117461116d6020880188612ddc565b9050611aff565b6111816020880188612ddc565b61118b6020611aff565b876040516020016111a2979695949392919061300b565b60408051601f19818403018152908290526111bc91613058565b602060405180830381855afa1580156111d9573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906111fc9190612ff2565b915060005b61120b8480612d5d565b90508110156113095760026112208580612d5d565b8381811061123057611230612da6565b9050602002810190611242919061306a565b61124c9080612ddc565b856112578880612d5d565b8681811061126757611267612da6565b9050602002810190611279919061306a565b611287906020810190612ddc565b60405160200161129b959493929190613080565b60408051601f19818403018152908290526112b591613058565b602060405180830381855afa1580156112d2573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906112f59190612ff2565b92508061130181612e74565b915050611201565b5050919050565b60608182601f0110156113565760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b604482015260640161079c565b8282840110156113995760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b604482015260640161079c565b818301845110156113e05760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b604482015260640161079c565b6060821580156113ff5760405191506000825260208201604052611449565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015611438578051835260209283019201611420565b5050858452601f01601f1916604052505b50949350505050565b606061037b61146083611b6e565b6060600080600061147085611bc1565b91945092509050600181600181111561148b5761148b61251b565b146114fe5760405162461bcd60e51b815260206004820152603860248201527f524c505265616465723a206465636f646564206974656d207479706520666f7260448201527f206c697374206973206e6f742061206c697374206974656d0000000000000000606482015260840161079c565b845161150a8385612f11565b146115725760405162461bcd60e51b815260206004820152603260248201527f524c505265616465723a206c697374206974656d2068617320616e20696e76616044820152713634b2103230ba30903932b6b0b4b73232b960711b606482015260840161079c565b604080516020808252610420820190925290816020015b60408051808201909152600080825260208201528152602001906001900390816115895790505093506000835b8651811015611677576000806115fc6040518060400160405280858c600001516115e09190612bb0565b8152602001858c602001516115f59190612f11565b9052611bc1565b5091509150604051806040016040528083836116189190612f11565b8152602001848b6020015161162d9190612f11565b81525088858151811061164257611642612da6565b6020908102919091010152611658600185612f11565b93506116648183612f11565b61166e9084612f11565b925050506115b6565b50845250919392505050565b6060600080600061169385611bc1565b9194509250905060008160018111156116ae576116ae61251b565b146117215760405162461bcd60e51b815260206004820152603960248201527f524c505265616465723a206465636f646564206974656d207479706520666f7260448201527f206279746573206973206e6f7420612064617461206974656d00000000000000606482015260840161079c565b61172b8284612f11565b8551146117975760405162461bcd60e51b815260206004820152603460248201527f524c505265616465723a2062797465732076616c756520636f6e7461696e732060448201527330b71034b73b30b634b2103932b6b0b4b73232b960611b606482015260840161079c565b6117a685602001518484612284565b95945050505050565b6060816000036117d65750506040805180820190915260018152600360fc1b602082015290565b6000825b801561180057816117ea81612e74565b92506117f99050600a826130a8565b90506117da565b6000826001600160401b0381111561181a5761181a61256c565b6040519080825280601f01601f191660200182016040528015611844576020820181803683370190505b509050825b80156114495761185a600a876130bc565b611865906030612f11565b60f81b82611874600184612bb0565b8151811061188457611884612da6565b60200101906001600160f81b031916908160001a9053506118a6600a876130a8565b9550806118b2816130d0565b915050611849565b8051606090806001600160401b038111156118d7576118d761256c565b60405190808252806020026020018201604052801561191c57816020015b60408051808201909152606080825260208201528152602001906001900390816118f55790505b50915060005b8181101561130957604051806040016040528085838151811061194757611947612da6565b6020026020010151815260200161197686848151811061196957611969612da6565b6020026020010151611452565b81525083828151811061198b5761198b612da6565b6020908102919091010152600101611922565b606080604051905082518060011b603f8101601f1916830160405280835250602084016020830160005b838110156119f6578060011b82018184015160001a8060041c8253600f8116600183015350506001016119c8565b509295945050505050565b60606020826000015110611a1d57611a1882611683565b61037b565b61037b82612317565b606061037b611a45836020015160008151811061067257610672612da6565b61199e565b606082518210611a69575060408051602081019091526000815261037b565b610fe383838486516106159190612bb0565b6000808251845110611a8e578251611a91565b83515b90505b8082108015611ae85750828281518110611ab057611ab0612da6565b602001015160f81c60f81b6001600160f81b031916848381518110611ad757611ad7612da6565b01602001516001600160f81b031916145b15611af857816001019150611a94565b5092915050565b6060805b60808310611b41578083607f1660801760f81b604051602001611b279291906130e7565b60408051601f198184030190525260079290921c91611b03565b808360f81b604051602001611b579291906130e7565b604051602081830303815290604052915050919050565b60408051808201909152600080825260208201526000825111611ba35760405162461bcd60e51b815260040161079c90613116565b50604080518082019091528151815260209182019181019190915290565b600080600080846000015111611be95760405162461bcd60e51b815260040161079c90613116565b6020840151805160001a607f8111611c0e57600060016000945094509450505061227d565b60b78111611d6b576000611c23608083612bb0565b905080876000015111611ca35760405162461bcd60e51b815260206004820152604e602482015260008051602061318783398151915260448201527f742062652067726561746572207468616e20737472696e67206c656e6774682060648201526d2873686f727420737472696e672960901b608482015260a40161079c565b6001838101516001600160f81b0319169082141580611cd05750600160ff1b6001600160f81b0319821610155b611d585760405162461bcd60e51b815260206004820152604d60248201527f524c505265616465723a20696e76616c6964207072656669782c2073696e676c60448201527f652062797465203c203078383020617265206e6f74207072656669786564202860648201526c73686f727420737472696e672960981b608482015260a40161079c565b506001955093506000925061227d915050565b60bf8111611fac576000611d8060b783612bb0565b905080876000015111611e035760405162461bcd60e51b8152602060048201526051602482015260008051602061318783398151915260448201527f74206265203e207468616e206c656e677468206f6620737472696e67206c656e60648201527067746820286c6f6e6720737472696e672960781b608482015260a40161079c565b60018301516001600160f81b0319166000819003611e8a5760405162461bcd60e51b815260206004820152604a602482015260008051602061318783398151915260448201527f74206e6f74206861766520616e79206c656164696e67207a65726f7320286c6f6064820152696e6720737472696e672960b01b608482015260a40161079c565b600184015160088302610100031c60378111611f0d5760405162461bcd60e51b8152602060048201526048602482015260008051602061318783398151915260448201527f742062652067726561746572207468616e20353520627974657320286c6f6e6760648201526720737472696e672960c01b608482015260a40161079c565b611f178184612f11565b895111611f8f5760405162461bcd60e51b815260206004820152604c602482015260008051602061318783398151915260448201527f742062652067726561746572207468616e20746f74616c206c656e677468202860648201526b6c6f6e6720737472696e672960a01b608482015260a40161079c565b611f9a836001612f11565b975095506000945061227d9350505050565b60f7811161204e576000611fc160c083612bb0565b90508087600001511161203d5760405162461bcd60e51b815260206004820152604a602482015260008051602061318783398151915260448201527f742062652067726561746572207468616e206c697374206c656e677468202873606482015269686f7274206c6973742960b01b608482015260a40161079c565b60019550935084925061227d915050565b600061205b60f783612bb0565b9050808760000151116120da5760405162461bcd60e51b815260206004820152604d602482015260008051602061318783398151915260448201527f74206265203e207468616e206c656e677468206f66206c697374206c656e677460648201526c6820286c6f6e67206c6973742960981b608482015260a40161079c565b60018301516001600160f81b031916600081900361215f5760405162461bcd60e51b8152602060048201526048602482015260008051602061318783398151915260448201527f74206e6f74206861766520616e79206c656164696e67207a65726f7320286c6f6064820152676e67206c6973742960c01b608482015260a40161079c565b600184015160088302610100031c603781116121e05760405162461bcd60e51b8152602060048201526046602482015260008051602061318783398151915260448201527f742062652067726561746572207468616e20353520627974657320286c6f6e67606482015265206c6973742960d01b608482015260a40161079c565b6121ea8184612f11565b8951116122605760405162461bcd60e51b815260206004820152604a602482015260008051602061318783398151915260448201527f742062652067726561746572207468616e20746f74616c206c656e67746820286064820152696c6f6e67206c6973742960b01b608482015260a40161079c565b61226b836001612f11565b975095506001945061227d9350505050565b9193909250565b6060816001600160401b0381111561229e5761229e61256c565b6040519080825280601f01601f1916602001820160405280156122c8576020820181803683370190505b5090508115610fe35760006122dd8486612f11565b90506020820160005b848110156122fe5782810151828201526020016122e6565b8481111561230d576000858301525b5050509392505050565b606061037b826020015160008460000151612284565b60008083601f84011261233f57600080fd5b5081356001600160401b0381111561235657600080fd5b60208301915083602082850101111561036457600080fd5b60006040828403121561238057600080fd5b50919050565b6000806000806060858703121561239c57600080fd5b8435935060208501356001600160401b03808211156123ba57600080fd5b6123c68883890161232d565b909550935060408701359150808211156123df57600080fd5b506123ec8782880161236e565b91505092959194509250565b6000806020838503121561240b57600080fd5b82356001600160401b0381111561242157600080fd5b61242d8582860161232d565b90969095509350505050565b60005b8381101561245457818101518382015260200161243c565b8381111561038d5750506000910152565b6000815180845261247d816020860160208601612439565b601f01601f19169290920160200192915050565b8281526040602082015260006124aa6040830184612465565b949350505050565b6000602082840312156124c457600080fd5b5035919050565b600080600080606085870312156124e157600080fd5b84356001600160401b038111156124f757600080fd5b6125038782880161232d565b90989097506020870135966040013595509350505050565b634e487b7160e01b600052602160045260246000fd5b602081016004831061255357634e487b7160e01b600052602160045260246000fd5b91905290565b602081526000610fe36020830184612465565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156125a4576125a461256c565b60405290565b604051608081016001600160401b03811182821017156125a4576125a461256c565b604051601f8201601f191681016001600160401b03811182821017156125f4576125f461256c565b604052919050565b600082601f83011261260d57600080fd5b81356001600160401b038111156126265761262661256c565b612639601f8201601f19166020016125cc565b81815284602083860101111561264e57600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000806080858703121561268157600080fd5b8435935060208501356001600160401b038082111561269f57600080fd5b6126ab888389016125fc565b945060408701359150808211156126c157600080fd5b6126cd888389016125fc565b935060608701359150808211156123df57600080fd5b6000806000604084860312156126f857600080fd5b8335925060208401356001600160401b0381111561271557600080fd5b6127218682870161232d565b9497909650939450505050565b600060808201868352602060018060a01b038716818501526080604085015281865180845260a08601915060a08160051b870101935082880160005b8281101561279857609f19888703018452612786868351612465565b9550928401929084019060010161276a565b505050505082810360608401526127af8185612465565b979650505050505050565b60006001600160401b038211156127d3576127d361256c565b5060051b60200190565b600082601f8301126127ee57600080fd5b813560206128036127fe836127ba565b6125cc565b82815260059290921b8401810191818101908684111561282257600080fd5b8286015b848110156128c05780356001600160401b03808211156128465760008081fd5b908801906040828b03601f19018113156128605760008081fd5b612868612582565b878401358381111561287a5760008081fd5b6128888d8a838801016125fc565b82525090830135908282111561289e5760008081fd5b6128ac8c89848701016125fc565b818901528652505050918301918301612826565b509695505050505050565b6000604082840312156128dd57600080fd5b6128e5612582565b905081356001600160401b03808211156128fe57600080fd5b818401915084601f83011261291257600080fd5b813560206129226127fe836127ba565b82815260059290921b8401810191818101908884111561294157600080fd5b8286015b84811015612a225780358681111561295d5760008081fd5b87016080818c03601f19018113156129755760008081fd5b61297d6125aa565b868301358981111561298f5760008081fd5b61299d8e89838701016127dd565b8252506040830135898111156129b35760008081fd5b6129c18e89838701016125fc565b88830152506060808401358a8111156129da5760008081fd5b6129e88f8a838801016125fc565b604084015250918301359189831115612a015760008081fd5b612a0f8e89858701016125fc565b9082015285525050918301918301612945565b50808752505080860135818601525050505092915050565b600082601f830112612a4b57600080fd5b81356020612a5b6127fe836127ba565b82815260059290921b84018101918181019086841115612a7a57600080fd5b8286015b848110156128c05780356001600160401b03811115612a9d5760008081fd5b612aab8986838b01016125fc565b845250918301918301612a7e565b80356001600160401b0381168114612ad057600080fd5b919050565b60008060008060008060c08789031215612aee57600080fd5b86356001600160401b0380821115612b0557600080fd5b612b118a838b016128cb565b97506020890135915080821115612b2757600080fd5b612b338a838b01612a3a565b965060408901359550612b4860608a01612ab9565b94506080890135915080821115612b5e57600080fd5b612b6a8a838b016125fc565b935060a0890135915080821115612b8057600080fd5b50612b8d89828a016125fc565b9150509295509295509295565b634e487b7160e01b600052601160045260246000fd5b600082821015612bc257612bc2612b9a565b500390565b600181811c90821680612bdb57607f821691505b60208210810361238057634e487b7160e01b600052602260045260246000fd5b84815260006020608081840152612c156080840187612465565b8381036040850152612c278187612465565b905083810360608501526040810185516040835281815180845260608501915060608160051b8601019350858301925060005b81811015612d4857605f19868603018352835180516080808852815190880181905260a0600582901b89018101928b01919089019060005b81811015612ce557609f198b86030183528351805160408752612cb86040880182612465565b90508e82015191508681038f880152612cd18183612465565b96505050928c0192918c0191600101612c92565b50505050888201518782038a890152612cfe8282612465565b91505060408201518782036040890152612d188282612465565b915050606082015191508681036060880152612d348183612465565b965050509286019291860191600101612c5a565b50505050948201519101525090949350505050565b6000808335601e19843603018112612d7457600080fd5b8301803591506001600160401b03821115612d8e57600080fd5b6020019150600581901b360382131561036457600080fd5b634e487b7160e01b600052603260045260246000fd5b60008235607e19833603018112612dd257600080fd5b9190910192915050565b6000808335601e19843603018112612df357600080fd5b8301803591506001600160401b03821115612e0d57600080fd5b60200191503681900382131561036457600080fd5b8183823760009101908152919050565b8035602083101561037b57600019602084900360031b1b1692915050565b805160208083015191908110156123805760001960209190910360031b1b16919050565b600060018201612e8657612e86612b9a565b5060010190565b65636861696e2f60d01b815260008451612eae816006850160208901612439565b672f636c69656e742f60c01b6006918401918201528451612ed681600e840160208901612439565b6c2f72656365697074526f6f742f60981b600e92909101918201528351612f0481601b840160208801612439565b01601b0195945050505050565b60008219821115612f2457612f24612b9a565b500190565b634e487b7160e01b600052601260045260246000fd5b600060ff831680612f5257612f52612f29565b8060ff84160691505092915050565b600060ff821660ff841680821015612f7b57612f7b612b9a565b90039392505050565b60008085851115612f9457600080fd5b83861115612fa157600080fd5b5050820193919092039150565b85815284602082015283604082015260806060820152816080820152818360a0830137600081830160a090810191909152601f909201601f19160101949350505050565b60006020828403121561300457600080fd5b5051919050565b868882376000878201600081528751613028818360208c01612439565b018587823760009086019081528451613045818360208901612439565b0192835250506020019695505050505050565b60008251612dd2818460208701612439565b60008235603e19833603018112612dd257600080fd5b8486823760008582018581528385602083013760009301602001928352509095945050505050565b6000826130b7576130b7612f29565b500490565b6000826130cb576130cb612f29565b500690565b6000816130df576130df612b9a565b506000190190565b600083516130f9818460208801612439565b6001600160f81b0319939093169190920190815260010192915050565b6020808252604a908201527f524c505265616465723a206c656e677468206f6620616e20524c50206974656d60408201527f206d7573742062652067726561746572207468616e207a65726f20746f206265606082015269206465636f6461626c6560b01b608082015260a0019056fe524c505265616465723a206c656e677468206f6620636f6e74656e74206d7573a2646970667358221220220652664a42da23e32336852bb7eb5958ddec4e860b1e9049e0ac06158696ac64736f6c634300080f0033";

type CrossL2ProverConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: CrossL2ProverConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class CrossL2Prover__factory extends ContractFactory {
  constructor(...args: CrossL2ProverConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    verifier_: AddressLike,
    clientType_: string,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(verifier_, clientType_, overrides || {});
  }
  override deploy(
    verifier_: AddressLike,
    clientType_: string,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(verifier_, clientType_, overrides || {}) as Promise<
      CrossL2Prover & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): CrossL2Prover__factory {
    return super.connect(runner) as CrossL2Prover__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): CrossL2ProverInterface {
    return new Interface(_abi) as CrossL2ProverInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): CrossL2Prover {
    return new Contract(address, _abi, runner) as unknown as CrossL2Prover;
  }
}
