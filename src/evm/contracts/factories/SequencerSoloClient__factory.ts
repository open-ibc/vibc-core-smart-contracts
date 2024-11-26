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
import type {
  SequencerSoloClient,
  SequencerSoloClientInterface,
} from "../SequencerSoloClient";

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
      {
        name: "",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [],
    stateMutability: "pure",
  },
  {
    type: "error",
    name: "AppHashHasNotPassedFraudProofWindow",
    inputs: [],
  },
  {
    type: "error",
    name: "CannotUpdateClientWithDifferentAppHash",
    inputs: [],
  },
  {
    type: "error",
    name: "InvalidL1Origin",
    inputs: [],
  },
  {
    type: "error",
    name: "NoConsensusStateAtHeight",
    inputs: [
      {
        name: "height",
        type: "uint256",
        internalType: "uint256",
      },
    ],
  },
  {
    type: "error",
    name: "NonMembershipProofsNotYetImplemented",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x60c06040526000805460ff1916600217905534801561001d57600080fd5b50604051610a88380380610a8883398101604081905261003c9161006b565b6001600160a01b039182166080521660a0526100a5565b6001600160a01b038116811461006857600080fd5b50565b6000806040838503121561007e57600080fd5b825161008981610053565b602084015190925061009a81610053565b809150509250929050565b60805160a0516109aa6100de6000396000818161015401526101b801526000818160c501528181610272015261038701526109aa6000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806357c1c5f41161005b57806357c1c5f414610127578063cb535ab51461013c578063eb7720581461014f578063fdaab4e51461017657600080fd5b80631b738a221461008d5780632b7ac3f3146100c057806344c9af28146100ff57806349ff245e14610112575b600080fd5b6100ad61009b366004610438565b60016020526000908152604090205481565b6040519081526020015b60405180910390f35b6100e77f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100b7565b6100ad61010d366004610438565b610189565b61012561012036600461049a565b61019f565b005b61012f600281565b6040516100b791906104eb565b61012561014a36600461052b565b610268565b6100e77f000000000000000000000000000000000000000000000000000000000000000081565b6101256101843660046105c0565b61031d565b6000818152600160205260408120545b92915050565b6101ad602060008587610629565b6101b691610653565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166309bd5a606040518163ffffffff1660e01b8152600401602060405180830381865afa158015610214573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102389190610671565b1461025657604051632b56254960e01b815260040160405180910390fd5b61026284848484610336565b50505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663c2f0329f6102bb6102a9600160208a013561068a565b60009081526001602052604090205490565b6040516001600160e01b031960e084901b1681526102e6919088908890889088908d90600401610767565b60006040518083038186803b1580156102fe57600080fd5b505afa158015610312573d6000803e3d6000fd5b505050505050505050565b604051634dfb272b60e11b815260040160405180910390fd5b6000828152600160205260409020541561037d57600082815260016020526040902054811461037857604051631549535560e01b815260040160405180910390fd5b610262565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001663e8d9003983836103bc60206000898b610629565b6103c591610653565b6103d2886020818c610629565b6040518663ffffffff1660e01b81526004016103f2959493929190610943565b600060405180830381600087803b15801561040c57600080fd5b505af1158015610420573d6000803e3d6000fd5b50505060009283525060016020526040909120555050565b60006020828403121561044a57600080fd5b5035919050565b60008083601f84011261046357600080fd5b50813567ffffffffffffffff81111561047b57600080fd5b60208301915083602082850101111561049357600080fd5b9250929050565b600080600080606085870312156104b057600080fd5b843567ffffffffffffffff8111156104c757600080fd5b6104d387828801610451565b90989097506020870135966040013595509350505050565b602081016004831061050d57634e487b7160e01b600052602160045260246000fd5b91905290565b60006040828403121561052557600080fd5b50919050565b60008060008060006060868803121561054357600080fd5b853567ffffffffffffffff8082111561055b57600080fd5b61056789838a01610513565b9650602088013591508082111561057d57600080fd5b61058989838a01610451565b909650945060408801359150808211156105a257600080fd5b506105af88828901610451565b969995985093965092949392505050565b6000806000604084860312156105d557600080fd5b833567ffffffffffffffff808211156105ed57600080fd5b6105f987838801610513565b9450602086013591508082111561060f57600080fd5b5061061c86828701610451565b9497909650939450505050565b6000808585111561063957600080fd5b8386111561064657600080fd5b5050820193919092039150565b8035602083101561019957600019602084900360031b1b1692915050565b60006020828403121561068357600080fd5b5051919050565b6000828210156106aa57634e487b7160e01b600052601160045260246000fd5b500390565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000808335601e198436030181126106ef57600080fd5b830160208101925035905067ffffffffffffffff81111561070f57600080fd5b8060051b360382131561049357600080fd5b6000808335601e1984360301811261073857600080fd5b830160208101925035905067ffffffffffffffff81111561075857600080fd5b80360382131561049357600080fd5b8681526080602082015260006107816080830187896106af565b82810360408401526107948186886106af565b905082810360608401526107a884856106d8565b604083526040830181815260608401905060608260051b8501018360005b8481101561092457868303605f1901845236869003607e19018235126107eb57600080fd5b858235016107f981826106d8565b608086526080860181815260a08701905060a08260051b8801018360005b8481101561089a57898303609f1901845236869003603e190182351261083c57600080fd5b8582350161084a8182610721565b6040865261085c6040870182846106af565b91505061086c6020830183610721565b925085820360208701526108818284836106af565b6020978801979096509490940193505050600101610817565b50506108a96020860186610721565b9450925087810360208901526108c08185856106af565b93505050506108d26040830183610721565b86830360408801526108e58382846106af565b925050506108f66060830183610721565b9250858203606087015261090b8284836106af565b60209788019790965094909401935050506001016107c6565b5050602088013560208601528095505050505050979650505050505050565b8581528460208201528360408201526080606082015260006109696080830184866106af565b97965050505050505056fea2646970667358221220a0a963ec5c4f6cf50ff5b13642bbf8da13c52fbbc0f400c16b397f761fc25cb764736f6c634300080f0033";

type SequencerSoloClientConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: SequencerSoloClientConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class SequencerSoloClient__factory extends ContractFactory {
  constructor(...args: SequencerSoloClientConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    verifier_: AddressLike,
    _l1BlockProvider: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(
      verifier_,
      _l1BlockProvider,
      overrides || {}
    );
  }
  override deploy(
    verifier_: AddressLike,
    _l1BlockProvider: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(
      verifier_,
      _l1BlockProvider,
      overrides || {}
    ) as Promise<
      SequencerSoloClient & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(
    runner: ContractRunner | null
  ): SequencerSoloClient__factory {
    return super.connect(runner) as SequencerSoloClient__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): SequencerSoloClientInterface {
    return new Interface(_abi) as SequencerSoloClientInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): SequencerSoloClient {
    return new Contract(
      address,
      _abi,
      runner
    ) as unknown as SequencerSoloClient;
  }
}
