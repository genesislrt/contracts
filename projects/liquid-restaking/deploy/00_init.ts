import { ethers, upgrades } from 'hardhat';
import { DeployFunction } from 'hardhat-deploy/types';
import { ozDeploy } from '../scripts/deploy-helpers';

const func: DeployFunction = async function ({
    getNamedAccounts,
    deployments,
}) {
    const { deploy, execute } = deployments;
    const {
        deployer,
        operator,
        governance,
        treasury,
        eigenPodManager,
        eigenDelegationManager,
    } = await getNamedAccounts();

    const config = await ozDeploy(deployments, 'ProtocolConfig', [
        governance,
        operator,
        treasury,
    ]);

    const ratioFeed = await ozDeploy(deployments, 'RatioFeed', [
        await config.getAddress(),
        '40000',
    ]);

    const cToken = await ozDeploy(deployments, 'cToken', [
        await config.getAddress(),
        'Genesis ETH',
        'genETH',
    ]);

    const restakingPool = await ozDeploy(deployments, 'RestakingPool', [
        await config.getAddress(),
        '200000',
    ]);

    const feeCollector = await ozDeploy(deployments, 'FeeCollector', [
        await config.getAddress(),
        '1500',
    ]);

    const executeCfg = {
        from: deployer,
        log: true,
    };

    await execute(
        'ProtocolConfig',
        executeCfg,
        'setRatioFeed',
        await ratioFeed.getAddress()
    );

    await execute(
        'ProtocolConfig',
        executeCfg,
        'setRestakingPool',
        await restakingPool.getAddress()
    );

    await execute(
        'ProtocolConfig',
        executeCfg,
        'setCToken',
        await cToken.getAddress()
    );

    // deploy restaker sub-protocol

    const restakerFacets = await ozDeploy(deployments, 'RestakerFacets', [
        governance,
        eigenPodManager,
        eigenDelegationManager,
    ]);

    const beacon = await upgrades.deployBeacon(
        await ethers.getContractFactory('Restaker')
    );

    const restakerDeployer = await deploy('RestakerDeployer', {
        log: true,
        from: deployer,
        args: [await beacon.getAddress(), await restakerFacets.getAddress()],
    });

    await execute(
        'ProtocolConfig',
        executeCfg,
        'setRestakerDeployer',
        restakerDeployer.address
    );

    return true;
};

module.exports = func;
module.exports.tags = ['00_init'];
module.exports.dependencies = [];
module.exports.skip = false;
module.exports.id = '00';
