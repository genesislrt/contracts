// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testeigenpodmanager

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BeaconChainETHDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"nonce\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegatedAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"}],\"name\":\"BeaconChainETHWithdrawalCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOracleAddress\",\"type\":\"address\"}],\"name\":\"BeaconOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MaxPodsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"eigenPod\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"}],\"name\":\"PodDeployed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"beaconChainETHStrategy\",\"outputs\":[{\"internalType\":\"contractIStrategy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beaconChainOracle\",\"outputs\":[{\"internalType\":\"contractIBeaconChainOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createPod\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenPodBeacon\",\"outputs\":[{\"internalType\":\"contractIBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethPOS\",\"outputs\":[{\"internalType\":\"contractIETHPOSDeposit\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"getBlockRootAtTimestamp\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"}],\"name\":\"getPod\",\"outputs\":[{\"internalType\":\"contractIEigenPod\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"}],\"name\":\"hasPod\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numPods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"}],\"name\":\"ownerToPod\",\"outputs\":[{\"internalType\":\"contractIEigenPod\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"podOwner\",\"type\":\"address\"}],\"name\":\"podOwnerShares\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasher\",\"outputs\":[{\"internalType\":\"contractISlasher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"depositDataRoot\",\"type\":\"bytes32\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyManager\",\"outputs\":[{\"internalType\":\"contractIStrategyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506103048061001d5f395ff3fe6080604052600436106100e4575f3560e01c80639ba0627511610087578063c052bd6111610057578063c052bd61146100e8578063c0ccbf1014610171578063d1c64cc914610183578063f6848d241461019d575f80fd5b80639ba0627514610157578063a38406a314610157578063a6a509be14610171578063b1344271146100e8575f80fd5b806374cdd798116100c257806374cdd798146100e857806384d81062146100e85780639104c319146100e85780639b4e463414610140575f80fd5b8063292b7b2b146100e857806339b70e38146100e857806360f4062b14610113575b5f80fd5b3480156100f3575f80fd5b505f5b6040516001600160a01b0390911681526020015b60405180910390f35b34801561011e575f80fd5b5061013261012d3660046101c7565b505f90565b60405190815260200161010a565b61015561014e366004610239565b5050505050565b005b348015610162575f80fd5b506100f661012d3660046101c7565b34801561017c575f80fd5b505f610132565b34801561018e575f80fd5b5061013261012d3660046102a7565b3480156101a8575f80fd5b506101b761012d3660046101c7565b604051901515815260200161010a565b5f602082840312156101d7575f80fd5b81356001600160a01b03811681146101ed575f80fd5b9392505050565b5f8083601f840112610204575f80fd5b50813567ffffffffffffffff81111561021b575f80fd5b602083019150836020828501011115610232575f80fd5b9250929050565b5f805f805f6060868803121561024d575f80fd5b853567ffffffffffffffff80821115610264575f80fd5b61027089838a016101f4565b90975095506020880135915080821115610288575f80fd5b50610295888289016101f4565b96999598509660400135949350505050565b5f602082840312156102b7575f80fd5b813567ffffffffffffffff811681146101ed575f80fdfea2646970667358221220964305bf61144c20f37e27c87475ab5b2d67f1cd8d2daa68c2d9185194c0104564736f6c63430008140033",
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

func (_Contract *ContractCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) BeaconChainETHStrategy() (common.Address, error) {
	return _Contract.Contract.BeaconChainETHStrategy(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _Contract.Contract.BeaconChainETHStrategy(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) BeaconChainOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "beaconChainOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) BeaconChainOracle() (common.Address, error) {
	return _Contract.Contract.BeaconChainOracle(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) BeaconChainOracle() (common.Address, error) {
	return _Contract.Contract.BeaconChainOracle(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) EigenPodBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "eigenPodBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) EigenPodBeacon() (common.Address, error) {
	return _Contract.Contract.EigenPodBeacon(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) EigenPodBeacon() (common.Address, error) {
	return _Contract.Contract.EigenPodBeacon(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) EthPOS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ethPOS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) EthPOS() (common.Address, error) {
	return _Contract.Contract.EthPOS(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) EthPOS() (common.Address, error) {
	return _Contract.Contract.EthPOS(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetBlockRootAtTimestamp(opts *bind.CallOpts, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBlockRootAtTimestamp", timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_Contract *ContractSession) GetBlockRootAtTimestamp(timestamp uint64) ([32]byte, error) {
	return _Contract.Contract.GetBlockRootAtTimestamp(&_Contract.CallOpts, timestamp)
}

func (_Contract *ContractCallerSession) GetBlockRootAtTimestamp(timestamp uint64) ([32]byte, error) {
	return _Contract.Contract.GetBlockRootAtTimestamp(&_Contract.CallOpts, timestamp)
}

func (_Contract *ContractCaller) GetPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPod", podOwner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetPod(podOwner common.Address) (common.Address, error) {
	return _Contract.Contract.GetPod(&_Contract.CallOpts, podOwner)
}

func (_Contract *ContractCallerSession) GetPod(podOwner common.Address) (common.Address, error) {
	return _Contract.Contract.GetPod(&_Contract.CallOpts, podOwner)
}

func (_Contract *ContractCaller) HasPod(opts *bind.CallOpts, podOwner common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "hasPod", podOwner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Contract *ContractSession) HasPod(podOwner common.Address) (bool, error) {
	return _Contract.Contract.HasPod(&_Contract.CallOpts, podOwner)
}

func (_Contract *ContractCallerSession) HasPod(podOwner common.Address) (bool, error) {
	return _Contract.Contract.HasPod(&_Contract.CallOpts, podOwner)
}

func (_Contract *ContractCaller) MaxPods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxPods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) MaxPods() (*big.Int, error) {
	return _Contract.Contract.MaxPods(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) MaxPods() (*big.Int, error) {
	return _Contract.Contract.MaxPods(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) NumPods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "numPods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) NumPods() (*big.Int, error) {
	return _Contract.Contract.NumPods(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) NumPods() (*big.Int, error) {
	return _Contract.Contract.NumPods(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) OwnerToPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ownerToPod", podOwner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) OwnerToPod(podOwner common.Address) (common.Address, error) {
	return _Contract.Contract.OwnerToPod(&_Contract.CallOpts, podOwner)
}

func (_Contract *ContractCallerSession) OwnerToPod(podOwner common.Address) (common.Address, error) {
	return _Contract.Contract.OwnerToPod(&_Contract.CallOpts, podOwner)
}

func (_Contract *ContractCaller) PodOwnerShares(opts *bind.CallOpts, podOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "podOwnerShares", podOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) PodOwnerShares(podOwner common.Address) (*big.Int, error) {
	return _Contract.Contract.PodOwnerShares(&_Contract.CallOpts, podOwner)
}

func (_Contract *ContractCallerSession) PodOwnerShares(podOwner common.Address) (*big.Int, error) {
	return _Contract.Contract.PodOwnerShares(&_Contract.CallOpts, podOwner)
}

func (_Contract *ContractCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) Slasher() (common.Address, error) {
	return _Contract.Contract.Slasher(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) Slasher() (common.Address, error) {
	return _Contract.Contract.Slasher(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) StrategyManager() (common.Address, error) {
	return _Contract.Contract.StrategyManager(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) StrategyManager() (common.Address, error) {
	return _Contract.Contract.StrategyManager(&_Contract.CallOpts)
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

type ContractBeaconChainETHDepositedIterator struct {
	Event *ContractBeaconChainETHDeposited

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractBeaconChainETHDepositedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBeaconChainETHDeposited)
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
		it.Event = new(ContractBeaconChainETHDeposited)
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

func (it *ContractBeaconChainETHDepositedIterator) Error() error {
	return it.fail
}

func (it *ContractBeaconChainETHDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractBeaconChainETHDeposited struct {
	PodOwner common.Address
	Amount   *big.Int
	Raw      types.Log
}

func (_Contract *ContractFilterer) FilterBeaconChainETHDeposited(opts *bind.FilterOpts, podOwner []common.Address) (*ContractBeaconChainETHDepositedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "BeaconChainETHDeposited", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractBeaconChainETHDepositedIterator{contract: _Contract.contract, event: "BeaconChainETHDeposited", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchBeaconChainETHDeposited(opts *bind.WatchOpts, sink chan<- *ContractBeaconChainETHDeposited, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "BeaconChainETHDeposited", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractBeaconChainETHDeposited)
				if err := _Contract.contract.UnpackLog(event, "BeaconChainETHDeposited", log); err != nil {
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

func (_Contract *ContractFilterer) ParseBeaconChainETHDeposited(log types.Log) (*ContractBeaconChainETHDeposited, error) {
	event := new(ContractBeaconChainETHDeposited)
	if err := _Contract.contract.UnpackLog(event, "BeaconChainETHDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractBeaconChainETHWithdrawalCompletedIterator struct {
	Event *ContractBeaconChainETHWithdrawalCompleted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractBeaconChainETHWithdrawalCompletedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBeaconChainETHWithdrawalCompleted)
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
		it.Event = new(ContractBeaconChainETHWithdrawalCompleted)
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

func (it *ContractBeaconChainETHWithdrawalCompletedIterator) Error() error {
	return it.fail
}

func (it *ContractBeaconChainETHWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractBeaconChainETHWithdrawalCompleted struct {
	PodOwner         common.Address
	Shares           *big.Int
	Nonce            *big.Int
	DelegatedAddress common.Address
	Withdrawer       common.Address
	WithdrawalRoot   [32]byte
	Raw              types.Log
}

func (_Contract *ContractFilterer) FilterBeaconChainETHWithdrawalCompleted(opts *bind.FilterOpts, podOwner []common.Address) (*ContractBeaconChainETHWithdrawalCompletedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "BeaconChainETHWithdrawalCompleted", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractBeaconChainETHWithdrawalCompletedIterator{contract: _Contract.contract, event: "BeaconChainETHWithdrawalCompleted", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchBeaconChainETHWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ContractBeaconChainETHWithdrawalCompleted, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "BeaconChainETHWithdrawalCompleted", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractBeaconChainETHWithdrawalCompleted)
				if err := _Contract.contract.UnpackLog(event, "BeaconChainETHWithdrawalCompleted", log); err != nil {
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

func (_Contract *ContractFilterer) ParseBeaconChainETHWithdrawalCompleted(log types.Log) (*ContractBeaconChainETHWithdrawalCompleted, error) {
	event := new(ContractBeaconChainETHWithdrawalCompleted)
	if err := _Contract.contract.UnpackLog(event, "BeaconChainETHWithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractBeaconOracleUpdatedIterator struct {
	Event *ContractBeaconOracleUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractBeaconOracleUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBeaconOracleUpdated)
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
		it.Event = new(ContractBeaconOracleUpdated)
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

func (it *ContractBeaconOracleUpdatedIterator) Error() error {
	return it.fail
}

func (it *ContractBeaconOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractBeaconOracleUpdated struct {
	NewOracleAddress common.Address
	Raw              types.Log
}

func (_Contract *ContractFilterer) FilterBeaconOracleUpdated(opts *bind.FilterOpts, newOracleAddress []common.Address) (*ContractBeaconOracleUpdatedIterator, error) {

	var newOracleAddressRule []interface{}
	for _, newOracleAddressItem := range newOracleAddress {
		newOracleAddressRule = append(newOracleAddressRule, newOracleAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "BeaconOracleUpdated", newOracleAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractBeaconOracleUpdatedIterator{contract: _Contract.contract, event: "BeaconOracleUpdated", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchBeaconOracleUpdated(opts *bind.WatchOpts, sink chan<- *ContractBeaconOracleUpdated, newOracleAddress []common.Address) (event.Subscription, error) {

	var newOracleAddressRule []interface{}
	for _, newOracleAddressItem := range newOracleAddress {
		newOracleAddressRule = append(newOracleAddressRule, newOracleAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "BeaconOracleUpdated", newOracleAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractBeaconOracleUpdated)
				if err := _Contract.contract.UnpackLog(event, "BeaconOracleUpdated", log); err != nil {
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

func (_Contract *ContractFilterer) ParseBeaconOracleUpdated(log types.Log) (*ContractBeaconOracleUpdated, error) {
	event := new(ContractBeaconOracleUpdated)
	if err := _Contract.contract.UnpackLog(event, "BeaconOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractMaxPodsUpdatedIterator struct {
	Event *ContractMaxPodsUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractMaxPodsUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMaxPodsUpdated)
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
		it.Event = new(ContractMaxPodsUpdated)
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

func (it *ContractMaxPodsUpdatedIterator) Error() error {
	return it.fail
}

func (it *ContractMaxPodsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractMaxPodsUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log
}

func (_Contract *ContractFilterer) FilterMaxPodsUpdated(opts *bind.FilterOpts) (*ContractMaxPodsUpdatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "MaxPodsUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractMaxPodsUpdatedIterator{contract: _Contract.contract, event: "MaxPodsUpdated", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchMaxPodsUpdated(opts *bind.WatchOpts, sink chan<- *ContractMaxPodsUpdated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "MaxPodsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractMaxPodsUpdated)
				if err := _Contract.contract.UnpackLog(event, "MaxPodsUpdated", log); err != nil {
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

func (_Contract *ContractFilterer) ParseMaxPodsUpdated(log types.Log) (*ContractMaxPodsUpdated, error) {
	event := new(ContractMaxPodsUpdated)
	if err := _Contract.contract.UnpackLog(event, "MaxPodsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractPodDeployedIterator struct {
	Event *ContractPodDeployed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractPodDeployedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPodDeployed)
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
		it.Event = new(ContractPodDeployed)
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

func (it *ContractPodDeployedIterator) Error() error {
	return it.fail
}

func (it *ContractPodDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractPodDeployed struct {
	EigenPod common.Address
	PodOwner common.Address
	Raw      types.Log
}

func (_Contract *ContractFilterer) FilterPodDeployed(opts *bind.FilterOpts, eigenPod []common.Address, podOwner []common.Address) (*ContractPodDeployedIterator, error) {

	var eigenPodRule []interface{}
	for _, eigenPodItem := range eigenPod {
		eigenPodRule = append(eigenPodRule, eigenPodItem)
	}
	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PodDeployed", eigenPodRule, podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractPodDeployedIterator{contract: _Contract.contract, event: "PodDeployed", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchPodDeployed(opts *bind.WatchOpts, sink chan<- *ContractPodDeployed, eigenPod []common.Address, podOwner []common.Address) (event.Subscription, error) {

	var eigenPodRule []interface{}
	for _, eigenPodItem := range eigenPod {
		eigenPodRule = append(eigenPodRule, eigenPodItem)
	}
	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PodDeployed", eigenPodRule, podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractPodDeployed)
				if err := _Contract.contract.UnpackLog(event, "PodDeployed", log); err != nil {
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

func (_Contract *ContractFilterer) ParsePodDeployed(log types.Log) (*ContractPodDeployed, error) {
	event := new(ContractPodDeployed)
	if err := _Contract.contract.UnpackLog(event, "PodDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Contract *Contract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Contract.abi.Events["BeaconChainETHDeposited"].ID:
		return _Contract.ParseBeaconChainETHDeposited(log)
	case _Contract.abi.Events["BeaconChainETHWithdrawalCompleted"].ID:
		return _Contract.ParseBeaconChainETHWithdrawalCompleted(log)
	case _Contract.abi.Events["BeaconOracleUpdated"].ID:
		return _Contract.ParseBeaconOracleUpdated(log)
	case _Contract.abi.Events["MaxPodsUpdated"].ID:
		return _Contract.ParseMaxPodsUpdated(log)
	case _Contract.abi.Events["PodDeployed"].ID:
		return _Contract.ParsePodDeployed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ContractBeaconChainETHDeposited) Topic() common.Hash {
	return common.HexToHash("0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d")
}

func (ContractBeaconChainETHWithdrawalCompleted) Topic() common.Hash {
	return common.HexToHash("0xa6bab1d55a361fcea2eee2bc9491e4f01e6cf333df03c9c4f2c144466429f7d6")
}

func (ContractBeaconOracleUpdated) Topic() common.Hash {
	return common.HexToHash("0x08f0470754946ccfbb446ff7fd2d6ae6af1bbdae19f85794c0cc5ed5e8ceb4f6")
}

func (ContractMaxPodsUpdated) Topic() common.Hash {
	return common.HexToHash("0x4e65c41a3597bda732ca64980235cf51494171d5853998763fb05db45afaacb3")
}

func (ContractPodDeployed) Topic() common.Hash {
	return common.HexToHash("0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a")
}

func (_Contract *Contract) Address() common.Address {
	return _Contract.address
}

type ContractInterface interface {
	BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error)

	BeaconChainOracle(opts *bind.CallOpts) (common.Address, error)

	EigenPodBeacon(opts *bind.CallOpts) (common.Address, error)

	EthPOS(opts *bind.CallOpts) (common.Address, error)

	GetBlockRootAtTimestamp(opts *bind.CallOpts, timestamp uint64) ([32]byte, error)

	GetPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error)

	HasPod(opts *bind.CallOpts, podOwner common.Address) (bool, error)

	MaxPods(opts *bind.CallOpts) (*big.Int, error)

	NumPods(opts *bind.CallOpts) (*big.Int, error)

	OwnerToPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error)

	PodOwnerShares(opts *bind.CallOpts, podOwner common.Address) (*big.Int, error)

	Slasher(opts *bind.CallOpts) (common.Address, error)

	StrategyManager(opts *bind.CallOpts) (common.Address, error)

	CreatePod(opts *bind.TransactOpts) (*types.Transaction, error)

	Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error)

	FilterBeaconChainETHDeposited(opts *bind.FilterOpts, podOwner []common.Address) (*ContractBeaconChainETHDepositedIterator, error)

	WatchBeaconChainETHDeposited(opts *bind.WatchOpts, sink chan<- *ContractBeaconChainETHDeposited, podOwner []common.Address) (event.Subscription, error)

	ParseBeaconChainETHDeposited(log types.Log) (*ContractBeaconChainETHDeposited, error)

	FilterBeaconChainETHWithdrawalCompleted(opts *bind.FilterOpts, podOwner []common.Address) (*ContractBeaconChainETHWithdrawalCompletedIterator, error)

	WatchBeaconChainETHWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ContractBeaconChainETHWithdrawalCompleted, podOwner []common.Address) (event.Subscription, error)

	ParseBeaconChainETHWithdrawalCompleted(log types.Log) (*ContractBeaconChainETHWithdrawalCompleted, error)

	FilterBeaconOracleUpdated(opts *bind.FilterOpts, newOracleAddress []common.Address) (*ContractBeaconOracleUpdatedIterator, error)

	WatchBeaconOracleUpdated(opts *bind.WatchOpts, sink chan<- *ContractBeaconOracleUpdated, newOracleAddress []common.Address) (event.Subscription, error)

	ParseBeaconOracleUpdated(log types.Log) (*ContractBeaconOracleUpdated, error)

	FilterMaxPodsUpdated(opts *bind.FilterOpts) (*ContractMaxPodsUpdatedIterator, error)

	WatchMaxPodsUpdated(opts *bind.WatchOpts, sink chan<- *ContractMaxPodsUpdated) (event.Subscription, error)

	ParseMaxPodsUpdated(log types.Log) (*ContractMaxPodsUpdated, error)

	FilterPodDeployed(opts *bind.FilterOpts, eigenPod []common.Address, podOwner []common.Address) (*ContractPodDeployedIterator, error)

	WatchPodDeployed(opts *bind.WatchOpts, sink chan<- *ContractPodDeployed, eigenPod []common.Address, podOwner []common.Address) (event.Subscription, error)

	ParsePodDeployed(log types.Log) (*ContractPodDeployed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
