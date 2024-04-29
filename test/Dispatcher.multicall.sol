// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "../contracts/libs/Ibc.sol";
import {Dispatcher} from "../contracts/core/Dispatcher.sol";
import {IbcEventsEmitter} from "../contracts/interfaces/IbcDispatcher.sol";
import {IbcReceiver, IbcReceiverBase} from "../contracts/interfaces/IbcReceiver.sol";
import {DummyLightClient} from "../contracts/utils/DummyLightClient.sol";
import {
    Mars,
    RevertingBytesMars,
    PanickingMars,
    RevertingEmptyMars,
    RevertingStringMars
} from "../contracts/examples/Mars.sol";
import "../contracts/core/OpLightClient.sol";
import "./Dispatcher.base.t.sol";
import {Earth} from "../contracts/examples/Earth.sol";
import {Multicall3} from "@multicall/Multicall3.sol";

// Minor test to ensure that multicall is compatible with reentrancyguard, and to judge gas overhead
contract DispatcherMultiCallTest is Base {
    string portId;
    LocalEnd _local;
    Mars mars;
    ChannelEnd _remote;
    Multicall3 multicall;
    ChannelEnd ch0;
    ChannelEnd ch1;

    function setUp() public override {
        (dispatcherProxy, dispatcherImplementation) =
            TestUtilsTest.deployDispatcherProxyAndImpl(portPrefix, dummyConsStateManager);
        mars = new Mars(dispatcherProxy);
        string memory marsPort = IbcUtils.addressToPortId(portPrefix, address(mars));
        ch0 = ChannelEnd(marsPort, IbcUtils.toBytes32("channel-0"), "1.0");
        ch1 = ChannelEnd(marsPort, IbcUtils.toBytes32("channel-0"), "1.0");
        multicall = new Multicall3();
    }

    function test_dispatcher_multicall() public {
        Multicall3.Call3[] memory calls = new Multicall3.Call3[](2);

        calls[0] = Multicall3.Call3(
            address(dispatcherProxy),
            false,
            abi.encodeWithSelector(
                Dispatcher.channelOpenTry.selector, ch1, ChannelOrder.NONE, false, connectionHops, ch0, validProof
            )
        );
        calls[1] = Multicall3.Call3(
            address(dispatcherProxy),
            false,
            abi.encodeWithSelector(
                Dispatcher.channelOpenTry.selector, ch1, ChannelOrder.NONE, false, connectionHops, ch0, validProof
            )
        );

        vm.expectEmit(true, true, true, true, address(dispatcherProxy));
        emit ChannelOpenTry(address(mars), "1.0", ChannelOrder.NONE, false, connectionHops, ch0.portId, ch0.channelId);
        vm.expectEmit(true, true, true, true, address(dispatcherProxy));
        emit ChannelOpenTry(address(mars), "1.0", ChannelOrder.NONE, false, connectionHops, ch0.portId, ch0.channelId);
        multicall.aggregate3(calls);
    }
}
