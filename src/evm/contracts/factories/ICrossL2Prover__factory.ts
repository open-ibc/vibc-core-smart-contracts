/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  ICrossL2Prover,
  ICrossL2ProverInterface,
} from "../ICrossL2Prover";

const _abi = [
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
    name: "updateClient",
    inputs: [
      {
        name: "proof",
        type: "bytes",
        internalType: "bytes",
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
        name: "srcChainId",
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
] as const;

export class ICrossL2Prover__factory {
  static readonly abi = _abi;
  static createInterface(): ICrossL2ProverInterface {
    return new Interface(_abi) as ICrossL2ProverInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): ICrossL2Prover {
    return new Contract(address, _abi, runner) as unknown as ICrossL2Prover;
  }
}
