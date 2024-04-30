/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type { ProofVerifier, ProofVerifierInterface } from "../ProofVerifier";

const _abi = [
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
    ],
    outputs: [],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "verifyNonMembership",
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
    ],
    outputs: [],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "verifyStateUpdate",
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
        name: "appHash",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "trustedL1BlockHash",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "trustedL1BlockNumber",
        type: "uint64",
        internalType: "uint64",
      },
    ],
    outputs: [],
    stateMutability: "view",
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

export class ProofVerifier__factory {
  static readonly abi = _abi;
  static createInterface(): ProofVerifierInterface {
    return new Interface(_abi) as ProofVerifierInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): ProofVerifier {
    return new Contract(address, _abi, runner) as unknown as ProofVerifier;
  }
}
