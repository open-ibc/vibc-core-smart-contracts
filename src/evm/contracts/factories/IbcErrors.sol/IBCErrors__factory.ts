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
import type { NonPayableOverrides } from "../../common";
import type {
  IBCErrors,
  IBCErrorsInterface,
} from "../../IbcErrors.sol/IBCErrors";

const _abi = [
  {
    type: "error",
    name: "ackPacketCommitmentAlreadyExists",
    inputs: [],
  },
  {
    type: "error",
    name: "channelIdNotFound",
    inputs: [
      {
        name: "channelId",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
  },
  {
    type: "error",
    name: "channelNotOwnedByPortAddress",
    inputs: [],
  },
  {
    type: "error",
    name: "channelNotOwnedBySender",
    inputs: [],
  },
  {
    type: "error",
    name: "clientAlreadyCreated",
    inputs: [],
  },
  {
    type: "error",
    name: "clientNotCreated",
    inputs: [],
  },
  {
    type: "error",
    name: "consensusStateVerificationFailed",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidAddress",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidChannelType",
    inputs: [
      {
        name: "channelType",
        type: "string",
        internalType: "string",
      },
    ],
  },
  {
    type: "error",
    name: "invalidCharacter",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidConnection",
    inputs: [
      {
        name: "connection",
        type: "string",
        internalType: "string",
      },
    ],
  },
  {
    type: "error",
    name: "invalidConnectionHops",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidCounterParty",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidCounterPartyPortId",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidHexStringLength",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidPacket",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidPacketSequence",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidPortPrefix",
    inputs: [],
  },
  {
    type: "error",
    name: "invalidRelayerAddress",
    inputs: [],
  },
  {
    type: "error",
    name: "lightClientNotFound",
    inputs: [
      {
        name: "connection",
        type: "string",
        internalType: "string",
      },
    ],
  },
  {
    type: "error",
    name: "notEnoughGas",
    inputs: [],
  },
  {
    type: "error",
    name: "packetCommitmentNotFound",
    inputs: [],
  },
  {
    type: "error",
    name: "packetNotTimedOut",
    inputs: [],
  },
  {
    type: "error",
    name: "packetReceiptAlreadyExists",
    inputs: [],
  },
  {
    type: "error",
    name: "receiverNotIntendedPacketDestination",
    inputs: [],
  },
  {
    type: "error",
    name: "receiverNotOriginPacketSender",
    inputs: [],
  },
  {
    type: "error",
    name: "unexpectedPacketSequence",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220956f69820b7b4e35fabdb845c5f93ab77c0261dbc53ef31f5dd0473a6c437b4864736f6c634300080f0033";

type IBCErrorsConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: IBCErrorsConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class IBCErrors__factory extends ContractFactory {
  constructor(...args: IBCErrorsConstructorParams) {
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
      IBCErrors & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): IBCErrors__factory {
    return super.connect(runner) as IBCErrors__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): IBCErrorsInterface {
    return new Interface(_abi) as IBCErrorsInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): IBCErrors {
    return new Contract(address, _abi, runner) as unknown as IBCErrors;
  }
}