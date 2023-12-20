// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feecollector

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
	ABI: "[{\"inputs\":[],\"name\":\"CommissionNotInRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"FeeCollectorTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernanceAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOperatorAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"prevValue\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newValue\",\"type\":\"uint16\"}],\"name\":\"CommissionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_COMMISSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commission\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIProtocolConfig\",\"name\":\"config\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"commission_\",\"type\":\"uint16\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newValue\",\"type\":\"uint16\"}],\"name\":\"setCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506109188061001d5f395ff3fe608060405260043610610057575f3560e01c80630572b0cc146100975780633ccfd60b146100be578063ae97dde8146100d4578063c594cc65146100fc578063e14891911461011b578063f39e69a41461013c575f80fd5b366100935760405134815233907f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f885258749060200160405180910390a2005b5f80fd5b3480156100a2575f80fd5b506100ab61015b565b6040519081526020015b60405180910390f35b3480156100c9575f80fd5b506100d261016b565b005b3480156100df575f80fd5b506100e961271081565b60405161ffff90911681526020016100b5565b348015610107575f80fd5b506100d26101163660046107b8565b6103dc565b348015610126575f80fd5b506033546100e990600160a01b900461ffff1681565b348015610147575f80fd5b506100d26101563660046107ec565b6104ca565b5f610165476105fb565b92915050565b610173610638565b4761271081106103d0575f80610188836105fb565b915091505f60335f9054906101000a90046001600160a01b03166001600160a01b0316637745165b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101dd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610201919061081f565b90505f60335f9054906101000a90046001600160a01b03166001600160a01b0316633b19e84a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610254573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610278919061081f565b90505f826001600160a01b0316846040515f6040518083038185875af1925050503d805f81146102c3576040519150601f19603f3d011682016040523d82523d5f602084013e6102c8565b606091505b50509050806102fa5760405163035106a760e31b81526001600160a01b03841660048201526024015b60405180910390fd5b6040516001600160a01b0383169086905f81818185875af1925050503d805f8114610340576040519150601f19603f3d011682016040523d82523d5f602084013e610345565b606091505b505080915050806103745760405163035106a760e31b81526001600160a01b03831660048201526024016102f1565b816001600160a01b0316836001600160a01b03167f91fb9d98b786c57d74c099ccd2beca1739e9f6a81fb49001ca465c4b7591bbe286886040516103c2929190918252602082015260400190565b60405180910390a350505050505b506103da60018055565b565b60335f9054906101000a90046001600160a01b03166001600160a01b031663289b3c0d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561042c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610450919061081f565b6001600160a01b0316336001600160a01b0316146104be5760405162461bcd60e51b815260206004820152602560248201527f466565436f6c6c6563746f723a206f6e6c7920676f7665726e616e636520616c6044820152641b1bddd95960da1b60648201526084016102f1565b6104c781610697565b50565b5f54610100900460ff16158080156104e857505f54600160ff909116105b806105015750303b15801561050157505f5460ff166001145b6105645760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102f1565b5f805460ff191660011790558015610585575f805461ff0019166101001790555b61058d610728565b603380546001600160a01b0319166001600160a01b0385161790556105b182610756565b80156105f6575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6033545f9081906127109061061b90600160a01b900461ffff168561084e565b6106259190610865565b91506106318284610884565b9050915091565b60026001540361068a5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016102f1565b6002600155565b60018055565b61271061ffff8216106106bd57604051631994455560e11b815260040160405180910390fd5b6033546040805161ffff600160a01b9093048316815291831660208301527f07d1ecb0b6fd1c556347c57a4035549a933093e3e5a0a63418dd3589c82749de910160405180910390a16033805461ffff909216600160a01b0261ffff60a01b19909216919091179055565b5f54610100900460ff1661074e5760405162461bcd60e51b81526004016102f190610897565b6103da61077c565b5f54610100900460ff166104be5760405162461bcd60e51b81526004016102f190610897565b5f54610100900460ff166106915760405162461bcd60e51b81526004016102f190610897565b803561ffff811681146107b3575f80fd5b919050565b5f602082840312156107c8575f80fd5b6107d1826107a2565b9392505050565b6001600160a01b03811681146104c7575f80fd5b5f80604083850312156107fd575f80fd5b8235610808816107d8565b9150610816602084016107a2565b90509250929050565b5f6020828403121561082f575f80fd5b81516107d1816107d8565b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176101655761016561083a565b5f8261087f57634e487b7160e01b5f52601260045260245ffd5b500490565b818103818111156101655761016561083a565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b60608201526080019056fea264697066735822122020e521f18ea0a5f0576b2b373d74bc31e2e178df02a118cdea8bceddf2fbfae564736f6c63430008140033",
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

func (_Contract *ContractCaller) MAXCOMMISSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "MAX_COMMISSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_Contract *ContractSession) MAXCOMMISSION() (uint16, error) {
	return _Contract.Contract.MAXCOMMISSION(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) MAXCOMMISSION() (uint16, error) {
	return _Contract.Contract.MAXCOMMISSION(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) Commission(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "commission")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

func (_Contract *ContractSession) Commission() (uint16, error) {
	return _Contract.Contract.Commission(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) Commission() (uint16, error) {
	return _Contract.Contract.Commission(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) GetRewards() (*big.Int, error) {
	return _Contract.Contract.GetRewards(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetRewards() (*big.Int, error) {
	return _Contract.Contract.GetRewards(&_Contract.CallOpts)
}

func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, config common.Address, commission_ uint16) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", config, commission_)
}

func (_Contract *ContractSession) Initialize(config common.Address, commission_ uint16) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, config, commission_)
}

func (_Contract *ContractTransactorSession) Initialize(config common.Address, commission_ uint16) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, config, commission_)
}

func (_Contract *ContractTransactor) SetCommission(opts *bind.TransactOpts, newValue uint16) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setCommission", newValue)
}

func (_Contract *ContractSession) SetCommission(newValue uint16) (*types.Transaction, error) {
	return _Contract.Contract.SetCommission(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetCommission(newValue uint16) (*types.Transaction, error) {
	return _Contract.Contract.SetCommission(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw")
}

func (_Contract *ContractSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil)
}

func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

type ContractCommissionChangedIterator struct {
	Event *ContractCommissionChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractCommissionChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCommissionChanged)
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
		it.Event = new(ContractCommissionChanged)
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

func (it *ContractCommissionChangedIterator) Error() error {
	return it.fail
}

func (it *ContractCommissionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractCommissionChanged struct {
	PrevValue uint16
	NewValue  uint16
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterCommissionChanged(opts *bind.FilterOpts) (*ContractCommissionChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CommissionChanged")
	if err != nil {
		return nil, err
	}
	return &ContractCommissionChangedIterator{contract: _Contract.contract, event: "CommissionChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchCommissionChanged(opts *bind.WatchOpts, sink chan<- *ContractCommissionChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CommissionChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractCommissionChanged)
				if err := _Contract.contract.UnpackLog(event, "CommissionChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseCommissionChanged(log types.Log) (*ContractCommissionChanged, error) {
	event := new(ContractCommissionChanged)
	if err := _Contract.contract.UnpackLog(event, "CommissionChanged", log); err != nil {
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
	Sender common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_Contract *ContractFilterer) FilterReceived(opts *bind.FilterOpts, sender []common.Address) (*ContractReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Received", senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractReceivedIterator{contract: _Contract.contract, event: "Received", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *ContractReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Received", senderRule)
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

type ContractWithdrawnIterator struct {
	Event *ContractWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdrawn)
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
		it.Event = new(ContractWithdrawn)
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

func (it *ContractWithdrawnIterator) Error() error {
	return it.fail
}

func (it *ContractWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractWithdrawn struct {
	Pool     common.Address
	Treasury common.Address
	Rewards  *big.Int
	Fee      *big.Int
	Raw      types.Log
}

func (_Contract *ContractFilterer) FilterWithdrawn(opts *bind.FilterOpts, pool []common.Address, treasury []common.Address) (*ContractWithdrawnIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Withdrawn", poolRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawnIterator{contract: _Contract.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractWithdrawn, pool []common.Address, treasury []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var treasuryRule []interface{}
	for _, treasuryItem := range treasury {
		treasuryRule = append(treasuryRule, treasuryItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Withdrawn", poolRule, treasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractWithdrawn)
				if err := _Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

func (_Contract *ContractFilterer) ParseWithdrawn(log types.Log) (*ContractWithdrawn, error) {
	event := new(ContractWithdrawn)
	if err := _Contract.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Contract *Contract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Contract.abi.Events["CommissionChanged"].ID:
		return _Contract.ParseCommissionChanged(log)
	case _Contract.abi.Events["Initialized"].ID:
		return _Contract.ParseInitialized(log)
	case _Contract.abi.Events["Received"].ID:
		return _Contract.ParseReceived(log)
	case _Contract.abi.Events["Withdrawn"].ID:
		return _Contract.ParseWithdrawn(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ContractCommissionChanged) Topic() common.Hash {
	return common.HexToHash("0x07d1ecb0b6fd1c556347c57a4035549a933093e3e5a0a63418dd3589c82749de")
}

func (ContractInitialized) Topic() common.Hash {
	return common.HexToHash("0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498")
}

func (ContractReceived) Topic() common.Hash {
	return common.HexToHash("0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874")
}

func (ContractWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x91fb9d98b786c57d74c099ccd2beca1739e9f6a81fb49001ca465c4b7591bbe2")
}

func (_Contract *Contract) Address() common.Address {
	return _Contract.address
}

type ContractInterface interface {
	MAXCOMMISSION(opts *bind.CallOpts) (uint16, error)

	Commission(opts *bind.CallOpts) (uint16, error)

	GetRewards(opts *bind.CallOpts) (*big.Int, error)

	Initialize(opts *bind.TransactOpts, config common.Address, commission_ uint16) (*types.Transaction, error)

	SetCommission(opts *bind.TransactOpts, newValue uint16) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts) (*types.Transaction, error)

	Receive(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterCommissionChanged(opts *bind.FilterOpts) (*ContractCommissionChangedIterator, error)

	WatchCommissionChanged(opts *bind.WatchOpts, sink chan<- *ContractCommissionChanged) (event.Subscription, error)

	ParseCommissionChanged(log types.Log) (*ContractCommissionChanged, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*ContractInitialized, error)

	FilterReceived(opts *bind.FilterOpts, sender []common.Address) (*ContractReceivedIterator, error)

	WatchReceived(opts *bind.WatchOpts, sink chan<- *ContractReceived, sender []common.Address) (event.Subscription, error)

	ParseReceived(log types.Log) (*ContractReceived, error)

	FilterWithdrawn(opts *bind.FilterOpts, pool []common.Address, treasury []common.Address) (*ContractWithdrawnIterator, error)

	WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractWithdrawn, pool []common.Address, treasury []common.Address) (event.Subscription, error)

	ParseWithdrawn(log types.Log) (*ContractWithdrawn, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
