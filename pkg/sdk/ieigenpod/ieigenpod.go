// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ieigenpod

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

type BeaconChainProofsBalanceUpdateProof struct {
	ValidatorBalanceProof []byte
	ValidatorFieldsProof  []byte
	BalanceRoot           [32]byte
}

type BeaconChainProofsStateRootProof struct {
	BeaconStateRoot [32]byte
	Proof           []byte
}

type BeaconChainProofsWithdrawalProof struct {
	WithdrawalProof                 []byte
	SlotProof                       []byte
	ExecutionPayloadProof           []byte
	TimestampProof                  []byte
	HistoricalSummaryBlockRootProof []byte
	BlockRootIndex                  uint64
	HistoricalSummaryIndex          uint64
	WithdrawalIndex                 uint64
	BlockRoot                       [32]byte
	SlotRoot                        [32]byte
	TimestampRoot                   [32]byte
	ExecutionPayloadRoot            [32]byte
}

type IEigenPodValidatorInfo struct {
	ValidatorIndex                   uint64
	RestakedBalanceGwei              uint64
	MostRecentBalanceUpdateTimestamp uint64
	Status                           uint8
}

var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"EigenPodStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"validatorIndex\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"withdrawalTimestamp\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"withdrawalAmountGwei\",\"type\":\"uint64\"}],\"name\":\"FullWithdrawalRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceived\",\"type\":\"uint256\"}],\"name\":\"NonBeaconChainETHReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"name\":\"NonBeaconChainETHWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"validatorIndex\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"withdrawalTimestamp\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"partialWithdrawalAmountGwei\",\"type\":\"uint64\"}],\"name\":\"PartialWithdrawalRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RestakedBeaconChainETHWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"}],\"name\":\"RestakingActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"validatorIndex\",\"type\":\"uint40\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"balanceTimestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newValidatorBalanceGwei\",\"type\":\"uint64\"}],\"name\":\"ValidatorBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"validatorIndex\",\"type\":\"uint40\"}],\"name\":\"ValidatorRestaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateRestaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenPodManager\",\"outputs\":[{\"internalType\":\"contractIEigenPodManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasRestaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_podOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mostRecentWithdrawalTimestamp\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonBeaconChainETHBalanceWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"podOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"provenWithdrawal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokenList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsToWithdraw\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"recoverTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"validatorPubkeyHash\",\"type\":\"bytes32\"}],\"name\":\"validatorPubkeyHashToInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"validatorIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"restakedBalanceGwei\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mostRecentBalanceUpdateTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIEigenPod.ValidatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pubkeyHash\",\"type\":\"bytes32\"}],\"name\":\"validatorStatus\",\"outputs\":[{\"internalType\":\"enumIEigenPod.VALIDATOR_STATUS\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"oracleTimestamp\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"beaconStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"name\":\"stateRootProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"withdrawalProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"slotProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"executionPayloadProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestampProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"historicalSummaryBlockRootProof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"blockRootIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"historicalSummaryIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"withdrawalIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blockRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"slotRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"timestampRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"executionPayloadRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconChainProofs.WithdrawalProof[]\",\"name\":\"withdrawalProofs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"validatorFields\",\"type\":\"bytes32[][]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"withdrawalFields\",\"type\":\"bytes32[][]\"}],\"name\":\"verifyAndProcessWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"oracleTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint40[]\",\"name\":\"validatorIndices\",\"type\":\"uint40[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"beaconStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"name\":\"stateRootProof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"validatorBalanceProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"validatorFieldsProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"balanceRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconChainProofs.BalanceUpdateProof[]\",\"name\":\"balanceUpdateProofs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"validatorFields\",\"type\":\"bytes32[][]\"}],\"name\":\"verifyBalanceUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"oracleTimestamp\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"beaconStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"name\":\"stateRootProof\",\"type\":\"tuple\"},{\"internalType\":\"uint40[]\",\"name\":\"validatorIndices\",\"type\":\"uint40[]\"},{\"internalType\":\"bytes[]\",\"name\":\"withdrawalCredentialProofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"validatorFields\",\"type\":\"bytes32[][]\"}],\"name\":\"verifyWithdrawalCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawBeforeRestaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"withdrawNonBeaconChainETHBalanceWei\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawableRestakedExecutionLayerGwei\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

func (_Contract *ContractCaller) MAXRESTAKEDBALANCEGWEIPERVALIDATOR(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_Contract *ContractSession) MAXRESTAKEDBALANCEGWEIPERVALIDATOR() (uint64, error) {
	return _Contract.Contract.MAXRESTAKEDBALANCEGWEIPERVALIDATOR(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) MAXRESTAKEDBALANCEGWEIPERVALIDATOR() (uint64, error) {
	return _Contract.Contract.MAXRESTAKEDBALANCEGWEIPERVALIDATOR(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) EigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "eigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) EigenPodManager() (common.Address, error) {
	return _Contract.Contract.EigenPodManager(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) EigenPodManager() (common.Address, error) {
	return _Contract.Contract.EigenPodManager(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) HasRestaked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "hasRestaked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Contract *ContractSession) HasRestaked() (bool, error) {
	return _Contract.Contract.HasRestaked(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) HasRestaked() (bool, error) {
	return _Contract.Contract.HasRestaked(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) MostRecentWithdrawalTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "mostRecentWithdrawalTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_Contract *ContractSession) MostRecentWithdrawalTimestamp() (uint64, error) {
	return _Contract.Contract.MostRecentWithdrawalTimestamp(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) MostRecentWithdrawalTimestamp() (uint64, error) {
	return _Contract.Contract.MostRecentWithdrawalTimestamp(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) NonBeaconChainETHBalanceWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nonBeaconChainETHBalanceWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) NonBeaconChainETHBalanceWei() (*big.Int, error) {
	return _Contract.Contract.NonBeaconChainETHBalanceWei(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) NonBeaconChainETHBalanceWei() (*big.Int, error) {
	return _Contract.Contract.NonBeaconChainETHBalanceWei(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) PodOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "podOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) PodOwner() (common.Address, error) {
	return _Contract.Contract.PodOwner(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) PodOwner() (common.Address, error) {
	return _Contract.Contract.PodOwner(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) ProvenWithdrawal(opts *bind.CallOpts, validatorPubkeyHash [32]byte, slot uint64) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "provenWithdrawal", validatorPubkeyHash, slot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Contract *ContractSession) ProvenWithdrawal(validatorPubkeyHash [32]byte, slot uint64) (bool, error) {
	return _Contract.Contract.ProvenWithdrawal(&_Contract.CallOpts, validatorPubkeyHash, slot)
}

func (_Contract *ContractCallerSession) ProvenWithdrawal(validatorPubkeyHash [32]byte, slot uint64) (bool, error) {
	return _Contract.Contract.ProvenWithdrawal(&_Contract.CallOpts, validatorPubkeyHash, slot)
}

func (_Contract *ContractCaller) ValidatorPubkeyHashToInfo(opts *bind.CallOpts, validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "validatorPubkeyHashToInfo", validatorPubkeyHash)

	if err != nil {
		return *new(IEigenPodValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IEigenPodValidatorInfo)).(*IEigenPodValidatorInfo)

	return out0, err

}

func (_Contract *ContractSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	return _Contract.Contract.ValidatorPubkeyHashToInfo(&_Contract.CallOpts, validatorPubkeyHash)
}

func (_Contract *ContractCallerSession) ValidatorPubkeyHashToInfo(validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error) {
	return _Contract.Contract.ValidatorPubkeyHashToInfo(&_Contract.CallOpts, validatorPubkeyHash)
}

func (_Contract *ContractCaller) ValidatorStatus(opts *bind.CallOpts, pubkeyHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "validatorStatus", pubkeyHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_Contract *ContractSession) ValidatorStatus(pubkeyHash [32]byte) (uint8, error) {
	return _Contract.Contract.ValidatorStatus(&_Contract.CallOpts, pubkeyHash)
}

func (_Contract *ContractCallerSession) ValidatorStatus(pubkeyHash [32]byte) (uint8, error) {
	return _Contract.Contract.ValidatorStatus(&_Contract.CallOpts, pubkeyHash)
}

func (_Contract *ContractCaller) WithdrawableRestakedExecutionLayerGwei(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "withdrawableRestakedExecutionLayerGwei")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_Contract *ContractSession) WithdrawableRestakedExecutionLayerGwei() (uint64, error) {
	return _Contract.Contract.WithdrawableRestakedExecutionLayerGwei(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) WithdrawableRestakedExecutionLayerGwei() (uint64, error) {
	return _Contract.Contract.WithdrawableRestakedExecutionLayerGwei(&_Contract.CallOpts)
}

func (_Contract *ContractTransactor) ActivateRestaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "activateRestaking")
}

func (_Contract *ContractSession) ActivateRestaking() (*types.Transaction, error) {
	return _Contract.Contract.ActivateRestaking(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactorSession) ActivateRestaking() (*types.Transaction, error) {
	return _Contract.Contract.ActivateRestaking(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, _podOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", _podOwner)
}

func (_Contract *ContractSession) Initialize(_podOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _podOwner)
}

func (_Contract *ContractTransactorSession) Initialize(_podOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, _podOwner)
}

func (_Contract *ContractTransactor) RecoverTokens(opts *bind.TransactOpts, tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "recoverTokens", tokenList, amountsToWithdraw, recipient)
}

func (_Contract *ContractSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RecoverTokens(&_Contract.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

func (_Contract *ContractTransactorSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RecoverTokens(&_Contract.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

func (_Contract *ContractTransactor) VerifyAndProcessWithdrawals(opts *bind.TransactOpts, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "verifyAndProcessWithdrawals", oracleTimestamp, stateRootProof, withdrawalProofs, validatorFieldsProofs, validatorFields, withdrawalFields)
}

func (_Contract *ContractSession) VerifyAndProcessWithdrawals(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.VerifyAndProcessWithdrawals(&_Contract.TransactOpts, oracleTimestamp, stateRootProof, withdrawalProofs, validatorFieldsProofs, validatorFields, withdrawalFields)
}

func (_Contract *ContractTransactorSession) VerifyAndProcessWithdrawals(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.VerifyAndProcessWithdrawals(&_Contract.TransactOpts, oracleTimestamp, stateRootProof, withdrawalProofs, validatorFieldsProofs, validatorFields, withdrawalFields)
}

func (_Contract *ContractTransactor) VerifyBalanceUpdates(opts *bind.TransactOpts, oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, balanceUpdateProofs []BeaconChainProofsBalanceUpdateProof, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "verifyBalanceUpdates", oracleTimestamp, validatorIndices, stateRootProof, balanceUpdateProofs, validatorFields)
}

func (_Contract *ContractSession) VerifyBalanceUpdates(oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, balanceUpdateProofs []BeaconChainProofsBalanceUpdateProof, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.VerifyBalanceUpdates(&_Contract.TransactOpts, oracleTimestamp, validatorIndices, stateRootProof, balanceUpdateProofs, validatorFields)
}

func (_Contract *ContractTransactorSession) VerifyBalanceUpdates(oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, balanceUpdateProofs []BeaconChainProofsBalanceUpdateProof, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.VerifyBalanceUpdates(&_Contract.TransactOpts, oracleTimestamp, validatorIndices, stateRootProof, balanceUpdateProofs, validatorFields)
}

func (_Contract *ContractTransactor) VerifyWithdrawalCredentials(opts *bind.TransactOpts, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, withdrawalCredentialProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "verifyWithdrawalCredentials", oracleTimestamp, stateRootProof, validatorIndices, withdrawalCredentialProofs, validatorFields)
}

func (_Contract *ContractSession) VerifyWithdrawalCredentials(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, withdrawalCredentialProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.VerifyWithdrawalCredentials(&_Contract.TransactOpts, oracleTimestamp, stateRootProof, validatorIndices, withdrawalCredentialProofs, validatorFields)
}

func (_Contract *ContractTransactorSession) VerifyWithdrawalCredentials(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, withdrawalCredentialProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Contract.Contract.VerifyWithdrawalCredentials(&_Contract.TransactOpts, oracleTimestamp, stateRootProof, validatorIndices, withdrawalCredentialProofs, validatorFields)
}

func (_Contract *ContractTransactor) WithdrawBeforeRestaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawBeforeRestaking")
}

func (_Contract *ContractSession) WithdrawBeforeRestaking() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawBeforeRestaking(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactorSession) WithdrawBeforeRestaking() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawBeforeRestaking(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactor) WithdrawNonBeaconChainETHBalanceWei(opts *bind.TransactOpts, recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawNonBeaconChainETHBalanceWei", recipient, amountToWithdraw)
}

func (_Contract *ContractSession) WithdrawNonBeaconChainETHBalanceWei(recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawNonBeaconChainETHBalanceWei(&_Contract.TransactOpts, recipient, amountToWithdraw)
}

func (_Contract *ContractTransactorSession) WithdrawNonBeaconChainETHBalanceWei(recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawNonBeaconChainETHBalanceWei(&_Contract.TransactOpts, recipient, amountToWithdraw)
}

type ContractEigenPodStakedIterator struct {
	Event *ContractEigenPodStaked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractEigenPodStakedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEigenPodStaked)
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
		it.Event = new(ContractEigenPodStaked)
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

func (it *ContractEigenPodStakedIterator) Error() error {
	return it.fail
}

func (it *ContractEigenPodStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractEigenPodStaked struct {
	Pubkey []byte
	Raw    types.Log
}

func (_Contract *ContractFilterer) FilterEigenPodStaked(opts *bind.FilterOpts) (*ContractEigenPodStakedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EigenPodStaked")
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodStakedIterator{contract: _Contract.contract, event: "EigenPodStaked", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchEigenPodStaked(opts *bind.WatchOpts, sink chan<- *ContractEigenPodStaked) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EigenPodStaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractEigenPodStaked)
				if err := _Contract.contract.UnpackLog(event, "EigenPodStaked", log); err != nil {
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

func (_Contract *ContractFilterer) ParseEigenPodStaked(log types.Log) (*ContractEigenPodStaked, error) {
	event := new(ContractEigenPodStaked)
	if err := _Contract.contract.UnpackLog(event, "EigenPodStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractFullWithdrawalRedeemedIterator struct {
	Event *ContractFullWithdrawalRedeemed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractFullWithdrawalRedeemedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractFullWithdrawalRedeemed)
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
		it.Event = new(ContractFullWithdrawalRedeemed)
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

func (it *ContractFullWithdrawalRedeemedIterator) Error() error {
	return it.fail
}

func (it *ContractFullWithdrawalRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractFullWithdrawalRedeemed struct {
	ValidatorIndex       *big.Int
	WithdrawalTimestamp  uint64
	Recipient            common.Address
	WithdrawalAmountGwei uint64
	Raw                  types.Log
}

func (_Contract *ContractFilterer) FilterFullWithdrawalRedeemed(opts *bind.FilterOpts, recipient []common.Address) (*ContractFullWithdrawalRedeemedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "FullWithdrawalRedeemed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractFullWithdrawalRedeemedIterator{contract: _Contract.contract, event: "FullWithdrawalRedeemed", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchFullWithdrawalRedeemed(opts *bind.WatchOpts, sink chan<- *ContractFullWithdrawalRedeemed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "FullWithdrawalRedeemed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractFullWithdrawalRedeemed)
				if err := _Contract.contract.UnpackLog(event, "FullWithdrawalRedeemed", log); err != nil {
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

func (_Contract *ContractFilterer) ParseFullWithdrawalRedeemed(log types.Log) (*ContractFullWithdrawalRedeemed, error) {
	event := new(ContractFullWithdrawalRedeemed)
	if err := _Contract.contract.UnpackLog(event, "FullWithdrawalRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractNonBeaconChainETHReceivedIterator struct {
	Event *ContractNonBeaconChainETHReceived

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractNonBeaconChainETHReceivedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNonBeaconChainETHReceived)
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
		it.Event = new(ContractNonBeaconChainETHReceived)
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

func (it *ContractNonBeaconChainETHReceivedIterator) Error() error {
	return it.fail
}

func (it *ContractNonBeaconChainETHReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractNonBeaconChainETHReceived struct {
	AmountReceived *big.Int
	Raw            types.Log
}

func (_Contract *ContractFilterer) FilterNonBeaconChainETHReceived(opts *bind.FilterOpts) (*ContractNonBeaconChainETHReceivedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NonBeaconChainETHReceived")
	if err != nil {
		return nil, err
	}
	return &ContractNonBeaconChainETHReceivedIterator{contract: _Contract.contract, event: "NonBeaconChainETHReceived", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchNonBeaconChainETHReceived(opts *bind.WatchOpts, sink chan<- *ContractNonBeaconChainETHReceived) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NonBeaconChainETHReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractNonBeaconChainETHReceived)
				if err := _Contract.contract.UnpackLog(event, "NonBeaconChainETHReceived", log); err != nil {
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

func (_Contract *ContractFilterer) ParseNonBeaconChainETHReceived(log types.Log) (*ContractNonBeaconChainETHReceived, error) {
	event := new(ContractNonBeaconChainETHReceived)
	if err := _Contract.contract.UnpackLog(event, "NonBeaconChainETHReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractNonBeaconChainETHWithdrawnIterator struct {
	Event *ContractNonBeaconChainETHWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractNonBeaconChainETHWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNonBeaconChainETHWithdrawn)
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
		it.Event = new(ContractNonBeaconChainETHWithdrawn)
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

func (it *ContractNonBeaconChainETHWithdrawnIterator) Error() error {
	return it.fail
}

func (it *ContractNonBeaconChainETHWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractNonBeaconChainETHWithdrawn struct {
	Recipient       common.Address
	AmountWithdrawn *big.Int
	Raw             types.Log
}

func (_Contract *ContractFilterer) FilterNonBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ContractNonBeaconChainETHWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NonBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractNonBeaconChainETHWithdrawnIterator{contract: _Contract.contract, event: "NonBeaconChainETHWithdrawn", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchNonBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractNonBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NonBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractNonBeaconChainETHWithdrawn)
				if err := _Contract.contract.UnpackLog(event, "NonBeaconChainETHWithdrawn", log); err != nil {
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

func (_Contract *ContractFilterer) ParseNonBeaconChainETHWithdrawn(log types.Log) (*ContractNonBeaconChainETHWithdrawn, error) {
	event := new(ContractNonBeaconChainETHWithdrawn)
	if err := _Contract.contract.UnpackLog(event, "NonBeaconChainETHWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractPartialWithdrawalRedeemedIterator struct {
	Event *ContractPartialWithdrawalRedeemed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractPartialWithdrawalRedeemedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPartialWithdrawalRedeemed)
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
		it.Event = new(ContractPartialWithdrawalRedeemed)
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

func (it *ContractPartialWithdrawalRedeemedIterator) Error() error {
	return it.fail
}

func (it *ContractPartialWithdrawalRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractPartialWithdrawalRedeemed struct {
	ValidatorIndex              *big.Int
	WithdrawalTimestamp         uint64
	Recipient                   common.Address
	PartialWithdrawalAmountGwei uint64
	Raw                         types.Log
}

func (_Contract *ContractFilterer) FilterPartialWithdrawalRedeemed(opts *bind.FilterOpts, recipient []common.Address) (*ContractPartialWithdrawalRedeemedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PartialWithdrawalRedeemed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractPartialWithdrawalRedeemedIterator{contract: _Contract.contract, event: "PartialWithdrawalRedeemed", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchPartialWithdrawalRedeemed(opts *bind.WatchOpts, sink chan<- *ContractPartialWithdrawalRedeemed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PartialWithdrawalRedeemed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractPartialWithdrawalRedeemed)
				if err := _Contract.contract.UnpackLog(event, "PartialWithdrawalRedeemed", log); err != nil {
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

func (_Contract *ContractFilterer) ParsePartialWithdrawalRedeemed(log types.Log) (*ContractPartialWithdrawalRedeemed, error) {
	event := new(ContractPartialWithdrawalRedeemed)
	if err := _Contract.contract.UnpackLog(event, "PartialWithdrawalRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractRestakedBeaconChainETHWithdrawnIterator struct {
	Event *ContractRestakedBeaconChainETHWithdrawn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractRestakedBeaconChainETHWithdrawnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRestakedBeaconChainETHWithdrawn)
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
		it.Event = new(ContractRestakedBeaconChainETHWithdrawn)
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

func (it *ContractRestakedBeaconChainETHWithdrawnIterator) Error() error {
	return it.fail
}

func (it *ContractRestakedBeaconChainETHWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractRestakedBeaconChainETHWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterRestakedBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ContractRestakedBeaconChainETHWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RestakedBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ContractRestakedBeaconChainETHWithdrawnIterator{contract: _Contract.contract, event: "RestakedBeaconChainETHWithdrawn", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchRestakedBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractRestakedBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RestakedBeaconChainETHWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractRestakedBeaconChainETHWithdrawn)
				if err := _Contract.contract.UnpackLog(event, "RestakedBeaconChainETHWithdrawn", log); err != nil {
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

func (_Contract *ContractFilterer) ParseRestakedBeaconChainETHWithdrawn(log types.Log) (*ContractRestakedBeaconChainETHWithdrawn, error) {
	event := new(ContractRestakedBeaconChainETHWithdrawn)
	if err := _Contract.contract.UnpackLog(event, "RestakedBeaconChainETHWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractRestakingActivatedIterator struct {
	Event *ContractRestakingActivated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractRestakingActivatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRestakingActivated)
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
		it.Event = new(ContractRestakingActivated)
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

func (it *ContractRestakingActivatedIterator) Error() error {
	return it.fail
}

func (it *ContractRestakingActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractRestakingActivated struct {
	PodOwner common.Address
	Raw      types.Log
}

func (_Contract *ContractFilterer) FilterRestakingActivated(opts *bind.FilterOpts, podOwner []common.Address) (*ContractRestakingActivatedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RestakingActivated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractRestakingActivatedIterator{contract: _Contract.contract, event: "RestakingActivated", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchRestakingActivated(opts *bind.WatchOpts, sink chan<- *ContractRestakingActivated, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RestakingActivated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractRestakingActivated)
				if err := _Contract.contract.UnpackLog(event, "RestakingActivated", log); err != nil {
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

func (_Contract *ContractFilterer) ParseRestakingActivated(log types.Log) (*ContractRestakingActivated, error) {
	event := new(ContractRestakingActivated)
	if err := _Contract.contract.UnpackLog(event, "RestakingActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractValidatorBalanceUpdatedIterator struct {
	Event *ContractValidatorBalanceUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractValidatorBalanceUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorBalanceUpdated)
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
		it.Event = new(ContractValidatorBalanceUpdated)
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

func (it *ContractValidatorBalanceUpdatedIterator) Error() error {
	return it.fail
}

func (it *ContractValidatorBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractValidatorBalanceUpdated struct {
	ValidatorIndex          *big.Int
	BalanceTimestamp        uint64
	NewValidatorBalanceGwei uint64
	Raw                     types.Log
}

func (_Contract *ContractFilterer) FilterValidatorBalanceUpdated(opts *bind.FilterOpts) (*ContractValidatorBalanceUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractValidatorBalanceUpdatedIterator{contract: _Contract.contract, event: "ValidatorBalanceUpdated", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchValidatorBalanceUpdated(opts *bind.WatchOpts, sink chan<- *ContractValidatorBalanceUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractValidatorBalanceUpdated)
				if err := _Contract.contract.UnpackLog(event, "ValidatorBalanceUpdated", log); err != nil {
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

func (_Contract *ContractFilterer) ParseValidatorBalanceUpdated(log types.Log) (*ContractValidatorBalanceUpdated, error) {
	event := new(ContractValidatorBalanceUpdated)
	if err := _Contract.contract.UnpackLog(event, "ValidatorBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractValidatorRestakedIterator struct {
	Event *ContractValidatorRestaked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractValidatorRestakedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorRestaked)
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
		it.Event = new(ContractValidatorRestaked)
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

func (it *ContractValidatorRestakedIterator) Error() error {
	return it.fail
}

func (it *ContractValidatorRestakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractValidatorRestaked struct {
	ValidatorIndex *big.Int
	Raw            types.Log
}

func (_Contract *ContractFilterer) FilterValidatorRestaked(opts *bind.FilterOpts) (*ContractValidatorRestakedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorRestaked")
	if err != nil {
		return nil, err
	}
	return &ContractValidatorRestakedIterator{contract: _Contract.contract, event: "ValidatorRestaked", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchValidatorRestaked(opts *bind.WatchOpts, sink chan<- *ContractValidatorRestaked) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorRestaked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractValidatorRestaked)
				if err := _Contract.contract.UnpackLog(event, "ValidatorRestaked", log); err != nil {
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

func (_Contract *ContractFilterer) ParseValidatorRestaked(log types.Log) (*ContractValidatorRestaked, error) {
	event := new(ContractValidatorRestaked)
	if err := _Contract.contract.UnpackLog(event, "ValidatorRestaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Contract *Contract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Contract.abi.Events["EigenPodStaked"].ID:
		return _Contract.ParseEigenPodStaked(log)
	case _Contract.abi.Events["FullWithdrawalRedeemed"].ID:
		return _Contract.ParseFullWithdrawalRedeemed(log)
	case _Contract.abi.Events["NonBeaconChainETHReceived"].ID:
		return _Contract.ParseNonBeaconChainETHReceived(log)
	case _Contract.abi.Events["NonBeaconChainETHWithdrawn"].ID:
		return _Contract.ParseNonBeaconChainETHWithdrawn(log)
	case _Contract.abi.Events["PartialWithdrawalRedeemed"].ID:
		return _Contract.ParsePartialWithdrawalRedeemed(log)
	case _Contract.abi.Events["RestakedBeaconChainETHWithdrawn"].ID:
		return _Contract.ParseRestakedBeaconChainETHWithdrawn(log)
	case _Contract.abi.Events["RestakingActivated"].ID:
		return _Contract.ParseRestakingActivated(log)
	case _Contract.abi.Events["ValidatorBalanceUpdated"].ID:
		return _Contract.ParseValidatorBalanceUpdated(log)
	case _Contract.abi.Events["ValidatorRestaked"].ID:
		return _Contract.ParseValidatorRestaked(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ContractEigenPodStaked) Topic() common.Hash {
	return common.HexToHash("0x606865b7934a25d4aed43f6cdb426403353fa4b3009c4d228407474581b01e23")
}

func (ContractFullWithdrawalRedeemed) Topic() common.Hash {
	return common.HexToHash("0xb76a93bb649ece524688f1a01d184e0bbebcda58eae80c28a898bec3fb5a0963")
}

func (ContractNonBeaconChainETHReceived) Topic() common.Hash {
	return common.HexToHash("0x6fdd3dbdb173299608c0aa9f368735857c8842b581f8389238bf05bd04b3bf49")
}

func (ContractNonBeaconChainETHWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x30420aacd028abb3c1fd03aba253ae725d6ddd52d16c9ac4cb5742cd43f53096")
}

func (ContractPartialWithdrawalRedeemed) Topic() common.Hash {
	return common.HexToHash("0x8a7335714231dbd551aaba6314f4a97a14c201e53a3e25e1140325cdf67d7a4e")
}

func (ContractRestakedBeaconChainETHWithdrawn) Topic() common.Hash {
	return common.HexToHash("0x8947fd2ce07ef9cc302c4e8f0461015615d91ce851564839e91cc804c2f49d8e")
}

func (ContractRestakingActivated) Topic() common.Hash {
	return common.HexToHash("0xca8dfc8c5e0a67a74501c072a3325f685259bebbae7cfd230ab85198a78b70cd")
}

func (ContractValidatorBalanceUpdated) Topic() common.Hash {
	return common.HexToHash("0x0e5fac175b83177cc047381e030d8fb3b42b37bd1c025e22c280facad62c32df")
}

func (ContractValidatorRestaked) Topic() common.Hash {
	return common.HexToHash("0x2d0800bbc377ea54a08c5db6a87aafff5e3e9c8fead0eda110e40e0c10441449")
}

func (_Contract *Contract) Address() common.Address {
	return _Contract.address
}

type ContractInterface interface {
	MAXRESTAKEDBALANCEGWEIPERVALIDATOR(opts *bind.CallOpts) (uint64, error)

	EigenPodManager(opts *bind.CallOpts) (common.Address, error)

	HasRestaked(opts *bind.CallOpts) (bool, error)

	MostRecentWithdrawalTimestamp(opts *bind.CallOpts) (uint64, error)

	NonBeaconChainETHBalanceWei(opts *bind.CallOpts) (*big.Int, error)

	PodOwner(opts *bind.CallOpts) (common.Address, error)

	ProvenWithdrawal(opts *bind.CallOpts, validatorPubkeyHash [32]byte, slot uint64) (bool, error)

	ValidatorPubkeyHashToInfo(opts *bind.CallOpts, validatorPubkeyHash [32]byte) (IEigenPodValidatorInfo, error)

	ValidatorStatus(opts *bind.CallOpts, pubkeyHash [32]byte) (uint8, error)

	WithdrawableRestakedExecutionLayerGwei(opts *bind.CallOpts) (uint64, error)

	ActivateRestaking(opts *bind.TransactOpts) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, _podOwner common.Address) (*types.Transaction, error)

	RecoverTokens(opts *bind.TransactOpts, tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error)

	VerifyAndProcessWithdrawals(opts *bind.TransactOpts, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error)

	VerifyBalanceUpdates(opts *bind.TransactOpts, oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, balanceUpdateProofs []BeaconChainProofsBalanceUpdateProof, validatorFields [][][32]byte) (*types.Transaction, error)

	VerifyWithdrawalCredentials(opts *bind.TransactOpts, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, withdrawalCredentialProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error)

	WithdrawBeforeRestaking(opts *bind.TransactOpts) (*types.Transaction, error)

	WithdrawNonBeaconChainETHBalanceWei(opts *bind.TransactOpts, recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error)

	FilterEigenPodStaked(opts *bind.FilterOpts) (*ContractEigenPodStakedIterator, error)

	WatchEigenPodStaked(opts *bind.WatchOpts, sink chan<- *ContractEigenPodStaked) (event.Subscription, error)

	ParseEigenPodStaked(log types.Log) (*ContractEigenPodStaked, error)

	FilterFullWithdrawalRedeemed(opts *bind.FilterOpts, recipient []common.Address) (*ContractFullWithdrawalRedeemedIterator, error)

	WatchFullWithdrawalRedeemed(opts *bind.WatchOpts, sink chan<- *ContractFullWithdrawalRedeemed, recipient []common.Address) (event.Subscription, error)

	ParseFullWithdrawalRedeemed(log types.Log) (*ContractFullWithdrawalRedeemed, error)

	FilterNonBeaconChainETHReceived(opts *bind.FilterOpts) (*ContractNonBeaconChainETHReceivedIterator, error)

	WatchNonBeaconChainETHReceived(opts *bind.WatchOpts, sink chan<- *ContractNonBeaconChainETHReceived) (event.Subscription, error)

	ParseNonBeaconChainETHReceived(log types.Log) (*ContractNonBeaconChainETHReceived, error)

	FilterNonBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ContractNonBeaconChainETHWithdrawnIterator, error)

	WatchNonBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractNonBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParseNonBeaconChainETHWithdrawn(log types.Log) (*ContractNonBeaconChainETHWithdrawn, error)

	FilterPartialWithdrawalRedeemed(opts *bind.FilterOpts, recipient []common.Address) (*ContractPartialWithdrawalRedeemedIterator, error)

	WatchPartialWithdrawalRedeemed(opts *bind.WatchOpts, sink chan<- *ContractPartialWithdrawalRedeemed, recipient []common.Address) (event.Subscription, error)

	ParsePartialWithdrawalRedeemed(log types.Log) (*ContractPartialWithdrawalRedeemed, error)

	FilterRestakedBeaconChainETHWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ContractRestakedBeaconChainETHWithdrawnIterator, error)

	WatchRestakedBeaconChainETHWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractRestakedBeaconChainETHWithdrawn, recipient []common.Address) (event.Subscription, error)

	ParseRestakedBeaconChainETHWithdrawn(log types.Log) (*ContractRestakedBeaconChainETHWithdrawn, error)

	FilterRestakingActivated(opts *bind.FilterOpts, podOwner []common.Address) (*ContractRestakingActivatedIterator, error)

	WatchRestakingActivated(opts *bind.WatchOpts, sink chan<- *ContractRestakingActivated, podOwner []common.Address) (event.Subscription, error)

	ParseRestakingActivated(log types.Log) (*ContractRestakingActivated, error)

	FilterValidatorBalanceUpdated(opts *bind.FilterOpts) (*ContractValidatorBalanceUpdatedIterator, error)

	WatchValidatorBalanceUpdated(opts *bind.WatchOpts, sink chan<- *ContractValidatorBalanceUpdated) (event.Subscription, error)

	ParseValidatorBalanceUpdated(log types.Log) (*ContractValidatorBalanceUpdated, error)

	FilterValidatorRestaked(opts *bind.FilterOpts) (*ContractValidatorRestakedIterator, error)

	WatchValidatorRestaked(opts *bind.WatchOpts, sink chan<- *ContractValidatorRestaked) (event.Subscription, error)

	ParseValidatorRestaked(log types.Log) (*ContractValidatorRestaked, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
