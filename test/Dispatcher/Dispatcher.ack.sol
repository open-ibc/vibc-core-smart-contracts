// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.15;

import "../utils/Dispatcher.base.t.sol";
import "../../contracts/examples/Mars.sol";
import "../../contracts/examples/Earth.sol";
import "../../contracts/libs/Ibc.sol";
import {IbcUtils} from "../../contracts/libs/IbcUtils.sol";
import {IBCErrors} from "../../contracts/libs/IbcErrors.sol";
import {IbcReceiver} from "../../contracts/interfaces/IbcReceiver.sol";
import {PacketSenderTestBase} from "./Dispatcher.t.sol";

// Test Chain A receives an acknowledgement packet from Chain B
contract DispatcherAckPacketTestSuite is PacketSenderTestBase {
    function test_success() public {
        for (uint64 index = 0; index < 3; index++) {
            sendPacket();

            vm.expectEmit(true, true, false, true, address(dispatcherProxy));
            emit Acknowledgement(address(mars), channelId, sentPacket.sequence);
            dispatcherProxy.acknowledgement(sentPacket, ackPacket, validProof);
            // confirm dapp recieved the ack
            (bool success, bytes memory data) = mars.ackPackets(sentPacket.sequence - 1);
            AckPacket memory parsed = Ibc.parseAckData(ackPacket);
            assertEq(success, parsed.success);
            assertEq(data, parsed.data);
        }
    }

    // cannot ack packets if packet commitment is missing
    function test_missingPacket() public {
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.packetCommitmentNotFound.selector));
        dispatcherProxy.acknowledgement(genPacket(1), genAckPacket("1"), validProof);

        sendPacket();
        dispatcherProxy.acknowledgement(sentPacket, ackPacket, validProof);

        // packet commitment is removed after ack
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.packetCommitmentNotFound.selector));
        dispatcherProxy.acknowledgement(sentPacket, ackPacket, validProof);
    }

    // cannot recieve ack packets out of order for ordered channel
    function test_outOfOrder() public {
        for (uint64 index = 0; index < 3; index++) {
            sendPacket();
        }
        // 1st ack is ok
        dispatcherProxy.acknowledgement(genPacket(1), genAckPacket("1"), validProof);

        // only 2nd ack is allowed; so the 3rd ack fails
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.unexpectedPacketSequence.selector));

        dispatcherProxy.acknowledgement(genPacket(3), genAckPacket("3"), validProof);
    }

    function test_invalidPort() public {
        Mars earth = new Mars(dispatcherProxy);
        string memory earthPort = string(abi.encodePacked(portPrefix, getHexBytes(address(earth))));
        IbcEndpoint memory earthEnd = IbcEndpoint(earthPort, channelId);

        sendPacket();

        // another valid packet but not the same port
        IbcPacket memory packetEarth = sentPacket;
        packetEarth.src = earthEnd;

        vm.expectRevert(abi.encodeWithSelector(IBCErrors.packetCommitmentNotFound.selector));
        dispatcherProxy.acknowledgement(packetEarth, ackPacket, validProof);
    }

    // ackPacket fails if channel doesn't match
    function test_invalidChannel() public {
        sendPacket();

        IbcEndpoint memory invalidSrc = IbcEndpoint(src.portId, "channel-invalid");
        IbcPacket memory packet = sentPacket;
        packet.src = invalidSrc;

        vm.expectRevert(abi.encodeWithSelector(IBCErrors.channelIdNotFound.selector, packet.src.channelId));

        dispatcherProxy.acknowledgement(packet, ackPacket, validProof);
    }


}
