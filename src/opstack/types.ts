export type FinalPeptideL1DeployConfig = {
    l1StartingBlockTag: string
    l1ChainID: number
    l2ChainID: number
    l2BlockTime: number
    l2OutputOracleSubmissionInterval: number
    l2OutputOracleStartingTimestamp: number
    l2OutputOracleStartingBlockNumber: number
    l2OutputOracleProposer: string
    l2OutputOracleChallenger: string
    finalizationPeriodSeconds: number
    proxyAdminOwner: string
    gasPriceOracleOverhead: number
    gasPriceOracleScalar: number
    l2GenesisBlockGasLimit: string
    batchInboxAddress: string
    batchSenderAddress: string
    maxSequencerDrift: number
    sequencerWindowSize: number
    channelTimeout: number
    l2GenesisRegolithTimeOffset?: string
    l2GenesisCanyonTimeOffset?: string
    l2GenesisDeltaTimeOffset?: string
    usePlasma: boolean
}

export type PeptideGenesis = {
    genesis_time: string
    genesis_block: {
      hash: string
      number: number
    }
    chain_id: string
    app_state: string
    l1?: {
      hash: string
      number: number
    }
    initial_height: number
  }

  export type PeptideL1DeployAddress = {
    // Reduced set of contracts from custom deploy script
    L2OutputOracle: string
    L2OutputOracleProxy: string
    OptimismPortalProxy: string
    ProxyAdmin: string
    SafeProxyFactory: string
    SafeSingleton: string
    SystemOwnerSafe: string
    DataAvailabilityChallengeProxy: string | undefined
  }
  export type OpNodeRollupConfig = {
    genesis: {
      l1: {
        hash: string
        number: number
      }
      l2: {
        hash: string
        number: number
      }
      l2_time: number
      system_config: {
        batcherAddr: string
        overhead: string
        scalar: string
        gasLimit: number
      }
    }
    block_time: number
    max_sequencer_drift: number
    seq_window_size: number
    channel_timeout: number
    l1_chain_id: number
    l2_chain_id: number
    regolith_time?: number
    canyon_time?: number
    delta_time?: number
    ecotone_time?: number
    fjord_time?: number
    batch_inbox_address: string
    deposit_contract_address: string
    l1_system_config_address: string
    protocol_versions_address: string
    /** only applicable if plasma is enabled */
    plasma_config?: {
      /** L1 DataAvailabilityChallenge contract proxy address */
      da_challenge_contract_address: string
      /** DA challenge window value set on the DAC contract. Used in plasma mode to compute when a commitment can no longer be challenged. */
      da_challenge_window: number
      /** DA resolve window value set on the DAC contract. Used in plasma mode to compute when a challenge expires and trigger a reorg if needed. */
      da_resolve_window: number
    }
  }