// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "forge-std/Script.sol";
import "../contracts/core/Dispatcher.sol";
import "../contracts/examples/Mars.sol";
import "../contracts/libs/Ibc.sol";
import "forge-std/console.sol";


contract BatchedEvents {
    Dispatcher public dispatcher;
    Ics23Proof validProof;
    string[] connectionHops;
    uint64 UINT64_MAX = 18_446_744_073_709_551_615;
    Height ZERO_HEIGHT = Height(0, 0);
    uint64 maxTimeout = UINT64_MAX;
    string srcChannelId;
    string destChannelId;
    string srcPortId;
    string destPortId;


    constructor(Dispatcher _dispatcher, string memory _srcChannelId, string memory _destChannelId) {
        dispatcher = _dispatcher;
        srcChannelId = _srcChannelId;
        destChannelId = _destChannelId;
        validProof.height = 10;

        srcPortId = IbcUtils.addressToPortId("polyibc.optimism-sim.", address(this));
        destPortId = IbcUtils.addressToPortId("polyibc.base-sim.", address(this));

        connectionHops = new string[](2);
        connectionHops[0] = "connection-12";
        connectionHops[1] = "connection-15";

    }

    function channelHandshake() public {
        openIbcChannel();
        connectIbcChannel();
    }

    function openIbcChannel() public {
        dispatcher.openIbcChannel(
            IbcChannelReceiver(address(this)),
            CounterParty(
                srcPortId,
                0x0000000000000000000000000000000000000000000000000000000000000000,
                "1.0"
            ),
            ChannelOrder.UNORDERED,
            false,
            connectionHops,
            CounterParty(
                destPortId,
                0x0000000000000000000000000000000000000000000000000000000000000000,
                ""
            ),
            validProof
        );
    }

    function onCloseIbcChannel(bytes32 channelId, string calldata, bytes32) external virtual {}

    function triggerChannelClose(bytes32 channelId) external {
        dispatcher.closeIbcChannel(channelId);
    }

    function onOpenIbcChannel(
        string calldata version,
        ChannelOrder,
        bool,
        string[] calldata,
        CounterParty calldata counterparty
    ) external view virtual returns (string memory selectedVersion) {
        selectedVersion =
            keccak256(abi.encodePacked(version)) == keccak256(abi.encodePacked("")) ? counterparty.version : version;

        return selectedVersion;
    }

    function connectIbcChannel() public {
        dispatcher.connectIbcChannel(IbcChannelReceiver(address(this)),
            CounterParty(
                srcPortId,
                IbcUtils.toBytes32(srcChannelId),
                "1.0"
            ),
            connectionHops,
            ChannelOrder.UNORDERED,
            false,
            false,
            CounterParty(
                destPortId,
                IbcUtils.toBytes32(destChannelId),
                "1.0"
            ),
            validProof
        );
    }

    function onConnectIbcChannel(bytes32 channelId, bytes32, string calldata counterpartyVersion) external virtual {}

    function sendPacket() public {
        dispatcher.sendPacket(IbcUtils.toBytes32(srcChannelId), bytes("payload"), maxTimeout);
    }

    function timeoutPacket() public {
        IbcEndpoint memory src = IbcEndpoint(srcPortId, IbcUtils.toBytes32(srcChannelId));
        IbcEndpoint memory dest = IbcEndpoint(destPortId, IbcUtils.toBytes32(destChannelId));

        IbcPacket memory sentPacket = IbcPacket(src, dest, 1, bytes("payload"), ZERO_HEIGHT, maxTimeout);
        dispatcher.timeout(IbcPacketReceiver(address(this)), sentPacket, validProof);
    }

    function packetEventsAndCloseChannel() public {
        sendPacket();
        timeoutPacket();
        closeChannel();
    }

    function closeChannel() public {
        this.triggerChannelClose(IbcUtils.toBytes32(srcChannelId));
    }

    function onRecvPacket(IbcPacket memory packet) external virtual returns (AckPacket memory ackPacket)
    {
        // solhint-disable-next-line quotes
        return AckPacket(true, abi.encodePacked('{ "account": "account", "reply": "got the message" }'));
    }

    function onAcknowledgementPacket(IbcPacket calldata, AckPacket calldata ack) external virtual {
    }

    function onTimeoutPacket(IbcPacket calldata packet) external virtual {
    }
}

contract GenerateEvents is Script {
    Dispatcher dispatcher;
    string srcChannelId;
    string destChannelId;


    constructor() {
        dispatcher = Dispatcher(0xC6c8a05aa18327DCA7c9f32dCb7da64aBfd9592F);
        srcChannelId = string(abi.encodePacked("channel-", vm.toString(block.number)));
        destChannelId = string(abi.encodePacked("channel-", vm.toString(block.number + 1)));
        console.log("srcChannelId: ", srcChannelId);
        console.log("destChannelId: ", destChannelId);
    }

    function run() public {
        vm.startBroadcast(msg.sender);

        BatchedEvents be = new BatchedEvents(dispatcher, srcChannelId, destChannelId);
        be.channelHandshake();
        be.packetEventsAndCloseChannel();
        vm.stopBroadcast();
    }
}