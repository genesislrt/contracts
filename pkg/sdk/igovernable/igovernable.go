// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igovernable

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"GovernanceAddressChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getGovernanceAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setGovernanceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

func (_Contract *ContractTransactor) SetGovernanceAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGovernanceAddress", newValue)
}

func (_Contract *ContractSession) SetGovernanceAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetGovernanceAddress(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetGovernanceAddress(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetGovernanceAddress(&_Contract.TransactOpts, newValue)
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

func (_Contract *Contract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Contract.abi.Events["GovernanceAddressChanged"].ID:
		return _Contract.ParseGovernanceAddressChanged(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ContractGovernanceAddressChanged) Topic() common.Hash {
	return common.HexToHash("0x023588d3d1dcaad34e471c9e01b616b794156174bc693539c8fe15c0bfd5d826")
}

func (_Contract *Contract) Address() common.Address {
	return _Contract.address
}

type ContractInterface interface {
	GetGovernanceAddress(opts *bind.CallOpts) (common.Address, error)

	SetGovernanceAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	FilterGovernanceAddressChanged(opts *bind.FilterOpts) (*ContractGovernanceAddressChangedIterator, error)

	WatchGovernanceAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractGovernanceAddressChanged) (event.Subscription, error)

	ParseGovernanceAddressChanged(log types.Log) (*ContractGovernanceAddressChanged, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
