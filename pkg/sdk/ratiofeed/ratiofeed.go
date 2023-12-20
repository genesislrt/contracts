// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ratiofeed

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/genesislrt/contracts/abigen/generated"
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
	ABI: "[{\"inputs\":[],\"name\":\"OnlyGovernanceAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOperatorAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"RatioNotUpdated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RatioThresholdNotInRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RatioThresholdNotSet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RatioThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRatio\",\"type\":\"uint256\"}],\"name\":\"RatioUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"day\",\"type\":\"uint8\"}],\"name\":\"averagePercentageRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"historicalRatios\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"lastUpdate\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingConfig\",\"name\":\"stakingConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ratioThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newRatio\",\"type\":\"uint256\"}],\"name\":\"repairRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setRatioThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newRatio\",\"type\":\"uint256\"}],\"name\":\"updateRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610cec8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610090575f3560e01c80633be19c03116100635780633be19c03146101025780637068ca0d1461010a578063754b27071461011d578063c4d66de814610145578063ec653c4b14610158575f80fd5b806308af54311461009457806311ad2955146100b95780632364753a146100ce57806327a0a544146100e1575b5f80fd5b61009f6305f5e10081565b60405163ffffffff90911681526020015b60405180910390f35b6100cc6100c7366004610a57565b610197565b005b6100cc6100dc366004610a81565b6102ff565b6100f46100ef366004610a98565b6103b0565b6040519081526020016100b0565b6003546100f4565b6100cc610118366004610a57565b610557565b6100f461012b366004610ad4565b6001600160a01b03165f9081526001602052604090205490565b6100cc610153366004610ad4565b6107d6565b610181610166366004610ad4565b60026020525f908152604090206003015464ffffffffff1681565b60405164ffffffffff90911681526020016100b0565b5f60029054906101000a90046001600160a01b03166001600160a01b031663289b3c0d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101e7573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061020b9190610af6565b6001600160a01b0316336001600160a01b03161461023c5760405163e2d4f15f60e01b815260040160405180910390fd5b670de0b6b3a7640000811180610250575080155b1561029257604051632481e5bd60e11b815260206004820152600c60248201526b6e6f7420696e2072616e676560a01b60448201526064015b60405180910390fd5b6001600160a01b0382165f818152600160209081526040918290205482519081529081018490527f4c5c23b4efbfea6d16c8453f565e165a02a22cda9a8dc7aac0a66f91d2304da6910160405180910390a26001600160a01b039091165f90815260016020526040902055565b5f60029054906101000a90046001600160a01b03166001600160a01b031663289b3c0d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561034f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103739190610af6565b6001600160a01b0316336001600160a01b0316146103a45760405163e2d4f15f60e01b815260040160405180910390fd5b6103ad816108fb565b50565b5f808260ff161180156103c6575060088260ff16105b6104125760405162461bcd60e51b815260206004820152601960248201527f6461792073686f756c642062652066726f6d203120746f2037000000000000006044820152606401610289565b6001600160a01b0383165f908152600260205260408120805490916001600160401b039091169082600861044960ff881685610b39565b6104539190610b74565b61045e906001610b99565b6001600160401b03166009811061047757610477610b11565b60048104909101546001600160401b036008600390931683026101000a9091041691505f9084906104a89085610b74565b6104b3906001610b99565b6001600160401b0316600981106104cc576104cc610b11565b600491828204019190066008029054906101000a90046001600160401b03166001600160401b031690508082101561050a575f945050505050610551565b61051760ff871683610bb9565b6105218284610bd0565b6105349068056bc75e2d63100000610bb9565b6105409061016d610bb9565b61054a9190610be3565b9450505050505b92915050565b5f60029054906101000a90046001600160a01b03166001600160a01b031663e7f43c686040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105a7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105cb9190610af6565b6001600160a01b0316336001600160a01b0316146105fc57604051633734611360e01b815260040160405180910390fd5b6003545f0361061e57604051633c70047d60e01b815260040160405180910390fd5b6001600160a01b0382165f90815260046020908152604080832054600190925282205490918061064f84868561096b565b91509150816106735780604051632481e5bd60e11b81526004016102899190610bf6565b6001600160a01b0386165f8181526001602090815260409182902088905581518681529081018890527f4c5c23b4efbfea6d16c8453f565e165a02a22cda9a8dc7aac0a66f91d2304da6910160405180910390a26001600160a01b0386165f90815260046020908152604080832064ffffffffff428181169092556002909352922060038101549092620151449261070c921690610bd0565b11156107cd5780546001600160401b03168682600861072c846001610b99565b6107369190610b74565b610741906001610b99565b6001600160401b03166009811061075a5761075a610b11565b600491828204019190066008026101000a8154816001600160401b0302191690836001600160401b031602179055508060016107969190610b99565b825467ffffffffffffffff19166001600160401b03919091161782555060038101805464ffffffffff19164264ffffffffff161790555b50505050505050565b5f54610100900460ff16158080156107f457505f54600160ff909116105b8061080d5750303b15801561080d57505f5460ff166001145b6108705760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610289565b5f805460ff191660011790558015610891575f805461ff0019166101001790555b5f805462010000600160b01b031916620100006001600160a01b0385160217905580156108f7575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6305f5e1008110158061090c575080155b1561092a57604051630a57072960e01b815260040160405180910390fd5b60035460408051918252602082018390527f661e4cadf2d36ec16a59d60dcfeebe23f9be2aec99852725798a4be99790840e910160405180910390a1600355565b5f6060825f0361097e5760019150610a3b565b61a8c061098b8642610bd0565b10156109b45781604051806060016040528060288152602001610c426028913991509150610a3b565b828411156109df5781604051806060016040528060248152602001610c6a6024913991509150610a3b565b6003545f906305f5e100906109f49086610bb9565b6109fe9190610be3565b9050610a0a8185610bd0565b851015610a355782604051806060016040528060298152602001610c8e602991399250925050610a3b565b60019250505b935093915050565b6001600160a01b03811681146103ad575f80fd5b5f8060408385031215610a68575f80fd5b8235610a7381610a43565b946020939093013593505050565b5f60208284031215610a91575f80fd5b5035919050565b5f8060408385031215610aa9575f80fd5b8235610ab481610a43565b9150602083013560ff81168114610ac9575f80fd5b809150509250929050565b5f60208284031215610ae4575f80fd5b8135610aef81610a43565b9392505050565b5f60208284031215610b06575f80fd5b8151610aef81610a43565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b6001600160401b03828116828216039080821115610b5957610b59610b25565b5092915050565b634e487b7160e01b5f52601260045260245ffd5b5f6001600160401b0380841680610b8d57610b8d610b60565b92169190910692915050565b6001600160401b03818116838216019080821115610b5957610b59610b25565b808202811582820484141761055157610551610b25565b8181038181111561055157610551610b25565b5f82610bf157610bf1610b60565b500490565b5f6020808352835180828501525f5b81811015610c2157858101830151858201604001528201610c05565b505f604082860101526040601f19601f830116850101925050509291505056fe726174696f207761732075706461746564206c657373207468616e20313220686f7572732061676f6e657720726174696f2063616e6e6f742062652067726561746572207468616e206f6c646e657720726174696f20746f6f206c6f772c206e6f7420696e207468726573686f6c642072616e6765a2646970667358221220ecaa0166668d8ce574ecc617a8ddcfea619f4aacb4b92e32fe524904a4abb0ba64736f6c63430008140033",
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

func (_Contract *ContractCaller) AveragePercentageRate(opts *bind.CallOpts, token common.Address, day uint8) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "averagePercentageRate", token, day)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) AveragePercentageRate(token common.Address, day uint8) (*big.Int, error) {
	return _Contract.Contract.AveragePercentageRate(&_Contract.CallOpts, token, day)
}

func (_Contract *ContractCallerSession) AveragePercentageRate(token common.Address, day uint8) (*big.Int, error) {
	return _Contract.Contract.AveragePercentageRate(&_Contract.CallOpts, token, day)
}

func (_Contract *ContractCaller) GetRatio(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRatio", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) GetRatio(token common.Address) (*big.Int, error) {
	return _Contract.Contract.GetRatio(&_Contract.CallOpts, token)
}

func (_Contract *ContractCallerSession) GetRatio(token common.Address) (*big.Int, error) {
	return _Contract.Contract.GetRatio(&_Contract.CallOpts, token)
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

func (_Contract *ContractCaller) RatioThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ratioThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) RatioThreshold() (*big.Int, error) {
	return _Contract.Contract.RatioThreshold(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) RatioThreshold() (*big.Int, error) {
	return _Contract.Contract.RatioThreshold(&_Contract.CallOpts)
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

func (_Contract *ContractTransactor) RepairRatio(opts *bind.TransactOpts, token common.Address, newRatio *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "repairRatio", token, newRatio)
}

func (_Contract *ContractSession) RepairRatio(token common.Address, newRatio *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RepairRatio(&_Contract.TransactOpts, token, newRatio)
}

func (_Contract *ContractTransactorSession) RepairRatio(token common.Address, newRatio *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.RepairRatio(&_Contract.TransactOpts, token, newRatio)
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

func (_Contract *ContractTransactor) UpdateRatio(opts *bind.TransactOpts, token common.Address, newRatio *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateRatio", token, newRatio)
}

func (_Contract *ContractSession) UpdateRatio(token common.Address, newRatio *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRatio(&_Contract.TransactOpts, token, newRatio)
}

func (_Contract *ContractTransactorSession) UpdateRatio(token common.Address, newRatio *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRatio(&_Contract.TransactOpts, token, newRatio)
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

	AveragePercentageRate(opts *bind.CallOpts, token common.Address, day uint8) (*big.Int, error)

	GetRatio(opts *bind.CallOpts, token common.Address) (*big.Int, error)

	HistoricalRatios(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	RatioThreshold(opts *bind.CallOpts) (*big.Int, error)

	Initialize(opts *bind.TransactOpts, stakingConfig common.Address) (*types.Transaction, error)

	RepairRatio(opts *bind.TransactOpts, token common.Address, newRatio *big.Int) (*types.Transaction, error)

	SetRatioThreshold(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error)

	UpdateRatio(opts *bind.TransactOpts, token common.Address, newRatio *big.Int) (*types.Transaction, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*ContractInitialized, error)

	FilterRatioThresholdChanged(opts *bind.FilterOpts) (*ContractRatioThresholdChangedIterator, error)

	WatchRatioThresholdChanged(opts *bind.WatchOpts, sink chan<- *ContractRatioThresholdChanged) (event.Subscription, error)

	ParseRatioThresholdChanged(log types.Log) (*ContractRatioThresholdChanged, error)

	FilterRatioUpdated(opts *bind.FilterOpts, tokenAddress []common.Address) (*ContractRatioUpdatedIterator, error)

	WatchRatioUpdated(opts *bind.WatchOpts, sink chan<- *ContractRatioUpdated, tokenAddress []common.Address) (event.Subscription, error)

	ParseRatioUpdated(log types.Log) (*ContractRatioUpdated, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
