// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Base} from "./Dispatcher.base.t.sol";
import {GasUsingMars} from "./mocks/GasUsingMars.sol";
import {IbcEndpoint, ChannelEnd, IbcUtils, IbcPacket, IBCErrors} from "../contracts/libs/Ibc.sol";
import {TestUtilsTest} from "./TestUtils.t.sol";

contract DispatcherGasGriefing is Base {
    IbcEndpoint src = IbcEndpoint("polyibc.bsc.58b604DB8886656695442374D8E940D814F2eDd4", "channel-99");
    IbcEndpoint dest;
    bytes payload = bytes("msgPayload");
    bytes appAck = abi.encodePacked('{ "account": "account", "reply": "got the message" }');

    GasUsingMars gasUsingMars;
    ChannelEnd ch0 =
        ChannelEnd("polyibc.eth.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-0"), "1.0");
    ChannelEnd ch1 =
        ChannelEnd("polyibc.eth.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-1"), "1.0");

    function setUp() public override {
        (dispatcherProxy, dispatcherImplementation) = TestUtilsTest.deployDispatcherProxyAndImpl(portPrefix);
        gasUsingMars = new GasUsingMars(3_000_000, dispatcherProxy); // Set arbitrarily high gas useage in mars contract
        bytes32 connectionStr = bytes32(0x636f6e6e656374696f6e2d310000000000000000000000000000000000000018); // connection-1
            // in hex
        _storeChannelidToConnectionMapping(ch1.channelId, connectionStr);
        dispatcherProxy.setNewConnection("connection-1", dummyLightClient);
    }

    function test_GasGriefing() public {
        IbcPacket memory packet;
        packet.data = bytes("packet-1");
        packet.timeoutTimestamp = 15_566_401_733_896_437_760;
        packet.dest.channelId = ch1.channelId;
        packet.dest.portId = string(abi.encodePacked(portPrefix, IbcUtils.toHexStr(address(gasUsingMars))));
        packet.src.portId = ch0.portId;
        packet.src.channelId = ch0.channelId;
        packet.sequence = 1;
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.notEnoughGas.selector));
        dispatcherProxy.recvPacket{gas: 2_000_000}(packet, validProof); // Should be enough gas to run out in the
            // callback but still finish execution
    }
}
