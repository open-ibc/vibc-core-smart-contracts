import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployer } = await hre.getNamedAccounts();
  const dispatcher = await hre.deployments.get(
    `OpDispatcherProxy-${hre.network.config.chainId}-v${process.env.DEPLOYMENT_VERSION}`
  );

  await hre.deployments.deploy(
    `UniversalChannelHandler-${hre.network.config.chainId}-v${process.env.DEPLOYMENT_VERSION}`,
    {
      from: deployer,
      args: [dispatcher.address],
      log: true,
    }
  );
};

func.tags = ["UniversalChannelHandler", "Middleware"];
func.dependencies = ["OpDispatcherProxy"];
export default func;
