/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumberish,
  BytesLike,
  FunctionFragment,
  Result,
  Interface,
  ContractRunner,
  ContractMethod,
  Listener,
} from "ethers";
import type {
  TypedContractEvent,
  TypedDeferredTopicFilter,
  TypedEventLog,
  TypedListener,
  TypedContractMethod,
} from "./common";

export type OpIcs23ProofPathStruct = { prefix: BytesLike; suffix: BytesLike };

export type OpIcs23ProofPathStructOutput = [prefix: string, suffix: string] & {
  prefix: string;
  suffix: string;
};

export type OpIcs23ProofStruct = {
  path: OpIcs23ProofPathStruct[];
  key: BytesLike;
  value: BytesLike;
  prefix: BytesLike;
};

export type OpIcs23ProofStructOutput = [
  path: OpIcs23ProofPathStructOutput[],
  key: string,
  value: string,
  prefix: string
] & {
  path: OpIcs23ProofPathStructOutput[];
  key: string;
  value: string;
  prefix: string;
};

export type Ics23ProofStruct = {
  proof: OpIcs23ProofStruct[];
  height: BigNumberish;
};

export type Ics23ProofStructOutput = [
  proof: OpIcs23ProofStructOutput[],
  height: bigint
] & { proof: OpIcs23ProofStructOutput[]; height: bigint };

export interface SequencerSignatureVerifierInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "CHAIN_ID"
      | "SEQUENCER"
      | "verifyMembership"
      | "verifyNonMembership"
      | "verifyStateUpdate"
  ): FunctionFragment;

  encodeFunctionData(functionFragment: "CHAIN_ID", values?: undefined): string;
  encodeFunctionData(functionFragment: "SEQUENCER", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "verifyMembership",
    values: [BytesLike, BytesLike, BytesLike, Ics23ProofStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "verifyNonMembership",
    values: [BytesLike, BytesLike, Ics23ProofStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "verifyStateUpdate",
    values: [BigNumberish, BytesLike, BytesLike, BytesLike]
  ): string;

  decodeFunctionResult(functionFragment: "CHAIN_ID", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "SEQUENCER", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "verifyMembership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "verifyNonMembership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "verifyStateUpdate",
    data: BytesLike
  ): Result;
}

export interface SequencerSignatureVerifier extends BaseContract {
  connect(runner?: ContractRunner | null): SequencerSignatureVerifier;
  waitForDeployment(): Promise<this>;

  interface: SequencerSignatureVerifierInterface;

  queryFilter<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;
  queryFilter<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;

  on<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  on<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  once<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  once<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  listeners<TCEvent extends TypedContractEvent>(
    event: TCEvent
  ): Promise<Array<TypedListener<TCEvent>>>;
  listeners(eventName?: string): Promise<Array<Listener>>;
  removeAllListeners<TCEvent extends TypedContractEvent>(
    event?: TCEvent
  ): Promise<this>;

  CHAIN_ID: TypedContractMethod<[], [string], "view">;

  SEQUENCER: TypedContractMethod<[], [string], "view">;

  verifyMembership: TypedContractMethod<
    [
      appHash: BytesLike,
      key: BytesLike,
      value: BytesLike,
      proofs: Ics23ProofStruct
    ],
    [void],
    "view"
  >;

  verifyNonMembership: TypedContractMethod<
    [arg0: BytesLike, arg1: BytesLike, arg2: Ics23ProofStruct],
    [void],
    "view"
  >;

  verifyStateUpdate: TypedContractMethod<
    [
      l2BlockNumber: BigNumberish,
      appHash: BytesLike,
      l1BlockHash: BytesLike,
      signature: BytesLike
    ],
    [void],
    "view"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "CHAIN_ID"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "SEQUENCER"
  ): TypedContractMethod<[], [string], "view">;
  getFunction(
    nameOrSignature: "verifyMembership"
  ): TypedContractMethod<
    [
      appHash: BytesLike,
      key: BytesLike,
      value: BytesLike,
      proofs: Ics23ProofStruct
    ],
    [void],
    "view"
  >;
  getFunction(
    nameOrSignature: "verifyNonMembership"
  ): TypedContractMethod<
    [arg0: BytesLike, arg1: BytesLike, arg2: Ics23ProofStruct],
    [void],
    "view"
  >;
  getFunction(
    nameOrSignature: "verifyStateUpdate"
  ): TypedContractMethod<
    [
      l2BlockNumber: BigNumberish,
      appHash: BytesLike,
      l1BlockHash: BytesLike,
      signature: BytesLike
    ],
    [void],
    "view"
  >;

  filters: {};
}
