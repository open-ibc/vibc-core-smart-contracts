// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import 'forge-std/Test.sol';
import '../contracts/Ibc.sol';
import {Dispatcher, InitClientMsg, UpgradeClientMsg} from '../contracts/Dispatcher.sol';
import {IbcReceiver} from '../contracts/IbcReceiver.sol';
import '../contracts/IbcVerifier.sol';
import '../contracts/Verifier.sol';
import '../contracts/Mars.sol';

contract Base is Test {
    //
    // channel events
    //

    event OpenIbcChannel(
        address indexed portAddress,
        string version,
        ChannelOrder ordering,
        string[] connectionHops,
        string counterpartyPortId,
        bytes32 counterpartyChannelId
    );

    event ConnectIbcChannel(address indexed portAddress, bytes32 channelId);

    event CloseIbcChannel(address indexed portAddress, bytes32 indexed channelId);

    //
    // packet events
    //

    event SendPacket(
        address indexed sourcePortAddress,
        bytes32 indexed sourceChannelId,
        bytes packet,
        uint64 sequence,
        // timeoutTimestamp is in UNIX nano seconds; packet will be rejected if
        // delivered after this timestamp on the receiving chain.
        // Timeout semantics is compliant to IBC spec and ibc-go implementation
        uint64 timeoutTimestamp,
        PacketFee fee
    );

    event Acknowledgement(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 sequence);

    event Timeout(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 indexed sequence);

    event RecvPacket(address indexed destPortAddress, bytes32 indexed destChannelId, uint64 sequence);

    event WriteAckPacket(
        address indexed writerPortAddress,
        bytes32 indexed writerChannelId,
        uint64 sequence,
        AckPacket ackPacket
    );

    event WriteTimeoutPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence);

    ConsensusState untrustedState =
        ConsensusState(
            80990,
            590199110038530808131927832294665177527506259518072095333098842116767351199,
            7000040,
            1000
        );
    ConsensusState trustedState =
        ConsensusState(
            10934,
            7064503680087416120706887577693908749828198688716609274705703517077803898371,
            7002040,
            1020
        );
    InitClientMsg initClientMsg = InitClientMsg(bytes('Polymer'), untrustedState);

    ZKMintVerifier verifier = new Verifier();
    ZkProof proof =
        ZkProof(
            [
                13449388914393258752883032560537386278857542833249142697090243871249761501123,
                5963894333042515966217276339656894890750758020607775733717462915787234629927
            ],
            [
                [
                    4811559872397934450173412387101758297072581261546553338353504577141293696514,
                    18400634037991283592418145553628778322894277765243742619561207628896194710939
                ],
                [
                    17903685207039300384995795331083569887497623740119108595975170464164316221644,
                    9246628133289276308945311797896077179503891414159382179119544904154776510385
                ]
            ],
            [
                17432552394458345841788798376121543520587716339044416231610790827968220675517,
                15082220514596158133191868403239442750535261032426474092101151620016661078026
            ]
        );
    address payable escrow = payable(vm.addr(uint256(keccak256(abi.encode('escrow')))));

    // Proofs from Polymer chain, to verify packet or channel state on Polymer
    Proof emptyProof;
    Proof invalidProof = Proof(42, bytes('')); // invalid proof with empty proof bytes
    Proof validProof = Proof(42, bytes('valid proof'));

    Dispatcher dispatcher;
}

contract DispatcherCreateClientTest is Test, Base {
    function setUp() public {
        dispatcher = new Dispatcher(verifier, escrow, 'polyibc.eth.');
    }

    function test_success() public {
        dispatcher.createClient(initClientMsg);
    }

    function test_mustByOwner() public {
        vm.prank(vm.addr(0x1));
        vm.expectRevert('Ownable: caller is not the owner');
        dispatcher.createClient(initClientMsg);
    }

    function test_onlyOnce() public {
        dispatcher.createClient(initClientMsg);
        vm.expectRevert('Client already created');
        dispatcher.createClient(initClientMsg);
    }
}

contract DispatcherUpdateClientTest is Test, Base {
    function setUp() public {
        dispatcher = new Dispatcher(verifier, escrow, 'polyibc.eth.');
        dispatcher.createClient(initClientMsg);
    }

    function test_success() public {
        dispatcher.updateClient(UpdateClientMsg(trustedState, proof));
    }

    function test_invalidConsensusState() public {
        vm.expectRevert('UpdateClientMsg proof verification failed');
        dispatcher.updateClient(UpdateClientMsg(untrustedState, proof));
        vm.expectRevert('UpdateClientMsg proof verification failed');
        ConsensusState memory invalidConsensusState;
        dispatcher.updateClient(UpdateClientMsg(invalidConsensusState, proof));
    }
}

contract DispatcherUpgradeClientTest is Test, Base {
    function setUp() public {
        dispatcher = new Dispatcher(verifier, escrow, 'polyibc.eth.');
        dispatcher.createClient(initClientMsg);
        dispatcher.updateClient(UpdateClientMsg(trustedState, proof));
    }

    function test_success() public {
        dispatcher.upgradeClient(UpgradeClientMsg(bytes('upgradeClientState'), trustedState));
    }

    function test_ownerOnly() public {
        vm.prank(vm.addr(0x1));
        vm.expectRevert('Ownable: caller is not the owner');
        dispatcher.upgradeClient(UpgradeClientMsg(bytes('upgradeClientState'), trustedState));
    }
}

struct VersionSet {
    string self;
    string counterparty;
    string expected;
}

contract ChannelTestBase is Test, Base {
    Mars mars = new Mars();
    string[] connectionHops;
    VersionSet version = VersionSet('1.0', '1.0', '1.0');
    ChannelOrder ordering = ChannelOrder.ORDERED;
    CounterParty cp;

    CounterParty cpBsc = CounterParty('polyibc.bsc.9876543210', bytes32('channel-99'), '1.0');

    function setUp() public virtual {
        connectionHops = new string[](2);
        connectionHops[0] = 'connection-1';
        connectionHops[1] = 'connection-2';

        dispatcher = new Dispatcher(verifier, escrow, 'polyibc.eth.');
        dispatcher.createClient(initClientMsg);
        dispatcher.updateClient(UpdateClientMsg(trustedState, proof));

        polymerProof = validProof;
        vm.startPrank(vm.addr(0x1));
    }

    Proof polymerProof;

    function setTestParams(VersionSet memory _version, CounterParty memory _cp, ChannelOrder _ordering) internal {
        version = _version;
        _cp.version = version.counterparty;
        cp = _cp;
        ordering = _ordering;

        // 1st handshake call with unknown counterparty version or channelId
        if (
            cp.channelId == bytes32(0x0) && keccak256(abi.encodePacked(cp.version)) == keccak256(abi.encodePacked(''))
        ) {
            polymerProof = emptyProof;
        } else {
            // all other handshake calls requires a polymer proof
            polymerProof = validProof;
        }
    }

    modifier goodCases() {
        // NOTE: counterparty version is set in VersionSet
        ChannelOrder[2] memory orderings = [ChannelOrder.ORDERED, ChannelOrder.UNORDERED];
        // 1st handshake with unknown counterparty version or channelId
        cp = CounterParty('polyibc.bsc.9876543210', bytes32(0x0), '');
        VersionSet[2] memory versions = [VersionSet('1.0', '', '1.0'), VersionSet('2.0', '', '2.0')];
        for (uint i = 0; i < versions.length; i++) {
            for (uint j = 0; j < orderings.length; j++) {
                setTestParams(versions[i], cp, orderings[j]);
                _;
            }
        }
        // 2nd handshake with known counterparty version and channelId
        cp = CounterParty('polyibc.bsc.9876543210', bytes32('channel-99'), '');
        versions = [VersionSet('', '1.0', '1.0'), VersionSet('', '2.0', '2.0')];
        for (uint i = 0; i < versions.length; i++) {
            for (uint j = 0; j < orderings.length; j++) {
                setTestParams(versions[i], cp, orderings[j]);
                _;
            }
        }
    }
}

contract DispatcherOpenIbcChannelTest is ChannelTestBase {
    function test_success() public goodCases {
        vm.expectEmit(true, true, true, true);
        emit OpenIbcChannel(address(mars), version.expected, ordering, connectionHops, cp.portId, cp.channelId);
        dispatcher.openIbcChannel(IbcReceiver(mars), version.self, ordering, connectionHops, cp, polymerProof);
    }

    modifier unsupportedVersions() {
        // NOTE: counterparty version is set in VersionSet
        ChannelOrder[2] memory orderings = [ChannelOrder.ORDERED, ChannelOrder.UNORDERED];
        // 1st handshake with unknown counterparty version or channelId
        cp = CounterParty('polyibc.bsc.9876543210', bytes32(0x0), '');
        VersionSet[2] memory versions = [
            VersionSet('', '', ''), // cannot be empty
            VersionSet('xxx', '', '')
        ];
        for (uint i = 0; i < versions.length; i++) {
            for (uint j = 0; j < orderings.length; j++) {
                setTestParams(versions[i], cp, orderings[j]);
                _;
            }
        }
        // 2nd handshake with known counterparty version and channelId
        cp = CounterParty('polyibc.bsc.9876543210', bytes32('channel-99'), '');
        versions = [VersionSet('', 'xxx', ''), VersionSet('', '', '')];
        for (uint i = 0; i < versions.length; i++) {
            for (uint j = 0; j < orderings.length; j++) {
                setTestParams(versions[i], cp, orderings[j]);
                _;
            }
        }
    }

    function test_unsupportedVersion() public unsupportedVersions {
        vm.expectRevert(bytes('Unsupported version'));
        dispatcher.openIbcChannel(IbcReceiver(mars), version.self, ordering, connectionHops, cp, polymerProof);
    }

    function test_invalidCounterpartyPortId() public {
        CounterParty[4] memory cps = [
            CounterParty('', bytes32(0x0), ''),
            CounterParty('portX', bytes32(0x0), ''),
            CounterParty('', bytes32('channel-99'), '1.0'),
            CounterParty('portX', bytes32('channel-99'), '1.0')
        ];
        for (uint i = 0; i < cps.length; i++) {
            cp = cps[i];
            vm.expectRevert('Invalid counterpartyPortId');
            dispatcher.openIbcChannel(IbcReceiver(mars), version.self, ordering, connectionHops, cp, polymerProof);
        }
    }

    function test_invalidProof() public {
        vm.expectRevert('Fail to prove channel state');
        dispatcher.openIbcChannel(IbcReceiver(mars), '', ChannelOrder.ORDERED, connectionHops, cpBsc, invalidProof);
    }
}

contract DispatcherConnectIbcChannelTest is ChannelTestBase {
    function test_success() public goodCases {
        dispatcher.openIbcChannel(IbcReceiver(mars), version.self, ordering, connectionHops, cp, polymerProof);
        bytes32 cpVersion = keccak256(abi.encode(cp.version)) == keccak256(abi.encode(bytes('')))
            ? bytes32(abi.encodePacked(version.self))
            : bytes32(abi.encodePacked(cp.version));

        bytes32 channelId = 'channel-1'; // generated by IBC core on Polymer chain
        vm.expectEmit(true, true, true, true);
        emit ConnectIbcChannel(address(mars), channelId);
        dispatcher.connectIbcChannel(
            IbcReceiver(mars),
            channelId,
            connectionHops,
            ordering,
            cpBsc.portId,
            bytes32('channel-99'),
            cpVersion,
            validProof
        );
    }

    function test_unsupportedVersions() public goodCases {
        dispatcher.openIbcChannel(IbcReceiver(mars), version.self, ordering, connectionHops, cp, polymerProof);
        bytes32 cpVersion = 'xxx';

        vm.expectRevert('Unsupported version');
        dispatcher.connectIbcChannel(
            IbcReceiver(mars),
            bytes32('channel-1'),
            connectionHops,
            ordering,
            cp.portId,
            bytes32('channel-99'),
            cpVersion,
            validProof
        );
    }

    function test_invalidProof() public {
        vm.expectRevert('Fail to prove channel state');
        dispatcher.connectIbcChannel(
            IbcReceiver(mars),
            bytes32('channel-1'),
            connectionHops,
            ordering,
            cpBsc.portId,
            cpBsc.channelId,
            bytes32(abi.encodePacked(cpBsc.version)),
            invalidProof
        );
    }
}

contract DispatcherCloseChannelTest is Test, Base {
    Mars mars;
    bytes32 channelId = 'channel-1';

    function setUp() public {
        string[] memory connectionHops = new string[](2);
        connectionHops[0] = 'connection-1';
        connectionHops[1] = 'connection-2';

        dispatcher = new Dispatcher(verifier, escrow, 'polyibc.eth.');
        dispatcher.createClient(initClientMsg);

        vm.startPrank(vm.addr(0x1));
        mars = new Mars();

        dispatcher.updateClient(UpdateClientMsg(trustedState, proof));

        // finish channel handshake as chain A
        CounterParty memory cp = CounterParty('polyibc.bsc.9876543210', bytes32(0x0), '');
        dispatcher.openIbcChannel(IbcReceiver(mars), '1.0', ChannelOrder.ORDERED, connectionHops, cp, emptyProof);
        CounterParty memory cp2 = CounterParty('polyibc.bsc.9876543210', bytes32('channel-99'), '1.0');
        dispatcher.connectIbcChannel(
            IbcReceiver(mars),
            channelId,
            connectionHops,
            ChannelOrder.ORDERED,
            cp2.portId,
            cp2.channelId,
            bytes32(abi.encodePacked(cp2.version)),
            validProof
        );
    }

    function test_closeChannelInit_success() public {
        vm.expectEmit(true, true, true, true);
        emit CloseIbcChannel(address(mars), channelId);
        mars.triggerChannelClose(channelId, IbcDispatcher(dispatcher));
    }

    function test_closeChannelInit_mustOwner() public {
        Mars earth = new Mars();
        vm.expectRevert('Channel not owned by msg.sender');
        earth.triggerChannelClose(channelId, IbcDispatcher(dispatcher));
    }

    function test_closeChannelConfirm_success() public {
        vm.expectEmit(true, true, true, true);
        emit CloseIbcChannel(address(mars), channelId);
        dispatcher.onCloseIbcChannel(address(mars), channelId, validProof);
    }

    function test_closeChannelConfirm_mustOwner() public {
        vm.expectRevert('Channel not owned by portAddress');
        dispatcher.onCloseIbcChannel(address(mars), 'channel-999', validProof);
    }

    function test_closeChannelConfirm_invalidProof() public {
        vm.expectRevert('Fail to prove channel state');
        dispatcher.onCloseIbcChannel(address(mars), channelId, invalidProof);
    }
}

contract DispatcherTest is Test, Base {
    function setUp() public {}
}
