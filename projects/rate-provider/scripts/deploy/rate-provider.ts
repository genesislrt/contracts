import { ethers, run } from "hardhat";

export const sleep = (ms: number) => new Promise((r) => setTimeout(r, ms));

async function main() {
  const args = ["0xf073bAC22DAb7FaF4a3Dd6c6189a70D54110525C"]; // InETH

  const rateProvider = await ethers.deployContract(
    "GenEthRateProvider",
    args,
  );
  await rateProvider.waitForDeployment();

  const block = await ethers.provider.getBlockNumber();
  while ((await ethers.provider.getBlockNumber()) < block + 5) {
    console.log("waiting before verification...");
    await sleep(6_000);
  }

  await run("verify:verify", {
    address: await rateProvider.getAddress(),
    constructorArguments: args,
    contract: "contracts/RateProvider.sol:GenEthRateProvider",
  });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
