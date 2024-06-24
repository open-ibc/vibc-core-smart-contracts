#!/usr/bin/env node

import { createWriteStream, WriteStream } from "fs";
import { $, ProcessPromise } from "zx";
import { deployToChain } from "../deploy";
import { ANVIL_PORT, MODULE_ROOT_PATH, RPC_URL } from "../utils/constants";
import {
  parseArgsFromCLI,
  parseObjFromFile,
  readDeploymentFilesIntoEnv,
} from "../utils/io";
import { ContractRegistryLoader } from "../evm/schemas/contract";
import { getMainLogger, getOutputLogger } from "../utils/cli";
import { sendTxToChain } from "../tx";
import { loadTxRegistry } from "../evm/schemas/tx";

const main = async () => {
  const { chain, accounts, args, deploySpecs, upgradeSpecs, anvilPort } =
    await parseArgsFromCLI();

  const { anvilProcess } = await startAnvilServer(
    chain.rpc,
    anvilPort,
    `anvil-${anvilPort}.out`
  );
  const anvilUrl = `http://127.0.0.1:${anvilPort}`;

  const forkedChain = { ...chain, rpc: anvilUrl };

  const parsedDeploySpecs = parseObjFromFile(deploySpecs);
  if (parsedDeploySpecs) {
    const contracts = ContractRegistryLoader.loadSingle(
      parseObjFromFile(deploySpecs)
    );

    await deployToChain(
      forkedChain,
      accounts.mustGet(chain.chainName),
      contracts.subset(),
      getMainLogger(),
      false,
      false,
      true
    );
  } else {
    console.log("empty deploy file detected, skipping fork deploy for test");
  }

  const parsedUpgradeSpec = parseObjFromFile(upgradeSpecs);
  if (parsedUpgradeSpec) {
    const upgradeTxs = loadTxRegistry(parsedUpgradeSpec);

    await sendTxToChain(
      forkedChain,
      accounts.mustGet(chain.chainName),
      ContractRegistryLoader.emptySingle(),
      upgradeTxs.subset(),
      getOutputLogger(),
      false
    );
  } else {
    console.log("empty upload file detected, skipping upgrade for test");
  }
  let env = {};
  env = await readDeploymentFilesIntoEnv(env, chain); // Read deployment files from non-forked chain to get live addresses

  $.env = {
    ...env,
    ...process.env,
  };

  await $`cd ${MODULE_ROOT_PATH} && forge test --match-contract DispatcherDeployTest --fork-url ${anvilUrl} -vvvv `.pipe(
    createWriteStream("fork-test.out")
  );

  anvilProcess.kill();
};

// Starts anvil server from an RPC fork, and waits until it's started
const startAnvilServer = async (
  rpcUrl: string,
  port: string,
  outFileName: string
) => {
  const p = $`anvil --port ${port} --fork-url ${rpcUrl}`.pipe(
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
