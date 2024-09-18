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

export interface ILightClientInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "getState"
      | "updateClient"
      | "verifyMembership"
      | "verifyNonMembership"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "getState",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "updateClient",
    values: [BytesLike, BigNumberish, BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "verifyMembership",
    values: [Ics23ProofStruct, BytesLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "verifyNonMembership",
    values: [Ics23ProofStruct, BytesLike]
  ): string;

  decodeFunctionResult(functionFragment: "getState", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "updateClient",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "verifyMembership",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "verifyNonMembership",
    data: BytesLike
  ): Result;
}

export interface ILightClient extends BaseContract {
  connect(runner?: ContractRunner | null): ILightClient;
  waitForDeployment(): Promise<this>;

  interface: ILightClientInterface;

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

  getState: TypedContractMethod<[height: BigNumberish], [bigint], "view">;

  updateClient: TypedContractMethod<
    [proof: BytesLike, height: BigNumberish, appHash: BigNumberish],
    [void],
    "nonpayable"
  >;

  verifyMembership: TypedContractMethod<
    [proof: Ics23ProofStruct, key: BytesLike, expectedValue: BytesLike],
    [void],
    "nonpayable"
  >;

  verifyNonMembership: TypedContractMethod<
    [proof: Ics23ProofStruct, key: BytesLike],
    [void],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "getState"
  ): TypedContractMethod<[height: BigNumberish], [bigint], "view">;
  getFunction(
    nameOrSignature: "updateClient"
  ): TypedContractMethod<
    [proof: BytesLike, height: BigNumberish, appHash: BigNumberish],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "verifyMembership"
  ): TypedContractMethod<
    [proof: Ics23ProofStruct, key: BytesLike, expectedValue: BytesLike],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "verifyNonMembership"
  ): TypedContractMethod<
    [proof: Ics23ProofStruct, key: BytesLike],
    [void],
    "nonpayable"
  >;

  filters: {};
}
