#!/usr/bin/env node
import {
  AccountRegistry,
  Chain,
  ContractRegistryLoader,
  parseObjFromFile,
} from "..";
import { loadTxRegistry } from "../evm/schemas/tx";
import { sendTxToChain } from "../tx";

import { getOutputLogger } from "../utils/cli";
import { DISPATCHER_SETUP_SPECS_PATH } from "../utils/constants";
import { parseArgsFromCLI } from "../utils/io";

async function main() {
  const { chain, accounts, args } = await parseArgsFromCLI();
  const txSpecs =
    (args.DISPATCHER_SETUP_SPECS_PATH as string) || DISPATCHER_SETUP_SPECS_PATH;

  const upgradeTxs = loadTxRegistry(parseObjFromFile(txSpecs));

  sendTxToChain(
    chain,
    accounts.mustGet(chain.chainName),
    ContractRegistryLoader.emptySingle(),
    upgradeTxs.subset(),
    getOutputLogger(),
    false
  );
}
main();
