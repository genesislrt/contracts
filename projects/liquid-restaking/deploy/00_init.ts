import { DeployFunction } from 'hardhat-deploy/types';

const { constants } = require('@openzeppelin/test-helpers');

const func: DeployFunction = async function ({
    getNamedAccounts,
    network,
    deployments,
}) {
    const { deploy, execute } = deployments;
    const { deployer, operator, governance, treasury, eigenPodManager } =
        await getNamedAccounts();

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
                    operator,
                    governance,
                    treasury,
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

    return true;
};
module.exports = func;
module.exports.tags = ['00_init'];
module.exports.dependencies = [];
module.exports.skip = true;
module.exports.id = '00';
