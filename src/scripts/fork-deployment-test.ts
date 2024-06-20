#!/usr/bin/env node

import { createWriteStream, WriteStream } from "fs";
import { $, ProcessPromise } from "zx";
import { deployToChain } from "../deploy";
import { RPC_URL } from "../utils/constants";

const main = async () => {
  console.log("running deployment test to fork from ", RPC_URL);
  await startAnvilServer(RPC_URL, "anvil.out");
  console.log("Anvil server started");
  console.log("running test");

  const anvilUrl = `http://127.0.0.1:8545`;

  const u =
    $`forge test --match-contract DispatcherUpgradeTest --fork-url ${anvilUrl} -vvvv `.pipe(
      createWriteStream("fork-test.out")
    );
};

// Starts anvil server, and waits until it's started
const startAnvilServer = async (rpcUrl: string, outFileName: string) => {
  const p = $`anvil --fork-url ${rpcUrl}`.pipe(createWriteStream(outFileName));
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

// const a
main();
