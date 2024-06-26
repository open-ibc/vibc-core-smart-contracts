/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type { IERC1967, IERC1967Interface } from "../IERC1967";

const _abi = [
  {
    type: "event",
    name: "AdminChanged",
    inputs: [
      {
        name: "previousAdmin",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "newAdmin",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "BeaconUpgraded",
    inputs: [
      {
        name: "beacon",
        type: "address",
        indexed: true,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "Upgraded",
    inputs: [
      {
        name: "implementation",
        type: "address",
        indexed: true,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
] as const;

export class IERC1967__factory {
  static readonly abi = _abi;
  static createInterface(): IERC1967Interface {
    return new Interface(_abi) as IERC1967Interface;
  }
  static connect(address: string, runner?: ContractRunner | null): IERC1967 {
    return new Contract(address, _abi, runner) as unknown as IERC1967;
  }
}
