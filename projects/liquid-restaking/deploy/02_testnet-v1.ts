import { DeployFunction } from 'hardhat-deploy/types';

const { constants } = require('@openzeppelin/test-helpers');

const func: DeployFunction = async function ({
    getNamedAccounts,
    network,
    deployments,
}) {
    const { deploy, get } = deployments;

    const { deployer } = await getNamedAccounts();

    const config = await get('StakingConfig');

    const feeCollector = await deploy('FeeCollector', {
        from: deployer,
        log: true,
        args: [],
        skipIfAlreadyDeployed: true,
        proxy: {
            upgradeIndex: 0,
            proxyContract: 'OpenZeppelinTransparentProxy',
            execute: {
                methodName: 'initialize',
                args: [config.address, '1000'],
            },
        },
    });

    await deploy('RatioFeed', {
        from: deployer,
        log: true,
        args: [],
        skipIfAlreadyDeployed: true,
        proxy: {
            proxyContract: 'OpenZeppelinTransparentProxy',
            upgradeIndex: 1,
        },
    });

    await deploy('StakingPool', {
        from: deployer,
        log: true,
        args: [],
        skipIfAlreadyDeployed: false,
        proxy: {
            proxyContract: 'OpenZeppelinTransparentProxy',
        },
    });

    await deploy('CertificateToken', {
        from: deployer,
        log: true,
        args: [],
        skipIfAlreadyDeployed: true,
        proxy: {
            proxyContract: 'OpenZeppelinTransparentProxy',
            upgradeIndex: 1,
        },
    });

    return true;
};

module.exports = func;
module.exports.tags = ['02_V1'];
module.exports.dependencies = [];
module.exports.skip = false;
module.exports.id = '02';
