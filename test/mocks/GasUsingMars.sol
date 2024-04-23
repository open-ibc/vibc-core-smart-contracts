// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {IDispatcher} from "../../contracts/interfaces/IDispatcher.sol";
import {Mars} from "../../contracts/examples/Mars.sol";
import {IbcPacket, AckPacket} from "../../contracts/libs/Ibc.sol";
import "forge-std/Test.sol";

/**
 * Testing mock contrantract to test gas useage.
 * This contract
 * We can label the true gas cost of the parent transaction sent to the dispatcher as g = a + b + c .
 *
 * Solidity automatically allocates 63/64 of the remaining gas left in a tx for low level calls, so (g-a)*63/64 will be
 * allocated for b,
 * and the remaining (g-a)/64 will be allocated for c.
 * If (g-a) * b > c and a malicious user sends a transaction with a gas limit g' such that  a + c < g'< (g' -
 * a)*(63/64),
 * this can result in a gas griefing attack.
 */
contract GasUsingMars is Mars {
    uint256 private gasToUse;

    constructor(uint256 _gasToUse, IDispatcher _ibcDispatcher) Mars(_ibcDispatcher) {
        gasToUse = _gasToUse;
    }

    function onRecvPacket(IbcPacket memory packet)
        external
        virtual
        override
        onlyIbcDispatcher
        returns (AckPacket memory ackPacket)
    {
        recvedPackets.push(packet);

        _useGas();
        // solhint-disable-next-line quotes
        return AckPacket(true, abi.encodePacked('{ "account": "account", "reply": "got the message" }'));
    }

    function _useGas() internal view {
        // This function is used to test the gas usage of the contract
        uint256 startingGas = gasleft();
        uint256 dummyInt = 0;

        // It will use the gasToUse amount of gas
        while (startingGas - gasleft() < gasToUse) {
            dummyInt += 1;
        }
    }
}
