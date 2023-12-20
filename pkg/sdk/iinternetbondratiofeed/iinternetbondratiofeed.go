// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iinternetbondratiofeed

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"failedRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"RatioNotUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"RatioThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRatio\",\"type\":\"uint256\"}],\"name\":\"RatioUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getRatioFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ratios\",\"type\":\"uint256[]\"}],\"name\":\"updateRatioBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

func (_Contract *ContractCaller) GetRatioFor(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRatioFor", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) GetRatioFor(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.GetRatioFor(&_Contract.CallOpts, arg0)
}

func (_Contract *ContractCallerSession) GetRatioFor(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.GetRatioFor(&_Contract.CallOpts, arg0)
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
	GetRatioFor(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	UpdateRatioBatch(opts *bind.TransactOpts, addresses []common.Address, ratios []*big.Int) (*types.Transaction, error)

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
