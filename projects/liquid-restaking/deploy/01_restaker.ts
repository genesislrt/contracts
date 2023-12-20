const { ethers, upgrades } = require('hardhat');
import { DeployFunction } from 'hardhat-deploy/types';

const { constants } = require('@openzeppelin/test-helpers');

const func: DeployFunction = async function ({
    getNamedAccounts,
    network,
    deployments,
}) {
    const { deploy, execute } = deployments;
    const { deployer, governance, eigenPodManager, delegationManager } =
        await getNamedAccounts();

    const beacon = await upgrades.deployBeacon(
        await ethers.getContractFactory('Restaker')
    );
    await beacon.deployed();

    await deploy('StakingConfig', {
        from: deployer,
        log: true,
        args: [],
        skipIfAlreadyDeployed: true,
        proxy: {
            proxyContract: 'OpenZeppelinTransparentProxy',
            upgradeIndex: 1,
        },
    });

    const restakerFacets = await deploy('RestakerFacets', {
        from: deployer,
        log: true,
        args: [],
        skipIfAlreadyDeployed: true,
        proxy: {
            upgradeIndex: 0,
            proxyContract: 'OpenZeppelinTransparentProxy',
            execute: {
                methodName: 'initialize',
                args: [governance, eigenPodManager, delegationManager],
            },
        },
    });

    const restakerDeployer = await deploy('RestakerDeployer', {
        from: deployer,
        log: true,
        args: [beacon.address, restakerFacets.address],
    });

    await execute(
        'StakingConfig',
        {
            from: deployer,
            log: true,
        },
        'setRestakerDeployer',
        restakerDeployer.address
    );

    return true;
};
module.exports = func;
module.exports.tags = ['01_restaker'];
module.exports.dependencies = [];
module.exports.skip = false;
module.exports.id = '01';
