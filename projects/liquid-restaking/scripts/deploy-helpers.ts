import { ethers, upgrades } from 'hardhat';
import { DeploymentsExtension, Receipt } from 'hardhat-deploy/types';
import { ContractTransactionReceipt } from 'ethers';

const TxResultToReceipt = async (
    result: undefined | null | ContractTransactionReceipt
): Promise<Receipt | undefined> => {
    if (!result) {
        return undefined;
    }
    return {
        ...result,
        transactionHash: result.hash,
        transactionIndex: result.index,
        status: result.status || undefined,
        confirmations: await result.confirmations(),
        cumulativeGasUsed: result.cumulativeGasUsed.toString(),
        gasUsed: result.gasUsed.toString(),
        contractAddress: result.contractAddress || undefined,
        to: result.to || undefined,
    };
};

export async function ozDeploy(
    { save, get }: DeploymentsExtension,
    contractName: string,
    args: unknown[],
    forceDeploy = false
) {
    if (!forceDeploy) {
        try {
            const existing = await get(contractName);
            return existing;
        } catch (e: any) {
            console.log(`No deployment found for ${contractName}`);
        }
    }

    const contractFactory = await ethers.getContractFactory(contractName);

    console.log(`deploying ${contractName}...`);

    const contract = await upgrades.deployProxy(contractFactory, args, {
        redeployImplementation: 'onchange',
    });
    await contract.waitForDeployment();

    console.log(`saving ${contractName}...`);

    await save(contractName, {
        abi: JSON.parse(contract.interface.formatJson()) as any[],
        address: await contract.getAddress(),
        receipt: await TxResultToReceipt(
            await contract.deploymentTransaction()?.wait()
        ),
        bytecode: contractFactory.bytecode,
        deployedBytecode: (await contract.getDeployedCode()) || undefined,
    });

    return get(contractName);
}

export const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms));

export async function transferAdminOwnership(
    { get }: DeploymentsExtension,
    contractName: string,
    to: string
) {
    const contract = await get(contractName);
    await upgrades.admin.transferProxyAdminOwnership(contract.address, to);
    console.log(`ProxyAdmin ownership of ${contractName} transferred to ${to}`);
    await sleep(24_000);
}
