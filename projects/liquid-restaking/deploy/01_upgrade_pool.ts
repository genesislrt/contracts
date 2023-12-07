import { DeployFunction } from 'hardhat-deploy/types';

const { constants } = require('@openzeppelin/test-helpers');

const func: DeployFunction = async function ({
    getUnnamedAccounts,
    network,
    deployments,
}) {
    const { deploy, execute, get } = deployments;
    const [deployer] = await getUnnamedAccounts();

    await deploy('StakingPool', {
        from: deployer,
        log: true,
        args: [],
        skipIfAlreadyDeployed: true,
        proxy: {
            implementationName: 'StakingPool_V1',
            proxyContract: 'OpenZeppelinTransparentProxy',
            upgradeIndex: 1,
        },
    });
};
export default func;
func.tags = ['upgrade_pool'];
