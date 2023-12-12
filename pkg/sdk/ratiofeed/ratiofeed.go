// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ratiofeed

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"failedRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"RatioNotUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RatioThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRatio\",\"type\":\"uint256\"}],\"name\":\"RatioUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"averagePercentageRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getRatioFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRatioThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"historicalRatios\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"lastUpdate\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingConfig\",\"name\":\"stakingConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"repairRatioFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setRatioThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ratios\",\"type\":\"uint256[]\"}],\"name\":\"updateRatioBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610f5a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063a1f1d48d11610066578063a1f1d48d14610106578063b037d56a1461012f578063ba34fa0514610142578063c4d66de81461014a578063ec653c4b1461015d57600080fd5b806308af5431146100985780632364753a146100bd5780632acaaff4146100d25780632ef86a1f146100f3575b600080fd5b6100a36305f5e10081565b60405163ffffffff90911681526020015b60405180910390f35b6100d06100cb366004610b9a565b61019d565b005b6100e56100e0366004610bcb565b6102f2565b6040519081526020016100b4565b6100d0610101366004610bcb565b610488565b6100e5610114366004610bf7565b6001600160a01b031660009081526001602052604090205490565b6100d061013d366004610c66565b6105ce565b6003546100e5565b6100d0610158366004610bf7565b610997565b61018761016b366004610bf7565b60026020526000908152604090206003015464ffffffffff1681565b60405164ffffffffff90911681526020016100b4565b600060029054906101000a90046001600160a01b03166001600160a01b031663732524946040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101f0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102149190610cd1565b6001600160a01b0316336001600160a01b03161461024d5760405162461bcd60e51b815260040161024490610cee565b60405180910390fd5b6305f5e100811080156102605750600081115b6102ac5760405162461bcd60e51b815260206004820152601f60248201527f77726f6e672076616c756520666f7220726174696f207468726573686f6c64006044820152606401610244565b600380549082905560408051828152602081018490527f661e4cadf2d36ec16a59d60dcfeebe23f9be2aec99852725798a4be99790840e91015b60405180910390a15050565b600080821180156103035750600882105b61034f5760405162461bcd60e51b815260206004820152601960248201527f6461792073686f756c642062652066726f6d203120746f2037000000000000006044820152606401610244565b6001600160a01b0383166000908152600260205260408120805490916001600160401b03909116908260086103848785610d5c565b61038e9190610d89565b610399906001610d9d565b600981106103a9576103a9610d30565b60048104909101546001600160401b036008600390931683026101000a90910416915060009084906103db9085610db5565b6103e6906001610ddb565b6001600160401b0316600981106103ff576103ff610d30565b600491828204019190066008029054906101000a90046001600160401b03166001600160401b031690508082101561043e576000945050505050610482565b6104488683610e06565b6104528284610d5c565b6104659068056bc75e2d63100000610e06565b6104719061016d610e06565b61047b9190610e25565b9450505050505b92915050565b600060029054906101000a90046001600160a01b03166001600160a01b031663732524946040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ff9190610cd1565b6001600160a01b0316336001600160a01b03161461052f5760405162461bcd60e51b815260040161024490610cee565b8060000361056f5760405162461bcd60e51b815260206004820152600d60248201526c726174696f206973207a65726f60981b6044820152606401610244565b6001600160a01b038216600081815260016020908152604091829020805490859055825181815291820185905292917f4c5c23b4efbfea6d16c8453f565e165a02a22cda9a8dc7aac0a66f91d2304da6910160405180910390a2505050565b600060029054906101000a90046001600160a01b03166001600160a01b0316632ec338ba6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610621573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106459190610cd1565b6001600160a01b0316336001600160a01b0316146106a55760405162461bcd60e51b815260206004820181905260248201527f526174696f466565643a206f6e6c79206f70657261746f7220616c6c6f7765646044820152606401610244565b8281146106eb5760405162461bcd60e51b8152602060048201526014602482015273636f7272757074656420726174696f206461746160601b6044820152606401610244565b60006003541161073d5760405162461bcd60e51b815260206004820152601a60248201527f726174696f207468726573686f6c64206973206e6f74207365740000000000006044820152606401610244565b60005b8381101561099057600085858381811061075c5761075c610d30565b90506020020160208101906107719190610bf7565b6001600160a01b038116600090815260046020908152604080832054600190925282205492935091908686868181106107ac576107ac610d30565b9050602002013590506000806107c3858486610abf565b915091508161081a57856001600160a01b03167f2471a7627ad27128888e46dfc72f5d674c7156d6e99c969a675492a558a0b0e08483604051610807929190610e39565b60405180910390a250505050505061097e565b6001600160a01b03861660008181526001602090815260409182902086905581518781529081018690527f4c5c23b4efbfea6d16c8453f565e165a02a22cda9a8dc7aac0a66f91d2304da6910160405180910390a26001600160a01b038616600090815260046020908152604080832064ffffffffff42818116909255600290935292206003810154909262015144926108b5921690610d5c565b11156109765780546001600160401b0316848260086108d5846001610ddb565b6108df9190610db5565b6108ea906001610ddb565b6001600160401b03166009811061090357610903610d30565b600491828204019190066008026101000a8154816001600160401b0302191690836001600160401b0316021790555080600161093f9190610ddb565b825467ffffffffffffffff19166001600160401b03919091161782555060038101805464ffffffffff19164264ffffffffff161790555b505050505050505b8061098881610e96565b915050610740565b5050505050565b600054610100900460ff16158080156109b75750600054600160ff909116105b806109d15750303b1580156109d1575060005460ff166001145b610a345760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610244565b6000805460ff191660011790558015610a57576000805461ff0019166101001790555b6000805462010000600160b01b031916620100006001600160a01b038516021790558015610abb576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020016102e6565b5050565b6000606082600003610ad45760019150610b92565b61a8c0610ae18642610d5c565b1015610b0a5781604051806060016040528060288152602001610eb06028913991509150610b92565b82841115610b355781604051806060016040528060248152602001610ed86024913991509150610b92565b6003546000906305f5e10090610b4b9086610e06565b610b559190610e25565b9050610b618185610d5c565b851015610b8c5782604051806060016040528060298152602001610efc602991399250925050610b92565b60019250505b935093915050565b600060208284031215610bac57600080fd5b5035919050565b6001600160a01b0381168114610bc857600080fd5b50565b60008060408385031215610bde57600080fd5b8235610be981610bb3565b946020939093013593505050565b600060208284031215610c0957600080fd5b8135610c1481610bb3565b9392505050565b60008083601f840112610c2d57600080fd5b5081356001600160401b03811115610c4457600080fd5b6020830191508360208260051b8501011115610c5f57600080fd5b9250929050565b60008060008060408587031215610c7c57600080fd5b84356001600160401b0380821115610c9357600080fd5b610c9f88838901610c1b565b90965094506020870135915080821115610cb857600080fd5b50610cc587828801610c1b565b95989497509550505050565b600060208284031215610ce357600080fd5b8151610c1481610bb3565b60208082526022908201527f526174696f466565643a206f6e6c7920676f7665726e616e636520616c6c6f77604082015261195960f21b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082821015610d6e57610d6e610d46565b500390565b634e487b7160e01b600052601260045260246000fd5b600082610d9857610d98610d73565b500690565b60008219821115610db057610db0610d46565b500190565b60006001600160401b0380841680610dcf57610dcf610d73565b92169190910692915050565b60006001600160401b03808316818516808303821115610dfd57610dfd610d46565b01949350505050565b6000816000190483118215151615610e2057610e20610d46565b500290565b600082610e3457610e34610d73565b500490565b82815260006020604081840152835180604085015260005b81811015610e6d57858101830151858201606001528201610e51565b81811115610e7f576000606083870101525b50601f01601f191692909201606001949350505050565b600060018201610ea857610ea8610d46565b506001019056fe726174696f207761732075706461746564206c657373207468616e20313220686f7572732061676f6e657720726174696f2063616e6e6f742062652067726561746572207468616e206f6c646e657720726174696f20746f6f206c6f772c206e6f7420696e207468726573686f6c642072616e6765a2646970667358221220c0e1f8bff704f076d42df0aa2c9d6f93ffa9be6f8ec5e7b56513a76ae3b5cb1964736f6c634300080f0033",
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

func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, stakingConfig common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", stakingConfig)
}

func (_Contract *ContractSession) Initialize(stakingConfig common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, stakingConfig)
}

func (_Contract *ContractTransactorSession) Initialize(stakingConfig common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, stakingConfig)
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

	Initialize(opts *bind.TransactOpts, stakingConfig common.Address) (*types.Transaction, error)

	RepairRatioFor(opts *bind.TransactOpts, token common.Address, ratio *big.Int) (*types.Transaction, error)

	SetRatioThreshold(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error)

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
