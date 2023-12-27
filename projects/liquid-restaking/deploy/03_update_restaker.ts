import { ethers, upgrades } from 'hardhat';
import { DeployFunction, DeploymentsExtension } from 'hardhat-deploy/types';

const func: DeployFunction = async function ({ deployments }) {
    const { get } = deployments;

    const restakerDeployer = await get('RestakerDeployer');

    if (!restakerDeployer.args) {
        throw Error('restakerDeployer deployment doesnt contain args');
    }

    const beacon = restakerDeployer.args[0];

    const upgradedBeacon = await upgrades.upgradeBeacon(
        beacon,
        await ethers.getContractFactory('Restaker')
    );
    await upgradedBeacon.waitForDeployment();

    return true;
};

module.exports = func;
module.exports.tags = ['03_upgrade_restaker'];
module.exports.dependencies = [];
module.exports.skip = false;
module.exports.id = '03';
