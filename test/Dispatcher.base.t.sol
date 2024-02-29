// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import 'forge-std/Test.sol';
import {ProofBase} from './Proof.base.t.sol';
import '../contracts/libs/Ibc.sol';
import {Dispatcher} from '../contracts/core/Dispatcher.sol';
import {IbcEventsEmitter} from '../contracts/interfaces/IbcDispatcher.sol';
import {IbcChannelReceiver} from '../contracts/interfaces/IbcReceiver.sol';
import '../contracts/examples/Mars.sol';
import '../contracts/core/OpLightClient.sol';
import '../contracts/utils/DummyLightClient.sol';
import '../contracts/core/OpProofVerifier.sol';

struct LocalEnd {
    IbcChannelReceiver receiver;
    string portId;
    // channelId is only used in connectIbcChannel
    bytes32 channelId;
    string[] connectionHops;
    string versionCall;
    string versionExpected;
}

struct ChannelHandshakeSetting {
    ChannelOrder ordering;
    bool feeEnabled;
    bool localInitiate;
    Ics23Proof proof;
}

// Base contract for testing Dispatcher
contract Base is IbcEventsEmitter, ProofBase {
    uint64 UINT64_MAX = 18_446_744_073_709_551_615;

    Height ZERO_HEIGHT = Height(0, 0);
    uint64 maxTimeout = UINT64_MAX;

    LightClient opLightClient =
        new OptimisticLightClient(1800, opProofVerifier, l1BlockProvider);

    LightClient dummyConsStateManager = new DummyLightClient();

    Dispatcher dispatcher;
    string portPrefix = "polyibc.eth.";
    string[] connectionHops = ["connection-1", "connection-2"];

    // ⬇️ Utility functions for testing

    // getHexStr converts an address to a hex string without the leading 0x
    function getHexBytes(address addr) internal pure returns (bytes memory) {
        bytes memory addrWithPrefix = abi.encodePacked(vm.toString(addr));
        bytes memory addrWithoutPrefix = new bytes(addrWithPrefix.length - 2);
        for (uint256 i = 0; i < addrWithoutPrefix.length; i++) {
            addrWithoutPrefix[i] = addrWithPrefix[i + 2];
        }
        return addrWithoutPrefix;
    }

    // deriveAddress derives an address from a given string deterministically for testing
    function deriveAddress(string memory str) internal pure returns (address) {
        return vm.addr(uint256(keccak256(abi.encodePacked(str))));
    }

    // ⬇️ IBC functions for testing

    /**
     * @dev Step-1 of the 4-step handshake to open an IBC channel.
     * @param le Local end settings for the channel.
     * @param re Remote end settings for the channel.
     * @param s Channel handshake settings.
     * @param expPass Expected pass status of the operation.
     * If expPass is false, `vm.expectRevert` should be called before this function.
     */
    function channelOpenInit(LocalEnd memory le, CounterParty memory re, ChannelHandshakeSetting memory s, bool expPass)
        public
    {
        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit ChannelOpenInit(
                address(le.receiver), le.versionExpected, s.ordering, s.feeEnabled, le.connectionHops, re.portId
            );
        }
        dispatcher.channelOpenInit(le.receiver, le.versionCall, s.ordering, s.feeEnabled, le.connectionHops, re.portId);
    }

    /**
     * @dev Step-2 of the 4-step handshake to open an IBC channel.
     * @param le Local end settings for the channel.
     * @param re Remote end settings for the channel.
     * @param s Channel handshake settings.
     * @param expPass Expected pass status of the operation.
     * If expPass is false, `vm.expectRevert` should be called before this function.
     */
    function channelOpenTry(LocalEnd memory le, CounterParty memory re, ChannelHandshakeSetting memory s, bool expPass)
        public
    {
        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit ChannelOpenTry(
                address(le.receiver),
                le.versionExpected,
                s.ordering,
                s.feeEnabled,
                le.connectionHops,
                re.portId,
                re.channelId
            );
        }
        CounterParty memory cp = CounterParty(re.portId, re.channelId, re.version);
        dispatcher.channelOpenTry(
            le.receiver,
            CounterParty(le.portId, le.channelId, le.versionCall),
            s.ordering,
            s.feeEnabled,
            le.connectionHops,
            cp,
            s.proof
        );
    }

    /**
     * @dev Step-3 of the 4-step handshake to open an IBC channel.
     * @param le Local end settings for the channel.
     * @param re Remote end settings for the channel.
     * @param s Channel handshake settings.
     * @param expPass Expected pass status of the operation.
     * If expPass is false, `vm.expectRevert` should be called before this function.
     */
    function channelOpenAck(LocalEnd memory le, CounterParty memory re, ChannelHandshakeSetting memory s, bool expPass)
        public
    {
        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit ChannelOpenAck(address(le.receiver), le.channelId);
        }
        dispatcher.channelOpenAck(
            le.receiver,
            CounterParty(le.portId, le.channelId, le.versionCall),
            le.connectionHops,
            s.ordering,
            s.feeEnabled,
            re,
            s.proof
        );
    }

    /**
     * @dev Step-4 of the 4-step handshake to open an IBC channel.
     * @param le Local end settings for the channel.
     * @param re Remote end settings for the channel.
     * @param s Channel handshake settings.
     * @param expPass Expected pass status of the operation.
     * If expPass is false, `vm.expectRevert` should be called before this function.
     */
    function channelOpenConfirm(
        LocalEnd memory le,
        CounterParty memory re,
        ChannelHandshakeSetting memory s,
        bool expPass
    ) public {
        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit ChannelOpenConfirm(address(le.receiver), le.channelId);
        }
        dispatcher.channelOpenConfirm(
            le.receiver,
            CounterParty(le.portId, le.channelId, le.versionCall),
            le.connectionHops,
            s.ordering,
            s.feeEnabled,
            re,
            s.proof
        );
    }

    // A helper function to expect revert with message when contract is called with non-owner.
    // Error msg is defined by OpenZeppelin Ownable contract.
    function expectRevertNonOwner() internal {
        vm.expectRevert("Ownable: caller is not the owner");
    }

    function ackToBytes(AckPacket memory ack) public pure returns (bytes memory) {
        return ack.success
            ? abi.encodePacked('{"result":"', Base64.encode(ack.data), '"}')
            : abi.encodePacked('{"error":"', ack.data, '"}');
    }
}
