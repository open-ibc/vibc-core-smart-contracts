/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  INonMembershipVerifier,
  INonMembershipVerifierInterface,
} from "../../ILightClient.sol/INonMembershipVerifier";

const _abi = [
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
    stateMutability: "nonpayable",
  },
] as const;

export class INonMembershipVerifier__factory {
  static readonly abi = _abi;
  static createInterface(): INonMembershipVerifierInterface {
    return new Interface(_abi) as INonMembershipVerifierInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): INonMembershipVerifier {
    return new Contract(
      address,
      _abi,
      runner
    ) as unknown as INonMembershipVerifier;
  }
}
