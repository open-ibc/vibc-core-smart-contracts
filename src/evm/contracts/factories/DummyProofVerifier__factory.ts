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
import type {
  DummyProofVerifier,
  DummyProofVerifierInterface,
} from "../DummyProofVerifier";

const _abi = [
  {
    type: "function",
    name: "verifyMembership",
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
    type: "function",
    name: "verifyStateUpdate",
    inputs: [
      {
        name: "",
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
        name: "",
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
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "",
        type: "uint64",
        internalType: "uint64",
      },
    ],
    outputs: [],
    stateMutability: "pure",
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
  "0x608060405234801561001057600080fd5b506102cf806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630a1bb8b5146100465780632a6ded741461005d578063c2f0329f14610071575b600080fd5b61005b610054366004610087565b5050505050565b005b61005b61006b36600461018a565b50505050565b61005b61007f3660046101fd565b505050505050565b600080600080600060a0868803121561009f57600080fd5b853567ffffffffffffffff808211156100b757600080fd5b908701906060828a0312156100cb57600080fd5b909550602087013590808211156100e157600080fd5b908701906080828a0312156100f557600080fd5b9094506040870135935060608701359250608087013590808216821461011a57600080fd5b50809150509295509295909350565b60008083601f84011261013b57600080fd5b50813567ffffffffffffffff81111561015357600080fd5b60208301915083602082850101111561016b57600080fd5b9250929050565b60006040828403121561018457600080fd5b50919050565b600080600080606085870312156101a057600080fd5b84359350602085013567ffffffffffffffff808211156101bf57600080fd5b6101cb88838901610129565b909550935060408701359150808211156101e457600080fd5b506101f187828801610172565b91505092959194509250565b6000806000806000806080878903121561021657600080fd5b86359550602087013567ffffffffffffffff8082111561023557600080fd5b6102418a838b01610129565b9097509550604089013591508082111561025a57600080fd5b6102668a838b01610129565b9095509350606089013591508082111561027f57600080fd5b5061028c89828a01610172565b915050929550929550929556fea2646970667358221220e0a3712a43ec910d057739751bae188269aec2b0ea5b0318d85bbf17cb64ec1864736f6c634300080f0033";

type DummyProofVerifierConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: DummyProofVerifierConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class DummyProofVerifier__factory extends ContractFactory {
  constructor(...args: DummyProofVerifierConstructorParams) {
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
      DummyProofVerifier & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): DummyProofVerifier__factory {
    return super.connect(runner) as DummyProofVerifier__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): DummyProofVerifierInterface {
    return new Interface(_abi) as DummyProofVerifierInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): DummyProofVerifier {
    return new Contract(address, _abi, runner) as unknown as DummyProofVerifier;
  }
}
