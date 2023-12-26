import { DeployFunction } from 'hardhat-deploy/types';

const { constants } = require('@openzeppelin/test-helpers');

const func: DeployFunction = async function ({
    getNamedAccounts,
    deployments,
}) {
    const { deploy, get } = deployments;

    const { deployer } = await getNamedAccounts();

    await deploy('StakingPool', {
        from: deployer,
        log: true,
        args: [],
        skipIfAlreadyDeployed: false,
        proxy: {
            proxyContract: 'OpenZeppelinTransparentProxy',
        },
    });

    return true;
};

module.exports = func;
module.exports.tags = ['03_delegate'];
module.exports.dependencies = [];
module.exports.skip = false;
module.exports.id = '03';
