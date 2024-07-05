#!/usr/bin/env node
import {
  parseVerifyArgsFromCLI,
  readFromDeploymentFile,
  readMetadata,
} from "../utils/io";
import { $, cd } from "zx";
import { MODULE_ROOT_PATH } from "../utils/constants";
/**
 * Take a deployment name, and verify the result to blockscout api
 */
async function main() {
  const { deploymentName, verifierUrl, chainFolder, etherscanApiKey } =
    await parseVerifyArgsFromCLI();

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

  console.log(
    `verifying ${deploymentName}'s deployment with ${deployment.factory} ${
      libraries ? `and libraries ${libraries}` : ``
    }`
  );

  // Run foundry verification command
  let command;
  cd(MODULE_ROOT_PATH); // cd to contracts directory
  try {
    command = await $`forge verify-contract ${args} 2>&1 | tee verify-out.txt`; // command = await $`cd ${MODULE_ROOT_PATH} && forge verify-contract ${
  } catch (e) {
    console.log("error ", e);
  }

  console.log("verification result ", command);
}

main();
