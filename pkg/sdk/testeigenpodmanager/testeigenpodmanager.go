// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testeigenpodmanager

import (
	"errors"
	"math/big"
	"strings"

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
	ABI: "[{\"inputs\":[],\"name\":\"createPod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"depositDataRoot\",\"type\":\"bytes32\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610146806100206000396000f3fe6080604052600436106100295760003560e01c806384d810621461002e5780639b4e46341461003e575b600080fd5b34801561003a57600080fd5b505b005b61003c61004c36600461009c565b5050505050565b60008083601f84011261006557600080fd5b50813567ffffffffffffffff81111561007d57600080fd5b60208301915083602082850101111561009557600080fd5b9250929050565b6000806000806000606086880312156100b457600080fd5b853567ffffffffffffffff808211156100cc57600080fd5b6100d889838a01610053565b909750955060208801359150808211156100f157600080fd5b506100fe88828901610053565b9699959850966040013594935050505056fea264697066735822122070c405a9fc57ba36d9d540ed7ebb15e9ee013259b3ed346d2bef099b91142e9f64736f6c634300080f0033",
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

func (_Contract *ContractTransactor) CreatePod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createPod")
}

func (_Contract *ContractSession) CreatePod() (*types.Transaction, error) {
	return _Contract.Contract.CreatePod(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactorSession) CreatePod() (*types.Transaction, error) {
	return _Contract.Contract.CreatePod(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactor) Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "stake", pubkey, signature, depositDataRoot)
}

func (_Contract *ContractSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Stake(&_Contract.TransactOpts, pubkey, signature, depositDataRoot)
}

func (_Contract *ContractTransactorSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Stake(&_Contract.TransactOpts, pubkey, signature, depositDataRoot)
}

func (_Contract *Contract) Address() common.Address {
	return _Contract.address
}

type ContractInterface interface {
	CreatePod(opts *bind.TransactOpts) (*types.Transaction, error)

	Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error)

	Address() common.Address
}
