import { loadFixture } from '@nomicfoundation/hardhat-network-helpers'
import { describe, it } from 'mocha'
import { ethers } from 'hardhat'
import { expect } from 'chai'
import { BigNumber } from 'ethers'
import { Dispatcher } from '../typechain-types'

const toBytes32 = ethers.utils.formatBytes32String
const toBytes = ethers.utils.toUtf8Bytes

interface Proof {
  a: [BigNumber, BigNumber]
  b: [[BigNumber, BigNumber], [BigNumber, BigNumber]]
  c: [BigNumber, BigNumber]
}

function getValidProof(): Proof {
  return {
    a: [
      BigNumber.from('13449388914393258752883032560537386278857542833249142697090243871249761501123'),
      BigNumber.from('5963894333042515966217276339656894890750758020607775733717462915787234629927')
    ],
    b: [
      [
        BigNumber.from('4811559872397934450173412387101758297072581261546553338353504577141293696514'),
        BigNumber.from('18400634037991283592418145553628778322894277765243742619561207628896194710939')
      ],
      [
        BigNumber.from('17903685207039300384995795331083569887497623740119108595975170464164316221644'),
        BigNumber.from('9246628133289276308945311797896077179503891414159382179119544904154776510385')
      ]
    ],
    c: [
      BigNumber.from('17432552394458345841788798376121543520587716339044416231610790827968220675517'),
      BigNumber.from('15082220514596158133191868403239442750535261032426474092101151620016661078026')
    ]
  }
}

function genState(appHash: any, valsetHash: any, time: any, height: any) {
  return {
    app_hash: BigNumber.from(appHash),
    valset_hash: BigNumber.from(valsetHash),
    time: BigNumber.from(time),
    height: BigNumber.from(height)
  }
}

async function getLatestConsensusState(dispatcher: Dispatcher) {
  const latestConsensusStateData = await dispatcher.latestConsensusState()
  return {
    app_hash: latestConsensusStateData.app_hash,
    valset_hash: latestConsensusStateData.valset_hash,
    time: latestConsensusStateData.time,
    height: latestConsensusStateData.height
  }
}

const proof = getValidProof()

describe('IBC Core Smart Contract', function () {
  // Constants for testing
  const trustedState = genState(
    80990,
    BigNumber.from('590199110038530808131927832294665177527506259518072095333098842116767351199'),
    7000040,
    1000
  )
  const untrustedState = genState(
    10934,
    BigNumber.from('7064503680087416120706887577693908749828198688716609274705703517077803898371'),
    7002040,
    1020
  )
  const C = {
    ClientState: toBytes32('clientState'),
    ConsensusStates: [trustedState, genState(1, 1, 1, 1)],
    InitClientMsg: {
      clientState: toBytes32('clientState'),
      consensusState: trustedState
    },
    UpdateClientMsg: {
      consensusState: untrustedState,
      zkProof: proof
    },
    UpgradeClientMsg: { clientState: toBytes32('clientState'), consensusState: genState(0, 0, 0, 0) },
    ConnHops1: ['connection-0', 'connection-2'],
    ConnHops2: ['connection-1', 'connection-3'],
    EmptyVersion: '',
    V1: '1.0',
    V2: '2.0',
    BEmptyVersion: toBytes32(''), // TODO: will be deleted after API re-design
    BV1: toBytes32('1.0'),// TODO: will be deleted after API re-design
    BV2: toBytes32('2.0'),// TODO: will be deleted after API re-design
    InvalidVersion: toBytes32('invalid-version'),
    Unordered: 0,
    Ordered: 1,
    InvalidProof: { proofHeight: 42, proof: ethers.utils.toUtf8Bytes('') },
    ValidProof: { proofHeight: 42, proof: ethers.utils.toUtf8Bytes('valid proof') },
    ChannelIds: ['channel-0', 'channel-1'].map(toBytes32),
    RemoteChannelIds: ['channel-100', 'channel-101'].map(toBytes32),
    EmptyChannelId: toBytes32(''),
    BscPortId: toBytes32('polyibc.bsc.9876543210'),
    Packets: [
      {
        msg: 'hello ibc',
        sequence: 0,
        timeout: ethers.BigNumber.from(123456789),
        fee: ethers.utils.parseEther('0.123')
      }
    ]
  }

  // Get all contract factories for testing
  const getContractFactories = async () => {
    const Dispatcher = await ethers.getContractFactory('Dispatcher')
    const Verifier = await ethers.getContractFactory('Verifier')
    const Mars = await ethers.getContractFactory('Mars')
    return { Dispatcher, Verifier, Mars }
  }

  const getSignerAccounts = async () => {
    const signers = await ethers.getSigners()
    return {
      owner: signers[0],
      user1: signers[1],
      user2: signers[2],
      relayer: signers[3],
      escrow: signers[4],
      otherUsers: signers.slice(5)
    }
  }

  /**
   * @description Deploy IBC Core SC and verifier contract
   */
  async function deployIbcCoreFixture() {
    const factories = await getContractFactories()
    const accounts = await getSignerAccounts()
    // Deploy Verifier and CoreSC contracts by owner
    const verifier = await factories.Verifier.deploy()
    const dispatcher = await factories.Dispatcher.deploy(verifier.address, accounts.escrow.address, 'polyibc.eth.')

    return { accounts, verifier, dispatcher, factories }
  }

  /**
   * @description Deploy IBC Core SC and verifier contract, create Polymer client and deploy a dApp Mars contract as an IBC-enabled contract
   */
  async function setupCoreClientFixture() {
    const { accounts, verifier, dispatcher, factories } = await loadFixture(deployIbcCoreFixture)

    // Deploy Mars contract by user1
    const mars = await factories.Mars.connect(accounts.user1).deploy()
    // Set up Polymer light client on CoreSC
    await dispatcher.createClient(C.InitClientMsg).then((tx) => tx.wait())

    return { accounts, verifier, dispatcher, mars }
  }

  /**
   * @description
   * Set up clients and establish a channel
   */
  async function setupChannelFixture() {
    const { accounts, dispatcher, mars } = await loadFixture(setupCoreClientFixture)
    const channel = {
      portAddress: mars.address,
      channelId: C.ChannelIds[0],
      version: C.BV1,
      ordering: C.Unordered,
      connectionHops: C.ConnHops1,
      counterpartyPortId: C.BscPortId,
      counterpartyChannelId: C.RemoteChannelIds[0]
    }
    await dispatcher
      .connect(accounts.relayer)
      .connectIbcChannel(
        channel.portAddress,
        channel.channelId,
        channel.connectionHops,
        channel.ordering,
        channel.counterpartyPortId,
        channel.counterpartyChannelId,
        channel.version,
        C.ValidProof
      )
      .then((tx) => tx.wait())
    return { accounts, dispatcher, mars, channel, packets: C.Packets }
  }

  describe('createClient', function () {
    it('only owner can create a new client', async function () {
      const { accounts, dispatcher } = await loadFixture(deployIbcCoreFixture)

      await expect(dispatcher.connect(accounts.user1).createClient(C.InitClientMsg)).to.be.revertedWith(
        'Ownable: caller is not the owner'
      )
    })

    it('should create a new client', async function () {
      const { dispatcher } = await loadFixture(setupCoreClientFixture)
      const latestConsensusState = await getLatestConsensusState(dispatcher)

      expect(latestConsensusState).to.deep.equal(C.ConsensusStates[0])
    })

    it('cannot create call creatClient twice', async function () {
      const { dispatcher } = await loadFixture(setupCoreClientFixture)
      await expect(dispatcher.createClient(C.InitClientMsg)).to.be.revertedWith('Client already created')
    })
  })

  describe('updateClient', function () {
    it('should update the consensus state of an existing client', async function () {
      const { dispatcher } = await loadFixture(setupCoreClientFixture)

      await dispatcher.updateClient(C.UpdateClientMsg)
      const latestConsensusState = await getLatestConsensusState(dispatcher)

      expect(latestConsensusState).to.deep.equal(C.UpdateClientMsg.consensusState)
    })

    it('cannot update client with invalid consensusState', async function () {
      const { dispatcher } = await loadFixture(setupCoreClientFixture)
      const invalidConsState = genState(0, 0, 0, 0)
      const invalidUpdateClientMsg = { ...C.UpdateClientMsg, consensusState: invalidConsState }
      await expect(dispatcher.updateClient(invalidUpdateClientMsg)).to.be.revertedWith(
        'UpdateClientMsg proof verification failed'
      )
    })
  })

  describe('upgradeClient', function () {
    it('should upgrade the client', async function () {
      const { dispatcher } = await loadFixture(setupCoreClientFixture)

      await dispatcher.upgradeClient(C.UpgradeClientMsg)
      const latestConsensusState = await getLatestConsensusState(dispatcher)

      expect(latestConsensusState).to.deep.equal(C.UpgradeClientMsg.consensusState)
    })

    it('cannot upgrade by non-owner', async function () {
      const { dispatcher, accounts } = await loadFixture(setupCoreClientFixture)

      await expect(dispatcher.connect(accounts.user1).upgradeClient(C.UpgradeClientMsg)).to.be.revertedWith(
        'Ownable: caller is not the owner'
      )
    })
  })

  describe('openIbcChannel', function () {
    it('ChanOpenInit', async function () {
      const { dispatcher, mars, accounts } = await loadFixture(setupCoreClientFixture)

      await expect(
        dispatcher
          .connect(accounts.relayer)
          .openIbcChannel(mars.address, C.V1, C.Unordered, C.ConnHops1, C.BscPortId, C.EmptyChannelId, C.EmptyVersion)
      )
        .to.emit(dispatcher, 'OpenIbcChannel')
        .withArgs(mars.address, C.V1, C.Unordered, C.ConnHops1, C.BscPortId, C.EmptyChannelId)
    })

    it('ChanOpenTry', async function () {
      const { dispatcher, mars, accounts } = await loadFixture(setupCoreClientFixture)

      await expect(
        dispatcher
          .connect(accounts.relayer)
          .openIbcChannel(
            mars.address,
            C.EmptyVersion,
            C.Ordered,
            C.ConnHops2,
            C.BscPortId,
            C.RemoteChannelIds[0],
            C.V2
          )
      )
        .to.emit(dispatcher, 'OpenIbcChannel')
        .withArgs(mars.address, C.V2, C.Ordered, C.ConnHops2, C.BscPortId, C.RemoteChannelIds[0])
    })

    it('unsupported version', async function () {
      const { dispatcher, mars, accounts } = await loadFixture(setupCoreClientFixture)

      // invalid version in ChanOpenInit
      await expect(
        dispatcher
          .connect(accounts.relayer)
          .openIbcChannel(
            mars.address,
            C.InvalidVersion,
            C.Unordered,
            C.ConnHops1,
            C.EmptyChannelId,
            C.BscPortId,
            C.EmptyVersion
          )
      ).to.be.revertedWith('Unsupported version')

      // invalid version in ChanOpenTry
      await expect(
        dispatcher
          .connect(accounts.relayer)
          .openIbcChannel(
            mars.address,
            C.EmptyVersion,
            C.Unordered,
            C.ConnHops2,
            C.RemoteChannelIds[0],
            C.BscPortId,
            C.InvalidVersion
          )
      ).to.be.revertedWith('Unsupported version')
    })

    it('onOpenIbcChannel callback error', async function () {
      const { dispatcher, mars, accounts } = await loadFixture(setupCoreClientFixture)

      await expect(
        dispatcher
          .connect(accounts.relayer)
          .openIbcChannel(mars.address, C.V1, C.Unordered, C.ConnHops1, 'portX', C.RemoteChannelIds[0], C.EmptyVersion)
      ).to.be.revertedWith('Invalid counterpartyPortId')
    })

    describe('connectIbcChannel', function () {
      it('ChanOpenAck/Confirm', async function () {
        const { dispatcher, mars, accounts } = await loadFixture(setupCoreClientFixture)

        await expect(
          dispatcher
            .connect(accounts.relayer)
            .connectIbcChannel(
              mars.address,
              C.ChannelIds[0],
              C.ConnHops1,
              C.Ordered,
              C.BscPortId,
              C.RemoteChannelIds[0],
              C.BV1,
              C.ValidProof
            )
        )
          .to.emit(dispatcher, 'ConnectIbcChannel')
          .withArgs(mars.address, C.ChannelIds[0], C.BscPortId, C.RemoteChannelIds[0], C.ConnHops1)

        // confirm channel is owned by the port address
        const channel = await dispatcher.getChannel(mars.address, C.ChannelIds[0])
        expect(channel).deep.equal([C.BV1, C.Ordered, C.ConnHops1, C.BscPortId, C.RemoteChannelIds[0]])
      })

      it('invalid proof', async function () {
        const { dispatcher, mars, accounts } = await loadFixture(setupCoreClientFixture)
        await expect(
          dispatcher
            .connect(accounts.relayer)
            .connectIbcChannel(
              mars.address,
              C.ChannelIds[0],
              C.ConnHops1,
              C.Ordered,
              C.BscPortId,
              C.RemoteChannelIds[0],
              C.BV1,
              C.InvalidProof
            )
        ).to.be.revertedWith('Fail to prove channel state')
      })

      it('unsupported version', async function () {
        const { dispatcher, mars, accounts } = await loadFixture(setupCoreClientFixture)
        await expect(
          dispatcher
            .connect(accounts.relayer)
            .connectIbcChannel(
              mars.address,
              C.ChannelIds[0],
              C.ConnHops1,
              C.Ordered,
              C.BscPortId,
              C.RemoteChannelIds[0],
              C.InvalidVersion,
              C.ValidProof
            )
        ).to.be.revertedWith('Unsupported version')
      })
    })
  })

  describe('closeIbcChannel', function () {
    it('ChanCloseInit', async function () {
      const { dispatcher, mars, accounts, channel } = await loadFixture(setupChannelFixture)
      await expect(mars.connect(accounts.user1).triggerChannelClose(channel.channelId, dispatcher.address))
        .to.emit(dispatcher, 'CloseIbcChannel')
        .withArgs(channel.portAddress, channel.channelId)
    })

    it('ChanCloseInit: cannot succeed if caller contract does not own the channel', async function () {
      const { dispatcher, channel } = await loadFixture(setupChannelFixture)
      const earth = await (await ethers.getContractFactory('Mars')).deploy()
      await expect(earth.triggerChannelClose(channel.channelId, dispatcher.address)).to.be.revertedWith(
        'Channel not owned by msg.sender'
      )
    })

    it('ChanCloseConfirm', async function () {
      const { dispatcher, accounts, channel } = await loadFixture(setupChannelFixture)
      await expect(
        dispatcher.connect(accounts.relayer).onCloseIbcChannel(channel.portAddress, channel.channelId, C.ValidProof)
      )
        .to.emit(dispatcher, 'CloseIbcChannel')
        .withArgs(channel.portAddress, channel.channelId)
    })

    it('ChanCloseConfirm fails if port does not own the channel', async function () {
      const { dispatcher, accounts, channel } = await loadFixture(setupChannelFixture)
      await expect(
        dispatcher.connect(accounts.relayer).onCloseIbcChannel(channel.portAddress, C.ChannelIds[1], C.ValidProof)
      ).to.be.revertedWith('Channel not owned by portAddress')
    })

    it('ChanCloseConfirm fails if channel proof invalid', async function () {
      const { dispatcher, accounts, channel } = await loadFixture(setupChannelFixture)
      await expect(
        dispatcher.connect(accounts.relayer).onCloseIbcChannel(channel.portAddress, C.ChannelIds[0], C.InvalidProof)
      ).to.be.revertedWith('Fail to prove channel state')
    })
  })

  // generate a packet with a sequence number for testing
  const getPacket = (packet: (typeof C.Packets)[0], sequence: number) => {
    return {
      ...packet,
      msg: `packet.msg-${sequence}`,
      sequence: sequence,
      fee: ethers.utils.parseEther('0.123').mul(sequence + 1)
    }
  }

  // generate a ack packet for testing
  const getAck = (packet: (typeof C.Packets)[0], ackSuccess: boolean) => {
    return { success: ackSuccess, data: toBytes(`$ack-${packet.sequence}-${packet.msg}`) }
  }

  const sendNPacket = async (N: number) => {
    const { dispatcher, mars, accounts, channel, packets } = await loadFixture(setupChannelFixture)
    for (let i = 1; i <= N; i++) {
      const packet = getPacket(packets[0], i)
      await expect(
        mars
          .connect(accounts.user1)
          .greet(dispatcher.address, packet.msg, channel.channelId, packet.timeout, packet.fee, {
            value: packet.fee
          })
      )
        .to.emit(dispatcher, 'SendPacket')
        .withArgs(
          channel.portAddress,
          channel.channelId,
          toBytes(packet.msg),
          packet.sequence,
          packet.timeout,
          packet.fee
        )
    }
    return { accounts, dispatcher, mars, channel, packets }
  }

  describe('recvPacket', function () {
    it('succeeds with a case insensitive address', async function () {
      const { dispatcher, mars, accounts, channel, packets } = await loadFixture(setupChannelFixture)

      const updatedPortPrefix = 'polyibc.eth.'
      await dispatcher.setPortPrefix(updatedPortPrefix)
      const packet = getPacket(packets[0], 0)

      const portId = `${updatedPortPrefix}${mars.address.slice(2)}`

      const ibcPacket = {
        src: {
          portId: C.BscPortId,
          channelId: channel.channelId
        },
        dest: { portId: portId, channelId: C.RemoteChannelIds[0] },
        sequence: 1,
        data: toBytes(packet.msg),
        timeout: { blockHeight: 0, timestamp: packet.timeout }
      }

      const txRecv = await dispatcher.connect(accounts.relayer).recvPacket(mars.address, ibcPacket, C.ValidProof)
      await expect(txRecv)
        .to.emit(dispatcher, 'RecvPacket')
        .withArgs(
          mars.address,
          ibcPacket.dest.channelId,
          ibcPacket.src.portId,
          ibcPacket.src.channelId,
          ibcPacket.sequence
        )
        .to.emit(dispatcher, 'WriteAckPacket')
        .withArgs(mars.address, ibcPacket.dest.channelId, ibcPacket.sequence, [
          true,
          ethers.utils.hexlify(toBytes(`{ "account": "account", "reply": "got the message" }`))
        ])
    })
  })

  describe('sendPacket', function () {
    it('succeeds', async function () {
      const { dispatcher, mars, accounts, channel, packets } = await loadFixture(setupChannelFixture)
      const packet = Object.assign({}, packets[0]) // make a copy
      const escrowBalance = () => accounts.escrow.getBalance().then((b) => b.toBigInt())

      const assertSendPacket = async (packet: (typeof C.Packets)[0]) => {
        const msg = `packet.msg-${packet.sequence}`
        const starttingEscrowBalance = await escrowBalance()
        await expect(
          mars.connect(accounts.user1).greet(dispatcher.address, msg, channel.channelId, packet.timeout, packet.fee, {
            // only fee is escrowed, if msg.value > fee. The overage is lost to miner.
            // So as a dApp dev, you should always set msg.value to the exact packet fee.
            value: packet.fee
          })
        )
          .to.emit(dispatcher, 'SendPacket')
          .withArgs(channel.portAddress, channel.channelId, toBytes(msg), packet.sequence, packet.timeout, packet.fee)
        // confirm Escrow balance changed
        const escrowIncrease = (await escrowBalance()) - starttingEscrowBalance
        expect(escrowIncrease).to.equal(packet.fee.toBigInt())
      }

      for (let i = 1; i <= 3; i++) {
        packet.sequence = i
        packet.fee = ethers.utils.parseEther('0.123').mul(i + 1)
        await assertSendPacket(packet)
      }
    })

    it('fails if tx value < packet fee', async function () {
      const {
        dispatcher,
        mars,
        accounts,
        channel,
        packets: [packet]
      } = await loadFixture(setupChannelFixture)

      await expect(
        mars
          .connect(accounts.user1)
          .greet(dispatcher.address, packet.msg, channel.channelId, packet.timeout, packet.fee, {
            value: packet.fee.sub(1)
          })
      ).to.be.reverted
    })

    it('fails if channel not owned by send dApp', async function () {
      const {
        dispatcher,
        mars,
        accounts,
        packets: [packet]
      } = await loadFixture(setupChannelFixture)

      await expect(
        mars
          .connect(accounts.user1)
          .greet(dispatcher.address, packet.msg, C.RemoteChannelIds[0], packet.timeout, packet.fee, {
            value: packet.fee
          })
      ).to.be.revertedWith('Channel not owned by sender')
    })
  })

  describe('acknowledge', function () {
    it('succeeds only if there is packet commitment and valid proof', async function () {
      const {
        dispatcher,
        mars,
        accounts,
        channel,
        packets: [packetTemplate]
      } = await sendNPacket(4)

      // unordered channel can ack packets in any order
      const assertAck = async (
        sequence: number,
        error?: string,
        setting: { ackError?: boolean; invalidReceiver?: boolean; invalidProof?: boolean } = {}
      ) => {
        const { ackError = false, invalidReceiver = false, invalidProof = false } = setting
        const packet = getPacket(packetTemplate, sequence)
        const ack = getAck(packet, !ackError)
        const srcAddr = invalidReceiver ? accounts.otherUsers[0] : mars
        const srcPortId = `polyibc.eth.${ethers.utils.hexlify(srcAddr.address).slice(2)}`

        const txAck = dispatcher.connect(accounts.relayer).acknowledgement(
          mars.address,
          {
            src: {
              portId: srcPortId,
              channelId: channel.channelId
            },
            dest: { portId: C.BscPortId, channelId: C.RemoteChannelIds[0] },
            sequence: sequence,
            data: toBytes(packet.msg),
            timeout: { blockHeight: 0, timestamp: packet.timeout }
          },
          ack,
          invalidProof ? C.InvalidProof : C.ValidProof
        )
        if (!error) {
          await expect(txAck)
            .to.emit(dispatcher, 'Acknowledgement')
            .withArgs(
              channel.portAddress,
              channel.channelId,
              [ack.success, ethers.utils.hexlify(ack.data)],
              packet.sequence
            )
        } else {
          await expect(txAck).to.be.revertedWith(error)
        }
      }

      await assertAck(3)
      // app-level ack error is still a successful ack at IBC level
      await assertAck(2, '', { ackError: true })
      // processed ackPacket cannot be acked again!
      await assertAck(3, 'Packet commitment not found')
      // cannot ack non-existing packet
      await assertAck(100, 'Packet commitment not found')
      await assertAck(4, 'Receiver is not the original packet sender', { invalidReceiver: true })
      await assertAck(4, 'Fail to prove ack', { invalidProof: true })
      await assertAck(4)
    })
  })

  describe('timeout', function () {
    it('succeeds only if there is packet commitment and valid proof', async function () {
      const {
        dispatcher,
        mars,
        accounts,
        channel,
        packets: [packetTemplate]
      } = await sendNPacket(4)

      // test setPortPrefix logic
      const updatedPortPrefix = 'polyibc.ethereumLongLongPrefix.'
      await expect(dispatcher.connect(accounts.relayer).setPortPrefix(updatedPortPrefix)).to.be.revertedWith(
        'Ownable: caller is not the owner'
      )
      await dispatcher.setPortPrefix(updatedPortPrefix)

      // unordered channel can timeout packets in any order
      const assertTimeout = async (
        sequence: number,
        error?: string,
        setting: { invalidReceiver?: boolean; invalidProof?: boolean } = {}
      ) => {
        const packet = getPacket(packetTemplate, sequence)
        const { invalidReceiver = false, invalidProof = false } = setting
        const srcAddr = invalidReceiver ? accounts.otherUsers[0] : mars
        const srcPortId = `${updatedPortPrefix}${ethers.utils.hexlify(srcAddr.address).slice(2)}`

        const txTimeout = dispatcher.connect(accounts.relayer).timeout(
          mars.address,
          {
            src: {
              portId: srcPortId,
              channelId: channel.channelId
            },
            dest: { portId: C.BscPortId, channelId: C.RemoteChannelIds[0] },
            sequence: sequence,
            data: toBytes(packet.msg),
            timeout: { blockHeight: 0, timestamp: packet.timeout }
          },
          invalidProof ? C.InvalidProof : C.ValidProof
        )
        if (!error) {
          await expect(txTimeout)
            .to.emit(dispatcher, 'Timeout')
            .withArgs(channel.portAddress, channel.channelId, packet.sequence)
        } else {
          await expect(txTimeout).to.be.revertedWith(error)
        }
      }

      await assertTimeout(2)
      // processed timeoutPacket cannot be timed out again!
      await assertTimeout(2, 'Packet commitment not found')
      // cannot timeout non-existing packet
      await assertTimeout(100, 'Packet commitment not found')
      await assertTimeout(3, 'Receiver is not the original packet sender', { invalidReceiver: true })
      await assertTimeout(3, 'Fail to prove timeout', { invalidProof: true })
      await assertTimeout(3)
    })
  })
  // end of tests
})
