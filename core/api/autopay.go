// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_client\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"E\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"N\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chalTags\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"client\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChIndex\",\"outputs\":[{\"internalType\":\"uint256[][]\",\"name\":\"\",\"type\":\"uint256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGs\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getN\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gs\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"indexs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proofs\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"sendViaCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"server\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_chaltags\",\"type\":\"bytes[]\"}],\"name\":\"setChals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[][]\",\"name\":\"_index\",\"type\":\"uint256[][]\"},{\"internalType\":\"bytes\",\"name\":\"_gs\",\"type\":\"bytes\"}],\"name\":\"setIndexAndGs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_N\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_E\",\"type\":\"bytes\"}],\"name\":\"setPk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_proof\",\"type\":\"bytes[]\"}],\"name\":\"setProve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405260405161139a38038061139a8339810160408190526100229161007f565b60008054336001600160a01b031991821617909155600180549091166001600160a01b0392909216919091179055346002556005805460ff199081169091556007805482169055600a805482169055600c805490911690556100af565b60006020828403121561009157600080fd5b81516001600160a01b03811681146100a857600080fd5b9392505050565b6112dc806100be6000396000f3fe6080604052600436106101135760003560e01c8063830c29ae116100a0578063c9e525df11610064578063c9e525df146102c8578063df5d093e146102dd578063e258eb7d146102fd578063fc735e991461031d578063fd922a421461032557600080fd5b8063830c29ae1461023e5780638947157e1461025357806392bbf6e8146102735780639ddaf5aa14610288578063a960f784146102a857600080fd5b806322e185e7116100e757806322e185e7146101bc57806322f9c360146101d157806338d52e0f146101f35780633e9552251461020957806371a76c7e1461021e57600080fd5b8062e71aaf146101185780630a36d18b14610143578063109e94cf1461017157806312065fe0146101a9575b600080fd5b34801561012457600080fd5b5061012d610345565b60405161013a9190610cbc565b60405180910390f35b34801561014f57600080fd5b5061016361015e366004610d11565b6103d3565b60405190815260200161013a565b34801561017d57600080fd5b50600154610191906001600160a01b031681565b6040516001600160a01b03909116815260200161013a565b3480156101b557600080fd5b5047610163565b3480156101c857600080fd5b5061012d610410565b3480156101dd57600080fd5b506101e66104a2565b60405161013a9190610d33565b3480156101ff57600080fd5b5061016360025481565b34801561021557600080fd5b5061012d61053b565b34801561022a57600080fd5b5061012d610239366004610dbd565b61054a565b61025161024c366004610dd6565b610575565b005b34801561025f57600080fd5b5061025161026e366004610ee1565b610622565b34801561027f57600080fd5b5061012d610700565b34801561029457600080fd5b5061012d6102a3366004610dbd565b61070d565b3480156102b457600080fd5b506102516102c3366004610ee1565b61071d565b3480156102d457600080fd5b5061012d61079b565b3480156102e957600080fd5b506102516102f8366004610f97565b6107a8565b34801561030957600080fd5b506102516103183660046110c1565b61084b565b6102516108cc565b34801561033157600080fd5b50600054610191906001600160a01b031681565b600680546103529061111b565b80601f016020809104026020016040519081016040528092919081815260200182805461037e9061111b565b80156103cb5780601f106103a0576101008083540402835291602001916103cb565b820191906000526020600020905b8154815290600101906020018083116103ae57829003601f168201915b505050505081565b600882815481106103e357600080fd5b9060005260206000200181815481106103fb57600080fd5b90600052602060002001600091509150505481565b60606006805461041f9061111b565b80601f016020809104026020016040519081016040528092919081815260200182805461044b9061111b565b80156104985780601f1061046d57610100808354040283529160200191610498565b820191906000526020600020905b81548152906001019060200180831161047b57829003601f168201915b5050505050905090565b60606008805480602002602001604051908101604052809291908181526020016000905b828210156105325760008481526020908190208301805460408051828502810185019091528181529283018282801561051e57602002820191906000526020600020905b81548152602001906001019080831161050a575b5050505050815260200190600101906104c6565b50505050905090565b60606004805461041f9061111b565b6009818154811061055a57600080fd5b9060005260206000200160009150905080546103529061111b565b600080826001600160a01b0316662386f26fc1000060405160006040518083038185875af1925050503d80600081146105ca576040519150601f19603f3d011682016040523d82523d6000602084013e6105cf565b606091505b50915091508161061d5760405162461bcd60e51b81526020600482015260146024820152732330b4b632b2103a379039b2b7321022ba3432b960611b60448201526064015b60405180910390fd5b505050565b6001546001600160a01b0316331461067c5760405162461bcd60e51b815260206004820152601e60248201527f4368616c6c656e67652073686f756c642073657420627920636c69656e7400006044820152606401610614565b600c5460ff166106dc5760405162461bcd60e51b815260206004820152602560248201527f50726f6f662073686f756c6420736574206265666f7265206368616c6c656e6760448201526419481cd95d60da1b6064820152608401610614565b80516106ef906009906020840190610aa5565b5050600a805460ff19166001179055565b600380546103529061111b565b600b818154811061055a57600080fd5b6000546001600160a01b031633146107775760405162461bcd60e51b815260206004820152601a60248201527f50726f6f662073686f756c6420736574206279207365727665720000000000006044820152606401610614565b805161078a90600b906020840190610aa5565b5050600c805460ff19166001179055565b600480546103529061111b565b6001546001600160a01b031633146108125760405162461bcd60e51b815260206004820152602760248201527f496e64657820616e642052616e646f6d47732073686f756c64207365742062796044820152660818db1a595b9d60ca1b6064820152608401610614565b8151610825906008906020850190610b02565b508051610839906006906020840190610b5b565b50506007805460ff1916600117905550565b6001546001600160a01b031633146108a55760405162461bcd60e51b815260206004820152601760248201527f506b2073686f756c642073657420627920636c69656e740000000000000000006044820152606401610614565b81516108b8906004906020850190610b5b565b50805161061d906003906020840190610b5b565b600c5460ff16151560011480156108ea5750600a5460ff1615156001145b6109425760405162461bcd60e51b815260206004820152602360248201527f50726f6f6620616e64206368616c6c656e6765546167206973206e6f7420656d60448201526270747960e81b6064820152608401610614565b600b54600954146109ad5760405162461bcd60e51b815260206004820152602f60248201527f5468652073697a65206f66206368616c6c656e676520616e642070726f6f662060448201526e1cda1bdd5b1908189948195c5d585b608a1b6064820152608401610614565b60005b600954811015610aa25773__$1110a454c665fa152c82a16407592c179f$__632ae0783d600b83815481106109e7576109e7611156565b9060005260206000200160098481548110610a0457610a04611156565b90600052602060002001600360046040518563ffffffff1660e01b8152600401610a31949392919061120c565b602060405180830381865af4158015610a4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a729190611264565b600d819055610a9057600154610a90906001600160a01b0316610575565b80610a9a8161127d565b9150506109b0565b50565b828054828255906000526020600020908101928215610af2579160200282015b82811115610af25782518051610ae2918491602090910190610b5b565b5091602001919060010190610ac5565b50610afe929150610bdb565b5090565b828054828255906000526020600020908101928215610b4f579160200282015b82811115610b4f5782518051610b3f918491602090910190610bf8565b5091602001919060010190610b22565b50610afe929150610c32565b828054610b679061111b565b90600052602060002090601f016020900481019282610b895760008555610bcf565b82601f10610ba257805160ff1916838001178555610bcf565b82800160010185558215610bcf579182015b82811115610bcf578251825591602001919060010190610bb4565b50610afe929150610c4f565b80821115610afe576000610bef8282610c64565b50600101610bdb565b828054828255906000526020600020908101928215610bcf5791602002820182811115610bcf578251825591602001919060010190610bb4565b80821115610afe576000610c468282610c9e565b50600101610c32565b5b80821115610afe5760008155600101610c50565b508054610c709061111b565b6000825580601f10610c80575050565b601f016020900490600052602060002090810190610aa29190610c4f565b5080546000825590600052602060002090810190610aa29190610c4f565b600060208083528351808285015260005b81811015610ce957858101830151858201604001528201610ccd565b81811115610cfb576000604083870101525b50601f01601f1916929092016040019392505050565b60008060408385031215610d2457600080fd5b50508035926020909101359150565b6000602080830181845280855180835260408601915060408160051b87010192508387016000805b83811015610daf57888603603f19018552825180518088529088019088880190845b81811015610d995783518352928a0192918a0191600101610d7d565b5090975050509386019391860191600101610d5b565b509398975050505050505050565b600060208284031215610dcf57600080fd5b5035919050565b600060208284031215610de857600080fd5b81356001600160a01b0381168114610dff57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610e4557610e45610e06565b604052919050565b600067ffffffffffffffff821115610e6757610e67610e06565b5060051b60200190565b600082601f830112610e8257600080fd5b813567ffffffffffffffff811115610e9c57610e9c610e06565b610eaf601f8201601f1916602001610e1c565b818152846020838601011115610ec457600080fd5b816020850160208301376000918101602001919091529392505050565b60006020808385031215610ef457600080fd5b823567ffffffffffffffff80821115610f0c57600080fd5b818501915085601f830112610f2057600080fd5b8135610f33610f2e82610e4d565b610e1c565b81815260059190911b83018401908481019088831115610f5257600080fd5b8585015b83811015610f8a57803585811115610f6e5760008081fd5b610f7c8b89838a0101610e71565b845250918601918601610f56565b5098975050505050505050565b6000806040808486031215610fab57600080fd5b833567ffffffffffffffff80821115610fc357600080fd5b818601915086601f830112610fd757600080fd5b81356020610fe7610f2e83610e4d565b82815260059290921b8401810191818101908a84111561100657600080fd5b8286015b84811015611092578035868111156110225760008081fd5b8701603f81018d136110345760008081fd5b84810135611044610f2e82610e4d565b81815260059190911b82018a0190868101908f8311156110645760008081fd5b928b01925b8284101561108257833582529287019290870190611069565b865250505091830191830161100a565b50975050870135935050808311156110a957600080fd5b50506110b785828601610e71565b9150509250929050565b600080604083850312156110d457600080fd5b823567ffffffffffffffff808211156110ec57600080fd5b6110f886838701610e71565b9350602085013591508082111561110e57600080fd5b506110b785828601610e71565b600181811c9082168061112f57607f821691505b6020821081141561115057634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b8054600090600181811c908083168061118657607f831692505b60208084108214156111a857634e487b7160e01b600052602260045260246000fd5b838852602088018280156111c357600181146111d4576111ff565b60ff198716825282820197506111ff565b60008981526020902060005b878110156111f9578154848201529086019084016111e0565b83019850505b5050505050505092915050565b60808152600061121f608083018761116c565b8281036020840152611231818761116c565b90508281036040840152611245818661116c565b90508281036060840152611259818561116c565b979650505050505050565b60006020828403121561127657600080fd5b5051919050565b600060001982141561129f57634e487b7160e01b600052601160045260246000fd5b506001019056fea26469706673582212200966aeb78fa61665f029965e26c800f34683df73e4d62090cdf56971b250448f64736f6c634300080a0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _client common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _client)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// E is a free data retrieval call binding the contract method 0x92bbf6e8.
//
// Solidity: function E() view returns(bytes)
func (_Api *ApiCaller) E(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "E")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// E is a free data retrieval call binding the contract method 0x92bbf6e8.
//
// Solidity: function E() view returns(bytes)
func (_Api *ApiSession) E() ([]byte, error) {
	return _Api.Contract.E(&_Api.CallOpts)
}

// E is a free data retrieval call binding the contract method 0x92bbf6e8.
//
// Solidity: function E() view returns(bytes)
func (_Api *ApiCallerSession) E() ([]byte, error) {
	return _Api.Contract.E(&_Api.CallOpts)
}

// N is a free data retrieval call binding the contract method 0xc9e525df.
//
// Solidity: function N() view returns(bytes)
func (_Api *ApiCaller) N(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "N")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// N is a free data retrieval call binding the contract method 0xc9e525df.
//
// Solidity: function N() view returns(bytes)
func (_Api *ApiSession) N() ([]byte, error) {
	return _Api.Contract.N(&_Api.CallOpts)
}

// N is a free data retrieval call binding the contract method 0xc9e525df.
//
// Solidity: function N() view returns(bytes)
func (_Api *ApiCallerSession) N() ([]byte, error) {
	return _Api.Contract.N(&_Api.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(uint256)
func (_Api *ApiCaller) Asset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(uint256)
func (_Api *ApiSession) Asset() (*big.Int, error) {
	return _Api.Contract.Asset(&_Api.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(uint256)
func (_Api *ApiCallerSession) Asset() (*big.Int, error) {
	return _Api.Contract.Asset(&_Api.CallOpts)
}

// ChalTags is a free data retrieval call binding the contract method 0x71a76c7e.
//
// Solidity: function chalTags(uint256 ) view returns(bytes)
func (_Api *ApiCaller) ChalTags(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "chalTags", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ChalTags is a free data retrieval call binding the contract method 0x71a76c7e.
//
// Solidity: function chalTags(uint256 ) view returns(bytes)
func (_Api *ApiSession) ChalTags(arg0 *big.Int) ([]byte, error) {
	return _Api.Contract.ChalTags(&_Api.CallOpts, arg0)
}

// ChalTags is a free data retrieval call binding the contract method 0x71a76c7e.
//
// Solidity: function chalTags(uint256 ) view returns(bytes)
func (_Api *ApiCallerSession) ChalTags(arg0 *big.Int) ([]byte, error) {
	return _Api.Contract.ChalTags(&_Api.CallOpts, arg0)
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() view returns(address)
func (_Api *ApiCaller) Client(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "client")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() view returns(address)
func (_Api *ApiSession) Client() (common.Address, error) {
	return _Api.Contract.Client(&_Api.CallOpts)
}

// Client is a free data retrieval call binding the contract method 0x109e94cf.
//
// Solidity: function client() view returns(address)
func (_Api *ApiCallerSession) Client() (common.Address, error) {
	return _Api.Contract.Client(&_Api.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Api *ApiCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Api *ApiSession) GetBalance() (*big.Int, error) {
	return _Api.Contract.GetBalance(&_Api.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Api *ApiCallerSession) GetBalance() (*big.Int, error) {
	return _Api.Contract.GetBalance(&_Api.CallOpts)
}

// GetChIndex is a free data retrieval call binding the contract method 0x22f9c360.
//
// Solidity: function getChIndex() view returns(uint256[][])
func (_Api *ApiCaller) GetChIndex(opts *bind.CallOpts) ([][]*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getChIndex")

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetChIndex is a free data retrieval call binding the contract method 0x22f9c360.
//
// Solidity: function getChIndex() view returns(uint256[][])
func (_Api *ApiSession) GetChIndex() ([][]*big.Int, error) {
	return _Api.Contract.GetChIndex(&_Api.CallOpts)
}

// GetChIndex is a free data retrieval call binding the contract method 0x22f9c360.
//
// Solidity: function getChIndex() view returns(uint256[][])
func (_Api *ApiCallerSession) GetChIndex() ([][]*big.Int, error) {
	return _Api.Contract.GetChIndex(&_Api.CallOpts)
}

// GetGs is a free data retrieval call binding the contract method 0x22e185e7.
//
// Solidity: function getGs() view returns(bytes)
func (_Api *ApiCaller) GetGs(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getGs")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetGs is a free data retrieval call binding the contract method 0x22e185e7.
//
// Solidity: function getGs() view returns(bytes)
func (_Api *ApiSession) GetGs() ([]byte, error) {
	return _Api.Contract.GetGs(&_Api.CallOpts)
}

// GetGs is a free data retrieval call binding the contract method 0x22e185e7.
//
// Solidity: function getGs() view returns(bytes)
func (_Api *ApiCallerSession) GetGs() ([]byte, error) {
	return _Api.Contract.GetGs(&_Api.CallOpts)
}

// GetN is a free data retrieval call binding the contract method 0x3e955225.
//
// Solidity: function getN() view returns(bytes)
func (_Api *ApiCaller) GetN(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getN")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetN is a free data retrieval call binding the contract method 0x3e955225.
//
// Solidity: function getN() view returns(bytes)
func (_Api *ApiSession) GetN() ([]byte, error) {
	return _Api.Contract.GetN(&_Api.CallOpts)
}

// GetN is a free data retrieval call binding the contract method 0x3e955225.
//
// Solidity: function getN() view returns(bytes)
func (_Api *ApiCallerSession) GetN() ([]byte, error) {
	return _Api.Contract.GetN(&_Api.CallOpts)
}

// Gs is a free data retrieval call binding the contract method 0x00e71aaf.
//
// Solidity: function gs() view returns(bytes)
func (_Api *ApiCaller) Gs(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "gs")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Gs is a free data retrieval call binding the contract method 0x00e71aaf.
//
// Solidity: function gs() view returns(bytes)
func (_Api *ApiSession) Gs() ([]byte, error) {
	return _Api.Contract.Gs(&_Api.CallOpts)
}

// Gs is a free data retrieval call binding the contract method 0x00e71aaf.
//
// Solidity: function gs() view returns(bytes)
func (_Api *ApiCallerSession) Gs() ([]byte, error) {
	return _Api.Contract.Gs(&_Api.CallOpts)
}

// Indexs is a free data retrieval call binding the contract method 0x0a36d18b.
//
// Solidity: function indexs(uint256 , uint256 ) view returns(uint256)
func (_Api *ApiCaller) Indexs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "indexs", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Indexs is a free data retrieval call binding the contract method 0x0a36d18b.
//
// Solidity: function indexs(uint256 , uint256 ) view returns(uint256)
func (_Api *ApiSession) Indexs(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Api.Contract.Indexs(&_Api.CallOpts, arg0, arg1)
}

// Indexs is a free data retrieval call binding the contract method 0x0a36d18b.
//
// Solidity: function indexs(uint256 , uint256 ) view returns(uint256)
func (_Api *ApiCallerSession) Indexs(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Api.Contract.Indexs(&_Api.CallOpts, arg0, arg1)
}

// Proofs is a free data retrieval call binding the contract method 0x9ddaf5aa.
//
// Solidity: function proofs(uint256 ) view returns(bytes)
func (_Api *ApiCaller) Proofs(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "proofs", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Proofs is a free data retrieval call binding the contract method 0x9ddaf5aa.
//
// Solidity: function proofs(uint256 ) view returns(bytes)
func (_Api *ApiSession) Proofs(arg0 *big.Int) ([]byte, error) {
	return _Api.Contract.Proofs(&_Api.CallOpts, arg0)
}

// Proofs is a free data retrieval call binding the contract method 0x9ddaf5aa.
//
// Solidity: function proofs(uint256 ) view returns(bytes)
func (_Api *ApiCallerSession) Proofs(arg0 *big.Int) ([]byte, error) {
	return _Api.Contract.Proofs(&_Api.CallOpts, arg0)
}

// Server is a free data retrieval call binding the contract method 0xfd922a42.
//
// Solidity: function server() view returns(address)
func (_Api *ApiCaller) Server(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "server")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Server is a free data retrieval call binding the contract method 0xfd922a42.
//
// Solidity: function server() view returns(address)
func (_Api *ApiSession) Server() (common.Address, error) {
	return _Api.Contract.Server(&_Api.CallOpts)
}

// Server is a free data retrieval call binding the contract method 0xfd922a42.
//
// Solidity: function server() view returns(address)
func (_Api *ApiCallerSession) Server() (common.Address, error) {
	return _Api.Contract.Server(&_Api.CallOpts)
}

// SendViaCall is a paid mutator transaction binding the contract method 0x830c29ae.
//
// Solidity: function sendViaCall(address _to) payable returns()
func (_Api *ApiTransactor) SendViaCall(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "sendViaCall", _to)
}

// SendViaCall is a paid mutator transaction binding the contract method 0x830c29ae.
//
// Solidity: function sendViaCall(address _to) payable returns()
func (_Api *ApiSession) SendViaCall(_to common.Address) (*types.Transaction, error) {
	return _Api.Contract.SendViaCall(&_Api.TransactOpts, _to)
}

// SendViaCall is a paid mutator transaction binding the contract method 0x830c29ae.
//
// Solidity: function sendViaCall(address _to) payable returns()
func (_Api *ApiTransactorSession) SendViaCall(_to common.Address) (*types.Transaction, error) {
	return _Api.Contract.SendViaCall(&_Api.TransactOpts, _to)
}

// SetChals is a paid mutator transaction binding the contract method 0x8947157e.
//
// Solidity: function setChals(bytes[] _chaltags) returns()
func (_Api *ApiTransactor) SetChals(opts *bind.TransactOpts, _chaltags [][]byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setChals", _chaltags)
}

// SetChals is a paid mutator transaction binding the contract method 0x8947157e.
//
// Solidity: function setChals(bytes[] _chaltags) returns()
func (_Api *ApiSession) SetChals(_chaltags [][]byte) (*types.Transaction, error) {
	return _Api.Contract.SetChals(&_Api.TransactOpts, _chaltags)
}

// SetChals is a paid mutator transaction binding the contract method 0x8947157e.
//
// Solidity: function setChals(bytes[] _chaltags) returns()
func (_Api *ApiTransactorSession) SetChals(_chaltags [][]byte) (*types.Transaction, error) {
	return _Api.Contract.SetChals(&_Api.TransactOpts, _chaltags)
}

// SetIndexAndGs is a paid mutator transaction binding the contract method 0xdf5d093e.
//
// Solidity: function setIndexAndGs(uint256[][] _index, bytes _gs) returns()
func (_Api *ApiTransactor) SetIndexAndGs(opts *bind.TransactOpts, _index [][]*big.Int, _gs []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setIndexAndGs", _index, _gs)
}

// SetIndexAndGs is a paid mutator transaction binding the contract method 0xdf5d093e.
//
// Solidity: function setIndexAndGs(uint256[][] _index, bytes _gs) returns()
func (_Api *ApiSession) SetIndexAndGs(_index [][]*big.Int, _gs []byte) (*types.Transaction, error) {
	return _Api.Contract.SetIndexAndGs(&_Api.TransactOpts, _index, _gs)
}

// SetIndexAndGs is a paid mutator transaction binding the contract method 0xdf5d093e.
//
// Solidity: function setIndexAndGs(uint256[][] _index, bytes _gs) returns()
func (_Api *ApiTransactorSession) SetIndexAndGs(_index [][]*big.Int, _gs []byte) (*types.Transaction, error) {
	return _Api.Contract.SetIndexAndGs(&_Api.TransactOpts, _index, _gs)
}

// SetPk is a paid mutator transaction binding the contract method 0xe258eb7d.
//
// Solidity: function setPk(bytes _N, bytes _E) returns()
func (_Api *ApiTransactor) SetPk(opts *bind.TransactOpts, _N []byte, _E []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setPk", _N, _E)
}

// SetPk is a paid mutator transaction binding the contract method 0xe258eb7d.
//
// Solidity: function setPk(bytes _N, bytes _E) returns()
func (_Api *ApiSession) SetPk(_N []byte, _E []byte) (*types.Transaction, error) {
	return _Api.Contract.SetPk(&_Api.TransactOpts, _N, _E)
}

// SetPk is a paid mutator transaction binding the contract method 0xe258eb7d.
//
// Solidity: function setPk(bytes _N, bytes _E) returns()
func (_Api *ApiTransactorSession) SetPk(_N []byte, _E []byte) (*types.Transaction, error) {
	return _Api.Contract.SetPk(&_Api.TransactOpts, _N, _E)
}

// SetProve is a paid mutator transaction binding the contract method 0xa960f784.
//
// Solidity: function setProve(bytes[] _proof) returns()
func (_Api *ApiTransactor) SetProve(opts *bind.TransactOpts, _proof [][]byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setProve", _proof)
}

// SetProve is a paid mutator transaction binding the contract method 0xa960f784.
//
// Solidity: function setProve(bytes[] _proof) returns()
func (_Api *ApiSession) SetProve(_proof [][]byte) (*types.Transaction, error) {
	return _Api.Contract.SetProve(&_Api.TransactOpts, _proof)
}

// SetProve is a paid mutator transaction binding the contract method 0xa960f784.
//
// Solidity: function setProve(bytes[] _proof) returns()
func (_Api *ApiTransactorSession) SetProve(_proof [][]byte) (*types.Transaction, error) {
	return _Api.Contract.SetProve(&_Api.TransactOpts, _proof)
}

// Verify is a paid mutator transaction binding the contract method 0xfc735e99.
//
// Solidity: function verify() payable returns()
func (_Api *ApiTransactor) Verify(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "verify")
}

// Verify is a paid mutator transaction binding the contract method 0xfc735e99.
//
// Solidity: function verify() payable returns()
func (_Api *ApiSession) Verify() (*types.Transaction, error) {
	return _Api.Contract.Verify(&_Api.TransactOpts)
}

// Verify is a paid mutator transaction binding the contract method 0xfc735e99.
//
// Solidity: function verify() payable returns()
func (_Api *ApiTransactorSession) Verify() (*types.Transaction, error) {
	return _Api.Contract.Verify(&_Api.TransactOpts)
}
