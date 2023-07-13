//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import './Ibc.sol';
import './IbcReceiver.sol';

struct Proof {
    // block height at which the proof is valid for a membership or non-membership at the given keyPath
    uint64 proofHeight;
    // ics23 merkle proof
    bytes proof;
}

/**
 * @title IbcDispatcher
 * @author Polymer Labs
 * @notice IBC dispatcher interface is the Polymer Core Smart Contract that implements the core IBC protocol.
 */
interface IbcDispatcher {
    function closeIbcChannel(bytes32 channelId) external;

    function sendPacket(
        bytes32 channelId,
        bytes calldata payload,
        uint64 timeoutTimestamp,
        PacketFee calldata fee
    ) external payable;
}
