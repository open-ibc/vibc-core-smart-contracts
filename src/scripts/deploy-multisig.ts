#!/usr/bin/env node
import { ethers, toBigInt } from "ethers";
import { SingleSigAccountRegistry, parseObjFromFile } from "..";
import { newSafeFromOwner } from "../multisig/safe";

import {
  parseMultisigInitArgsFromCLI,
  saveMultisigAddressToAccountsSpec,
} from "../utils/io";
import { SendingAccountRegistry } from "../evm/schemas/sendingAccount";
import { isUninitializedMultisig } from "../evm/schemas/multisig";

async function main() {
  const { rpcUrl, initiator, accountsSpecPath } =
    await parseMultisigInitArgsFromCLI();

  const accountConfigFromYaml = {
    name: "multisig-accounts",
    registry: parseObjFromFile(accountsSpecPath),
  };

  const accounts = SendingAccountRegistry.loadMultiple([
    accountConfigFromYaml,
  ]).mustGet("multisig-accounts");

  const multisigAccount = accounts.mustGet(initiator);
  if (!isUninitializedMultisig(multisigAccount)) {
    throw new Error(
      "Account read from yaml but isn't a multisig account that needs to be initialized."
    );
  }

  const senderPrivateKey = accounts.getSinglePrivateKeyFromAccount(initiator);
  if (!senderPrivateKey) {
    throw new Error(`Could not find private key for owner ${initiator}`);
  }

  const provider = new ethers.JsonRpcProvider(rpcUrl);
  const chainId = (await provider.getNetwork()).chainId;
  if (!chainId || chainId !== toBigInt(multisigAccount.chainId)) {
    throw new Error(
      `Chain id mismatch between multisig account and rpc url. ${multisigAccount.chainId} is specified in accounts spec, but ${chainId} is the chain id of the rpc url`
    );
  }

  const newSafeAddress = await newSafeFromOwner(
    rpcUrl,
    senderPrivateKey,
    multisigAccount.owners,
    multisigAccount.threshold
  );

  await saveMultisigAddressToAccountsSpec(
    newSafeAddress,
    accountsSpecPath,
    initiator
  );
}

main();
