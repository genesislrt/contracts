require('dotenv').config();

const accounts = process.env.DEPLOYER_PRIVATE_KEY
    ? [process.env.DEPLOYER_PRIVATE_KEY]
    : ['1495992B2A5CC4DD53E231157BBF401329BD1B7EE355CEAB55A791398921CA17'];
const gasPrice = process.env.GAS_PRICE
    ? parseInt(process.env.GAS_PRICE)
    : 'auto';

export const CONFIG = {
    networks: {
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
        hardhat: {
            // forking: {
            //     url: process.env.MAINNET_RPC,
            //     blockNumber: 18871509,
            // },
        },
    },
    mocha: {
        enableTimeouts: false,
        before_timeout: 120000,
    },
    gasReporter: {
        enabled: true,
        currency: 'USD',
        token: 'ETH',
        coinmarketcap: process.env.COINMARKETCAP_API_KEY,
    },
    namedAccounts: {
        deployer: {
            default: 0,
        },
        treasury: {
            goerli: '0x064B9a8cd35ad4dB117617A3773F8129E9515967',
            mainnet: '0x00Fd4edEd5BB37d19F98Ab49722Ef51E84745928',
        },
        operator: {
            goerli: '0x064B9a8cd35ad4dB117617A3773F8129E9515967',
            mainnet: '0x078dc682083132b4E86731062FCF95A729Bac067',
        },
        governance: {
            goerli: 0,
            mainnet: '0x03D7aaa453D9e7048101d425e73848e16c534DFD',
        },
        elPodManager: {
            goerli: '0xa286b84C96aF280a49Fe1F40B9627C2A2827df41',
            mainnet: '0x91E677b07F7AF907ec9a428aafA9fc14a0d3A338',
        },
        elDelegationManager: {
            goerli: '0x1b7b8F6b258f95Cf9596EabB9aa18B62940Eb0a8',
            mainnet: '0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A',
        },
        proxyAdminTimelock: {
            goerli: 0,
            mainnet: 'fill',
        },
    },
    verify: {
        etherscan: {
            apiKey: process.env.ETHERSCAN_API_KEY,
        },
        sourcify: {
            enabled: false,
        },
    },
    etherscan: {
        apiKey: process.env.ETHERSCAN_API_KEY,
    },
    sourcify: {
        enabled: false,
    },
};
