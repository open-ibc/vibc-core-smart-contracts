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
import type { IbcUtils, IbcUtilsInterface } from "../IbcUtils";

const _abi = [
  {
    type: "function",
    name: "fromUniversalPacketBytes",
    inputs: [
      {
        name: "data",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [
      {
        name: "universalPacketData",
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
    ],
    stateMutability: "pure",
  },
  {
    type: "function",
    name: "hexStrToAddress",
    inputs: [
      {
        name: "hexStr",
        type: "string",
        internalType: "string",
      },
    ],
    outputs: [
      {
        name: "addr",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "pure",
  },
  {
    type: "error",
    name: "StringTooLong",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidCharacter",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidHexStringLength",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x61067e61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c8063a1ef9a9814610045578063d5c39a9d14610075575b600080fd5b6100586100533660046103a1565b610095565b6040516001600160a01b0390911681526020015b60405180910390f35b610088610083366004610452565b610260565b60405161006c91906104c4565b805160009082906028146100bc576040516305229b2360e41b815260040160405180910390fd5b6000600160285b801561025657806100d381610555565b915050600060308583815181106100ec576100ec61056c565b016020015160f81c1080159061011c575060398583815181106101115761011161056c565b016020015160f81c11155b156101515760308583815181106101355761013561056c565b0160200151610147919060f81c610582565b60ff169050610224565b60418583815181106101655761016561056c565b016020015160f81c108015906101955750604685838151811061018a5761018a61056c565b016020015160f81c11155b156101ae5760378583815181106101355761013561056c565b60618583815181106101c2576101c261056c565b016020015160f81c108015906101f2575060668583815181106101e7576101e761056c565b016020015160f81c11155b1561020b5760578583815181106101355761013561056c565b60405163f379095160e01b815260040160405180910390fd5b61022d8361031a565b61023790826105a5565b61024190856105d4565b935061024e6010846105ff565b9250506100c3565b5090949350505050565b604080516080810182526000808252602082018190529181019190915260608082015260008060006020866000376000519250602080870160003760005191506020604087016000375060005160408051608081018252848152602081018490529081018290526060808201906102da908890818b61061e565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509152509695505050505050565b60006001600160a01b038211156103875760405162461bcd60e51b815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e20316044820152663630206269747360c81b606482015260840160405180910390fd5b5090565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156103b357600080fd5b813567ffffffffffffffff808211156103cb57600080fd5b818401915084601f8301126103df57600080fd5b8135818111156103f1576103f161038b565b604051601f8201601f19908116603f011681019083821181831017156104195761041961038b565b8160405282815287602084870101111561043257600080fd5b826020860160208301376000928101602001929092525095945050505050565b6000806020838503121561046557600080fd5b823567ffffffffffffffff8082111561047d57600080fd5b818501915085601f83011261049157600080fd5b8135818111156104a057600080fd5b8660208285010111156104b257600080fd5b60209290920196919550909350505050565b6000602080835283518184015280840151604084015260408401516060840152606084015160808085015280518060a086015260005b818110156105165782810184015186820160c0015283016104fa565b8181111561052857600060c083880101525b50601f01601f19169390930160c001949350505050565b634e487b7160e01b600052601160045260246000fd5b6000816105645761056461053f565b506000190190565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff84168082101561059c5761059c61053f565b90039392505050565b60006001600160a01b03828116848216811515828404821116156105cb576105cb61053f565b02949350505050565b60006001600160a01b038281168482168083038211156105f6576105f661053f565b01949350505050565b60008160001904831182151516156106195761061961053f565b500290565b6000808585111561062e57600080fd5b8386111561063b57600080fd5b505082019391909203915056fea2646970667358221220ca70ad57cf2f0719e16a40e446bf6054d89790c3945e65c0542a60ed6f75f9c564736f6c634300080f0033";

type IbcUtilsConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: IbcUtilsConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class IbcUtils__factory extends ContractFactory {
  constructor(...args: IbcUtilsConstructorParams) {
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
      IbcUtils & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): IbcUtils__factory {
    return super.connect(runner) as IbcUtils__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): IbcUtilsInterface {
    return new Interface(_abi) as IbcUtilsInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): IbcUtils {
    return new Contract(address, _abi, runner) as unknown as IbcUtils;
  }
}
