/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  IbcMwPacketReceiver,
  IbcMwPacketReceiverInterface,
} from "../../IbcMiddleware.sol/IbcMwPacketReceiver";

const _abi = [
  {
    type: "function",
    name: "onRecvMWAck",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "ucPacket",
        type: "tuple",
        internalType: "struct UniversalPacket",
        components: [
          {
            name: "srcPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "mwBitmap",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "destPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "appData",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
      {
        name: "mwIndex",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "mwAddrs",
        type: "address[]",
        internalType: "address[]",
      },
      {
        name: "ack",
        type: "tuple",
        internalType: "struct AckPacket",
        components: [
          {
            name: "success",
            type: "bool",
            internalType: "bool",
          },
          {
            name: "data",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "onRecvMWPacket",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "ucPacket",
        type: "tuple",
        internalType: "struct UniversalPacket",
        components: [
          {
            name: "srcPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "mwBitmap",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "destPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "appData",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
      {
        name: "mwIndex",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "mwAddrs",
        type: "address[]",
        internalType: "address[]",
      },
    ],
    outputs: [
      {
        name: "ackPacket",
        type: "tuple",
        internalType: "struct AckPacket",
        components: [
          {
            name: "success",
            type: "bool",
            internalType: "bool",
          },
          {
            name: "data",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "onRecvMWTimeout",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
      {
        name: "ucPacket",
        type: "tuple",
        internalType: "struct UniversalPacket",
        components: [
          {
            name: "srcPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "mwBitmap",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "destPortAddr",
            type: "bytes32",
            internalType: "bytes32",
          },
          {
            name: "appData",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
      {
        name: "mwIndex",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "mwAddrs",
        type: "address[]",
        internalType: "address[]",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
] as const;

export class IbcMwPacketReceiver__factory {
  static readonly abi = _abi;
  static createInterface(): IbcMwPacketReceiverInterface {
    return new Interface(_abi) as IbcMwPacketReceiverInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): IbcMwPacketReceiver {
    return new Contract(
      address,
      _abi,
      runner
    ) as unknown as IbcMwPacketReceiver;
  }
}
