import { z } from "zod";
import { ethers } from "ethers";
import fs from "fs";
import path from "path";
import { Registry } from "../../utils/registry";
import { renderString } from "../../utils/io";
import { initializedMultisigConfig, uninitializedMultisigConfig } from "./multisig";
import {
  isMnemonic,
  isPrivateKey,
  isSingleSigAccount,
  SingleSigAccount,
  singleSigAccount,
  Wallet,
} from "./wallet";

// ethers wallet with encryption
// geth compatible keystore
const keyStore = z.object({
  dir: z.string().min(1),
  password: z.optional(z.string()),
});

export type KeyStore = z.infer<typeof keyStore>;

export const evmAccounts = z.array(
  z.union([singleSigAccount, initializedMultisigConfig, uninitializedMultisigConfig])
); // Type of account that one can send transactions from
export type EvmAccounts = z.infer<typeof evmAccounts>;
export const EvmAccountsConfig = z.union([evmAccounts, keyStore]);
export type EvmAccountConfig = z.infer<typeof EvmAccountsConfig>;

export const isKeyStore = (account: unknown): account is KeyStore => {
  return keyStore.safeParse(account).success;
};

export const isEvmAccounts = (account: unknown): account is EvmAccounts => {
  return evmAccounts.safeParse(account).success;
};
export const isEvmAccountsConfig = (
  account: unknown
): account is EvmAccountConfig => {
  return EvmAccountsConfig.safeParse(account).success;
};


export class SingleSigAccountRegistry extends Registry<Wallet> {
  static load(config: unknown[], name: string): SingleSigAccountRegistry {
    return new SingleSigAccountRegistry(loadEvmAccounts(config), config, name);
  }

  static loadMultiple(registryItems: { name: string; registry: any }[]) {
    const result = new Registry([] as SingleSigAccountRegistry[], {
      toObj: (t) => {
        return { name: t.name, registry: t.serialize() };
      },
    });
    for (const item of registryItems) {
      result.set(
        item.name,
        SingleSigAccountRegistry.load(item.registry, item.name)
      );
    }
    return result;
  }

  constructor(r: Registry<Wallet>, private config: any[], name: string) {
    super([], { nameInParent: name });
    for (const [name, wallet] of r.entries()) {
      this.set(name, wallet);
    }
  }

  // return the same config obj that was used to load the accounts, but filtered by current account names
  public serialize() {
    const wallets = this.toList();
    return this.config.map((item, index) => {
      return {
        name: item.name,
        privateKey: wallets[index].privateKey,
        address: wallets[index].address,
        ...item,
      };
    });
  }

  public getSinglePrivateKeyFromAccount = (accountName: string) => {
    const account = this.mustGet(accountName);
    if (!isMnemonic(account)) {
      return account.privateKey;
    }
    return account.mnemonic;
  };
  // Connect all accounts to the provider
  public connectProviderAccounts = (rpc: string) => {
    const provider = ethers.getDefaultProvider(rpc);
    for (const [name, account] of this.entries()) {
      this.set(name, account.connect(provider), true);
    }
    return this;
  };
}

// load a Map of { [name: string]: Wallet } from EvmAccountsSchema object
export function loadEvmAccounts(config: unknown): Registry<Wallet> {
  if (!isEvmAccountsConfig(config)) {
    throw new Error(`Error parsing schema: ${config}: \n ${EvmAccountsConfig.safeParse(config).error}`);
  }
  const walletMap = new Registry<Wallet>([]);

  if (isEvmAccounts(config)) {
    for (const account of config) {
      if (isSingleSigAccount(account)) {
        walletMap.set(account.name, createWallet(account));
      }
    }
  } else if (isKeyStore(config)) {
    const files = fs.readdirSync(config.dir);
    for (const file of files) {
      const filePath = path.join(config.dir, file);
      const json = fs.readFileSync(filePath, "utf8");
      const wallet = ethers.Wallet.fromEncryptedJsonSync(
        json,
        config.password ?? ""
      );
      walletMap.set(wallet.address, wallet);
    }
  }
  return walletMap;
}

export function createWallet(opt: SingleSigAccount): Wallet {
  if (isPrivateKey(opt)) {
    let renderedPrivatekey = opt.privateKey;
    if (!ethers.isHexString(renderedPrivatekey, 32)) {
      // check if is a valid private key - if not, see if it represents an env variable that represents a private key.
      try {
        renderedPrivatekey = renderString(opt.privateKey, process.env); // look up in env if not.
      } catch (e) {
        console.log("no valid private key found for account spec", e);
      }
    }
    return new ethers.Wallet(renderedPrivatekey);
  }

  // Otherwise is a mnemonic
  if (opt.mnemonic && typeof opt.mnemonic === "string") {
    let wallet = ethers.Wallet.fromPhrase(opt.mnemonic);
    if (typeof opt.path === "string" && opt.path.length > 0) {
      wallet = ethers.HDNodeWallet.fromPhrase(
        opt.mnemonic,
        undefined,
        opt.path
      );
    }
    // if account.index is specified, derive the child wallet by index
    if (Number.isInteger(opt.index)) {
      wallet = wallet.deriveChild(opt.index!);
    }
    return wallet;
  }
  throw new Error(
    `invalid wallet config, must provide at least one of {privateKey, mnemonic}, but got: ${JSON.stringify(
      opt
    )}`
  );
}
