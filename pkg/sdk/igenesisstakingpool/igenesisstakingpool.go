// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igenesisstakingpool

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CertificateTokenChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinimumStakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinimumUnstakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"PoolOnGoing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getFreeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinUnstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"deposit_data_roots\",\"type\":\"bytes32[]\"}],\"name\":\"pushToBeaconMulti\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setCertificateToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMinimumStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMinimumUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeCerts\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

func (_Contract *ContractTransactor) PushToBeaconMulti(opts *bind.TransactOpts, pubkeys [][]byte, signatures [][]byte, deposit_data_roots [][32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "pushToBeaconMulti", pubkeys, signatures, deposit_data_roots)
}

func (_Contract *ContractSession) PushToBeaconMulti(pubkeys [][]byte, signatures [][]byte, deposit_data_roots [][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.PushToBeaconMulti(&_Contract.TransactOpts, pubkeys, signatures, deposit_data_roots)
}

func (_Contract *ContractTransactorSession) PushToBeaconMulti(pubkeys [][]byte, signatures [][]byte, deposit_data_roots [][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.PushToBeaconMulti(&_Contract.TransactOpts, pubkeys, signatures, deposit_data_roots)
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
	case _Contract.abi.Events["MinimumStakeChanged"].ID:
		return _Contract.ParseMinimumStakeChanged(log)
	case _Contract.abi.Events["MinimumUnstakeChanged"].ID:
		return _Contract.ParseMinimumUnstakeChanged(log)
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

func (ContractMinimumStakeChanged) Topic() common.Hash {
	return common.HexToHash("0xdc4a0b2dc1fa27da98de2ac6f8fa373b4be405e1bf69fc3976597b6d56b79abc")
}

func (ContractMinimumUnstakeChanged) Topic() common.Hash {
	return common.HexToHash("0x976e4c5a2502181884199a44267e9ba339f2d071066e4f0ae97600ce4975935c")
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
	GetFreeBalance(opts *bind.CallOpts) (*big.Int, error)

	GetMinStake(opts *bind.CallOpts) (*big.Int, error)

	GetMinUnstake(opts *bind.CallOpts) (*big.Int, error)

	PushToBeaconMulti(opts *bind.TransactOpts, pubkeys [][]byte, signatures [][]byte, deposit_data_roots [][32]byte) (*types.Transaction, error)

	SetCertificateToken(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	SetMinimumStake(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error)

	SetMinimumUnstake(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error)

	StakeCerts(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterCertificateTokenChanged(opts *bind.FilterOpts) (*ContractCertificateTokenChangedIterator, error)

	WatchCertificateTokenChanged(opts *bind.WatchOpts, sink chan<- *ContractCertificateTokenChanged) (event.Subscription, error)

	ParseCertificateTokenChanged(log types.Log) (*ContractCertificateTokenChanged, error)

	FilterMinimumStakeChanged(opts *bind.FilterOpts) (*ContractMinimumStakeChangedIterator, error)

	WatchMinimumStakeChanged(opts *bind.WatchOpts, sink chan<- *ContractMinimumStakeChanged) (event.Subscription, error)

	ParseMinimumStakeChanged(log types.Log) (*ContractMinimumStakeChanged, error)

	FilterMinimumUnstakeChanged(opts *bind.FilterOpts) (*ContractMinimumUnstakeChangedIterator, error)

	WatchMinimumUnstakeChanged(opts *bind.WatchOpts, sink chan<- *ContractMinimumUnstakeChanged) (event.Subscription, error)

	ParseMinimumUnstakeChanged(log types.Log) (*ContractMinimumUnstakeChanged, error)

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
