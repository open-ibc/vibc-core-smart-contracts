import { ethers } from "ethers"
import { fs } from "zx"
import { FinalPeptideL1DeployConfig, OpNodeRollupConfig, PeptideGenesis, PeptideL1DeployAddress } from "./types"
import { Logger } from "winston"
import { Provider } from "ethers"
import path from "path"

  
export  async function generatePeptideL1Origin(l1DeployConfig: FinalPeptideL1DeployConfig, provider: Provider, logger: Logger, L1OriginOutPath: string) {
    // get the latest L1 block
    const tag = l1DeployConfig['l1StartingBlockTag']
    const l1Block = await provider.getBlock(tag)
    if(!l1Block || !l1Block.hash){
        throw new Error(`L1 block not found at ${tag}`)
    }

    await fs.writeFile(L1OriginOutPath, JSON.stringify({ hash: l1Block.hash, number: l1Block.number }, null, 2))
    logger.info(`L1 origin written to ${L1OriginOutPath}`)
    return { hash: l1Block.hash, number: l1Block.number }
  }

export async function generateRollupConfig(l1DeployConfig: FinalPeptideL1DeployConfig, genesis: PeptideGenesis, addresses: PeptideL1DeployAddress, baseOutPath: string, provider: Provider, logger: Logger) {

    const toNumber = (s: string | undefined) => (s ? ethers.toNumber(s) : undefined)

    const l1OutPath = path.resolve(baseOutPath, 'l1_origin.json')
    const rollupConfigOutPath = path.resolve(baseOutPath, 'rollup_config.json')
    
    const l1Block = genesis.l1?? await generatePeptideL1Origin(l1DeployConfig, provider, logger, l1OutPath)
    const rollupConfig: OpNodeRollupConfig= {
      genesis: {
        l1: l1Block,
        l2: genesis.genesis_block,
        l2_time: Math.floor(Date.parse(genesis.genesis_time) / 1000), // l2 time is in seconds!
        system_config: {
          batcherAddr: l1DeployConfig.batchSenderAddress,
          overhead: ethers.zeroPadValue(ethers.toBeHex(l1DeployConfig.gasPriceOracleOverhead), 32),
          scalar: ethers.zeroPadValue(ethers.toBeHex(l1DeployConfig.gasPriceOracleScalar), 32),
          gasLimit: ethers.toNumber(l1DeployConfig.l2GenesisBlockGasLimit),
        },
      },
      block_time: l1DeployConfig.l2BlockTime,
      max_sequencer_drift: l1DeployConfig.maxSequencerDrift,
      seq_window_size: l1DeployConfig.sequencerWindowSize,
      channel_timeout: l1DeployConfig.channelTimeout,
      l1_chain_id: l1DeployConfig.l1ChainID,
      l2_chain_id: l1DeployConfig.l2ChainID,
      regolith_time: toNumber(l1DeployConfig.l2GenesisRegolithTimeOffset),
      canyon_time: toNumber(l1DeployConfig.l2GenesisCanyonTimeOffset),
      delta_time: toNumber(l1DeployConfig.l2GenesisDeltaTimeOffset),
      batch_inbox_address: l1DeployConfig.batchInboxAddress,
      deposit_contract_address: addresses.OptimismPortalProxy,
      l1_system_config_address: '0x0000000000000000000000000000000000000000', // We don't deploy a system config so we leave this as uninitialized address
      protocol_versions_address: '0x0000000000000000000000000000000000000000', // We don't deploy a protocol_versions address so we leave this as uninitialized address
    }
    if (l1DeployConfig.usePlasma) {
      if (!addresses.DataAvailabilityChallengeProxy) {
        throw new Error('DataAvailabilityChallengeProxy is not found in deployment but usePlasma is true')
      }
      rollupConfig.plasma_config = {
        da_challenge_contract_address: addresses.DataAvailabilityChallengeProxy,
        da_challenge_window: 160,
        da_resolve_window: 160,
      }
    }

    await fs.writeFile(rollupConfigOutPath, JSON.stringify(rollupConfig, null, 2))
    logger.info(`Peptide OpNode Rollup config written to ${rollupConfigOutPath}`)
  }