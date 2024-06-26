// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../utils/Dispatcher.base.t.sol";
import "../../contracts/examples/Mars.sol";
import "../../contracts/examples/Earth.sol";
import "../../contracts/libs/Ibc.sol";
import {IbcUtils} from "../../contracts/libs/IbcUtils.sol";
import {IBCErrors} from "../../contracts/libs/IbcErrors.sol";

import {IbcReceiver} from "../../contracts/interfaces/IbcReceiver.sol";
import {PacketSenderTestBase} from "./Dispatcher.t.sol";

// Test Chain A receives a timeout packet from Chain B
contract DispatcherTimeoutPacketTestSuite is PacketSenderTestBase {
    // preconditions for timeout packet
    // - packet commitment exists
    // - packet timeout is verified by Polymer client
    function test_success() public {
        vm.skip(true);
        for (uint64 index = 0; index < 3; index++) {
            sendPacket();

            vm.expectEmit(true, true, true, true, address(dispatcherProxy));
            emit Timeout(address(mars), channelId, sentPacket.sequence);
            dispatcherProxy.timeout(sentPacket, validProof);
        }
    }

    function test_timeout_dapp_revert() public {
        vm.skip(true);
        sentPacket = IbcPacket(srcRevertingMars, dest, 1, payload, ZERO_HEIGHT, maxTimeout);
        revertingBytesMars.greet(payloadStr, channelId, maxTimeout);
        nextSendSeq += 1;
        vm.expectEmit(true, true, true, true, address(dispatcherProxy));
        emit TimeoutError(
            address(revertingBytesMars), abi.encodeWithSelector(RevertingBytesMars.OnTimeoutPacket.selector)
        );
        dispatcherProxy.timeout(sentPacket, validProof);
    }

    // cannot timeout packets if packet commitment is missing
    function test_missingPacket() public {
        vm.skip(true);
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.packetCommitmentNotFound.selector));
        dispatcherProxy.timeout(genPacket(1), validProof);

        sendPacket();
        dispatcherProxy.timeout(sentPacket, validProof);

        // packet commitment is removed after timeout
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.packetCommitmentNotFound.selector));
        dispatcherProxy.timeout(sentPacket, validProof);
    }

    // cannot timeout packets if original packet port doesn't match current port
    function test_invalidPort() public {
        vm.skip(true);
        Mars earth = new Mars(dispatcherProxy);
        string memory earthPort = string(abi.encodePacked(portPrefix, getHexBytes(address(earth))));
        IbcEndpoint memory earthEnd = IbcEndpoint(earthPort, channelId);

        sendPacket();

        // another valid packet but not the same port
        IbcPacket memory packetEarth = sentPacket;
        packetEarth.src = earthEnd;

        vm.expectRevert(IBCErrors.packetCommitmentNotFound.selector);
        dispatcherProxy.timeout(packetEarth, validProof);
    }

    // cannot timeout packetsfails if channel doesn't match
    function test_invalidChannel() public {
        vm.skip(true);
        sendPacket();

        IbcEndpoint memory invalidSrc = IbcEndpoint(src.portId, "channel-invalid");
        IbcPacket memory packet = sentPacket;
        packet.src = invalidSrc;

        bytes32 connectionStr = bytes32(0x636f6e6e656374696f6e2d310000000000000000000000000000000000000018); //
        // Connection-1
        _storeChannelidToConnectionMapping(packet.src.channelId, connectionStr);

        vm.expectRevert(abi.encodeWithSelector(IBCErrors.packetCommitmentNotFound.selector));
        /* vm.expectRevert('Packet commitment not found'); */
        dispatcherProxy.timeout(packet, validProof);
    }

    // cannot timeout packets if proof from Polymer is invalid
    function test_invalidProof() public {
        vm.skip(true);
        sendPacket();
        vm.expectRevert(DummyLightClient.InvalidDummyNonMembershipProof.selector);
        dispatcherProxy.timeout(sentPacket, invalidProof);
    }

    function test_invalidSendPacket() public {
        sentPacket = genPacket(nextSendSeq);
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.invalidPacket.selector));
        mars.greet(payloadStr, channelId, uint64(block.timestamp) - 1);
    }

    function test_writeTimeoutPacket() public {
        IbcEndpoint memory packetSrc = IbcEndpoint("polyibc.bsc.71C95911E9a5D330f4D621842EC243EE1343292e", "channel-99");
        IbcEndpoint memory packetDest = IbcEndpoint(marsPort, channelId);
        uint64 timeoutTimestamp = (uint64(block.timestamp));
        IbcPacket memory packet = IbcPacket(packetSrc, packetDest, 1, payload, ZERO_HEIGHT, timeoutTimestamp);

        vm.warp(block.timestamp + 100);
        vm.expectEmit(true, true, true, true, address(dispatcherProxy));
        emit WriteTimeoutPacket(address(mars), packet.dest.channelId, 1, ZERO_HEIGHT, timeoutTimestamp);
        dispatcherProxy.writeTimeoutPacket(packet, validProof);
    }

    function test_writeTimeoutPacket_invalid_PacketReceiptAlreadyExists() public {
        // This test ensures that we can't call writeTimeout for a packet that has already been received on the dest
        // chain.

        // First, we receive a packet from another chain, with the Mars contract on this chain as the destination
        IbcEndpoint memory packetSrc = IbcEndpoint("polyibc.bsc.71C95911E9a5D330f4D621842EC243EE1343292e", "channel-99");
        IbcEndpoint memory packetDest = IbcEndpoint(marsPort, channelId);

        IbcPacket memory packet = IbcPacket(packetSrc, packetDest, 1, payload, ZERO_HEIGHT, uint64(block.timestamp) + 1);
        dispatcherProxy.recvPacket(packet, validProof);

        vm.warp(block.timestamp + 100);
        // Then we ensure that writeTimeoutPacket can't be called once the packet has been received successfully, even
        // for a block timestap that has passed the packets original timeouttimestmap.
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.packetReceiptAlreadyExists.selector));
        dispatcherProxy.writeTimeoutPacket(packet, validProof);
    }
}
