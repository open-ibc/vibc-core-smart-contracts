//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import {OwnableUpgradeable} from "@openzeppelin-upgradeable/contracts/access/OwnableUpgradeable.sol";
import {IbcDispatcher} from "./IbcDispatcher.sol";
import {ChannelOrder, ChannelEnd, IbcPacket, AckPacket} from "../libs/Ibc.sol";

contract IbcReceiverBaseUpgradeable is OwnableUpgradeable {
    IbcDispatcher public dispatcher;

    error notIbcDispatcher();
    error UnsupportedVersion();
    error ChannelNotFound();

    /**
     * @dev Modifier to restrict access to only the IBC dispatcher.
     * Only the address with the IBC_ROLE can execute the function.
     * Should add this modifier to all IBC-related callback functions.
     */
    modifier onlyIbcDispatcher() {
        if (msg.sender != address(dispatcher)) {
            revert notIbcDispatcher();
        }
        _;
    }

    constructor() {
        _disableInitializers();
    }

    /// This function is called for plain Ether transfers, i.e. for every call with empty calldata.
    // An empty function body is sufficient to receive packet fee refunds.
    receive() external payable {}

    /**
     * @dev initializer function that takes an IbcDispatcher address and grants the IBC_ROLE to the Polymer IBC
     * Dispatcher.
     * @param _dispatcher The address of the IbcDispatcher contract.
     */
    function __IbcReceiverBase_init(IbcDispatcher _dispatcher) internal onlyInitializing {
        __Ownable_init();
        dispatcher = _dispatcher;
    }
}
