import { ethers, upgrades } from 'hardhat';
import { DeployFunction } from 'hardhat-deploy/types';
import { ozDeploy } from '../scripts/deploy-helpers';

const func: DeployFunction = async function ({
    getNamedAccounts,
    deployments,
}) {
    const { execute, get } = deployments;
    const { deployer, governance, treasury } = await getNamedAccounts();

    const executeCfg = {
        from: deployer,
        log: true,
    };

    const protocolConfig = await get('ProtocolConfig');

    await upgrades.admin.transferProxyAdminOwnership(
        protocolConfig.address,
        governance
    );

    await execute('ProtocolConfig', executeCfg, 'setGovernance', governance);

    return true;
};

module.exports = func;
module.exports.tags = ['01_ownership'];
module.exports.dependencies = [];
module.exports.skip = false;
module.exports.id = '01';
