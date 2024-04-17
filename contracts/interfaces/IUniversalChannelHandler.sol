// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IbcDispatcher, IbcEventsEmitter} from "./IbcDispatcher.sol";

import {L1Header, OpL2StateProof, Ics23Proof} from "./ProofVerifier.sol";
import {IbcUniversalChannelMW} from "./IbcMiddleware.sol";
import {
    Channel,
    ChannelEnd,
    ChannelOrder,
    IbcPacket,
    ChannelState,
    AckPacket,
    IBCErrors,
    IbcUtils,
    Ibc
} from "../libs/Ibc.sol";

interface IUniversalChannelHandler is IbcUniversalChannelMW {
    function registerMwStack(uint256 mwBitmap, address[] calldata mwAddrs) external;
    function openChannel(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortIdentifier
    ) external;
    function closeChannel(bytes32 channelId) external;
    function setDispatcher(IbcDispatcher dispatcher) external;
    function dispatcher() external returns (IbcDispatcher dispatcher);
    function connectedChannels(uint256) external view returns (bytes32 channel);
}
