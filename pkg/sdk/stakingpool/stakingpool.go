// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakingpool

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/TagusLabs/genesis-smart-contracts/abigen/generated"
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"DistributeGasLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"PendingUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"PoolOnGoing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getCert\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDistributeGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEigenPodManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFreeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinUnstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"getPendingRequestsOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"}],\"name\":\"getPendingUnstakesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPendingUnstakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingConfig\",\"name\":\"stakingConfig\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"distributeGasLimit\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"deposit_data_root\",\"type\":\"bytes32\"}],\"name\":\"pushToBeacon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"deposit_data_roots\",\"type\":\"bytes32[]\"}],\"name\":\"pushToBeaconMulti\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setDistributeGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeCerts\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"unstakeCerts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611951806100206000396000f3fe6080604052600436106100e85760003560e01c806356a3b5fa1161008a578063c664d38b11610059578063c664d38b1461025d578063cd6dc68714610272578063d0652e4214610292578063e93c4f15146102a757600080fd5b806356a3b5fa14610200578063ac76d45014610215578063afbcc4481461021d578063be9631e81461023d57600080fd5b8063126d5df6116100c6578063126d5df614610168578063183cdf241461017d578063409cb613146101aa57806346faa33c146101ca57600080fd5b806301ff923c146100ed578063034c445414610123578063042fc3a614610146575b600080fd5b3480156100f957600080fd5b5061010d6101083660046114a3565b6102ba565b60405161011a91906114c7565b60405180910390f35b34801561012f57600080fd5b506101386103d8565b60405190815260200161011a565b34801561015257600080fd5b5061016661016136600461150b565b61044a565b005b34801561017457600080fd5b50603654610138565b34801561018957600080fd5b50610192610577565b6040516001600160a01b03909116815260200161011a565b3480156101b657600080fd5b506101666101c5366004611524565b6105e5565b3480156101d657600080fd5b506101386101e53660046114a3565b6001600160a01b031660009081526038602052604090205490565b34801561020c57600080fd5b506101386108ad565b6101666108f7565b34801561022957600080fd5b5061016661023836600461159c565b610b14565b34801561024957600080fd5b50610166610258366004611678565b610e33565b34801561026957600080fd5b5061019261104b565b34801561027e57600080fd5b5061016661028d366004611524565b611095565b34801561029e57600080fd5b50603454610138565b3480156102b357600080fd5b5047610138565b60355460375460609160009182916102d191611702565b67ffffffffffffffff8111156102e9576102e961171b565b604051908082528060200260200182016040528015610312578160200160208202803683370190505b506035549091505b6037548110156103b457846001600160a01b03166037828154811061034157610341611731565b6000918252602090912001546001600160a01b0316036103a2576039818154811061036e5761036e611731565b906000526020600020015482848151811061038b5761038b611731565b602090810291909101015261039f83611747565b92505b806103ac81611747565b91505061031a565b5060008282516103c49190611702565b905080156103d0578282525b509392505050565b6033546040805162d3111560e21b815290516000926001600160a01b03169163034c44549160048083019260209291908290030181865afa158015610421573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104459190611760565b905090565b603360009054906101000a90046001600160a01b03166001600160a01b031663732524946040518163ffffffff1660e01b8152600401602060405180830381865afa15801561049d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104c19190611779565b6001600160a01b0316336001600160a01b0316146105325760405162461bcd60e51b8152602060048201526024808201527f5374616b696e67506f6f6c3a206f6e6c7920676f7665726e616e636520616c6c6044820152631bddd95960e21b60648201526084015b60405180910390fd5b603480549082905560408051828152602081018490527f3bb2da990d30b0bc98e39d632b60814d66b3bae55947927dec7a75719de577de910160405180910390a15050565b60335460408051635eaad42760e11b815290516000926001600160a01b03169163bd55a84e9160048083019260209291908290030181865afa1580156105c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104459190611779565b60335460408051635eaad42760e11b8152905133926000926001600160a01b039091169163bd55a84e916004808201926020929091908290030181865afa158015610634573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106589190611779565b604051636c58d43d60e01b8152600481018590529091506000906001600160a01b03831690636c58d43d90602401602060405180830381865afa1580156106a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106c79190611760565b9050603360009054906101000a90046001600160a01b03166001600160a01b031663034c44546040518163ffffffff1660e01b8152600401602060405180830381865afa15801561071c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107409190611760565b81101561075f5760405162461bcd60e51b815260040161052990611796565b6040516370a0823160e01b81526001600160a01b0384811660048301528591908416906370a0823190602401602060405180830381865afa1580156107a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107cc9190611760565b10156108385760405162461bcd60e51b815260206004820152603560248201527f5374616b696e67506f6f6c3a2063616e6e6f7420756e7374616b65206d6f7265604482015274207468616e2068617665206f6e206164647265737360581b6064820152608401610529565b604051632770a7eb60e21b81526001600160a01b03848116600483015260248201869052831690639dc29fac90604401600060405180830381600087803b15801561088257600080fd5b505af1158015610896573d6000803e3d6000fd5b505050506108a683868684611278565b5050505050565b60335460408051632b51dafd60e11b815290516000926001600160a01b0316916356a3b5fa9160048083019260209291908290030181865afa158015610421573d6000803e3d6000fd5b6108ff6113f0565b60335460408051632b51dafd60e11b8152905134926001600160a01b0316916356a3b5fa9160048083019260209291908290030181865afa158015610948573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096c9190611760565b81101561098b5760405162461bcd60e51b815260040161052990611796565b60335460408051635eaad42760e11b815290516000926001600160a01b03169163bd55a84e9160048083019260209291908290030181865afa1580156109d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109f99190611779565b604051635361637360e01b8152600481018490529091506000906001600160a01b03831690635361637390602401602060405180830381865afa158015610a44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a689190611760565b6040516340c10f1960e01b8152336004820152602481018290529091506001600160a01b038316906340c10f1990604401600060405180830381600087803b158015610ab357600080fd5b505af1158015610ac7573d6000803e3d6000fd5b505060408051868152602081018590523393507f1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee9092500160405180910390a2505050610b1260018055565b565b603360009054906101000a90046001600160a01b03166001600160a01b0316632ec338ba6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8b9190611779565b6001600160a01b0316336001600160a01b031614610bbb5760405162461bcd60e51b8152600401610529906117e8565b848381148015610bca57508382145b610c205760405162461bcd60e51b815260206004820152602160248201527f5374616b696e67506f6f6c3a206c656e67746820617265206e6f7420657175616044820152601b60fa1b6064820152608401610529565b610c33816801bc16d674ec80000061182b565b471015610c7e5760405162461bcd60e51b81526020600482015260196024820152780e0cadcc8d2dcce40cae8d0cae4e640dcdee840cadcdeeaced603b1b6044820152606401610529565b6033546040805163476c397360e11b815290516000926001600160a01b031691638ed872e69160048083019260209291908290030181865afa158015610cc8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cec9190611779565b905060005b82811015610e2857816001600160a01b0316639b4e46348a8a84818110610d1a57610d1a611731565b9050602002810190610d2c9190611842565b8a8a86818110610d3e57610d3e611731565b9050602002810190610d509190611842565b8a8a88818110610d6257610d62611731565b905060200201356040518663ffffffff1660e01b8152600401610d899594939291906118b2565b600060405180830381600087803b158015610da357600080fd5b505af1158015610db7573d6000803e3d6000fd5b505050507f09bfa94ea4cf30558b9b9914b0029f04369c211b0e1f965f9fa7a29536b300df898983818110610dee57610dee611731565b9050602002810190610e009190611842565b604051610e0e9291906118ec565b60405180910390a180610e2081611747565b915050610cf1565b505050505050505050565b603360009054906101000a90046001600160a01b03166001600160a01b0316632ec338ba6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eaa9190611779565b6001600160a01b0316336001600160a01b031614610eda5760405162461bcd60e51b8152600401610529906117e8565b6801bc16d674ec800000471015610f2f5760405162461bcd60e51b81526020600482015260196024820152780e0cadcc8d2dcce40cae8d0cae4e640dcdee840cadcdeeaced603b1b6044820152606401610529565b603360009054906101000a90046001600160a01b03166001600160a01b0316638ed872e66040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fa69190611779565b6001600160a01b0316639b4e463486868686866040518663ffffffff1660e01b8152600401610fd99594939291906118b2565b600060405180830381600087803b158015610ff357600080fd5b505af1158015611007573d6000803e3d6000fd5b505050507f09bfa94ea4cf30558b9b9914b0029f04369c211b0e1f965f9fa7a29536b300df858560405161103c9291906118ec565b60405180910390a15050505050565b6033546040805163476c397360e11b815290516000926001600160a01b031691638ed872e69160048083019260209291908290030181865afa1580156105c1573d6000803e3d6000fd5b600054610100900460ff16158080156110b55750600054600160ff909116105b806110cf5750303b1580156110cf575060005460ff166001145b6111325760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610529565b6000805460ff191660011790558015611155576000805461ff0019166101001790555b603380546001600160a01b0319166001600160a01b0385169081179091556040805163476c397360e11b81529051638ed872e6916004808201926020929091908290030181865afa1580156111ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111d29190611779565b6001600160a01b03166384d810626040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561120c57600080fd5b505af1158015611220573d6000803e3d6000fd5b5050505061122d82611449565b8015611273576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b801580159061128f57506001600160a01b03831615155b6112ed5760405162461bcd60e51b815260206004820152602960248201527f4c6971756964546f6b656e5374616b696e67506f6f6c3a207a65726f20696e7060448201526875742076616c75657360b81b6064820152608401610529565b80603660008282546112ff9190611908565b90915550506037805460018082019092557f42a7b7dd785cd69714a189dffb3fd7d7174edc9ece837694ce50f7078f7c31ae0180546001600160a01b0319166001600160a01b0386169081179091556039805492830190557fdc16fef70f8d5ddbc01ee3d903d1e69c18a3c7be080eb86a81e0578814ee58d39091018290556000908152603860205260408120805483929061139c908490611908565b909155505060408051828152602081018490526001600160a01b0380861692908716917fe8f73d529f5ced08581a2c18456a6530dbd0dddf94d8c98e0ab8f9883e2f4482910160405180910390a350505050565b6002600154036114425760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610529565b6002600155565b60348190556040805160008152602081018390527f3bb2da990d30b0bc98e39d632b60814d66b3bae55947927dec7a75719de577de910160405180910390a150565b6001600160a01b03811681146114a057600080fd5b50565b6000602082840312156114b557600080fd5b81356114c08161148b565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156114ff578351835292840192918401916001016114e3565b50909695505050505050565b60006020828403121561151d57600080fd5b5035919050565b6000806040838503121561153757600080fd5b82356115428161148b565b946020939093013593505050565b60008083601f84011261156257600080fd5b50813567ffffffffffffffff81111561157a57600080fd5b6020830191508360208260051b850101111561159557600080fd5b9250929050565b600080600080600080606087890312156115b557600080fd5b863567ffffffffffffffff808211156115cd57600080fd5b6115d98a838b01611550565b909850965060208901359150808211156115f257600080fd5b6115fe8a838b01611550565b9096509450604089013591508082111561161757600080fd5b5061162489828a01611550565b979a9699509497509295939492505050565b60008083601f84011261164857600080fd5b50813567ffffffffffffffff81111561166057600080fd5b60208301915083602082850101111561159557600080fd5b60008060008060006060868803121561169057600080fd5b853567ffffffffffffffff808211156116a857600080fd5b6116b489838a01611636565b909750955060208801359150808211156116cd57600080fd5b506116da88828901611636565b96999598509660400135949350505050565b634e487b7160e01b600052601160045260246000fd5b81810381811115611715576117156116ec565b92915050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600060018201611759576117596116ec565b5060010190565b60006020828403121561177257600080fd5b5051919050565b60006020828403121561178b57600080fd5b81516114c08161148b565b60208082526032908201527f5374616b696e67506f6f6c3a2076616c7565206d757374206265206772656174604082015271195c881d1a185b881b5a5b88185b5bdd5b9d60721b606082015260800190565b60208082526023908201527f5374616b696e67506f6f6c3a206f6e6c7920636f6e73656e73757320616c6c6f6040820152621dd95960ea1b606082015260800190565b8082028115828204841417611715576117156116ec565b6000808335601e1984360301811261185957600080fd5b83018035915067ffffffffffffffff82111561187457600080fd5b60200191503681900382131561159557600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6060815260006118c6606083018789611889565b82810360208401526118d9818688611889565b9150508260408301529695505050505050565b602081526000611900602083018486611889565b949350505050565b80820180821115611715576117156116ec56fea2646970667358221220bb6d5f4a5600687da478a6f275314b2f67a0f93e5a54864985870dc7316352de64736f6c63430008130033",
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

	FilterStaked(opts *bind.FilterOpts, staker []common.Address) (*ContractStakedIterator, error)

	WatchStaked(opts *bind.WatchOpts, sink chan<- *ContractStaked, staker []common.Address) (event.Subscription, error)

	ParseStaked(log types.Log) (*ContractStaked, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
