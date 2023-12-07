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

    await ratioFeed.setRatioThreshold(_ratioThreshold);
    await ratioFeed
        .connect(operator)
        .updateRatioBatch([certificateToken.address], [e18]);
    await increaseChainTimeForSeconds(60 * 60 * 12 + 1); //+12h
    return [
        podManager,
        stakingConfig,
        ratioFeed,
        certificateToken,
        stakingPool,
    ];
};

describe('Staking pool', function () {
    this.timeout(150000);

    let podManager, stakingConfig, ratioFeed, certificateToken, stakingPool;

    describe('Getters and Setters', function () {
        before(async function () {
            [
                podManager,
                stakingConfig,
                ratioFeed,
                certificateToken,
                stakingPool,
            ] = await init();
        });

        it('getCert()', async function () {
            expect(await stakingPool.getCert()).to.be.eq(
                certificateToken.address
            );
        });

        it('getEigenPodManager()', async function () {
            expect(await stakingPool.getEigenPodManager()).to.be.eq(
                podManager.address
            );
        });

        it('getMinStake()', async function () {
            expect(await stakingPool.getMinStake()).to.be.eq(_minimumStake);
        });

        it('getMinUnstake()', async function () {
            expect(await stakingPool.getMinUnstake()).to.be.eq(_minimumUnstake);
        });

        it('getDistributeGasLimit()', async function () {
            expect(await stakingPool.getDistributeGasLimit()).to.be.eq(
                _DISTRIBUTE_GAS_LIMIT
            );
        });

        it('setDistributeGasLimit()', async function () {
            const oldValue = await stakingPool.getDistributeGasLimit();
            const newValue = randomBN(3);
            const tx = await stakingPool.setDistributeGasLimit(newValue);
            const rec = await tx.wait();
            const events = rec.events?.filter((e) => {
                return e.event === 'DistributeGasLimitChanged';
            });
            expect(events.length).to.be.eq(1);
            expect(events[0].args['prevValue']).to.be.eq(oldValue);
            expect(events[0].args['newValue']).to.be.eq(newValue);

            expect(await stakingPool.getDistributeGasLimit()).to.be.eq(
                newValue
            );
        });

        it('setDistributeGasLimit(): reverts: only governance can', async function () {
            const newValue = randomBN(3);
            await expect(
                stakingPool.connect(operator).setDistributeGasLimit(newValue)
            ).to.be.revertedWith('StakingPool: only governance allowed');
        });
    });

    describe('stakeCerts()', function () {
        before(async function () {
            [
                podManager,
                stakingConfig,
                ratioFeed,
                certificateToken,
                stakingPool,
            ] = await init();
        });

        const amounts = [
            { name: 'Random value', amount: async (x) => randomBN(19) },
            {
                name: '999999999999999999',
                amount: async (x) => toBN('999999999999999999'),
            },
            {
                name: '888888888888888888',
                amount: async (x) => toBN('888888888888888888'),
            },
            {
                name: '777777777777777777',
                amount: async (x) => toBN('777777777777777777'),
            },
            {
                name: '666666666666666666',
                amount: async (x) => toBN('666666666666666666'),
            },
            {
                name: '555555555555555555',
                amount: async (x) => toBN('555555555555555555'),
            },
            {
                name: '444444444444444444',
                amount: async (x) => toBN('444444444444444444'),
            },
            {
                name: '333333333333333333',
                amount: async (x) => toBN('333333333333333333'),
            },
            {
                name: '222222222222222222',
                amount: async (x) => toBN('222222222222222222'),
            },
            {
                name: '111111111111111111',
                amount: async (x) => toBN('111111111111111111'),
            },
            { name: 'Min amount', amount: async (x) => await x.getMinStake() },
        ];

        amounts.forEach((param) => {
            it(`Stake: ${param.name}`, async function () {
                //Update ratio
                const ratio = (await certificateToken.ratio()).sub(1);
                await updateRatio(ratioFeed, certificateToken, ratio);
                const signerBalanceBefore = await certificateToken.balanceOf(
                    signer1.address
                );
                const totalSupplyBefore = await certificateToken.totalSupply();

                //Stake
                const amount = await param.amount(stakingPool);
                const expectedShares = amount.mul(ratio).div(e18);
                const tx = await stakingPool
                    .connect(signer1)
                    .stakeCerts({ value: amount });
                const rec = await tx.wait();
                const events = rec.events?.filter((e) => {
                    return e.event === 'Staked';
                });
                expect(events.length).to.be.eq(1);
                expect(events[0].args['staker']).to.be.eq(signer1.address);
                expect(events[0].args['amount']).to.be.eq(amount);
                expect(events[0].args['shares']).to.be.closeTo(
                    expectedShares,
                    1
                );

                const signerBalanceAfter = await certificateToken.balanceOf(
                    signer1.address
                );
                const totalSupplyAfter = await certificateToken.totalSupply();

                expect(
                    signerBalanceAfter.sub(signerBalanceBefore)
                ).to.be.closeTo(expectedShares, 1);
                expect(totalSupplyAfter.sub(totalSupplyBefore)).to.be.closeTo(
                    expectedShares,
                    1
                );
            });
        });

        it('Reverts: when amount < min', async function () {
            const amount = (await stakingPool.getMinStake()).sub(1);
            await expect(
                stakingPool.connect(signer1).stakeCerts({ value: amount })
            ).to.be.revertedWith(
                'StakingPool: value must be greater than min amount'
            );
        });

        //Stake many times with different ratio values
        it('Stake many times with different signers and ratio', async function () {
            await hhReset();
            [
                podManager,
                stakingConfig,
                ratioFeed,
                certificateToken,
                stakingPool,
            ] = await init();
            const signers = [signer1, signer2, signer3];
            const signersShares = new Map();
            signersShares.set(signer1.address, toBN(0));
            signersShares.set(signer2.address, toBN(0));
            signersShares.set(signer3.address, toBN(0));
            let ratio,
                expectedTotalSupply = toBN(0),
                expectedPoolBalance = toBN(0);

            const iterations = 50;
            for (let i = 0; i < iterations; i++) {
                ratio = (await certificateToken.ratio()).sub(randomBN(15));
                await updateRatio(ratioFeed, certificateToken, ratio);

                for (const signer of signers) {
                    const amount = randomBNbyMax(e18.sub(_minimumStake)).add(
                        _minimumUnstake
                    );
                    expectedPoolBalance = expectedPoolBalance.add(amount);
                    await stakingPool
                        .connect(signer)
                        .stakeCerts({ value: amount });
                    const shares = amount.mul(ratio).div(e18);
                    const currentShares = signersShares.get(signer.address);
                    signersShares.set(
                        signer.address,
                        currentShares.add(shares)
                    );
                    expectedTotalSupply = expectedTotalSupply.add(shares);
                }
            }
            expect(
                await certificateToken.balanceOf(signer1.address)
            ).to.be.closeTo(signersShares.get(signer1.address), 100);
            expect(
                await certificateToken.balanceOf(signer2.address)
            ).to.be.closeTo(signersShares.get(signer2.address), 100);
            expect(
                await certificateToken.balanceOf(signer3.address)
            ).to.be.closeTo(signersShares.get(signer3.address), 100);
            expect(await certificateToken.totalSupply()).to.be.closeTo(
                expectedTotalSupply,
                300
            );
            expect(await stakingPool.getFreeBalance()).to.be.eq(
                expectedPoolBalance
            );
        });
    });

    describe('unstakeCerts()', function () {
        before(async function () {
            [
                podManager,
                stakingConfig,
                ratioFeed,
                certificateToken,
                stakingPool,
            ] = await init();
            await stakingPool.setDistributeGasLimit('23000');
        });

        const amounts = [
            {
                name: 'Random value to another address',
                shares: async (x) => randomBN(19),
                receiver: () => signer2,
            },
            {
                name: '999999999999999999',
                shares: async (x) => toBN('999999999999999999'),
                receiver: () => signer1,
            },
            {
                name: '888888888888888888',
                shares: async (x) => toBN('888888888888888888'),
                receiver: () => signer1,
            },
            {
                name: '777777777777777777',
                shares: async (x) => toBN('777777777777777777'),
                receiver: () => signer1,
            },
            {
                name: '666666666666666666',
                shares: async (x) => toBN('666666666666666666'),
                receiver: () => signer1,
            },
            {
                name: '555555555555555555',
                shares: async (x) => toBN('555555555555555555'),
                receiver: () => signer2,
            },
            {
                name: '444444444444444444',
                shares: async (x) => toBN('444444444444444444'),
                receiver: () => signer1,
            },
            {
                name: '333333333333333333',
                shares: async (x) => toBN('333333333333333333'),
                receiver: () => signer1,
            },
            {
                name: '222222222222222222',
                shares: async (x) => toBN('222222222222222222'),
                receiver: () => signer1,
            },
            {
                name: '111111111111111111',
                shares: async (x) => toBN('111111111111111111'),
                receiver: () => signer1,
            },
            {
                name: 'Min amount',
                shares: async (x) => await x.getMinUnstake(),
                receiver: () => signer1,
            },
        ];

        const receiverUnstakesMap = new Map();
        amounts.forEach((param) => {
            it(`Unstake: ${param.name}`, async function () {
                //Stake once
                if ((await certificateToken.balanceOf(signer1.address)).eq(0)) {
                    await stakingPool
                        .connect(signer1)
                        .stakeCerts({ value: e18.mul(50) });
                }
                //Update ratio
                const ratio = await certificateToken.ratio();
                await updateRatio(ratioFeed, certificateToken, ratio);
                //Values before
                const receiver = param.receiver();
                const ownerBalanceBefore = await certificateToken.balanceOf(
                    signer1.address
                );
                const totalSupplyBefore = await certificateToken.totalSupply();
                const totalPendingUnstakesBefore =
                    await stakingPool.getTotalPendingUnstakes();
                const receiverPendingUnstakesBefore =
                    await stakingPool.getPendingUnstakesOf(receiver.address);

                //Unstake
                const shares = await param.shares(stakingPool);
                const expectedAsset = shares.mul(e18).div(ratio);
                const tx = await stakingPool
                    .connect(signer1)
                    .unstakeCerts(receiver.address, shares);
                const rec = await tx.wait();
                const events = rec.events?.filter((e) => {
                    return e.event === 'PendingUnstake';
                });
                expect(events.length).to.be.eq(1);
                expect(events[0].args['ownerAddress']).to.be.eq(
                    signer1.address
                );
                expect(events[0].args['receiverAddress']).to.be.eq(
                    receiver.address
                );
                expect(events[0].args['amount']).to.be.closeTo(
                    expectedAsset,
                    1
                );
                expect(events[0].args['shares']).to.be.eq(shares);

                const ownerBalanceAfter = await certificateToken.balanceOf(
                    signer1.address
                );
                const totalSupplyAfter = await certificateToken.totalSupply();
                const totalPendingUnstakesAfter =
                    await stakingPool.getTotalPendingUnstakes();
                const receiverPendingUnstakesAfter =
                    await stakingPool.getPendingUnstakesOf(receiver.address);
                const receiverPendingRequests = (
                    await stakingPool.getPendingRequestsOf(receiver.address)
                ).map((bn) => bn.toString());
                if (!receiverUnstakesMap.has(receiver.address)) {
                    receiverUnstakesMap.set(receiver.address, [
                        expectedAsset.toString(),
                    ]);
                } else {
                    receiverUnstakesMap
                        .get(receiver.address)
                        .push(expectedAsset.toString());
                }

                expect(ownerBalanceBefore.sub(ownerBalanceAfter)).to.be.closeTo(
                    shares,
                    1
                );
                expect(totalSupplyBefore.sub(totalSupplyAfter)).to.be.closeTo(
                    shares,
                    1
                );
                expect(
                    totalPendingUnstakesAfter.sub(totalPendingUnstakesBefore)
                ).to.be.closeTo(expectedAsset, 1);
                expect(
                    receiverPendingUnstakesAfter.sub(
                        receiverPendingUnstakesBefore
                    )
                ).to.be.closeTo(expectedAsset, 1);

                console.log(
                    `getPendingRequestsOf: ${await stakingPool.getPendingRequestsOf(
                        receiver.address
                    )}`
                );
                expect(receiverPendingRequests).to.include.members(
                    receiverUnstakesMap.get(receiver.address)
                );
            });
        });

        it('Reverts: when shares < min', async function () {
            const ratio = await certificateToken.ratio();
            const shares = (await stakingPool.getMinUnstake())
                .mul(ratio)
                .div(e18)
                .sub(1);
            await expect(
                stakingPool
                    .connect(signer1)
                    .unstakeCerts(signer1.address, shares)
            ).to.be.revertedWith(
                'StakingPool: value must be greater than min amount'
            );
        });

        it('Reverts: when exceed user balance', async function () {
            await stakingPool.connect(signer1).stakeCerts({ value: e18 });
            const shares = (
                await certificateToken.balanceOf(signer1.address)
            ).add(1);
            await expect(
                stakingPool
                    .connect(signer1)
                    .unstakeCerts(signer1.address, shares)
            ).to.be.revertedWith(
                'StakingPool: cannot unstake more than have on address'
            );
        });

        it('Reverts: receiver is zero address', async function () {
            await stakingPool.connect(signer1).stakeCerts({ value: e18 });
            const shares = await certificateToken.bondsToShares(e18);
            await expect(
                stakingPool.connect(signer1).unstakeCerts(zeroAddress, shares)
            ).to.be.revertedWith('LiquidTokenStakingPool: zero input values');
        });

        it('Unstake all', async function () {
            const ratio = await certificateToken.ratio();

            const shares = await certificateToken.balanceOf(signer1.address);
            const expectedAsset = shares.mul(e18).div(ratio).add(1); //Rounding up
            const totalPendingUnstakesBefore =
                await stakingPool.getTotalPendingUnstakes();
            const receiverPendingUnstakesBefore =
                await stakingPool.getPendingUnstakesOf(signer1.address);
            await stakingPool
                .connect(signer1)
                .unstakeCerts(signer1.address, shares);
            console.log(
                `getPendingRequestsOf: ${await stakingPool.getPendingRequestsOf(
                    signer1.address
                )}`
            );

            const ownerBalanceAfter = await certificateToken.balanceOf(
                signer1.address
            );
            const totalSupplyAfter = await certificateToken.totalSupply();
            const totalPendingUnstakesAfter =
                await stakingPool.getTotalPendingUnstakes();
            const receiverPendingUnstakesAfter =
                await stakingPool.getPendingUnstakesOf(signer1.address);

            expect(ownerBalanceAfter).to.be.eq(0);
            expect(totalSupplyAfter).to.be.eq(0);
            expect(
                totalPendingUnstakesAfter.sub(totalPendingUnstakesBefore)
            ).to.be.closeTo(expectedAsset, 1);
            expect(
                receiverPendingUnstakesAfter.sub(receiverPendingUnstakesBefore)
            ).to.be.closeTo(expectedAsset, 1);
        });

        it('distributePendingRewards(): when gas is not enough', async function () {
            const totalPendingUnstakesBefore =
                await stakingPool.getTotalPendingUnstakes();
            console.log(
                `Pending unstakes before: ${totalPendingUnstakesBefore}`
            );
            await stakingPool.distributePendingRewards({ gasLimit: 50_000 });
            const totalPendingUnstakesAfter =
                await stakingPool.getTotalPendingUnstakes();
            console.log(`Pending unstakes after: ${totalPendingUnstakesAfter}`);
        });

        it('distributePendingRewards()', async function () {
            console.log(`Ratio: ${await certificateToken.ratio()}`);
            const poolBalanceBefore = await stakingPool.getFreeBalance();
            console.log(`Pool free balance before: ${poolBalanceBefore}`);
            const signer1BalanceBefore = await ethers.provider.getBalance(
                signer1.address
            );
            const signer2BalanceBefore = await ethers.provider.getBalance(
                signer2.address
            );
            const signer1PendingUnstakesBefore =
                await stakingPool.getPendingUnstakesOf(signer1.address);
            console.log(
                `Signer1 pending unstakes before: ${signer1PendingUnstakesBefore}`
            );
            const signer2PendingUnstakesBefore =
                await stakingPool.getPendingUnstakesOf(signer2.address);
            console.log(
                `Signer2 pending unstakes before: ${signer2PendingUnstakesBefore}`
            );
            const totalPendingUnstakesBefore =
                await stakingPool.getTotalPendingUnstakes();

            //Distribute
            await stakingPool.connect(signer3).distributePendingRewards();

            const signer1BalanceAfter = await ethers.provider.getBalance(
                signer1.address
            );
            const signer2BalanceAfter = await ethers.provider.getBalance(
                signer2.address
            );
            const totalPendingUnstakesAfter =
                await stakingPool.getTotalPendingUnstakes();
            const signer1PendingUnstakesAfter =
                await stakingPool.getPendingUnstakesOf(signer1.address);
            console.log(
                `Signer1 pending unstakes after: ${signer1PendingUnstakesAfter}`
            );
            const signer2PendingUnstakesAfter =
                await stakingPool.getPendingUnstakesOf(signer2.address);
            console.log(
                `Signer2 pending unstakes after: ${signer2PendingUnstakesAfter}`
            );
            console.log(
                `Signer1 balance diff: ${signer1BalanceAfter.sub(
                    signer1BalanceBefore
                )}`
            );
            console.log(
                `Signer2 balance diff: ${signer2BalanceAfter.sub(
                    signer2BalanceBefore
                )}`
            );
            const poolBalanceAfter = await stakingPool.getFreeBalance();
            console.log(
                `Pool free balance after: ${await stakingPool.getFreeBalance()}`
            );

            expect(signer1PendingUnstakesBefore).to.be.closeTo(
                signer1BalanceAfter.sub(signer1BalanceBefore),
                1
            );
            expect(signer2PendingUnstakesBefore).to.be.closeTo(
                signer2BalanceAfter.sub(signer2BalanceBefore),
                1
            );
            expect(signer1PendingUnstakesAfter).to.be.eq(0);
            expect(signer2PendingUnstakesAfter).to.be.eq(0);
            expect(totalPendingUnstakesAfter).to.be.eq(0);
            expect(poolBalanceBefore.sub(poolBalanceAfter)).to.be.eq(
                totalPendingUnstakesBefore
            );
        });
    });

    describe('Push to beacon', function () {
        before(async function () {
            [
                podManager,
                stakingConfig,
                ratioFeed,
                certificateToken,
                stakingPool,
            ] = await init();
        });

        it('pushToBeacon(): Reverts: when pool balance < 32Eth', async function () {
            await expect(
                stakingPool
                    .connect(operator)
                    .pushToBeacon(
                        '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25b',
                        '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                        '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97'
                    )
            ).to.be.revertedWith('pending ethers not enough');
        });

        it('pushToBeaconMulti(): Reverts: when pool balance not enough', async function () {
            await stakingPool
                .connect(signer2)
                .stakeCerts({ value: e18.mul(33) });
            const pubkey1 =
                '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25b';
            const pubkey2 =
                '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25a';
            await expect(
                stakingPool
                    .connect(operator)
                    .pushToBeaconMulti(
                        [pubkey1, pubkey2],
                        [
                            '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                            '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                        ],
                        [
                            '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97',
                            '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97',
                        ]
                    )
            ).to.be.revertedWith('pending ethers not enough');
        });

        it('pushToBeacon()', async function () {
            await stakingPool
                .connect(signer1)
                .stakeCerts({ value: e18.mul(33) });

            const pubkey =
                '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25b';
            const tx = await stakingPool
                .connect(operator)
                .pushToBeacon(
                    '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25b',
                    '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                    '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97'
                );
            const rec = await tx.wait();
            const events = rec.events?.filter((e) => {
                return e.event === 'PoolOnGoing';
            });
            expect(events.length).to.be.eq(1);
            expect(events[0].args['pubkey']).to.be.eq(pubkey);
        });

        it('pushToBeaconMulti()', async function () {
            await stakingPool
                .connect(signer2)
                .stakeCerts({ value: e18.mul(65) });
            const pubkey1 =
                '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25b';
            const pubkey2 =
                '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25a';
            const tx = await stakingPool
                .connect(operator)
                .pushToBeaconMulti(
                    [pubkey1, pubkey2],
                    [
                        '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                        '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                    ],
                    [
                        '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97',
                        '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97',
                    ]
                );
            const rec = await tx.wait();
            const events = rec.events?.filter((e) => {
                return e.event === 'PoolOnGoing';
            });
            expect(events.length).to.be.eq(2);
            expect(events[0].args['pubkey']).to.be.eq(pubkey1);
            expect(events[1].args['pubkey']).to.be.eq(pubkey2);
        });

        it('pushToBeacon(): Only operator can', async function () {
            await expect(
                stakingPool
                    .connect(governance)
                    .pushToBeacon(
                        '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25b',
                        '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                        '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97'
                    )
            ).to.be.revertedWith('StakingPool: only consensus allowed');
        });

        it('pushToBeaconMulti(): Only operator can', async function () {
            const pubkey1 =
                '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25b';
            const pubkey2 =
                '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25a';
            await expect(
                stakingPool
                    .connect(governance)
                    .pushToBeaconMulti(
                        [pubkey1, pubkey2],
                        [
                            '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                            '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                        ],
                        [
                            '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97',
                            '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97',
                        ]
                    )
            ).to.be.revertedWith('StakingPool: only consensus allowed');
        });
    });
});

async function updateRatio(ratioFeed, token, ratio) {
    await increaseChainTimeForSeconds(60 * 60 * 12 + 1); //+12h
    await ratioFeed
        .connect(operator)
        .updateRatioBatch([token.address], [ratio]);
    console.log(`### Ratio updated: ${await token.ratio()}`);
}

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

async function hhReset() {
    await network.provider.request({
        method: 'hardhat_reset',
        params: [],
    });
}
