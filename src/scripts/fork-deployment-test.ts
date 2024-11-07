#!/usr/bin/env node

import { createWriteStream } from "fs";
import { $ } from "zx";
import { MODULE_ROOT_PATH } from "../utils/constants";
import {
  parseArgsFromCLI,
  parseObjFromFile,
  readDeploymentFilesIntoEnv,
} from "../utils/io";
import { ContractRegistryLoader } from "../evm/schemas/contract";
import { getOutputLogger } from "../utils/cli";
import { loadContractUpdateRegistry } from "../evm/schemas/contractUpdate";
import { updateContractsForChain } from "../updateContract";

const main = async () => {
  const { chain, accounts, updateSpecs, anvilPort } = await parseArgsFromCLI();

  const { anvilProcess } = await startAnvilServer(
    chain.rpc,
    anvilPort,
    `anvil-${anvilPort}.out`
  );
  const anvilUrl = `http://127.0.0.1:${anvilPort}`;

  const forkedChain = { ...chain, rpc: anvilUrl };

  const contractUpdates = loadContractUpdateRegistry(
    parseObjFromFile(updateSpecs)
  );

  await updateContractsForChain(
    forkedChain,
    accounts.mustGet(chain.chainName),
    ContractRegistryLoader.emptySingle(),
    contractUpdates,
    getOutputLogger(),
    {
      dryRun: false,
      forceDeployNewContracts: false,
      writeContracts: true,
    }
  );

  let env = {};
  env = await readDeploymentFilesIntoEnv(env, chain); // Read deployment files from non-forked chain to get live addresses

  $.env = {
    ...process.env,
    ...env,
  };

  await $`cd ${MODULE_ROOT_PATH} && forge test --match-contract DispatcherDeployTest --fork-url ${anvilUrl} -vvvv `.pipe(
    createWriteStream("fork-test.out")
  );

  await anvilProcess.kill();
};

// Starts anvil server from an RPC fork, and waits until it's started
const startAnvilServer = async (
  rpcUrl: string,
  port: string,
  outFileName: string
) => {
  const p = $`anvil --port ${port} --fork-url ${rpcUrl}`

  p.pipe(
    createWriteStream(outFileName)
  );
  await waitForAnvilServer(outFileName);
  return { anvilProcess: p };
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
