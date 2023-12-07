// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package internetbondratiofeed

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"failedRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"RatioNotUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RatioThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRatio\",\"type\":\"uint256\"}],\"name\":\"RatioUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"averagePercentageRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getRatioFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRatioThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"historicalRatios\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"lastUpdate\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"repairRatioFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setRatioThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ratios\",\"type\":\"uint256[]\"}],\"name\":\"updateRatioBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506111b0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063a1f1d48d1161008c578063ba34fa0511610066578063ba34fa05146101e2578063c4d66de8146101ea578063ec653c4b146101fd578063f2fde38b1461023d57600080fd5b8063a1f1d48d14610193578063ac8a584a146101bc578063b037d56a146101cf57600080fd5b80632ef86a1f116100c85780632ef86a1f1461014a578063715018a61461015d5780638da5cb5b146101655780639870d7fe1461018057600080fd5b806308af5431146100ef5780632364753a146101145780632acaaff414610129575b600080fd5b6100fa6305f5e10081565b60405163ffffffff90911681526020015b60405180910390f35b610127610122366004610e04565b610250565b005b61013c610137366004610e39565b610302565b60405190815260200161010b565b610127610158366004610e39565b610498565b61012761053f565b6033546040516001600160a01b03909116815260200161010b565b61012761018e366004610e63565b610553565b61013c6101a1366004610e63565b6001600160a01b031660009081526066602052604090205490565b6101276101ca366004610e63565b610668565b6101276101dd366004610ed0565b61071b565b60685461013c565b6101276101f8366004610e63565b610a79565b61022761020b366004610e63565b60676020526000908152604090206003015464ffffffffff1681565b60405164ffffffffff909116815260200161010b565b61012761024b366004610e63565b610baa565b610258610c23565b6305f5e1008110801561026b5750600081115b6102bc5760405162461bcd60e51b815260206004820152601f60248201527f77726f6e672076616c756520666f7220726174696f207468726573686f6c640060448201526064015b60405180910390fd5b606880549082905560408051828152602081018490527f661e4cadf2d36ec16a59d60dcfeebe23f9be2aec99852725798a4be99790840e91015b60405180910390a15050565b600080821180156103135750600882105b61035f5760405162461bcd60e51b815260206004820152601960248201527f6461792073686f756c642062652066726f6d203120746f20370000000000000060448201526064016102b3565b6001600160a01b0383166000908152606760205260408120805490916001600160401b03909116908260086103948785610f67565b61039e9190610f94565b6103a9906001610fa8565b600981106103b9576103b9610f3b565b60048104909101546001600160401b036008600390931683026101000a90910416915060009084906103eb9085610fc0565b6103f6906001610fe6565b6001600160401b03166009811061040f5761040f610f3b565b600491828204019190066008029054906101000a90046001600160401b03166001600160401b031690508082101561044e576000945050505050610492565b6104588683611011565b6104628284610f67565b6104759068056bc75e2d63100000611011565b6104819061016d611011565b61048b9190611030565b9450505050505b92915050565b6104a0610c23565b806000036104e05760405162461bcd60e51b815260206004820152600d60248201526c726174696f206973207a65726f60981b60448201526064016102b3565b6001600160a01b038216600081815260666020908152604091829020805490859055825181815291820185905292917f4c5c23b4efbfea6d16c8453f565e165a02a22cda9a8dc7aac0a66f91d2304da6910160405180910390a2505050565b610547610c23565b6105516000610c7d565b565b61055b610c23565b6001600160a01b0381166105b15760405162461bcd60e51b815260206004820152601960248201527f6f70657261746f72206d757374206265206e6f6e2d7a65726f0000000000000060448201526064016102b3565b6001600160a01b03811660009081526065602052604090205460ff161561060d5760405162461bcd60e51b815260206004820152601060248201526f30b63932b0b23c9037b832b930ba37b960811b60448201526064016102b3565b6001600160a01b038116600081815260656020908152604091829020805460ff1916600117905590519182527fac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d91015b60405180910390a150565b610670610c23565b6001600160a01b03811660009081526065602052604090205460ff166106ca5760405162461bcd60e51b815260206004820152600f60248201526e3737ba1030b71037b832b930ba37b960891b60448201526064016102b3565b6001600160a01b038116600081815260656020908152604091829020805460ff1916905590519182527f80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d910161065d565b6033546001600160a01b031633148061074357503360009081526065602052604090205460ff165b6107875760405162461bcd60e51b815260206004820152601560248201527413dc195c985d1bdc8e881b9bdd08185b1b1bddd959605a1b60448201526064016102b3565b8281146107cd5760405162461bcd60e51b8152602060048201526014602482015273636f7272757074656420726174696f206461746160601b60448201526064016102b3565b60006068541161081f5760405162461bcd60e51b815260206004820152601a60248201527f726174696f207468726573686f6c64206973206e6f742073657400000000000060448201526064016102b3565b60005b83811015610a7257600085858381811061083e5761083e610f3b565b90506020020160208101906108539190610e63565b6001600160a01b0381166000908152606960209081526040808320546066909252822054929350919086868681811061088e5761088e610f3b565b9050602002013590506000806108a5858486610ccf565b91509150816108fc57856001600160a01b03167f2471a7627ad27128888e46dfc72f5d674c7156d6e99c969a675492a558a0b0e084836040516108e9929190611044565b60405180910390a2505050505050610a60565b6001600160a01b03861660008181526066602090815260409182902086905581518781529081018690527f4c5c23b4efbfea6d16c8453f565e165a02a22cda9a8dc7aac0a66f91d2304da6910160405180910390a26001600160a01b038616600090815260696020908152604080832064ffffffffff4281811690925560679093529220600381015490926201514492610997921690610f67565b1115610a585780546001600160401b0316848260086109b7846001610fe6565b6109c19190610fc0565b6109cc906001610fe6565b6001600160401b0316600981106109e5576109e5610f3b565b600491828204019190066008026101000a8154816001600160401b0302191690836001600160401b03160217905550806001610a219190610fe6565b825467ffffffffffffffff19166001600160401b03919091161782555060038101805464ffffffffff19164264ffffffffff161790555b505050505050505b80610a6a816110a1565b915050610822565b5050505050565b600054610100900460ff1615808015610a995750600054600160ff909116105b80610ab35750303b158015610ab3575060005460ff166001145b610b165760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102b3565b6000805460ff191660011790558015610b39576000805461ff0019166101001790555b610b41610daa565b6001600160a01b0382166000908152606560205260409020805460ff191660011790558015610ba6576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020016102f6565b5050565b610bb2610c23565b6001600160a01b038116610c175760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102b3565b610c2081610c7d565b50565b6033546001600160a01b031633146105515760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b3565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000606082600003610ce45760019150610da2565b61a8c0610cf18642610f67565b1015610d1a57816040518060600160405280602881526020016111066028913991509150610da2565b82841115610d45578160405180606001604052806024815260200161112e6024913991509150610da2565b6068546000906305f5e10090610d5b9086611011565b610d659190611030565b9050610d718185610f67565b851015610d9c5782604051806060016040528060298152602001611152602991399250925050610da2565b60019250505b935093915050565b600054610100900460ff16610dd15760405162461bcd60e51b81526004016102b3906110ba565b610551600054610100900460ff16610dfb5760405162461bcd60e51b81526004016102b3906110ba565b61055133610c7d565b600060208284031215610e1657600080fd5b5035919050565b80356001600160a01b0381168114610e3457600080fd5b919050565b60008060408385031215610e4c57600080fd5b610e5583610e1d565b946020939093013593505050565b600060208284031215610e7557600080fd5b610e7e82610e1d565b9392505050565b60008083601f840112610e9757600080fd5b5081356001600160401b03811115610eae57600080fd5b6020830191508360208260051b8501011115610ec957600080fd5b9250929050565b60008060008060408587031215610ee657600080fd5b84356001600160401b0380821115610efd57600080fd5b610f0988838901610e85565b90965094506020870135915080821115610f2257600080fd5b50610f2f87828801610e85565b95989497509550505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082821015610f7957610f79610f51565b500390565b634e487b7160e01b600052601260045260246000fd5b600082610fa357610fa3610f7e565b500690565b60008219821115610fbb57610fbb610f51565b500190565b60006001600160401b0380841680610fda57610fda610f7e565b92169190910692915050565b60006001600160401b0380831681851680830382111561100857611008610f51565b01949350505050565b600081600019048311821515161561102b5761102b610f51565b500290565b60008261103f5761103f610f7e565b500490565b82815260006020604081840152835180604085015260005b818110156110785785810183015185820160600152820161105c565b8181111561108a576000606083870101525b50601f01601f191692909201606001949350505050565b6000600182016110b3576110b3610f51565b5060010190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b60608201526080019056fe726174696f207761732075706461746564206c657373207468616e20313220686f7572732061676f6e657720726174696f2063616e6e6f742062652067726561746572207468616e206f6c646e657720726174696f20746f6f206c6f772c206e6f7420696e207468726573686f6c642072616e6765a26469706673582212207bfe76aeaa4e7ad470d3b2c8dcdc899581c4ba242fab74239973593455368ddf64736f6c634300080f0033",
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

func (_Contract *ContractCaller) MAXTHRESHOLD(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "MAX_THRESHOLD")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

func (_Contract *ContractSession) MAXTHRESHOLD() (uint32, error) {
	return _Contract.Contract.MAXTHRESHOLD(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) MAXTHRESHOLD() (uint32, error) {
	return _Contract.Contract.MAXTHRESHOLD(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) AveragePercentageRate(opts *bind.CallOpts, addr common.Address, day *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "averagePercentageRate", addr, day)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) AveragePercentageRate(addr common.Address, day *big.Int) (*big.Int, error) {
	return _Contract.Contract.AveragePercentageRate(&_Contract.CallOpts, addr, day)
}

func (_Contract *ContractCallerSession) AveragePercentageRate(addr common.Address, day *big.Int) (*big.Int, error) {
	return _Contract.Contract.AveragePercentageRate(&_Contract.CallOpts, addr, day)
}

func (_Contract *ContractCaller) GetRatioFor(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRatioFor", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) GetRatioFor(token common.Address) (*big.Int, error) {
	return _Contract.Contract.GetRatioFor(&_Contract.CallOpts, token)
}

func (_Contract *ContractCallerSession) GetRatioFor(token common.Address) (*big.Int, error) {
	return _Contract.Contract.GetRatioFor(&_Contract.CallOpts, token)
}

func (_Contract *ContractCaller) GetRatioThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRatioThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) GetRatioThreshold() (*big.Int, error) {
	return _Contract.Contract.GetRatioThreshold(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetRatioThreshold() (*big.Int, error) {
	return _Contract.Contract.GetRatioThreshold(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) HistoricalRatios(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "historicalRatios", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) HistoricalRatios(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.HistoricalRatios(&_Contract.CallOpts, arg0)
}

func (_Contract *ContractCallerSession) HistoricalRatios(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.HistoricalRatios(&_Contract.CallOpts, arg0)
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

func (_Contract *ContractTransactor) AddOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addOperator", operator)
}

func (_Contract *ContractSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddOperator(&_Contract.TransactOpts, operator)
}

func (_Contract *ContractTransactorSession) AddOperator(operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddOperator(&_Contract.TransactOpts, operator)
}

func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", operator)
}

func (_Contract *ContractSession) Initialize(operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, operator)
}

func (_Contract *ContractTransactorSession) Initialize(operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, operator)
}

func (_Contract *ContractTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeOperator", operator)
}

func (_Contract *ContractSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveOperator(&_Contract.TransactOpts, operator)
}

func (_Contract *ContractTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveOperator(&_Contract.TransactOpts, operator)
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

func (_Contract *ContractTransactor) RepairRatioFor(opts *bind.TransactOpts, token common.Address, ratio *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "repairRatioFor", token, ratio)
}

func (_Contract *ContractSession) RepairRatioFor(token common.Address, ratio *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RepairRatioFor(&_Contract.TransactOpts, token, ratio)
}

func (_Contract *ContractTransactorSession) RepairRatioFor(token common.Address, ratio *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RepairRatioFor(&_Contract.TransactOpts, token, ratio)
}

func (_Contract *ContractTransactor) SetRatioThreshold(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setRatioThreshold", newValue)
}

func (_Contract *ContractSession) SetRatioThreshold(newValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetRatioThreshold(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetRatioThreshold(newValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetRatioThreshold(&_Contract.TransactOpts, newValue)
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

func (_Contract *ContractTransactor) UpdateRatioBatch(opts *bind.TransactOpts, addresses []common.Address, ratios []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateRatioBatch", addresses, ratios)
}

func (_Contract *ContractSession) UpdateRatioBatch(addresses []common.Address, ratios []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRatioBatch(&_Contract.TransactOpts, addresses, ratios)
}

func (_Contract *ContractTransactorSession) UpdateRatioBatch(addresses []common.Address, ratios []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRatioBatch(&_Contract.TransactOpts, addresses, ratios)
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

type ContractOperatorAddedIterator struct {
	Event *ContractOperatorAdded

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractOperatorAddedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOperatorAdded)
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
		it.Event = new(ContractOperatorAdded)
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

func (it *ContractOperatorAddedIterator) Error() error {
	return it.fail
}

func (it *ContractOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractOperatorAdded struct {
	Operator common.Address
	Raw      types.Log
}

func (_Contract *ContractFilterer) FilterOperatorAdded(opts *bind.FilterOpts) (*ContractOperatorAddedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return &ContractOperatorAddedIterator{contract: _Contract.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *ContractOperatorAdded) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OperatorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractOperatorAdded)
				if err := _Contract.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

func (_Contract *ContractFilterer) ParseOperatorAdded(log types.Log) (*ContractOperatorAdded, error) {
	event := new(ContractOperatorAdded)
	if err := _Contract.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractOperatorRemovedIterator struct {
	Event *ContractOperatorRemoved

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractOperatorRemovedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOperatorRemoved)
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
		it.Event = new(ContractOperatorRemoved)
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

func (it *ContractOperatorRemovedIterator) Error() error {
	return it.fail
}

func (it *ContractOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractOperatorRemoved struct {
	Operator common.Address
	Raw      types.Log
}

func (_Contract *ContractFilterer) FilterOperatorRemoved(opts *bind.FilterOpts) (*ContractOperatorRemovedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OperatorRemoved")
	if err != nil {
		return nil, err
	}
	return &ContractOperatorRemovedIterator{contract: _Contract.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *ContractOperatorRemoved) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OperatorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractOperatorRemoved)
				if err := _Contract.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

func (_Contract *ContractFilterer) ParseOperatorRemoved(log types.Log) (*ContractOperatorRemoved, error) {
	event := new(ContractOperatorRemoved)
	if err := _Contract.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

type ContractRatioNotUpdatedIterator struct {
	Event *ContractRatioNotUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractRatioNotUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRatioNotUpdated)
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
		it.Event = new(ContractRatioNotUpdated)
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

func (it *ContractRatioNotUpdatedIterator) Error() error {
	return it.fail
}

func (it *ContractRatioNotUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractRatioNotUpdated struct {
	TokenAddress common.Address
	FailedRatio  *big.Int
	Reason       string
	Raw          types.Log
}

func (_Contract *ContractFilterer) FilterRatioNotUpdated(opts *bind.FilterOpts, tokenAddress []common.Address) (*ContractRatioNotUpdatedIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RatioNotUpdated", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractRatioNotUpdatedIterator{contract: _Contract.contract, event: "RatioNotUpdated", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchRatioNotUpdated(opts *bind.WatchOpts, sink chan<- *ContractRatioNotUpdated, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RatioNotUpdated", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractRatioNotUpdated)
				if err := _Contract.contract.UnpackLog(event, "RatioNotUpdated", log); err != nil {
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

func (_Contract *ContractFilterer) ParseRatioNotUpdated(log types.Log) (*ContractRatioNotUpdated, error) {
	event := new(ContractRatioNotUpdated)
	if err := _Contract.contract.UnpackLog(event, "RatioNotUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractRatioThresholdChangedIterator struct {
	Event *ContractRatioThresholdChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractRatioThresholdChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRatioThresholdChanged)
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
		it.Event = new(ContractRatioThresholdChanged)
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

func (it *ContractRatioThresholdChangedIterator) Error() error {
	return it.fail
}

func (it *ContractRatioThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractRatioThresholdChanged struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log
}

func (_Contract *ContractFilterer) FilterRatioThresholdChanged(opts *bind.FilterOpts) (*ContractRatioThresholdChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RatioThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &ContractRatioThresholdChangedIterator{contract: _Contract.contract, event: "RatioThresholdChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchRatioThresholdChanged(opts *bind.WatchOpts, sink chan<- *ContractRatioThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RatioThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractRatioThresholdChanged)
				if err := _Contract.contract.UnpackLog(event, "RatioThresholdChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseRatioThresholdChanged(log types.Log) (*ContractRatioThresholdChanged, error) {
	event := new(ContractRatioThresholdChanged)
	if err := _Contract.contract.UnpackLog(event, "RatioThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractRatioUpdatedIterator struct {
	Event *ContractRatioUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractRatioUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRatioUpdated)
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
		it.Event = new(ContractRatioUpdated)
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

func (it *ContractRatioUpdatedIterator) Error() error {
	return it.fail
}

func (it *ContractRatioUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractRatioUpdated struct {
	TokenAddress common.Address
	OldRatio     *big.Int
	NewRatio     *big.Int
	Raw          types.Log
}

func (_Contract *ContractFilterer) FilterRatioUpdated(opts *bind.FilterOpts, tokenAddress []common.Address) (*ContractRatioUpdatedIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RatioUpdated", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractRatioUpdatedIterator{contract: _Contract.contract, event: "RatioUpdated", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchRatioUpdated(opts *bind.WatchOpts, sink chan<- *ContractRatioUpdated, tokenAddress []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RatioUpdated", tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractRatioUpdated)
				if err := _Contract.contract.UnpackLog(event, "RatioUpdated", log); err != nil {
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

func (_Contract *ContractFilterer) ParseRatioUpdated(log types.Log) (*ContractRatioUpdated, error) {
	event := new(ContractRatioUpdated)
	if err := _Contract.contract.UnpackLog(event, "RatioUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Contract *Contract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Contract.abi.Events["Initialized"].ID:
		return _Contract.ParseInitialized(log)
	case _Contract.abi.Events["OperatorAdded"].ID:
		return _Contract.ParseOperatorAdded(log)
	case _Contract.abi.Events["OperatorRemoved"].ID:
		return _Contract.ParseOperatorRemoved(log)
	case _Contract.abi.Events["OwnershipTransferred"].ID:
		return _Contract.ParseOwnershipTransferred(log)
	case _Contract.abi.Events["RatioNotUpdated"].ID:
		return _Contract.ParseRatioNotUpdated(log)
	case _Contract.abi.Events["RatioThresholdChanged"].ID:
		return _Contract.ParseRatioThresholdChanged(log)
	case _Contract.abi.Events["RatioUpdated"].ID:
		return _Contract.ParseRatioUpdated(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ContractInitialized) Topic() common.Hash {
	return common.HexToHash("0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498")
}

func (ContractOperatorAdded) Topic() common.Hash {
	return common.HexToHash("0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d")
}

func (ContractOperatorRemoved) Topic() common.Hash {
	return common.HexToHash("0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d")
}

func (ContractOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ContractRatioNotUpdated) Topic() common.Hash {
	return common.HexToHash("0x2471a7627ad27128888e46dfc72f5d674c7156d6e99c969a675492a558a0b0e0")
}

func (ContractRatioThresholdChanged) Topic() common.Hash {
	return common.HexToHash("0x661e4cadf2d36ec16a59d60dcfeebe23f9be2aec99852725798a4be99790840e")
}

func (ContractRatioUpdated) Topic() common.Hash {
	return common.HexToHash("0x4c5c23b4efbfea6d16c8453f565e165a02a22cda9a8dc7aac0a66f91d2304da6")
}

func (_Contract *Contract) Address() common.Address {
	return _Contract.address
}

type ContractInterface interface {
	MAXTHRESHOLD(opts *bind.CallOpts) (uint32, error)

	AveragePercentageRate(opts *bind.CallOpts, addr common.Address, day *big.Int) (*big.Int, error)

	GetRatioFor(opts *bind.CallOpts, token common.Address) (*big.Int, error)

	GetRatioThreshold(opts *bind.CallOpts) (*big.Int, error)

	HistoricalRatios(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AddOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error)

	RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	RepairRatioFor(opts *bind.TransactOpts, token common.Address, ratio *big.Int) (*types.Transaction, error)

	SetRatioThreshold(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	UpdateRatioBatch(opts *bind.TransactOpts, addresses []common.Address, ratios []*big.Int) (*types.Transaction, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*ContractInitialized, error)

	FilterOperatorAdded(opts *bind.FilterOpts) (*ContractOperatorAddedIterator, error)

	WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *ContractOperatorAdded) (event.Subscription, error)

	ParseOperatorAdded(log types.Log) (*ContractOperatorAdded, error)

	FilterOperatorRemoved(opts *bind.FilterOpts) (*ContractOperatorRemovedIterator, error)

	WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *ContractOperatorRemoved) (event.Subscription, error)

	ParseOperatorRemoved(log types.Log) (*ContractOperatorRemoved, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error)

	FilterRatioNotUpdated(opts *bind.FilterOpts, tokenAddress []common.Address) (*ContractRatioNotUpdatedIterator, error)

	WatchRatioNotUpdated(opts *bind.WatchOpts, sink chan<- *ContractRatioNotUpdated, tokenAddress []common.Address) (event.Subscription, error)

	ParseRatioNotUpdated(log types.Log) (*ContractRatioNotUpdated, error)

	FilterRatioThresholdChanged(opts *bind.FilterOpts) (*ContractRatioThresholdChangedIterator, error)

	WatchRatioThresholdChanged(opts *bind.WatchOpts, sink chan<- *ContractRatioThresholdChanged) (event.Subscription, error)

	ParseRatioThresholdChanged(log types.Log) (*ContractRatioThresholdChanged, error)

	FilterRatioUpdated(opts *bind.FilterOpts, tokenAddress []common.Address) (*ContractRatioUpdatedIterator, error)

	WatchRatioUpdated(opts *bind.WatchOpts, sink chan<- *ContractRatioUpdated, tokenAddress []common.Address) (event.Subscription, error)

	ParseRatioUpdated(log types.Log) (*ContractRatioUpdated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
