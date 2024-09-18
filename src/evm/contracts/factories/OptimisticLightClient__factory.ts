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
        name: "peptideHeight",
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
        name: "peptideHeight",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "peptideAppHash",
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
        name: "peptideHeight",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "peptideAppHash",
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
  "0x608060405234801561001057600080fd5b50604051610fbf380380610fbf83398101604081905261002f91610084565b63ffffffff92909216600255600380546001600160a01b039283166001600160a01b031991821617909155600480549290931691161790556100da565b6001600160a01b038116811461008157600080fd5b50565b60008060006060848603121561009957600080fd5b835163ffffffff811681146100ad57600080fd5b60208501519093506100be8161006c565b60408501519092506100cf8161006c565b809150509250925092565b610ed6806100e96000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806363042720116100715780636304272014610161578063b9a1e87b1461016a578063cb535ab51461019a578063d56ff842146101ad578063eb772058146101d8578063fdaab4e5146101eb57600080fd5b80631b738a22146100ae5780632b7ac3f3146100e157806334b80a411461010c57806344c9af281461012c57806349ff245e1461014c575b600080fd5b6100ce6100bc36600461057a565b60006020819052908152604090205481565b6040519081526020015b60405180910390f35b6003546100f4906001600160a01b031681565b6040516001600160a01b0390911681526020016100d8565b6100ce61011a36600461057a565b60016020526000908152604090205481565b6100ce61013a36600461057a565b60009081526020819052604090205490565b61015f61015a3660046105db565b6101fe565b005b6100ce60025481565b61017d61017836600461057a565b6103be565b6040805193845260208401929092521515908201526060016100d8565b61015f6101a8366004610643565b6103d9565b6100ce6101bb36600461057a565b600090815260208181526040808320548352600190915290205490565b6004546100f4906001600160a01b031681565b61015f6101f93660046106d7565b610489565b60008061020d85870187610956565b600086815260208190526040812054929450909250819003610395576003546004805460408051624dead360e51b815290516001600160a01b0394851694630a1bb8b594899489948c9493909116926309bd5a60928281019260209291908290030181865afa158015610284573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102a89190610a09565b60048054604080516341c0fac560e11b815290516001600160a01b0390921692638381f58a9282820192602092908290030181865afa1580156102ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103139190610a22565b6040518663ffffffff1660e01b8152600401610333959493929190610ad4565b60006040518083038186803b15801561034b57600080fd5b505afa15801561035f573d6000803e3d6000fd5b5050506000868152602081905260409020859055506002546103819042610ba1565b6000858152600160205260409020556103b5565b8381146103b55760405163f0cd4ed960e01b815260040160405180910390fd5b50505050505050565b60008060006103cc8461052e565b9250925092509193909250565b6000806103f36103ee600160208a0135610bb9565b61052e565b92505091508061041657604051631234d8dd60e01b815260040160405180910390fd5b60035460405163c2f0329f60e01b81526001600160a01b039091169063c2f0329f906104509085908a908a908a908a908f90600401610e28565b60006040518083038186803b15801561046857600080fd5b505afa15801561047c573d6000803e3d6000fd5b5050505050505050505050565b60008061049e6103ee60016020880135610bb9565b9250509150806104c157604051631234d8dd60e01b815260040160405180910390fd5b600354604051630a9b7b5d60e21b81526001600160a01b0390911690632a6ded74906104f7908590889088908b90600401610e69565b60006040518083038186803b15801561050f57600080fd5b505afa158015610523573d6000803e3d6000fd5b505050505050505050565b600081815260208181526040808320548084526001909252822054829182918190811580159061056c57506000838152600160205260409020544210155b935093509350509193909250565b60006020828403121561058c57600080fd5b5035919050565b60008083601f8401126105a557600080fd5b5081356001600160401b038111156105bc57600080fd5b6020830191508360208285010111156105d457600080fd5b9250929050565b600080600080606085870312156105f157600080fd5b84356001600160401b0381111561060757600080fd5b61061387828801610593565b90989097506020870135966040013595509350505050565b60006040828403121561063d57600080fd5b50919050565b60008060008060006060868803121561065b57600080fd5b85356001600160401b038082111561067257600080fd5b61067e89838a0161062b565b9650602088013591508082111561069457600080fd5b6106a089838a01610593565b909650945060408801359150808211156106b957600080fd5b506106c688828901610593565b969995985093965092949392505050565b6000806000604084860312156106ec57600080fd5b83356001600160401b038082111561070357600080fd5b61070f8783880161062b565b9450602086013591508082111561072557600080fd5b5061073286828701610593565b9497909650939450505050565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156107775761077761073f565b60405290565b604051601f8201601f191681016001600160401b03811182821017156107a5576107a561073f565b604052919050565b6000601f83818401126107bf57600080fd5b823560206001600160401b03808311156107db576107db61073f565b8260051b6107ea83820161077d565b938452868101830193838101908986111561080457600080fd5b84890192505b85831015610894578235848111156108225760008081fd5b8901603f81018b136108345760008081fd5b8581013560408682111561084a5761084a61073f565b61085b828b01601f1916890161077d565b8281528d828486010111156108705760008081fd5b828285018a830137600092810189019290925250835250918401919084019061080a565b9998505050505050505050565b6001600160401b03811681146108b657600080fd5b50565b6000608082840312156108cb57600080fd5b604051608081016001600160401b0382821081831117156108ee576108ee61073f565b81604052829350843591508082111561090657600080fd5b610912868387016107ad565b8352602085013591508082111561092857600080fd5b50610935858286016107ad565b60208301525060408301356040820152606083013560608201525092915050565b6000806040838503121561096957600080fd5b82356001600160401b038082111561098057600080fd5b908401906060828703121561099457600080fd5b61099c610755565b8235828111156109ab57600080fd5b6109b7888286016107ad565b82525060208301356020820152604083013592506109d4836108a1565b8260408201528094505060208501359150808211156109f257600080fd5b506109ff858286016108b9565b9150509250929050565b600060208284031215610a1b57600080fd5b5051919050565b600060208284031215610a3457600080fd5b8151610a3f816108a1565b9392505050565b600081518084526020808501808196508360051b810191508286016000805b86811015610ac6578385038a5282518051808752835b81811015610a96578281018901518882018a01528801610a7b565b81811115610aa6578489838a0101525b509a87019a601f01601f1916959095018601945091850191600101610a65565b509298975050505050505050565b60a0815260008651606060a0840152610af1610100840182610a46565b9050602088015160c08401526001600160401b0360408901511660e08401528281036020840152865160808252610b2b6080830182610a46565b905060208801518282036020840152610b448282610a46565b91505060408801516040830152606088015160608301528092505050846040830152836060830152610b8160808301846001600160401b03169052565b9695505050505050565b634e487b7160e01b600052601160045260246000fd5b60008219821115610bb457610bb4610b8b565b500190565b600082821015610bcb57610bcb610b8b565b500390565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e19843603018112610c1057600080fd5b83016020810192503590506001600160401b03811115610c2f57600080fd5b8060051b36038213156105d457600080fd5b6000808335601e19843603018112610c5857600080fd5b83016020810192503590506001600160401b03811115610c7757600080fd5b8036038213156105d457600080fd5b600060408301610c968384610bf9565b604086528281845260608701905060608260051b88010193508260005b83811015610e0f57888603605f1901835236859003607e1901823512610cd857600080fd5b8482350160808701610cea8283610bf9565b60808a528281845260a08b01905060a08260051b8c010193508260005b83811015610d88578c8603609f19018352813536869003603e19018112610d2d57600080fd5b8501610d398180610c41565b60408952610d4b60408a018284610bd0565b915050610d5b6020830183610c41565b925088820360208a0152610d70828483610bd0565b98505050602093840193929092019150600101610d07565b5050505050610d9a6020830183610c41565b89830360208b0152610dad838284610bd0565b92505050610dbe6040830183610c41565b89830360408b0152610dd1838284610bd0565b92505050610de26060830183610c41565b925088820360608a0152610df7828483610bd0565b98505050602093840193929092019150600101610cb3565b5050505050602083013560208501528091505092915050565b868152608060208201526000610e42608083018789610bd0565b8281036040840152610e55818688610bd0565b905082810360608401526108948185610c86565b848152606060208201526000610e83606083018587610bd0565b8281036040840152610e958185610c86565b97965050505050505056fea2646970667358221220853ec805731e7eed5df82abdd45c9074563932fa41032ccff8917caac32afd9f64736f6c634300080f0033";

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
