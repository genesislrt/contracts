const { ethers, network, upgrades } = require('hardhat');
const helpers = require('@nomicfoundation/hardhat-network-helpers');
const { expect } = require('chai');
const toBN = ethers.BigNumber.from;

const e18 = new toBN('1000000000000000000');
const zeroAddress = '0x0000000000000000000000000000000000000000';
const _minimumStake = toBN('1000000000');
const _minimumUnstake = toBN('10000000000');
const _DISTRIBUTE_GAS_LIMIT = toBN(100);
const _ratioThreshold = toBN(10_000_000); //10%

let governance, operator, treasury, signer1, signer2, signer3;

const init = async () => {
    [governance, operator, treasury, signer1, signer2, signer3] =
        await ethers.getSigners();

    console.log('... Start initialization ...');
    const block = await ethers.provider.getBlock('latest');
    console.log(`Starting block number: ${block.number}`);

    console.log(`- TestEigenPodManager`);
    const podManagerFactory = await ethers.getContractFactory(
        'TestEigenPodManager'
    );
    const podManager = await podManagerFactory.deploy();
    console.log(`pod manager address: ${podManager.address}`);

    console.log(`- CertificateToken`);
    const certificateTokenFactory =
        await ethers.getContractFactory('CertificateToken');
    const certificateToken = await certificateTokenFactory.deploy();
    await certificateToken.deployed();

    console.log(`- StakingConfig`);
    const stakingConfigFactory =
        await ethers.getContractFactory('StakingConfig');
    const stakingConfig = await upgrades.deployProxy(
        stakingConfigFactory,
        [
            _minimumUnstake,
            _minimumStake,
            operator.address,
            governance.address,
            treasury.address,
            zeroAddress,
            certificateToken.address,
            zeroAddress,
            podManager.address,
        ],
        { initializer: 'initialize' }
    );
    await stakingConfig.deployed();
    await certificateToken.initialize(
        stakingConfig.address,
        'aETHc',
        'Certificate token'
    );

    console.log(`- RatioFeed`);
    const ratioFeedFactory = await ethers.getContractFactory('RatioFeed');
    const ratioFeed = await upgrades.deployProxy(
        ratioFeedFactory,
        [stakingConfig.address],
        { initializer: 'initialize' }
    );
    await ratioFeed.deployed();
    await stakingConfig.setRatioFeedAddress(ratioFeed.address);

    console.log(`- StakingPool`);
    const stakingPoolFactory =
        await ethers.getContractFactory('StakingPool_V1');
    const stakingPool = await upgrades.deployProxy(
        stakingPoolFactory,
        [stakingConfig.address, _DISTRIBUTE_GAS_LIMIT],
        { initializer: 'initialize' }
    );
    await stakingPool.deployed();
    await stakingConfig.setStakingPoolAddress(stakingPool.address);

    console.log('... Initialization completed ...');
    return [
        podManager,
        stakingConfig,
        ratioFeed,
        certificateToken,
        stakingPool,
    ];
};

describe('RatioFeed', function () {
    let podManager, stakingConfig, ratioFeed, certificateToken, stakingPool;

    describe('setRatioThreshold()', function () {
        before(async function () {
            [
                podManager,
                stakingConfig,
                ratioFeed,
                certificateToken,
                stakingPool,
            ] = await init();
        });

        it('Changes threshold value', async function () {
            const oldThreshold = await ratioFeed.getRatioThreshold();
            const newThreshold = randomBN(7);
            const tx = await ratioFeed.setRatioThreshold(newThreshold);
            const receipt = await tx.wait();

            const events = receipt.events?.filter((e) => {
                return e.event === 'RatioThresholdChanged';
            });
            expect(events.length).to.be.eq(1);
            expect(events[0].args['oldValue']).to.be.eq(oldThreshold);
            expect(events[0].args['newValue']).to.be.eq(newThreshold);

            expect(await ratioFeed.getRatioThreshold()).to.be.eq(newThreshold);
        });

        it('Changes threshold one more time', async function () {
            const oldThreshold = await ratioFeed.getRatioThreshold();
            const newThreshold = randomBN(7);
            const tx = await ratioFeed.setRatioThreshold(newThreshold);
            const receipt = await tx.wait();

            const events = receipt.events?.filter((e) => {
                return e.event === 'RatioThresholdChanged';
            });
            expect(events.length).to.be.eq(1);
            expect(events[0].args['oldValue']).to.be.eq(oldThreshold);
            expect(events[0].args['newValue']).to.be.eq(newThreshold);

            expect(await ratioFeed.getRatioThreshold()).to.be.eq(newThreshold);
        });

        it('Reverts: only governance can modify', async function () {
            const newThreshold = randomBN(7);
            await expect(
                ratioFeed.connect(signer1).setRatioThreshold(newThreshold)
            ).to.be.revertedWith('RatioFeed: only governance allowed');
        });

        it('Reverts: when threshold >= MAX_THRESHOLD', async function () {
            await expect(
                ratioFeed.setRatioThreshold(await ratioFeed.MAX_THRESHOLD())
            ).to.be.revertedWith('wrong value for ratio threshold');
        });

        it('Reverts: when threshold = 0', async function () {
            await expect(ratioFeed.setRatioThreshold('0')).to.be.revertedWith(
                'wrong value for ratio threshold'
            );
        });
    });

    describe('updateRatioBatch(): publishes ratio for listed addresses', function () {
        before(async function () {
            [
                podManager,
                stakingConfig,
                ratioFeed,
                certificateToken,
                stakingPool,
            ] = await init();
        });

        it('Reverts: when ratio threshold is not set', async function () {
            await expect(
                ratioFeed
                    .connect(operator)
                    .updateRatioBatch(
                        [certificateToken.address],
                        [randomBN(18)]
                    )
            ).to.be.revertedWith('ratio threshold is not set');
        });

        it('Set threshold value', async function () {
            await ratioFeed.setRatioThreshold(randomBN(7));
        });

        it('Publish ratio for the first time', async function () {
            const ratioBefore = await ratioFeed.getRatioFor(
                certificateToken.address
            );
            console.log(`Ratio before: ${ratioBefore}`);
            const newRatio = e18;
            const tx = await ratioFeed
                .connect(operator)
                .updateRatioBatch([certificateToken.address], [newRatio]);
            const receipt = await tx.wait();
            const events = receipt.events?.filter((e) => {
                return e.event === 'RatioUpdated';
            });
            expect(events.length).to.be.eq(1);
            expect(events[0].args['tokenAddress']).to.be.eq(
                certificateToken.address
            );
            expect(events[0].args['oldRatio']).to.be.eq(ratioBefore);
            expect(events[0].args['newRatio']).to.be.eq(newRatio);

            const ratioAfter = await ratioFeed.getRatioFor(
                certificateToken.address
            );
            console.log(`Ratio after: ${ratioAfter}`);
            expect(ratioAfter).to.be.eq(newRatio);
        });

        it('Publish ratio within threshold', async function () {
            await increaseChainTimeForSeconds(60 * 60 * 12 + 1); //+12h
            const ratioBefore = await ratioFeed.getRatioFor(
                certificateToken.address
            );
            console.log(`Ratio before: ${ratioBefore}`);
            const threshold = ratioBefore
                .mul(await ratioFeed.getRatioThreshold())
                .div(await ratioFeed.MAX_THRESHOLD());
            const newRatio = ratioBefore.sub(randomBNbyMax(threshold));
            const tx = await ratioFeed
                .connect(operator)
                .updateRatioBatch([certificateToken.address], [newRatio]);
            const receipt = await tx.wait();
            const events = receipt.events?.filter((e) => {
                return e.event === 'RatioUpdated';
            });
            expect(events.length).to.be.eq(1);
            expect(events[0].args['tokenAddress']).to.be.eq(
                certificateToken.address
            );
            expect(events[0].args['oldRatio']).to.be.eq(ratioBefore);
            expect(events[0].args['newRatio']).to.be.eq(newRatio);

            const ratioAfter = await ratioFeed.getRatioFor(
                certificateToken.address
            );
            console.log(`Ratio after: ${ratioAfter}`);
            expect(ratioAfter).to.be.eq(newRatio);
        });

        it('Publish the same ratio', async function () {
            await increaseChainTimeForSeconds(60 * 60 * 12 + 1); //+12h
            const ratioBefore = await ratioFeed.getRatioFor(
                certificateToken.address
            );
            const tx = await ratioFeed
                .connect(operator)
                .updateRatioBatch([certificateToken.address], [ratioBefore]);
            const receipt = await tx.wait();
            const events = receipt.events?.filter((e) => {
                return e.event === 'RatioUpdated';
            });
            expect(events.length).to.be.eq(1);
            expect(events[0].args['tokenAddress']).to.be.eq(
                certificateToken.address
            );
            expect(events[0].args['oldRatio']).to.be.eq(ratioBefore);
            expect(events[0].args['newRatio']).to.be.eq(ratioBefore);

            expect(
                await ratioFeed.getRatioFor(certificateToken.address)
            ).to.be.eq(ratioBefore);
        });

        it('Skips: when publish second time in less than 12h', async function () {
            await increaseChainTimeForSeconds(60 * 60 * 11.5);
            const ratioBefore = await ratioFeed.getRatioFor(
                certificateToken.address
            );
            const newRatio = ratioBefore.sub(1);
            const tx = await ratioFeed
                .connect(operator)
                .updateRatioBatch([certificateToken.address], [newRatio]);
            const rec = await tx.wait();
            const events = rec.events?.filter((e) => {
                return e.event === 'RatioNotUpdated';
            });
            expect(events.length).to.be.eq(1);
            expect(events[0].args['tokenAddress'].toLowerCase()).to.be.eq(
                certificateToken.address.toLowerCase()
            );
            expect(events[0].args['failedRatio']).to.be.eq(newRatio);
            expect(events[0].args['reason']).to.be.eq(
                'ratio was updated less than 12 hours ago'
            );
            expect(
                await ratioFeed.getRatioFor(certificateToken.address)
            ).to.be.eq(ratioBefore);
        });

        it('Skips: when new ratio > current ratio', async function () {
            await increaseChainTimeForSeconds(60 * 60 * 13);
            const ratioBefore = await ratioFeed.getRatioFor(
                certificateToken.address
            );
            const newRatio = ratioBefore.add(1);
            const tx = await ratioFeed
                .connect(operator)
                .updateRatioBatch([certificateToken.address], [newRatio]);
            const rec = await tx.wait();
            const events = rec.events?.filter((e) => {
                return e.event === 'RatioNotUpdated';
            });
            expect(events.length).to.be.eq(1);
            expect(events[0].args['tokenAddress'].toLowerCase()).to.be.eq(
                certificateToken.address.toLowerCase()
            );
            expect(events[0].args['failedRatio']).to.be.eq(newRatio);
            expect(events[0].args['reason']).to.be.eq(
                'new ratio cannot be greater than old'
            );
            expect(
                await ratioFeed.getRatioFor(certificateToken.address)
            ).to.be.eq(ratioBefore);
        });

        it('Skips: when current - new ratio > threshold', async function () {
            await increaseChainTimeForSeconds(60 * 60 * 13);
            const ratioBefore = await ratioFeed.getRatioFor(
                certificateToken.address
            );
            const threshold = ratioBefore
                .mul(await ratioFeed.getRatioThreshold())
                .div(await ratioFeed.MAX_THRESHOLD());
            const newRatio = ratioBefore.sub(threshold.add(1));
            const tx = await ratioFeed
                .connect(operator)
                .updateRatioBatch([certificateToken.address], [newRatio]);
            const rec = await tx.wait();
            const events = rec.events?.filter((e) => {
                return e.event === 'RatioNotUpdated';
            });
            expect(events.length).to.be.eq(1);
            expect(events[0].args['tokenAddress'].toLowerCase()).to.be.eq(
                certificateToken.address.toLowerCase()
            );
            expect(events[0].args['failedRatio']).to.be.eq(newRatio);
            expect(events[0].args['reason']).to.be.eq(
                'new ratio too low, not in threshold range'
            );
            expect(
                await ratioFeed.getRatioFor(certificateToken.address)
            ).to.be.eq(ratioBefore);
        });

        it('Reverts: only operator can', async function () {
            const ratioBefore = await ratioFeed.getRatioFor(
                certificateToken.address
            );
            const newRatio = ratioBefore.sub(1);
            await expect(
                ratioFeed
                    .connect(signer1)
                    .updateRatioBatch([certificateToken.address], [newRatio])
            ).to.be.revertedWith('RatioFeed: only operator allowed');
        });
    });

    describe('repairRatioFor(): sets specified value', function () {
        before(async function () {
            [
                podManager,
                stakingConfig,
                ratioFeed,
                certificateToken,
                stakingPool,
            ] = await init();

            await ratioFeed.setRatioThreshold('10000000');
            await ratioFeed
                .connect(operator)
                .updateRatioBatch([certificateToken.address], [e18]);
        });

        it('Decrease ratio', async function () {
            const ratioBefore = await ratioFeed.getRatioFor(
                certificateToken.address
            );
            console.log(`Ratio before: ${ratioBefore}`);
            const newRatio = ratioBefore.sub(1);
            const tx = await ratioFeed.repairRatioFor(
                certificateToken.address,
                newRatio
            );
            const receipt = await tx.wait();
            const events = receipt.events?.filter((e) => {
                return e.event === 'RatioUpdated';
            });
            expect(events.length).to.be.eq(1);
            expect(events[0].args['tokenAddress']).to.be.eq(
                certificateToken.address
            );
            expect(events[0].args['oldRatio']).to.be.eq(ratioBefore);
            expect(events[0].args['newRatio']).to.be.eq(newRatio);

            const ratioAfter = await ratioFeed.getRatioFor(
                certificateToken.address
            );
            console.log(`Ratio after: ${ratioAfter}`);
            expect(ratioAfter).to.be.eq(newRatio);
        });

        it('Increase ratio', async function () {
            const ratioBefore = await ratioFeed.getRatioFor(
                certificateToken.address
            );
            console.log(`Ratio before: ${ratioBefore}`);
            const newRatio = ratioBefore.add(1);
            const tx = await ratioFeed.repairRatioFor(
                certificateToken.address,
                newRatio
            );
            const receipt = await tx.wait();
            const events = receipt.events?.filter((e) => {
                return e.event === 'RatioUpdated';
            });
            expect(events.length).to.be.eq(1);
            expect(events[0].args['tokenAddress']).to.be.eq(
                certificateToken.address
            );
            expect(events[0].args['oldRatio']).to.be.eq(ratioBefore);
            expect(events[0].args['newRatio']).to.be.eq(newRatio);

            const ratioAfter = await ratioFeed.getRatioFor(
                certificateToken.address
            );
            console.log(`Ratio after: ${ratioAfter}`);
            expect(ratioAfter).to.be.eq(newRatio);
        });

        it('Reverts: when new ratio = 0', async function () {
            await expect(
                ratioFeed.repairRatioFor(certificateToken.address, '0')
            ).to.be.revertedWith('ratio is zero');
        });

        it('Reverts: only governance can', async function () {
            await expect(
                ratioFeed
                    .connect(signer1)
                    .repairRatioFor(certificateToken.address, randomBN(18))
            ).to.be.revertedWith('RatioFeed: only governance allowed');
        });
    });

    //TODO: when formula will be confirmed (Divisor must be newRatio)
    describe('averagePercentageRate()', function () {
        before(async function () {
            [
                podManager,
                stakingConfig,
                ratioFeed,
                certificateToken,
                stakingPool,
            ] = await init();

            await ratioFeed.setRatioThreshold('1000000');
            await ratioFeed
                .connect(operator)
                .updateRatioBatch([certificateToken.address], [e18]);
        });

        it('Publish ratio for 8days', async function () {
            for (let i = 1; i < 8; i++) {
                const ts = await increaseChainTimeForSeconds(60 * 60 * 24); //12h
                const ratioBefore = await ratioFeed.getRatioFor(
                    certificateToken.address
                );
                // const newRatio = ratioBefore.sub(randomBNbyMax(toBN(1000000)));
                const newRatio = ratioBefore.sub(toBN('27396509684666').mul(i));
                console.log(`${i}. Ratio:\t${ts}\t${ratioBefore}`);
                await ratioFeed
                    .connect(operator)
                    .updateRatioBatch([certificateToken.address], [newRatio]);
            }
        });

        it('Get apr', async function () {
            for (let i = 1; i < 8; i++) {
                const apr = await ratioFeed.averagePercentageRate(
                    certificateToken.address,
                    i
                );
                const aprPercent =
                    apr.div(toBN('1000000000')).toNumber() / 1000000000;
                console.log(`${i}. APR: ${aprPercent}%`);
                //APR: 0.990287047%
            }
        });
    });
});

function randomBN(length) {
    if (length > 0) {
        let randomNum = '';
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
        return toBN(0);
    }
}

const increaseChainTimeForSeconds = async (seconds) => {
    if (seconds > 0) {
        await ethers.provider.send('evm_increaseTime', [seconds]);
        await mineBlock();
        const nextBlockNumber = await ethers.provider.getBlockNumber();
        const nextBlock = await ethers.provider.getBlock(nextBlockNumber);
        console.log(
            `Moved to the block #${nextBlockNumber} and timestamp: ${nextBlock.timestamp}`
        );
        return nextBlock;
    } else {
        return await ethers.provider.getBlock('latest');
    }
};

const mineBlock = async () => {
    network.provider.send('evm_mine');
};
