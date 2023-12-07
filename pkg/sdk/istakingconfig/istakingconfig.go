// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package istakingconfig

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CertTokenAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"EigenManagerAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"GovernanceAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinStakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinUnstakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"OperatorAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"RatioFeedAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"StakingPoolAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"TreasuryAddressChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getCertTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEigenPodManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGovernanceAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinUnstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperatorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRatioFeedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setCertTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setGovernanceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMinStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMinUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setOperatorAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setRatioFeedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setStakingPoolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setTreasuryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

var ContractABI = ContractMetaData.ABI

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

func (_Contract *ContractCaller) GetCertTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCertTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetCertTokenAddress() (common.Address, error) {
	return _Contract.Contract.GetCertTokenAddress(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetCertTokenAddress() (common.Address, error) {
	return _Contract.Contract.GetCertTokenAddress(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetEigenPodManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEigenPodManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetEigenPodManagerAddress() (common.Address, error) {
	return _Contract.Contract.GetEigenPodManagerAddress(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetEigenPodManagerAddress() (common.Address, error) {
	return _Contract.Contract.GetEigenPodManagerAddress(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetGovernanceAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getGovernanceAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetGovernanceAddress() (common.Address, error) {
	return _Contract.Contract.GetGovernanceAddress(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetGovernanceAddress() (common.Address, error) {
	return _Contract.Contract.GetGovernanceAddress(&_Contract.CallOpts)
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

func (_Contract *ContractCaller) GetOperatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetOperatorAddress() (common.Address, error) {
	return _Contract.Contract.GetOperatorAddress(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetOperatorAddress() (common.Address, error) {
	return _Contract.Contract.GetOperatorAddress(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetRatioFeedAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRatioFeedAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetRatioFeedAddress() (common.Address, error) {
	return _Contract.Contract.GetRatioFeedAddress(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetRatioFeedAddress() (common.Address, error) {
	return _Contract.Contract.GetRatioFeedAddress(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetStakingPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getStakingPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetStakingPoolAddress() (common.Address, error) {
	return _Contract.Contract.GetStakingPoolAddress(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetStakingPoolAddress() (common.Address, error) {
	return _Contract.Contract.GetStakingPoolAddress(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetTreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTreasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetTreasuryAddress() (common.Address, error) {
	return _Contract.Contract.GetTreasuryAddress(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetTreasuryAddress() (common.Address, error) {
	return _Contract.Contract.GetTreasuryAddress(&_Contract.CallOpts)
}

func (_Contract *ContractTransactor) SetCertTokenAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setCertTokenAddress", newValue)
}

func (_Contract *ContractSession) SetCertTokenAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetCertTokenAddress(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetCertTokenAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetCertTokenAddress(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) SetGovernanceAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGovernanceAddress", newValue)
}

func (_Contract *ContractSession) SetGovernanceAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetGovernanceAddress(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetGovernanceAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetGovernanceAddress(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) SetMinStake(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMinStake", newValue)
}

func (_Contract *ContractSession) SetMinStake(newValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinStake(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetMinStake(newValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinStake(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) SetMinUnstake(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMinUnstake", newValue)
}

func (_Contract *ContractSession) SetMinUnstake(newValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinUnstake(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetMinUnstake(newValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinUnstake(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) SetOperatorAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setOperatorAddress", newValue)
}

func (_Contract *ContractSession) SetOperatorAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOperatorAddress(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetOperatorAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOperatorAddress(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) SetRatioFeedAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setRatioFeedAddress", newValue)
}

func (_Contract *ContractSession) SetRatioFeedAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetRatioFeedAddress(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetRatioFeedAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetRatioFeedAddress(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) SetStakingPoolAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setStakingPoolAddress", newValue)
}

func (_Contract *ContractSession) SetStakingPoolAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetStakingPoolAddress(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetStakingPoolAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetStakingPoolAddress(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) SetTreasuryAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setTreasuryAddress", newValue)
}

func (_Contract *ContractSession) SetTreasuryAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetTreasuryAddress(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetTreasuryAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetTreasuryAddress(&_Contract.TransactOpts, newValue)
}

type ContractCertTokenAddressChangedIterator struct {
	Event *ContractCertTokenAddressChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractCertTokenAddressChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCertTokenAddressChanged)
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
		it.Event = new(ContractCertTokenAddressChanged)
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

func (it *ContractCertTokenAddressChangedIterator) Error() error {
	return it.fail
}

func (it *ContractCertTokenAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractCertTokenAddressChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterCertTokenAddressChanged(opts *bind.FilterOpts) (*ContractCertTokenAddressChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CertTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return &ContractCertTokenAddressChangedIterator{contract: _Contract.contract, event: "CertTokenAddressChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchCertTokenAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractCertTokenAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CertTokenAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractCertTokenAddressChanged)
				if err := _Contract.contract.UnpackLog(event, "CertTokenAddressChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseCertTokenAddressChanged(log types.Log) (*ContractCertTokenAddressChanged, error) {
	event := new(ContractCertTokenAddressChanged)
	if err := _Contract.contract.UnpackLog(event, "CertTokenAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractEigenManagerAddressChangedIterator struct {
	Event *ContractEigenManagerAddressChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractEigenManagerAddressChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEigenManagerAddressChanged)
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
		it.Event = new(ContractEigenManagerAddressChanged)
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

func (it *ContractEigenManagerAddressChangedIterator) Error() error {
	return it.fail
}

func (it *ContractEigenManagerAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractEigenManagerAddressChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterEigenManagerAddressChanged(opts *bind.FilterOpts) (*ContractEigenManagerAddressChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EigenManagerAddressChanged")
	if err != nil {
		return nil, err
	}
	return &ContractEigenManagerAddressChangedIterator{contract: _Contract.contract, event: "EigenManagerAddressChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchEigenManagerAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractEigenManagerAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EigenManagerAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractEigenManagerAddressChanged)
				if err := _Contract.contract.UnpackLog(event, "EigenManagerAddressChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseEigenManagerAddressChanged(log types.Log) (*ContractEigenManagerAddressChanged, error) {
	event := new(ContractEigenManagerAddressChanged)
	if err := _Contract.contract.UnpackLog(event, "EigenManagerAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractGovernanceAddressChangedIterator struct {
	Event *ContractGovernanceAddressChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractGovernanceAddressChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractGovernanceAddressChanged)
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
		it.Event = new(ContractGovernanceAddressChanged)
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

func (it *ContractGovernanceAddressChangedIterator) Error() error {
	return it.fail
}

func (it *ContractGovernanceAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractGovernanceAddressChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterGovernanceAddressChanged(opts *bind.FilterOpts) (*ContractGovernanceAddressChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "GovernanceAddressChanged")
	if err != nil {
		return nil, err
	}
	return &ContractGovernanceAddressChangedIterator{contract: _Contract.contract, event: "GovernanceAddressChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchGovernanceAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractGovernanceAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "GovernanceAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractGovernanceAddressChanged)
				if err := _Contract.contract.UnpackLog(event, "GovernanceAddressChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseGovernanceAddressChanged(log types.Log) (*ContractGovernanceAddressChanged, error) {
	event := new(ContractGovernanceAddressChanged)
	if err := _Contract.contract.UnpackLog(event, "GovernanceAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractMinStakeChangedIterator struct {
	Event *ContractMinStakeChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractMinStakeChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMinStakeChanged)
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
		it.Event = new(ContractMinStakeChanged)
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

func (it *ContractMinStakeChangedIterator) Error() error {
	return it.fail
}

func (it *ContractMinStakeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractMinStakeChanged struct {
	PrevValue *big.Int
	NewValue  *big.Int
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterMinStakeChanged(opts *bind.FilterOpts) (*ContractMinStakeChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MinStakeChanged")
	if err != nil {
		return nil, err
	}
	return &ContractMinStakeChangedIterator{contract: _Contract.contract, event: "MinStakeChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchMinStakeChanged(opts *bind.WatchOpts, sink chan<- *ContractMinStakeChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MinStakeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractMinStakeChanged)
				if err := _Contract.contract.UnpackLog(event, "MinStakeChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseMinStakeChanged(log types.Log) (*ContractMinStakeChanged, error) {
	event := new(ContractMinStakeChanged)
	if err := _Contract.contract.UnpackLog(event, "MinStakeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractMinUnstakeChangedIterator struct {
	Event *ContractMinUnstakeChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractMinUnstakeChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMinUnstakeChanged)
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
		it.Event = new(ContractMinUnstakeChanged)
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

func (it *ContractMinUnstakeChangedIterator) Error() error {
	return it.fail
}

func (it *ContractMinUnstakeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractMinUnstakeChanged struct {
	PrevValue *big.Int
	NewValue  *big.Int
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterMinUnstakeChanged(opts *bind.FilterOpts) (*ContractMinUnstakeChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MinUnstakeChanged")
	if err != nil {
		return nil, err
	}
	return &ContractMinUnstakeChangedIterator{contract: _Contract.contract, event: "MinUnstakeChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchMinUnstakeChanged(opts *bind.WatchOpts, sink chan<- *ContractMinUnstakeChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MinUnstakeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractMinUnstakeChanged)
				if err := _Contract.contract.UnpackLog(event, "MinUnstakeChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseMinUnstakeChanged(log types.Log) (*ContractMinUnstakeChanged, error) {
	event := new(ContractMinUnstakeChanged)
	if err := _Contract.contract.UnpackLog(event, "MinUnstakeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractOperatorAddressChangedIterator struct {
	Event *ContractOperatorAddressChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractOperatorAddressChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOperatorAddressChanged)
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
		it.Event = new(ContractOperatorAddressChanged)
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

func (it *ContractOperatorAddressChangedIterator) Error() error {
	return it.fail
}

func (it *ContractOperatorAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractOperatorAddressChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterOperatorAddressChanged(opts *bind.FilterOpts) (*ContractOperatorAddressChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OperatorAddressChanged")
	if err != nil {
		return nil, err
	}
	return &ContractOperatorAddressChangedIterator{contract: _Contract.contract, event: "OperatorAddressChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchOperatorAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractOperatorAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OperatorAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractOperatorAddressChanged)
				if err := _Contract.contract.UnpackLog(event, "OperatorAddressChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseOperatorAddressChanged(log types.Log) (*ContractOperatorAddressChanged, error) {
	event := new(ContractOperatorAddressChanged)
	if err := _Contract.contract.UnpackLog(event, "OperatorAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractRatioFeedAddressChangedIterator struct {
	Event *ContractRatioFeedAddressChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractRatioFeedAddressChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRatioFeedAddressChanged)
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
		it.Event = new(ContractRatioFeedAddressChanged)
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

func (it *ContractRatioFeedAddressChangedIterator) Error() error {
	return it.fail
}

func (it *ContractRatioFeedAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractRatioFeedAddressChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterRatioFeedAddressChanged(opts *bind.FilterOpts) (*ContractRatioFeedAddressChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RatioFeedAddressChanged")
	if err != nil {
		return nil, err
	}
	return &ContractRatioFeedAddressChangedIterator{contract: _Contract.contract, event: "RatioFeedAddressChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchRatioFeedAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractRatioFeedAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RatioFeedAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractRatioFeedAddressChanged)
				if err := _Contract.contract.UnpackLog(event, "RatioFeedAddressChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseRatioFeedAddressChanged(log types.Log) (*ContractRatioFeedAddressChanged, error) {
	event := new(ContractRatioFeedAddressChanged)
	if err := _Contract.contract.UnpackLog(event, "RatioFeedAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractStakingPoolAddressChangedIterator struct {
	Event *ContractStakingPoolAddressChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractStakingPoolAddressChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractStakingPoolAddressChanged)
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
		it.Event = new(ContractStakingPoolAddressChanged)
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

func (it *ContractStakingPoolAddressChangedIterator) Error() error {
	return it.fail
}

func (it *ContractStakingPoolAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractStakingPoolAddressChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterStakingPoolAddressChanged(opts *bind.FilterOpts) (*ContractStakingPoolAddressChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "StakingPoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return &ContractStakingPoolAddressChangedIterator{contract: _Contract.contract, event: "StakingPoolAddressChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchStakingPoolAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractStakingPoolAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "StakingPoolAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractStakingPoolAddressChanged)
				if err := _Contract.contract.UnpackLog(event, "StakingPoolAddressChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseStakingPoolAddressChanged(log types.Log) (*ContractStakingPoolAddressChanged, error) {
	event := new(ContractStakingPoolAddressChanged)
	if err := _Contract.contract.UnpackLog(event, "StakingPoolAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractTreasuryAddressChangedIterator struct {
	Event *ContractTreasuryAddressChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractTreasuryAddressChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTreasuryAddressChanged)
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
		it.Event = new(ContractTreasuryAddressChanged)
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

func (it *ContractTreasuryAddressChangedIterator) Error() error {
	return it.fail
}

func (it *ContractTreasuryAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractTreasuryAddressChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterTreasuryAddressChanged(opts *bind.FilterOpts) (*ContractTreasuryAddressChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TreasuryAddressChanged")
	if err != nil {
		return nil, err
	}
	return &ContractTreasuryAddressChangedIterator{contract: _Contract.contract, event: "TreasuryAddressChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchTreasuryAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractTreasuryAddressChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TreasuryAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractTreasuryAddressChanged)
				if err := _Contract.contract.UnpackLog(event, "TreasuryAddressChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseTreasuryAddressChanged(log types.Log) (*ContractTreasuryAddressChanged, error) {
	event := new(ContractTreasuryAddressChanged)
	if err := _Contract.contract.UnpackLog(event, "TreasuryAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Contract *Contract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Contract.abi.Events["CertTokenAddressChanged"].ID:
		return _Contract.ParseCertTokenAddressChanged(log)
	case _Contract.abi.Events["EigenManagerAddressChanged"].ID:
		return _Contract.ParseEigenManagerAddressChanged(log)
	case _Contract.abi.Events["GovernanceAddressChanged"].ID:
		return _Contract.ParseGovernanceAddressChanged(log)
	case _Contract.abi.Events["MinStakeChanged"].ID:
		return _Contract.ParseMinStakeChanged(log)
	case _Contract.abi.Events["MinUnstakeChanged"].ID:
		return _Contract.ParseMinUnstakeChanged(log)
	case _Contract.abi.Events["OperatorAddressChanged"].ID:
		return _Contract.ParseOperatorAddressChanged(log)
	case _Contract.abi.Events["RatioFeedAddressChanged"].ID:
		return _Contract.ParseRatioFeedAddressChanged(log)
	case _Contract.abi.Events["StakingPoolAddressChanged"].ID:
		return _Contract.ParseStakingPoolAddressChanged(log)
	case _Contract.abi.Events["TreasuryAddressChanged"].ID:
		return _Contract.ParseTreasuryAddressChanged(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ContractCertTokenAddressChanged) Topic() common.Hash {
	return common.HexToHash("0x5eb18df2507d560d200b3f0500f2c8f0a1451987559d67ccd2a9f803d729447d")
}

func (ContractEigenManagerAddressChanged) Topic() common.Hash {
	return common.HexToHash("0x68d0770ff8eaea943d0b2f2020510cc23bbedaefd94add71cda3b4a4054b11c3")
}

func (ContractGovernanceAddressChanged) Topic() common.Hash {
	return common.HexToHash("0x023588d3d1dcaad34e471c9e01b616b794156174bc693539c8fe15c0bfd5d826")
}

func (ContractMinStakeChanged) Topic() common.Hash {
	return common.HexToHash("0xca11c8a4c461b60c9f485404c272650c2aaae260b2067d72e9924abb68556593")
}

func (ContractMinUnstakeChanged) Topic() common.Hash {
	return common.HexToHash("0x984016d328adef81f3cc09f8ea9e3420f85d390635be94215c432e83837aa0a2")
}

func (ContractOperatorAddressChanged) Topic() common.Hash {
	return common.HexToHash("0x6d83175096b0de3f5de98ab86d95234fb1a95bd44b8e97696d5b195b27bf42fd")
}

func (ContractRatioFeedAddressChanged) Topic() common.Hash {
	return common.HexToHash("0xeb06cc25da8e6462dfb53cbdcffa08eacbc3833c4f299500a8944ad43e4fd3d2")
}

func (ContractStakingPoolAddressChanged) Topic() common.Hash {
	return common.HexToHash("0x9ce404b121e13a051d915e2e8ffd02c961417a8882f46fec013411e84c7f82d3")
}

func (ContractTreasuryAddressChanged) Topic() common.Hash {
	return common.HexToHash("0x4fc6e7a37aea21888550b60360992adb6a9b3b4da644d63e9f3a420c2d86e282")
}

func (_Contract *Contract) Address() common.Address {
	return _Contract.address
}

type ContractInterface interface {
	GetCertTokenAddress(opts *bind.CallOpts) (common.Address, error)

	GetEigenPodManagerAddress(opts *bind.CallOpts) (common.Address, error)

	GetGovernanceAddress(opts *bind.CallOpts) (common.Address, error)

	GetMinStake(opts *bind.CallOpts) (*big.Int, error)

	GetMinUnstake(opts *bind.CallOpts) (*big.Int, error)

	GetOperatorAddress(opts *bind.CallOpts) (common.Address, error)

	GetRatioFeedAddress(opts *bind.CallOpts) (common.Address, error)

	GetStakingPoolAddress(opts *bind.CallOpts) (common.Address, error)

	GetTreasuryAddress(opts *bind.CallOpts) (common.Address, error)

	SetCertTokenAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	SetGovernanceAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	SetMinStake(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error)

	SetMinUnstake(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error)

	SetOperatorAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	SetRatioFeedAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	SetStakingPoolAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	SetTreasuryAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	FilterCertTokenAddressChanged(opts *bind.FilterOpts) (*ContractCertTokenAddressChangedIterator, error)

	WatchCertTokenAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractCertTokenAddressChanged) (event.Subscription, error)

	ParseCertTokenAddressChanged(log types.Log) (*ContractCertTokenAddressChanged, error)

	FilterEigenManagerAddressChanged(opts *bind.FilterOpts) (*ContractEigenManagerAddressChangedIterator, error)

	WatchEigenManagerAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractEigenManagerAddressChanged) (event.Subscription, error)

	ParseEigenManagerAddressChanged(log types.Log) (*ContractEigenManagerAddressChanged, error)

	FilterGovernanceAddressChanged(opts *bind.FilterOpts) (*ContractGovernanceAddressChangedIterator, error)

	WatchGovernanceAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractGovernanceAddressChanged) (event.Subscription, error)

	ParseGovernanceAddressChanged(log types.Log) (*ContractGovernanceAddressChanged, error)

	FilterMinStakeChanged(opts *bind.FilterOpts) (*ContractMinStakeChangedIterator, error)

	WatchMinStakeChanged(opts *bind.WatchOpts, sink chan<- *ContractMinStakeChanged) (event.Subscription, error)

	ParseMinStakeChanged(log types.Log) (*ContractMinStakeChanged, error)

	FilterMinUnstakeChanged(opts *bind.FilterOpts) (*ContractMinUnstakeChangedIterator, error)

	WatchMinUnstakeChanged(opts *bind.WatchOpts, sink chan<- *ContractMinUnstakeChanged) (event.Subscription, error)

	ParseMinUnstakeChanged(log types.Log) (*ContractMinUnstakeChanged, error)

	FilterOperatorAddressChanged(opts *bind.FilterOpts) (*ContractOperatorAddressChangedIterator, error)

	WatchOperatorAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractOperatorAddressChanged) (event.Subscription, error)

	ParseOperatorAddressChanged(log types.Log) (*ContractOperatorAddressChanged, error)

	FilterRatioFeedAddressChanged(opts *bind.FilterOpts) (*ContractRatioFeedAddressChangedIterator, error)

	WatchRatioFeedAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractRatioFeedAddressChanged) (event.Subscription, error)

	ParseRatioFeedAddressChanged(log types.Log) (*ContractRatioFeedAddressChanged, error)

	FilterStakingPoolAddressChanged(opts *bind.FilterOpts) (*ContractStakingPoolAddressChangedIterator, error)

	WatchStakingPoolAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractStakingPoolAddressChanged) (event.Subscription, error)

	ParseStakingPoolAddressChanged(log types.Log) (*ContractStakingPoolAddressChanged, error)

	FilterTreasuryAddressChanged(opts *bind.FilterOpts) (*ContractTreasuryAddressChangedIterator, error)

	WatchTreasuryAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractTreasuryAddressChanged) (event.Subscription, error)

	ParseTreasuryAddressChanged(log types.Log) (*ContractTreasuryAddressChanged, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
