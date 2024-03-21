import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const { deployer } = await hre.getNamedAccounts();
  console.log(hre.network.config.chainId);
  console.log(
    process.env[`L2_OUTPUT_ORACLE_ADDRESS-${hre.network.config.chainId}`]
  );
  await hre.deployments.deploy(
    `OpProofVerifier-${hre.network.config.chainId}-v${process.env.DEPLOYMENT_VERSION}`,
    {
      from: deployer,
      contract: `OpProofVerifier`,
      args: [
        process.env[`L2_OUTPUT_ORACLE_ADDRESS-${hre.network.config.chainId}`],
      ],
      log: true,
    }
  );
};

func.tags = ["OpProofVerifier", "Core"];
func.dependencies = ["ProofVerifierConfig"];
export default func;
