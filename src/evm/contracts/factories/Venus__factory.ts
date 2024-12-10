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
    name: "emitEvent",
    inputs: [
      {
        name: "sourceChainID",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "destinationChainID",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [],
    stateMutability: "payable",
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
    name: "SendHealthCheckEvent",
    inputs: [
      {
        name: "id",
        type: "bytes32",
        indexed: false,
        internalType: "bytes32",
      },
      {
        name: "sourceChainID",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "destChainID",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
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
  "0x60a060405234801561001057600080fd5b5060405161081c38038061081c83398101604081905261002f91610067565b6001600160a01b03811661005657604051637169bd2b60e11b815260040160405180910390fd5b6001600160a01b0316608052610097565b60006020828403121561007957600080fd5b81516001600160a01b038116811461009057600080fd5b9392505050565b6080516107566100c66000396000818160ab01528181610100015281816101b6015261033301526107566000f3fe60806040526004361061003f5760003560e01c8063163be0d8146100445780631a9c5187146100665780632004bbf01461007957806332a8f30f14610099575b600080fd5b34801561005057600080fd5b5061006461005f366004610460565b6100e9565b005b610064610074366004610530565b61029e565b34801561008557600080fd5b50610064610094366004610552565b61031c565b3480156100a557600080fd5b506100cd7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b6040516342a9943960e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063855328729061013f908c908c908c908c9089908990600401610615565b6020604051808303816000875af115801561015e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610182919061065e565b61019f576040516319e2b9bb60e31b815260040160405180910390fd5b604051632073934d60e21b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906381ce4d34906101fb908c908c908c908c908c908c908c908c908c90600401610687565b6020604051808303816000875af115801561021a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061023e919061065e565b61025b57604051638bf0861d60e01b815260040160405180910390fd5b604080518681523360208201527f49a61a3517534657df66eaaaef62f55a830c07d22ca1760e0eff4a2c823e0bc9910160405180910390a1505050505050505050565b604080516020810184905290810182905242606082015243608082015260009060a00160408051601f19818403018152828252805160209182012080845290830186905290820184905291507f2aa5701af4665db6d0d94b04beb302db15e10382304fe5195aeefb9c7dd1808b9060600160405180910390a1505050565b6040516342a9943960e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063855328729061037290899089908990899089908990600401610615565b6020604051808303816000875af1158015610391573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b5919061065e565b6103d2576040516319e2b9bb60e31b815260040160405180910390fd5b7fd6728eaf25dd1431eb8afabc6f371b3379a6c7beb972e468f5bf992fc6e822d58686868660405161040794939291906106ee565b60405180910390a1505050505050565b60008083601f84011261042957600080fd5b50813567ffffffffffffffff81111561044157600080fd5b60208301915083602082850101111561045957600080fd5b9250929050565b600080600080600080600080600060a08a8c03121561047e57600080fd5b893567ffffffffffffffff8082111561049657600080fd5b6104a28d838e01610417565b909b50995060208c01359150808211156104bb57600080fd5b6104c78d838e01610417565b909950975060408c0135965060608c01359150808211156104e757600080fd5b6104f38d838e01610417565b909650945060808c013591508082111561050c57600080fd5b506105198c828d01610417565b915080935050809150509295985092959850929598565b6000806040838503121561054357600080fd5b50508035926020909101359150565b6000806000806000806060878903121561056b57600080fd5b863567ffffffffffffffff8082111561058357600080fd5b61058f8a838b01610417565b909850965060208901359150808211156105a857600080fd5b6105b48a838b01610417565b909650945060408901359150808211156105cd57600080fd5b506105da89828a01610417565b979a9699509497509295939492505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60608152600061062960608301888a6105ec565b828103602084015261063c8187896105ec565b905082810360408401526106518185876105ec565b9998505050505050505050565b60006020828403121561067057600080fd5b8151801515811461068057600080fd5b9392505050565b60a08152600061069b60a083018b8d6105ec565b82810360208401526106ae818a8c6105ec565b905087604084015282810360608401526106c98187896105ec565b905082810360808401526106de8185876105ec565b9c9b505050505050505050505050565b6040815260006107026040830186886105ec565b82810360208401526107158185876105ec565b97965050505050505056fea26469706673582212201ea28c1bd14bae1bcd88487e24578960e8fd6d3deb9870f58f980ca2e2533a2564736f6c634300080f0033";

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
