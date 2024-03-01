import {IDispatcher} from "../contracts/interfaces/IDispatcher.sol";
import {LightClient} from "../contracts/interfaces/LightClient.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Dispatcher} from "../contracts/core/Dispatcher.sol";

pragma solidity ^0.8.0;

library DeploymentUtils {
    function deployDispatcherProxyAndImpl(string memory initPortPrefix, LightClient lightClient)
        public
        returns (IDispatcher proxy)
    {
        Dispatcher impl = new Dispatcher();
        proxy = IDispatcher(
            address(
                new ERC1967Proxy(
                    address(impl), abi.encodeWithSelector(Dispatcher.initialize.selector, initPortPrefix, lightClient)
                )
            )
        );
    }
}
