import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployer } = await hre.getNamedAccounts();
  await hre.deployments.deploy(
    `DispatcherImplementation-${hre.network.config.chainId}-v${process.env.DEPLOYMENT_VERSION}`,
    {
      from: deployer,
      args: [],
      libraries: {
        Ibc: (await hre.deployments.get("Ibc")).address,
      },
      contract: "Dispatcher",
      log: true,
    }
  );
};

func.tags = ["DispatcherImplementation", "Dispatcher", "Core"];
func.dependencies = ["Ibc"];
export default func;
