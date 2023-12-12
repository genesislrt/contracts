import '@nomicfoundation/hardhat-toolbox';
import '@openzeppelin/hardhat-upgrades';
import 'hardhat-deploy';

require('dotenv').config();

const accounts = process.env.DEPLOYER_PRIVATE_KEY
    ? [process.env.DEPLOYER_PRIVATE_KEY]
    : ['1495992B2A5CC4DD53E231157BBF401329BD1B7EE355CEAB55A791398921CA17'];
const gasPrice = process.env.GAS_PRICE
    ? parseInt(process.env.GAS_PRICE)
    : 'auto';

require('dotenv').config();

module.exports = {
    hardhat: {
        gas: 8000000,
        gasPrice,
        allowUnlimitedContractSize: false,
    },
    networks: {
        // Ethereum
        mainnet: {
            url: process.env.MAINNET_RPC || 'https://rpc.ankr.com/eth',
            chainId: 1,
            gas: 8000000,
            gasPrice,
            accounts,
        },
        goerli: {
            url: process.env.GOERLI_RPC || 'https://rpc.ankr.com/eth_goerli',
            chainId: 5,
            gasPrice,
            gas: 8000000,
            accounts,
        },
        local: {
            url: process.env.LOCAL_RPC || 'http://127.0.0.1:8545',
            chainId: 1337,
            gasPrice: 20000000000,
            gas: 6721975,
        },
    },
    solidity: {
        version: '0.8.20',
        settings: {
            optimizer: {
                enabled: true,
                runs: 200,
            },
        },
    },
    mocha: {
        enableTimeouts: false,
        before_timeout: 120000,
    },
    gasReporter: {
        enabled: true,
    },
};
