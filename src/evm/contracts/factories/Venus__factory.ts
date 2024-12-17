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
  BytesLike,
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
      {
        name: "_counterParty",
        type: "address",
        internalType: "address",
      },
      {
        name: "_chainId",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "chainId",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "counterParty",
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
    name: "lastReceivedTransmission",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    stateMutability: "view",
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
        name: "logIndex",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "proof",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "expectedEmitter",
        type: "address",
        internalType: "address",
      },
      {
        name: "expectedTopics",
        type: "bytes[]",
        internalType: "bytes[]",
      },
      {
        name: "expectedUnindexedData",
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
    name: "receiveTransmissionEvent",
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
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    name: "SuccessfulReceipt",
    inputs: [
      {
        name: "srcChainId",
        type: "bytes32",
        indexed: false,
        internalType: "bytes32",
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
    type: "event",
    name: "TransmissionReceived",
    inputs: [
      {
        name: "message",
        type: "bytes32",
        indexed: false,
        internalType: "bytes32",
      },
      {
        name: "timestamp",
        type: "uint64",
        indexed: false,
        internalType: "uint64",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "TransmitToHouston",
    inputs: [
      {
        name: "message",
        type: "bytes32",
        indexed: false,
        internalType: "bytes32",
      },
      {
        name: "timestamp",
        type: "uint64",
        indexed: false,
        internalType: "uint64",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "ValidCounterpartyEvent",
    inputs: [
      {
        name: "counterParty",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "topics",
        type: "bytes[]",
        indexed: false,
        internalType: "bytes[]",
      },
      {
        name: "unindexed",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
    ],
    anonymous: false,
  },
  {
    type: "error",
    name: "invalidChainId",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidCounterpartyEvent",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidEventProof",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidEventSender",
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
  "0x60e060405234801561001057600080fd5b50604051610e7b380380610e7b83398101604081905261002f91610089565b6001600160a01b03831661005657604051637169bd2b60e11b815260040160405180910390fd5b6001600160a01b03928316608052911660a05260c0526100cc565b6001600160a01b038116811461008657600080fd5b50565b60008060006060848603121561009e57600080fd5b83516100a981610071565b60208501519093506100ba81610071565b80925050604084015190509250925092565b60805160a05160c051610d586101236000396000818161011a015281816102e50152610517015260008181609c015261055601526000818160e00152818161016b0152818161024701526104790152610d586000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063932e6cb21161005b578063932e6cb2146101025780639a8a059214610115578063a17d42d41461014a578063e03d0aac1461015d57600080fd5b8063273533e1146100825780632c453bdf1461009757806332a8f30f146100db575b600080fd5b6100956100903660046106fd565b610166565b005b6100be7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6100be7f000000000000000000000000000000000000000000000000000000000000000081565b6100956101103660046107ea565b61023f565b61013c7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016100d2565b610095610158366004610961565b610471565b61013c60005481565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632cd78e7785856040518363ffffffff1660e01b81526004016101b79291906109d6565b600060405180830381865afa1580156101d4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526101fc9190810190610a67565b915091507ff771e200728f07a661e482ccfec772812dfe16ba618778746183c6622500118b8282604051610231929190610ada565b60405180910390a150505050565b6000806000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fc396ddb8c8c8c6040518463ffffffff1660e01b815260040161029593929190610af3565b600060405180830381865afa1580156102b2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526102da9190810190610b16565b9350935093509350837f00000000000000000000000000000000000000000000000000000000000000001461032257604051632da6f64760e11b815260040160405180910390fd5b876001600160a01b0316836001600160a01b031614610354576040516304b5dc5f60e11b815260040160405180910390fd5b6103ad826040516020016103689190610c5f565b604051602081830303815290604052886040516020016103889190610c5f565b6040516020818303038152906040528051602091820120825192909101919091201490565b6103ca5760405163596128d160e01b815260040160405180910390fd5b61040a8187878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061069e92505050565b6104275760405163596128d160e01b815260040160405180910390fd5b7fe683aa4424de0611da9ae18b5fc355ae68fc1dca14a14c5d7e2ad9ba3711d4468383888860405161045c9493929190610c79565b60405180910390a15050505050505050505050565b6000806000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fc396ddb8888886040518463ffffffff1660e01b81526004016104c793929190610af3565b600060405180830381865afa1580156104e4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261050c9190810190610b16565b9350935093509350837f00000000000000000000000000000000000000000000000000000000000000001461055457604051632da6f64760e11b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316146105a6576040516304b5dc5f60e11b815260040160405180910390fd5b7fc134cee5dddcc539de40a587edc2d37fb04e8ec090cb8dcf8ef0c9b23d41908d826000815181106105da576105da610cbb565b60200260200101516105eb90610cd1565b146106095760405163596128d160e01b815260040160405180910390fd5b60008260028151811061061e5761061e610cbb565b602002602001015161062f90610cd1565b90508060008190555060008280602001905181019061064e9190610cf8565b6040805184815267ffffffffffffffff831660208201529192507f777932fb4709c8bfb29f9190e22e536eaf00fd97a76a15a41b54e1b619fe5863910160405180910390a1505050505050505050565b8051602091820120825192909101919091201490565b60008083601f8401126106c657600080fd5b50813567ffffffffffffffff8111156106de57600080fd5b6020830191508360208285010111156106f657600080fd5b9250929050565b6000806020838503121561071057600080fd5b823567ffffffffffffffff81111561072757600080fd5b610733858286016106b4565b90969095509350505050565b6001600160a01b038116811461075457600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561079657610796610757565b604052919050565b600067ffffffffffffffff8211156107b8576107b8610757565b5060051b60200190565b600067ffffffffffffffff8211156107dc576107dc610757565b50601f01601f191660200190565b600080600080600080600060a0888a03121561080557600080fd5b8735965067ffffffffffffffff8060208a0135111561082357600080fd5b6108338a60208b01358b016106b4565b909750955061084560408a013561073f565b604089013594508060608a0135111561085d57600080fd5b606089013589018a601f82011261087357600080fd5b610885610880823561079e565b61076d565b81358082526020808301929160051b8401018d10156108a357600080fd5b602083015b6020843560051b8501018110156109285784813511156108c757600080fd5b8d603f8235860101126108d957600080fd5b60208135850101356108ed610880826107c2565b8181528f604083853589010101111561090557600080fd5b8160408435880101602083013760006020928201830152845292830192016108a8565b5095505050608089013581101561093e57600080fd5b5061094f8960808a01358a016106b4565b979a9699509497509295919491925050565b60008060006040848603121561097657600080fd5b83359250602084013567ffffffffffffffff81111561099457600080fd5b6109a0868287016106b4565b9497909650939450505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6020815260006109ea6020830184866109ad565b949350505050565b60005b83811015610a0d5781810151838201526020016109f5565b83811115610a1c576000848401525b50505050565b600082601f830112610a3357600080fd5b8151610a41610880826107c2565b818152846020838601011115610a5657600080fd5b6109ea8260208301602087016109f2565b60008060408385031215610a7a57600080fd5b82519150602083015167ffffffffffffffff811115610a9857600080fd5b610aa485828601610a22565b9150509250929050565b60008151808452610ac68160208601602086016109f2565b601f01601f19169290920160200192915050565b8281526040602082015260006109ea6040830184610aae565b838152604060208201526000610b0d6040830184866109ad565b95945050505050565b60008060008060808587031215610b2c57600080fd5b84519350602080860151610b3f8161073f565b604087015190945067ffffffffffffffff80821115610b5d57600080fd5b818801915088601f830112610b7157600080fd5b8151610b7f6108808261079e565b81815260059190911b8301840190848101908b831115610b9e57600080fd5b8585015b83811015610bd657805185811115610bba5760008081fd5b610bc88e89838a0101610a22565b845250918601918601610ba2565b5060608b01519097509450505080831115610bf057600080fd5b5050610bfe87828801610a22565b91505092959194509250565b600081518084526020808501808196508360051b8101915082860160005b85811015610c52578284038952610c40848351610aae565b98850198935090840190600101610c28565b5091979650505050505050565b602081526000610c726020830184610c0a565b9392505050565b6001600160a01b0385168152606060208201819052600090610c9d90830186610c0a565b8281036040840152610cb08185876109ad565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b80516020808301519190811015610cf2576000198160200360031b1b821691505b50919050565b600060208284031215610d0a57600080fd5b815167ffffffffffffffff81168114610c7257600080fdfea2646970667358221220406a1d451c50756bf6472cd8297c4b40cacc118e9fb8ad895369f927d758855264736f6c634300080f0033";

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
    _counterParty: AddressLike,
    _chainId: BytesLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(
      _prover,
      _counterParty,
      _chainId,
      overrides || {}
    );
  }
  override deploy(
    _prover: AddressLike,
    _counterParty: AddressLike,
    _chainId: BytesLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(
      _prover,
      _counterParty,
      _chainId,
      overrides || {}
    ) as Promise<
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
