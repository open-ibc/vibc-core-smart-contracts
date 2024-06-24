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
  "0x6105c261003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c8063a1ef9a9814610045578063d5c39a9d14610075575b600080fd5b610058610053366004610327565b610095565b6040516001600160a01b0390911681526020015b60405180910390f35b6100886100833660046103d8565b610257565b60405161006c919061044a565b805160009082906028146100bc576040516305229b2360e41b815260040160405180910390fd5b6000600160285b801561024d57806100d3816104db565b915050600060308583815181106100ec576100ec6104f2565b016020015160f81c1080159061011c57506039858381518110610111576101116104f2565b016020015160f81c11155b15610151576030858381518110610135576101356104f2565b0160200151610147919060f81c610508565b60ff169050610224565b6041858381518110610165576101656104f2565b016020015160f81c108015906101955750604685838151811061018a5761018a6104f2565b016020015160f81c11155b156101ae576037858381518110610135576101356104f2565b60618583815181106101c2576101c26104f2565b016020015160f81c108015906101f2575060668583815181106101e7576101e76104f2565b016020015160f81c11155b1561020b576057858381518110610135576101356104f2565b60405163f379095160e01b815260040160405180910390fd5b61022e838261052b565b610238908561054a565b935061024560108461052b565b9250506100c3565b5090949350505050565b604080516080810182526000808252602082018190529181019190915260608082015260008060006020866000376000519250602080870160003760005191506020604087016000375060005160408051608081018252848152602081018490529081018290526060808201906102d1908890818b610562565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509152509695505050505050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561033957600080fd5b813567ffffffffffffffff8082111561035157600080fd5b818401915084601f83011261036557600080fd5b81358181111561037757610377610311565b604051601f8201601f19908116603f0116810190838211818310171561039f5761039f610311565b816040528281528760208487010111156103b857600080fd5b826020860160208301376000928101602001929092525095945050505050565b600080602083850312156103eb57600080fd5b823567ffffffffffffffff8082111561040357600080fd5b818501915085601f83011261041757600080fd5b81358181111561042657600080fd5b86602082850101111561043857600080fd5b60209290920196919550909350505050565b6000602080835283518184015280840151604084015260408401516060840152606084015160808085015280518060a086015260005b8181101561049c5782810184015186820160c001528301610480565b818111156104ae57600060c083880101525b50601f01601f19169390930160c001949350505050565b634e487b7160e01b600052601160045260246000fd5b6000816104ea576104ea6104c5565b506000190190565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff841680821015610522576105226104c5565b90039392505050565b6000816000190483118215151615610545576105456104c5565b500290565b6000821982111561055d5761055d6104c5565b500190565b6000808585111561057257600080fd5b8386111561057f57600080fd5b505082019391909203915056fea2646970667358221220920cf9b2dba6cf22ef2d78571c1e71926572dd4bfba9bffdc98f5961817354c164736f6c634300080f0033";

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