import { HardhatRuntimeEnvironment } from "hardhat/types";

export const encodeInitData = async (
  hre: HardhatRuntimeEnvironment,
  portPrefix: string,
  clientType: string
) => {
  const proxyAbi = (await hre.deployments.getArtifact("ERC1967Proxy")).abi;
  const client = await hre.deployments.get(
    `${clientType}LightClient-${hre.network.config.chainId}-v${process.env.DEPLOYMENT_VERSION}`
  );
  return new hre.ethers.Interface(proxyAbi).encodeFunctionData("initialize", [
    portPrefix,
    client.address,
  ]);
};
