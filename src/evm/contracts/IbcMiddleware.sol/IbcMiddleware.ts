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
  AddressLike,
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
} from "../common";

export type UniversalPacketStruct = {
  srcPortAddr: BytesLike;
  mwBitmap: BigNumberish;
  destPortAddr: BytesLike;
  appData: BytesLike;
};

export type UniversalPacketStructOutput = [
  srcPortAddr: string,
  mwBitmap: bigint,
  destPortAddr: string,
  appData: string
] & {
  srcPortAddr: string;
  mwBitmap: bigint;
  destPortAddr: string;
  appData: string;
};

export type AckPacketStruct = { success: boolean; data: BytesLike };

export type AckPacketStructOutput = [success: boolean, data: string] & {
  success: boolean;
  data: string;
};

export interface IbcMiddlewareInterface extends Interface {
  getFunction(
    nameOrSignature:
      | "MW_ID"
      | "onRecvMWAck"
      | "onRecvMWPacket"
      | "onRecvMWTimeout"
      | "onRecvUniversalPacket"
      | "onTimeoutUniversalPacket"
      | "onUniversalAcknowledgement"
      | "sendUniversalPacket"
      | "sendUniversalPacketWithFee"
  ): FunctionFragment;

  encodeFunctionData(functionFragment: "MW_ID", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "onRecvMWAck",
    values: [
      BytesLike,
      UniversalPacketStruct,
      BigNumberish,
      AddressLike[],
      AckPacketStruct
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "onRecvMWPacket",
    values: [BytesLike, UniversalPacketStruct, BigNumberish, AddressLike[]]
  ): string;
  encodeFunctionData(
    functionFragment: "onRecvMWTimeout",
    values: [BytesLike, UniversalPacketStruct, BigNumberish, AddressLike[]]
  ): string;
  encodeFunctionData(
    functionFragment: "onRecvUniversalPacket",
    values: [BytesLike, UniversalPacketStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "onTimeoutUniversalPacket",
    values: [BytesLike, UniversalPacketStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "onUniversalAcknowledgement",
    values: [BytesLike, UniversalPacketStruct, AckPacketStruct]
  ): string;
  encodeFunctionData(
    functionFragment: "sendUniversalPacket",
    values: [BytesLike, BytesLike, BytesLike, BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "sendUniversalPacketWithFee",
    values: [
      BytesLike,
      BytesLike,
      BytesLike,
      BigNumberish,
      [BigNumberish, BigNumberish],
      [BigNumberish, BigNumberish]
    ]
  ): string;

  decodeFunctionResult(functionFragment: "MW_ID", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "onRecvMWAck",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onRecvMWPacket",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onRecvMWTimeout",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onRecvUniversalPacket",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onTimeoutUniversalPacket",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "onUniversalAcknowledgement",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "sendUniversalPacket",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "sendUniversalPacketWithFee",
    data: BytesLike
  ): Result;
}

export interface IbcMiddleware extends BaseContract {
  connect(runner?: ContractRunner | null): IbcMiddleware;
  waitForDeployment(): Promise<this>;

  interface: IbcMiddlewareInterface;

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

  MW_ID: TypedContractMethod<[], [bigint], "view">;

  onRecvMWAck: TypedContractMethod<
    [
      channelId: BytesLike,
      ucPacket: UniversalPacketStruct,
      mwIndex: BigNumberish,
      mwAddrs: AddressLike[],
      ack: AckPacketStruct
    ],
    [void],
    "nonpayable"
  >;

  onRecvMWPacket: TypedContractMethod<
    [
      channelId: BytesLike,
      ucPacket: UniversalPacketStruct,
      mwIndex: BigNumberish,
      mwAddrs: AddressLike[]
    ],
    [AckPacketStructOutput],
    "nonpayable"
  >;

  onRecvMWTimeout: TypedContractMethod<
    [
      channelId: BytesLike,
      ucPacket: UniversalPacketStruct,
      mwIndex: BigNumberish,
      mwAddrs: AddressLike[]
    ],
    [void],
    "nonpayable"
  >;

  onRecvUniversalPacket: TypedContractMethod<
    [channelId: BytesLike, ucPacket: UniversalPacketStruct],
    [AckPacketStructOutput],
    "nonpayable"
  >;

  onTimeoutUniversalPacket: TypedContractMethod<
    [channelId: BytesLike, packet: UniversalPacketStruct],
    [void],
    "nonpayable"
  >;

  onUniversalAcknowledgement: TypedContractMethod<
    [channelId: BytesLike, packet: UniversalPacketStruct, ack: AckPacketStruct],
    [void],
    "nonpayable"
  >;

  sendUniversalPacket: TypedContractMethod<
    [
      channelId: BytesLike,
      destPortAddr: BytesLike,
      appData: BytesLike,
      timeoutTimestamp: BigNumberish
    ],
    [bigint],
    "nonpayable"
  >;

  sendUniversalPacketWithFee: TypedContractMethod<
    [
      channelId: BytesLike,
      destPortAddr: BytesLike,
      appData: BytesLike,
      timeoutTimestamp: BigNumberish,
      gasLimits: [BigNumberish, BigNumberish],
      gasPrices: [BigNumberish, BigNumberish]
    ],
    [bigint],
    "payable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "MW_ID"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "onRecvMWAck"
  ): TypedContractMethod<
    [
      channelId: BytesLike,
      ucPacket: UniversalPacketStruct,
      mwIndex: BigNumberish,
      mwAddrs: AddressLike[],
      ack: AckPacketStruct
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onRecvMWPacket"
  ): TypedContractMethod<
    [
      channelId: BytesLike,
      ucPacket: UniversalPacketStruct,
      mwIndex: BigNumberish,
      mwAddrs: AddressLike[]
    ],
    [AckPacketStructOutput],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onRecvMWTimeout"
  ): TypedContractMethod<
    [
      channelId: BytesLike,
      ucPacket: UniversalPacketStruct,
      mwIndex: BigNumberish,
      mwAddrs: AddressLike[]
    ],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onRecvUniversalPacket"
  ): TypedContractMethod<
    [channelId: BytesLike, ucPacket: UniversalPacketStruct],
    [AckPacketStructOutput],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onTimeoutUniversalPacket"
  ): TypedContractMethod<
    [channelId: BytesLike, packet: UniversalPacketStruct],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "onUniversalAcknowledgement"
  ): TypedContractMethod<
    [channelId: BytesLike, packet: UniversalPacketStruct, ack: AckPacketStruct],
    [void],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "sendUniversalPacket"
  ): TypedContractMethod<
    [
      channelId: BytesLike,
      destPortAddr: BytesLike,
      appData: BytesLike,
      timeoutTimestamp: BigNumberish
    ],
    [bigint],
    "nonpayable"
  >;
  getFunction(
    nameOrSignature: "sendUniversalPacketWithFee"
  ): TypedContractMethod<
    [
      channelId: BytesLike,
      destPortAddr: BytesLike,
      appData: BytesLike,
      timeoutTimestamp: BigNumberish,
      gasLimits: [BigNumberish, BigNumberish],
      gasPrices: [BigNumberish, BigNumberish]
    ],
    [bigint],
    "payable"
  >;

  filters: {};
}
