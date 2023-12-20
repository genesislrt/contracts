// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakingconfig

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
	ABI: "[{\"inputs\":[],\"name\":\"OnlyGovernanceAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOperatorAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractICToken\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractICToken\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"CTokenChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"EigenManagerAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIEigenPodManager\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIEigenPodManager\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"EigenManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"GovernanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinStakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinUnstakeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"OperatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIRatioFeed\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIRatioFeed\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"RatioFeedChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIRestakerDeployer\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIRestakerDeployer\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"RestakerDeployerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIRestakingPool\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIRestakingPool\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"RestakingPoolChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevValue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"TreasuryChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getCToken\",\"outputs\":[{\"internalType\":\"contractICToken\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCertTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEigenPodManager\",\"outputs\":[{\"internalType\":\"contractIEigenPodManager\",\"name\":\"manager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEigenPodManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinUnstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRatioFeed\",\"outputs\":[{\"internalType\":\"contractIRatioFeed\",\"name\":\"feed\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRatioFeedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakerDeployer\",\"outputs\":[{\"internalType\":\"contractIRestakerDeployer\",\"name\":\"deployer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakingPool\",\"outputs\":[{\"internalType\":\"contractIRestakingPool\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumUnstake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governanceAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasuryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakingPoolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"certTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ratioFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"eigenPodManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setCToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setGovernanceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMinStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMinUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setOperatorAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setRatioFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRestakerDeployer\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setRestakerDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setRestakingPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506110bb8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610153575f3560e01c8063b6217be7116100bf578063e7f43c6811610079578063e7f43c6814610293578063e855f74a146102a4578063e87ff344146102b7578063ec6c350c146102c8578063f0e35a9a146102d0578063f0f44260146102e3575f80fd5b8063b6217be71461023b578063bd55a84e1461024c578063c5db8a7a1461025d578063c664d38b14610265578063cfc162541461026d578063dd2b442f14610280575f80fd5b806356a3b5fa1161011057806356a3b5fa146101e15780637745165b146101e9578063845b9150146101f15780638c80fd90146102045780638ce4ff16146102175780638ed872e61461022a575f80fd5b8063034c445414610157578063289b3c0d146101725780632f1d5a60146101975780633b19e84a146101ac5780633f69e0f7146101bd5780634cb71222146101ce575b5f80fd5b61015f6102f6565b6040519081526020015b60405180910390f35b6002546001600160a01b03165b6040516001600160a01b039091168152602001610169565b6101aa6101a5366004610e2c565b61031c565b005b6003546001600160a01b031661017f565b6008546001600160a01b031661017f565b6101aa6101dc366004610e2c565b6103d4565b61015f610477565b61017f610499565b6101aa6101ff366004610e4e565b6104ac565b6101aa610212366004610e4e565b6104e2565b6101aa610225366004610e65565b610515565b6006546001600160a01b031661017f565b6005546001600160a01b031661017f565b6007546001600160a01b031661017f565b61017f610885565b61017f610898565b6101aa61027b366004610e2c565b6108ab565b6101aa61028e366004610e2c565b61094e565b6001546001600160a01b031661017f565b6101aa6102b2366004610e2c565b610a07565b6004546001600160a01b031661017f565b61017f610aa9565b6101aa6102de366004610e2c565b610abc565b6101aa6102f1366004610e2c565b610b5f565b5f80546103179062010000900467ffffffffffffffff16633b9aca00610f1d565b905090565b6002546001600160a01b0316331461034f5760405162461bcd60e51b815260040161034690610f3a565b60405180910390fd5b6001600160a01b0381166103755760405162461bcd60e51b815260040161034690610f71565b600180546001600160a01b038381166001600160a01b03198316179092556040519116907fd58299b712891143e76310d5e664c4203c940a67db37cf856bdaa3c5c76a802c906103c89083908590610fb4565b60405180910390a15050565b6002546001600160a01b031633146103fe5760405162461bcd60e51b815260040161034690610f3a565b6001600160a01b0381166104245760405162461bcd60e51b815260040161034690610f71565b600780546001600160a01b038381166001600160a01b03198316179092556040519116907f604ae4b80bb1d3cb1b6f8fd99500a3203337ec3cdd83cb343ee91f8960f634df906103c89083908590610fb4565b5f805461031790600160501b900467ffffffffffffffff16633b9aca00610f1d565b5f6103176005546001600160a01b031690565b6002546001600160a01b031633146104d65760405162461bcd60e51b815260040161034690610f3a565b6104df81610c02565b50565b6002546001600160a01b0316331461050c5760405162461bcd60e51b815260040161034690610f3a565b6104df81610d0c565b5f54610100900460ff161580801561053357505f54600160ff909116105b8061054c5750303b15801561054c57505f5460ff166001145b6105af5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610346565b5f805460ff1916600117905580156105d0575f805461ff0019166101001790555b6105d989610d0c565b6105e28a610c02565b600180546001600160a01b0319166001600160a01b038a161790556040517fd58299b712891143e76310d5e664c4203c940a67db37cf856bdaa3c5c76a802c9061062f905f908b90610fb4565b60405180910390a1600780546001600160a01b0319166001600160a01b0386161790556040517f604ae4b80bb1d3cb1b6f8fd99500a3203337ec3cdd83cb343ee91f8960f634df90610684905f908790610fb4565b60405180910390a1600280546001600160a01b0319166001600160a01b0389161790556040517f3aaaebeb4821d6a7e5c77ece53cff0afcc56c82add2c978dbbb7f73e84cbcfd2906106d9905f908a90610fb4565b60405180910390a1600380546001600160a01b0319166001600160a01b0388161790556040517f8c3aa5f43a388513435861bf27dfad7829cd248696fed367c62d441f629544969061072e905f908990610fb4565b60405180910390a1600580546001600160a01b0319166001600160a01b0387161790556040517eae48a6cddea33b0b408d1f3e36bef3e47379bdc069c2f6662786c5bec83e1090610782905f908890610fb4565b60405180910390a1600680546001600160a01b0319166001600160a01b0384161790556040517f68d0770ff8eaea943d0b2f2020510cc23bbedaefd94add71cda3b4a4054b11c3906107d7905f908590610fb4565b60405180910390a1600480546001600160a01b0319166001600160a01b0385161790556040517fdb29c30d5fa0d3da86f28fcd1e16611171e924d291c7ef82f03cffb0bfa056529061082c905f908690610fb4565b60405180910390a18015610879575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050505050565b5f6103176004546001600160a01b031690565b5f6103176006546001600160a01b031690565b6002546001600160a01b031633146108d55760405162461bcd60e51b815260040161034690610f3a565b6001600160a01b0381166108fb5760405162461bcd60e51b815260040161034690610f71565b600280546001600160a01b038381166001600160a01b03198316179092556040519116907f3aaaebeb4821d6a7e5c77ece53cff0afcc56c82add2c978dbbb7f73e84cbcfd2906103c89083908590610fb4565b6002546001600160a01b031633146109785760405162461bcd60e51b815260040161034690610f3a565b6001600160a01b03811661099e5760405162461bcd60e51b815260040161034690610f71565b6008546040517f37910025f99bc7fdc07bdf77dee21c246391d5e3f98e8c6e3b0306dfaf8f24fa916109dd916001600160a01b03909116908490610fb4565b60405180910390a1600880546001600160a01b0319166001600160a01b0392909216919091179055565b6002546001600160a01b03163314610a315760405162461bcd60e51b815260040161034690610f3a565b6001600160a01b038116610a575760405162461bcd60e51b815260040161034690610f71565b600580546001600160a01b038381166001600160a01b03198316179092556040519116907eae48a6cddea33b0b408d1f3e36bef3e47379bdc069c2f6662786c5bec83e10906103c89083908590610fb4565b5f6103176007546001600160a01b031690565b6002546001600160a01b03163314610ae65760405162461bcd60e51b815260040161034690610f3a565b6001600160a01b038116610b0c5760405162461bcd60e51b815260040161034690610f71565b600480546001600160a01b038381166001600160a01b03198316179092556040519116907fdb29c30d5fa0d3da86f28fcd1e16611171e924d291c7ef82f03cffb0bfa05652906103c89083908590610fb4565b6002546001600160a01b03163314610b895760405162461bcd60e51b815260040161034690610f3a565b6001600160a01b038116610baf5760405162461bcd60e51b815260040161034690610f71565b600380546001600160a01b038381166001600160a01b03198316179092556040519116907f8c3aa5f43a388513435861bf27dfad7829cd248696fed367c62d441f62954496906103c89083908590610fb4565b610c10633b9aca0082610fe2565b15610c2d5760405162461bcd60e51b815260040161034690610ff5565b5f610c366102f6565b9050610c46633b9aca0083611046565b5f805469ffffffffffffffff000019166201000067ffffffffffffffff938416810291909117918290558492610c85929190910416633b9aca00611059565b67ffffffffffffffff1614610cd65760405162461bcd60e51b81526020600482015260176024820152765374616b696e67436f6e6669673a206f766572666c6f7760481b6044820152606401610346565b60408051828152602081018490527f984016d328adef81f3cc09f8ea9e3420f85d390635be94215c432e83837aa0a291016103c8565b610d1a633b9aca0082610fe2565b15610d375760405162461bcd60e51b815260040161034690610ff5565b5f610d40610477565b9050610d50633b9aca0083611046565b5f805467ffffffffffffffff60501b1916600160501b67ffffffffffffffff938416810291909117918290558492610d91929190910416633b9aca00611059565b67ffffffffffffffff1614610de25760405162461bcd60e51b81526020600482015260176024820152765374616b696e67436f6e6669673a206f766572666c6f7760481b6044820152606401610346565b60408051828152602081018490527fca11c8a4c461b60c9f485404c272650c2aaae260b2067d72e9924abb6855659391016103c8565b6001600160a01b03811681146104df575f80fd5b5f60208284031215610e3c575f80fd5b8135610e4781610e18565b9392505050565b5f60208284031215610e5e575f80fd5b5035919050565b5f805f805f805f805f6101208a8c031215610e7e575f80fd5b8935985060208a0135975060408a0135610e9781610e18565b965060608a0135610ea781610e18565b955060808a0135610eb781610e18565b945060a08a0135610ec781610e18565b935060c08a0135610ed781610e18565b925060e08a0135610ee781610e18565b91506101008a0135610ef881610e18565b809150509295985092959850929598565b634e487b7160e01b5f52601160045260245ffd5b8082028115828204841417610f3457610f34610f09565b92915050565b6020808252601e908201527f5374616b696e67436f6e6669673a206f6e6c7920676f7665726e616e63650000604082015260600190565b60208082526023908201527f5374616b696e67436f6e6669673a20616464726573732063616e2774206265206040820152621b9a5b60ea1b606082015260800190565b6001600160a01b0392831681529116602082015260400190565b634e487b7160e01b5f52601260045260245ffd5b5f82610ff057610ff0610fce565b500690565b60208082526031908201527f5374616b696e67436f6e6669673a2076616c75652073686f756c64206265206d604082015270756c7469706c696564206f66206777656960781b606082015260800190565b5f8261105457611054610fce565b500490565b67ffffffffffffffff81811683821602808216919082811461107d5761107d610f09565b50509291505056fea264697066735822122037bb4e1f41cede0c037f480e28be5b9dbb7c7d655e732da29d6d28e5625a8d2164736f6c63430008140033",
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

func (_Contract *ContractCaller) GetCToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getCToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetCToken() (common.Address, error) {
	return _Contract.Contract.GetCToken(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetCToken() (common.Address, error) {
	return _Contract.Contract.GetCToken(&_Contract.CallOpts)
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

func (_Contract *ContractCaller) GetEigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetEigenPodManager() (common.Address, error) {
	return _Contract.Contract.GetEigenPodManager(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetEigenPodManager() (common.Address, error) {
	return _Contract.Contract.GetEigenPodManager(&_Contract.CallOpts)
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

func (_Contract *ContractCaller) GetGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetGovernance() (common.Address, error) {
	return _Contract.Contract.GetGovernance(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetGovernance() (common.Address, error) {
	return _Contract.Contract.GetGovernance(&_Contract.CallOpts)
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

func (_Contract *ContractCaller) GetOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetOperator() (common.Address, error) {
	return _Contract.Contract.GetOperator(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetOperator() (common.Address, error) {
	return _Contract.Contract.GetOperator(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetRatioFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRatioFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetRatioFeed() (common.Address, error) {
	return _Contract.Contract.GetRatioFeed(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetRatioFeed() (common.Address, error) {
	return _Contract.Contract.GetRatioFeed(&_Contract.CallOpts)
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

func (_Contract *ContractCaller) GetRestakerDeployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRestakerDeployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetRestakerDeployer() (common.Address, error) {
	return _Contract.Contract.GetRestakerDeployer(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetRestakerDeployer() (common.Address, error) {
	return _Contract.Contract.GetRestakerDeployer(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) GetRestakingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRestakingPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetRestakingPool() (common.Address, error) {
	return _Contract.Contract.GetRestakingPool(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetRestakingPool() (common.Address, error) {
	return _Contract.Contract.GetRestakingPool(&_Contract.CallOpts)
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

func (_Contract *ContractCaller) GetTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) GetTreasury() (common.Address, error) {
	return _Contract.Contract.GetTreasury(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) GetTreasury() (common.Address, error) {
	return _Contract.Contract.GetTreasury(&_Contract.CallOpts)
}

func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, minimumUnstake *big.Int, minimumStake *big.Int, operatorAddress common.Address, governanceAddress common.Address, treasuryAddress common.Address, stakingPoolAddress common.Address, certTokenAddress common.Address, ratioFeed common.Address, eigenPodManager common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", minimumUnstake, minimumStake, operatorAddress, governanceAddress, treasuryAddress, stakingPoolAddress, certTokenAddress, ratioFeed, eigenPodManager)
}

func (_Contract *ContractSession) Initialize(minimumUnstake *big.Int, minimumStake *big.Int, operatorAddress common.Address, governanceAddress common.Address, treasuryAddress common.Address, stakingPoolAddress common.Address, certTokenAddress common.Address, ratioFeed common.Address, eigenPodManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, minimumUnstake, minimumStake, operatorAddress, governanceAddress, treasuryAddress, stakingPoolAddress, certTokenAddress, ratioFeed, eigenPodManager)
}

func (_Contract *ContractTransactorSession) Initialize(minimumUnstake *big.Int, minimumStake *big.Int, operatorAddress common.Address, governanceAddress common.Address, treasuryAddress common.Address, stakingPoolAddress common.Address, certTokenAddress common.Address, ratioFeed common.Address, eigenPodManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, minimumUnstake, minimumStake, operatorAddress, governanceAddress, treasuryAddress, stakingPoolAddress, certTokenAddress, ratioFeed, eigenPodManager)
}

func (_Contract *ContractTransactor) SetCToken(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setCToken", newValue)
}

func (_Contract *ContractSession) SetCToken(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetCToken(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetCToken(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetCToken(&_Contract.TransactOpts, newValue)
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

func (_Contract *ContractTransactor) SetRatioFeed(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setRatioFeed", newValue)
}

func (_Contract *ContractSession) SetRatioFeed(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetRatioFeed(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetRatioFeed(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetRatioFeed(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) SetRestakerDeployer(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setRestakerDeployer", newValue)
}

func (_Contract *ContractSession) SetRestakerDeployer(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetRestakerDeployer(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetRestakerDeployer(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetRestakerDeployer(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) SetRestakingPool(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setRestakingPool", newValue)
}

func (_Contract *ContractSession) SetRestakingPool(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetRestakingPool(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetRestakingPool(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetRestakingPool(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactor) SetTreasury(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setTreasury", newValue)
}

func (_Contract *ContractSession) SetTreasury(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetTreasury(&_Contract.TransactOpts, newValue)
}

func (_Contract *ContractTransactorSession) SetTreasury(newValue common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetTreasury(&_Contract.TransactOpts, newValue)
}

type ContractCTokenChangedIterator struct {
	Event *ContractCTokenChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractCTokenChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractCTokenChanged)
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
		it.Event = new(ContractCTokenChanged)
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

func (it *ContractCTokenChangedIterator) Error() error {
	return it.fail
}

func (it *ContractCTokenChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractCTokenChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterCTokenChanged(opts *bind.FilterOpts) (*ContractCTokenChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "CTokenChanged")
	if err != nil {
		return nil, err
	}
	return &ContractCTokenChangedIterator{contract: _Contract.contract, event: "CTokenChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchCTokenChanged(opts *bind.WatchOpts, sink chan<- *ContractCTokenChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "CTokenChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractCTokenChanged)
				if err := _Contract.contract.UnpackLog(event, "CTokenChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseCTokenChanged(log types.Log) (*ContractCTokenChanged, error) {
	event := new(ContractCTokenChanged)
	if err := _Contract.contract.UnpackLog(event, "CTokenChanged", log); err != nil {
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

type ContractEigenManagerChangedIterator struct {
	Event *ContractEigenManagerChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractEigenManagerChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEigenManagerChanged)
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
		it.Event = new(ContractEigenManagerChanged)
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

func (it *ContractEigenManagerChangedIterator) Error() error {
	return it.fail
}

func (it *ContractEigenManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractEigenManagerChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterEigenManagerChanged(opts *bind.FilterOpts) (*ContractEigenManagerChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EigenManagerChanged")
	if err != nil {
		return nil, err
	}
	return &ContractEigenManagerChangedIterator{contract: _Contract.contract, event: "EigenManagerChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchEigenManagerChanged(opts *bind.WatchOpts, sink chan<- *ContractEigenManagerChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EigenManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractEigenManagerChanged)
				if err := _Contract.contract.UnpackLog(event, "EigenManagerChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseEigenManagerChanged(log types.Log) (*ContractEigenManagerChanged, error) {
	event := new(ContractEigenManagerChanged)
	if err := _Contract.contract.UnpackLog(event, "EigenManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractGovernanceChangedIterator struct {
	Event *ContractGovernanceChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractGovernanceChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractGovernanceChanged)
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
		it.Event = new(ContractGovernanceChanged)
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

func (it *ContractGovernanceChangedIterator) Error() error {
	return it.fail
}

func (it *ContractGovernanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractGovernanceChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterGovernanceChanged(opts *bind.FilterOpts) (*ContractGovernanceChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "GovernanceChanged")
	if err != nil {
		return nil, err
	}
	return &ContractGovernanceChangedIterator{contract: _Contract.contract, event: "GovernanceChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchGovernanceChanged(opts *bind.WatchOpts, sink chan<- *ContractGovernanceChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "GovernanceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractGovernanceChanged)
				if err := _Contract.contract.UnpackLog(event, "GovernanceChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseGovernanceChanged(log types.Log) (*ContractGovernanceChanged, error) {
	event := new(ContractGovernanceChanged)
	if err := _Contract.contract.UnpackLog(event, "GovernanceChanged", log); err != nil {
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

type ContractOperatorChangedIterator struct {
	Event *ContractOperatorChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractOperatorChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOperatorChanged)
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
		it.Event = new(ContractOperatorChanged)
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

func (it *ContractOperatorChangedIterator) Error() error {
	return it.fail
}

func (it *ContractOperatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractOperatorChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterOperatorChanged(opts *bind.FilterOpts) (*ContractOperatorChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OperatorChanged")
	if err != nil {
		return nil, err
	}
	return &ContractOperatorChangedIterator{contract: _Contract.contract, event: "OperatorChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchOperatorChanged(opts *bind.WatchOpts, sink chan<- *ContractOperatorChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OperatorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractOperatorChanged)
				if err := _Contract.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseOperatorChanged(log types.Log) (*ContractOperatorChanged, error) {
	event := new(ContractOperatorChanged)
	if err := _Contract.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractRatioFeedChangedIterator struct {
	Event *ContractRatioFeedChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractRatioFeedChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRatioFeedChanged)
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
		it.Event = new(ContractRatioFeedChanged)
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

func (it *ContractRatioFeedChangedIterator) Error() error {
	return it.fail
}

func (it *ContractRatioFeedChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractRatioFeedChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterRatioFeedChanged(opts *bind.FilterOpts) (*ContractRatioFeedChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RatioFeedChanged")
	if err != nil {
		return nil, err
	}
	return &ContractRatioFeedChangedIterator{contract: _Contract.contract, event: "RatioFeedChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchRatioFeedChanged(opts *bind.WatchOpts, sink chan<- *ContractRatioFeedChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RatioFeedChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractRatioFeedChanged)
				if err := _Contract.contract.UnpackLog(event, "RatioFeedChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseRatioFeedChanged(log types.Log) (*ContractRatioFeedChanged, error) {
	event := new(ContractRatioFeedChanged)
	if err := _Contract.contract.UnpackLog(event, "RatioFeedChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractRestakerDeployerChangedIterator struct {
	Event *ContractRestakerDeployerChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractRestakerDeployerChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRestakerDeployerChanged)
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
		it.Event = new(ContractRestakerDeployerChanged)
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

func (it *ContractRestakerDeployerChangedIterator) Error() error {
	return it.fail
}

func (it *ContractRestakerDeployerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractRestakerDeployerChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterRestakerDeployerChanged(opts *bind.FilterOpts) (*ContractRestakerDeployerChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RestakerDeployerChanged")
	if err != nil {
		return nil, err
	}
	return &ContractRestakerDeployerChangedIterator{contract: _Contract.contract, event: "RestakerDeployerChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchRestakerDeployerChanged(opts *bind.WatchOpts, sink chan<- *ContractRestakerDeployerChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RestakerDeployerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractRestakerDeployerChanged)
				if err := _Contract.contract.UnpackLog(event, "RestakerDeployerChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseRestakerDeployerChanged(log types.Log) (*ContractRestakerDeployerChanged, error) {
	event := new(ContractRestakerDeployerChanged)
	if err := _Contract.contract.UnpackLog(event, "RestakerDeployerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractRestakingPoolChangedIterator struct {
	Event *ContractRestakingPoolChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractRestakingPoolChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRestakingPoolChanged)
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
		it.Event = new(ContractRestakingPoolChanged)
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

func (it *ContractRestakingPoolChangedIterator) Error() error {
	return it.fail
}

func (it *ContractRestakingPoolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractRestakingPoolChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterRestakingPoolChanged(opts *bind.FilterOpts) (*ContractRestakingPoolChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RestakingPoolChanged")
	if err != nil {
		return nil, err
	}
	return &ContractRestakingPoolChangedIterator{contract: _Contract.contract, event: "RestakingPoolChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchRestakingPoolChanged(opts *bind.WatchOpts, sink chan<- *ContractRestakingPoolChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RestakingPoolChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractRestakingPoolChanged)
				if err := _Contract.contract.UnpackLog(event, "RestakingPoolChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseRestakingPoolChanged(log types.Log) (*ContractRestakingPoolChanged, error) {
	event := new(ContractRestakingPoolChanged)
	if err := _Contract.contract.UnpackLog(event, "RestakingPoolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractTreasuryChangedIterator struct {
	Event *ContractTreasuryChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractTreasuryChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTreasuryChanged)
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
		it.Event = new(ContractTreasuryChanged)
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

func (it *ContractTreasuryChangedIterator) Error() error {
	return it.fail
}

func (it *ContractTreasuryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractTreasuryChanged struct {
	PrevValue common.Address
	NewValue  common.Address
	Raw       types.Log
}

func (_Contract *ContractFilterer) FilterTreasuryChanged(opts *bind.FilterOpts) (*ContractTreasuryChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TreasuryChanged")
	if err != nil {
		return nil, err
	}
	return &ContractTreasuryChangedIterator{contract: _Contract.contract, event: "TreasuryChanged", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchTreasuryChanged(opts *bind.WatchOpts, sink chan<- *ContractTreasuryChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TreasuryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractTreasuryChanged)
				if err := _Contract.contract.UnpackLog(event, "TreasuryChanged", log); err != nil {
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

func (_Contract *ContractFilterer) ParseTreasuryChanged(log types.Log) (*ContractTreasuryChanged, error) {
	event := new(ContractTreasuryChanged)
	if err := _Contract.contract.UnpackLog(event, "TreasuryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Contract *Contract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Contract.abi.Events["CTokenChanged"].ID:
		return _Contract.ParseCTokenChanged(log)
	case _Contract.abi.Events["EigenManagerAddressChanged"].ID:
		return _Contract.ParseEigenManagerAddressChanged(log)
	case _Contract.abi.Events["EigenManagerChanged"].ID:
		return _Contract.ParseEigenManagerChanged(log)
	case _Contract.abi.Events["GovernanceChanged"].ID:
		return _Contract.ParseGovernanceChanged(log)
	case _Contract.abi.Events["Initialized"].ID:
		return _Contract.ParseInitialized(log)
	case _Contract.abi.Events["MinStakeChanged"].ID:
		return _Contract.ParseMinStakeChanged(log)
	case _Contract.abi.Events["MinUnstakeChanged"].ID:
		return _Contract.ParseMinUnstakeChanged(log)
	case _Contract.abi.Events["OperatorChanged"].ID:
		return _Contract.ParseOperatorChanged(log)
	case _Contract.abi.Events["RatioFeedChanged"].ID:
		return _Contract.ParseRatioFeedChanged(log)
	case _Contract.abi.Events["RestakerDeployerChanged"].ID:
		return _Contract.ParseRestakerDeployerChanged(log)
	case _Contract.abi.Events["RestakingPoolChanged"].ID:
		return _Contract.ParseRestakingPoolChanged(log)
	case _Contract.abi.Events["TreasuryChanged"].ID:
		return _Contract.ParseTreasuryChanged(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ContractCTokenChanged) Topic() common.Hash {
	return common.HexToHash("0x604ae4b80bb1d3cb1b6f8fd99500a3203337ec3cdd83cb343ee91f8960f634df")
}

func (ContractEigenManagerAddressChanged) Topic() common.Hash {
	return common.HexToHash("0x68d0770ff8eaea943d0b2f2020510cc23bbedaefd94add71cda3b4a4054b11c3")
}

func (ContractEigenManagerChanged) Topic() common.Hash {
	return common.HexToHash("0xdce0fb87deaee9a5681d737cdbc7ba6fc10d8fac2e72e52589fff9dbbae06abb")
}

func (ContractGovernanceChanged) Topic() common.Hash {
	return common.HexToHash("0x3aaaebeb4821d6a7e5c77ece53cff0afcc56c82add2c978dbbb7f73e84cbcfd2")
}

func (ContractInitialized) Topic() common.Hash {
	return common.HexToHash("0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498")
}

func (ContractMinStakeChanged) Topic() common.Hash {
	return common.HexToHash("0xca11c8a4c461b60c9f485404c272650c2aaae260b2067d72e9924abb68556593")
}

func (ContractMinUnstakeChanged) Topic() common.Hash {
	return common.HexToHash("0x984016d328adef81f3cc09f8ea9e3420f85d390635be94215c432e83837aa0a2")
}

func (ContractOperatorChanged) Topic() common.Hash {
	return common.HexToHash("0xd58299b712891143e76310d5e664c4203c940a67db37cf856bdaa3c5c76a802c")
}

func (ContractRatioFeedChanged) Topic() common.Hash {
	return common.HexToHash("0xdb29c30d5fa0d3da86f28fcd1e16611171e924d291c7ef82f03cffb0bfa05652")
}

func (ContractRestakerDeployerChanged) Topic() common.Hash {
	return common.HexToHash("0x37910025f99bc7fdc07bdf77dee21c246391d5e3f98e8c6e3b0306dfaf8f24fa")
}

func (ContractRestakingPoolChanged) Topic() common.Hash {
	return common.HexToHash("0x00ae48a6cddea33b0b408d1f3e36bef3e47379bdc069c2f6662786c5bec83e10")
}

func (ContractTreasuryChanged) Topic() common.Hash {
	return common.HexToHash("0x8c3aa5f43a388513435861bf27dfad7829cd248696fed367c62d441f62954496")
}

func (_Contract *Contract) Address() common.Address {
	return _Contract.address
}

type ContractInterface interface {
	GetCToken(opts *bind.CallOpts) (common.Address, error)

	GetCertTokenAddress(opts *bind.CallOpts) (common.Address, error)

	GetEigenPodManager(opts *bind.CallOpts) (common.Address, error)

	GetEigenPodManagerAddress(opts *bind.CallOpts) (common.Address, error)

	GetGovernance(opts *bind.CallOpts) (common.Address, error)

	GetMinStake(opts *bind.CallOpts) (*big.Int, error)

	GetMinUnstake(opts *bind.CallOpts) (*big.Int, error)

	GetOperator(opts *bind.CallOpts) (common.Address, error)

	GetRatioFeed(opts *bind.CallOpts) (common.Address, error)

	GetRatioFeedAddress(opts *bind.CallOpts) (common.Address, error)

	GetRestakerDeployer(opts *bind.CallOpts) (common.Address, error)

	GetRestakingPool(opts *bind.CallOpts) (common.Address, error)

	GetStakingPoolAddress(opts *bind.CallOpts) (common.Address, error)

	GetTreasury(opts *bind.CallOpts) (common.Address, error)

	Initialize(opts *bind.TransactOpts, minimumUnstake *big.Int, minimumStake *big.Int, operatorAddress common.Address, governanceAddress common.Address, treasuryAddress common.Address, stakingPoolAddress common.Address, certTokenAddress common.Address, ratioFeed common.Address, eigenPodManager common.Address) (*types.Transaction, error)

	SetCToken(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	SetGovernanceAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	SetMinStake(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error)

	SetMinUnstake(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error)

	SetOperatorAddress(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	SetRatioFeed(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	SetRestakerDeployer(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	SetRestakingPool(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	SetTreasury(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error)

	FilterCTokenChanged(opts *bind.FilterOpts) (*ContractCTokenChangedIterator, error)

	WatchCTokenChanged(opts *bind.WatchOpts, sink chan<- *ContractCTokenChanged) (event.Subscription, error)

	ParseCTokenChanged(log types.Log) (*ContractCTokenChanged, error)

	FilterEigenManagerAddressChanged(opts *bind.FilterOpts) (*ContractEigenManagerAddressChangedIterator, error)

	WatchEigenManagerAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractEigenManagerAddressChanged) (event.Subscription, error)

	ParseEigenManagerAddressChanged(log types.Log) (*ContractEigenManagerAddressChanged, error)

	FilterEigenManagerChanged(opts *bind.FilterOpts) (*ContractEigenManagerChangedIterator, error)

	WatchEigenManagerChanged(opts *bind.WatchOpts, sink chan<- *ContractEigenManagerChanged) (event.Subscription, error)

	ParseEigenManagerChanged(log types.Log) (*ContractEigenManagerChanged, error)

	FilterGovernanceChanged(opts *bind.FilterOpts) (*ContractGovernanceChangedIterator, error)

	WatchGovernanceChanged(opts *bind.WatchOpts, sink chan<- *ContractGovernanceChanged) (event.Subscription, error)

	ParseGovernanceChanged(log types.Log) (*ContractGovernanceChanged, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*ContractInitialized, error)

	FilterMinStakeChanged(opts *bind.FilterOpts) (*ContractMinStakeChangedIterator, error)

	WatchMinStakeChanged(opts *bind.WatchOpts, sink chan<- *ContractMinStakeChanged) (event.Subscription, error)

	ParseMinStakeChanged(log types.Log) (*ContractMinStakeChanged, error)

	FilterMinUnstakeChanged(opts *bind.FilterOpts) (*ContractMinUnstakeChangedIterator, error)

	WatchMinUnstakeChanged(opts *bind.WatchOpts, sink chan<- *ContractMinUnstakeChanged) (event.Subscription, error)

	ParseMinUnstakeChanged(log types.Log) (*ContractMinUnstakeChanged, error)

	FilterOperatorChanged(opts *bind.FilterOpts) (*ContractOperatorChangedIterator, error)

	WatchOperatorChanged(opts *bind.WatchOpts, sink chan<- *ContractOperatorChanged) (event.Subscription, error)

	ParseOperatorChanged(log types.Log) (*ContractOperatorChanged, error)

	FilterRatioFeedChanged(opts *bind.FilterOpts) (*ContractRatioFeedChangedIterator, error)

	WatchRatioFeedChanged(opts *bind.WatchOpts, sink chan<- *ContractRatioFeedChanged) (event.Subscription, error)

	ParseRatioFeedChanged(log types.Log) (*ContractRatioFeedChanged, error)

	FilterRestakerDeployerChanged(opts *bind.FilterOpts) (*ContractRestakerDeployerChangedIterator, error)

	WatchRestakerDeployerChanged(opts *bind.WatchOpts, sink chan<- *ContractRestakerDeployerChanged) (event.Subscription, error)

	ParseRestakerDeployerChanged(log types.Log) (*ContractRestakerDeployerChanged, error)

	FilterRestakingPoolChanged(opts *bind.FilterOpts) (*ContractRestakingPoolChangedIterator, error)

	WatchRestakingPoolChanged(opts *bind.WatchOpts, sink chan<- *ContractRestakingPoolChanged) (event.Subscription, error)

	ParseRestakingPoolChanged(log types.Log) (*ContractRestakingPoolChanged, error)

	FilterTreasuryChanged(opts *bind.FilterOpts) (*ContractTreasuryChangedIterator, error)

	WatchTreasuryChanged(opts *bind.WatchOpts, sink chan<- *ContractTreasuryChanged) (event.Subscription, error)

	ParseTreasuryChanged(log types.Log) (*ContractTreasuryChanged, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
