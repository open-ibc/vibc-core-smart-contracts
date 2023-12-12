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

    // anyone can open channels, universal or not. But only "official channels" are shown in public IBC dashboards
    function test_channel_ok_by_anyone() public {
        VirtualChain eth1 = new VirtualChain(100, "polyibc.eth1.");
        VirtualChain eth2 = new VirtualChain(200, "polyibc.eth2.");
        ChannelSetting memory setting = ChannelSetting(ChannelOrder.UNORDERED, "1.0", true, validProof);

        IbcChannelReceiver ucHandler1 = eth1.ucHandler();
        IbcChannelReceiver ucHandler2 = eth2.ucHandler();
        eth1.assignChanelIds(ucHandler1, ucHandler2, eth2);

        address unauthorized = deriveAddress("unauthorized");
        vm.deal(unauthorized, 100 ether);
        vm.prank(unauthorized);
        eth1.channelOpenInit(ucHandler1, eth2, ucHandler2, setting, true);

        vm.prank(unauthorized);
        eth2.channelOpenTry(ucHandler2, eth1, ucHandler1, setting, true);

        vm.prank(unauthorized);
        eth1.channelOpenAckOrConfirm(ucHandler1, eth2, ucHandler2, setting, true);

        vm.prank(unauthorized);
        eth2.channelOpenAckOrConfirm(ucHandler2, eth1, ucHandler1, setting, true);
    }
}

struct UcPacket {
    bytes32 channelId;
    address srcPortId;
    bytes appData;
}

contract UniversalChannelPacketTest is Base {
    VirtualChain eth1;
    VirtualChain eth2;
    ChannelSetting setting = ChannelSetting(ChannelOrder.UNORDERED, "1.0", true, validProof);
    VirtualChainData v1;
    VirtualChainData v2;

    bytes appData = bytes("msg-1");
    bytes packetData;
    bytes ackPacketBytes;
    AckPacket ackPacket;
    UniversalPacket ucData;
    // IBC packet received by dispatcher
    IbcPacket recvPacket;
    // actual universal channel packet received by dApp
    IbcPacket ucPacket;

    function setUp() public virtual {
        eth1 = new VirtualChain(100, "polyibc.eth1.");
        eth2 = new VirtualChain(200, "polyibc.eth2.");
        eth1.finishHandshake(eth1.ucHandler(), eth2, eth2.ucHandler(), setting);
        v1 = eth1.getVirtualChainData();
        v2 = eth2.getVirtualChainData();
    }

    function test_sendPacket_ok() public {
        // universal channelIDs
        bytes32 channelId1 = eth1.channelIds(address(eth1.ucHandler()), address(eth2.ucHandler()));
        bytes32 channelId2 = eth2.channelIds(address(eth2.ucHandler()), address(eth1.ucHandler()));

        // send packet from earth1 to earth2
        for (uint64 packetSeq = 1; packetSeq <= 5; packetSeq++) {
            uint64 factor = packetSeq; // change packet settings for each iteration
            uint64 timeout = 1 days * 10 ** 9 * factor;
            appData = abi.encodePacked("msg-", packetSeq);

            ucData = UniversalPacket(address(v1.earth), v1.ucHandler.MW_ID(), address(v2.earth), appData);
            packetData = Ibc.toUniversalPacketBytes(ucData);
            vm.expectEmit(true, true, true, true);
            emit SendPacket(address(v1.ucHandler), channelId1, packetData, packetSeq, timeout);
            v1.earth.greet(address(v2.earth), channelId1, appData, timeout);

            // simulate relayer calling dispatcher.recvPacket on chain B
            // recvPacket is an IBC packet
            recvPacket = IbcPacket(
                IbcEndpoint(eth1.portIds(address(v1.ucHandler)), channelId1),
                IbcEndpoint(eth2.portIds(address(v2.ucHandler)), channelId2),
                packetSeq,
                packetData,
                Height(0, 0),
                timeout
            );
            ucPacket = IbcPacket(
                IbcEndpoint(eth1.portIds(address(v1.ucHandler)), channelId1),
                IbcEndpoint(eth2.portIds(address(v2.ucHandler)), channelId2),
                packetSeq,
                appData,
                Height(0, 0),
                timeout
            );
            ackPacket = v2.earth.generateAckPacket(channelId2, address(v1.earth), appData);
            vm.expectEmit(true, true, true, true);
            emit RecvPacket(address(v2.ucHandler), channelId2, packetSeq);
            emit WriteAckPacket(address(v2.ucHandler), channelId2, packetSeq, ackPacket);
            v2.dispatcher.recvPacket(v2.ucHandler, recvPacket, setting.proof);

            assertDappUcPacket(v2.earth, packetSeq - 1, UcPacket(channelId2, address(v1.earth), appData));

            //
            // simulate relayer calling dispatcher.acknowledgePacket on chain A
            //
            vm.expectEmit(true, true, true, true);
            emit Acknowledgement(address(v1.ucHandler), channelId1, packetSeq);
            v1.dispatcher.acknowledgement(v1.ucHandler, recvPacket, ackPacket, setting.proof);

            (bool success, bytes memory data) = v1.earth.ackPackets(packetSeq - 1);
            assertEq(success, true);
            assertEq(data, ackPacket.data);
        }
    }

    // TODO: test timeout
    function test_timeoutPacket_ok() public {}

    function assertDappUcPacket(Earth earth, uint256 index, UcPacket memory expectedPacket) internal {
        (bytes32 channelId, address srcPortId, bytes memory _appData) = earth.recvedPackets(index);
        assertEq(channelId, expectedPacket.channelId);
        assertEq(srcPortId, expectedPacket.srcPortId);
        assertEq(_appData, expectedPacket.appData);
    }
}
