#!/usr/bin/env node
import { AccountRegistry, parseObjFromFile } from "..";
import { executeMultisigTx} from "../multisig/safe";

import {
  parseExecuteMultisigTxArgsFromCLI,
} from "../utils/io";
import { isParsedMultiSigWallet } from "../evm/schemas/account";

async function main() {
  const { executor, rpcUrl, txIndex, accountsSpecPath } =
    await parseExecuteMultisigTxArgsFromCLI();

  const accounts = AccountRegistry.load(
    parseObjFromFile(accountsSpecPath),
    "multisig-accounts"
  );

  const multisigAccount = accounts.mustGet(executor);
  if (!isParsedMultiSigWallet(multisigAccount)) {
    throw new Error("Can only execute transactions on a multisig wallet");
  }

  const privateKey = accounts.getSinglePrivateKeyFromAccount(executor);

  executeMultisigTx(
    multisigAccount.safeAddress,
    privateKey,
    BigInt(multisigAccount.chainId),
    rpcUrl,
    txIndex
  );
}

main();
