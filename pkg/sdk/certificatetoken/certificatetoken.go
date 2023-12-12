// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package certificatetoken

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bondsToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakingConfig\",\"name\":\"stakingConfig\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRebasing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ratio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sharesToBonds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506119db806100206000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c806370a08231116100c357806395d89b411161007c57806395d89b41146102b05780639dc29fac146102b8578063a457c2d7146102cb578063a9059cbb146102de578063dd62ed3e146102f1578063f2fde38b1461030457600080fd5b806370a0823114610241578063715018a61461026a57806371ca337d146102725780638456cb591461027a5780638da5cb5b14610282578063906571471461029d57600080fd5b80633f4ba83a116101155780633f4ba83a146101e557806340c10f19146101ef57806353616373146102025780635c975abb146102155780635dfba115146102275780636c58d43d1461022e57600080fd5b806306fdde031461015d578063095ea7b31461017b57806318160ddd1461019e57806323b872dd146101b0578063313ce567146101c357806339509351146101d2575b600080fd5b610165610317565b6040516101729190611411565b60405180910390f35b61018e61018936600461147b565b6103a9565b6040519015158152602001610172565b6067545b604051908152602001610172565b61018e6101be3660046114a7565b6103f7565b60405160128152602001610172565b61018e6101e036600461147b565b610446565b6101ed610468565b005b6101ed6101fd36600461147b565b6105b5565b6101a26102103660046114e8565b610694565b609754600160a01b900460ff1661018e565b600061018e565b6101a261023c3660046114e8565b6106b0565b6101a261024f366004611501565b6001600160a01b031660009081526065602052604090205490565b6101ed6106cc565b6101a26106e0565b6101ed6107cf565b6033546040516001600160a01b039091168152602001610172565b6101ed6102ab3660046115c1565b6108db565b610165610a14565b6101ed6102c636600461147b565b610a23565b61018e6102d936600461147b565b610afe565b61018e6102ec36600461147b565b610b79565b6101a26102ff366004611637565b610bb2565b6101ed610312366004611501565b610bdd565b60606068805461032690611670565b80601f016020809104026020016040519081016040528092919081815260200182805461035290611670565b801561039f5780601f106103745761010080835404028352916020019161039f565b820191906000526020600020905b81548152906001019060200180831161038257829003601f168201915b5050505050905090565b609754600090600160a01b900460ff16156103df5760405162461bcd60e51b81526004016103d6906116aa565b60405180910390fd5b336103eb818585610c56565b60019150505b92915050565b609754600090600160a01b900460ff16156104245760405162461bcd60e51b81526004016103d6906116aa565b33610430858285610d7b565b61043b858585610def565b506001949350505050565b6000336103eb8185856104598383610bb2565b61046391906116f7565b610c56565b609754600160a01b900460ff166104c15760405162461bcd60e51b815260206004820152601c60248201527f4365727469666963617465546f6b656e3a206e6f74207061757365640000000060448201526064016103d6565b609760009054906101000a90046001600160a01b03166001600160a01b031663732524946040518163ffffffff1660e01b8152600401602060405180830381865afa158015610514573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610538919061170f565b6001600160a01b0316336001600160a01b0316146105685760405162461bcd60e51b81526004016103d69061172c565b6097805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b609754600160a01b900460ff16156105df5760405162461bcd60e51b81526004016103d6906116aa565b609760009054906101000a90046001600160a01b03166001600160a01b031663b6217be76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610632573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610656919061170f565b6001600160a01b0316336001600160a01b0316146106865760405162461bcd60e51b81526004016103d690611775565b6106908282610f9a565b5050565b60006103f1826106a26106e0565b670de0b6b3a764000061105b565b60006103f182670de0b6b3a76400006106c76106e0565b6110a0565b6106d4611116565b6106de6000611170565b565b600080609760009054906101000a90046001600160a01b03166001600160a01b031663e87ff3446040518163ffffffff1660e01b8152600401602060405180830381865afa158015610736573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061075a919061170f565b60405163a1f1d48d60e01b81523060048201529091506000906001600160a01b0383169063a1f1d48d90602401602060405180830381865afa1580156107a4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107c891906117bf565b9392505050565b609754600160a01b900460ff16156107f95760405162461bcd60e51b81526004016103d6906116aa565b609760009054906101000a90046001600160a01b03166001600160a01b031663732524946040518163ffffffff1660e01b8152600401602060405180830381865afa15801561084c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610870919061170f565b6001600160a01b0316336001600160a01b0316146108a05760405162461bcd60e51b81526004016103d69061172c565b6097805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586105983390565b600054610100900460ff16158080156108fb5750600054600160ff909116105b806109155750303b158015610915575060005460ff166001145b6109785760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103d6565b6000805460ff19166001179055801561099b576000805461ff0019166101001790555b6109a36111c2565b6109ad83836111f1565b609780546001600160a01b0319166001600160a01b0386161790558015610a0e576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b60606069805461032690611670565b609754600160a01b900460ff1615610a4d5760405162461bcd60e51b81526004016103d6906116aa565b609760009054906101000a90046001600160a01b03166001600160a01b031663b6217be76040518163ffffffff1660e01b8152600401602060405180830381865afa158015610aa0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac4919061170f565b6001600160a01b0316336001600160a01b031614610af45760405162461bcd60e51b81526004016103d690611775565b6106908282611222565b60003381610b0c8286610bb2565b905083811015610b6c5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016103d6565b61043b8286868403610c56565b609754600090600160a01b900460ff1615610ba65760405162461bcd60e51b81526004016103d6906116aa565b336103eb818585610def565b6001600160a01b03918216600090815260666020908152604080832093909416825291909152205490565b610be5611116565b6001600160a01b038116610c4a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103d6565b610c5381611170565b50565b6001600160a01b038316610cb85760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016103d6565b6001600160a01b038216610d195760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103d6565b6001600160a01b0383811660008181526066602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6000610d878484610bb2565b90506000198114610a0e5781811015610de25760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103d6565b610a0e8484848403610c56565b6001600160a01b038316610e535760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103d6565b6001600160a01b038216610eb55760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103d6565b6001600160a01b03831660009081526065602052604090205481811015610f2d5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016103d6565b6001600160a01b0380851660008181526065602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90610f8d9086815260200190565b60405180910390a3610a0e565b6001600160a01b038216610ff05760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016103d6565b806067600082825461100291906116f7565b90915550506001600160a01b0382166000818152606560209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b600061109861107361106d84876117ee565b85611353565b838561107f8289611802565b6110899190611816565b61109391906117ee565b611389565b949350505050565b6000816000036110db5760405162461bcd60e51b8152602060048201526006602482015265063203d3d20360d41b60448201526064016103d6565b6110986110eb61106d84876117ee565b836110f7600182611835565b86611102878a611802565b61110c9190611816565b61108991906116f7565b6033546001600160a01b031633146106de5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103d6565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff166111e95760405162461bcd60e51b81526004016103d69061184c565b6106de6113a1565b600054610100900460ff166112185760405162461bcd60e51b81526004016103d69061184c565b61069082826113d1565b6001600160a01b0382166112825760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b60648201526084016103d6565b6001600160a01b038216600090815260656020526040902054818110156112f65760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b60648201526084016103d6565b6001600160a01b03831660008181526065602090815260408083208686039055606780548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610d6e565b505050565b600082600003611365575060006103f1565b82820282848281611378576113786117d8565b04146107c8576000199150506103f1565b6000828201838110156107c8576000199150506103f1565b600054610100900460ff166113c85760405162461bcd60e51b81526004016103d69061184c565b6106de33611170565b600054610100900460ff166113f85760405162461bcd60e51b81526004016103d69061184c565b606861140483826118e5565b50606961134e82826118e5565b600060208083528351808285015260005b8181101561143e57858101830151858201604001528201611422565b81811115611450576000604083870101525b50601f01601f1916929092016040019392505050565b6001600160a01b0381168114610c5357600080fd5b6000806040838503121561148e57600080fd5b823561149981611466565b946020939093013593505050565b6000806000606084860312156114bc57600080fd5b83356114c781611466565b925060208401356114d781611466565b929592945050506040919091013590565b6000602082840312156114fa57600080fd5b5035919050565b60006020828403121561151357600080fd5b81356107c881611466565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261154557600080fd5b813567ffffffffffffffff808211156115605761156061151e565b604051601f8301601f19908116603f011681019082821181831017156115885761158861151e565b816040528381528660208588010111156115a157600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156115d657600080fd5b83356115e181611466565b9250602084013567ffffffffffffffff808211156115fe57600080fd5b61160a87838801611534565b9350604086013591508082111561162057600080fd5b5061162d86828701611534565b9150509250925092565b6000806040838503121561164a57600080fd5b823561165581611466565b9150602083013561166581611466565b809150509250929050565b600181811c9082168061168457607f821691505b6020821081036116a457634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526018908201527f4365727469666963617465546f6b656e3a207061757365640000000000000000604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6000821982111561170a5761170a6116e1565b500190565b60006020828403121561172157600080fd5b81516107c881611466565b60208082526029908201527f4365727469666963617465546f6b656e3a206f6e6c7920676f7665726e616e636040820152681948185b1b1bddd95960ba1b606082015260800190565b6020808252602a908201527f4365727469666963617465546f6b656e3a206f6e6c79206c6971756964207374604082015269185ada5b99c81c1bdbdb60b21b606082015260800190565b6000602082840312156117d157600080fd5b5051919050565b634e487b7160e01b600052601260045260246000fd5b6000826117fd576117fd6117d8565b500490565b600082611811576118116117d8565b500690565b6000816000190483118215151615611830576118306116e1565b500290565b600082821015611847576118476116e1565b500390565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b601f82111561134e57600081815260208120601f850160051c810160208610156118be5750805b601f850160051c820191505b818110156118dd578281556001016118ca565b505050505050565b815167ffffffffffffffff8111156118ff576118ff61151e565b6119138161190d8454611670565b84611897565b602080601f83116001811461194857600084156119305750858301515b600019600386901b1c1916600185901b1785556118dd565b600085815260208120601f198616915b8281101561197757888601518255948401946001909101908401611958565b50858210156119955787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220f326a4c78d46536331ab98d43ddaff9c05e118ebfb8d9adc80becbf252154ab864736f6c634300080f0033",
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

func (_Contract *ContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, owner, spender)
}

func (_Contract *ContractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Contract.Contract.Allowance(&_Contract.CallOpts, owner, spender)
}

func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, account)
}

func (_Contract *ContractCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, account)
}

func (_Contract *ContractCaller) BondsToShares(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "bondsToShares", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) BondsToShares(amount *big.Int) (*big.Int, error) {
	return _Contract.Contract.BondsToShares(&_Contract.CallOpts, amount)
}

func (_Contract *ContractCallerSession) BondsToShares(amount *big.Int) (*big.Int, error) {
	return _Contract.Contract.BondsToShares(&_Contract.CallOpts, amount)
}

func (_Contract *ContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_Contract *ContractSession) Decimals() (uint8, error) {
	return _Contract.Contract.Decimals(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) Decimals() (uint8, error) {
	return _Contract.Contract.Decimals(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) IsRebasing(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isRebasing")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Contract *ContractSession) IsRebasing() (bool, error) {
	return _Contract.Contract.IsRebasing(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) IsRebasing() (bool, error) {
	return _Contract.Contract.IsRebasing(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_Contract *ContractSession) Paused() (bool, error) {
	return _Contract.Contract.Paused(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) Paused() (bool, error) {
	return _Contract.Contract.Paused(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) Ratio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ratio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) Ratio() (*big.Int, error) {
	return _Contract.Contract.Ratio(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) Ratio() (*big.Int, error) {
	return _Contract.Contract.Ratio(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) SharesToBonds(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "sharesToBonds", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) SharesToBonds(amount *big.Int) (*big.Int, error) {
	return _Contract.Contract.SharesToBonds(&_Contract.CallOpts, amount)
}

func (_Contract *ContractCallerSession) SharesToBonds(amount *big.Int) (*big.Int, error) {
	return _Contract.Contract.SharesToBonds(&_Contract.CallOpts, amount)
}

func (_Contract *ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

func (_Contract *ContractSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", spender, amount)
}

func (_Contract *ContractSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, spender, amount)
}

func (_Contract *ContractTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, spender, amount)
}

func (_Contract *ContractTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "burn", account, amount)
}

func (_Contract *ContractSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Burn(&_Contract.TransactOpts, account, amount)
}

func (_Contract *ContractTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Burn(&_Contract.TransactOpts, account, amount)
}

func (_Contract *ContractTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

func (_Contract *ContractSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DecreaseAllowance(&_Contract.TransactOpts, spender, subtractedValue)
}

func (_Contract *ContractTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DecreaseAllowance(&_Contract.TransactOpts, spender, subtractedValue)
}

func (_Contract *ContractTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

func (_Contract *ContractSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncreaseAllowance(&_Contract.TransactOpts, spender, addedValue)
}

func (_Contract *ContractTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.IncreaseAllowance(&_Contract.TransactOpts, spender, addedValue)
}

func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, stakingConfig common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", stakingConfig, name, symbol)
}

func (_Contract *ContractSession) Initialize(stakingConfig common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, stakingConfig, name, symbol)
}

func (_Contract *ContractTransactorSession) Initialize(stakingConfig common.Address, name string, symbol string) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, stakingConfig, name, symbol)
}

func (_Contract *ContractTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mint", account, amount)
}

func (_Contract *ContractSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Mint(&_Contract.TransactOpts, account, amount)
}

func (_Contract *ContractTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Mint(&_Contract.TransactOpts, account, amount)
}

func (_Contract *ContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "pause")
}

func (_Contract *ContractSession) Pause() (*types.Transaction, error) {
	return _Contract.Contract.Pause(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactorSession) Pause() (*types.Transaction, error) {
	return _Contract.Contract.Pause(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transfer", to, amount)
}

func (_Contract *ContractSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, to, amount)
}

func (_Contract *ContractTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Transfer(&_Contract.TransactOpts, to, amount)
}

func (_Contract *ContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom", from, to, amount)
}

func (_Contract *ContractSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, amount)
}

func (_Contract *ContractTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, from, to, amount)
}

func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

func (_Contract *ContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "unpause")
}

func (_Contract *ContractSession) Unpause() (*types.Transaction, error) {
	return _Contract.Contract.Unpause(&_Contract.TransactOpts)
}

func (_Contract *ContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _Contract.Contract.Unpause(&_Contract.TransactOpts)
}

type ContractApprovalIterator struct {
	Event *ContractApproval

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractApprovalIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
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
		it.Event = new(ContractApproval)
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

func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log
}

func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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

func (_Contract *ContractFilterer) ParseApproval(log types.Log) (*ContractApproval, error) {
	event := new(ContractApproval)
	if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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

type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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

func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log
}

func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractPausedIterator struct {
	Event *ContractPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPaused)
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
		it.Event = new(ContractPaused)
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

func (it *ContractPausedIterator) Error() error {
	return it.fail
}

func (it *ContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractPaused struct {
	Account common.Address
	Raw     types.Log
}

func (_Contract *ContractFilterer) FilterPaused(opts *bind.FilterOpts) (*ContractPausedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ContractPausedIterator{contract: _Contract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractPaused) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractPaused)
				if err := _Contract.contract.UnpackLog(event, "Paused", log); err != nil {
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

func (_Contract *ContractFilterer) ParsePaused(log types.Log) (*ContractPaused, error) {
	event := new(ContractPaused)
	if err := _Contract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractTransferIterator struct {
	Event *ContractTransfer

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractTransferIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransfer)
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
		it.Event = new(ContractTransfer)
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

func (it *ContractTransferIterator) Error() error {
	return it.fail
}

func (it *ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log
}

func (_Contract *ContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferIterator{contract: _Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractTransfer)
				if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

func (_Contract *ContractFilterer) ParseTransfer(log types.Log) (*ContractTransfer, error) {
	event := new(ContractTransfer)
	if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ContractUnpausedIterator struct {
	Event *ContractUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ContractUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUnpaused)
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
		it.Event = new(ContractUnpaused)
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

func (it *ContractUnpausedIterator) Error() error {
	return it.fail
}

func (it *ContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ContractUnpaused struct {
	Account common.Address
	Raw     types.Log
}

func (_Contract *ContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ContractUnpausedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ContractUnpausedIterator{contract: _Contract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

func (_Contract *ContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ContractUnpaused)
				if err := _Contract.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

func (_Contract *ContractFilterer) ParseUnpaused(log types.Log) (*ContractUnpaused, error) {
	event := new(ContractUnpaused)
	if err := _Contract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Contract *Contract) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _Contract.abi.Events["Approval"].ID:
		return _Contract.ParseApproval(log)
	case _Contract.abi.Events["Initialized"].ID:
		return _Contract.ParseInitialized(log)
	case _Contract.abi.Events["OwnershipTransferred"].ID:
		return _Contract.ParseOwnershipTransferred(log)
	case _Contract.abi.Events["Paused"].ID:
		return _Contract.ParsePaused(log)
	case _Contract.abi.Events["Transfer"].ID:
		return _Contract.ParseTransfer(log)
	case _Contract.abi.Events["Unpaused"].ID:
		return _Contract.ParseUnpaused(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ContractApproval) Topic() common.Hash {
	return common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
}

func (ContractInitialized) Topic() common.Hash {
	return common.HexToHash("0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498")
}

func (ContractOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (ContractPaused) Topic() common.Hash {
	return common.HexToHash("0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258")
}

func (ContractTransfer) Topic() common.Hash {
	return common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
}

func (ContractUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa")
}

func (_Contract *Contract) Address() common.Address {
	return _Contract.address
}

type ContractInterface interface {
	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)

	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)

	BondsToShares(opts *bind.CallOpts, amount *big.Int) (*big.Int, error)

	Decimals(opts *bind.CallOpts) (uint8, error)

	IsRebasing(opts *bind.CallOpts) (bool, error)

	Name(opts *bind.CallOpts) (string, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	Paused(opts *bind.CallOpts) (bool, error)

	Ratio(opts *bind.CallOpts) (*big.Int, error)

	SharesToBonds(opts *bind.CallOpts, amount *big.Int) (*big.Int, error)

	Symbol(opts *bind.CallOpts) (string, error)

	TotalSupply(opts *bind.CallOpts) (*big.Int, error)

	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error)

	IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, stakingConfig common.Address, name string, symbol string) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractApprovalIterator, error)

	WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error)

	ParseApproval(log types.Log) (*ContractApproval, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error)

	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error)

	ParseInitialized(log types.Log) (*ContractInitialized, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts) (*ContractPausedIterator, error)

	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractPaused) (event.Subscription, error)

	ParsePaused(log types.Log) (*ContractPaused, error)

	FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractTransferIterator, error)

	WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseTransfer(log types.Log) (*ContractTransfer, error)

	FilterUnpaused(opts *bind.FilterOpts) (*ContractUnpausedIterator, error)

	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractUnpaused) (event.Subscription, error)

	ParseUnpaused(log types.Log) (*ContractUnpaused, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
