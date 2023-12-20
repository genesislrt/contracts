// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package genesisstakingpool

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CertificateTokenChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinimumStakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinimumUnstakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"PoolOnGoing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getCert\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFreeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinUnstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingConfig\",\"name\":\"earnConfig\",\"type\":\"address\"},{\"internalType\":\"contractIEigenPodManager\",\"name\":\"podManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"deposit_data_root\",\"type\":\"bytes32\"}],\"name\":\"pushToBeacon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"deposit_data_roots\",\"type\":\"bytes32[]\"}],\"name\":\"pushToBeaconMulti\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setCertificateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMinimumStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMinimumUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeCerts\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061152e806100206000396000f3fe6080604052600436106100dd5760003560e01c80638da5cb5b1161007f578063be9631e811610059578063be9631e81461020e578063d05b01ae1461022e578063e93c4f151461024e578063f2fde38b1461026157600080fd5b80638da5cb5b146101c8578063ac76d450146101e6578063afbcc448146101ee57600080fd5b8063233e9903116100bb578063233e99031461015e578063485cc9551461017e57806356a3b5fa1461019e578063715018a6146101b357600080fd5b8063034c4454146100e2578063183cdf241461010a57806321ce57cf1461013c575b600080fd5b3480156100ee57600080fd5b506100f7610281565b6040519081526020015b60405180910390f35b34801561011657600080fd5b506097546001600160a01b03165b6040516001600160a01b039091168152602001610101565b34801561014857600080fd5b5061015c610157366004611018565b6102a4565b005b34801561016a57600080fd5b5061015c610179366004611018565b61045f565b34801561018a57600080fd5b5061015c610199366004611046565b61061a565b3480156101aa57600080fd5b506100f76107ba565b3480156101bf57600080fd5b5061015c6107df565b3480156101d457600080fd5b506033546001600160a01b0316610124565b61015c6107f3565b3480156101fa57600080fd5b5061015c6102093660046110cb565b61080f565b34801561021a57600080fd5b5061015c6102293660046111a7565b610a25565b34801561023a57600080fd5b5061015c61024936600461121b565b610b4a565b34801561025a57600080fd5b50476100f7565b34801561026d57600080fd5b5061015c61027c36600461121b565b610c4b565b609a5460009061029f9067ffffffffffffffff16633b9aca00611255565b905090565b609860009054906101000a90046001600160a01b03166001600160a01b031663732524946040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031b9190611274565b6001600160a01b0316336001600160a01b0316146103545760405162461bcd60e51b815260040161034b90611291565b60405180910390fd5b610362633b9aca00826112f2565b1561037f5760405162461bcd60e51b815260040161034b90611306565b6000610389610281565b9050610399633b9aca008361135c565b609a805467ffffffffffffffff191667ffffffffffffffff92909216918217905582906103ca90633b9aca00611370565b67ffffffffffffffff16146104215760405162461bcd60e51b815260206004820152601c60248201527f47656e657369735374616b696e67506f6f6c3a206f766572666c6f7700000000604482015260640161034b565b60408051828152602081018490527f976e4c5a2502181884199a44267e9ba339f2d071066e4f0ae97600ce4975935c91015b60405180910390a15050565b609860009054906101000a90046001600160a01b03166001600160a01b031663732524946040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d69190611274565b6001600160a01b0316336001600160a01b0316146105065760405162461bcd60e51b815260040161034b90611291565b610514633b9aca00826112f2565b156105315760405162461bcd60e51b815260040161034b90611306565b600061053b6107ba565b905061054b633b9aca008361135c565b6099805467ffffffffffffffff60a01b1916600160a01b67ffffffffffffffff93841681029190911791829055849261058d929190910416633b9aca00611370565b67ffffffffffffffff16146105e45760405162461bcd60e51b815260206004820152601c60248201527f47656e657369735374616b696e67506f6f6c3a206f766572666c6f7700000000604482015260640161034b565b60408051828152602081018490527fdc4a0b2dc1fa27da98de2ac6f8fa373b4be405e1bf69fc3976597b6d56b79abc9101610453565b600054610100900460ff161580801561063a5750600054600160ff909116105b806106545750303b158015610654575060005460ff166001145b6106b75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161034b565b6000805460ff1916600117905580156106da576000805461ff0019166101001790555b82609860006101000a8154816001600160a01b0302191690836001600160a01b03160217905550816001600160a01b03166384d810626040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561073c57600080fd5b505af1158015610750573d6000803e3d6000fd5b5050609980546001600160a01b0319166001600160a01b038616179055505080156107b5576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60995460009061029f90600160a01b900467ffffffffffffffff16633b9aca00611255565b6107e7610cc4565b6107f16000610d1e565b565b6107fb610d70565b6108053334610dc9565b6107f16001606555565b6033546001600160a01b03163314806108af5750609860009054906101000a90046001600160a01b03166001600160a01b031663dddce8c16040518163ffffffff1660e01b8152600401602060405180830381865afa158015610876573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089a9190611274565b6001600160a01b0316336001600160a01b0316145b6108cb5760405162461bcd60e51b815260040161034b906113a0565b8483811480156108da57508382145b6109375760405162461bcd60e51b815260206004820152602860248201527f47656e657369735374616b696e67506f6f6c3a206c656e67746820617265206e6044820152671bdd08195c5d585b60c21b606482015260840161034b565b61094a816801bc16d674ec800000611255565b4710156109955760405162461bcd60e51b81526020600482015260196024820152780e0cadcc8d2dcce40cae8d0cae4e640dcdee840cadcdeeaced603b1b604482015260640161034b565b60005b81811015610a1b57610a098888838181106109b5576109b56113ea565b90506020028101906109c79190611400565b8888858181106109d9576109d96113ea565b90506020028101906109eb9190611400565b8888878181106109fd576109fd6113ea565b90506020020135610e44565b80610a1381611447565b915050610998565b5050505050505050565b6033546001600160a01b0316331480610ac55750609860009054906101000a90046001600160a01b03166001600160a01b031663dddce8c16040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab09190611274565b6001600160a01b0316336001600160a01b0316145b610ae15760405162461bcd60e51b815260040161034b906113a0565b6801bc16d674ec800000471015610b365760405162461bcd60e51b81526020600482015260196024820152780e0cadcc8d2dcce40cae8d0cae4e640dcdee840cadcdeeaced603b1b604482015260640161034b565b610b438585858585610e44565b5050505050565b609860009054906101000a90046001600160a01b03166001600160a01b031663732524946040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc19190611274565b6001600160a01b0316336001600160a01b031614610bf15760405162461bcd60e51b815260040161034b90611291565b609780546001600160a01b038381166001600160a01b031983168117909355604080519190921680825260208201939093527f63d8371ad7e38cfb2f86b20688da999f9750c68f15d078ac59c9a64c410b44329101610453565b610c53610cc4565b6001600160a01b038116610cb85760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161034b565b610cc181610d1e565b50565b6033546001600160a01b031633146107f15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161034b565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600260655403610dc25760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015260640161034b565b6002606555565b609754604051635361637360e01b8152600481018390526000916001600160a01b031690635361637390602401602060405180830381865afa158015610e13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e379190611460565b90506107b5838383610eee565b6099546040516326d3918d60e21b81526001600160a01b0390911690639b4e463490610e7c90889088908890889088906004016114a2565b600060405180830381600087803b158015610e9657600080fd5b505af1158015610eaa573d6000803e3d6000fd5b505050507f09bfa94ea4cf30558b9b9914b0029f04369c211b0e1f965f9fa7a29536b300df8585604051610edf9291906114dc565b60405180910390a15050505050565b610ef66107ba565b821015610f6b5760405162461bcd60e51b815260206004820152603960248201527f47656e657369735374616b696e67506f6f6c3a2076616c7565206d757374206260448201527f652067726561746572207468616e206d696e20616d6f756e7400000000000000606482015260840161034b565b6097546040516340c10f1960e01b8152336004820152602481018390526001600160a01b03909116906340c10f1990604401600060405180830381600087803b158015610fb757600080fd5b505af1158015610fcb573d6000803e3d6000fd5b505060408051858152602081018590526001600160a01b03871693507f1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee9092500160405180910390a2505050565b60006020828403121561102a57600080fd5b5035919050565b6001600160a01b0381168114610cc157600080fd5b6000806040838503121561105957600080fd5b823561106481611031565b9150602083013561107481611031565b809150509250929050565b60008083601f84011261109157600080fd5b50813567ffffffffffffffff8111156110a957600080fd5b6020830191508360208260051b85010111156110c457600080fd5b9250929050565b600080600080600080606087890312156110e457600080fd5b863567ffffffffffffffff808211156110fc57600080fd5b6111088a838b0161107f565b9098509650602089013591508082111561112157600080fd5b61112d8a838b0161107f565b9096509450604089013591508082111561114657600080fd5b5061115389828a0161107f565b979a9699509497509295939492505050565b60008083601f84011261117757600080fd5b50813567ffffffffffffffff81111561118f57600080fd5b6020830191508360208285010111156110c457600080fd5b6000806000806000606086880312156111bf57600080fd5b853567ffffffffffffffff808211156111d757600080fd5b6111e389838a01611165565b909750955060208801359150808211156111fc57600080fd5b5061120988828901611165565b96999598509660400135949350505050565b60006020828403121561122d57600080fd5b813561123881611031565b9392505050565b634e487b7160e01b600052601160045260246000fd5b600081600019048311821515161561126f5761126f61123f565b500290565b60006020828403121561128657600080fd5b815161123881611031565b6020808252602b908201527f47656e657369735374616b696e67506f6f6c3a206f6e6c7920676f7665726e6160408201526a1b98d948185b1b1bddd95960aa1b606082015260800190565b634e487b7160e01b600052601260045260246000fd5b600082611301576113016112dc565b500690565b60208082526036908201527f47656e657369735374616b696e67506f6f6c3a2076616c75652073686f756c64604082015275206265206d756c7469706c696564206f66206777656960501b606082015260800190565b60008261136b5761136b6112dc565b500490565b600067ffffffffffffffff808316818516818304811182151516156113975761139761123f565b02949350505050565b6020808252602a908201527f47656e657369735374616b696e67506f6f6c3a206f6e6c7920636f6e73656e736040820152691d5cc8185b1b1bddd95960b21b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000808335601e1984360301811261141757600080fd5b83018035915067ffffffffffffffff82111561143257600080fd5b6020019150368190038213156110c457600080fd5b6000600182016114595761145961123f565b5060010190565b60006020828403121561147257600080fd5b5051919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6060815260006114b6606083018789611479565b82810360208401526114c9818688611479565b9150508260408301529695505050505050565b6020815260006114f0602083018486611479565b94935050505056fea2646970667358221220442115317f7f84e71d8f733d04f7665fea9e46fd89d78c1c50c43082dd65864d64736f6c634300080f0033",
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

func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, earnConfig common.Address, podManager common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", earnConfig, podManager)
}

func (_Contract *ContractSession) Initialize(earnConfig common.Address, podManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, earnConfig, podManager)
}

func (_Contract *ContractTransactorSession) Initialize(earnConfig common.Address, podManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, earnConfig, podManager)
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

func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactor) SetCertificateToken(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setCertificateToken", newValue)
}

func (_Contract *ContractSession) SetCertificateToken(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetCertificateToken(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetCertificateToken(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetCertificateToken(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) SetMinimumStake(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMinimumStake", newValue)
}

func (_Contract *ContractSession) SetMinimumStake(newValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinimumStake(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetMinimumStake(newValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinimumStake(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) SetMinimumUnstake(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMinimumUnstake", newValue)
}

func (_Contract *ContractSession) SetMinimumUnstake(newValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinimumUnstake(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetMinimumUnstake(newValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinimumUnstake(&_Contract.TransactOpts, newValue)
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

func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

type ContractCertificateTokenChangedIterator struct {
	Event *ContractCertificateTokenChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractCertificateTokenChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCertificateTokenChanged)
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
		it.Event = new(ContractCertificateTokenChanged)
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

func (it *ContractCertificateTokenChangedIterator) Error() error {
	return it.fail
}

func (it *ContractCertificateTokenChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractCertificateTokenChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterCertificateTokenChanged(opts *bind.FilterOpts) (*ContractCertificateTokenChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CertificateTokenChanged")
	if err != nil {
		return nil, err
	}
	return &ContractCertificateTokenChangedIterator{contract: _Contract.contract, event: "CertificateTokenChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchCertificateTokenChanged(opts *bind.WatchOpts, sink chan<- *ContractCertificateTokenChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CertificateTokenChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractCertificateTokenChanged)
				if err := _Contract.contract.UnpackLog(event, "CertificateTokenChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseCertificateTokenChanged(log types.Log) (*ContractCertificateTokenChanged, error) {
	event := new(ContractCertificateTokenChanged)
	if err := _Contract.contract.UnpackLog(event, "CertificateTokenChanged", log); err != nil {
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

type ContractMinimumStakeChangedIterator struct {
	Event *ContractMinimumStakeChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractMinimumStakeChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMinimumStakeChanged)
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
		it.Event = new(ContractMinimumStakeChanged)
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

func (it *ContractMinimumStakeChangedIterator) Error() error {
	return it.fail
}

func (it *ContractMinimumStakeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractMinimumStakeChanged struct {
	PrevValue *big.Int
	NewValue  *big.Int
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterMinimumStakeChanged(opts *bind.FilterOpts) (*ContractMinimumStakeChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MinimumStakeChanged")
	if err != nil {
		return nil, err
	}
	return &ContractMinimumStakeChangedIterator{contract: _Contract.contract, event: "MinimumStakeChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchMinimumStakeChanged(opts *bind.WatchOpts, sink chan<- *ContractMinimumStakeChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MinimumStakeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractMinimumStakeChanged)
				if err := _Contract.contract.UnpackLog(event, "MinimumStakeChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseMinimumStakeChanged(log types.Log) (*ContractMinimumStakeChanged, error) {
	event := new(ContractMinimumStakeChanged)
	if err := _Contract.contract.UnpackLog(event, "MinimumStakeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractMinimumUnstakeChangedIterator struct {
	Event *ContractMinimumUnstakeChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractMinimumUnstakeChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMinimumUnstakeChanged)
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
		it.Event = new(ContractMinimumUnstakeChanged)
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

func (it *ContractMinimumUnstakeChangedIterator) Error() error {
	return it.fail
}

func (it *ContractMinimumUnstakeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractMinimumUnstakeChanged struct {
	PrevValue *big.Int
	NewValue  *big.Int
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterMinimumUnstakeChanged(opts *bind.FilterOpts) (*ContractMinimumUnstakeChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MinimumUnstakeChanged")
	if err != nil {
		return nil, err
	}
	return &ContractMinimumUnstakeChangedIterator{contract: _Contract.contract, event: "MinimumUnstakeChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchMinimumUnstakeChanged(opts *bind.WatchOpts, sink chan<- *ContractMinimumUnstakeChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MinimumUnstakeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractMinimumUnstakeChanged)
				if err := _Contract.contract.UnpackLog(event, "MinimumUnstakeChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseMinimumUnstakeChanged(log types.Log) (*ContractMinimumUnstakeChanged, error) {
	event := new(ContractMinimumUnstakeChanged)
	if err := _Contract.contract.UnpackLog(event, "MinimumUnstakeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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

func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log
}

func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

type ContractUnstakedIterator struct {
	Event *ContractUnstaked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractUnstakedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUnstaked)
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
		it.Event = new(ContractUnstaked)
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

func (it *ContractUnstakedIterator) Error() error {
	return it.fail
}

func (it *ContractUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractUnstaked struct {
	OwnerAddress    common.Address
	ReceiverAddress common.Address
	Amount          *big.Int
	Shares          *big.Int
	Raw             types.Log
}

func (_Contract *ContractFilterer) FilterUnstaked(opts *bind.FilterOpts, ownerAddress []common.Address, receiverAddress []common.Address) (*ContractUnstakedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}
	var receiverAddressRule []interface{}
	for _, receiverAddressItem := range receiverAddress {
		receiverAddressRule = append(receiverAddressRule, receiverAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Unstaked", ownerAddressRule, receiverAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractUnstakedIterator{contract: _Contract.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *ContractUnstaked, ownerAddress []common.Address, receiverAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}
	var receiverAddressRule []interface{}
	for _, receiverAddressItem := range receiverAddress {
		receiverAddressRule = append(receiverAddressRule, receiverAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Unstaked", ownerAddressRule, receiverAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractUnstaked)
				if err := _Contract.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

func (_Contract *ContractFilterer) ParseUnstaked(log types.Log) (*ContractUnstaked, error) {
	event := new(ContractUnstaked)
	if err := _Contract.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Contract *Contract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Contract.abi.Events["CertificateTokenChanged"].ID:
		return _Contract.ParseCertificateTokenChanged(log)
	case _Contract.abi.Events["Initialized"].ID:
		return _Contract.ParseInitialized(log)
	case _Contract.abi.Events["MinimumStakeChanged"].ID:
		return _Contract.ParseMinimumStakeChanged(log)
	case _Contract.abi.Events["MinimumUnstakeChanged"].ID:
		return _Contract.ParseMinimumUnstakeChanged(log)
	case _Contract.abi.Events["OwnershipTransferred"].ID:
		return _Contract.ParseOwnershipTransferred(log)
	case _Contract.abi.Events["PoolOnGoing"].ID:
		return _Contract.ParsePoolOnGoing(log)
	case _Contract.abi.Events["Received"].ID:
		return _Contract.ParseReceived(log)
	case _Contract.abi.Events["Staked"].ID:
		return _Contract.ParseStaked(log)
	case _Contract.abi.Events["Unstaked"].ID:
		return _Contract.ParseUnstaked(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ContractCertificateTokenChanged) Topic() common.Hash {
	return common.HexToHash("0x63d8371ad7e38cfb2f86b20688da999f9750c68f15d078ac59c9a64c410b4432")
}

func (ContractInitialized) Topic() common.Hash {
	return common.HexToHash("0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498")
}

func (ContractMinimumStakeChanged) Topic() common.Hash {
	return common.HexToHash("0xdc4a0b2dc1fa27da98de2ac6f8fa373b4be405e1bf69fc3976597b6d56b79abc")
}

func (ContractMinimumUnstakeChanged) Topic() common.Hash {
	return common.HexToHash("0x976e4c5a2502181884199a44267e9ba339f2d071066e4f0ae97600ce4975935c")
}

func (ContractOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
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

func (ContractUnstaked) Topic() common.Hash {
	return common.HexToHash("0x06cc7e90b4f2b554a9614b0caa84f909f3498c820ae47c731f490c28c07f7d3b")
}

func (_Contract *Contract) Address() common.Address {
	return _Contract.address
}

type ContractInterface interface {
	GetCert(opts *bind.CallOpts) (common.Address, error)

	GetFreeBalance(opts *bind.CallOpts) (*big.Int, error)

	GetMinStake(opts *bind.CallOpts) (*big.Int, error)

	GetMinUnstake(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Initialize(opts *bind.TransactOpts, earnConfig common.Address, podManager common.Address) (*types.Transaction, error)

	PushToBeacon(opts *bind.TransactOpts, pubkey []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error)

	PushToBeaconMulti(opts *bind.TransactOpts, pubkeys [][]byte, signatures [][]byte, deposit_data_roots [][32]byte) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetCertificateToken(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	SetMinimumStake(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error)

	SetMinimumUnstake(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error)

	StakeCerts(opts *bind.TransactOpts) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	FilterCertificateTokenChanged(opts *bind.FilterOpts) (*ContractCertificateTokenChangedIterator, error)

	WatchCertificateTokenChanged(opts *bind.WatchOpts, sink chan<- *ContractCertificateTokenChanged) (event.Subscription, error)

	ParseCertificateTokenChanged(log types.Log) (*ContractCertificateTokenChanged, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*ContractInitialized, error)

	FilterMinimumStakeChanged(opts *bind.FilterOpts) (*ContractMinimumStakeChangedIterator, error)

	WatchMinimumStakeChanged(opts *bind.WatchOpts, sink chan<- *ContractMinimumStakeChanged) (event.Subscription, error)

	ParseMinimumStakeChanged(log types.Log) (*ContractMinimumStakeChanged, error)

	FilterMinimumUnstakeChanged(opts *bind.FilterOpts) (*ContractMinimumUnstakeChangedIterator, error)

	WatchMinimumUnstakeChanged(opts *bind.WatchOpts, sink chan<- *ContractMinimumUnstakeChanged) (event.Subscription, error)

	ParseMinimumUnstakeChanged(log types.Log) (*ContractMinimumUnstakeChanged, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error)

	FilterPoolOnGoing(opts *bind.FilterOpts) (*ContractPoolOnGoingIterator, error)

	WatchPoolOnGoing(opts *bind.WatchOpts, sink chan<- *ContractPoolOnGoing) (event.Subscription, error)

	ParsePoolOnGoing(log types.Log) (*ContractPoolOnGoing, error)

	FilterReceived(opts *bind.FilterOpts, from []common.Address) (*ContractReceivedIterator, error)

	WatchReceived(opts *bind.WatchOpts, sink chan<- *ContractReceived, from []common.Address) (event.Subscription, error)

	ParseReceived(log types.Log) (*ContractReceived, error)

	FilterStaked(opts *bind.FilterOpts, staker []common.Address) (*ContractStakedIterator, error)

	WatchStaked(opts *bind.WatchOpts, sink chan<- *ContractStaked, staker []common.Address) (event.Subscription, error)

	ParseStaked(log types.Log) (*ContractStaked, error)

	FilterUnstaked(opts *bind.FilterOpts, ownerAddress []common.Address, receiverAddress []common.Address) (*ContractUnstakedIterator, error)

	WatchUnstaked(opts *bind.WatchOpts, sink chan<- *ContractUnstaked, ownerAddress []common.Address, receiverAddress []common.Address) (event.Subscription, error)

	ParseUnstaked(log types.Log) (*ContractUnstaked, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
