import { ethers, upgrades } from 'hardhat';
import { DeployFunction } from 'hardhat-deploy/types';

const func: DeployFunction = async function ({ deployments }) {
    const result = await upgrades.prepareUpgrade(
        'RestakingPool',
        await ethers.getContractFactory('RestakingPool')
    );
    console.log(result);
    return true;
};

module.exports = func;
module.exports.tags = ['04_upgrade_pool'];
module.exports.dependencies = [];
module.exports.skip = false;
module.exports.id = '04';
