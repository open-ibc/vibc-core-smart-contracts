// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import '../contracts/Ibc.sol';
import {Dispatcher} from '../contracts/Dispatcher.sol';
import {IbcEventsEmitter} from '../contracts/IbcDispatcher.sol';
import {IbcReceiver} from '../contracts/IbcReceiver.sol';
import '../contracts/UniversalChannelHandler.sol';
import '../contracts/Mars.sol';
import '../contracts/OpConsensusStateManager.sol';
import './Dispatcher.base.t.sol';
import './VirtualChain.sol';

contract UniversalChannelTest is Base {
    function test_channel_settings_ok() public {
        ChannelOrder[2] memory orders = [ChannelOrder.UNORDERED, ChannelOrder.ORDERED];
        bool[2] memory feesEnabled = [true, false];

        for (uint256 i = 0; i < orders.length; i++) {
            for (uint256 j = 0; j < feesEnabled.length; j++) {
                VirtualChain eth1 = new VirtualChain(100, 'polyibc.eth1.');
                VirtualChain eth2 = new VirtualChain(200, 'polyibc.eth2.');
                ChannelSetting memory setting = ChannelSetting(
                    orders[i],
                    '1.0',
                    'polyibc.eth1.7E5F4552091A69125d5DfCb7b8C2659029395Bdf',
                    IbcUtils.toBytes32('channel-0'),
                    feesEnabled[j],
                    validProof
                );
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
        Channel memory channel2Expected = vc1.expectedChannel(
            address(vc1.ucHandler()),
            channelId1,
            vc2.getConnectionHops(),
            setting
        );
        Channel memory channel1Expected = vc2.expectedChannel(
            address(vc2.ucHandler()),
            channelId2,
            vc1.getConnectionHops(),
            setting
        );
        assertEq(abi.encode(channel1), abi.encode(channel1Expected));
        assertEq(abi.encode(channel2), abi.encode(channel2Expected));
    }

    // anyone can open channels, universal or not. But only "official channels" are shown in public IBC dashboards
    function test_channel_ok_by_anyone() public {
        VirtualChain eth1 = new VirtualChain(100, 'polyibc.eth1.');
        VirtualChain eth2 = new VirtualChain(200, 'polyibc.eth2.');
        ChannelSetting memory setting = ChannelSetting(
            ChannelOrder.UNORDERED,
            '1.0',
            'polyibc.eth1.7E5F4552091A69125d5DfCb7b8C2659029395Bdf',
            IbcUtils.toBytes32('channel-0'),
            true,
            validProof
        );

        IbcChannelReceiver ucHandler1 = eth1.ucHandler();
        IbcChannelReceiver ucHandler2 = eth2.ucHandler();
        eth1.assignChanelIds(ucHandler1, ucHandler2, eth2);

        address unauthorized = deriveAddress('unauthorized');
        vm.deal(unauthorized, 100 ether);
        vm.prank(unauthorized);
        eth1.channelOpenInit(ucHandler1, eth2, ucHandler2, setting, true);

        vm.prank(unauthorized);
        eth2.channelOpenTry(ucHandler2, eth1, ucHandler1, setting, true);

        vm.prank(unauthorized);
        eth1.channelOpenAckOrConfirm(ucHandler1, eth2, ucHandler2, setting, false, true);

        vm.prank(unauthorized);
        eth2.channelOpenAckOrConfirm(ucHandler2, eth1, ucHandler1, setting, true, true);
    }
}

struct UcPacket {
    bytes32 channelId;
    address srcPortId;
    bytes appData;
}

contract UniversalChannelPacketTest is Base, IbcMwEventsEmitter {
    VirtualChain eth1;
    VirtualChain eth2;
    VirtualChainData v1;
    VirtualChainData v2;

    bytes appData = bytes('msg-1');
    bytes packetData;
    bytes ackPacketBytes;
    // AckPacket generated by chainB.Earth and received by chainA.Earth
    AckPacket ackPacket;
    // UniversalPacket sent by chainA.Earth and received by chainB.Earth
    UniversalPacket ucPacket;
    // IBC packet received by chainB.Dispatcher
    IbcPacket recvPacket;

    // universal channel packet stored in Earth for test assertion
    UniversalPacket gotUcPacket;
    // ack packet stored in Earth for test assertion
    AckPacket gotAckPacket;
    // channelId stored in Earth for test assertion
    bytes32 gotChannelId;

    function setUp() public virtual override {
        ChannelSetting memory setting = ChannelSetting(
            ChannelOrder.UNORDERED,
            '1.0',
            'polyibc.eth1.7E5F4552091A69125d5DfCb7b8C2659029395Bdf',
            IbcUtils.toBytes32('channel-0'),
            true,
            validProof
        );
        eth1 = new VirtualChain(100, 'polyibc.eth1.');
        eth2 = new VirtualChain(200, 'polyibc.eth2.');
        eth1.finishHandshake(eth1.ucHandler(), eth2, eth2.ucHandler(), setting);
        v1 = eth1.getVirtualChainData();
        v2 = eth2.getVirtualChainData();
    }

    // register middleware stack: dApp -> mw1 -> UniversalChannel MW
    function register_mw1() internal returns (uint256) {
        uint256 mwBitmap = v1.ucHandler.MW_ID() | v1.mw1.MW_ID();
        address[] memory mwAddrs = new address[](1);

        // change Earth's default middleware to mw1, which sits on top of UniversalChannel MW
        vm.startPrank(address(eth1));
        v1.earth.setDefaultMw(address(v1.mw1));
        // register mw1 as the only middleware in the stack
        mwAddrs[0] = address(v1.mw1);
        v1.ucHandler.registerMwStack(mwBitmap, mwAddrs);
        vm.stopPrank();

        vm.startPrank(address(eth2));
        v2.earth.setDefaultMw(address(v2.mw1));
        // register mw1 as the only middleware in the stack
        mwAddrs[0] = address(v2.mw1);
        v2.ucHandler.registerMwStack(mwBitmap, mwAddrs);
        vm.stopPrank();

        return mwBitmap;
    }

    // register middleware stack: dApp -> mw1 -> mw2 -> UniversalChannel MW
    function register_mw1_mw2() internal returns (uint256) {
        address[] memory mwAddrs = new address[](2);
        uint256 mwBitmap = v1.ucHandler.MW_ID() | v1.mw1.MW_ID() | v2.mw2.MW_ID();

        // change Earth's default middleware to mw1, which calls mw2, then UniversalChannel MW
        vm.startPrank(address(eth1));
        v1.earth.setDefaultMw(address(v1.mw1));
        v1.mw1.setDefaultMw(address(v1.mw2));
        // register middleware stack
        mwAddrs[0] = address(v1.mw2);
        mwAddrs[1] = address(v1.mw1);

        v1.ucHandler.registerMwStack(mwBitmap, mwAddrs);
        vm.stopPrank();

        vm.startPrank(address(eth2));
        v2.earth.setDefaultMw(address(v2.mw1));
        v2.mw1.setDefaultMw(address(v2.mw2));
        // register middleware stack
        mwAddrs[0] = address(v2.mw2);
        mwAddrs[1] = address(v2.mw1);
        v2.ucHandler.registerMwStack(mwBitmap, mwAddrs);
        vm.stopPrank();

        return mwBitmap;
    }

    // test address/bytes32 conversion
    function test_address_conversion() public {
        assertEq(address(eth1), IbcUtils.toAddress(IbcUtils.toBytes32(address(eth1))));
        assertEq(address(eth2), IbcUtils.toAddress(IbcUtils.toBytes32(address(eth2))));
    }

    // packet flow: Earth -> UC -> Dispatcher -> (Relayer) -> Dispatcher -> UC -> Earth
    function test_packetFlow_via_universal_channel_ok() public {
        uint256 mwBitmap = v1.ucHandler.MW_ID();
        verifyPacketFlow(5, mwBitmap);
    }

    // packet flow: Earth -> MW1 -> UC -> Dispatcher -> (Relayer) -> Dispatcher -> UC -> MW1 -> Earth
    function test_packetFlow_via_mw1_ok() public {
        uint256 mwBitmap = register_mw1();
        verifyPacketFlow(5, mwBitmap);
    }

    // packet flow: Earth -> MW1 -> MW2 -> UC -> Dispatcher -> (Relayer) -> Dispatcher -> UC -> MW2 -> MW1 -> Earth
    function test_packetFlow_via_mw2_ok() public {
        uint256 mwBitmap = register_mw1_mw2();
        verifyPacketFlow(5, mwBitmap);
    }

    function test_timeout_via_universal_channel_ok() public {
        uint256 mwBitmap = v1.ucHandler.MW_ID();
        verifyTimeoutFlow(5, mwBitmap);
    }

    function test_timeout_via_mw1_ok() public {
        uint256 mwBitmap = register_mw1();
        verifyTimeoutFlow(5, mwBitmap);
    }

    function test_timeout_via_mw2_ok() public {
        uint256 mwBitmap = register_mw1_mw2();
        verifyTimeoutFlow(5, mwBitmap);
    }

    /**
     * Test packet flow from chain A to chain B via UniversalChannel MW and optionally other MW that sits on top of UniversalChannel MW.
     * @param numOfPackets packet sequence starts from 1, and ends at numOfPackets
     * @param mwBitmap bit OR of all MW_IDs of all MWs in the packet flow
     */
    function verifyTimeoutFlow(uint64 numOfPackets, uint256 mwBitmap) internal {
        // universal channelIDs
        bytes32 channelId1 = eth1.channelIds(address(eth1.ucHandler()), address(eth2.ucHandler()));
        bytes32 channelId2 = eth2.channelIds(address(eth2.ucHandler()), address(eth1.ucHandler()));
        GeneralMiddleware[2] memory senderMws = [v1.mw1, v1.mw2];

        for (uint64 packetSeq = 1; packetSeq <= numOfPackets; packetSeq++) {
            uint64 timeout = 1 seconds * 10 ** 9;
            appData = abi.encodePacked('msg-', packetSeq);

            ucPacket = UniversalPacket(
                IbcUtils.toBytes32(address(v1.earth)),
                mwBitmap,
                IbcUtils.toBytes32(address(v2.earth)),
                appData
            );
            packetData = IbcUtils.toUniversalPacketBytes(ucPacket);

            // iterate over sending middleware contracts to verify each MW has witnessed the packet
            for (uint256 i = 0; i < senderMws.length; i++) {
                if (senderMws[i].MW_ID() == (senderMws[i].MW_ID() & mwBitmap)) {
                    vm.expectEmit(true, true, true, true);
                    emit SendMWPacket(
                        channelId1,
                        IbcUtils.toBytes32(address(v1.earth)),
                        IbcUtils.toBytes32(address(v2.earth)),
                        senderMws[i].MW_ID(),
                        appData,
                        timeout,
                        abi.encodePacked(senderMws[i].MW_ID())
                    );
                }
            }
            // Verify event emitted by Dispatcher
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

            //
            // simulate relayer calling dispatcher.timeout on chain A
            //

            // iterate over sending middleware contracts to verify each MW has witnessed the ack
            for (uint256 j = 0; j < senderMws.length; j++) {
                // order is reversed than the sending path. Now chain A receives ack from chain B
                uint256 i = senderMws.length - j - 1;
                if (senderMws[i].MW_ID() == (senderMws[i].MW_ID() & mwBitmap)) {
                    vm.expectEmit(true, true, true, true);
                    emit RecvMWTimeout(
                        channelId1,
                        IbcUtils.toBytes32(address(v1.earth)),
                        IbcUtils.toBytes32(address(v2.earth)),
                        senderMws[i].MW_ID(),
                        appData,
                        abi.encodePacked(senderMws[i].MW_ID())
                    );
                }
            }
            // verify event emitted by Dispatcher
            vm.expectEmit(true, true, true, true);
            emit Timeout(address(v1.ucHandler), channelId1, packetSeq);
            // receive ack on chain A, triggering expected events
            v1.dispatcher.timeout(v1.ucHandler, recvPacket, validProof);

            // verify timeout packet received by Earth on chain A
            (gotChannelId, gotUcPacket) = v1.earth.timeoutPackets(packetSeq - 1);
            assertEq(gotChannelId, channelId1);
            assertEq(abi.encode(gotUcPacket), abi.encode(ucPacket));
        }
    }

    /**
     * Test packet flow from chain A to chain B via UniversalChannel MW and optionally other MW that sits on top of UniversalChannel MW.
     * @param numOfPackets packet sequence starts from 1, and ends at numOfPackets
     * @param mwBitmap bit OR of all MW_IDs of all MWs in the packet flow
     */
    function verifyPacketFlow(uint64 numOfPackets, uint256 mwBitmap) internal {
        // universal channelIDs
        bytes32 channelId1 = eth1.channelIds(address(eth1.ucHandler()), address(eth2.ucHandler()));
        bytes32 channelId2 = eth2.channelIds(address(eth2.ucHandler()), address(eth1.ucHandler()));
        GeneralMiddleware[2] memory senderMws = [v1.mw1, v1.mw2];
        GeneralMiddleware[2] memory recvMws = [v2.mw2, v1.mw1];

        for (uint64 packetSeq = 1; packetSeq <= numOfPackets; packetSeq++) {
            uint64 timeout = 1 days * 10 ** 9 * packetSeq; // change packet settings for each iteration
            appData = abi.encodePacked('msg-', packetSeq);

            ucPacket = UniversalPacket(
                IbcUtils.toBytes32(address(v1.earth)),
                mwBitmap,
                IbcUtils.toBytes32(address(v2.earth)),
                appData
            );
            packetData = IbcUtils.toUniversalPacketBytes(ucPacket);

            // iterate over sending middleware contracts to verify each MW has witnessed the packet
            for (uint256 i = 0; i < senderMws.length; i++) {
                if (senderMws[i].MW_ID() == (senderMws[i].MW_ID() & mwBitmap)) {
                    vm.expectEmit(true, true, true, true);
                    emit SendMWPacket(
                        channelId1,
                        IbcUtils.toBytes32(address(v1.earth)),
                        IbcUtils.toBytes32(address(v2.earth)),
                        senderMws[i].MW_ID(),
                        appData,
                        timeout,
                        abi.encodePacked(senderMws[i].MW_ID())
                    );
                }
            }
            // Verify event emitted by Dispatcher
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

            ackPacket = v2.earth.generateAckPacket(channelId2, address(v1.earth), appData);
            // verify event emitted by Dispatcher
            vm.expectEmit(true, true, true, true);
            emit RecvPacket(address(v2.ucHandler), channelId2, packetSeq);
            // iterate over receiving middleware contracts to verify each MW has witnessed the packet
            for (uint256 i = 0; i < recvMws.length; i++) {
                if (recvMws[i].MW_ID() == (recvMws[i].MW_ID() & mwBitmap)) {
                    vm.expectEmit(true, true, true, true);
                    emit RecvMWPacket(
                        channelId2,
                        IbcUtils.toBytes32(address(v1.earth)),
                        IbcUtils.toBytes32(address(v2.earth)),
                        recvMws[i].MW_ID(),
                        appData,
                        abi.encodePacked(recvMws[i].MW_ID())
                    );
                }
            }
            // verify event emitted by Dispatcher
            vm.expectEmit(true, true, true, true);
            emit WriteAckPacket(address(v2.ucHandler), channelId2, packetSeq, ackPacket);
            v2.dispatcher.recvPacket(v2.ucHandler, recvPacket, validProof);

            // verify packet received by Earth on chain B
            (gotChannelId, gotUcPacket) = v2.earth.recvedPackets(packetSeq - 1);
            assertEq(gotChannelId, channelId2);
            assertEq(abi.encode(gotUcPacket), abi.encode(ucPacket));

            //
            // simulate relayer calling dispatcher.acknowledgePacket on chain A
            //

            // iterate over sending middleware contracts to verify each MW has witnessed the ack
            for (uint256 j = 0; j < senderMws.length; j++) {
                // order is reversed than the sending path. Now chain A receives ack from chain B
                uint256 i = senderMws.length - j - 1;
                if (senderMws[i].MW_ID() == (senderMws[i].MW_ID() & mwBitmap)) {
                    vm.expectEmit(true, true, true, true);
                    emit RecvMWAck(
                        channelId1,
                        IbcUtils.toBytes32(address(v1.earth)),
                        IbcUtils.toBytes32(address(v2.earth)),
                        senderMws[i].MW_ID(),
                        appData,
                        abi.encodePacked(senderMws[i].MW_ID()),
                        ackPacket
                    );
                }
            }
            // verify event emitted by Dispatcher
            vm.expectEmit(true, true, true, true);
            emit Acknowledgement(address(v1.ucHandler), channelId1, packetSeq);
            // receive ack on chain A, triggering expected events
            v1.dispatcher.acknowledgement(v1.ucHandler, recvPacket, ackToBytes(ackPacket), validProof);

            // verify ack packet received by Earth on chain A
            (gotChannelId, gotUcPacket, gotAckPacket) = v1.earth.ackPackets(packetSeq - 1);
            assertEq(gotChannelId, channelId1);
            assertEq(abi.encode(gotUcPacket), abi.encode(ucPacket));
            assertEq(abi.encode(gotAckPacket), abi.encode(ackPacket));
        }
    }
}
