#!/usr/bin/env node

import { createWriteStream, WriteStream } from "fs";
import { $, ProcessPromise } from "zx";
import { deployToChain } from "../deploy";
import { ANVIL_PORT, RPC_URL } from "../utils/constants";
import {
  parseArgsFromCLI,
  parseObjFromFile,
  readDeploymentFilesIntoEnv,
} from "../utils/io";
import { ContractRegistryLoader } from "../evm/schemas/contract";
import { getMainLogger } from "../utils/cli";

const main = async () => {
  await startAnvilServer(RPC_URL, ANVIL_PORT, "anvil.out");
  const anvilUrl = `http://127.0.0.1:8545`;

  const { chain, accounts, args, deploySpecs } = await parseArgsFromCLI();

  console.log("deploying from", deploySpecs);
  const contracts = ContractRegistryLoader.loadSingle(
    parseObjFromFile(deploySpecs)
  );

  chain.rpc = anvilUrl;

  await deployToChain(
    chain,
    accounts.mustGet(chain.chainName),
    contracts.subset(),
    getMainLogger(),
    false,
    false,
    true
  );
  let env = {};
  env = await readDeploymentFilesIntoEnv(env);

  $.env = {
    ...env,
    OWNER: "0xD2b654e3FD89237F8C8a5d7E1AfB5989A13C886e",
    ...process.env,
  };

  await $`forge test --match-contract DispatcherUpgradeTest --fork-url ${anvilUrl} -vvvv `.pipe(
    createWriteStream("fork-test.out")
  );

  console.log("done running tests");
};

// Starts anvil server, and waits until it's started
const startAnvilServer = async (
  rpcUrl: string,
  port: string,
  outFileName: string
) => {
  const p = $`anvil --port ${port} --fork-url ${rpcUrl}`.pipe(
    createWriteStream(outFileName)
  );
  await waitForAnvilServer(outFileName);
  return;
};

const waitForAnvilServer = async (anvilOutFile: string) => {
  const checkOutPutContainsListening = async () => {
    const out = await $`cat ${anvilOutFile}`;
    if (out.stdout.includes("Listening")) {
      return true;
    }
    return false;
  };
  // var intervalId = await setInterval(checkOutPutContainsListening, 1000);
  while (!(await checkOutPutContainsListening())) {
    await setTimeout(() => {}, 1000);
  }
  return;
};

main();
