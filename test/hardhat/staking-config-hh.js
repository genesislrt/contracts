const { ethers, network, upgrades } = require("hardhat");
const helpers = require("@nomicfoundation/hardhat-network-helpers");
const { expect } = require("chai");
const toBN = ethers.BigNumber.from;

const e18 = new toBN("1000000000000000000");
const zeroAddress = "0x0000000000000000000000000000000000000000";
const _minimumStake = toBN("1000000000");
const _minimumUnstake = toBN("10000000000");
const _DISTRIBUTE_GAS_LIMIT = toBN(100);
const _ratioThreshold = toBN(10_000_000); //10%

let governance, operator, treasury, signer1, signer2, signer3;

const init = async () => {
  [governance, operator, treasury, signer1, signer2, signer3] = await ethers.getSigners();

  console.log('... Start initialization ...');
  const block = await ethers.provider.getBlock("latest");
  console.log(`Starting block number: ${block.number}`);

  console.log(`- TestEigenPodManager`);
  const podManagerFactory = await ethers.getContractFactory("TestEigenPodManager");
  const podManager = await podManagerFactory.deploy();
  console.log(`pod manager address: ${podManager.address}`);

  console.log(`- CertificateToken`);
  const certificateTokenFactory = await ethers.getContractFactory("CertificateToken");
  const certificateToken = await certificateTokenFactory.deploy();
  await certificateToken.deployed();

  console.log(`- StakingConfig`);
  const stakingConfigFactory = await ethers.getContractFactory("StakingConfig");
  const stakingConfig = await upgrades.deployProxy(stakingConfigFactory,
    [_minimumUnstake, _minimumStake, operator.address, governance.address, treasury.address, zeroAddress, certificateToken.address, zeroAddress, podManager.address],
    { initializer: "initialize" });
  await stakingConfig.deployed();
  await certificateToken.initialize(stakingConfig.address, "aETHc", "Certificate token");

  console.log(`- RatioFeed`);
  const ratioFeedFactory = await ethers.getContractFactory("RatioFeed");
  const ratioFeed = await upgrades.deployProxy(ratioFeedFactory,
    [stakingConfig.address], { initializer: "initialize" });
  await ratioFeed.deployed();
  await stakingConfig.setRatioFeedAddress(ratioFeed.address);

  console.log(`- StakingPool`);
  const stakingPoolFactory = await ethers.getContractFactory("StakingPool_V1");
  const stakingPool = await upgrades.deployProxy(stakingPoolFactory,
    [stakingConfig.address, _DISTRIBUTE_GAS_LIMIT], { initializer: "initialize" });
  await stakingPool.deployed();
  await stakingConfig.setStakingPoolAddress(stakingPool.address);

  console.log('... Initialization completed ...');

  await ratioFeed.setRatioThreshold(_ratioThreshold);
  await ratioFeed.connect(operator).updateRatioBatch([certificateToken.address], [e18]);
  await increaseChainTimeForSeconds(60 * 60 * 12 + 1); //+12h
  return [podManager, stakingConfig, ratioFeed, certificateToken, stakingPool]
}

describe("Staking config", function () {
  let podManager, stakingConfig, ratioFeed, certificateToken, stakingPool;

  describe("Operator", function () {
    before(async function () {
      [podManager, stakingConfig, ratioFeed, certificateToken, stakingPool] = await init();
    });

    it("getOperatorAddress()", async function () {
      expect(await stakingConfig.getOperatorAddress()).to.be.eq(operator.address);
    });

    it("setOperatorAddress()", async function () {
      const newAddress = randomAddress();
      const tx = await stakingConfig.setOperatorAddress(newAddress);
      const rec = await tx.wait();
      const events = rec.events?.filter((e) => {
        return e.event === "OperatorAddressChanged";
      });
      expect(events.length).to.be.eq(1);
      expect(events[0].args["prevValue"].toLowerCase()).to.be.eq(operator.address.toLowerCase());
      expect(events[0].args["newValue"].toLowerCase()).to.be.eq(newAddress.toLowerCase());

      expect((await stakingConfig.getOperatorAddress()).toLowerCase()).to.be.eq(newAddress);
    });

    it("setOperatorAddress(): only governance can", async function () {
      await expect(stakingConfig.connect(signer1).setOperatorAddress(randomAddress()))
        .to.be.revertedWith("StakingConfig: only governance");
    });

    it("setOperatorAddress(): reverts: when change to zero address", async function () {
      await expect(stakingConfig.setOperatorAddress(zeroAddress))
        .to.be.revertedWith("StakingConfig: address can't be nil");
    });
  });

  describe("Governance", function () {
    before(async function () {
      [podManager, stakingConfig, ratioFeed, certificateToken, stakingPool] = await init();
    });

    it("getGovernanceAddress()", async function () {
      expect(await stakingConfig.getGovernanceAddress()).to.be.eq(governance.address);
    });

    it("setGovernanceAddress(): only governance can", async function () {
      await expect(stakingConfig.connect(signer1).setGovernanceAddress(randomAddress()))
        .to.be.revertedWith("StakingConfig: only governance");
    });

    it("setGovernanceAddress(): reverts: when change to zero address", async function () {
      await expect(stakingConfig.setGovernanceAddress(zeroAddress))
        .to.be.revertedWith("StakingConfig: address can't be nil");
    });

    it("setGovernanceAddress()", async function () {
      const newAddress = randomAddress();
      const tx = await stakingConfig.setGovernanceAddress(newAddress);
      const rec = await tx.wait();
      const events = rec.events?.filter((e) => {
        return e.event === "GovernanceAddressChanged";
      });
      expect(events.length).to.be.eq(1);
      expect(events[0].args["prevValue"].toLowerCase()).to.be.eq(governance.address.toLowerCase());
      expect(events[0].args["newValue"].toLowerCase()).to.be.eq(newAddress.toLowerCase());

      expect((await stakingConfig.getGovernanceAddress()).toLowerCase()).to.be.eq(newAddress);
    });
  });

  describe("RatioFeed", function () {
    before(async function () {
      [podManager, stakingConfig, ratioFeed, certificateToken, stakingPool] = await init();
    });

    it("getRatioFeedAddress()", async function () {
      expect(await stakingConfig.getRatioFeedAddress()).to.be.eq(ratioFeed.address);
    });

    it("setRatioFeedAddress(): only governance can", async function () {
      await expect(stakingConfig.connect(signer1).setRatioFeedAddress(randomAddress()))
        .to.be.revertedWith("StakingConfig: only governance");
    });

    it("setRatioFeedAddress(): reverts: when change to zero address", async function () {
      await expect(stakingConfig.setRatioFeedAddress(zeroAddress))
        .to.be.revertedWith("StakingConfig: address can't be nil");
    });

    it("setRatioFeedAddress()", async function () {
      const newAddress = randomAddress();
      const tx = await stakingConfig.setRatioFeedAddress(newAddress);
      const rec = await tx.wait();
      const events = rec.events?.filter((e) => {
        return e.event === "RatioFeedAddressChanged";
      });
      expect(events.length).to.be.eq(1);
      expect(events[0].args["prevValue"].toLowerCase()).to.be.eq(ratioFeed.address.toLowerCase());
      expect(events[0].args["newValue"].toLowerCase()).to.be.eq(newAddress.toLowerCase());

      expect((await stakingConfig.getRatioFeedAddress()).toLowerCase()).to.be.eq(newAddress);
    });
  });

  describe("Treasury", function () {
    before(async function () {
      [podManager, stakingConfig, ratioFeed, certificateToken, stakingPool] = await init();
    });

    it("getTreasuryAddress()", async function () {
      expect(await stakingConfig.getTreasuryAddress()).to.be.eq(treasury.address);
    });

    it("setTreasuryAddress(): only governance can", async function () {
      await expect(stakingConfig.connect(signer1).setTreasuryAddress(randomAddress()))
        .to.be.revertedWith("StakingConfig: only governance");
    });

    it("setTreasuryAddress(): reverts: when change to zero address", async function () {
      await expect(stakingConfig.setTreasuryAddress(zeroAddress))
        .to.be.revertedWith("StakingConfig: address can't be nil");
    });

    it("setTreasuryAddress()", async function () {
      const newAddress = randomAddress();
      const tx = await stakingConfig.setTreasuryAddress(newAddress);
      const rec = await tx.wait();
      const events = rec.events?.filter((e) => {
        return e.event === "TreasuryAddressChanged";
      });
      expect(events.length).to.be.eq(1);
      expect(events[0].args["prevValue"].toLowerCase()).to.be.eq(treasury.address.toLowerCase());
      expect(events[0].args["newValue"].toLowerCase()).to.be.eq(newAddress.toLowerCase());

      expect((await stakingConfig.getTreasuryAddress()).toLowerCase()).to.be.eq(newAddress);
    });
  });

  describe("StakingPool", function () {
    before(async function () {
      [podManager, stakingConfig, ratioFeed, certificateToken, stakingPool] = await init();
    });

    it("getStakingPoolAddress()", async function () {
      expect(await stakingConfig.getStakingPoolAddress()).to.be.eq(stakingPool.address);
    });

    it("setStakingPoolAddress(): only governance can", async function () {
      await expect(stakingConfig.connect(signer1).setStakingPoolAddress(randomAddress()))
        .to.be.revertedWith("StakingConfig: only governance");
    });

    it("setStakingPoolAddress(): reverts: when change to zero address", async function () {
      await expect(stakingConfig.setStakingPoolAddress(zeroAddress))
        .to.be.revertedWith("StakingConfig: address can't be nil");
    });

    it("setStakingPoolAddress()", async function () {
      const newAddress = randomAddress();
      const tx = await stakingConfig.setStakingPoolAddress(newAddress);
      const rec = await tx.wait();
      const events = rec.events?.filter((e) => {
        return e.event === "StakingPoolAddressChanged";
      });
      expect(events.length).to.be.eq(1);
      expect(events[0].args["prevValue"].toLowerCase()).to.be.eq(stakingPool.address.toLowerCase());
      expect(events[0].args["newValue"].toLowerCase()).to.be.eq(newAddress.toLowerCase());

      expect((await stakingConfig.getStakingPoolAddress()).toLowerCase()).to.be.eq(newAddress);
    });
  });

  describe("CertToken", function () {
    before(async function () {
      [podManager, stakingConfig, ratioFeed, certificateToken, stakingPool] = await init();
    });

    it("getCertTokenAddress()", async function () {
      expect(await stakingConfig.getCertTokenAddress()).to.be.eq(certificateToken.address);
    });

    it("setCertTokenAddress(): only governance can", async function () {
      await expect(stakingConfig.connect(signer1).setCertTokenAddress(randomAddress()))
        .to.be.revertedWith("StakingConfig: only governance");
    });

    it("setCertTokenAddress(): reverts: when change to zero address", async function () {
      await expect(stakingConfig.setCertTokenAddress(zeroAddress))
        .to.be.revertedWith("StakingConfig: address can't be nil");
    });

    it("setCertTokenAddress()", async function () {
      const newAddress = randomAddress();
      const tx = await stakingConfig.setCertTokenAddress(newAddress);
      const rec = await tx.wait();
      const events = rec.events?.filter((e) => {
        return e.event === "CertTokenAddressChanged";
      });
      expect(events.length).to.be.eq(1);
      expect(events[0].args["prevValue"].toLowerCase()).to.be.eq(certificateToken.address.toLowerCase());
      expect(events[0].args["newValue"].toLowerCase()).to.be.eq(newAddress.toLowerCase());

      expect((await stakingConfig.getCertTokenAddress()).toLowerCase()).to.be.eq(newAddress);
    });



  });

  describe("MinStake", function () {
    before(async function () {
      [podManager, stakingConfig, ratioFeed, certificateToken, stakingPool] = await init();
    });

    it("getMinStake()", async function () {
      expect(await stakingConfig.getMinStake()).to.be.eq(_minimumStake);
    });

    it("setMinStake(): only governance can", async function () {
      await expect(stakingConfig.connect(signer1).setMinStake(_minimumStake.mul(2)))
        .to.be.revertedWith("StakingConfig: only governance");
    });

    it("setMinStake(): reverts: when has a reminder from division by 1gwei", async function () {
      await expect(stakingConfig.setMinStake("1234567890"))
        .to.be.revertedWith("StakingConfig: value should be multiplied of gwei");
    });

    it("setMinStake()", async function () {
      const newAmount = _minimumStake.mul(2);
      const tx = await stakingConfig.setMinStake(newAmount);
      const rec = await tx.wait();
      const events = rec.events?.filter((e) => {
        return e.event === "MinStakeChanged";
      });
      expect(events.length).to.be.eq(1);
      expect(events[0].args["prevValue"]).to.be.eq(_minimumStake);
      expect(events[0].args["newValue"]).to.be.eq(newAmount);

      expect(await stakingConfig.getMinStake()).to.be.eq(newAmount);
    });
  });

  describe("MinUnstake", function () {
    before(async function () {
      [podManager, stakingConfig, ratioFeed, certificateToken, stakingPool] = await init();
    });

    it("getMinUnstake()", async function () {
      expect(await stakingConfig.getMinUnstake()).to.be.eq(_minimumUnstake);
    });

    it("setMinUnstake(): only governance can", async function () {
      await expect(stakingConfig.connect(signer1).setMinUnstake(_minimumUnstake.mul(2)))
        .to.be.revertedWith("StakingConfig: only governance");
    });

    it("setMinUnstake(): reverts: when has a reminder from division by 1gwei", async function () {
      await expect(stakingConfig.setMinUnstake("1234567890"))
        .to.be.revertedWith("StakingConfig: value should be multiplied of gwei");
    });

    it("setMinUnstake()", async function () {
      const newAmount = _minimumUnstake.mul(2);
      const tx = await stakingConfig.setMinUnstake(newAmount);
      const rec = await tx.wait();
      const events = rec.events?.filter((e) => {
        return e.event === "MinUnstakeChanged";
      });
      expect(events.length).to.be.eq(1);
      expect(events[0].args["prevValue"]).to.be.eq(_minimumUnstake);
      expect(events[0].args["newValue"]).to.be.eq(newAmount);

      expect(await stakingConfig.getMinUnstake()).to.be.eq(newAmount);
    });
  });

})


async function updateRatio(ratioFeed, token, ratio) {
  await increaseChainTimeForSeconds(60 * 60 * 12 + 1); //+12h
  await ratioFeed.connect(operator).updateRatioBatch([token.address], [ratio]);
  console.log(`### Ratio updated: ${await token.ratio()}`);
}

function randomBN(length) {
  if (length > 0) {
    let randomNum = "";
    randomNum += Math.floor(Math.random() * 9) + 1; // generates a random digit 1-9
    for (let i = 0; i < length - 1; i++) {
      randomNum += Math.floor(Math.random() * 10); // generates a random digit 0-9
    }
    return toBN(randomNum);
  } else {
    return toBN(0);
  }
}

function randomBNbyMax(max) {
  const maxRandom = 1000_000_000_000;
  if (max.gt(0)) {
    const r = toBN(Math.floor(Math.random() * maxRandom));
    return max.mul(r).div(maxRandom);
  } else {
    return toBN(0)
  }
}

const increaseChainTimeForSeconds = async (seconds) => {
  if (seconds > 0) {
    await ethers.provider.send('evm_increaseTime', [seconds]);
    await mineBlock();
    const nextBlockNumber = await ethers.provider.getBlockNumber();
    const nextBlock = await ethers.provider.getBlock(nextBlockNumber);
    console.log(`Moved to the block #${nextBlockNumber} and timestamp: ${nextBlock.timestamp}`);
    return nextBlock;
  } else {
    return await ethers.provider.getBlock("latest");
  }
}

const mineBlock = async () => {
  network.provider.send("evm_mine");
};

const randomAddress = () => {
  let bigNumber = randomBN(18);
  let hexString = bigNumber.toHexString();
  return  "0x" + hexString.replace("0x", "").padStart(40, '0');
}
