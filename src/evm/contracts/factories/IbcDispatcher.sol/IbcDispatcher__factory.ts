/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  IbcDispatcher,
  IbcDispatcherInterface,
} from "../../IbcDispatcher.sol/IbcDispatcher";

const _abi = [
  {
    type: "function",
    name: "channelCloseConfirm",
    inputs: [
      {
        name: "portAddress",
        type: "address",
        internalType: "address",
      },
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
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
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "channelCloseInit",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "channelOpenInit",
    inputs: [
      {
        name: "version",
        type: "string",
        internalType: "string",
      },
      {
        name: "ordering",
        type: "uint8",
        internalType: "enum ChannelOrder",
      },
      {
        name: "feeEnabled",
        type: "bool",
        internalType: "bool",
      },
      {
        name: "connectionHops",
        type: "string[]",
        internalType: "string[]",
      },
      {
        name: "counterpartyPortId",
        type: "string",
        internalType: "string",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "feeVault",
    inputs: [],
    outputs: [
      {
        name: "feeVault",
        type: "address",
        internalType: "contract IFeeVault",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "portPrefix",
    inputs: [],
    outputs: [
      {
        name: "portPrefix",
        type: "string",
        internalType: "string",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "sendPacket",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "payload",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "timeoutTimestamp",
        type: "uint64",
        internalType: "uint64",
      },
    ],
    outputs: [
      {
        name: "sequence",
        type: "uint64",
        internalType: "uint64",
      },
    ],
    stateMutability: "nonpayable",
  },
] as const;

export class IbcDispatcher__factory {
  static readonly abi = _abi;
  static createInterface(): IbcDispatcherInterface {
    return new Interface(_abi) as IbcDispatcherInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): IbcDispatcher {
    return new Contract(address, _abi, runner) as unknown as IbcDispatcher;
  }
}
