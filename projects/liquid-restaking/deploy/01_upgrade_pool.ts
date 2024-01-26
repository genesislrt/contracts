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
        // skipIfAlreadyDeployed: true,
        proxy: {
            proxyContract: 'OpenZeppelinTransparentProxy',
        },
    });
};
export default func;
func.tags = ['upgrade_pool_01'];
