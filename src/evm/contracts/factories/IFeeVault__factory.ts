/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type { IFeeVault, IFeeVaultInterface } from "../IFeeVault";

const _abi = [
  {
    type: "function",
    name: "depositOpenChannelFee",
    inputs: [
      {
        name: "sender",
        type: "address",
        internalType: "address",
      },
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
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "depositSendPacketFee",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "sequence",
        type: "uint64",
        internalType: "uint64",
      },
      {
        name: "gasLimits",
        type: "uint256[2]",
        internalType: "uint256[2]",
      },
      {
        name: "gasPrices",
        type: "uint256[2]",
        internalType: "uint256[2]",
      },
    ],
    outputs: [],
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "withdrawFeesToOwner",
    inputs: [],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    name: "OpenChannelFeeDeposited",
    inputs: [
      {
        name: "sourceAddress",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "version",
        type: "string",
        indexed: false,
        internalType: "string",
      },
      {
        name: "ordering",
        type: "uint8",
        indexed: false,
        internalType: "enum ChannelOrder",
      },
      {
        name: "connectionHops",
        type: "string[]",
        indexed: false,
        internalType: "string[]",
      },
      {
        name: "counterpartyPortId",
        type: "string",
        indexed: false,
        internalType: "string",
      },
      {
        name: "feeAmount",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "SendPacketFeeDeposited",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        indexed: true,
        internalType: "bytes32",
      },
      {
        name: "sequence",
        type: "uint64",
        indexed: true,
        internalType: "uint64",
      },
      {
        name: "gasLimits",
        type: "uint256[2]",
        indexed: false,
        internalType: "uint256[2]",
      },
      {
        name: "gasPrices",
        type: "uint256[2]",
        indexed: false,
        internalType: "uint256[2]",
      },
    ],
    anonymous: false,
  },
  {
    type: "error",
    name: "FeeThresholdNotMet",
    inputs: [],
  },
  {
    type: "error",
    name: "IncorrectFeeSent",
    inputs: [
      {
        name: "expected",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "sent",
        type: "uint256",
        internalType: "uint256",
      },
    ],
  },
  {
    type: "error",
    name: "SenderNotDispatcher",
    inputs: [],
  },
] as const;

export class IFeeVault__factory {
  static readonly abi = _abi;
  static createInterface(): IFeeVaultInterface {
    return new Interface(_abi) as IFeeVaultInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): IFeeVault {
    return new Contract(address, _abi, runner) as unknown as IFeeVault;
  }
}
