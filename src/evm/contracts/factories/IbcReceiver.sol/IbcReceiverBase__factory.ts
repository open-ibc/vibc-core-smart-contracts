/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Contract,
  ContractFactory,
  ContractTransactionResponse,
  Interface,
} from "ethers";
import type {
  Signer,
  AddressLike,
  ContractDeployTransaction,
  ContractRunner,
} from "ethers";
import type { NonPayableOverrides } from "../../common";
import type {
  IbcReceiverBase,
  IbcReceiverBaseInterface,
} from "../../IbcReceiver.sol/IbcReceiverBase";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "_dispatcher",
        type: "address",
        internalType: "contract IbcDispatcher",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "receive",
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "dispatcher",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "contract IbcDispatcher",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "owner",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "renounceOwnership",
    inputs: [],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "transferOwnership",
    inputs: [
      {
        name: "newOwner",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    name: "OwnershipTransferred",
    inputs: [
      {
        name: "previousOwner",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "newOwner",
        type: "address",
        indexed: true,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "error",
    name: "ChannelNotFound",
    inputs: [],
  },
  {
    type: "error",
    name: "UnsupportedVersion",
    inputs: [],
  },
  {
    type: "error",
    name: "notIbcDispatcher",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x608060405234801561001057600080fd5b5060405161036a38038061036a83398101604081905261002f916100ad565b6100383361005d565b600180546001600160a01b0319166001600160a01b03929092169190911790556100dd565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100bf57600080fd5b81516001600160a01b03811681146100d657600080fd5b9392505050565b61027e806100ec6000396000f3fe6080604052600436106100435760003560e01c8063715018a61461004f5780638da5cb5b14610066578063cb7e90571461009c578063f2fde38b146100bc57600080fd5b3661004a57005b600080fd5b34801561005b57600080fd5b506100646100dc565b005b34801561007257600080fd5b506000546001600160a01b03165b6040516001600160a01b03909116815260200160405180910390f35b3480156100a857600080fd5b50600154610080906001600160a01b031681565b3480156100c857600080fd5b506100646100d7366004610218565b6100f0565b6100e461016e565b6100ee60006101c8565b565b6100f861016e565b6001600160a01b0381166101625760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b61016b816101c8565b50565b6000546001600160a01b031633146100ee5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610159565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561022a57600080fd5b81356001600160a01b038116811461024157600080fd5b939250505056fea2646970667358221220298e19616943284d6862fcdc37f08ddd079032c462ca1b0e0cec6150837ef0db64736f6c634300080f0033";

type IbcReceiverBaseConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: IbcReceiverBaseConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class IbcReceiverBase__factory extends ContractFactory {
  constructor(...args: IbcReceiverBaseConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    _dispatcher: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(_dispatcher, overrides || {});
  }
  override deploy(
    _dispatcher: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(_dispatcher, overrides || {}) as Promise<
      IbcReceiverBase & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): IbcReceiverBase__factory {
    return super.connect(runner) as IbcReceiverBase__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): IbcReceiverBaseInterface {
    return new Interface(_abi) as IbcReceiverBaseInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): IbcReceiverBase {
    return new Contract(address, _abi, runner) as unknown as IbcReceiverBase;
  }
}
