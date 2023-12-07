import "@nomicfoundation/hardhat-toolbox";
import "@openzeppelin/hardhat-upgrades";
import "hardhat-deploy";

require('dotenv').config()
const gasPrice = parseInt(process.env.GAS_PRICE) || 'auto';

module.exports = {
    hardhat: {
        gas: 2100000,
        gasPrice: 8000000000,
        allowUnlimitedContractSize: true,
    },
    networks: {
        // Ethereum
        mainnet: {
            url: process.env.MAINNET_RPC || 'https://rpc.ankr.com/eth',
            chainId: 1,
            gas: 8000000,
            gasPrice: gasPrice,
            accounts: [process.env.DEPLOYER_PRIVATE_KEY],
        },
        goerli: {
            url: process.env.GOERLI_RPC || 'https://rpc.ankr.com/eth_goerli',
            chainId: 5,
            gasPrice: gasPrice,
            gas: 8000000,
            accounts: [process.env.DEPLOYER_PRIVATE_KEY],
        },
        local: {
            url: process.env.LOCAL_RPC || 'http://127.0.0.1:8545',
            chainId: 1337,
            gasPrice: 20000000000,
            gas: 6721975,
        }
    },
    solidity: {
        version: "0.8.19",
        settings: {
            optimizer: {
                enabled: true,
                runs: 200
            }
        }
    },
    mocha: {
        enableTimeouts: false,
        before_timeout: 120000
    },
    gasReporter: {
        enabled: true
    },
};
