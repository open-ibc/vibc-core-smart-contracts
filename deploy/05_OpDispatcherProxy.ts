import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import { encodeInitData } from "./utils";

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployer } = await hre.getNamedAccounts();
  const dispatcherImplementation = await hre.deployments.get(
    `DispatcherImplementation-${hre.network.config.chainId}-v${process.env.DEPLOYMENT_VERSION}`
  );
  const initData = encodeInitData(
    hre,
    process.env.DISPATCHER_PORT_PREFIX!,
    "OP"
  );
  await hre.deployments.deploy(
    `OpDispatcherProxy-${hre.network.config.chainId}-v${process.env.DEPLOYMENT_VERSION}`,
    {
      from: deployer,
      args: [dispatcherImplementation.address, initData],
      log: true,
      contract: "ERC1967Proxy",
    }
  );
};

func.tags = ["OpDispatcherProxy", "Dispatcher", "Core"];
// func.dependencies = ["DispatcherConfig"];
export default func;
