import { ethers, upgrades } from 'hardhat';
import {
    DelegationManagerMock,
    EigenPodManagerMock,
    ProtocolConfig,
    RestakingPool,
    CToken,
    RatioFeed,
    FeeCollector,
} from '../../typechain-types';
import {
    HardhatEthersSigner,
    SignerWithAddress,
} from '@nomicfoundation/hardhat-ethers/signers';
import { _1E18 } from './constants';

export async function deployEigenMocks({
    protocolConfig,
}: {
    protocolConfig: ProtocolConfig;
}) {
    const beacon = await upgrades.deployBeacon(
        await ethers.getContractFactory('EigenPodMock'),
        {
            redeployImplementation: 'always',
        }
    );
    await beacon.waitForDeployment();

    const eigenPodManager = await ethers.deployContract('EigenPodManagerMock', [
        await beacon.getAddress(),
    ]);
    await eigenPodManager.waitForDeployment();

    const delegationManager = await ethers.deployContract(
        'DelegationManagerMock'
    );
    await delegationManager.waitForDeployment();

    return { eigenPodManager, delegationManager };
}

export async function deployRestakerContacts({
    eigenPodManager,
    delegationManager,
    owner,
    protocolConfig,
}: {
    eigenPodManager: EigenPodManagerMock;
    delegationManager: DelegationManagerMock;
    owner: string;
    protocolConfig: ProtocolConfig;
}) {
    // deploy facets
    const restakerFacets = await upgrades.deployProxy(
        await ethers.getContractFactory('RestakerFacets'),
        [
            owner,
            await eigenPodManager.getAddress(),
            await delegationManager.getAddress(),
        ],
        { redeployImplementation: 'always' }
    );
    await restakerFacets.waitForDeployment();

    // deploy restaker beacon
    const Restaker = await ethers.getContractFactory('Restaker');
    const beacon = await upgrades.deployBeacon(Restaker, {
        redeployImplementation: 'always',
    });
    await beacon.waitForDeployment();

    // deploy RestakerDeployer
    const restakerDeployer = await ethers.deployContract('RestakerDeployer', [
        await beacon.getAddress(),
        await restakerFacets.getAddress(),
    ]);
    await restakerDeployer.waitForDeployment();
    await protocolConfig.setRestakerDeployer(
        await restakerDeployer.getAddress()
    );

    return {
        restakerDeployer,
        restakerFacets,
        beacon,
    };
}

export async function deployConfig(wallets: HardhatEthersSigner[]) {
    const [governance, operator, treasury] = wallets;
    const config = await upgrades.deployProxy(
        await ethers.getContractFactory('ProtocolConfig'),
        [governance.address, operator.address, treasury.address],
        { redeployImplementation: 'always' }
    );
    await config.waitForDeployment();

    return config as unknown as ProtocolConfig;
}

export async function deployRatioFeed(
    protocolConfig: ProtocolConfig,
    ratioThreshold = 10_000_000n
) {
    const ratioFeed = await upgrades.deployProxy(
        await ethers.getContractFactory('RatioFeed'),
        [await protocolConfig.getAddress(), ratioThreshold],
        { redeployImplementation: 'always' }
    );
    await ratioFeed.waitForDeployment();
    await protocolConfig.setRatioFeed(await ratioFeed.getAddress());

    return ratioFeed as unknown as RatioFeed;
}

export async function deployFeeCollector(
    protocolConfig: ProtocolConfig,
    commission = 1000 // 10%
) {
    const feeCollector = await upgrades.deployProxy(
        await ethers.getContractFactory('FeeCollector'),
        [await protocolConfig.getAddress(), commission],
        { redeployImplementation: 'always' }
    );
    await feeCollector.waitForDeployment();

    return feeCollector as unknown as FeeCollector;
}

export async function deployLiquidRestaking({
    protocolConfig,
    tokenName,
    tokenSymbol,
    distributeGasLimit,
    ratioThreshold,
}: {
    protocolConfig: ProtocolConfig;
    tokenName: string;
    tokenSymbol: string;
    distributeGasLimit: bigint;
    ratioThreshold?: bigint;
}) {
    // cToken
    const cToken = await upgrades.deployProxy(
        await ethers.getContractFactory('cToken'),
        [await protocolConfig.getAddress(), tokenName, tokenSymbol],
        { redeployImplementation: 'always' }
    );
    await cToken.waitForDeployment();
    await protocolConfig.setCToken(await cToken.getAddress());

    // Pool
    const restakingPool = await upgrades.deployProxy(
        await ethers.getContractFactory('RestakingPool'),
        [await protocolConfig.getAddress(), distributeGasLimit],
        { redeployImplementation: 'always' }
    );
    await restakingPool.waitForDeployment();
    await protocolConfig.setRestakingPool(await restakingPool.getAddress());

    const ratioFeed = await deployRatioFeed(protocolConfig, ratioThreshold);
    await ratioFeed.repairRatio(await cToken.getAddress(), _1E18); // force update ratio to 1e18

    return {
        cToken: cToken as unknown as CToken,
        restakingPool: restakingPool as unknown as RestakingPool,
        ratioFeed,
    };
}