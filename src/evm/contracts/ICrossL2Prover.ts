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

export interface ICrossL2ProverInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "LIGHT_CLIENT_TYPE"
      | "getState"
      | "updateClient"
      | "validateEvent"
      | "validateReceipt"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "LIGHT_CLIENT_TYPE",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "getState",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "updateClient",
    values: [BytesLike, BigNumberish, BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "validateEvent",
    values: [BytesLike, BytesLike, BigNumberish, BytesLike, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "validateReceipt",
    values: [BytesLike, BytesLike, BytesLike]
  ): string;

  decodeFunctionResult(
    functionFragment: "LIGHT_CLIENT_TYPE",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "getState", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "updateClient",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "validateEvent",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "validateReceipt",
    data: BytesLike
  ): Result;
}

export interface ICrossL2Prover extends BaseContract {
  connect(runner?: ContractRunner | null): ICrossL2Prover;
  waitForDeployment(): Promise<this>;

  interface: ICrossL2ProverInterface;

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

  LIGHT_CLIENT_TYPE: TypedContractMethod<[], [bigint], "view">;

  getState: TypedContractMethod<[height: BigNumberish], [bigint], "view">;

  updateClient: TypedContractMethod<
    [proof: BytesLike, height: BigNumberish, appHash: BigNumberish],
    [void],
    "nonpayable"
  >;

  validateEvent: TypedContractMethod<
    [
      receiptIndex: BytesLike,
      receiptRLPEncodedBytes: BytesLike,
      logIndex: BigNumberish,
      logBytes: BytesLike,
      proof: BytesLike
    ],
    [boolean],
    "nonpayable"
  >;

  validateReceipt: TypedContractMethod<
    [
      receiptIndex: BytesLike,
      receiptRLPEncodedBytes: BytesLike,
      proof: BytesLike
    ],
    [boolean],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "LIGHT_CLIENT_TYPE"
  ): TypedContractMethod<[], [bigint], "view">;
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
    nameOrSignature: "validateEvent"
  ): TypedContractMethod<
    [
      receiptIndex: BytesLike,
      receiptRLPEncodedBytes: BytesLike,
      logIndex: BigNumberish,
      logBytes: BytesLike,
      proof: BytesLike
    ],
    [boolean],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "validateReceipt"
  ): TypedContractMethod<
    [
      receiptIndex: BytesLike,
      receiptRLPEncodedBytes: BytesLike,
      proof: BytesLike
    ],
    [boolean],
    "nonpayable"
  >;

  filters: {};
}
