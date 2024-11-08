import SafeApiKit from "@safe-global/api-kit";
import Safe, { SafeFactory } from "@safe-global/protocol-kit";
import {
  MetaTransactionData,
  OperationType,
  TransactionResult,
} from "@safe-global/safe-core-sdk-types";
import { InitializedMultisig } from "../evm/schemas/multisig";

/*
 * Init a safe owner from a given private key to a give RPC's network
 * Submits an evm tx to create a new safe proxy
 */
export const newSafeFromOwner = async (
  RPC_URL: string,
  ownerKey: string,
  owners: string[],
  threshold: number
) => {
  const safeFactory = await SafeFactory.init({
    provider: RPC_URL,
    signer: ownerKey,
  });

  const safe = await safeFactory.deploySafe({
    safeAccountConfig: { owners, threshold },
  });
  const safeAddress = await safe.getAddress();

  console.log(`Safe has been deployed to: ${safeAddress}`);

  return safeAddress;
};

/**
 *
 * @param data Propose a transaction from a signer account
 */
export async function proposeTransaction(
  initializedMultisig: InitializedMultisig,
  toAddress: string,
  txData: string,
  rpcUrl: string,
  nonce: number,
  value = "0",
  operation = OperationType.Call
) {
  console.log(
    `proposing transaction from \n \n @safe ${initializedMultisig.safeAddress} to ${toAddress} with value: ${value} and data: ${txData}, and nonce : ${nonce}`
  );

  const txProposer = await Safe.init({
    provider: rpcUrl,
    signer: initializedMultisig.wallet.privateKey,
    safeAddress: initializedMultisig.safeAddress,
  });

  const txProposerAddress = initializedMultisig.wallet.address;

  const safeTransactionData: MetaTransactionData = {
    to: toAddress,
    value,
    data: txData,
    operation: operation,
  };

  const safeTransaction = await txProposer.createTransaction({
    transactions: [safeTransactionData],
  });

  const safeTxHash = await txProposer.getTransactionHash(safeTransaction);

  const proposerSignature = await txProposer.signHash(safeTxHash);
  if (initializedMultisig.txServiceUrl) {
    // For custom tx services, we have to submit our own http request since it seems like the safeApiKit.proposeTransaction method isn't compatible with some tx service urls

    const url = `${initializedMultisig.txServiceUrl}/v1/chains/${initializedMultisig.chainId}/transactions/${initializedMultisig.safeAddress}/propose`;
    const body = JSON.stringify({
      ...safeTransaction.data,
      nonce: nonce.toString(),
      signature: proposerSignature.data,
      sender: txProposerAddress,
      safeTxHash,
    });

    console.log(`proposing transaction to tx service:${url} \n `);
    console.log(`with body${body}`);
    try {
      const res = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body,
      });

      const json = await res.json();

      if (!res.ok) {
        console.log(
          `received non 200 response from tx service ${res.status} ${json.message}`
        );
      }
      return json;
    } catch (e) {
      console.log(
        `error sending pending multisig tx to tx service ${initializedMultisig.txServiceUrl}`,
        e
      );

      return;
    }
  } else {
    const apiKit = new SafeApiKit({
      chainId: BigInt(initializedMultisig.chainId),
    });
    // Propose transaction to the service
    await apiKit.proposeTransaction({
      safeAddress: initializedMultisig.safeAddress,
      safeTransactionData: safeTransaction.data,
      safeTxHash,
      senderAddress: txProposerAddress,
      senderSignature: proposerSignature.data,
    });
    console.log(
      `Transaction has been proposed with hash: ${safeTxHash} from address ${txProposerAddress}`
    );
    return await apiKit.getTransaction(safeTxHash);
  }
}

/**
 * Execute a multisig tx from an account generated from proposeTransaction
 */
export const executeMultisigTx = async (
  initializedMultisig: InitializedMultisig,
  rpcUrl: string,
  pendingTransactionIndex: number
): Promise<TransactionResult> => {
  const apiKit = new SafeApiKit({
    chainId: BigInt(initializedMultisig.chainId),
    txServiceUrl: initializedMultisig.txServiceUrl,
  });

  const executor = await Safe.init({
    provider: rpcUrl,
    signer: initializedMultisig.wallet.privateKey,
    safeAddress: initializedMultisig.safeAddress,
  });

  const transactions = await apiKit.getPendingTransactions(
    initializedMultisig.safeAddress
  );

  if (transactions.results.length <= pendingTransactionIndex) {
    throw new Error(
      `Invalid transaction index - trying to access index ${pendingTransactionIndex} but only ${transactions.results.length} pending txs in safe`
    );
  }

  const tx = transactions.results[pendingTransactionIndex];
  return await executor.executeTransaction(tx);
};

export async function fetchNonceFromSafeAddress(
  rpcUrl: string,
  safeAddress: string
) {
  const safe = await Safe.init({
    provider: rpcUrl,
    safeAddress,
  });

  return safe.getNonce();
}
