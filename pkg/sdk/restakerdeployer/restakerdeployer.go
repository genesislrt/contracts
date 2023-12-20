// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package restakerdeployer

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beacon_\",\"type\":\"address\"},{\"internalType\":\"contractIRestakerFacets\",\"name\":\"facets_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIRestaker\",\"name\":\"restaker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"RestakerDeployed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BEACON_PROXY_BYTECODE\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beacon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployRestaker\",\"outputs\":[{\"internalType\":\"contractIRestaker\",\"name\":\"restaker\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"internalType\":\"contractIRestakerFacets\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getRestaker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561000f575f80fd5b50604051610d3b380380610d3b83398101604081905261002e9161005c565b6001600160a01b039182166080521660a052610094565b6001600160a01b0381168114610059575f80fd5b50565b5f806040838503121561006d575f80fd5b825161007881610045565b602084015190925061008981610045565b809150509250929050565b60805160a051610c736100c85f395f818160b701526101ee01525f818160720152818161016001526103040152610c735ff3fe608060405234801562000010575f80fd5b506004361062000068575f3560e01c806359659e90146200006c5780637a0ed62714620000b15780639ce953c814620000d9578063affed0e014620000e3578063cda3ef3614620000fb578063d52fc7101462000114575b5f80fd5b620000947f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b620000947f000000000000000000000000000000000000000000000000000000000000000081565b620000946200012b565b620000ec5f5481565b604051908152602001620000a8565b62000105620002ac565b604051620000a89190620004f6565b62000094620001253660046200052a565b620002d3565b5f80546040513391620001cf918491906200014960208201620004c4565b601f1982820381018352601f9091011660408181527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166020830152808201525f606082015260800160408051601f1981840301815290829052620001ba929160200162000542565b60405160208183030381529060405262000380565b60405163485cc95560e01b81526001600160a01b0383811660048301527f0000000000000000000000000000000000000000000000000000000000000000811660248301529193509083169063485cc955906044015f604051808303815f87803b1580156200023c575f80fd5b505af11580156200024f573d5f803e3d5ffd5b50505f80546001600160a01b038087169450851692507f66b1c85e3aa7b590e4fcd19543377d320772af5d49300c8c50653f46253ee99f9180620002938362000574565b9091555060405190815260200160405180910390a35090565b604051620002bd60208201620004c4565b601f1982820381018352601f9091011660405281565b5f6200037a825f1b60405180602001620002ed90620004c4565b601f1982820381018352601f9091011660408181527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166020830152808201525f606082015260800160408051601f19818403018152908290526200035e929160200162000542565b6040516020818303038152906040528051906020012062000493565b92915050565b5f83471015620003d75760405162461bcd60e51b815260206004820152601d60248201527f437265617465323a20696e73756666696369656e742062616c616e636500000060448201526064015b60405180910390fd5b81515f03620004295760405162461bcd60e51b815260206004820181905260248201527f437265617465323a2062797465636f6465206c656e677468206973207a65726f6044820152606401620003ce565b8282516020840186f590506001600160a01b0381166200048c5760405162461bcd60e51b815260206004820152601960248201527f437265617465323a204661696c6564206f6e206465706c6f79000000000000006044820152606401620003ce565b9392505050565b5f6200048c8383305f604051836040820152846020820152828152600b8101905060ff815360559020949350505050565b6106a4806200059a83390190565b5f5b83811015620004ee578181015183820152602001620004d4565b50505f910152565b602081525f825180602084015262000516816040850160208701620004d2565b601f01601f19169190910160400192915050565b5f602082840312156200053b575f80fd5b5035919050565b5f835162000555818460208801620004d2565b8351908301906200056b818360208801620004d2565b01949350505050565b5f600182016200059257634e487b7160e01b5f52601160045260245ffd5b506001019056fe60806040526040516106a43803806106a48339810160408190526100229161040f565b61002d82825f610034565b5050610530565b61003d836100f1565b6040516001600160a01b038416907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e905f90a25f8251118061007c5750805b156100ec576100ea836001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100e491906104ca565b83610273565b505b505050565b6001600160a01b0381163b61015b5760405162461bcd60e51b815260206004820152602560248201527f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e6044820152641d1c9858dd60da1b60648201526084015b60405180910390fd5b6101cd816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561019a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906101be91906104ca565b6001600160a01b03163b151590565b6102325760405162461bcd60e51b815260206004820152603060248201527f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960448201526f1cc81b9bdd08184818dbdb9d1c9858dd60821b6064820152608401610152565b7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5080546001600160a01b0319166001600160a01b0392909216919091179055565b6060610298838360405180606001604052806027815260200161067d6027913961029f565b9392505050565b60605f80856001600160a01b0316856040516102bb91906104e3565b5f60405180830381855af49150503d805f81146102f3576040519150601f19603f3d011682016040523d82523d5f602084013e6102f8565b606091505b50909250905061030a86838387610314565b9695505050505050565b606083156103825782515f0361037b576001600160a01b0385163b61037b5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610152565b508161038c565b61038c8383610394565b949350505050565b8151156103a45781518083602001fd5b8060405162461bcd60e51b815260040161015291906104fe565b80516001600160a01b03811681146103d4575f80fd5b919050565b634e487b7160e01b5f52604160045260245ffd5b5f5b838110156104075781810151838201526020016103ef565b50505f910152565b5f8060408385031215610420575f80fd5b610429836103be565b60208401519092506001600160401b0380821115610445575f80fd5b818501915085601f830112610458575f80fd5b81518181111561046a5761046a6103d9565b604051601f8201601f19908116603f01168101908382118183101715610492576104926103d9565b816040528281528860208487010111156104aa575f80fd5b6104bb8360208301602088016103ed565b80955050505050509250929050565b5f602082840312156104da575f80fd5b610298826103be565b5f82516104f48184602087016103ed565b9190910192915050565b602081525f825180602084015261051c8160408501602087016103ed565b601f01601f19169190910160400192915050565b6101408061053d5f395ff3fe60806040523661001357610011610017565b005b6100115b610027610022610029565b6100bf565b565b5f61005b7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50546001600160a01b031690565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610096573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906100ba91906100dd565b905090565b365f80375f80365f845af43d5f803e8080156100d9573d5ff35b3d5ffd5b5f602082840312156100ed575f80fd5b81516001600160a01b0381168114610103575f80fd5b939250505056fea264697066735822122012704b9aff4d502a7f596e970cc1747d07e2f4ac947e4b4045739d6d94e880b264736f6c63430008140033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212201692c5fb9672d3951a2be9f57ba7a8ce637306411ecb91482025cf9a0cb81cf264736f6c63430008140033",
}

var ContractABI = ContractMetaData.ABI

var ContractBin = ContractMetaData.Bin

func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, beacon_ common.Address, facets_ common.Address) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend, beacon_, facets_)
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

func (_Contract *ContractCaller) BEACONPROXYBYTECODE(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "BEACON_PROXY_BYTECODE")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_Contract *ContractSession) BEACONPROXYBYTECODE() ([]byte, error) {
	return _Contract.Contract.BEACONPROXYBYTECODE(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) BEACONPROXYBYTECODE() ([]byte, error) {
	return _Contract.Contract.BEACONPROXYBYTECODE(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) Beacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "beacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) Beacon() (common.Address, error) {
	return _Contract.Contract.Beacon(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) Beacon() (common.Address, error) {
	return _Contract.Contract.Beacon(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) Facets(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) Facets() (common.Address, error) {
	return _Contract.Contract.Facets(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) Facets() (common.Address, error) {
	return _Contract.Contract.Facets(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetRestaker(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRestaker", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetRestaker(id *big.Int) (common.Address, error) {
	return _Contract.Contract.GetRestaker(&_Contract.CallOpts, id)
}

func (_Contract *ContractCallerSession) GetRestaker(id *big.Int) (common.Address, error) {
	return _Contract.Contract.GetRestaker(&_Contract.CallOpts, id)
}

func (_Contract *ContractCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) Nonce() (*big.Int, error) {
	return _Contract.Contract.Nonce(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) Nonce() (*big.Int, error) {
	return _Contract.Contract.Nonce(&_Contract.CallOpts)
}

func (_Contract *ContractTransactor) DeployRestaker(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deployRestaker")
}

func (_Contract *ContractSession) DeployRestaker() (*types.Transaction, error) {
	return _Contract.Contract.DeployRestaker(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactorSession) DeployRestaker() (*types.Transaction, error) {
	return _Contract.Contract.DeployRestaker(&_Contract.TransactOpts)
}

type ContractRestakerDeployedIterator struct {
	Event *ContractRestakerDeployed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractRestakerDeployedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRestakerDeployed)
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
		it.Event = new(ContractRestakerDeployed)
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

func (it *ContractRestakerDeployedIterator) Error() error {
	return it.fail
}

func (it *ContractRestakerDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractRestakerDeployed struct {
	Creator  common.Address
	Restaker common.Address
	Id       *big.Int
	Raw      types.Log
}

func (_Contract *ContractFilterer) FilterRestakerDeployed(opts *bind.FilterOpts, creator []common.Address, restaker []common.Address) (*ContractRestakerDeployedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var restakerRule []interface{}
	for _, restakerItem := range restaker {
		restakerRule = append(restakerRule, restakerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RestakerDeployed", creatorRule, restakerRule)
	if err != nil {
		return nil, err
	}
	return &ContractRestakerDeployedIterator{contract: _Contract.contract, event: "RestakerDeployed", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchRestakerDeployed(opts *bind.WatchOpts, sink chan<- *ContractRestakerDeployed, creator []common.Address, restaker []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var restakerRule []interface{}
	for _, restakerItem := range restaker {
		restakerRule = append(restakerRule, restakerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RestakerDeployed", creatorRule, restakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractRestakerDeployed)
				if err := _Contract.contract.UnpackLog(event, "RestakerDeployed", log); err != nil {
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

func (_Contract *ContractFilterer) ParseRestakerDeployed(log types.Log) (*ContractRestakerDeployed, error) {
	event := new(ContractRestakerDeployed)
	if err := _Contract.contract.UnpackLog(event, "RestakerDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Contract *Contract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Contract.abi.Events["RestakerDeployed"].ID:
		return _Contract.ParseRestakerDeployed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ContractRestakerDeployed) Topic() common.Hash {
	return common.HexToHash("0x66b1c85e3aa7b590e4fcd19543377d320772af5d49300c8c50653f46253ee99f")
}

func (_Contract *Contract) Address() common.Address {
	return _Contract.address
}

type ContractInterface interface {
	BEACONPROXYBYTECODE(opts *bind.CallOpts) ([]byte, error)

	Beacon(opts *bind.CallOpts) (common.Address, error)

	Facets(opts *bind.CallOpts) (common.Address, error)

	GetRestaker(opts *bind.CallOpts, id *big.Int) (common.Address, error)

	Nonce(opts *bind.CallOpts) (*big.Int, error)

	DeployRestaker(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterRestakerDeployed(opts *bind.FilterOpts, creator []common.Address, restaker []common.Address) (*ContractRestakerDeployedIterator, error)

	WatchRestakerDeployed(opts *bind.WatchOpts, sink chan<- *ContractRestakerDeployed, creator []common.Address, restaker []common.Address) (event.Subscription, error)

	ParseRestakerDeployed(log types.Log) (*ContractRestakerDeployed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
