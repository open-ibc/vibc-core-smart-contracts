//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import {DispatcherV2} from "./DispatcherV2.sol";
import {LightClient} from "../../../contracts/interfaces/LightClient.sol";
import {IBCErrors} from "../../../contracts/libs/Ibc.sol";
/**
 * @title Dispatcher
 * @author Polymer Labs
 * @notice
 *     Contract callers call this contract to send IBC-like msg,
 *     which can be relayed to a rollup module on the Polymerase chain
 */

contract DispatcherV2Initializable is DispatcherV2 {
    function initialize(string memory initPortPrefix) public override reinitializer(2) {
        __Ownable_init();
        portPrefix = initPortPrefix;
        portPrefixLen = uint32(bytes(initPortPrefix).length);
    }
}
