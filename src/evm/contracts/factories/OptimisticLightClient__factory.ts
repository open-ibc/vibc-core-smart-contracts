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
  BigNumberish,
  AddressLike,
  ContractDeployTransaction,
  ContractRunner,
} from "ethers";
import type { NonPayableOverrides } from "../common";
import type {
  OptimisticLightClient,
  OptimisticLightClientInterface,
} from "../OptimisticLightClient";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "fraudProofWindowSeconds_",
        type: "uint32",
        internalType: "uint32",
      },
      {
        name: "verifier_",
        type: "address",
        internalType: "contract IProofVerifier",
      },
      {
        name: "_l1BlockProvider",
        type: "address",
        internalType: "contract L1Block",
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
    name: "consensusStates",
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
    name: "fraudProofEndtime",
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
    name: "fraudProofWindowSeconds",
    inputs: [],
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
    name: "getFraudProofEndtime",
    inputs: [
      {
        name: "polymerHeight",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "fraudProofEndTime",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "getState",
    inputs: [
      {
        name: "polymerHeight",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "polymerAppHash",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "getStateAndEndTime",
    inputs: [
      {
        name: "polymerHeight",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "polymerAppHash",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "fraudProofEndTime",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "ended",
        type: "bool",
        internalType: "bool",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "l1BlockProvider",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "contract L1Block",
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
        name: "polymerHeight",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "polymerAppHash",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "verifier",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "contract IProofVerifier",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "verifyMembership",
    inputs: [
      {
        name: "proof",
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
      {
        name: "key",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "expectedValue",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "verifyNonMembership",
    inputs: [
      {
        name: "proof",
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
      {
        name: "key",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [],
    stateMutability: "view",
  },
  {
    type: "error",
    name: "AppHashHasNotPassedFraudProofWindow",
    inputs: [],
  },
  {
    type: "error",
    name: "CannotUpdatePendingOptimisticConsensusState",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x60806040526000805460ff1916600117905534801561001d57600080fd5b5060405161103638038061103683398101604081905261003c91610091565b63ffffffff92909216600355600480546001600160a01b039283166001600160a01b031991821617909155600580549290931691161790556100e7565b6001600160a01b038116811461008e57600080fd5b50565b6000806000606084860312156100a657600080fd5b835163ffffffff811681146100ba57600080fd5b60208501519093506100cb81610079565b60408501519092506100dc81610079565b809150509250925092565b610f40806100f66000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806363042720116100715780636304272014610181578063b9a1e87b1461018a578063cb535ab5146101ba578063d56ff842146101cd578063eb772058146101fa578063fdaab4e51461020d57600080fd5b80631b738a22146100b95780632b7ac3f3146100ec57806334b80a411461011757806344c9af281461013757806349ff245e1461015757806357c1c5f41461016c575b600080fd5b6100d96100c73660046105bc565b60016020526000908152604090205481565b6040519081526020015b60405180910390f35b6004546100ff906001600160a01b031681565b6040516001600160a01b0390911681526020016100e3565b6100d96101253660046105bc565b60026020526000908152604090205481565b6100d96101453660046105bc565b60009081526001602052604090205490565b61016a61016536600461061d565b610220565b005b610174600181565b6040516100e3919061066d565b6100d960035481565b61019d6101983660046105bc565b6103eb565b6040805193845260208401929092521515908201526060016100e3565b61016a6101c83660046106ad565b610406565b6100d96101db3660046105bc565b6000908152600160209081526040808320548352600290915290205490565b6005546100ff906001600160a01b031681565b61016a61021b366004610741565b6104ca565b60008061022f858701876109c0565b6000868152600160205260408120549294509092508190036103c2576004805460055460408051624dead360e51b815290516001600160a01b0393841694630a1bb8b594899489948c9492909116926309bd5a609282820192602092908290030181865afa1580156102a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c99190610a73565b600560009054906101000a90046001600160a01b03166001600160a01b0316638381f58a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561031c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103409190610a8c565b6040518663ffffffff1660e01b8152600401610360959493929190610b3e565b60006040518083038186803b15801561037857600080fd5b505afa15801561038c573d6000803e3d6000fd5b5050506000868152600160205260409020859055506003546103ae9042610c0b565b6000858152600260205260409020556103e2565b8381146103e25760405163f0cd4ed960e01b815260040160405180910390fd5b50505050505050565b60008060006103f98461056e565b9250925092509193909250565b60008061042061041b600160208a0135610c23565b61056e565b92505091508061044357604051631234d8dd60e01b815260040160405180910390fd5b600460009054906101000a90046001600160a01b03166001600160a01b031663c2f0329f8360001b888888888d6040518763ffffffff1660e01b815260040161049196959493929190610e92565b60006040518083038186803b1580156104a957600080fd5b505afa1580156104bd573d6000803e3d6000fd5b5050505050505050505050565b6000806104df61041b60016020880135610c23565b92505091508061050257604051631234d8dd60e01b815260040160405180910390fd5b60048054604051630a9b7b5d60e21b81526001600160a01b0390911691632a6ded7491610537918691899189918c9101610ed3565b60006040518083038186803b15801561054f57600080fd5b505afa158015610563573d6000803e3d6000fd5b505050505050505050565b600081815260016020908152604080832054808452600290925282205482918291819081158015906105ae57506000838152600260205260409020544210155b935093509350509193909250565b6000602082840312156105ce57600080fd5b5035919050565b60008083601f8401126105e757600080fd5b5081356001600160401b038111156105fe57600080fd5b60208301915083602082850101111561061657600080fd5b9250929050565b6000806000806060858703121561063357600080fd5b84356001600160401b0381111561064957600080fd5b610655878288016105d5565b90989097506020870135966040013595509350505050565b602081016003831061068f57634e487b7160e01b600052602160045260246000fd5b91905290565b6000604082840312156106a757600080fd5b50919050565b6000806000806000606086880312156106c557600080fd5b85356001600160401b03808211156106dc57600080fd5b6106e889838a01610695565b965060208801359150808211156106fe57600080fd5b61070a89838a016105d5565b9096509450604088013591508082111561072357600080fd5b50610730888289016105d5565b969995985093965092949392505050565b60008060006040848603121561075657600080fd5b83356001600160401b038082111561076d57600080fd5b61077987838801610695565b9450602086013591508082111561078f57600080fd5b5061079c868287016105d5565b9497909650939450505050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156107e1576107e16107a9565b60405290565b604051601f8201601f191681016001600160401b038111828210171561080f5761080f6107a9565b604052919050565b6000601f838184011261082957600080fd5b823560206001600160401b0380831115610845576108456107a9565b8260051b6108548382016107e7565b938452868101830193838101908986111561086e57600080fd5b84890192505b858310156108fe5782358481111561088c5760008081fd5b8901603f81018b1361089e5760008081fd5b858101356040868211156108b4576108b46107a9565b6108c5828b01601f191689016107e7565b8281528d828486010111156108da5760008081fd5b828285018a8301376000928101890192909252508352509184019190840190610874565b9998505050505050505050565b6001600160401b038116811461092057600080fd5b50565b60006080828403121561093557600080fd5b604051608081016001600160401b038282108183111715610958576109586107a9565b81604052829350843591508082111561097057600080fd5b61097c86838701610817565b8352602085013591508082111561099257600080fd5b5061099f85828601610817565b60208301525060408301356040820152606083013560608201525092915050565b600080604083850312156109d357600080fd5b82356001600160401b03808211156109ea57600080fd5b90840190606082870312156109fe57600080fd5b610a066107bf565b823582811115610a1557600080fd5b610a2188828601610817565b8252506020830135602082015260408301359250610a3e8361090b565b826040820152809450506020850135915080821115610a5c57600080fd5b50610a6985828601610923565b9150509250929050565b600060208284031215610a8557600080fd5b5051919050565b600060208284031215610a9e57600080fd5b8151610aa98161090b565b9392505050565b600081518084526020808501808196508360051b810191508286016000805b86811015610b30578385038a5282518051808752835b81811015610b00578281018901518882018a01528801610ae5565b81811115610b10578489838a0101525b509a87019a601f01601f1916959095018601945091850191600101610acf565b509298975050505050505050565b60a0815260008651606060a0840152610b5b610100840182610ab0565b9050602088015160c08401526001600160401b0360408901511660e08401528281036020840152865160808252610b956080830182610ab0565b905060208801518282036020840152610bae8282610ab0565b91505060408801516040830152606088015160608301528092505050846040830152836060830152610beb60808301846001600160401b03169052565b9695505050505050565b634e487b7160e01b600052601160045260246000fd5b60008219821115610c1e57610c1e610bf5565b500190565b600082821015610c3557610c35610bf5565b500390565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e19843603018112610c7a57600080fd5b83016020810192503590506001600160401b03811115610c9957600080fd5b8060051b360382131561061657600080fd5b6000808335601e19843603018112610cc257600080fd5b83016020810192503590506001600160401b03811115610ce157600080fd5b80360382131561061657600080fd5b600060408301610d008384610c63565b604086528281845260608701905060608260051b88010193508260005b83811015610e7957888603605f1901835236859003607e1901823512610d4257600080fd5b8482350160808701610d548283610c63565b60808a528281845260a08b01905060a08260051b8c010193508260005b83811015610df2578c8603609f19018352813536869003603e19018112610d9757600080fd5b8501610da38180610cab565b60408952610db560408a018284610c3a565b915050610dc56020830183610cab565b925088820360208a0152610dda828483610c3a565b98505050602093840193929092019150600101610d71565b5050505050610e046020830183610cab565b89830360208b0152610e17838284610c3a565b92505050610e286040830183610cab565b89830360408b0152610e3b838284610c3a565b92505050610e4c6060830183610cab565b925088820360608a0152610e61828483610c3a565b98505050602093840193929092019150600101610d1d565b5050505050602083013560208501528091505092915050565b868152608060208201526000610eac608083018789610c3a565b8281036040840152610ebf818688610c3a565b905082810360608401526108fe8185610cf0565b848152606060208201526000610eed606083018587610c3a565b8281036040840152610eff8185610cf0565b97965050505050505056fea264697066735822122002704f39522eba69937029ffb29e1a22a80c40ab8029a5a547a981d55cb5e56b64736f6c634300080f0033";

type OptimisticLightClientConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: OptimisticLightClientConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class OptimisticLightClient__factory extends ContractFactory {
  constructor(...args: OptimisticLightClientConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    fraudProofWindowSeconds_: BigNumberish,
    verifier_: AddressLike,
    _l1BlockProvider: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(
      fraudProofWindowSeconds_,
      verifier_,
      _l1BlockProvider,
      overrides || {}
    );
  }
  override deploy(
    fraudProofWindowSeconds_: BigNumberish,
    verifier_: AddressLike,
    _l1BlockProvider: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(
      fraudProofWindowSeconds_,
      verifier_,
      _l1BlockProvider,
      overrides || {}
    ) as Promise<
      OptimisticLightClient & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(
    runner: ContractRunner | null
  ): OptimisticLightClient__factory {
    return super.connect(runner) as OptimisticLightClient__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): OptimisticLightClientInterface {
    return new Interface(_abi) as OptimisticLightClientInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): OptimisticLightClient {
    return new Contract(
      address,
      _abi,
      runner
    ) as unknown as OptimisticLightClient;
  }
}
