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

export type L1HeaderStruct = {
  header: BytesLike[];
  stateRoot: BytesLike;
  number: BigNumberish;
};

export type L1HeaderStructOutput = [
  header: string[],
  stateRoot: string,
  number: bigint
] & { header: string[]; stateRoot: string; number: bigint };

export type OpL2StateProofStruct = {
  accountProof: BytesLike[];
  outputRootProof: BytesLike[];
  l2OutputProposalKey: BytesLike;
  l2BlockHash: BytesLike;
};

export type OpL2StateProofStructOutput = [
  accountProof: string[],
  outputRootProof: string[],
  l2OutputProposalKey: string,
  l2BlockHash: string
] & {
  accountProof: string[];
  outputRootProof: string[];
  l2OutputProposalKey: string;
  l2BlockHash: string;
};

export interface DummyProofVerifierInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "verifyMembership"
      | "verifyNonMembership"
      | "verifyStateUpdate"
  ): FunctionFragment;

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
    values: [
      L1HeaderStruct,
      OpL2StateProofStruct,
      BytesLike,
      BytesLike,
      BigNumberish
    ]
  ): string;

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

export interface DummyProofVerifier extends BaseContract {
  connect(runner?: ContractRunner | null): DummyProofVerifier;
  waitForDeployment(): Promise<this>;

  interface: DummyProofVerifierInterface;

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

  verifyMembership: TypedContractMethod<
    [arg0: BytesLike, arg1: BytesLike, arg2: BytesLike, arg3: Ics23ProofStruct],
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
      arg0: L1HeaderStruct,
      arg1: OpL2StateProofStruct,
      arg2: BytesLike,
      arg3: BytesLike,
      arg4: BigNumberish
    ],
    [void],
    "view"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "verifyMembership"
  ): TypedContractMethod<
    [arg0: BytesLike, arg1: BytesLike, arg2: BytesLike, arg3: Ics23ProofStruct],
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
      arg0: L1HeaderStruct,
      arg1: OpL2StateProofStruct,
      arg2: BytesLike,
      arg3: BytesLike,
      arg4: BigNumberish
    ],
    [void],
    "view"
  >;

  filters: {};
}
