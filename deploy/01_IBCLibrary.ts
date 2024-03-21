import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployer } = await hre.getNamedAccounts();
  const IBCLibrary = await hre.deployments.deploy("Ibc", {
    from: deployer,
    log: true,
  });
};

func.tags = ["Ibc", "Libraries"];
export default func;
