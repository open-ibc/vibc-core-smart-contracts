//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "./Ibc.sol";
import "./IbcDispatcher.sol";
import "./IbcReceiver.sol";

/**
 * @title IbcMiddleware
 * @notice IBC middleware interface must be implemented by a IBC-middleware contract, similar to ICS-30
 *  https://github.com/cosmos/ibc/tree/main/spec/app/ics-030-middleware.
 */
interface IbcMiddleware is IbcPacketSender, IbcPacketReceiver {}

/**
 * @title IbcUniversalChannelMW
 * @notice This interface must be implemented by IBC universal channel middleware contracts.
 */
interface IbcUniversalChannelMW is IbcMiddleware, IbcChannelReceiver {}
