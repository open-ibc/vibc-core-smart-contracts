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
    name: "addOpConsensusState",
    inputs: [
      {
        name: "l1header",
        type: "tuple",
        internalType: "struct L1Header",
        components: [
          {
            name: "header",
            type: "bytes[]",
            internalType: "bytes[]",
          },
          {
            name: "stateRoot",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "number",
            type: "uint64",
            internalType: "uint64",
          },
        ],
      },
      {
        name: "proof",
        type: "tuple",
        internalType: "struct OpL2StateProof",
        components: [
          {
            name: "accountProof",
            type: "bytes[]",
            internalType: "bytes[]",
          },
          {
            name: "outputRootProof",
            type: "bytes[]",
            internalType: "bytes[]",
          },
          {
            name: "l2OutputProposalKey",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "l2BlockHash",
            type: "bytes32",
            internalType: "bytes32",
          },
        ],
      },
      {
        name: "height",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "appHash",
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
      {
        name: "ended",
        type: "bool",
        internalType: "bool",
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
        name: "height",
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
    name: "getInternalState",
    inputs: [
      {
        name: "height",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "appHash",
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
        name: "appHash",
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
  "0x608060405234801561001057600080fd5b50604051610d5f380380610d5f83398101604081905261002f91610084565b63ffffffff92909216600255600380546001600160a01b039283166001600160a01b031991821617909155600480549290931691161790556100da565b6001600160a01b038116811461008157600080fd5b50565b60008060006060848603121561009957600080fd5b835163ffffffff811681146100ad57600080fd5b60208501519093506100be8161006c565b60408501519092506100cf8161006c565b809150509250925092565b610c76806100e96000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80635922f420116100715780635922f4201461016f5780636304272014610197578063cb535ab5146101a0578063d56ff842146101b5578063eb772058146101e0578063fdaab4e5146101f357600080fd5b80631b738a22146100ae5780631bc97a78146100e15780632b7ac3f31461011157806334b80a411461013c57806344c9af281461015c575b600080fd5b6100ce6100bc366004610591565b60006020819052908152604090205481565b6040519081526020015b60405180910390f35b6100f46100ef366004610591565b610206565b6040805193845260208401929092521515908201526060016100d8565b600354610124906001600160a01b031681565b6040516001600160a01b0390911681526020016100d8565b6100ce61014a366004610591565b60016020526000908152604090205481565b6100f461016a366004610591565b610252565b61018261017d3660046105aa565b61026d565b604080519283529015156020830152016100d8565b6100ce60025481565b6101b36101ae36600461068b565b610441565b005b6100ce6101c3366004610591565b600090815260208181526040808320548352600190915290205490565b600454610124906001600160a01b031681565b6101b3610201366004610720565b6104ec565b600081815260208181526040808320548084526001909252822054829182918190811580159061024457506000838152600160205260409020544210155b935093509350509193909250565b600080600061026084610206565b9250925092509193909250565b60008281526020819052604081205481908082036103fd576003546004805460408051624dead360e51b815290516001600160a01b0394851694630a1bb8b5948d948d948c9493909116926309bd5a60928281019260209291908290030181865afa1580156102e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103049190610789565b60048054604080516341c0fac560e11b815290516001600160a01b0390921692638381f58a9282820192602092908290030181865afa15801561034b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036f91906107bb565b6040518663ffffffff1660e01b815260040161038f9594939291906108f4565b60006040518083038186803b1580156103a757600080fd5b505afa1580156103bb573d6000803e3d6000fd5b50505060008681526020819052604081208690556002549091506103df90426109ea565b60008681526001602052604081208290559094509250610438915050565b83810361041f5760009081526001602052604090205491505042811115610438565b60405163f0cd4ed960e01b815260040160405180910390fd5b94509492505050565b6000806104566100ef600160208a0135610a02565b92505091508061047957604051631234d8dd60e01b815260040160405180910390fd5b60035460405163c2f0329f60e01b81526001600160a01b039091169063c2f0329f906104b39085908a908a908a908a908f90600401610bbb565b60006040518083038186803b1580156104cb57600080fd5b505afa1580156104df573d6000803e3d6000fd5b5050505050505050505050565b6000806105016100ef60016020880135610a02565b92505091508061052457604051631234d8dd60e01b815260040160405180910390fd5b600354604051630a9b7b5d60e21b81526001600160a01b0390911690632a6ded749061055a908590889088908b90600401610c09565b60006040518083038186803b15801561057257600080fd5b505afa158015610586573d6000803e3d6000fd5b505050505050505050565b6000602082840312156105a357600080fd5b5035919050565b600080600080608085870312156105c057600080fd5b843567ffffffffffffffff808211156105d857600080fd5b90860190606082890312156105ec57600080fd5b9094506020860135908082111561060257600080fd5b5085016080818803121561061557600080fd5b93969395505050506040820135916060013590565b60006040828403121561063c57600080fd5b50919050565b60008083601f84011261065457600080fd5b50813567ffffffffffffffff81111561066c57600080fd5b60208301915083602082850101111561068457600080fd5b9250929050565b6000806000806000606086880312156106a357600080fd5b853567ffffffffffffffff808211156106bb57600080fd5b6106c789838a0161062a565b965060208801359150808211156106dd57600080fd5b6106e989838a01610642565b9096509450604088013591508082111561070257600080fd5b5061070f88828901610642565b969995985093965092949392505050565b60008060006040848603121561073557600080fd5b833567ffffffffffffffff8082111561074d57600080fd5b6107598783880161062a565b9450602086013591508082111561076f57600080fd5b5061077c86828701610642565b9497909650939450505050565b60006020828403121561079b57600080fd5b5051919050565b67ffffffffffffffff811681146107b857600080fd5b50565b6000602082840312156107cd57600080fd5b81516107d8816107a2565b9392505050565b6000808335601e198436030181126107f657600080fd5b830160208101925035905067ffffffffffffffff81111561081657600080fd5b8060051b360382131561068457600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e1984360301811261086857600080fd5b830160208101925035905067ffffffffffffffff81111561088857600080fd5b80360382131561068457600080fd5b81835260006020808501808196508560051b810191508460005b878110156108e75782840389526108c88288610851565b6108d3868284610828565b9a87019a95505050908401906001016108b1565b5091979650505050505050565b60a08152600061090487886107df565b606060a085015261091a61010085018284610897565b915050602088013560c08401526040880135610935816107a2565b67ffffffffffffffff1660e0840152828103602084015261095687806107df565b60808352610968608084018284610897565b91505061097860208901896107df565b838303602085015261098b838284610897565b92505050604088013560408301526060880135606083015280925050508460408301528360608301526109ca608083018467ffffffffffffffff169052565b9695505050505050565b634e487b7160e01b600052601160045260246000fd5b600082198211156109fd576109fd6109d4565b500190565b600082821015610a1457610a146109d4565b500390565b600060408301610a2983846107df565b604086528281845260608701905060608260051b88010193508260005b83811015610ba257888603605f1901835236859003607e1901823512610a6b57600080fd5b8482350160808701610a7d82836107df565b60808a528281845260a08b01905060a08260051b8c010193508260005b83811015610b1b578c8603609f19018352813536869003603e19018112610ac057600080fd5b8501610acc8180610851565b60408952610ade60408a018284610828565b915050610aee6020830183610851565b925088820360208a0152610b03828483610828565b98505050602093840193929092019150600101610a9a565b5050505050610b2d6020830183610851565b89830360208b0152610b40838284610828565b92505050610b516040830183610851565b89830360408b0152610b64838284610828565b92505050610b756060830183610851565b925088820360608a0152610b8a828483610828565b98505050602093840193929092019150600101610a46565b5050505050602083013560208501528091505092915050565b868152608060208201526000610bd5608083018789610828565b8281036040840152610be8818688610828565b90508281036060840152610bfc8185610a19565b9998505050505050505050565b848152606060208201526000610c23606083018587610828565b8281036040840152610c358185610a19565b97965050505050505056fea2646970667358221220e4a5f132f7cf78072488590227d73cd3e57b68ef63275bfbc05c60a6b7a458a364736f6c634300080f0033";

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
