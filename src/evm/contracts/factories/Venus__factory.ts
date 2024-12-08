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
import type { NonPayableOverrides } from "../common";
import type { Venus, VenusInterface } from "../Venus";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "_prover",
        type: "address",
        internalType: "contract ICrossL2Prover",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "prover",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "contract ICrossL2Prover",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "receiveEvent",
    inputs: [
      {
        name: "receiptIndex",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "receiptRLPEncodedBytes",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "logIndex",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "logBytes",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "proof",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "receiveReceipt",
    inputs: [
      {
        name: "receiptIndex",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "receiptRLPEncodedBytes",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "proof",
        type: "bytes",
        internalType: "bytes",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    name: "SuccessfulEvent",
    inputs: [
      {
        name: "eventIndex",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "sender",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "SuccessfulReceipt",
    inputs: [
      {
        name: "receiptIndex",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "receiptRLP",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
    ],
    anonymous: false,
  },
  {
    type: "error",
    name: "invalidEventProof",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidProverAddress",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidReceiptProof",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x60a060405234801561001057600080fd5b5060405161074338038061074383398101604081905261002f91610067565b6001600160a01b03811661005657604051637169bd2b60e11b815260040160405180910390fd5b6001600160a01b0316608052610097565b60006020828403121561007957600080fd5b81516001600160a01b038116811461009057600080fd5b9392505050565b60805161067e6100c56000396000818160730152818160c80152818161017e015261027d015261067e6000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063163be0d8146100465780632004bbf01461005b57806332a8f30f1461006e575b600080fd5b6100596100543660046103aa565b6100b1565b005b61005961006936600461047a565b610266565b6100957f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b6040516342a9943960e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638553287290610107908c908c908c908c908990899060040161053d565b6020604051808303816000875af1158015610126573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061014a9190610586565b610167576040516319e2b9bb60e31b815260040160405180910390fd5b604051632073934d60e21b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906381ce4d34906101c3908c908c908c908c908c908c908c908c908c906004016105af565b6020604051808303816000875af11580156101e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102069190610586565b61022357604051638bf0861d60e01b815260040160405180910390fd5b604080518681523360208201527f49a61a3517534657df66eaaaef62f55a830c07d22ca1760e0eff4a2c823e0bc9910160405180910390a1505050505050505050565b6040516342a9943960e11b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906385532872906102bc9089908990899089908990899060040161053d565b6020604051808303816000875af11580156102db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102ff9190610586565b61031c576040516319e2b9bb60e31b815260040160405180910390fd5b7fd6728eaf25dd1431eb8afabc6f371b3379a6c7beb972e468f5bf992fc6e822d5868686866040516103519493929190610616565b60405180910390a1505050505050565b60008083601f84011261037357600080fd5b50813567ffffffffffffffff81111561038b57600080fd5b6020830191508360208285010111156103a357600080fd5b9250929050565b600080600080600080600080600060a08a8c0312156103c857600080fd5b893567ffffffffffffffff808211156103e057600080fd5b6103ec8d838e01610361565b909b50995060208c013591508082111561040557600080fd5b6104118d838e01610361565b909950975060408c0135965060608c013591508082111561043157600080fd5b61043d8d838e01610361565b909650945060808c013591508082111561045657600080fd5b506104638c828d01610361565b915080935050809150509295985092959850929598565b6000806000806000806060878903121561049357600080fd5b863567ffffffffffffffff808211156104ab57600080fd5b6104b78a838b01610361565b909850965060208901359150808211156104d057600080fd5b6104dc8a838b01610361565b909650945060408901359150808211156104f557600080fd5b5061050289828a01610361565b979a9699509497509295939492505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60608152600061055160608301888a610514565b8281036020840152610564818789610514565b90508281036040840152610579818587610514565b9998505050505050505050565b60006020828403121561059857600080fd5b815180151581146105a857600080fd5b9392505050565b60a0815260006105c360a083018b8d610514565b82810360208401526105d6818a8c610514565b905087604084015282810360608401526105f1818789610514565b90508281036080840152610606818587610514565b9c9b505050505050505050505050565b60408152600061062a604083018688610514565b828103602084015261063d818587610514565b97965050505050505056fea2646970667358221220fc3a98c7a6dc3f76a337caab2118923fec04bb8a171652adfaaa6f7a21b7fa0a64736f6c634300080f0033";

type VenusConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: VenusConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class Venus__factory extends ContractFactory {
  constructor(...args: VenusConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    _prover: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(_prover, overrides || {});
  }
  override deploy(
    _prover: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(_prover, overrides || {}) as Promise<
      Venus & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): Venus__factory {
    return super.connect(runner) as Venus__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): VenusInterface {
    return new Interface(_abi) as VenusInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): Venus {
    return new Contract(address, _abi, runner) as unknown as Venus;
  }
}
