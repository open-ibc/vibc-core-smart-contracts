// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../../contracts/libs/Ibc.sol";
import {Dispatcher} from "../../contracts/core/Dispatcher.sol";
import {IbcEventsEmitter} from "../../contracts/interfaces/IbcDispatcher.sol";
import {IbcReceiver} from "../../contracts/interfaces/IbcReceiver.sol";
import "../../contracts/examples/Mars.sol";
import "../../contracts/core/OpLightClient.sol";
import "../Dispatcher.base.t.sol";
import {DispatcherV2} from "./upgrades/DispatcherV2.sol";
import {DispatcherV2Initializable} from "./upgrades/DispatcherV2Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import {TestUtils} from "../TestUtils.sol";
import {DummyLightClient} from "../../contracts/utils/DummyLightClient.sol";

contract DispatcherUUPSAccessControl is Base {
    string public portPrefix2 = "IIpolyibc.eth.";
    LightClient lightClient2 = new DummyLightClient();
    address public notOwner = vm.addr(1);
    DispatcherV2 impl2;
    DispatcherV2Initializable impl3;

    function setUp() public override {
        (dispatcher, impl) = TestUtils.deployDispatcherProxyAndImpl(portPrefix, dummyConsStateManager);
        impl2 = new DispatcherV2();
        impl3 = new DispatcherV2Initializable();
    }

    function test_Dispatcher_Allows_Upgrade() public {
        assertEq(address(impl), TestUtils.getProxyImplementation(address(dispatcher), vm));
        UUPSUpgradeable(address(dispatcher)).upgradeTo(address(impl2));
        assertEq(address(impl2), TestUtils.getProxyImplementation(address(dispatcher), vm));
        assertEq(dispatcher.portPrefix(), portPrefix);
        assertEq(
            address(uint160(uint256(vm.load(address(dispatcher), bytes32(uint256(110)))))),
            address(dummyConsStateManager)
        );
    }

    function test_Dispatcher_Allows_Upgrade_To_And_Call() public {
        assertEq(address(impl), TestUtils.getProxyImplementation(address(dispatcher), vm));
        bytes memory initData = abi.encodeWithSignature("initialize(string,address)", portPrefix2, lightClient2);
        UUPSUpgradeable(address(dispatcher)).upgradeToAndCall(address(impl3), initData);
        assertEq(address(impl3), TestUtils.getProxyImplementation(address(dispatcher), vm));
        assertEq(dispatcher.portPrefix(), portPrefix2);
        assertEq(address(uint160(uint256(vm.load(address(dispatcher), bytes32(uint256(110)))))), address(lightClient2));
    }

    function test_Dispatcher_Prevents_Non_Owner_Updgrade() public {
        vm.startPrank(notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        UUPSUpgradeable(address(dispatcher)).upgradeTo(address(impl2));
    }

    function test_Dispatcher_Prevents_Non_Owner_UpdgradeToAndCall() public {
        vm.expectRevert("Function must be called through delegatecall");
        UUPSUpgradeable(address(impl)).upgradeTo(address(impl2));
    }

    function test_Dispatcher_Prevents_Reinit_Attacks() public {
        vm.expectRevert("Initializable: contract is already initialized");
        impl.initialize("IIpolyibc.eth.", dummyConsStateManager);
    }
}
