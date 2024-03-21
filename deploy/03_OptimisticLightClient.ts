import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployer } = await hre.getNamedAccounts();
  const OpProofVerifier = await hre.deployments.get(
    `OpProofVerifier-${hre.network.config.chainId}-v${process.env.DEPLOYMENT_VERSION}`
  );
  await hre.deployments.deploy(
    `OptimisticLightClient-${hre.network.config.chainId}-v${process.env.DEPLOYMENT_VERSION}`,
    {
      from: deployer,
      contract: `OptimisticLightClient`,
      args: [
        process.env.FRAUD_PROOF_WINDOW,
        OpProofVerifier.address,
        process.env[`L1_BLOCK_PROVIDER-${hre.network.config.chainId}`],
      ],
      log: true,
    }
  );
};

func.tags = ["OptimisticLightClient", "Core"];
func.dependencies = ["OpProofVerifier"];
export default func;
