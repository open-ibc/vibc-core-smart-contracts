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
  "0x60a06040523480156200001157600080fd5b50604051620034f0380380620034f083398101604081905262000034916200006e565b6001600160a01b03821660805260006200004f8282620001fd565b505050620002c9565b634e487b7160e01b600052604160045260246000fd5b600080604083850312156200008257600080fd5b82516001600160a01b03811681146200009a57600080fd5b602084810151919350906001600160401b0380821115620000ba57600080fd5b818601915086601f830112620000cf57600080fd5b815181811115620000e457620000e462000058565b604051601f8201601f19908116603f011681019083821181831017156200010f576200010f62000058565b8160405282815289868487010111156200012857600080fd5b600093505b828410156200014c57848401860151818501870152928501926200012d565b828411156200015e5760008684830101525b8096505050505050509250929050565b600181811c908216806200018357607f821691505b602082108103620001a457634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115620001f857600081815260208120601f850160051c81016020861015620001d35750805b601f850160051c820191505b81811015620001f457828155600101620001df565b5050505b505050565b81516001600160401b0381111562000219576200021962000058565b62000231816200022a84546200016e565b84620001aa565b602080601f831160018114620002695760008415620002505750858301515b600019600386901b1c1916600185901b178555620001f4565b600085815260208120601f198616915b828110156200029a5788860151825594840194600190910190840162000279565b5085821015620002b95787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b608051613205620002eb6000396000818160bd0152610f3e01526132056000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c806357c1c5f41161006657806357c1c5f414610151578063b3768f0d14610166578063c2f0329f1461017b578063c67e15f71461018e578063fc396ddb146101ae57600080fd5b80632a6ded74146100a35780632b7ac3f3146100b85780632cd78e77146100fc57806344c9af281461011d57806349ff245e1461013e575b600080fd5b6100b66100b13660046123a7565b6101d1565b005b6100df7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b61010f61010a366004612419565b6101ea565b6040516100f39291906124b2565b61013061012b3660046124d3565b61036b565b6040519081526020016100f3565b6100b661014c3660046124ec565b610381565b610159600281565b6040516100f39190612552565b61016e610393565b6040516100f3919061257a565b6100b661018936600461268c565b610421565b61013061019c3660046124d3565b60016020526000908152604090205481565b6101c16101bc366004612704565b6105e3565b6040516100f3949392919061274f565b604051632974974360e01b815260040160405180910390fd5b60006060818080808080610200898b018b612af6565b955095509550955095509550306001600160a01b031663c2f0329f6102416001896020015161022f9190612bd1565b60009081526001602052604090205490565b60001b6102e2856000805461025590612be8565b80601f016020809104026020016040519081016040528092919081815260200182805461028190612be8565b80156102ce5780601f106102a3576101008083540402835291602001916102ce565b820191906000526020600020905b8154815290600101906020018083116102b157829003601f168201915b5050505050886001600160401b0316610614565b60408051602081018a9052016040516020818303038152906040528a6040518563ffffffff1660e01b815260040161031d9493929190612c1c565b60006040518083038186803b15801561033557600080fd5b505afa158015610349573d6000803e3d6000fd5b505050508761035982878761064b565b975097505050505050505b9250929050565b6000818152600160205260408120545b92915050565b61038d84848484610eed565b50505050565b600080546103a090612be8565b80601f01602080910402602001604051908101604052809291908181526020018280546103cc90612be8565b80156104195780601f106103ee57610100808354040283529160200191610419565b820191906000526020600020905b8154815290600101906020018083116103fc57829003601f168201915b505050505081565b61042b8180612d7e565b600081811061043c5761043c612dc7565b905060200281019061044e9190612ddd565b61045c906020810190612dfd565b60405161046a929190612e43565b60405180910390208380519060200120146104985760405163026a287560e51b815260040160405180910390fd5b6104a28180612d7e565b60008181106104b3576104b3612dc7565b90506020028101906104c59190612ddd565b6104d3906040810190612dfd565b6040516104e1929190612e43565b604051809103902082805190602001201461050f576040516310d9300f60e11b815260040160405180910390fd5b61054461051c8280612d7e565b600081811061052d5761052d612dc7565b905060200281019061053f9190612ddd565b610fef565b61054e8280612d7e565b600181811061055f5761055f612dc7565b90506020028101906105719190612ddd565b61057f906040810190612dfd565b61058891612e53565b146105a657604051636589f0e160e11b815260040160405180910390fd5b6105c46105b38280612d7e565b600181811061052d5761052d612dc7565b841461038d576040516392cb8fbb60e01b815260040160405180910390fd5b60008060608060606105f587876101ea565b90955090506106048882611213565b969a919950975094955050505050565b6060838361062184611329565b60405160200161063393929190612e71565b60405160208183030381529060405290509392505050565b6060600084511161069b5760405162461bcd60e51b81526020600482015260156024820152744d65726b6c65547269653a20656d707479206b657960581b60448201526064015b60405180910390fd5b60006106a68461143d565b905060006106b386611521565b90506000846040516020016106ca91815260200190565b60405160208183030381529060405290506000805b8451811015610e8f5760008582815181106106fc576106fc612dc7565b60200260200101519050845183111561076e5760405162461bcd60e51b815260206004820152602e60248201527f4d65726b6c65547269653a206b657920696e646578206578636565647320746f60448201526d0e8c2d840d6caf240d8cadccee8d60931b6064820152608401610692565b8260000361080d57805180516020918201206040516107bc9261079692910190815260200190565b604051602081830303815290604052858051602091820120825192909101919091201490565b6108085760405162461bcd60e51b815260206004820152601d60248201527f4d65726b6c65547269653a20696e76616c696420726f6f7420686173680000006044820152606401610692565b610903565b80515160201161089357805180516020918201206040516108379261079692910190815260200190565b6108085760405162461bcd60e51b815260206004820152602760248201527f4d65726b6c65547269653a20696e76616c6964206c6172676520696e7465726e6044820152660c2d840d0c2e6d60cb1b6064820152608401610692565b8051845160208087019190912082519190920120146109035760405162461bcd60e51b815260206004820152602660248201527f4d65726b6c65547269653a20696e76616c696420696e7465726e616c206e6f646044820152650ca40d0c2e6d60d31b6064820152608401610692565b61090f60106001612efd565b81602001515103610ab75784518303610a4f57610949816020015160108151811061093c5761093c612dc7565b6020026020010151611584565b965060008751116109c25760405162461bcd60e51b815260206004820152603b60248201527f4d65726b6c65547269653a2076616c7565206c656e677468206d75737420626560448201527f2067726561746572207468616e207a65726f20286272616e63682900000000006064820152608401610692565b600186516109d09190612bd1565b8214610a445760405162461bcd60e51b815260206004820152603a60248201527f4d65726b6c65547269653a2076616c7565206e6f6465206d757374206265206c60448201527f617374206e6f646520696e2070726f6f6620286272616e6368290000000000006064820152608401610692565b505050505050610ee6565b6000858481518110610a6357610a63612dc7565b602001015160f81c60f81b60f81c9050600082602001518260ff1681518110610a8e57610a8e612dc7565b60200260200101519050610aa1816116b0565b9550610aae600186612efd565b94505050610e7c565b600281602001515103610e23576000610acf826116d5565b9050600081600081518110610ae657610ae6612dc7565b016020015160f81c90506000610afd600283612f2b565b610b08906002612f4d565b90506000610b19848360ff166116f9565b90506000610b278a896116f9565b90506000610b35838361172a565b905080835114610bad5760405162461bcd60e51b815260206004820152603a60248201527f4d65726b6c65547269653a20706174682072656d61696e646572206d7573742060448201527f736861726520616c6c206e6962626c65732077697468206b65790000000000006064820152608401610692565b60ff851660021480610bc2575060ff85166003145b15610d635780825114610c3d5760405162461bcd60e51b815260206004820152603d60248201527f4d65726b6c65547269653a206b65792072656d61696e646572206d757374206260448201527f65206964656e746963616c20746f20706174682072656d61696e6465720000006064820152608401610692565b610c57876020015160018151811061093c5761093c612dc7565b9c5060008d5111610cd05760405162461bcd60e51b815260206004820152603960248201527f4d65726b6c65547269653a2076616c7565206c656e677468206d75737420626560448201527f2067726561746572207468616e207a65726f20286c65616629000000000000006064820152608401610692565b60018c51610cde9190612bd1565b8814610d525760405162461bcd60e51b815260206004820152603860248201527f4d65726b6c65547269653a2076616c7565206e6f6465206d757374206265206c60448201527f617374206e6f646520696e2070726f6f6620286c6561662900000000000000006064820152608401610692565b505050505050505050505050610ee6565b60ff85161580610d76575060ff85166001145b15610db557610da28760200151600181518110610d9557610d95612dc7565b60200260200101516116b0565b9950610dae818a612efd565b9850610e18565b60405162461bcd60e51b815260206004820152603260248201527f4d65726b6c65547269653a2072656365697665642061206e6f64652077697468604482015271040c2dc40eadcd6dcdeeedc40e0e4caccd2f60731b6064820152608401610692565b505050505050610e7c565b60405162461bcd60e51b815260206004820152602860248201527f4d65726b6c65547269653a20726563656976656420616e20756e706172736561604482015267626c65206e6f646560c01b6064820152608401610692565b5080610e8781612f70565b9150506106df565b5060405162461bcd60e51b815260206004820152602560248201527f4d65726b6c65547269653a2072616e206f7574206f662070726f6f6620656c656044820152646d656e747360d81b6064820152608401610692565b9392505050565b60008281526001602052604090205415610f34576000828152600160205260409020548114610f2f57604051631549535560e01b815260040160405180910390fd5b61038d565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663e8d900398383610f7360206000898b612f89565b610f7c91612e53565b610f89886020818c612f89565b6040518663ffffffff1660e01b8152600401610fa9959493929190612fb3565b600060405180830381600087803b158015610fc357600080fd5b505af1158015610fd7573d6000803e3d6000fd5b50505060009283525060016020526040909120555050565b60008060026110016040850185612dfd565b60405161100f929190612e43565b602060405180830381855afa15801561102c573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061104f9190612ff7565b905060026110606060850185612dfd565b6110776110706020880188612dfd565b90506117ae565b6110846020880188612dfd565b61108e60206117ae565b876040516020016110a59796959493929190613010565b60408051601f19818403018152908290526110bf9161305d565b602060405180830381855afa1580156110dc573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906110ff9190612ff7565b915060005b61110e8480612d7e565b905081101561120c5760026111238580612d7e565b8381811061113357611133612dc7565b9050602002810190611145919061306f565b61114f9080612dfd565b8561115a8880612d7e565b8681811061116a5761116a612dc7565b905060200281019061117c919061306f565b61118a906020810190612dfd565b60405160200161119e959493929190613085565b60408051601f19818403018152908290526111b89161305d565b602060405180830381855afa1580156111d5573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906111f89190612ff7565b92508061120481612f70565b915050611104565b5050919050565b6000606080600061123a6112358660018089516112309190612bd1565b61181d565b61195e565b905060006112756112648360038151811061125757611257612dc7565b602002602001015161196c565b888151811061125757611257612dc7565b905061128d8160008151811061093c5761093c612dc7565b611296906130ad565b60001c945060006112b38260018151811061125757611257612dc7565b90506112cb8260028151811061093c5761093c612dc7565b935060005b815181101561131e576112ee82828151811061093c5761093c612dc7565b86828151811061130057611300612dc7565b6020026020010181905250808061131690612f70565b9150506112d0565b505050509250925092565b6060816000036113505750506040805180820190915260018152600360fc1b602082015290565b6000825b801561137a578161136481612f70565b92506113739050600a826130d1565b9050611354565b6000826001600160401b038111156113945761139461258d565b6040519080825280601f01601f1916602001820160405280156113be576020820181803683370190505b509050825b8015611434576113d4600a876130e5565b6113df906030612efd565b60f81b826113ee600184612bd1565b815181106113fe576113fe612dc7565b60200101906001600160f81b031916908160001a905350611420600a876130d1565b95508061142c816130f9565b9150506113c3565b50949350505050565b8051606090806001600160401b0381111561145a5761145a61258d565b60405190808252806020026020018201604052801561149f57816020015b60408051808201909152606080825260208201528152602001906001900390816114785790505b50915060005b8181101561120c5760405180604001604052808583815181106114ca576114ca612dc7565b602002602001015181526020016114f98684815181106114ec576114ec612dc7565b602002602001015161195e565b81525083828151811061150e5761150e612dc7565b60209081029190910101526001016114a5565b606080604051905082518060011b603f8101601f1916830160405280835250602084016020830160005b83811015611579578060011b82018184015160001a8060041c8253600f81166001830153505060010161154b565b509295945050505050565b6060600080600061159485611b8f565b9194509250905060008160018111156115af576115af61253c565b146116225760405162461bcd60e51b815260206004820152603960248201527f524c505265616465723a206465636f646564206974656d207479706520666f7260448201527f206279746573206973206e6f7420612064617461206974656d000000000000006064820152608401610692565b61162c8284612efd565b8551146116985760405162461bcd60e51b815260206004820152603460248201527f524c505265616465723a2062797465732076616c756520636f6e7461696e732060448201527330b71034b73b30b634b2103932b6b0b4b73232b960611b6064820152608401610692565b6116a785602001518484612252565b95945050505050565b606060208260000151106116cc576116c782611584565b61037b565b61037b826122e5565b606061037b6116f4836020015160008151811061093c5761093c612dc7565b611521565b606082518210611718575060408051602081019091526000815261037b565b610ee683838486516112309190612bd1565b600080825184511061173d578251611740565b83515b90505b8082108015611797575082828151811061175f5761175f612dc7565b602001015160f81c60f81b6001600160f81b03191684838151811061178657611786612dc7565b01602001516001600160f81b031916145b156117a757816001019150611743565b5092915050565b6060805b608083106117f0578083607f1660801760f81b6040516020016117d6929190613110565b60408051601f198184030190525260079290921c916117b2565b808360f81b604051602001611806929190613110565b604051602081830303815290604052915050919050565b60608182601f0110156118635760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b6044820152606401610692565b8282840110156118a65760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b6044820152606401610692565b818301845110156118ed5760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b6044820152606401610692565b60608215801561190c5760405191506000825260208201604052611434565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561194557805183526020928301920161192d565b5050858452601f01601f19166040525050949350505050565b606061037b61196c836122fb565b6060600080600061197c85611b8f565b9194509250905060018160018111156119975761199761253c565b14611a0a5760405162461bcd60e51b815260206004820152603860248201527f524c505265616465723a206465636f646564206974656d207479706520666f7260448201527f206c697374206973206e6f742061206c697374206974656d00000000000000006064820152608401610692565b8451611a168385612efd565b14611a7e5760405162461bcd60e51b815260206004820152603260248201527f524c505265616465723a206c697374206974656d2068617320616e20696e76616044820152713634b2103230ba30903932b6b0b4b73232b960711b6064820152608401610692565b604080516020808252610420820190925290816020015b6040805180820190915260008082526020820152815260200190600190039081611a955790505093506000835b8651811015611b8357600080611b086040518060400160405280858c60000151611aec9190612bd1565b8152602001858c60200151611b019190612efd565b9052611b8f565b509150915060405180604001604052808383611b249190612efd565b8152602001848b60200151611b399190612efd565b815250888581518110611b4e57611b4e612dc7565b6020908102919091010152611b64600185612efd565b9350611b708183612efd565b611b7a9084612efd565b92505050611ac2565b50845250919392505050565b600080600080846000015111611bb75760405162461bcd60e51b81526004016106929061313f565b6020840151805160001a607f8111611bdc57600060016000945094509450505061224b565b60b78111611d39576000611bf1608083612bd1565b905080876000015111611c715760405162461bcd60e51b815260206004820152604e60248201526000805160206131b083398151915260448201527f742062652067726561746572207468616e20737472696e67206c656e6774682060648201526d2873686f727420737472696e672960901b608482015260a401610692565b6001838101516001600160f81b0319169082141580611c9e5750600160ff1b6001600160f81b0319821610155b611d265760405162461bcd60e51b815260206004820152604d60248201527f524c505265616465723a20696e76616c6964207072656669782c2073696e676c60448201527f652062797465203c203078383020617265206e6f74207072656669786564202860648201526c73686f727420737472696e672960981b608482015260a401610692565b506001955093506000925061224b915050565b60bf8111611f7a576000611d4e60b783612bd1565b905080876000015111611dd15760405162461bcd60e51b815260206004820152605160248201526000805160206131b083398151915260448201527f74206265203e207468616e206c656e677468206f6620737472696e67206c656e60648201527067746820286c6f6e6720737472696e672960781b608482015260a401610692565b60018301516001600160f81b0319166000819003611e585760405162461bcd60e51b815260206004820152604a60248201526000805160206131b083398151915260448201527f74206e6f74206861766520616e79206c656164696e67207a65726f7320286c6f6064820152696e6720737472696e672960b01b608482015260a401610692565b600184015160088302610100031c60378111611edb5760405162461bcd60e51b815260206004820152604860248201526000805160206131b083398151915260448201527f742062652067726561746572207468616e20353520627974657320286c6f6e6760648201526720737472696e672960c01b608482015260a401610692565b611ee58184612efd565b895111611f5d5760405162461bcd60e51b815260206004820152604c60248201526000805160206131b083398151915260448201527f742062652067726561746572207468616e20746f74616c206c656e677468202860648201526b6c6f6e6720737472696e672960a01b608482015260a401610692565b611f68836001612efd565b975095506000945061224b9350505050565b60f7811161201c576000611f8f60c083612bd1565b90508087600001511161200b5760405162461bcd60e51b815260206004820152604a60248201526000805160206131b083398151915260448201527f742062652067726561746572207468616e206c697374206c656e677468202873606482015269686f7274206c6973742960b01b608482015260a401610692565b60019550935084925061224b915050565b600061202960f783612bd1565b9050808760000151116120a85760405162461bcd60e51b815260206004820152604d60248201526000805160206131b083398151915260448201527f74206265203e207468616e206c656e677468206f66206c697374206c656e677460648201526c6820286c6f6e67206c6973742960981b608482015260a401610692565b60018301516001600160f81b031916600081900361212d5760405162461bcd60e51b815260206004820152604860248201526000805160206131b083398151915260448201527f74206e6f74206861766520616e79206c656164696e67207a65726f7320286c6f6064820152676e67206c6973742960c01b608482015260a401610692565b600184015160088302610100031c603781116121ae5760405162461bcd60e51b815260206004820152604660248201526000805160206131b083398151915260448201527f742062652067726561746572207468616e20353520627974657320286c6f6e67606482015265206c6973742960d01b608482015260a401610692565b6121b88184612efd565b89511161222e5760405162461bcd60e51b815260206004820152604a60248201526000805160206131b083398151915260448201527f742062652067726561746572207468616e20746f74616c206c656e67746820286064820152696c6f6e67206c6973742960b01b608482015260a401610692565b612239836001612efd565b975095506001945061224b9350505050565b9193909250565b6060816001600160401b0381111561226c5761226c61258d565b6040519080825280601f01601f191660200182016040528015612296576020820181803683370190505b5090508115610ee65760006122ab8486612efd565b90506020820160005b848110156122cc5782810151828201526020016122b4565b848111156122db576000858301525b5050509392505050565b606061037b826020015160008460000151612252565b604080518082019091526000808252602082015260008251116123305760405162461bcd60e51b81526004016106929061313f565b50604080518082019091528151815260209182019181019190915290565b60008083601f84011261236057600080fd5b5081356001600160401b0381111561237757600080fd5b60208301915083602082850101111561036457600080fd5b6000604082840312156123a157600080fd5b50919050565b600080600080606085870312156123bd57600080fd5b8435935060208501356001600160401b03808211156123db57600080fd5b6123e78883890161234e565b9095509350604087013591508082111561240057600080fd5b5061240d8782880161238f565b91505092959194509250565b6000806020838503121561242c57600080fd5b82356001600160401b0381111561244257600080fd5b61244e8582860161234e565b90969095509350505050565b60005b8381101561247557818101518382015260200161245d565b8381111561038d5750506000910152565b6000815180845261249e81602086016020860161245a565b601f01601f19169290920160200192915050565b8281526040602082015260006124cb6040830184612486565b949350505050565b6000602082840312156124e557600080fd5b5035919050565b6000806000806060858703121561250257600080fd5b84356001600160401b0381111561251857600080fd5b6125248782880161234e565b90989097506020870135966040013595509350505050565b634e487b7160e01b600052602160045260246000fd5b602081016004831061257457634e487b7160e01b600052602160045260246000fd5b91905290565b602081526000610ee66020830184612486565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b03811182821017156125c5576125c561258d565b60405290565b604051608081016001600160401b03811182821017156125c5576125c561258d565b604051601f8201601f191681016001600160401b03811182821017156126155761261561258d565b604052919050565b600082601f83011261262e57600080fd5b81356001600160401b038111156126475761264761258d565b61265a601f8201601f19166020016125ed565b81815284602083860101111561266f57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080608085870312156126a257600080fd5b8435935060208501356001600160401b03808211156126c057600080fd5b6126cc8883890161261d565b945060408701359150808211156126e257600080fd5b6126ee8883890161261d565b9350606087013591508082111561240057600080fd5b60008060006040848603121561271957600080fd5b8335925060208401356001600160401b0381111561273657600080fd5b6127428682870161234e565b9497909650939450505050565b600060808201868352602060018060a01b038716818501526080604085015281865180845260a08601915060a08160051b870101935082880160005b828110156127b957609f198887030184526127a7868351612486565b9550928401929084019060010161278b565b505050505082810360608401526127d08185612486565b979650505050505050565b60006001600160401b038211156127f4576127f461258d565b5060051b60200190565b600082601f83011261280f57600080fd5b8135602061282461281f836127db565b6125ed565b82815260059290921b8401810191818101908684111561284357600080fd5b8286015b848110156128e15780356001600160401b03808211156128675760008081fd5b908801906040828b03601f19018113156128815760008081fd5b6128896125a3565b878401358381111561289b5760008081fd5b6128a98d8a8388010161261d565b8252509083013590828211156128bf5760008081fd5b6128cd8c898487010161261d565b818901528652505050918301918301612847565b509695505050505050565b6000604082840312156128fe57600080fd5b6129066125a3565b905081356001600160401b038082111561291f57600080fd5b818401915084601f83011261293357600080fd5b8135602061294361281f836127db565b82815260059290921b8401810191818101908884111561296257600080fd5b8286015b84811015612a435780358681111561297e5760008081fd5b87016080818c03601f19018113156129965760008081fd5b61299e6125cb565b86830135898111156129b05760008081fd5b6129be8e89838701016127fe565b8252506040830135898111156129d45760008081fd5b6129e28e898387010161261d565b88830152506060808401358a8111156129fb5760008081fd5b612a098f8a8388010161261d565b604084015250918301359189831115612a225760008081fd5b612a308e898587010161261d565b9082015285525050918301918301612966565b50808752505080860135818601525050505092915050565b600082601f830112612a6c57600080fd5b81356020612a7c61281f836127db565b82815260059290921b84018101918181019086841115612a9b57600080fd5b8286015b848110156128e15780356001600160401b03811115612abe5760008081fd5b612acc8986838b010161261d565b845250918301918301612a9f565b80356001600160401b0381168114612af157600080fd5b919050565b60008060008060008060c08789031215612b0f57600080fd5b86356001600160401b0380821115612b2657600080fd5b612b328a838b016128ec565b97506020890135915080821115612b4857600080fd5b612b548a838b01612a5b565b965060408901359550612b6960608a01612ada565b94506080890135915080821115612b7f57600080fd5b612b8b8a838b0161261d565b935060a0890135915080821115612ba157600080fd5b50612bae89828a0161261d565b9150509295509295509295565b634e487b7160e01b600052601160045260246000fd5b600082821015612be357612be3612bbb565b500390565b600181811c90821680612bfc57607f821691505b6020821081036123a157634e487b7160e01b600052602260045260246000fd5b84815260006020608081840152612c366080840187612486565b8381036040850152612c488187612486565b905083810360608501526040810185516040835281815180845260608501915060608160051b8601019350858301925060005b81811015612d6957605f19868603018352835180516080808852815190880181905260a0600582901b89018101928b01919089019060005b81811015612d0657609f198b86030183528351805160408752612cd96040880182612486565b90508e82015191508681038f880152612cf28183612486565b96505050928c0192918c0191600101612cb3565b50505050888201518782038a890152612d1f8282612486565b91505060408201518782036040890152612d398282612486565b915050606082015191508681036060880152612d558183612486565b965050509286019291860191600101612c7b565b50505050948201519101525090949350505050565b6000808335601e19843603018112612d9557600080fd5b8301803591506001600160401b03821115612daf57600080fd5b6020019150600581901b360382131561036457600080fd5b634e487b7160e01b600052603260045260246000fd5b60008235607e19833603018112612df357600080fd5b9190910192915050565b6000808335601e19843603018112612e1457600080fd5b8301803591506001600160401b03821115612e2e57600080fd5b60200191503681900382131561036457600080fd5b8183823760009101908152919050565b8035602083101561037b57600019602084900360031b1b1692915050565b65636861696e2f60d01b815260008451612e9281600685016020890161245a565b6f2f73746f72656452656365697074732f60801b6006918401918201528451612ec281601684016020890161245a565b6c2f72656365697074526f6f742f60981b601692909101918201528351612ef081602384016020880161245a565b0160230195945050505050565b60008219821115612f1057612f10612bbb565b500190565b634e487b7160e01b600052601260045260246000fd5b600060ff831680612f3e57612f3e612f15565b8060ff84160691505092915050565b600060ff821660ff841680821015612f6757612f67612bbb565b90039392505050565b600060018201612f8257612f82612bbb565b5060010190565b60008085851115612f9957600080fd5b83861115612fa657600080fd5b5050820193919092039150565b85815284602082015283604082015260806060820152816080820152818360a0830137600081830160a090810191909152601f909201601f19160101949350505050565b60006020828403121561300957600080fd5b5051919050565b86888237600087820160008152875161302d818360208c0161245a565b01858782376000908601908152845161304a81836020890161245a565b0192835250506020019695505050505050565b60008251612df381846020870161245a565b60008235603e19833603018112612df357600080fd5b8486823760008582018581528385602083013760009301602001928352509095945050505050565b805160208083015191908110156123a15760001960209190910360031b1b16919050565b6000826130e0576130e0612f15565b500490565b6000826130f4576130f4612f15565b500690565b60008161310857613108612bbb565b506000190190565b6000835161312281846020880161245a565b6001600160f81b0319939093169190920190815260010192915050565b6020808252604a908201527f524c505265616465723a206c656e677468206f6620616e20524c50206974656d60408201527f206d7573742062652067726561746572207468616e207a65726f20746f206265606082015269206465636f6461626c6560b01b608082015260a0019056fe524c505265616465723a206c656e677468206f6620636f6e74656e74206d7573a26469706673582212208f44f259f2331e19673df95d1f4d8ccd3d09c8481fe0c2929b548a0bc591da8164736f6c634300080f0033";

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
