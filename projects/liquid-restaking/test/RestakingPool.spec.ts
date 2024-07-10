import { takeSnapshot, time } from "@nomicfoundation/hardhat-network-helpers";
import { ethers } from "hardhat";
import { expect } from "chai";
import { deployConfig, deployEigenMocks, deployLiquidRestaking, deployRestakerContacts } from "./helpers/deploy";
import {
  CToken,
  ExpensiveStakerMock,
  ProtocolConfig,
  RatioFeed,
  RestakerDeployer,
  RestakingPool,
} from "../typechain-types";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/signers";
import { _1E18, pubkeys, signature, dataRoot } from "./helpers/constants";
import { randomBN, randomBNbyMax } from "./helpers/math";
import { increaseChainTimeForSeconds } from "./helpers/evmutils";
import { SnapshotRestorer } from "@nomicfoundation/hardhat-network-helpers/src/helpers/takeSnapshot";
BigInt.prototype.format = function () {
  return this.toLocaleString("de-DE");
};

const TOKEN_NAME = "Token Name",
  TOKEN_SYMBOL = "Token Symbol",
  TEST_PROVIDER = "TEST_PROVIDER",
  DISTRIBUTE_GAS_LIMIT = 250_000n,
  MIN_UNSTAKE = 10_000_000_000n,
  MIN_STAKE = 1000_000_000n,
  MAX_TVL = 32n * _1E18;

const ceilN = (n: bigint, d: bigint) => n / d + (n % d ? 1n : 0n);

let governance: HardhatEthersSigner,
  operator: HardhatEthersSigner,
  treasury: HardhatEthersSigner,
  signer1: HardhatEthersSigner,
  signer2: HardhatEthersSigner,
  signer3: HardhatEthersSigner;

const init = async () => {
  [governance, operator, treasury, signer1, signer2, signer3] = await ethers.getSigners();

  const protocolConfig = await deployConfig([governance, operator, treasury]);

  // EigenLayr
  const el = await deployEigenMocks();

  // Restaker
  const { restakerDeployer } = await deployRestakerContacts({
    ...el,
    owner: governance.address,
    protocolConfig,
  });

  const { restakingPool, ratioFeed, cToken } = await deployLiquidRestaking({
    protocolConfig,
    tokenName: TOKEN_NAME,
    tokenSymbol: TOKEN_SYMBOL,
    distributeGasLimit: DISTRIBUTE_GAS_LIMIT,
    maxTVL: MAX_TVL,
  });

  return [protocolConfig, restakingPool, cToken, ratioFeed, restakerDeployer];
};

describe("RestakingPool", function () {
  this.timeout(15_000);
  let config: ProtocolConfig,
    pool: RestakingPool,
    cToken: CToken,
    feed: RatioFeed,
    deployer: RestakerDeployer,
    expensiveStaker: ExpensiveStakerMock;
  let snapshot: SnapshotRestorer;
  let MAX_PERCENT, MAX_TARGET_PERCENT;

  before(async function () {
    [config, pool, cToken, feed, deployer] = await init();
    snapshot = await takeSnapshot();
    MAX_PERCENT = await pool.MAX_PERCENT();
    MAX_TARGET_PERCENT = await pool.MAX_TARGET_PERCENT();
  });

  describe("Getters and Setters", function () {
    before(async function () {
      await snapshot.restore();
    });

    it("getMinStake()", async function () {
      expect(await pool.getMinStake()).to.be.eq("1");
    });

    it("getMinUnstake()", async function () {
      expect(await pool.getMinUnstake()).to.be.eq("1");
    });

    it("setDistributeGasLimit()", async function () {
      await expect(pool.setDistributeGasLimit("30"))
        .to.emit(pool, "DistributeGasLimitChanged")
        .withArgs(DISTRIBUTE_GAS_LIMIT, "30");
    });

    it("setDistributeGasLimit(): reverts: only governance can", async function () {
      await expect(pool.connect(operator).setDistributeGasLimit("30")).to.be.revertedWithCustomError(
        pool,
        "OnlyGovernanceAllowed",
      );
    });

    // TODO: check that distribute gas limit cannot be greater than max

    // TODO: set target capacity

    it("setProtocolFee(): sets share of flashWithdrawFee that goes to treasury", async function () {
      const prevValue = await pool.protocolFee();
      const newValue = randomBN(10);
      await expect(pool.setProtocolFee(newValue)).to.emit(pool, "ProtocolFeeChanged").withArgs(prevValue, newValue);
      expect(await pool.protocolFee()).to.be.eq(newValue);
    });

    it("setProtocolFee(): reverts when > MAX_PERCENT", async function () {
      const newValue = (await pool.MAX_PERCENT()) + 1n;
      await expect(pool.setProtocolFee(newValue))
        .to.be.revertedWithCustomError(pool, "ParameterExceedsLimits")
        .withArgs(newValue);
    });

    it("setProtocolFee(): reverts when caller is not an owner", async function () {
      const newValue = randomBN(10);
      await expect(pool.connect(signer1).setProtocolFee(newValue)).to.be.revertedWithCustomError(
        pool,
        "OnlyGovernanceAllowed",
      );
    });
  });

  describe("stake", function () {
    before(async function () {
      await snapshot.restore();
    });

    it("Reverts: when amount > available", async () => {
      const available = await pool.availableToStake();
      expect(available).to.be.eq(MAX_TVL);
      await expect(pool.connect(signer1)["stake()"]({ value: available + 1n })).to.be.revertedWithCustomError(
        pool,
        "PoolStakeAmGreaterThanAvailable",
      );
    });

    const amounts = [
      { name: "Random value", amount: async x => randomBN(19) },
      {
        name: "999999999999999999",
        amount: async x => 999999999999999999n,
      },
      {
        name: "888888888888888888",
        amount: async x => 888888888888888888n,
      },
      {
        name: "777777777777777777",
        amount: async x => 777777777777777777n,
      },
      {
        name: "666666666666666666",
        amount: async x => 666666666666666666n,
      },
      {
        name: "555555555555555555",
        amount: async x => 555555555555555555n,
      },
      {
        name: "444444444444444444",
        amount: async x => 444444444444444444n,
      },
      {
        name: "333333333333333333",
        amount: async x => 333333333333333333n,
      },
      {
        name: "222222222222222222",
        amount: async x => 222222222222222222n,
      },
      {
        name: "111111111111111111",
        amount: async x => 111111111111111111n,
      },
      { name: "Min amount", amount: async x => await x.getMinStake() },
    ];

    amounts.forEach(param => {
      it(`Stake: ${param.name}`, async function () {
        // Update ratio
        const ratio = (await cToken.ratio()) - 1n;
        await updateRatio(feed, cToken, ratio);
        const signerBalanceBefore = await cToken.balanceOf(signer1.address);
        const totalSupplyBefore = await cToken.totalSupply();
        const available = await pool.availableToStake();
        expect(available).to.be.eq(MAX_TVL - ceilN(totalSupplyBefore * _1E18, ratio));

        // Stake
        const amount = await param.amount(pool);
        const expectedShares = (amount * ratio) / _1E18;
        await expect(pool.connect(signer1)["stake()"]({ value: amount }))
          .to.emit(pool, "Staked")
          .withArgs(signer1.address, amount.toString(), expectedShares.toString());

        const signerBalanceAfter = await cToken.balanceOf(signer1.address);
        const totalSupplyAfter = await cToken.totalSupply();

        expect(signerBalanceAfter - signerBalanceBefore).to.be.closeTo(expectedShares, 1);
        expect(totalSupplyAfter - totalSupplyBefore).to.be.closeTo(expectedShares, 1);
      });
    });

    it("Stake with referral code", async function () {
      // Update ratio
      const ratio = (await cToken.ratio()) - 1n;
      await updateRatio(feed, cToken, ratio);
      const signerBalanceBefore = await cToken.balanceOf(signer1.address);
      const totalSupplyBefore = await cToken.totalSupply();
      const available = await pool.availableToStake();
      expect(available).to.be.eq(MAX_TVL - ceilN(totalSupplyBefore * _1E18, ratio));

      // Stake
      const amount = _1E18;
      const expectedShares = (amount * ratio) / _1E18;
      const code = ethers.encodeBytes32String("promo");
      await expect(pool.connect(signer1)["stake(bytes32)"](code, { value: amount }))
        .to.emit(pool, "Staked")
        .withArgs(signer1.address, amount.toString(), expectedShares.toString())
        .and.to.emit(pool, "ReferralStake")
        .withArgs(code);

      const signerBalanceAfter = await cToken.balanceOf(signer1.address);
      const totalSupplyAfter = await cToken.totalSupply();

      expect(signerBalanceAfter - signerBalanceBefore).to.be.closeTo(expectedShares, 1);
      expect(totalSupplyAfter - totalSupplyBefore).to.be.closeTo(expectedShares, 1);
    });

    it("Reverts: when amount < min", async function () {
      const amount = (await pool.getMinStake()) - 1n;
      await expect(pool.connect(signer1)["stake()"]({ value: amount })).to.be.revertedWithCustomError(
        pool,
        "PoolStakeAmLessThanMin",
      );
    });

    it("Reverts: when amount > available", async () => {
      await expect(
        pool.connect(signer1)["stake()"]({ value: (await pool.availableToStake()) + 1n }),
      ).to.be.revertedWithCustomError(pool, "PoolStakeAmGreaterThanAvailable");
    });

    it("Increase max tvl", async () => {
      const newMax = 2_000_000_000_000_000n * _1E18;
      await expect(pool.setMaxTVL(newMax))
        .to.emit(pool, "MaxTVLChanged")
        .withArgs(32n * _1E18, newMax);
    });

    //Stake many times with different ratio values
    it("Stake many times with different signers and ratio", async function () {
      await snapshot.restore();
      const signers = [signer1, signer2, signer3];
      const signersShares = new Map();
      signersShares.set(signer1.address, 0n);
      signersShares.set(signer2.address, 0n);
      signersShares.set(signer3.address, 0n);
      let ratio,
        expectedTotalSupply = 0n,
        expectedPoolBalance = 0n;

      const iterations = 50;
      await pool.setMaxTVL(BigInt(iterations) * 10n * _1E18);
      for (let i = 0; i < iterations; i++) {
        ratio = (await cToken.ratio()) - randomBN(15);
        await updateRatio(feed, cToken, ratio);

        for (const signer of signers) {
          const amount = randomBNbyMax(10n ** 18n - MIN_STAKE + MIN_UNSTAKE);
          expectedPoolBalance = expectedPoolBalance + amount;
          await pool.connect(signer)["stake()"]({ value: amount });
          const shares = (amount * ratio) / _1E18;
          const currentShares = signersShares.get(signer.address);
          signersShares.set(signer.address, currentShares + shares);
          expectedTotalSupply = expectedTotalSupply + shares;
        }
      }
      expect(await cToken.balanceOf(signer1.address)).to.be.closeTo(signersShares.get(signer1.address), 100);
      expect(await cToken.balanceOf(signer2.address)).to.be.closeTo(signersShares.get(signer2.address), 100);
      expect(await cToken.balanceOf(signer3.address)).to.be.closeTo(signersShares.get(signer3.address), 100);
      expect(await cToken.totalSupply()).to.be.closeTo(expectedTotalSupply, 300);
      expect(await ethers.provider.getBalance(await pool.getAddress())).to.be.eq(expectedPoolBalance);
    });
  });

  describe("stake bonus params setter and calculation", function () {
    let localSnapshot;

    const depositBonusSegment = [
      {
        fromUtilization: async () => 0n,
        fromPercent: async () => await pool.maxBonusRate(),
        toUtilization: async () => await pool.stakeUtilizationKink(),
        toPercent: async () => await pool.optimalBonusRate(),
      },
      {
        fromUtilization: async () => await pool.stakeUtilizationKink(),
        fromPercent: async () => await pool.optimalBonusRate(),
        toUtilization: async () => await pool.MAX_PERCENT(),
        toPercent: async () => await pool.optimalBonusRate(),
      },
      {
        fromUtilization: async () => await pool.MAX_PERCENT(),
        fromPercent: async () => 0n,
        toUtilization: async () => ethers.MaxUint256,
        toPercent: async () => 0n,
      },
    ];

    const args = [
      {
        name: "Normal bonus rewards profile > 0",
        newMaxBonusRate: BigInt(2 * 10 ** 8), //2%
        newOptimalBonusRate: BigInt(0.2 * 10 ** 8), //0.2%
        newstakeUtilizationKink: BigInt(25 * 10 ** 8), //25%
      },
      {
        name: "Optimal utilization = 0 => always optimal rate",
        newMaxBonusRate: BigInt(2 * 10 ** 8),
        newOptimalBonusRate: BigInt(10 ** 8), //1%
        newstakeUtilizationKink: 0n,
      },
      {
        name: "Optimal bonus rate = 0",
        newMaxBonusRate: BigInt(2 * 10 ** 8),
        newOptimalBonusRate: 0n,
        newstakeUtilizationKink: BigInt(25 * 10 ** 8),
      },
      {
        name: "Optimal bonus rate = max > 0 => rate is constant over utilization",
        newMaxBonusRate: BigInt(2 * 10 ** 8),
        newOptimalBonusRate: BigInt(2 * 10 ** 8),
        newstakeUtilizationKink: BigInt(25 * 10 ** 8),
      },
      {
        name: "Optimal bonus rate = max = 0 => no bonus",
        newMaxBonusRate: 0n,
        newOptimalBonusRate: 0n,
        newstakeUtilizationKink: BigInt(25 * 10 ** 8),
      },
      //Will fail when OptimalBonusRate > MaxBonusRate
    ];

    const amounts = [
      {
        name: "min amount from 0",
        flashCapacity: targetCapacity => 0n,
        amount: async (targetCapacity) => (await cToken.convertToAmount(await pool.getMinStake())) + 1n,
      },
      {
        name: "1 wei from 0",
        flashCapacity: targetCapacity => 0n,
        amount: async (targetCapacity) => 1n,
      },
      {
        name: "from 0 to 25% of TARGET",
        flashCapacity: targetCapacity => 0n,
        amount: async (targetCapacity) => (targetCapacity * 25n) / 100n,
      },
      {
        name: "from 0 to 25% + 1wei of TARGET",
        flashCapacity: targetCapacity => 0n,
        amount: async (targetCapacity) => (targetCapacity * 25n) / 100n,
      },
      {
        name: "from 25% to 100% of TARGET",
        flashCapacity: targetCapacity => (targetCapacity * 25n) / 100n,
        amount: async (targetCapacity) => (targetCapacity * 75n) / 100n,
      },
      {
        name: "from 0% to 100% of TARGET",
        flashCapacity: targetCapacity => 0n,
        amount: async (targetCapacity) => targetCapacity,
      },
      {
        name: "from 0% to 200% of TARGET",
        flashCapacity: targetCapacity => 0n,
        amount: async (targetCapacity) => targetCapacity * 2n,
      },
    ];

    args.forEach(function (arg) {
      it(`setStakeBonusParams: ${arg.name}`, async function () {
        await snapshot.restore();
        await pool.addRestaker(TEST_PROVIDER);
        await pool.connect(governance).setMaxTVL(64n * _1E18);
        await expect(
          pool.setStakeBonusParams(arg.newMaxBonusRate, arg.newOptimalBonusRate, arg.newstakeUtilizationKink),
        )
          .to.emit(pool, "StakeBonusParamsChanged")
          .withArgs(arg.newMaxBonusRate, arg.newOptimalBonusRate, arg.newstakeUtilizationKink);

        expect(await pool.maxBonusRate()).to.be.eq(arg.newMaxBonusRate);
        expect(await pool.optimalBonusRate()).to.be.eq(arg.newOptimalBonusRate);
        expect(await pool.stakeUtilizationKink()).to.be.eq(arg.newstakeUtilizationKink);
        localSnapshot = await takeSnapshot();
      });

      amounts.forEach(function (amount) {
        it(`calculateDepositBonus for ${amount.name}`, async function () {
          await localSnapshot.restore();
          const batchDeposited = _1E18 * 32n;
          const targetCapacity = _1E18;
          const targetCapacityPercent = (targetCapacity * MAX_TARGET_PERCENT) / (targetCapacity + batchDeposited);
          console.log(`Target capacity percent:\t${targetCapacityPercent}`);
          console.log(`Default capacity percent:\t${await pool.targetCapacity()}`);

          let flashCapacity = amount.flashCapacity(targetCapacity);

          console.log(`available to stake: ${await pool.availableToStake()}`);
          await pool.connect(signer1)["stake()"]({ value: batchDeposited + flashCapacity + 1n });
          await pool.connect(operator).batchDeposit(TEST_PROVIDER, [pubkeys[0]], [signature], [dataRoot]);
          await pool.connect(governance).setTargetFlashCapacity(targetCapacityPercent);
          console.log(`Flash capacity:\t\t${await pool.getFlashCapacity()}`);
          console.log(`Total assets:\t\t${await cToken.totalAssets()}`);

          let _amount = await amount.amount(targetCapacity);
          let depositBonus = 0n;
          while (_amount > 0n) {
            for (const feeFunc of depositBonusSegment) {
              const utilization = (flashCapacity * MAX_PERCENT) / targetCapacity;
              const fromUtilization = await feeFunc.fromUtilization();
              const toUtilization = await feeFunc.toUtilization();
              if (_amount > 0n && fromUtilization <= utilization && utilization < toUtilization) {
                const fromPercent = await feeFunc.fromPercent();
                const toPercent = await feeFunc.toPercent();
                const upperBound = (toUtilization * targetCapacity) / MAX_PERCENT;
                const replenished = upperBound > flashCapacity + _amount ? _amount : upperBound - flashCapacity;
                const slope = ((toPercent - fromPercent) * MAX_PERCENT) / (toUtilization - fromUtilization);
                const bonusPercent = fromPercent + (slope * (flashCapacity + replenished / 2n)) / targetCapacity;
                const bonus = (replenished * bonusPercent) / MAX_PERCENT;
                console.log(`Replenished:\t\t\t${replenished.format()}`);
                console.log(`Bonus percent:\t\t\t${bonusPercent.format()}`);
                console.log(`Bonus:\t\t\t\t\t${bonus.format()}`);
                flashCapacity += replenished;
                _amount -= replenished;
                depositBonus += bonus;
              }
            }
          }
          let contractBonus = await pool.calculateStakeBonus(await amount.amount(targetCapacity));
          console.log(`Expected deposit bonus:\t${depositBonus.format()}`);
          console.log(`Contract deposit bonus:\t${contractBonus.format()}`);
          expect(contractBonus).to.be.closeTo(depositBonus, 1n);
        });
      });
    });

    const invalidArgs = [
      {
        name: "MaxBonusRate > MAX_PERCENT",
        newMaxBonusRate: () => MAX_PERCENT + 1n,
        newOptimalBonusRate: () => BigInt(0.2 * 10 ** 8), //0.2%
        newstakeUtilizationKink: () => BigInt(25 * 10 ** 8),
        customError: "ParameterExceedsLimits",
      },
      {
        name: "OptimalBonusRate > MAX_PERCENT",
        newMaxBonusRate: () => BigInt(2 * 10 ** 8),
        newOptimalBonusRate: () => MAX_PERCENT + 1n,
        newstakeUtilizationKink: () => BigInt(25 * 10 ** 8),
        customError: "ParameterExceedsLimits",
      },
      {
        name: "stakeUtilizationKink > MAX_PERCENT",
        newMaxBonusRate: () => BigInt(2 * 10 ** 8),
        newOptimalBonusRate: () => BigInt(0.2 * 10 ** 8), //0.2%
        newstakeUtilizationKink: () => MAX_PERCENT + 1n,
        customError: "ParameterExceedsLimits",
      },
    ];
    invalidArgs.forEach(function (arg) {
      it(`setStakeBonusParams reverts when ${arg.name}`, async function () {
        await expect(pool.setStakeBonusParams(arg.newMaxBonusRate(), arg.newOptimalBonusRate(), arg.newstakeUtilizationKink()))
            .to.be.revertedWithCustomError(pool, arg.customError);
      });
    });

    it("setDepositBonusParams reverts when caller is not an owner", async function () {
      await expect(pool.connect(signer1).setStakeBonusParams(BigInt(2 * 10 ** 8), BigInt(0.2 * 10 ** 8), BigInt(25 * 10 ** 8)),)
          .to.be.revertedWithCustomError(pool, "OnlyGovernanceAllowed");
    });
  });

  describe("unstake()", function () {
    before(async function () {
      await snapshot.restore();
      await pool.setMaxTVL(200n * _1E18);
    });

    const amounts = [
      {
        name: "Random value to another address",
        shares: async () => randomBN(19),
        receiver: () => signer2,
      },
      {
        name: "999999999999999999",
        shares: async () => 999999999999999999n,
        receiver: () => signer1,
      },
      {
        name: "888888888888888888",
        shares: async () => 888888888888888888n,
        receiver: () => signer1,
      },
      {
        name: "777777777777777777",
        shares: async () => 777777777777777777n,
        receiver: () => signer1,
      },
      {
        name: "666666666666666666",
        shares: async () => 666666666666666666n,
        receiver: () => signer1,
      },
      {
        name: "555555555555555555",
        shares: async () => 555555555555555555n,
        receiver: () => signer1,
      },
      {
        name: "444444444444444444",
        shares: async () => 444444444444444444n,
        receiver: () => signer1,
      },
      {
        name: "333333333333333333",
        shares: async () => 333333333333333333n,
        receiver: () => signer1,
      },
      {
        name: "222222222222222222",
        shares: async () => 222222222222222222n,
        receiver: () => signer1,
      },
      {
        name: "111111111111111111",
        shares: async () => 111111111111111111n,
        receiver: () => signer1,
      },
      {
        name: "Min amount",
        shares: async (x: RestakingPool) => await x.getMinUnstake(),
        receiver: () => signer1,
      },
    ];

    const receiverUnstakesMap = new Map();
    amounts.forEach(param => {
      it(`Unstake: ${param.name}`, async function () {
        // Stake once
        if ((await cToken.balanceOf(signer1.address)) === 0n) {
          await pool.connect(signer1)["stake()"]({ value: 50n * _1E18 });
        }
        //Update ratio
        const ratio = (await cToken.ratio()) - 1000n;
        await updateRatio(feed, cToken, ratio);
        //Values before

        const receiver = param.receiver();
        const ownerBalanceBefore = await cToken.balanceOf(signer1.address);
        const totalSupplyBefore = await cToken.totalSupply();
        const totalPendingUnstakesBefore = await pool.getTotalPendingUnstakes();
        const receiverPendingUnstakesBefore = await pool.getTotalUnstakesOf(receiver.address);

        //Unstake
        const shares = await param.shares(pool);
        const expectedAsset = (shares * _1E18) / ratio + 1n; //Rounding up

        await expect(pool.connect(signer1).unstake(receiver.address, shares))
          .to.emit(pool, "Unstaked")
          .withArgs(signer1.address, receiver.address, expectedAsset, shares);

        const ownerBalanceAfter = await cToken.balanceOf(signer1.address);
        const totalSupplyAfter = await cToken.totalSupply();
        const totalPendingUnstakesAfter = await pool.getTotalPendingUnstakes();
        const receiverPendingUnstakesAfter = await pool.getTotalUnstakesOf(receiver.address);
        const receiverPendingRequests = (await pool.getUnstakesOf(receiver.address)).map(bn => bn.toString());
        if (!receiverUnstakesMap.has(receiver.address)) {
          receiverUnstakesMap.set(receiver.address, [`${receiver.address},${expectedAsset.toString()}`]);
        } else {
          receiverUnstakesMap.get(receiver.address).push(`${receiver.address},${expectedAsset.toString()}`);
        }

        expect(ownerBalanceBefore - ownerBalanceAfter).to.be.closeTo(shares, 1);
        expect(totalSupplyBefore - totalSupplyAfter).to.be.closeTo(shares, 1);
        expect(totalPendingUnstakesAfter - totalPendingUnstakesBefore).to.be.closeTo(expectedAsset, 1);
        expect(receiverPendingUnstakesAfter - receiverPendingUnstakesBefore).to.be.closeTo(expectedAsset, 1);

        expect(receiverPendingRequests).to.include.members(receiverUnstakesMap.get(receiver.address));
      });
    });

    it("Reverts: when shares < min", async function () {
      const ratio = await cToken.ratio();
      const shares = (await pool.getMinUnstake()) - 1n;
      await expect(pool.connect(signer1).unstake(signer1.address, shares)).to.be.revertedWithCustomError(
        pool,
        "PoolUnstakeAmLessThanMin",
      );
    });

    it("Reverts: when exceed user balance", async function () {
      await pool.connect(signer1)["stake()"]({ value: 10n ** 18n });
      const shares = (await cToken.balanceOf(signer1.address)) + 1n;
      await expect(pool.connect(signer1).unstake(signer1.address, shares)).to.be.revertedWithCustomError(
        cToken,
        "ERC20InsufficientBalance",
      );
    });

    it("Reverts: receiver is zero address", async function () {
      await pool.connect(signer1)["stake()"]({ value: 10n ** 18n });
      const shares = await cToken.convertToShares(10n ** 18n);
      await expect(pool.connect(signer1).unstake(ethers.ZeroAddress, shares)).to.be.revertedWithCustomError(
        pool,
        "PoolZeroAddress",
      );
    });

    it("Unstake all", async function () {
      const ratio = await cToken.ratio();

      const shares = await cToken.balanceOf(signer1.address);
      const expectedAsset = (shares * _1E18) / ratio + 1n; //Rounding up
      const totalPendingUnstakesBefore = await pool.getTotalPendingUnstakes();
      const receiverPendingUnstakesBefore = await pool.getTotalUnstakesOf(signer1.address);
      await pool.connect(signer1).unstake(signer1.address, shares);

      const ownerBalanceAfter = await cToken.balanceOf(signer1.address);
      const totalSupplyAfter = await cToken.totalSupply();
      const totalPendingUnstakesAfter = await pool.getTotalPendingUnstakes();
      const receiverPendingUnstakesAfter = await pool.getTotalUnstakesOf(signer1.address);

      expect(ownerBalanceAfter).to.be.eq(0);
      expect(totalSupplyAfter).to.be.eq(0);
      expect(totalPendingUnstakesAfter - totalPendingUnstakesBefore).to.be.closeTo(expectedAsset, 1);
      expect(receiverPendingUnstakesAfter - receiverPendingUnstakesBefore).to.be.closeTo(expectedAsset, 1);
    });
  });

  describe("Unstake fee params setter and calculation", function () {
    let localSnapshot;

    const withdrawFeeSegment = [
      {
        fromUtilization: async () => 0n,
        fromPercent: async () => await pool.maxFlashFeeRate(),
        toUtilization: async () => await pool.unstakeUtilizationKink(),
        toPercent: async () => await pool.optimalUnstakeRate(),
      },
      {
        fromUtilization: async () => await pool.unstakeUtilizationKink(),
        fromPercent: async () => await pool.optimalUnstakeRate(),
        toUtilization: async () => ethers.MaxUint256,
        toPercent: async () => await pool.optimalUnstakeRate(),
      },
    ];

    const args = [
      {
        name: "Normal withdraw fee profile > 0",
        maxFlashFeeRate: BigInt(2 * 10 ** 8), //2%
        optimalUnstakeRate: BigInt(0.2 * 10 ** 8), //0.2%
        unstakeUtilizationKink: BigInt(25 * 10 ** 8),
      },
      {
        name: "Optimal utilization = 0 => always optimal rate",
        maxFlashFeeRate: BigInt(2 * 10 ** 8),
        optimalUnstakeRate: BigInt(10 ** 8), //1%
        unstakeUtilizationKink: 0n,
      },
      {
        name: "Optimal withdraw rate = 0",
        maxFlashFeeRate: BigInt(2 * 10 ** 8),
        optimalUnstakeRate: 0n,
        unstakeUtilizationKink: BigInt(25 * 10 ** 8),
      },
      {
        name: "Optimal withdraw rate = max > 0 => rate is constant over utilization",
        maxFlashFeeRate: BigInt(2 * 10 ** 8),
        optimalUnstakeRate: BigInt(2 * 10 ** 8),
        unstakeUtilizationKink: BigInt(25 * 10 ** 8),
      },
      {
        name: "Optimal withdraw rate = max = 0 => no fee",
        maxFlashFeeRate: 0n,
        optimalUnstakeRate: 0n,
        unstakeUtilizationKink: BigInt(25 * 10 ** 8),
      },
      //Will fail when optimalWithdrawalRate > MaxFlashFeeRate
    ];

    const amounts = [
      {
        name: "from 200% to 0% of TARGET",
        flashCapacity: targetCapacity => targetCapacity * 2n,
        amount: async (targetCapacity) => await pool.getFlashCapacity(),
      },
      {
        name: "from 200% to 100% of TARGET",
        flashCapacity: targetCapacity => targetCapacity * 2n,
        amount: async (targetCapacity) => targetCapacity,
      },
      {
        name: "from 100% to 0% of TARGET",
        flashCapacity: targetCapacity => targetCapacity,
        amount: async (targetCapacity) => await pool.getFlashCapacity(),
      },
      {
        name: "1 wei from 100%",
        flashCapacity: targetCapacity => targetCapacity,
        amount: async (targetCapacity) => 1n,
      },
      {
        name: "min amount from 100%",
        flashCapacity: targetCapacity => targetCapacity,
        amount: async (targetCapacity) => (await cToken.convertToAmount(await pool.getMinUnstake())) + 1n,
      },
      {
        name: "from 100% to 25% of TARGET",
        flashCapacity: targetCapacity => targetCapacity,
        amount: async (targetCapacity) => (targetCapacity * 75n) / 100n,
      },
      {
        name: "from 100% to 25% - 1wei of TARGET",
        flashCapacity: targetCapacity => targetCapacity,
        amount: async (targetCapacity) => (targetCapacity * 75n) / 100n + 1n,
      },
      {
        name: "from 25% to 0% of TARGET",
        flashCapacity: targetCapacity => (targetCapacity * 25n) / 100n,
        amount: async (targetCapacity) => await pool.getFlashCapacity(),
      },
    ];

    args.forEach(function (arg) {
      it(`setFlashWithdrawFeeParams: ${arg.name}`, async function () {
        await snapshot.restore();
        await pool.addRestaker(TEST_PROVIDER);
        await pool.connect(governance).setMaxTVL(64n * _1E18);
        await expect(
            pool.setFlashUnstakeFeeParams(
                arg.maxFlashFeeRate,
                arg.optimalUnstakeRate,
                arg.unstakeUtilizationKink,
            ))
            .to.emit(pool, "UnstakeFeeParamsChanged")
            .withArgs(arg.maxFlashFeeRate, arg.optimalUnstakeRate, arg.unstakeUtilizationKink);

        expect(await pool.maxFlashFeeRate()).to.be.eq(arg.maxFlashFeeRate);
        expect(await pool.optimalUnstakeRate()).to.be.eq(arg.optimalUnstakeRate);
        expect(await pool.unstakeUtilizationKink()).to.be.eq(arg.unstakeUtilizationKink);
        localSnapshot = await takeSnapshot();
      });

      amounts.forEach(function (amount) {
        it(`calculateFlashWithdrawFee for: ${amount.name}`, async function () {
          await localSnapshot.restore();
          const batchDeposited = _1E18 * 32n;
          const targetCapacity = _1E18;
          let flashCapacity = amount.flashCapacity(targetCapacity);
          const targetCapacityPercent = (targetCapacity * MAX_TARGET_PERCENT) / (flashCapacity + batchDeposited);
          console.log(`Target capacity percent:\t${targetCapacityPercent}`);
          console.log(`Default capacity percent:\t${await pool.targetCapacity()}`);


          console.log(`available to stake: ${await pool.availableToStake()}`);
          await pool.connect(signer1)["stake()"]({ value: batchDeposited + flashCapacity + 1n });
          await pool.connect(operator).batchDeposit(TEST_PROVIDER, [pubkeys[0]], [signature], [dataRoot]);
          await pool.connect(governance).setTargetFlashCapacity(targetCapacityPercent);
          console.log(`Flash capacity:\t\t${await pool.getFlashCapacity()}`);
          console.log(`Total assets:\t\t${await cToken.totalAssets()}`);

          let _amount = await amount.amount(targetCapacity);
          let withdrawFee = 0n;
          while (_amount > 1n) {
            for (const feeFunc of withdrawFeeSegment) {
              const utilization = (flashCapacity * MAX_PERCENT) / targetCapacity;
              const fromUtilization = await feeFunc.fromUtilization();
              const toUtilization = await feeFunc.toUtilization();
              if (_amount > 0n && fromUtilization < utilization && utilization <= toUtilization) {
                console.log(`Utilization:\t\t\t${utilization.format()}`);
                const fromPercent = await feeFunc.fromPercent();
                const toPercent = await feeFunc.toPercent();
                const lowerBound = (fromUtilization * targetCapacity) / MAX_PERCENT;
                const replenished = lowerBound > flashCapacity - _amount ? flashCapacity - lowerBound : _amount;
                const slope = ((toPercent - fromPercent) * MAX_PERCENT) / (toUtilization - fromUtilization);
                const withdrawFeePercent =
                    fromPercent + (slope * (flashCapacity - replenished / 2n)) / targetCapacity;
                const fee = (replenished * withdrawFeePercent) / MAX_PERCENT;
                console.log(`Replenished:\t\t\t${replenished.format()}`);
                console.log(`Fee percent:\t\t\t${withdrawFeePercent.format()}`);
                console.log(`Fee:\t\t\t\t\t${fee.format()}`);
                flashCapacity -= replenished;
                _amount -= replenished;
                withdrawFee += fee;
              }
            }
          }
          let contractFee = await pool.calculateFlashUnstakeFee(await amount.amount(targetCapacity));
          console.log(`Expected withdraw fee:\t${withdrawFee.format()}`);
          console.log(`Contract withdraw fee:\t${contractFee.format()}`);
          expect(contractFee).to.be.closeTo(withdrawFee, 1n);
          expect(contractFee).to.be.gt(0n); //flashWithdraw fee is always greater than 0
        });
      });
    });

    const invalidArgs = [
      {
        name: "MaxBonusRate > MAX_PERCENT",
        maxFlashFeeRate: () => MAX_PERCENT + 1n,
        optimalUnstakeRate: () => BigInt(0.2 * 10 ** 8), //0.2%
        unstakeUtilizationKink: () => BigInt(25 * 10 ** 8),
        customError: "ParameterExceedsLimits",
      },
      {
        name: "OptimalBonusRate > MAX_PERCENT",
        maxFlashFeeRate: () => BigInt(2 * 10 ** 8),
        optimalUnstakeRate: () => MAX_PERCENT + 1n,
        unstakeUtilizationKink: () => BigInt(25 * 10 ** 8),
        customError: "ParameterExceedsLimits",
      },
      {
        name: "DepositUtilizationKink > MAX_PERCENT",
        maxFlashFeeRate: () => BigInt(2 * 10 ** 8),
        optimalUnstakeRate: () => BigInt(0.2 * 10 ** 8), //0.2%
        unstakeUtilizationKink: () => MAX_PERCENT + 1n,
        customError: "ParameterExceedsLimits",
      },
    ];
    invalidArgs.forEach(function (arg) {
      it(`setFlashWithdrawFeeParams reverts when ${arg.name}`, async function () {
        await expect(pool.setFlashUnstakeFeeParams(
                arg.maxFlashFeeRate(),
                arg.optimalUnstakeRate(),
                arg.unstakeUtilizationKink()),
        ).to.be.revertedWithCustomError(pool, arg.customError);
      });
    });

    it("calculateFlashWithdrawFee reverts when capacity is not sufficient", async function () {
      await snapshot.restore();
      await pool.connect(signer1)["stake()"]({value: randomBN(19) });
      const capacity = await pool.getFlashCapacity();
      await expect(pool.calculateFlashUnstakeFee(capacity + 1n))
          .to.be.revertedWithCustomError(pool, "InsufficientCapacity")
          .withArgs(capacity);
    });

    it("setFlashWithdrawFeeParams reverts when caller is not an owner", async function () {
      await expect(pool.connect(signer1)
              .setFlashUnstakeFeeParams(BigInt(2 * 10 ** 8), BigInt(0.2 * 10 ** 8), BigInt(25 * 10 ** 8)),
      ).to.be.revertedWith("Ownable: caller is not the owner");
    });
  });

  describe("Deposit", function () {
    before(async function () {
      await snapshot.restore();
      await pool.addRestaker(TEST_PROVIDER);
      await pool.setMaxTVL(200n * _1E18);
    });

    it("Cannot add one provider twice", async () => {
      await expect(pool.addRestaker(TEST_PROVIDER)).to.be.revertedWithCustomError(pool, "PoolRestakerExists");
    });

    it("batchDeposit(): Reverts: when pool balance < 32Eth", async function () {
      await expect(
        pool.connect(operator).batchDeposit(TEST_PROVIDER, [pubkeys[0]], [signature], [dataRoot]),
      ).to.be.revertedWithCustomError(pool, "PoolInsufficientBalance");
    });

    it("batchDeposit()", async function () {
      await pool.connect(signer1)["stake()"]({ value: _1E18 * 33n });
      await expect(pool.connect(operator).batchDeposit(TEST_PROVIDER, [pubkeys[0]], [signature], [dataRoot]))
        .to.emit(pool, "Deposited")
        .withArgs(TEST_PROVIDER, [pubkeys[0]]);
    });

    it("batchDeposit()", async function () {
      await pool.connect(signer2)["stake()"]({ value: _1E18 * 65n });
      await expect(
        pool.connect(operator).batchDeposit(TEST_PROVIDER, pubkeys, [signature, signature], [dataRoot, dataRoot]),
      )
        .to.emit(pool, "Deposited")
        .withArgs(TEST_PROVIDER, pubkeys);
    });

    it("batchDeposit(): Only operator can", async function () {
      await expect(
        pool.connect(governance).batchDeposit(TEST_PROVIDER, [pubkeys[0]], [signature], [dataRoot]),
      ).to.be.revertedWithCustomError(pool, "OnlyOperatorAllowed");
    });

    it("batchDeposit(): provider not exists", async function () {
      await pool.connect(signer2)["stake()"]({ value: _1E18 * 32n });
      await expect(
        pool.connect(operator).batchDeposit("provider", [pubkeys[0]], [signature], [dataRoot]),
      ).to.be.revertedWithCustomError(pool, "PoolRestakerNotExists");
    });
  });

  describe("distribute unstakes and claims", function () {
    before(async function () {
      await snapshot.restore();
      await pool.setMaxTVL(_1E18 * _1E18);
      await pool.setMinStake(0);
      await pool.setMinUnstake(0);
    });

    const difficult = 25n;
    const signers = [() => signer1, () => signer2, () => signer3];

    for (let i = 0n; i < difficult; i++) {
      const signerIndex = Number(i) % signers.length;

      it(`unstake from contract (${i}/${difficult})`, async () => {
        const minStake = await pool.getMinStake();
        const signer = signers[signerIndex]();
        await pool.connect(signer)["stake()"]({ value: minStake * (i + 10n) });
        /// get tokens amount and unstake a bit less
        const tokensAm = (await cToken.balanceOf(signer.address)) - i;
        await time.increase(60 * 60 * 12);
        await feed.connect(operator).updateRatio(await cToken.getAddress(), (await cToken.ratio()) - 100n);
        await pool.connect(signer).unstake(signer, tokensAm - i);
        // check how much is pending
      });
    }

    it("distributeUnstakes", async () => {
      /// get current amounts
      const signer1Expected = await pool.getTotalUnstakesOf(signer1.address);
      const signer2Expected = await pool.getTotalUnstakesOf(signer2.address);
      const signer3Expected = await pool.getTotalUnstakesOf(signer3.address);

      await pool.connect(operator).distributeUnstakes();

      const signer1Distributed = await pool.claimableOf(signer1.address);
      const signer2Distributed = await pool.claimableOf(signer2.address);
      const signer3Distributed = await pool.claimableOf(signer3.address);

      expect(signer1Distributed).to.be.eq(signer1Expected);
      expect(signer2Distributed).to.be.eq(signer2Expected);
      expect(signer3Distributed).to.be.eq(signer3Expected);
    });

    for (let i = 0; i < signers.length; i++) {
      it(`claim unstake of signer${i}`, async () => {
        const claimable = await pool.claimableOf(signers[i]().address);
        // get how much expected
        await expect(pool.claimUnstake(signers[i]()))
          .to.emit(pool, "UnstakeClaimed")
          .withArgs(signers[i]().address, governance.address, claimable);
      });
    }
  });

  describe("claim rewards from restaker", function () {
    before(async function () {
      await snapshot.restore();
      await pool.addRestaker(TEST_PROVIDER);
    });

    it("only operator allowed", async () => {
      await expect(pool.claimRestaker(TEST_PROVIDER, "0")).to.be.revertedWithCustomError(pool, "OnlyOperatorAllowed");
    });

    it("claim without fee", async () => {
      const restakerAddr = await pool.getRestaker(TEST_PROVIDER);
      // send ETH to restaker
      await operator.sendTransaction({
        to: restakerAddr,
        value: _1E18,
      });

      const balanceBefore = await ethers.provider.getBalance(await pool.getAddress());

      await expect(pool.connect(operator).claimRestaker(TEST_PROVIDER, "0"))
        .to.emit(pool, "FeeClaimed")
        .withArgs(restakerAddr, treasury.address, 0, _1E18);
    });

    it("claim with fee", async () => {
      const restakerAddr = await pool.getRestaker(TEST_PROVIDER);
      // send ETH to restaker
      await operator.sendTransaction({
        to: restakerAddr,
        value: _1E18,
      });

      const balanceBefore = await ethers.provider.getBalance(await pool.getAddress());

      await expect(pool.connect(operator).claimRestaker(TEST_PROVIDER, "0"))
        .to.emit(pool, "FeeClaimed")
        .withArgs(restakerAddr, treasury.address, 0, _1E18);
    });

    it("fee is ambiguous", async () => {
      await expect(pool.connect(operator).claimRestaker(TEST_PROVIDER, "1000"))
        .to.be.revertedWithCustomError(pool, "AmbiguousFee")
        .withArgs(0, "1000");
    });
  });
});

async function updateRatio(ratioFeed: RatioFeed, token: CToken, ratio: BigInt) {
  await increaseChainTimeForSeconds(60 * 60 * 12 + 1); //+12h
  await ratioFeed.connect(operator).updateRatio(token, ratio.toString());
  expect(await token.ratio()).to.be.eq(ratio);
}
