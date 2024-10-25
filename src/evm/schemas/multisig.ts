import { z } from 'zod';
import {singleSigAccount, wallet} from './wallet';

// defined in an account spec, which will be cconvertedi nto an initialized multisig config once we deploy the multisig contract
export const uninitializedMultisigConfig = z.object({
  name: z.string().min(1),
  chainId: z.number(),
  owners: z.array(z.string().min(1)), 
  signer: singleSigAccount 
}) .strict()

// Defined in an account spec, which is not necessarily in a config
export const initializedMultisigConfig = z.object({
  name: z.string().min(1),
  chainId: z.number(),
  safeAddress: z.string().min(1),
  signer: singleSigAccount 
}) .strict()


// Multisig which is described in an account spec but is not yet initialized. (i.e. multisig contract has not been deployed yet)
export const unInitializedMultisig = z.intersection(
    uninitializedMultisigConfig,
    z.object({wallet: wallet})
)

// Multisig which has been deployed & can be used to propose transactions. This is the type that loadEvmAccounts will return for multisig types
export const initializedMultisig = z.intersection(
  initializedMultisigConfig,
    z.object({wallet: wallet})
);

export type UninitializedMultisigConfig = z.infer<typeof uninitializedMultisigConfig>;
export type InitializedMultisigConfig = z.infer<typeof initializedMultisigConfig>;
export type UninitializedMultisig = z.infer<typeof unInitializedMultisig>;
export type InitializedMultisig = z.infer<typeof initializedMultisig>;

// Type Guards
export const isUninitializedMultisigConfig = (
  account: unknown
): account is UninitializedMultisigConfig => {
  return uninitializedMultisigConfig.safeParse(account).success;
};
  
export const isUninitializedMultisig = (
  account: unknown
): account is UninitializedMultisig => {
  return unInitializedMultisig.safeParse(account).success;
};

export const isInitializedMultisigConfig = (
  account: unknown
): account is InitializedMultisigConfig => {
  return initializedMultisigConfig.safeParse(account).success;
};

export const isInitializedMultisig = (
  account: unknown
): account is InitializedMultisig => {
  return initializedMultisig.safeParse(account).success;
};

export const isMultisig = (account: unknown): account is InitializedMultisig | UninitializedMultisig => {
  return isInitializedMultisig(account) || isUninitializedMultisig(account);
};

export const isMultisigConfig = (account: unknown): account is InitializedMultisigConfig | UninitializedMultisigConfig => {
  return isInitializedMultisigConfig(account) || isUninitializedMultisigConfig(account);
}
