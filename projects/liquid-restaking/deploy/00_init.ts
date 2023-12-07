import { DeployFunction } from 'hardhat-deploy/types';

const { constants } = require('@openzeppelin/test-helpers');

const func: DeployFunction = async function ({
    getUnnamedAccounts,
    network,
    deployments,
}) {
    const { deploy, execute } = deployments;
    const [deployer] = await getUnnamedAccounts();
    let eigenPodManager;
    let operatorAddress;
    let consensusAddress;
    let governanceAddress;
    let treasureAddress;

    switch (network.name) {
        case 'goerli':
        case 'hardhat':
            eigenPodManager = '0xa286b84C96aF280a49Fe1F40B9627C2A2827df41';
            operatorAddress = deployer;
            governanceAddress = deployer;
            treasureAddress = deployer;
            break;
        default: {
            throw new Error(`Not supported network (${network.name})`);
        }
    }

    const config = await deploy('StakingConfig', {
        from: deployer,
        log: true,
        args: [],
        skipIfAlreadyDeployed: true,
        proxy: {
            upgradeIndex: 0,
            proxyContract: 'OpenZeppelinTransparentProxy',
            execute: {
                methodName: 'initialize',
                args: [
                    0,
                    0,
                    operatorAddress,
                    governanceAddress,
                    treasureAddress,
                    constants.ZERO_ADDRESS,
                    constants.ZERO_ADDRESS,
                    constants.ZERO_ADDRESS,
                    eigenPodManager,
                ],
            },
        },
    });

    const feed = await deploy('RatioFeed', {
        from: deployer,
        log: true,
        args: [],
        skipIfAlreadyDeployed: true,
        proxy: {
            upgradeIndex: 0,
            proxyContract: 'OpenZeppelinTransparentProxy',
            execute: {
                methodName: 'initialize',
                args: [config.address],
            },
        },
    });

    const certToken = await deploy('CertificateToken', {
        from: deployer,
        log: true,
        args: [],
        skipIfAlreadyDeployed: true,
        proxy: {
            upgradeIndex: 0,
            proxyContract: 'OpenZeppelinTransparentProxy',
            execute: {
                methodName: 'initialize',
                args: [config.address, 'Genesis ETH', 'genETH'],
            },
        },
    });

    if (certToken.newlyDeployed) {
        await execute(
            'RatioFeed',
            {
                from: deployer,
                log: true,
            },
            'setRatioThreshold',
            1
        );
        await execute(
            'RatioFeed',
            {
                from: deployer,
                log: true,
            },
            'updateRatioBatch',
            [certToken.address],
            ['1000000000000000000']
        );
    }

    const stakingPool = await deploy('StakingPool', {
        from: deployer,
        log: true,
        args: [],
        skipIfAlreadyDeployed: true,
        proxy: {
            upgradeIndex: 0,
            proxyContract: 'OpenZeppelinTransparentProxy',
            execute: {
                methodName: 'initialize',
                args: [config.address, 0],
            },
        },
    });

    if (config.newlyDeployed) {
        await execute(
            'StakingConfig',
            {
                from: deployer,
                log: true,
            },
            'setRatioFeedAddress',
            feed.address
        );
        await execute(
            'StakingConfig',
            {
                from: deployer,
                log: true,
            },
            'setStakingPoolAddress',
            stakingPool.address
        );
        await execute(
            'StakingConfig',
            {
                from: deployer,
                log: true,
            },
            'setCertTokenAddress',
            certToken.address
        );
    }
};
export default func;
func.tags = ['init'];
