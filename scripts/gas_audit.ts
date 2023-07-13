import * as hre from 'hardhat'

const ethers = hre.ethers

async function getTxGasCost(txHash, txName) {
  const trace = await hre.network.provider.send('debug_traceTransaction', [txHash])
  console.log(`tx [${txName}] gas cost: ${trace.gas}`)
}

async function main() {
  const [deployer] = await ethers.getSigners()

  // deploy GasAudit contract
  const GasAuditFactory = await ethers.getContractFactory('GasAudit')
  const gasAudit = await GasAuditFactory.deploy()

  console.log('GasAudit deployed to:', gasAudit.address)

  const callArgs = {
    address: deployer.address,
    channel: 'channel-1',
    counterpartyChannel: 'channel-123',
    port: 'transfer'
  }

  const tx1 = await gasAudit
    .connect(deployer)
    .callWithBytes32(
      callArgs.address,
      ethers.utils.formatBytes32String(callArgs.channel),
      callArgs.port,
      ethers.utils.formatBytes32String(callArgs.counterpartyChannel)
    )

  const tx2 = await gasAudit
    .connect(deployer)
    .callWithString(deployer.address, callArgs.channel, callArgs.port, callArgs.counterpartyChannel)

  await getTxGasCost(tx1.hash, 'callWithBytes32')
  await getTxGasCost(tx2.hash, 'callWithString')

  /**
   * Output
    GasAudit deployed to: 0x5FbDB2315678afecb367f032d93F642f64180aa3
    tx [callWithBytes32] gas cost: 48046
    tx [callWithString] gas cost: 50396
   */
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error)
    process.exit(1)
  })
