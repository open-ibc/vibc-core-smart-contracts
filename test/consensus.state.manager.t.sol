// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../contracts/Ibc.sol";
import {Dispatcher, InitClientMsg, UpgradeClientMsg} from "../contracts/Dispatcher.sol";
import {IbcEventsEmitter} from "../contracts/IbcDispatcher.sol";
import {Escrow} from "../contracts/Escrow.sol";
import {IbcReceiver} from "../contracts/IbcReceiver.sol";
import "../contracts/IbcVerifier.sol";
import "../contracts/Verifier.sol";
import "../contracts/Mars.sol";
import "../contracts/OpConsensusStateManager.sol";
import "./Dispatcher.base.t.sol";

contract StateManagerTest is Base {
    LocalEnd local;
    RemoteEnd remote;
    ChannelHandshakeSetting setting;

    Mars mars;

    function setUp() public {
        dispatcher = new Dispatcher(verifier, escrow, portPrefix, dummyConsStateManager);
        mars = new Mars(dispatcher);
        local = LocalEnd(mars, "channel-1", new string[](2), "1.0", "1.0");
        local.connectionHops[0] = "connection-1";
        local.connectionHops[1] = "connection-2";

        remote = RemoteEnd("eth2.7E5F4552091A69125d5DfCb7b8C2659029395Bdf", "channel-2", "1.0");
        setting = ChannelHandshakeSetting(ChannelOrder.UNORDERED, false, false, validProof);
    }

    function test_ok() public {
        openChannel(local, remote, setting, true);
        connectionChannel(local, remote, setting, true);
    }

    function test_fail() public {
        LocalEnd memory le = local;
        RemoteEnd memory re = remote;
        le.versionCall = "";
        ChannelHandshakeSetting memory s = setting;
        s.localInitiate = true;
        vm.expectRevert(bytes("Unsupported version"));
        openChannel(le, re, s, false);
        // connectionChannel(local, remote, setting, false);
    }
}
