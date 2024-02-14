import { ethers, run } from 'hardhat';

export const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms));

async function main() {
    const args = ['0x51046E9a40E217Fb63f27440e8E7949bbAcF9309']; // genETH

    const rateProvider = await ethers.deployContract(
        'GenEthRateProvider',
        args
    );
    await rateProvider.waitForDeployment();

    const block = await ethers.provider.getBlockNumber();
    while ((await ethers.provider.getBlockNumber()) < block + 5) {
        console.log('waiting before verification...');
        await sleep(6_000);
    }

    await run('verify:verify', {
        address: await rateProvider.getAddress(),
        constructorArguments: args,
        contract: 'contracts/RateProvider.sol:GenEthRateProvider',
    });
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
