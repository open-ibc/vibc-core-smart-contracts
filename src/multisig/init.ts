import SafeApiKit from "@safe-global/api-kit";
import Safe, { SafeFactory } from "@safe-global/protocol-kit";
import {
  MetaTransactionData,
  OperationType,
} from "@safe-global/safe-core-sdk-types";

/*
 * Init a safe owner from a given private key to a give RPC's network
 * Submits an evm tx to create a new safe proxy
 */
export const newSafeFromOwner = async (
  RPC_URL: string,
  ownerKey: string,
  owners: string[],
  threshold: number,
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
export const proposeTransaction = async (
  safeAddress: string,
  toAddress: string,
  txData: string,
  proposerPrivateKey: string,
  chainId: bigint,
  rpcUrl: string,
  value = "0",
  operation = OperationType.Call
) => {
  const apiKit = new SafeApiKit({
    chainId,
  });

  const txProposer = await Safe.init({
    provider: rpcUrl,
    signer: proposerPrivateKey,
    safeAddress: safeAddress,
  });

  console.log(
    `proposing transaction from ${txProposer} @safe ${safeAddress} to ${toAddress} with value: ${value} and data: ${txData}`
  );

  const txProposerAddress = await txProposer.getAddress();

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

  // Propose transaction to the service
  await apiKit.proposeTransaction({
    safeAddress,
    safeTransactionData: safeTransaction.data,
    safeTxHash,
    senderAddress: await txProposerAddress,
    senderSignature: proposerSignature.data,
  });

  console.log(
    `Transaction has been proposed with hash: ${safeTxHash} from address ${txProposerAddress}`
  );
  return await apiKit.getTransaction(safeTxHash);
};

export const executeTransaction = async (
  safeAddress: string,
  executorPrivateKey: string,
  chainId: bigint,
  rpcUrl: string,
  pendingtransactionIndex: number
) => {
  const apiKit = new SafeApiKit({
    chainId,
  });

  const executor = await Safe.init({
    provider: rpcUrl,
    signer: executorPrivateKey,
    safeAddress,
  });

  const transactions = await apiKit.getPendingTransactions(safeAddress);
  if (transactions.results.length <= pendingtransactionIndex) {
    throw new Error(
      `Invalid transaction index - trying to access index ${pendingtransactionIndex} but only ${transactions.results.length} pending txs in safe`
    );
  }

  const tx = transactions.results[pendingtransactionIndex];
  return await executor.executeTransaction(tx);
};
