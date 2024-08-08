#!/usr/bin/env node
import {
  ChainFolder,
  parseObjFromFile,
  parseVerifyArgsFromCLI,
  readFromDeploymentFile,
  readMetadata,
} from "../utils/io";
import { $, cd } from "zx";
import { MODULE_ROOT_PATH } from "../utils/constants";
import { Logger } from "winston";
import { getMainLogger } from "../utils/cli";
import { ContractItemSchema } from "../evm/schemas/contract";
import { loadContractUpdateRegistry } from "../evm/schemas/contractUpdate";

const verifyContract = async (
  deploymentName: string,
  chainFolder: ChainFolder,
  etherscanApiKey: string,
  verifierUrl: string,
  logger: Logger
) => {
  // Read deployment file, so that we can find path to artifact
  const deployment = await readFromDeploymentFile(deploymentName, chainFolder);
  const metadata = await readMetadata(deployment.factory);
  const compilationTarget = JSON.parse(metadata).settings.compilationTarget;
  const contractFile = Object.keys(compilationTarget)[0];
  const contractPath = `${contractFile}:${compilationTarget[contractFile]}`;

  const args = [
    `${deployment.address}`,
    `${contractPath}`,
    `--etherscan-api-key`,
    `${etherscanApiKey}`,
    `--verifier-url=${verifierUrl}`,
  ];
  // Find libraries into string that foundry verification expects if there are any
  let libraries: Record<string, string> | null = null;
  if (deployment.libraries && deployment.libraries.length > 0) {
    libraries = deployment.libraries[0] as Record<string, string>; // Have to parse this out because typechain = weird and stores libs in a 1-length array
    Object.keys(libraries).forEach((lib) => {
      args.push(`--libraries=${lib}:${libraries![lib]}`);
    });
  }

  logger.info(
    `verifying ${deploymentName}'s deployment with ${deployment.factory} ${
      libraries ? `and libraries ${libraries}` : ``
    }`
  );

  // Run foundry verification command
  let command;

   // TODO : Add check here for if already in MODULE_ROOT 

  cd(MODULE_ROOT_PATH); // cd to contracts directory
  try {
    command = await $`forge verify-contract ${args} 2>&1 | tee verify-out.txt`; // command = await $`cd ${MODULE_ROOT_PATH} && forge verify-contract ${
  } catch (e) {
    logger.error("error ", e);
  }

  logger.info("verification result ", command);
};
/**
 * Take a deployment name, and verify the result to blockscout api
 */
async function main() {
  const { verifierUrl, chainFolder, etherscanApiKey, updateSpecs } =
    await parseVerifyArgsFromCLI();

  // Fetch spec from contract

  const contractUpdates = loadContractUpdateRegistry(
    parseObjFromFile(updateSpecs)
  );

  const logger = getMainLogger();
  for (const contractUpdate of contractUpdates.values()) {
    const parsed = ContractItemSchema.safeParse(contractUpdate);
    if (parsed.success) {
      // Only try to verify contractName if it matches the deploymentName
      try {
        await verifyContract(
          parsed.data.name,
          chainFolder,
          etherscanApiKey,
          verifierUrl,
          logger
        );
      } catch (e) {
        logger.error(
          `Failed to verify contract ${parsed.data.name} with error: ${e}`
        );
      }
    }
  }
}

main();
