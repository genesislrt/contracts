import { ethers, upgrades } from 'hardhat';
import { DeployFunction, DeploymentsExtension } from 'hardhat-deploy/types';

const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms));

const transferAdminOwnership = async (
    { get }: DeploymentsExtension,
    contractName: string,
    to: string
) => {
    const contract = await get(contractName);
    await upgrades.admin.transferProxyAdminOwnership(contract.address, to);
    console.log(`ProxyAdmin ownership of ${contractName} transferred to ${to}`);
    await sleep(12_000);
};

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

    await execute('ProtocolConfig', executeCfg, 'setGovernance', governance);

    await transferAdminOwnership(deployments, 'ProtocolConfig', governance);
    await transferAdminOwnership(deployments, 'cToken', governance);
    await transferAdminOwnership(deployments, 'FeeCollector', governance);
    await transferAdminOwnership(deployments, 'RatioFeed', governance);
    await transferAdminOwnership(deployments, 'RestakingPool', governance);

    return true;
};

module.exports = func;
module.exports.tags = ['01_ownership'];
module.exports.dependencies = [];
module.exports.skip = false;
module.exports.id = '01';
