// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../contracts/Ibc.sol";
import {Dispatcher, InitClientMsg, UpgradeClientMsg} from "../contracts/Dispatcher.sol";
import {IbcEventsEmitter} from "../contracts/IbcDispatcher.sol";
import {IbcReceiver} from "../contracts/IbcReceiver.sol";
import "../contracts/IbcVerifier.sol";
import "../contracts/UniversalChannelHandler.sol";
import "../contracts/Verifier.sol";
import "../contracts/Mars.sol";
import "../contracts/OpConsensusStateManager.sol";
import "./Dispatcher.base.t.sol";
import "./VirtualChain.sol";

contract UniversalChannelTest is Base {
    function test_channel_settings_ok() public {
        ChannelOrder[2] memory orders = [ChannelOrder.UNORDERED, ChannelOrder.ORDERED];
        bool[2] memory feesEnabled = [true, false];

        for (uint256 i = 0; i < orders.length; i++) {
            for (uint256 j = 0; j < feesEnabled.length; j++) {
                VirtualChain eth1 = new VirtualChain(100, "polyibc.eth1.");
                VirtualChain eth2 = new VirtualChain(200, "polyibc.eth2.");
                ChannelSetting memory setting = ChannelSetting(orders[i], "1.0", feesEnabled[j], validProof);
                eth1.finishHandshake(eth1.ucHandler(), eth2, eth2.ucHandler(), setting);

                assert_channel(eth1, eth2, setting);
            }
        }
    }

    function assert_channel(VirtualChain vc1, VirtualChain vc2, ChannelSetting memory setting) internal {
        bytes32 channelId1 = vc1.channelIds(address(vc1.ucHandler()), address(vc2.ucHandler()));
        bytes32 channelId2 = vc2.channelIds(address(vc2.ucHandler()), address(vc1.ucHandler()));
        assertEq(vc1.ucHandler().connectedChannels(0), channelId1);
        assertEq(vc2.ucHandler().connectedChannels(0), channelId2);

        Channel memory channel1 = vc1.dispatcher().getChannel(address(vc1.ucHandler()), channelId1);
        Channel memory channel2 = vc2.dispatcher().getChannel(address(vc2.ucHandler()), channelId2);
        Channel memory channel2Expected =
            vc1.expectedChannel(address(vc1.ucHandler()), channelId1, vc2.getConnectionHops(), setting);
        Channel memory channel1Expected =
            vc2.expectedChannel(address(vc2.ucHandler()), channelId2, vc1.getConnectionHops(), setting);
        assertEq(abi.encode(channel1), abi.encode(channel1Expected));
        assertEq(abi.encode(channel2), abi.encode(channel2Expected));
    }

    function test_channel_fail_unauthorized() public {
        VirtualChain eth1 = new VirtualChain(100, "polyibc.eth1.");
        VirtualChain eth2 = new VirtualChain(200, "polyibc.eth2.");
        ChannelSetting memory setting = ChannelSetting(ChannelOrder.UNORDERED, "1.0", true, validProof);

        IbcChannelHandler ucHandler1 = eth1.ucHandler();
        IbcChannelHandler ucHandler2 = eth2.ucHandler();
        eth1.assignChanelIds(ucHandler1, ucHandler2, eth2);

        address unauthorized = deriveAddress("unauthorized");
        vm.deal(unauthorized, 100 ether);
        vm.prank(unauthorized);
        expectRevertNonOwner();
        eth1.channelOpenInit(ucHandler1, eth2, ucHandler2, setting, false);
        vm.prank(address(eth1));
        eth1.channelOpenInit(ucHandler1, eth2, ucHandler2, setting, true);

        vm.prank(unauthorized);
        expectRevertNonOwner();
        eth2.channelOpenTry(ucHandler2, eth1, ucHandler1, setting, false);
        vm.prank(address(eth2));
        eth2.channelOpenTry(ucHandler2, eth1, ucHandler1, setting, true);

        vm.prank(unauthorized);
        expectRevertNonOwner();
        eth1.channelOpenAckOrConfirm(ucHandler1, eth2, ucHandler2, setting, false);
        vm.prank(address(eth1));
        eth1.channelOpenAckOrConfirm(ucHandler1, eth2, ucHandler2, setting, true);

        vm.prank(unauthorized);
        expectRevertNonOwner();
        eth2.channelOpenAckOrConfirm(ucHandler2, eth1, ucHandler1, setting, false);
        vm.prank(address(eth2));
        eth2.channelOpenAckOrConfirm(ucHandler2, eth1, ucHandler1, setting, true);
    }
}

contract UniversalChannelPacketTest is Base {
    VirtualChain eth1;
    VirtualChain eth2;
    ChannelSetting setting = ChannelSetting(ChannelOrder.UNORDERED, "1.0", true, validProof);
    Earth earth1;
    Earth earth2;
    UniversalChannelHandler ucHandler1;
    UniversalChannelHandler ucHandler2;

    bytes appData = bytes("msg-1");
    bytes packetData;
    bytes ackPacket;
    UniversalPacketData ucData;
    // IBC packet received by dispatcher
    IbcPacket recvPacket;
    // actual universal channel packet received by dApp
    IbcPacket ucPacket;

    function setUp() public virtual {
        eth1 = new VirtualChain(100, "polyibc.eth1.");
        eth2 = new VirtualChain(200, "polyibc.eth2.");
        eth1.finishHandshake(eth1.ucHandler(), eth2, eth2.ucHandler(), setting);
        earth1 = eth1.earth();
        earth2 = eth2.earth();
        ucHandler1 = eth1.ucHandler();
        ucHandler2 = eth2.ucHandler();
    }

    function test_sendPacket_ok() public {
        bytes32 channelId1 = eth1.channelIds(address(eth1.ucHandler()), address(eth2.ucHandler()));
        bytes32 channelId2 = eth2.channelIds(address(eth2.ucHandler()), address(eth1.ucHandler()));
        string memory earth1PortId = eth1.portIds(address(eth1.earth()));
        string memory earth2PortId = eth2.portIds(address(eth2.earth()));

        // send packet from earth1 to earth2
        for (uint64 packetSeq = 1; packetSeq <= 5; packetSeq++) {
            uint64 factor = packetSeq; // change packet settings for each iteration
            uint64 timeout = 1 days * 10 ** 9 * factor;
            appData = abi.encodePacked("msg-", packetSeq);

            ucData = UniversalPacketData(
                Ibc.addressToPortId(eth1.dispatcher().portPrefix(), address(earth1)), earth2PortId, appData
            );
            packetData = Ibc.toUniversalPacketDataBytes(ucData);
            vm.expectEmit(true, true, true, true);
            emit SendPacket(address(ucHandler1), channelId1, packetData, packetSeq, timeout);
            earth1.greet(earth2PortId, channelId1, appData, timeout);

            // simulate relayer calling dispatcher.recvPacket on chain B
            // recvPacket is an IBC packet
            recvPacket = IbcPacket(
                IbcEndpoint(eth1.portIds(address(ucHandler1)), channelId1),
                IbcEndpoint(eth2.portIds(address(ucHandler2)), channelId2),
                packetSeq,
                packetData,
                Height(0, 0),
                timeout
            );
            ucPacket = IbcPacket(
                IbcEndpoint(earth1PortId, channelId1),
                IbcEndpoint(earth2PortId, channelId2),
                packetSeq,
                appData,
                Height(0, 0),
                timeout
            );
            ackPacket = abi.encodePacked("ack-", appData);
            vm.expectEmit(true, true, true, true);
            emit RecvPacket(address(earth2), channelId2, packetSeq);
            emit WriteAckPacket(
                address(earth2), channelId2, packetSeq, AckPacket(true, abi.encodePacked("ack-", appData))
            );
            eth2.dispatcher().recvPacket(earth2, recvPacket, setting.proof);

            IbcPacket memory dappPacket = getIbcPacket(earth2, packetSeq - 1);
            assertEq(abi.encode(dappPacket), abi.encode(ucPacket));
        }
    }

    // getIbcPacket returns an IBC packet by index
    function getIbcPacket(Earth earth, uint256 index) internal view returns (IbcPacket memory) {
        (
            IbcEndpoint memory _src,
            IbcEndpoint memory _dest,
            uint64 _seq,
            bytes memory _data,
            Height memory _height,
            uint64 _timeoutTimestamp
        ) = earth.recvedPackets(index);
        return IbcPacket(_src, _dest, _seq, _data, _height, _timeoutTimestamp);
    }
}
