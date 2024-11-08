#!/usr/bin/env node
import { parseObjFromFile } from "..";
import { executeMultisigTx } from "../multisig/safe";

import { parseExecuteMultisigTxArgsFromCLI } from "../utils/io";
import { isInitializedMultisig } from "../evm/schemas/multisig";
import { SendingAccountRegistry } from "../evm/schemas/sendingAccount";

async function main() {
  const { executor, rpcUrl, txIndex, accountsSpecPath } =
    await parseExecuteMultisigTxArgsFromCLI();

  const accounts = SendingAccountRegistry.load(
    parseObjFromFile(accountsSpecPath),
    "multisig-accounts"
  );

  const multisigAccount = accounts.mustGet(executor);
  if (!isInitializedMultisig(multisigAccount)) {
    throw new Error("Can only execute transactions on a multisig wallet");
  }

  executeMultisigTx(
    multisigAccount,
    rpcUrl,
    txIndex
  );
}

main();
