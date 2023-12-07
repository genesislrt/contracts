require('dotenv').config()
const HDWalletProvider = require('@truffle/hdwallet-provider');

module.exports = {
    networks: {
        local: {
            host: "127.0.0.1",     // Localhost (default: none)
            port: 8545,            // Standard Ethereum port (default: none)
            network_id: "*",       // Any network (default: none)
        },
        goerli: {
            provider: () =>
                new HDWalletProvider({
                    privateKeys: [
                        process.env.DEPLOYER_PRIVATE_KEY,
                    ],
                    providerOrUrl: process.env.PROVIDER_URL,
                }),
            network_id: 5,
            confirmations: 0,
            gas: 8000000,
            timeoutBlocks: 400,
            skipDryRun: true,
            networkCheckTimeout: 100000000,
        },
    },

    mocha: {
        enableTimeouts: false,
        before_timeout: 12000000,
        reporterOptions: {
            showTimeSpent: true,
            showMethodSig: true,
        },
        reporter: 'eth-gas-reporter'
    },
    compilers: {
        solc: {
            version: "0.8.20",
            settings: {
                optimizer: {
                    enabled: true,
                    runs: 200
                },
            }
        }
    },
    plugins: [
        'truffle-plugin-verify'
    ],
    api_keys: {
        etherscan: process.env.ETHERSCAN_API_KEY
    }
};
