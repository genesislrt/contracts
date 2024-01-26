import { loadFixture } from '@nomicfoundation/hardhat-network-helpers';
import { ethers, network, upgrades } from 'hardhat';
import { expect } from 'chai';
import {
    deployConfig,
    deployEigenMocks,
    deployLiquidRestaking,
    deployRestakerContacts,
} from './helpers/deploy';
import {
    CToken,
    ExpensiveStakerMock,
    ProtocolConfig,
    RatioFeed,
    RestakerDeployer,
    RestakingPool,
} from '../typechain-types';
import { HardhatEthersSigner } from '@nomicfoundation/hardhat-ethers/signers';
import { _1E18 } from './helpers/constants';
import { randomBN, randomBNbyMax } from './helpers/math';
import { increaseChainTimeForSeconds } from './helpers/evmutils';

const TOKEN_NAME = 'Token Name',
    TOKEN_SYMBOL = 'Token Symbol',
    TEST_PROVIDER = 'TEST_PROVIDER',
    DISTRIBUTE_GAS_LIMIT = 250000n,
    MIN_UNSTAKE = 10000000000n,
    MIN_STAKE = 1000000000n,
    MAX_TVL = 32_000_000_000_000_000_000n;

const ceilN = (n: bigint, d: bigint) => n / d + (n % d ? 1n : 0n);

let governance: HardhatEthersSigner,
    operator: HardhatEthersSigner,
    treasury: HardhatEthersSigner,
    signer1: HardhatEthersSigner,
    signer2: HardhatEthersSigner,
    signer3: HardhatEthersSigner;

const init = async () => {
    [governance, operator, treasury, signer1, signer2, signer3] =
        await ethers.getSigners();

    const protocolConfig = await deployConfig([governance, operator, treasury]);

    // EigenLayr
    const el = await deployEigenMocks({ protocolConfig });

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

    const expensiveStaker = await ethers.deployContract('ExpensiveStakerMock');
    await expensiveStaker.waitForDeployment();

    return [
        protocolConfig,
        restakingPool,
        cToken,
        ratioFeed,
        restakerDeployer,
        expensiveStaker,
    ];
};

describe('RestakingPool', function () {
    let config: ProtocolConfig,
        pool: RestakingPool,
        cToken: CToken,
        feed: RatioFeed,
        deployer: RestakerDeployer,
        expensiveStaker: ExpensiveStakerMock;

    describe('Getters and Setters', function () {
        before(async function () {
            [config, pool, cToken, feed, deployer] = await loadFixture(init);
        });

        it('getMinStake()', async function () {
            expect(await pool.getMinStake()).to.be.eq('1');
        });

        it('getMinUnstake()', async function () {
            expect(await pool.getMinUnstake()).to.be.eq('1');
        });

        it('setDistributeGasLimit()', async function () {
            await expect(pool.setDistributeGasLimit('30'))
                .to.emit(pool, 'DistributeGasLimitChanged')
                .withArgs(DISTRIBUTE_GAS_LIMIT, '30');
        });

        it('setDistributeGasLimit(): reverts: only governance can', async function () {
            await expect(
                pool.connect(operator).setDistributeGasLimit('30')
            ).to.be.revertedWithCustomError(pool, 'OnlyGovernanceAllowed');
        });

        // TODO: check that distribute gas limit cannot be greater than max
    });

    describe('stake', function () {
        before(async function () {
            [config, pool, cToken, feed] = await loadFixture(init);
        });

        it('Reverts: when amount > available', async () => {
            const available = await pool.availableToStake();
            expect(available).to.be.eq(MAX_TVL);
            await expect(
                pool.connect(signer1).stake({ value: available + 1n })
            ).to.be.revertedWithCustomError(
                pool,
                'PoolStakeAmGreaterThanAvailable'
            );
        });

        const amounts = [
            { name: 'Random value', amount: async (x) => randomBN(19) },
            {
                name: '999999999999999999',
                amount: async (x) => 999999999999999999n,
            },
            {
                name: '888888888888888888',
                amount: async (x) => 888888888888888888n,
            },
            {
                name: '777777777777777777',
                amount: async (x) => 777777777777777777n,
            },
            {
                name: '666666666666666666',
                amount: async (x) => 666666666666666666n,
            },
            {
                name: '555555555555555555',
                amount: async (x) => 555555555555555555n,
            },
            {
                name: '444444444444444444',
                amount: async (x) => 444444444444444444n,
            },
            {
                name: '333333333333333333',
                amount: async (x) => 333333333333333333n,
            },
            {
                name: '222222222222222222',
                amount: async (x) => 222222222222222222n,
            },
            {
                name: '111111111111111111',
                amount: async (x) => 111111111111111111n,
            },
            { name: 'Min amount', amount: async (x) => await x.getMinStake() },
        ];

        amounts.forEach((param) => {
            it(`Stake: ${param.name}`, async function () {
                // Update ratio
                const ratio = (await cToken.ratio()) - 1n;
                await updateRatio(feed, cToken, ratio);
                const signerBalanceBefore = await cToken.balanceOf(
                    signer1.address
                );
                const totalSupplyBefore = await cToken.totalSupply();
                const available = await pool.availableToStake();
                expect(available).to.be.eq(
                    MAX_TVL - ceilN(totalSupplyBefore * _1E18, ratio)
                );

                // Stake
                const amount = await param.amount(pool);
                const expectedShares = (amount * ratio) / _1E18;
                await expect(pool.connect(signer1).stake({ value: amount }))
                    .to.emit(pool, 'Staked')
                    .withArgs(
                        signer1.address,
                        amount.toString(),
                        expectedShares.toString()
                    );

                const signerBalanceAfter = await cToken.balanceOf(
                    signer1.address
                );
                const totalSupplyAfter = await cToken.totalSupply();

                expect(signerBalanceAfter - signerBalanceBefore).to.be.closeTo(
                    expectedShares,
                    1
                );
                expect(totalSupplyAfter - totalSupplyBefore).to.be.closeTo(
                    expectedShares,
                    1
                );
            });
        });

        it('Reverts: when amount < min', async function () {
            const amount = (await pool.getMinStake()) - 1n;
            await expect(
                pool.connect(signer1).stake({ value: amount })
            ).to.be.revertedWithCustomError(pool, 'PoolStakeAmLessThanMin');
        });

        it('Reverts: when amount > available', async () => {
            await expect(
                pool
                    .connect(signer1)
                    .stake({ value: (await pool.availableToStake()) + 1n })
            ).to.be.revertedWithCustomError(
                pool,
                'PoolStakeAmGreaterThanAvailable'
            );
        });

        it('Increase max tvl', async () => {
            const newMax = 2_000_000_000_000_000n * _1E18;
            await expect(pool.setMaxTVL(newMax))
                .to.emit(pool, 'MaxTVLChanged')
                .withArgs(32n * _1E18, newMax);
        });

        //Stake many times with different ratio values
        it('Stake many times with different signers and ratio', async function () {
            [config, pool, cToken, feed] = await loadFixture(init);
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
                    const amount = randomBNbyMax(
                        10n ** 18n - MIN_STAKE + MIN_UNSTAKE
                    );
                    expectedPoolBalance = expectedPoolBalance + amount;
                    await pool.connect(signer).stake({ value: amount });
                    const shares = (amount * ratio) / _1E18;
                    const currentShares = signersShares.get(signer.address);
                    signersShares.set(signer.address, currentShares + shares);
                    expectedTotalSupply = expectedTotalSupply + shares;
                }
            }
            expect(await cToken.balanceOf(signer1.address)).to.be.closeTo(
                signersShares.get(signer1.address),
                100
            );
            expect(await cToken.balanceOf(signer2.address)).to.be.closeTo(
                signersShares.get(signer2.address),
                100
            );
            expect(await cToken.balanceOf(signer3.address)).to.be.closeTo(
                signersShares.get(signer3.address),
                100
            );
            expect(await cToken.totalSupply()).to.be.closeTo(
                expectedTotalSupply,
                300
            );
            expect(
                await ethers.provider.getBalance(await pool.getAddress())
            ).to.be.eq(expectedPoolBalance);
        });
    });

    describe('unstake()', function () {
        before(async function () {
            [config, pool, cToken, feed] = await loadFixture(init);
            await pool.setMaxTVL(200n * _1E18);
        });

        const amounts = [
            {
                name: 'Random value to another address',
                shares: async () => randomBN(19),
                receiver: () => signer2,
            },
            {
                name: '999999999999999999',
                shares: async () => 999999999999999999n,
                receiver: () => signer1,
            },
            {
                name: '888888888888888888',
                shares: async () => 888888888888888888n,
                receiver: () => signer1,
            },
            {
                name: '777777777777777777',
                shares: async () => 777777777777777777n,
                receiver: () => signer1,
            },
            {
                name: '666666666666666666',
                shares: async () => 666666666666666666n,
                receiver: () => signer1,
            },
            {
                name: '555555555555555555',
                shares: async () => 555555555555555555n,
                receiver: () => signer1,
            },
            {
                name: '444444444444444444',
                shares: async () => 444444444444444444n,
                receiver: () => signer1,
            },
            {
                name: '333333333333333333',
                shares: async () => 333333333333333333n,
                receiver: () => signer1,
            },
            {
                name: '222222222222222222',
                shares: async () => 222222222222222222n,
                receiver: () => signer1,
            },
            {
                name: '111111111111111111',
                shares: async () => 111111111111111111n,
                receiver: () => signer1,
            },
            {
                name: 'Min amount',
                shares: async (x: RestakingPool) => await x.getMinUnstake(),
                receiver: () => signer1,
            },
        ];

        const receiverUnstakesMap = new Map();
        amounts.forEach((param) => {
            it(`Unstake: ${param.name}`, async function () {
                // Stake once
                if ((await cToken.balanceOf(signer1.address)) === 0n) {
                    await pool.connect(signer1).stake({ value: 50n * _1E18 });
                }
                //Update ratio
                const ratio = (await cToken.ratio()) - 1000n;
                await updateRatio(feed, cToken, ratio);
                //Values before

                const receiver = param.receiver();
                const ownerBalanceBefore = await cToken.balanceOf(
                    signer1.address
                );
                const totalSupplyBefore = await cToken.totalSupply();
                const totalPendingUnstakesBefore =
                    await pool.getTotalPendingUnstakes();
                const receiverPendingUnstakesBefore =
                    await pool.getTotalUnstakesOf(receiver.address);

                //Unstake
                const shares = await param.shares(pool);
                const expectedAsset = (shares * _1E18) / ratio + 1n; //Rounding up

                await expect(
                    pool.connect(signer1).unstake(receiver.address, shares)
                )
                    .to.emit(pool, 'Unstaked')
                    .withArgs(
                        signer1.address,
                        receiver.address,
                        expectedAsset,
                        shares
                    );

                const ownerBalanceAfter = await cToken.balanceOf(
                    signer1.address
                );
                const totalSupplyAfter = await cToken.totalSupply();
                const totalPendingUnstakesAfter =
                    await pool.getTotalPendingUnstakes();
                const receiverPendingUnstakesAfter =
                    await pool.getTotalUnstakesOf(receiver.address);
                const receiverPendingRequests = (
                    await pool.getUnstakesOf(receiver.address)
                ).map((bn) => bn.toString());
                if (!receiverUnstakesMap.has(receiver.address)) {
                    receiverUnstakesMap.set(receiver.address, [
                        `${receiver.address},${expectedAsset.toString()}`,
                    ]);
                } else {
                    receiverUnstakesMap
                        .get(receiver.address)
                        .push(
                            `${receiver.address},${expectedAsset.toString()}`
                        );
                }

                expect(ownerBalanceBefore - ownerBalanceAfter).to.be.closeTo(
                    shares,
                    1
                );
                expect(totalSupplyBefore - totalSupplyAfter).to.be.closeTo(
                    shares,
                    1
                );
                expect(
                    totalPendingUnstakesAfter - totalPendingUnstakesBefore
                ).to.be.closeTo(expectedAsset, 1);
                expect(
                    receiverPendingUnstakesAfter - receiverPendingUnstakesBefore
                ).to.be.closeTo(expectedAsset, 1);

                expect(receiverPendingRequests).to.include.members(
                    receiverUnstakesMap.get(receiver.address)
                );
            });
        });

        it('Reverts: when shares < min', async function () {
            const ratio = await cToken.ratio();
            const shares = (await pool.getMinUnstake()) - 1n;
            await expect(
                pool.connect(signer1).unstake(signer1.address, shares)
            ).to.be.revertedWithCustomError(pool, 'PoolUnstakeAmLessThanMin');
        });

        it('Reverts: when exceed user balance', async function () {
            await pool.connect(signer1).stake({ value: 10n ** 18n });
            const shares = (await cToken.balanceOf(signer1.address)) + 1n;
            await expect(
                pool.connect(signer1).unstake(signer1.address, shares)
            ).to.be.revertedWithCustomError(cToken, 'ERC20InsufficientBalance');
        });

        it('Reverts: receiver is zero address', async function () {
            await pool.connect(signer1).stake({ value: 10n ** 18n });
            const shares = await cToken.convertToShares(10n ** 18n);
            await expect(
                pool.connect(signer1).unstake(ethers.ZeroAddress, shares)
            ).to.be.revertedWithCustomError(pool, 'PoolZeroAddress');
        });

        it('Unstake all', async function () {
            const ratio = await cToken.ratio();

            const shares = await cToken.balanceOf(signer1.address);
            const expectedAsset = (shares * _1E18) / ratio + 1n; //Rounding up
            const totalPendingUnstakesBefore =
                await pool.getTotalPendingUnstakes();
            const receiverPendingUnstakesBefore = await pool.getTotalUnstakesOf(
                signer1.address
            );
            await pool.connect(signer1).unstake(signer1.address, shares);

            const ownerBalanceAfter = await cToken.balanceOf(signer1.address);
            const totalSupplyAfter = await cToken.totalSupply();
            const totalPendingUnstakesAfter =
                await pool.getTotalPendingUnstakes();
            const receiverPendingUnstakesAfter = await pool.getTotalUnstakesOf(
                signer1.address
            );

            expect(ownerBalanceAfter).to.be.eq(0);
            expect(totalSupplyAfter).to.be.eq(0);
            expect(
                totalPendingUnstakesAfter - totalPendingUnstakesBefore
            ).to.be.closeTo(expectedAsset, 1);
            expect(
                receiverPendingUnstakesAfter - receiverPendingUnstakesBefore
            ).to.be.closeTo(expectedAsset, 1);
        });
    });

    describe('Deposit', function () {
        before(async function () {
            [config, pool, cToken, feed] = await loadFixture(init);
            await pool.addRestaker(TEST_PROVIDER);
            await pool.setMaxTVL(200n * _1E18);
        });

        it('Cannot add one provider twice', async () => {
            await expect(
                pool.addRestaker(TEST_PROVIDER)
            ).to.be.revertedWithCustomError(pool, 'PoolRestakerExists');
        });

        it('batchDeposit(): Reverts: when pool balance < 32Eth', async function () {
            await expect(
                pool
                    .connect(operator)
                    .batchDeposit(
                        TEST_PROVIDER,
                        [
                            '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25b',
                        ],
                        [
                            '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                        ],
                        [
                            '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97',
                        ]
                    )
            ).to.be.revertedWithCustomError(pool, 'PoolInsufficientBalance');
        });

        it('batchDeposit()', async function () {
            await pool.connect(signer1).stake({ value: _1E18 * 33n });

            const pubkey =
                '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25b';
            await expect(
                pool
                    .connect(operator)
                    .batchDeposit(
                        TEST_PROVIDER,
                        [pubkey],
                        [
                            '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                        ],
                        [
                            '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97',
                        ]
                    )
            )
                .to.emit(pool, 'Deposited')
                .withArgs(TEST_PROVIDER, [pubkey]);
        });

        it('batchDeposit()', async function () {
            await pool.connect(signer2).stake({ value: _1E18 * 65n });
            const pubkeys = [
                '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25b',
                '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25a',
            ];
            await expect(
                pool
                    .connect(operator)
                    .batchDeposit(
                        TEST_PROVIDER,
                        pubkeys,
                        [
                            '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                            '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                        ],
                        [
                            '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97',
                            '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97',
                        ]
                    )
            )
                .to.emit(pool, 'Deposited')
                .withArgs(TEST_PROVIDER, pubkeys);
        });

        it('batchDeposit(): Only operator can', async function () {
            await expect(
                pool
                    .connect(governance)
                    .batchDeposit(
                        TEST_PROVIDER,
                        [
                            '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25b',
                        ],
                        [
                            '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                        ],
                        [
                            '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97',
                        ]
                    )
            ).to.be.revertedWithCustomError(pool, 'OnlyOperatorAllowed');
        });

        it('batchDeposit(): provider not exists', async function () {
            await pool.connect(signer2).stake({ value: _1E18 * 32n });
            await expect(
                pool
                    .connect(operator)
                    .batchDeposit(
                        'provider',
                        [
                            '0xb8ed0276c4c631f3901bafa668916720f2606f58e0befab541f0cf9e0ec67a8066577e9a01ce58d4e47fba56c516f25b',
                        ],
                        [
                            '0x927b16171b51ca4ccab59de07ea20dacc33baa0f89f06b6a762051cac07233eb613a6c272b724a46b8145850b8851e4a12eb470bfb140e028ae0ac794f3a890ec4fac33910d338343f059d93a6d688238510c147f155d984de7c01daa0d3241b',
                        ],
                        [
                            '0x50021ea68edb12aaa54fc8a2706b2f4b1d35d1406512fc6de230e0ea0391cf97',
                        ]
                    )
            ).to.be.revertedWithCustomError(pool, 'PoolRestakerNotExists');
        });
    });

    describe('distribute unstakes and claims', function () {
        before(async function () {
            [config, pool, cToken, feed, deployer] = await loadFixture(init);
            await pool.setMaxTVL(_1E18 * _1E18);
            await pool.setMinStake(0);
            await pool.setMinUnstake(0);
        });

        const difficult = 50n;
        const signers = [() => signer1, () => signer2, () => signer3];

        for (let i = 0n; i < difficult; i++) {
            const signerIndex = Number(i) % signers.length;

            it(`unstake from contract (${i}/${difficult})`, async () => {
                const minStake = await pool.getMinStake();
                const signer = signers[signerIndex]();
                await pool
                    .connect(signer)
                    .stake({ value: minStake * (i + 10n) });
                /// get tokens amount and unstake a bit less
                const minUnstake = await pool.getMinUnstake();
                const tokensAm = (await cToken.balanceOf(signer.address)) - i;
                await pool.connect(signer).unstake(signer, tokensAm - i);
                // check how much is pending
            });
        }

        it('distributeUnstakes', async () => {
            /// get current amounts
            const signer1Expected = await pool.getTotalUnstakesOf(
                signer1.address
            );
            const signer2Expected = await pool.getTotalUnstakesOf(
                signer2.address
            );
            const signer3Expected = await pool.getTotalUnstakesOf(
                signer3.address
            );

            await pool.connect(operator).distributeUnstakes('0');

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
                    .to.emit(pool, 'UnstakeClaimed')
                    .withArgs(
                        signers[i]().address,
                        governance.address,
                        claimable
                    );
            });
        }
    });
});

async function updateRatio(ratioFeed: RatioFeed, token: CToken, ratio: BigInt) {
    await increaseChainTimeForSeconds(60 * 60 * 12 + 1); //+12h
    await ratioFeed.connect(operator).updateRatio(token, ratio.toString());
    expect(await token.ratio()).to.be.eq(ratio);
}
