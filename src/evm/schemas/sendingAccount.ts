import { fs } from "zx";
import { Registry } from "../../utils/registry";
import {
  createWallet,
  isEvmAccounts,
  isEvmAccountsConfig,
  isKeyStore,
} from "./account";
import { isPrivateKey, isSingleSigAccount, Wallet } from "./wallet";
import {
  InitializedMultisig,
  isInitializedMultisig,
  isMultisig,
  isMultisigConfig,
  isUninitializedMultisig,
  UninitializedMultisig,
} from "./multisig";
import path from "path";
import { ethers } from "ethers";

export type SendingAccount =
  | Wallet
  | InitializedMultisig
  | UninitializedMultisig;

export class SendingAccountRegistry extends Registry<SendingAccount> {
  static load(config: unknown[], name: string): SendingAccountRegistry {
    return new SendingAccountRegistry(
      loadSendingAccounts(config),
      config,
      name
    );
  }

  static loadMultiple(registryItems: { name: string; registry: any }[]) {
    const result = new Registry([] as SendingAccountRegistry[], {
      toObj: (t) => {
        return { name: t.name, registry: t.serialize() };
      },
    });
    for (const item of registryItems) {
      result.set(
        item.name,
        SendingAccountRegistry.load(item.registry, item.name)
      );
    }
    return result;
  }

  constructor(
    r: Registry<SendingAccount>,
    private config: any[],
    name: string
  ) {
    super([], { nameInParent: name });
    for (const [name, wallet] of r.entries()) {
      this.set(name, wallet);
    }
  }

  // return the same config obj that was used to load the accounts, but filtered by current account names
  public serialize() {
    const wallets = this.toList();
    return this.config.map((item, index) => {
      const walletItem = wallets[index];
      return {
        name: item.name,
        privateKey: this.getSinglePrivateKeyFromAccount(item.name),
        address: isMultisig(walletItem)
          ? walletItem.wallet.address
          : walletItem.address,
        ...item,
      };
    });
  }

  public getSinglePrivateKeyFromAccount = (accountName: string) => {
    const account = this.mustGet(accountName);
    if (isSingleSigAccount(account) && isPrivateKey(account)) {
      return account.privateKey;
    } else if (
      isInitializedMultisig(account) ||
      isUninitializedMultisig(account)
    ) {
      return account.wallet.privateKey;
    }
    throw new Error(
      `Can't find private key for ${accountName} in this registry`
    );
  };
  // Connect all accounts to the provider
  public connectProviderAccounts = (rpc: string) => {
    const provider = ethers.getDefaultProvider(rpc);
    // const newAccounts = this.subset([]);
    for (const [name, account] of this.entries()) {
      if (isMultisig(account)) {
        const newMultisigWallet = {
          ...account,
          wallet: account.wallet.connect(provider),
        };
        this.set(name, newMultisigWallet, true);
      } else {
        this.set(name, account.connect(provider), true);
      }
    }
    return this;
  };
}

// Load a map of evm accounts from a config through connecting wallets, can either take in sending accounts or not
export function loadSendingAccounts(config: unknown): Registry<SendingAccount> {
  if (!isEvmAccountsConfig(config)) {
    throw new Error(`Error parsing schema: ${config}`);
  }

  const walletMap = new Registry<SendingAccount>([]);

  if (isEvmAccounts(config)) {
    for (const account of config) {
      if (isMultisigConfig(account)) {
        const wallet = createWallet(account.signer);
        const multisigAccount: InitializedMultisig = {
          ...account,
          wallet: wallet,
        };
        walletMap.set(account.name, multisigAccount);
      } else if (isSingleSigAccount(account)) {
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
