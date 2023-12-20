// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakingpool_v1

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/genesislrt/contracts/abigen/generated"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"DistributeGasLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"PendingUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"PoolOnGoing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"claimers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"RewardsDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"distributePendingRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCert\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDistributeGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEigenPodManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFreeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinUnstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"getPendingRequestsOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"getPendingUnstakesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPendingUnstakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingConfig\",\"name\":\"stakingConfig\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"distributeGasLimit\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"deposit_data_root\",\"type\":\"bytes32\"}],\"name\":\"pushToBeacon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"deposit_data_roots\",\"type\":\"bytes32[]\"}],\"name\":\"pushToBeaconMulti\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setDistributeGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeCerts\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"unstakeCerts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611dad806100206000396000f3fe6080604052600436106100f35760003560e01c80636d186d7b1161008a578063c664d38b11610059578063c664d38b1461027d578063cd6dc68714610292578063d0652e42146102b2578063e93c4f15146102c757600080fd5b80636d186d7b14610220578063ac76d45014610235578063afbcc4481461023d578063be9631e81461025d57600080fd5b8063183cdf24116100c6578063183cdf2414610188578063409cb613146101b557806346faa33c146101d557806356a3b5fa1461020b57600080fd5b806301ff923c146100f8578063034c44541461012e578063042fc3a614610151578063126d5df614610173575b600080fd5b34801561010457600080fd5b5061011861011336600461189c565b6102da565b60405161012591906118f4565b60405180910390f35b34801561013a57600080fd5b506101436103f8565b604051908152602001610125565b34801561015d57600080fd5b5061017161016c366004611907565b61046a565b005b34801561017f57600080fd5b50603654610143565b34801561019457600080fd5b5061019d610597565b6040516001600160a01b039091168152602001610125565b3480156101c157600080fd5b506101716101d0366004611920565b610605565b3480156101e157600080fd5b506101436101f036600461189c565b6001600160a01b031660009081526038602052604090205490565b34801561021757600080fd5b506101436108cd565b34801561022c57600080fd5b50610171610917565b610171610c73565b34801561024957600080fd5b50610171610258366004611998565b610e90565b34801561026957600080fd5b50610171610278366004611a74565b6111af565b34801561028957600080fd5b5061019d6113c7565b34801561029e57600080fd5b506101716102ad366004611920565b611411565b3480156102be57600080fd5b50603454610143565b3480156102d357600080fd5b5047610143565b60355460375460609160009182916102f191611afe565b67ffffffffffffffff81111561030957610309611b17565b604051908082528060200260200182016040528015610332578160200160208202803683370190505b506035549091505b6037548110156103d457846001600160a01b03166037828154811061036157610361611b2d565b6000918252602090912001546001600160a01b0316036103c2576039818154811061038e5761038e611b2d565b90600052602060002001548284815181106103ab576103ab611b2d565b60209081029190910101526103bf83611b43565b92505b806103cc81611b43565b91505061033a565b5060008282516103e49190611afe565b905080156103f0578282525b509392505050565b6033546040805162d3111560e21b815290516000926001600160a01b03169163034c44549160048083019260209291908290030181865afa158015610441573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104659190611b5c565b905090565b603360009054906101000a90046001600160a01b03166001600160a01b031663732524946040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e19190611b75565b6001600160a01b0316336001600160a01b0316146105525760405162461bcd60e51b8152602060048201526024808201527f5374616b696e67506f6f6c3a206f6e6c7920676f7665726e616e636520616c6c6044820152631bddd95960e21b60648201526084015b60405180910390fd5b603480549082905560408051828152602081018490527f3bb2da990d30b0bc98e39d632b60814d66b3bae55947927dec7a75719de577de910160405180910390a15050565b60335460408051635eaad42760e11b815290516000926001600160a01b03169163bd55a84e9160048083019260209291908290030181865afa1580156105e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104659190611b75565b60335460408051635eaad42760e11b8152905133926000926001600160a01b039091169163bd55a84e916004808201926020929091908290030181865afa158015610654573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106789190611b75565b604051636c58d43d60e01b8152600481018590529091506000906001600160a01b03831690636c58d43d90602401602060405180830381865afa1580156106c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e79190611b5c565b9050603360009054906101000a90046001600160a01b03166001600160a01b031663034c44546040518163ffffffff1660e01b8152600401602060405180830381865afa15801561073c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107609190611b5c565b81101561077f5760405162461bcd60e51b815260040161054990611b92565b6040516370a0823160e01b81526001600160a01b0384811660048301528591908416906370a0823190602401602060405180830381865afa1580156107c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ec9190611b5c565b10156108585760405162461bcd60e51b815260206004820152603560248201527f5374616b696e67506f6f6c3a2063616e6e6f7420756e7374616b65206d6f7265604482015274207468616e2068617665206f6e206164647265737360581b6064820152608401610549565b604051632770a7eb60e21b81526001600160a01b03848116600483015260248201869052831690639dc29fac90604401600060405180830381600087803b1580156108a257600080fd5b505af11580156108b6573d6000803e3d6000fd5b505050506108c6838686846115f4565b5050505050565b60335460408051632b51dafd60e11b815290516000926001600160a01b0316916356a3b5fa9160048083019260209291908290030181865afa158015610441573d6000803e3d6000fd5b60006034541161097e5760405162461bcd60e51b815260206004820152602c60248201527f5374616b696e67506f6f6c3a20444953545249425554455f4741535f4c494d4960448201526b15081a5cc81b9bdd081cd95d60a21b6064820152608401610549565b60355460375447916000916109939190611afe565b67ffffffffffffffff8111156109ab576109ab611b17565b6040519080825280602002602001820160405280156109d4578160200160208202803683370190505b506035546037549192506000916109eb9190611afe565b67ffffffffffffffff811115610a0357610a03611b17565b604051908082528060200260200182016040528015610a2c578160200160208202803683370190505b506035549091506000905b60375481108015610a485750600085115b8015610a5557506034545a115b15610c0e57600060378281548110610a6f57610a6f611b2d565b6000918252602082200154603980546001600160a01b0390921693509084908110610a9c57610a9c611b2d565b60009182526020909120015490506001600160a01b0382161580610abe575080155b15610ad557610acc83611b43565b92505050610a37565b80871015610ae4575050610c0e565b6001600160a01b03821660009081526038602052604081208054839290610b0c908490611afe565b925050819055508060366000828254610b259190611afe565b90915550610b3590508188611afe565b965060378381548110610b4a57610b4a611b2d565b600091825260209091200180546001600160a01b03191690556039805484908110610b7757610b77611b2d565b6000918252602082200155610b8b83611b43565b92506000610b9b8383600161176c565b905080610baa57505050610a37565b82878681518110610bbd57610bbd611b2d565b60200260200101906001600160a01b031690816001600160a01b03168152505081868681518110610bf057610bf0611b2d565b6020908102919091010152610c0485611b43565b9450505050610a37565b60358190558351600090610c23908490611afe565b90508015610c32578285528284525b7fe69d325558610ba73c441901deb46d7f251108348dc5dc9447e8866774c12edc8585604051610c63929190611be4565b60405180910390a1505050505050565b610c7b6117e9565b60335460408051632b51dafd60e11b8152905134926001600160a01b0316916356a3b5fa9160048083019260209291908290030181865afa158015610cc4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce89190611b5c565b811015610d075760405162461bcd60e51b815260040161054990611b92565b60335460408051635eaad42760e11b815290516000926001600160a01b03169163bd55a84e9160048083019260209291908290030181865afa158015610d51573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d759190611b75565b604051635361637360e01b8152600481018490529091506000906001600160a01b03831690635361637390602401602060405180830381865afa158015610dc0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610de49190611b5c565b6040516340c10f1960e01b8152336004820152602481018290529091506001600160a01b038316906340c10f1990604401600060405180830381600087803b158015610e2f57600080fd5b505af1158015610e43573d6000803e3d6000fd5b505060408051868152602081018590523393507f1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee9092500160405180910390a2505050610e8e60018055565b565b603360009054906101000a90046001600160a01b03166001600160a01b0316632ec338ba6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ee3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f079190611b75565b6001600160a01b0316336001600160a01b031614610f375760405162461bcd60e51b815260040161054990611c44565b848381148015610f4657508382145b610f9c5760405162461bcd60e51b815260206004820152602160248201527f5374616b696e67506f6f6c3a206c656e67746820617265206e6f7420657175616044820152601b60fa1b6064820152608401610549565b610faf816801bc16d674ec800000611c87565b471015610ffa5760405162461bcd60e51b81526020600482015260196024820152780e0cadcc8d2dcce40cae8d0cae4e640dcdee840cadcdeeaced603b1b6044820152606401610549565b6033546040805163476c397360e11b815290516000926001600160a01b031691638ed872e69160048083019260209291908290030181865afa158015611044573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110689190611b75565b905060005b828110156111a457816001600160a01b0316639b4e46348a8a8481811061109657611096611b2d565b90506020028101906110a89190611c9e565b8a8a868181106110ba576110ba611b2d565b90506020028101906110cc9190611c9e565b8a8a888181106110de576110de611b2d565b905060200201356040518663ffffffff1660e01b8152600401611105959493929190611d0e565b600060405180830381600087803b15801561111f57600080fd5b505af1158015611133573d6000803e3d6000fd5b505050507f09bfa94ea4cf30558b9b9914b0029f04369c211b0e1f965f9fa7a29536b300df89898381811061116a5761116a611b2d565b905060200281019061117c9190611c9e565b60405161118a929190611d48565b60405180910390a18061119c81611b43565b91505061106d565b505050505050505050565b603360009054906101000a90046001600160a01b03166001600160a01b0316632ec338ba6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611202573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112269190611b75565b6001600160a01b0316336001600160a01b0316146112565760405162461bcd60e51b815260040161054990611c44565b6801bc16d674ec8000004710156112ab5760405162461bcd60e51b81526020600482015260196024820152780e0cadcc8d2dcce40cae8d0cae4e640dcdee840cadcdeeaced603b1b6044820152606401610549565b603360009054906101000a90046001600160a01b03166001600160a01b0316638ed872e66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113229190611b75565b6001600160a01b0316639b4e463486868686866040518663ffffffff1660e01b8152600401611355959493929190611d0e565b600060405180830381600087803b15801561136f57600080fd5b505af1158015611383573d6000803e3d6000fd5b505050507f09bfa94ea4cf30558b9b9914b0029f04369c211b0e1f965f9fa7a29536b300df85856040516113b8929190611d48565b60405180910390a15050505050565b6033546040805163476c397360e11b815290516000926001600160a01b031691638ed872e69160048083019260209291908290030181865afa1580156105e1573d6000803e3d6000fd5b600054610100900460ff16158080156114315750600054600160ff909116105b8061144b5750303b15801561144b575060005460ff166001145b6114ae5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610549565b6000805460ff1916600117905580156114d1576000805461ff0019166101001790555b603380546001600160a01b0319166001600160a01b0385169081179091556040805163476c397360e11b81529051638ed872e6916004808201926020929091908290030181865afa15801561152a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061154e9190611b75565b6001600160a01b03166384d810626040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561158857600080fd5b505af115801561159c573d6000803e3d6000fd5b505050506115a982611842565b80156115ef576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b801580159061160b57506001600160a01b03831615155b6116695760405162461bcd60e51b815260206004820152602960248201527f4c6971756964546f6b656e5374616b696e67506f6f6c3a207a65726f20696e7060448201526875742076616c75657360b81b6064820152608401610549565b806036600082825461167b9190611d64565b90915550506037805460018082019092557f42a7b7dd785cd69714a189dffb3fd7d7174edc9ece837694ce50f7078f7c31ae0180546001600160a01b0319166001600160a01b0386169081179091556039805492830190557fdc16fef70f8d5ddbc01ee3d903d1e69c18a3c7be080eb86a81e0578814ee58d390910182905560009081526038602052604081208054839290611718908490611d64565b909155505060408051828152602081018490526001600160a01b0380861692908716917fe8f73d529f5ced08581a2c18456a6530dbd0dddf94d8c98e0ab8f9883e2f4482910160405180910390a350505050565b60008381831561178c576000806000808886612710f192506117e2915050565b6040516001600160a01b038316908690600081818185875af1925050503d80600081146117d5576040519150601f19603f3d011682016040523d82523d6000602084013e6117da565b606091505b509093505050505b9392505050565b60026001540361183b5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610549565b6002600155565b60348190556040805160008152602081018390527f3bb2da990d30b0bc98e39d632b60814d66b3bae55947927dec7a75719de577de910160405180910390a150565b6001600160a01b038116811461189957600080fd5b50565b6000602082840312156118ae57600080fd5b81356117e281611884565b600081518084526020808501945080840160005b838110156118e9578151875295820195908201906001016118cd565b509495945050505050565b6020815260006117e260208301846118b9565b60006020828403121561191957600080fd5b5035919050565b6000806040838503121561193357600080fd5b823561193e81611884565b946020939093013593505050565b60008083601f84011261195e57600080fd5b50813567ffffffffffffffff81111561197657600080fd5b6020830191508360208260051b850101111561199157600080fd5b9250929050565b600080600080600080606087890312156119b157600080fd5b863567ffffffffffffffff808211156119c957600080fd5b6119d58a838b0161194c565b909850965060208901359150808211156119ee57600080fd5b6119fa8a838b0161194c565b90965094506040890135915080821115611a1357600080fd5b50611a2089828a0161194c565b979a9699509497509295939492505050565b60008083601f840112611a4457600080fd5b50813567ffffffffffffffff811115611a5c57600080fd5b60208301915083602082850101111561199157600080fd5b600080600080600060608688031215611a8c57600080fd5b853567ffffffffffffffff80821115611aa457600080fd5b611ab089838a01611a32565b90975095506020880135915080821115611ac957600080fd5b50611ad688828901611a32565b96999598509660400135949350505050565b634e487b7160e01b600052601160045260246000fd5b81810381811115611b1157611b11611ae8565b92915050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600060018201611b5557611b55611ae8565b5060010190565b600060208284031215611b6e57600080fd5b5051919050565b600060208284031215611b8757600080fd5b81516117e281611884565b60208082526032908201527f5374616b696e67506f6f6c3a2076616c7565206d757374206265206772656174604082015271195c881d1a185b881b5a5b88185b5bdd5b9d60721b606082015260800190565b604080825283519082018190526000906020906060840190828701845b82811015611c265781516001600160a01b031684529284019290840190600101611c01565b50505083810382850152611c3a81866118b9565b9695505050505050565b60208082526023908201527f5374616b696e67506f6f6c3a206f6e6c7920636f6e73656e73757320616c6c6f6040820152621dd95960ea1b606082015260800190565b8082028115828204841417611b1157611b11611ae8565b6000808335601e19843603018112611cb557600080fd5b83018035915067ffffffffffffffff821115611cd057600080fd5b60200191503681900382131561199157600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b606081526000611d22606083018789611ce5565b8281036020840152611d35818688611ce5565b9150508260408301529695505050505050565b602081526000611d5c602083018486611ce5565b949350505050565b80820180821115611b1157611b11611ae856fea2646970667358221220233b5b6efc5aade38bf5c35a7fdc658f6650451696f125613d8e6784a76d227f64736f6c63430008130033",
}

var ContractABI = ContractMetaData.ABI

var ContractBin = ContractMetaData.Bin

func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

type Contract struct {
	address common.Address
	abi     abi.ABI
	ContractCaller
	ContractTransactor
	ContractFilterer
}

type ContractCaller struct {
	contract *bind.BoundContract
}

type ContractTransactor struct {
	contract *bind.BoundContract
}

type ContractFilterer struct {
	contract *bind.BoundContract
}

type ContractSession struct {
	Contract     *Contract
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ContractCallerSession struct {
	Contract *ContractCaller
	CallOpts bind.CallOpts
}

type ContractTransactorSession struct {
	Contract     *ContractTransactor
	TransactOpts bind.TransactOpts
}

type ContractRaw struct {
	Contract *Contract
}

type ContractCallerRaw struct {
	Contract *ContractCaller
}

type ContractTransactorRaw struct {
	Contract *ContractTransactor
}

func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	abi, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{address: address, abi: abi, ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

func (_Contract *ContractCaller) GetCert(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCert")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetCert() (common.Address, error) {
	return _Contract.Contract.GetCert(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetCert() (common.Address, error) {
	return _Contract.Contract.GetCert(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetDistributeGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDistributeGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) GetDistributeGasLimit() (*big.Int, error) {
	return _Contract.Contract.GetDistributeGasLimit(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetDistributeGasLimit() (*big.Int, error) {
	return _Contract.Contract.GetDistributeGasLimit(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetEigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetEigenPodManager() (common.Address, error) {
	return _Contract.Contract.GetEigenPodManager(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetEigenPodManager() (common.Address, error) {
	return _Contract.Contract.GetEigenPodManager(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetFreeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getFreeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) GetFreeBalance() (*big.Int, error) {
	return _Contract.Contract.GetFreeBalance(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetFreeBalance() (*big.Int, error) {
	return _Contract.Contract.GetFreeBalance(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetMinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getMinStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) GetMinStake() (*big.Int, error) {
	return _Contract.Contract.GetMinStake(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetMinStake() (*big.Int, error) {
	return _Contract.Contract.GetMinStake(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetMinUnstake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getMinUnstake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) GetMinUnstake() (*big.Int, error) {
	return _Contract.Contract.GetMinUnstake(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetMinUnstake() (*big.Int, error) {
	return _Contract.Contract.GetMinUnstake(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetPendingRequestsOf(opts *bind.CallOpts, claimer common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPendingRequestsOf", claimer)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

func (_Contract *ContractSession) GetPendingRequestsOf(claimer common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetPendingRequestsOf(&_Contract.CallOpts, claimer)
}

func (_Contract *ContractCallerSession) GetPendingRequestsOf(claimer common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetPendingRequestsOf(&_Contract.CallOpts, claimer)
}

func (_Contract *ContractCaller) GetPendingUnstakesOf(opts *bind.CallOpts, claimer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPendingUnstakesOf", claimer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) GetPendingUnstakesOf(claimer common.Address) (*big.Int, error) {
	return _Contract.Contract.GetPendingUnstakesOf(&_Contract.CallOpts, claimer)
}

func (_Contract *ContractCallerSession) GetPendingUnstakesOf(claimer common.Address) (*big.Int, error) {
	return _Contract.Contract.GetPendingUnstakesOf(&_Contract.CallOpts, claimer)
}

func (_Contract *ContractCaller) GetTotalPendingUnstakes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTotalPendingUnstakes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) GetTotalPendingUnstakes() (*big.Int, error) {
	return _Contract.Contract.GetTotalPendingUnstakes(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetTotalPendingUnstakes() (*big.Int, error) {
	return _Contract.Contract.GetTotalPendingUnstakes(&_Contract.CallOpts)
}

func (_Contract *ContractTransactor) DistributePendingRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "distributePendingRewards")
}

func (_Contract *ContractSession) DistributePendingRewards() (*types.Transaction, error) {
	return _Contract.Contract.DistributePendingRewards(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactorSession) DistributePendingRewards() (*types.Transaction, error) {
	return _Contract.Contract.DistributePendingRewards(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, stakingConfig common.Address, distributeGasLimit *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", stakingConfig, distributeGasLimit)
}

func (_Contract *ContractSession) Initialize(stakingConfig common.Address, distributeGasLimit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, stakingConfig, distributeGasLimit)
}

func (_Contract *ContractTransactorSession) Initialize(stakingConfig common.Address, distributeGasLimit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, stakingConfig, distributeGasLimit)
}

func (_Contract *ContractTransactor) PushToBeacon(opts *bind.TransactOpts, pubkey []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "pushToBeacon", pubkey, signature, deposit_data_root)
}

func (_Contract *ContractSession) PushToBeacon(pubkey []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.PushToBeacon(&_Contract.TransactOpts, pubkey, signature, deposit_data_root)
}

func (_Contract *ContractTransactorSession) PushToBeacon(pubkey []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.PushToBeacon(&_Contract.TransactOpts, pubkey, signature, deposit_data_root)
}

func (_Contract *ContractTransactor) PushToBeaconMulti(opts *bind.TransactOpts, pubkeys [][]byte, signatures [][]byte, deposit_data_roots [][32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "pushToBeaconMulti", pubkeys, signatures, deposit_data_roots)
}

func (_Contract *ContractSession) PushToBeaconMulti(pubkeys [][]byte, signatures [][]byte, deposit_data_roots [][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.PushToBeaconMulti(&_Contract.TransactOpts, pubkeys, signatures, deposit_data_roots)
}

func (_Contract *ContractTransactorSession) PushToBeaconMulti(pubkeys [][]byte, signatures [][]byte, deposit_data_roots [][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.PushToBeaconMulti(&_Contract.TransactOpts, pubkeys, signatures, deposit_data_roots)
}

func (_Contract *ContractTransactor) SetDistributeGasLimit(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDistributeGasLimit", newValue)
}

func (_Contract *ContractSession) SetDistributeGasLimit(newValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetDistributeGasLimit(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetDistributeGasLimit(newValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetDistributeGasLimit(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) StakeCerts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "stakeCerts")
}

func (_Contract *ContractSession) StakeCerts() (*types.Transaction, error) {
	return _Contract.Contract.StakeCerts(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactorSession) StakeCerts() (*types.Transaction, error) {
	return _Contract.Contract.StakeCerts(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactor) UnstakeCerts(opts *bind.TransactOpts, receiverAddress common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unstakeCerts", receiverAddress, shares)
}

func (_Contract *ContractSession) UnstakeCerts(receiverAddress common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UnstakeCerts(&_Contract.TransactOpts, receiverAddress, shares)
}

func (_Contract *ContractTransactorSession) UnstakeCerts(receiverAddress common.Address, shares *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UnstakeCerts(&_Contract.TransactOpts, receiverAddress, shares)
}

type ContractDistributeGasLimitChangedIterator struct {
	Event *ContractDistributeGasLimitChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractDistributeGasLimitChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDistributeGasLimitChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(ContractDistributeGasLimitChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *ContractDistributeGasLimitChangedIterator) Error() error {
	return it.fail
}

func (it *ContractDistributeGasLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractDistributeGasLimitChanged struct {
	PrevValue *big.Int
	NewValue  *big.Int
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterDistributeGasLimitChanged(opts *bind.FilterOpts) (*ContractDistributeGasLimitChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DistributeGasLimitChanged")
	if err != nil {
		return nil, err
	}
	return &ContractDistributeGasLimitChangedIterator{contract: _Contract.contract, event: "DistributeGasLimitChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchDistributeGasLimitChanged(opts *bind.WatchOpts, sink chan<- *ContractDistributeGasLimitChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DistributeGasLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractDistributeGasLimitChanged)
				if err := _Contract.contract.UnpackLog(event, "DistributeGasLimitChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Contract *ContractFilterer) ParseDistributeGasLimitChanged(log types.Log) (*ContractDistributeGasLimitChanged, error) {
	event := new(ContractDistributeGasLimitChanged)
	if err := _Contract.contract.UnpackLog(event, "DistributeGasLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractInitializedIterator struct {
	Event *ContractInitialized

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractInitializedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(ContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractInitialized struct {
	Version uint8
	Raw     types.Log
}

func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractPendingUnstakeIterator struct {
	Event *ContractPendingUnstake

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractPendingUnstakeIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPendingUnstake)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(ContractPendingUnstake)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *ContractPendingUnstakeIterator) Error() error {
	return it.fail
}

func (it *ContractPendingUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractPendingUnstake struct {
	OwnerAddress    common.Address
	ReceiverAddress common.Address
	Amount          *big.Int
	Shares          *big.Int
	Raw             types.Log
}

func (_Contract *ContractFilterer) FilterPendingUnstake(opts *bind.FilterOpts, ownerAddress []common.Address, receiverAddress []common.Address) (*ContractPendingUnstakeIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}
	var receiverAddressRule []interface{}
	for _, receiverAddressItem := range receiverAddress {
		receiverAddressRule = append(receiverAddressRule, receiverAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PendingUnstake", ownerAddressRule, receiverAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractPendingUnstakeIterator{contract: _Contract.contract, event: "PendingUnstake", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchPendingUnstake(opts *bind.WatchOpts, sink chan<- *ContractPendingUnstake, ownerAddress []common.Address, receiverAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}
	var receiverAddressRule []interface{}
	for _, receiverAddressItem := range receiverAddress {
		receiverAddressRule = append(receiverAddressRule, receiverAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PendingUnstake", ownerAddressRule, receiverAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractPendingUnstake)
				if err := _Contract.contract.UnpackLog(event, "PendingUnstake", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Contract *ContractFilterer) ParsePendingUnstake(log types.Log) (*ContractPendingUnstake, error) {
	event := new(ContractPendingUnstake)
	if err := _Contract.contract.UnpackLog(event, "PendingUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractPoolOnGoingIterator struct {
	Event *ContractPoolOnGoing

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractPoolOnGoingIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPoolOnGoing)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(ContractPoolOnGoing)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *ContractPoolOnGoingIterator) Error() error {
	return it.fail
}

func (it *ContractPoolOnGoingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractPoolOnGoing struct {
	Pubkey []byte
	Raw    types.Log
}

func (_Contract *ContractFilterer) FilterPoolOnGoing(opts *bind.FilterOpts) (*ContractPoolOnGoingIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PoolOnGoing")
	if err != nil {
		return nil, err
	}
	return &ContractPoolOnGoingIterator{contract: _Contract.contract, event: "PoolOnGoing", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchPoolOnGoing(opts *bind.WatchOpts, sink chan<- *ContractPoolOnGoing) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PoolOnGoing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractPoolOnGoing)
				if err := _Contract.contract.UnpackLog(event, "PoolOnGoing", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Contract *ContractFilterer) ParsePoolOnGoing(log types.Log) (*ContractPoolOnGoing, error) {
	event := new(ContractPoolOnGoing)
	if err := _Contract.contract.UnpackLog(event, "PoolOnGoing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractReceivedIterator struct {
	Event *ContractReceived

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractReceivedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(ContractReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *ContractReceivedIterator) Error() error {
	return it.fail
}

func (it *ContractReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractReceived struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log
}

func (_Contract *ContractFilterer) FilterReceived(opts *bind.FilterOpts, from []common.Address) (*ContractReceivedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Received", fromRule)
	if err != nil {
		return nil, err
	}
	return &ContractReceivedIterator{contract: _Contract.contract, event: "Received", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *ContractReceived, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Received", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractReceived)
				if err := _Contract.contract.UnpackLog(event, "Received", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Contract *ContractFilterer) ParseReceived(log types.Log) (*ContractReceived, error) {
	event := new(ContractReceived)
	if err := _Contract.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractRewardsDistributedIterator struct {
	Event *ContractRewardsDistributed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractRewardsDistributedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsDistributed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(ContractRewardsDistributed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *ContractRewardsDistributedIterator) Error() error {
	return it.fail
}

func (it *ContractRewardsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractRewardsDistributed struct {
	Claimers []common.Address
	Amounts  []*big.Int
	Raw      types.Log
}

func (_Contract *ContractFilterer) FilterRewardsDistributed(opts *bind.FilterOpts) (*ContractRewardsDistributedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RewardsDistributed")
	if err != nil {
		return nil, err
	}
	return &ContractRewardsDistributedIterator{contract: _Contract.contract, event: "RewardsDistributed", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchRewardsDistributed(opts *bind.WatchOpts, sink chan<- *ContractRewardsDistributed) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RewardsDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractRewardsDistributed)
				if err := _Contract.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Contract *ContractFilterer) ParseRewardsDistributed(log types.Log) (*ContractRewardsDistributed, error) {
	event := new(ContractRewardsDistributed)
	if err := _Contract.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractStakedIterator struct {
	Event *ContractStaked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractStakedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(ContractStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *ContractStakedIterator) Error() error {
	return it.fail
}

func (it *ContractStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractStaked struct {
	Staker common.Address
	Amount *big.Int
	Shares *big.Int
	Raw    types.Log
}

func (_Contract *ContractFilterer) FilterStaked(opts *bind.FilterOpts, staker []common.Address) (*ContractStakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return &ContractStakedIterator{contract: _Contract.contract, event: "Staked", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *ContractStaked, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractStaked)
				if err := _Contract.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_Contract *ContractFilterer) ParseStaked(log types.Log) (*ContractStaked, error) {
	event := new(ContractStaked)
	if err := _Contract.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Contract *Contract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Contract.abi.Events["DistributeGasLimitChanged"].ID:
		return _Contract.ParseDistributeGasLimitChanged(log)
	case _Contract.abi.Events["Initialized"].ID:
		return _Contract.ParseInitialized(log)
	case _Contract.abi.Events["PendingUnstake"].ID:
		return _Contract.ParsePendingUnstake(log)
	case _Contract.abi.Events["PoolOnGoing"].ID:
		return _Contract.ParsePoolOnGoing(log)
	case _Contract.abi.Events["Received"].ID:
		return _Contract.ParseReceived(log)
	case _Contract.abi.Events["RewardsDistributed"].ID:
		return _Contract.ParseRewardsDistributed(log)
	case _Contract.abi.Events["Staked"].ID:
		return _Contract.ParseStaked(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ContractDistributeGasLimitChanged) Topic() common.Hash {
	return common.HexToHash("0x3bb2da990d30b0bc98e39d632b60814d66b3bae55947927dec7a75719de577de")
}

func (ContractInitialized) Topic() common.Hash {
	return common.HexToHash("0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498")
}

func (ContractPendingUnstake) Topic() common.Hash {
	return common.HexToHash("0xe8f73d529f5ced08581a2c18456a6530dbd0dddf94d8c98e0ab8f9883e2f4482")
}

func (ContractPoolOnGoing) Topic() common.Hash {
	return common.HexToHash("0x09bfa94ea4cf30558b9b9914b0029f04369c211b0e1f965f9fa7a29536b300df")
}

func (ContractReceived) Topic() common.Hash {
	return common.HexToHash("0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874")
}

func (ContractRewardsDistributed) Topic() common.Hash {
	return common.HexToHash("0xe69d325558610ba73c441901deb46d7f251108348dc5dc9447e8866774c12edc")
}

func (ContractStaked) Topic() common.Hash {
	return common.HexToHash("0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90")
}

func (_Contract *Contract) Address() common.Address {
	return _Contract.address
}

type ContractInterface interface {
	GetCert(opts *bind.CallOpts) (common.Address, error)

	GetDistributeGasLimit(opts *bind.CallOpts) (*big.Int, error)

	GetEigenPodManager(opts *bind.CallOpts) (common.Address, error)

	GetFreeBalance(opts *bind.CallOpts) (*big.Int, error)

	GetMinStake(opts *bind.CallOpts) (*big.Int, error)

	GetMinUnstake(opts *bind.CallOpts) (*big.Int, error)

	GetPendingRequestsOf(opts *bind.CallOpts, claimer common.Address) ([]*big.Int, error)

	GetPendingUnstakesOf(opts *bind.CallOpts, claimer common.Address) (*big.Int, error)

	GetTotalPendingUnstakes(opts *bind.CallOpts) (*big.Int, error)

	DistributePendingRewards(opts *bind.TransactOpts) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, stakingConfig common.Address, distributeGasLimit *big.Int) (*types.Transaction, error)

	PushToBeacon(opts *bind.TransactOpts, pubkey []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error)

	PushToBeaconMulti(opts *bind.TransactOpts, pubkeys [][]byte, signatures [][]byte, deposit_data_roots [][32]byte) (*types.Transaction, error)

	SetDistributeGasLimit(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error)

	StakeCerts(opts *bind.TransactOpts) (*types.Transaction, error)

	UnstakeCerts(opts *bind.TransactOpts, receiverAddress common.Address, shares *big.Int) (*types.Transaction, error)

	FilterDistributeGasLimitChanged(opts *bind.FilterOpts) (*ContractDistributeGasLimitChangedIterator, error)

	WatchDistributeGasLimitChanged(opts *bind.WatchOpts, sink chan<- *ContractDistributeGasLimitChanged) (event.Subscription, error)

	ParseDistributeGasLimitChanged(log types.Log) (*ContractDistributeGasLimitChanged, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*ContractInitialized, error)

	FilterPendingUnstake(opts *bind.FilterOpts, ownerAddress []common.Address, receiverAddress []common.Address) (*ContractPendingUnstakeIterator, error)

	WatchPendingUnstake(opts *bind.WatchOpts, sink chan<- *ContractPendingUnstake, ownerAddress []common.Address, receiverAddress []common.Address) (event.Subscription, error)

	ParsePendingUnstake(log types.Log) (*ContractPendingUnstake, error)

	FilterPoolOnGoing(opts *bind.FilterOpts) (*ContractPoolOnGoingIterator, error)

	WatchPoolOnGoing(opts *bind.WatchOpts, sink chan<- *ContractPoolOnGoing) (event.Subscription, error)

	ParsePoolOnGoing(log types.Log) (*ContractPoolOnGoing, error)

	FilterReceived(opts *bind.FilterOpts, from []common.Address) (*ContractReceivedIterator, error)

	WatchReceived(opts *bind.WatchOpts, sink chan<- *ContractReceived, from []common.Address) (event.Subscription, error)

	ParseReceived(log types.Log) (*ContractReceived, error)

	FilterRewardsDistributed(opts *bind.FilterOpts) (*ContractRewardsDistributedIterator, error)

	WatchRewardsDistributed(opts *bind.WatchOpts, sink chan<- *ContractRewardsDistributed) (event.Subscription, error)

	ParseRewardsDistributed(log types.Log) (*ContractRewardsDistributed, error)

	FilterStaked(opts *bind.FilterOpts, staker []common.Address) (*ContractStakedIterator, error)

	WatchStaked(opts *bind.WatchOpts, sink chan<- *ContractStaked, staker []common.Address) (event.Subscription, error)

	ParseStaked(log types.Log) (*ContractStaked, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
