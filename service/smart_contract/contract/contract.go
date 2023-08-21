// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package data

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
	_ = abi.ConvertType
)

// DataIdentitas is an auto generated low-level Go binding around an user-defined struct.
type DataIdentitas struct {
	Alamat string
	Nomor  *big.Int
	Nama   string
	Status string
}

// DataMetaData contains all meta data concerning the Data contract.
var DataMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dokumen\",\"type\":\"string\"}],\"name\":\"addDokumen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_alamat\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_nomor\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_nama\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_status\",\"type\":\"string\"}],\"name\":\"addIdentitas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_namaDokumen\",\"type\":\"string\"}],\"name\":\"deleteDokumen\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_alamat\",\"type\":\"string\"}],\"name\":\"deleteIdentitas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donateContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDokumens\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIdentitas\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"alamat\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nomor\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nama\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"internalType\":\"structData.Identitas[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611808806100616000396000f3fe6080604052600436106100915760003560e01c80639a2682bd116100595780639a2682bd1461012b578063c824f13f14610154578063d431d8991461017d578063f851a440146101a6578063fca2817e146101d157610091565b806314f6c3be1461009657806353e4bb84146100a0578063569f9b77146100cb57806362964936146100d55780638b7afe2e14610100575b600080fd5b61009e6101fa565b005b3480156100ac57600080fd5b506100b5610278565b6040516100c29190610c5e565b60405180910390f35b6100d3610351565b005b3480156100e157600080fd5b506100ea610353565b6040516100f79190610dcc565b60405180910390f35b34801561010c57600080fd5b50610115610572565b6040516101229190610dfd565b60405180910390f35b34801561013757600080fd5b50610152600480360381019061014d9190610f8d565b61057a565b005b34801561016057600080fd5b5061017b60048036038101906101769190611048565b610674565b005b34801561018957600080fd5b506101a4600480360381019061019f9190611048565b6107f4565b005b3480156101b257600080fd5b506101bb6109f7565b6040516101c891906110d2565b60405180910390f35b3480156101dd57600080fd5b506101f860048036038101906101f39190611048565b610a1d565b005b68056bc75e2d631000003373ffffffffffffffffffffffffffffffffffffffff16311061022657600080fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc68056bc75e2d631000009081150290604051600060405180830381858888f19350505050158015610275573d6000803e3d6000fd5b50565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156103485783829060005260206000200180546102bb9061111c565b80601f01602080910402602001604051908101604052809291908181526020018280546102e79061111c565b80156103345780601f1061030957610100808354040283529160200191610334565b820191906000526020600020905b81548152906001019060200180831161031757829003601f168201915b50505050508152602001906001019061029c565b50505050905090565b565b60606000805480602002602001604051908101604052809291908181526020016000905b8282101561056957838290600052602060002090600402016040518060800160405290816000820180546103aa9061111c565b80601f01602080910402602001604051908101604052809291908181526020018280546103d69061111c565b80156104235780601f106103f857610100808354040283529160200191610423565b820191906000526020600020905b81548152906001019060200180831161040657829003601f168201915b50505050508152602001600182015481526020016002820180546104469061111c565b80601f01602080910402602001604051908101604052809291908181526020018280546104729061111c565b80156104bf5780601f10610494576101008083540402835291602001916104bf565b820191906000526020600020905b8154815290600101906020018083116104a257829003601f168201915b505050505081526020016003820180546104d89061111c565b80601f01602080910402602001604051908101604052809291908181526020018280546105049061111c565b80156105515780601f1061052657610100808354040283529160200191610551565b820191906000526020600020905b81548152906001019060200180831161053457829003601f168201915b50505050508152505081526020019060010190610377565b50505050905090565b600047905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105d457600080fd5b60006040518060800160405280868152602001858152602001848152602001838152509080600181540180825580915050600190039060005260206000209060040201600090919091909150600082015181600001908161063591906112f9565b5060208201518160010155604082015181600201908161065591906112f9565b50606082015181600301908161066b91906112f9565b50505050505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ce57600080fd5b60005b6001805490508110156107ef57816040516020016106ef9190611407565b60405160208183030381529060405280519060200120600182815481106107195761071861141e565b5b9060005260206000200160405160200161073391906114d0565b60405160208183030381529060405280519060200120036107dc5760018080805490506107609190611516565b815481106107715761077061141e565b5b906000526020600020016001828154811061078f5761078e61141e565b5b9060005260206000200190816107a59190611575565b5060018054806107b8576107b761165d565b5b6001900381819060005260206000200160006107d49190610aaf565b9055506107f1565b80806107e79061168c565b9150506106d1565b505b50565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461084e57600080fd5b60005b6000805490508110156109f2578160405160200161086f9190611407565b60405160208183030381529060405280519060200120600082815481106108995761089861141e565b5b90600052602060002090600402016000016040516020016108ba91906114d0565b60405160208183030381529060405280519060200120036109df57600060016000805490506108e99190611516565b815481106108fa576108f961141e565b5b90600052602060002090600402016000828154811061091c5761091b61141e565b5b90600052602060002090600402016000820181600001908161093e91906116ea565b50600182015481600101556002820181600201908161095d91906116ea565b506003820181600301908161097291906116ea565b5090505060008054806109885761098761165d565b5b6001900381819060005260206000209060040201600080820160006109ad9190610aaf565b60018201600090556002820160006109c59190610aaf565b6003820160006109d59190610aaf565b50509055506109f4565b80806109ea9061168c565b915050610851565b505b50565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a7757600080fd5b600181908060018154018082558091505060019003906000526020600020016000909190919091509081610aab91906112f9565b5050565b508054610abb9061111c565b6000825580601f10610acd5750610aec565b601f016020900490600052602060002090810190610aeb9190610aef565b5b50565b5b80821115610b08576000816000905550600101610af0565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610b72578082015181840152602081019050610b57565b60008484015250505050565b6000601f19601f8301169050919050565b6000610b9a82610b38565b610ba48185610b43565b9350610bb4818560208601610b54565b610bbd81610b7e565b840191505092915050565b6000610bd48383610b8f565b905092915050565b6000602082019050919050565b6000610bf482610b0c565b610bfe8185610b17565b935083602082028501610c1085610b28565b8060005b85811015610c4c5784840389528151610c2d8582610bc8565b9450610c3883610bdc565b925060208a01995050600181019050610c14565b50829750879550505050505092915050565b60006020820190508181036000830152610c788184610be9565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b610cbf81610cac565b82525050565b60006080830160008301518482036000860152610ce28282610b8f565b9150506020830151610cf76020860182610cb6565b5060408301518482036040860152610d0f8282610b8f565b91505060608301518482036060860152610d298282610b8f565b9150508091505092915050565b6000610d428383610cc5565b905092915050565b6000602082019050919050565b6000610d6282610c80565b610d6c8185610c8b565b935083602082028501610d7e85610c9c565b8060005b85811015610dba5784840389528151610d9b8582610d36565b9450610da683610d4a565b925060208a01995050600181019050610d82565b50829750879550505050505092915050565b60006020820190508181036000830152610de68184610d57565b905092915050565b610df781610cac565b82525050565b6000602082019050610e126000830184610dee565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610e6e82610b7e565b810181811067ffffffffffffffff82111715610e8d57610e8c610e36565b5b80604052505050565b6000610ea0610e18565b9050610eac8282610e65565b919050565b600067ffffffffffffffff821115610ecc57610ecb610e36565b5b610ed582610b7e565b9050602081019050919050565b82818337600083830152505050565b6000610f04610eff84610eb1565b610e96565b905082815260208101848484011115610f2057610f1f610e31565b5b610f2b848285610ee2565b509392505050565b600082601f830112610f4857610f47610e2c565b5b8135610f58848260208601610ef1565b91505092915050565b610f6a81610cac565b8114610f7557600080fd5b50565b600081359050610f8781610f61565b92915050565b60008060008060808587031215610fa757610fa6610e22565b5b600085013567ffffffffffffffff811115610fc557610fc4610e27565b5b610fd187828801610f33565b9450506020610fe287828801610f78565b935050604085013567ffffffffffffffff81111561100357611002610e27565b5b61100f87828801610f33565b925050606085013567ffffffffffffffff8111156110305761102f610e27565b5b61103c87828801610f33565b91505092959194509250565b60006020828403121561105e5761105d610e22565b5b600082013567ffffffffffffffff81111561107c5761107b610e27565b5b61108884828501610f33565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006110bc82611091565b9050919050565b6110cc816110b1565b82525050565b60006020820190506110e760008301846110c3565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061113457607f821691505b602082108103611147576111466110ed565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026111af7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611172565b6111b98683611172565b95508019841693508086168417925050509392505050565b6000819050919050565b60006111f66111f16111ec84610cac565b6111d1565b610cac565b9050919050565b6000819050919050565b611210836111db565b61122461121c826111fd565b84845461117f565b825550505050565b600090565b61123961122c565b611244818484611207565b505050565b5b818110156112685761125d600082611231565b60018101905061124a565b5050565b601f8211156112ad5761127e8161114d565b61128784611162565b81016020851015611296578190505b6112aa6112a285611162565b830182611249565b50505b505050565b600082821c905092915050565b60006112d0600019846008026112b2565b1980831691505092915050565b60006112e983836112bf565b9150826002028217905092915050565b61130282610b38565b67ffffffffffffffff81111561131b5761131a610e36565b5b611325825461111c565b61133082828561126c565b600060209050601f8311600181146113635760008415611351578287015190505b61135b85826112dd565b8655506113c3565b601f1984166113718661114d565b60005b8281101561139957848901518255600182019150602085019450602081019050611374565b868310156113b657848901516113b2601f8916826112bf565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b60006113e182610b38565b6113eb81856113cb565b93506113fb818560208601610b54565b80840191505092915050565b600061141382846113d6565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000815461145a8161111c565b61146481866113cb565b9450600182166000811461147f5760018114611494576114c7565b60ff19831686528115158202860193506114c7565b61149d8561114d565b60005b838110156114bf578154818901526001820191506020810190506114a0565b838801955050505b50505092915050565b60006114dc828461144d565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061152182610cac565b915061152c83610cac565b9250828203905081811115611544576115436114e7565b5b92915050565b6000815490506115598161111c565b9050919050565b60008190508160005260206000209050919050565b81810361158357505061165b565b61158c8261154a565b67ffffffffffffffff8111156115a5576115a4610e36565b5b6115af825461111c565b6115ba82828561126c565b6000601f8311600181146115e957600084156115d7578287015490505b6115e185826112dd565b865550611654565b601f1984166115f787611560565b96506116028661114d565b60005b8281101561162a57848901548255600182019150600185019450602081019050611605565b868310156116475784890154611643601f8916826112bf565b8355505b6001600288020188555050505b5050505050505b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600061169782610cac565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036116c9576116c86114e7565b5b600182019050919050565b6000815490506116e38161111c565b9050919050565b8181036116f85750506117d0565b611701826116d4565b67ffffffffffffffff81111561171a57611719610e36565b5b611724825461111c565b61172f82828561126c565b6000601f83116001811461175e576000841561174c578287015490505b61175685826112dd565b8655506117c9565b601f19841661176c8761114d565b96506117778661114d565b60005b8281101561179f5784890154825560018201915060018501945060208101905061177a565b868310156117bc57848901546117b8601f8916826112bf565b8355505b6001600288020188555050505b5050505050505b56fea2646970667358221220530076d18b04cb05cf57348c9cfcd98c6bbee4ff6220981de1665453f42d97dd64736f6c63430008120033",
}

// DataABI is the input ABI used to generate the binding from.
// Deprecated: Use DataMetaData.ABI instead.
var DataABI = DataMetaData.ABI

// DataBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DataMetaData.Bin instead.
var DataBin = DataMetaData.Bin

// DeployData deploys a new Ethereum contract, binding an instance of Data to it.
func DeployData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Data, error) {
	parsed, err := DataMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Data{DataCaller: DataCaller{contract: contract}, DataTransactor: DataTransactor{contract: contract}, DataFilterer: DataFilterer{contract: contract}}, nil
}

// Data is an auto generated Go binding around an Ethereum contract.
type Data struct {
	DataCaller     // Read-only binding to the contract
	DataTransactor // Write-only binding to the contract
	DataFilterer   // Log filterer for contract events
}

// DataCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataSession struct {
	Contract     *Data             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataCallerSession struct {
	Contract *DataCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataTransactorSession struct {
	Contract     *DataTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataRaw struct {
	Contract *Data // Generic contract binding to access the raw methods on
}

// DataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataCallerRaw struct {
	Contract *DataCaller // Generic read-only contract binding to access the raw methods on
}

// DataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataTransactorRaw struct {
	Contract *DataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewData creates a new instance of Data, bound to a specific deployed contract.
func NewData(address common.Address, backend bind.ContractBackend) (*Data, error) {
	contract, err := bindData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Data{DataCaller: DataCaller{contract: contract}, DataTransactor: DataTransactor{contract: contract}, DataFilterer: DataFilterer{contract: contract}}, nil
}

// NewDataCaller creates a new read-only instance of Data, bound to a specific deployed contract.
func NewDataCaller(address common.Address, caller bind.ContractCaller) (*DataCaller, error) {
	contract, err := bindData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataCaller{contract: contract}, nil
}

// NewDataTransactor creates a new write-only instance of Data, bound to a specific deployed contract.
func NewDataTransactor(address common.Address, transactor bind.ContractTransactor) (*DataTransactor, error) {
	contract, err := bindData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataTransactor{contract: contract}, nil
}

// NewDataFilterer creates a new log filterer instance of Data, bound to a specific deployed contract.
func NewDataFilterer(address common.Address, filterer bind.ContractFilterer) (*DataFilterer, error) {
	contract, err := bindData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataFilterer{contract: contract}, nil
}

// bindData binds a generic wrapper to an already deployed contract.
func bindData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Data *DataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Data.Contract.DataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Data *DataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Data.Contract.DataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Data *DataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Data.Contract.DataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Data *DataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Data.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Data *DataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Data.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Data *DataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Data.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Data *DataCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Data *DataSession) Admin() (common.Address, error) {
	return _Data.Contract.Admin(&_Data.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Data *DataCallerSession) Admin() (common.Address, error) {
	return _Data.Contract.Admin(&_Data.CallOpts)
}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_Data *DataCaller) ContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "contractBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_Data *DataSession) ContractBalance() (*big.Int, error) {
	return _Data.Contract.ContractBalance(&_Data.CallOpts)
}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_Data *DataCallerSession) ContractBalance() (*big.Int, error) {
	return _Data.Contract.ContractBalance(&_Data.CallOpts)
}

// GetDokumens is a free data retrieval call binding the contract method 0x53e4bb84.
//
// Solidity: function getDokumens() view returns(string[])
func (_Data *DataCaller) GetDokumens(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getDokumens")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetDokumens is a free data retrieval call binding the contract method 0x53e4bb84.
//
// Solidity: function getDokumens() view returns(string[])
func (_Data *DataSession) GetDokumens() ([]string, error) {
	return _Data.Contract.GetDokumens(&_Data.CallOpts)
}

// GetDokumens is a free data retrieval call binding the contract method 0x53e4bb84.
//
// Solidity: function getDokumens() view returns(string[])
func (_Data *DataCallerSession) GetDokumens() ([]string, error) {
	return _Data.Contract.GetDokumens(&_Data.CallOpts)
}

// GetIdentitas is a free data retrieval call binding the contract method 0x62964936.
//
// Solidity: function getIdentitas() view returns((string,uint256,string,string)[])
func (_Data *DataCaller) GetIdentitas(opts *bind.CallOpts) ([]DataIdentitas, error) {
	var out []interface{}
	err := _Data.contract.Call(opts, &out, "getIdentitas")

	if err != nil {
		return *new([]DataIdentitas), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataIdentitas)).(*[]DataIdentitas)

	return out0, err

}

// GetIdentitas is a free data retrieval call binding the contract method 0x62964936.
//
// Solidity: function getIdentitas() view returns((string,uint256,string,string)[])
func (_Data *DataSession) GetIdentitas() ([]DataIdentitas, error) {
	return _Data.Contract.GetIdentitas(&_Data.CallOpts)
}

// GetIdentitas is a free data retrieval call binding the contract method 0x62964936.
//
// Solidity: function getIdentitas() view returns((string,uint256,string,string)[])
func (_Data *DataCallerSession) GetIdentitas() ([]DataIdentitas, error) {
	return _Data.Contract.GetIdentitas(&_Data.CallOpts)
}

// AddDokumen is a paid mutator transaction binding the contract method 0xfca2817e.
//
// Solidity: function addDokumen(string dokumen) returns()
func (_Data *DataTransactor) AddDokumen(opts *bind.TransactOpts, dokumen string) (*types.Transaction, error) {
	return _Data.contract.Transact(opts, "addDokumen", dokumen)
}

// AddDokumen is a paid mutator transaction binding the contract method 0xfca2817e.
//
// Solidity: function addDokumen(string dokumen) returns()
func (_Data *DataSession) AddDokumen(dokumen string) (*types.Transaction, error) {
	return _Data.Contract.AddDokumen(&_Data.TransactOpts, dokumen)
}

// AddDokumen is a paid mutator transaction binding the contract method 0xfca2817e.
//
// Solidity: function addDokumen(string dokumen) returns()
func (_Data *DataTransactorSession) AddDokumen(dokumen string) (*types.Transaction, error) {
	return _Data.Contract.AddDokumen(&_Data.TransactOpts, dokumen)
}

// AddIdentitas is a paid mutator transaction binding the contract method 0x9a2682bd.
//
// Solidity: function addIdentitas(string _alamat, uint256 _nomor, string _nama, string _status) returns()
func (_Data *DataTransactor) AddIdentitas(opts *bind.TransactOpts, _alamat string, _nomor *big.Int, _nama string, _status string) (*types.Transaction, error) {
	return _Data.contract.Transact(opts, "addIdentitas", _alamat, _nomor, _nama, _status)
}

// AddIdentitas is a paid mutator transaction binding the contract method 0x9a2682bd.
//
// Solidity: function addIdentitas(string _alamat, uint256 _nomor, string _nama, string _status) returns()
func (_Data *DataSession) AddIdentitas(_alamat string, _nomor *big.Int, _nama string, _status string) (*types.Transaction, error) {
	return _Data.Contract.AddIdentitas(&_Data.TransactOpts, _alamat, _nomor, _nama, _status)
}

// AddIdentitas is a paid mutator transaction binding the contract method 0x9a2682bd.
//
// Solidity: function addIdentitas(string _alamat, uint256 _nomor, string _nama, string _status) returns()
func (_Data *DataTransactorSession) AddIdentitas(_alamat string, _nomor *big.Int, _nama string, _status string) (*types.Transaction, error) {
	return _Data.Contract.AddIdentitas(&_Data.TransactOpts, _alamat, _nomor, _nama, _status)
}

// DeleteDokumen is a paid mutator transaction binding the contract method 0xc824f13f.
//
// Solidity: function deleteDokumen(string _namaDokumen) returns()
func (_Data *DataTransactor) DeleteDokumen(opts *bind.TransactOpts, _namaDokumen string) (*types.Transaction, error) {
	return _Data.contract.Transact(opts, "deleteDokumen", _namaDokumen)
}

// DeleteDokumen is a paid mutator transaction binding the contract method 0xc824f13f.
//
// Solidity: function deleteDokumen(string _namaDokumen) returns()
func (_Data *DataSession) DeleteDokumen(_namaDokumen string) (*types.Transaction, error) {
	return _Data.Contract.DeleteDokumen(&_Data.TransactOpts, _namaDokumen)
}

// DeleteDokumen is a paid mutator transaction binding the contract method 0xc824f13f.
//
// Solidity: function deleteDokumen(string _namaDokumen) returns()
func (_Data *DataTransactorSession) DeleteDokumen(_namaDokumen string) (*types.Transaction, error) {
	return _Data.Contract.DeleteDokumen(&_Data.TransactOpts, _namaDokumen)
}

// DeleteIdentitas is a paid mutator transaction binding the contract method 0xd431d899.
//
// Solidity: function deleteIdentitas(string _alamat) returns()
func (_Data *DataTransactor) DeleteIdentitas(opts *bind.TransactOpts, _alamat string) (*types.Transaction, error) {
	return _Data.contract.Transact(opts, "deleteIdentitas", _alamat)
}

// DeleteIdentitas is a paid mutator transaction binding the contract method 0xd431d899.
//
// Solidity: function deleteIdentitas(string _alamat) returns()
func (_Data *DataSession) DeleteIdentitas(_alamat string) (*types.Transaction, error) {
	return _Data.Contract.DeleteIdentitas(&_Data.TransactOpts, _alamat)
}

// DeleteIdentitas is a paid mutator transaction binding the contract method 0xd431d899.
//
// Solidity: function deleteIdentitas(string _alamat) returns()
func (_Data *DataTransactorSession) DeleteIdentitas(_alamat string) (*types.Transaction, error) {
	return _Data.Contract.DeleteIdentitas(&_Data.TransactOpts, _alamat)
}

// DonateContract is a paid mutator transaction binding the contract method 0x569f9b77.
//
// Solidity: function donateContract() payable returns()
func (_Data *DataTransactor) DonateContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Data.contract.Transact(opts, "donateContract")
}

// DonateContract is a paid mutator transaction binding the contract method 0x569f9b77.
//
// Solidity: function donateContract() payable returns()
func (_Data *DataSession) DonateContract() (*types.Transaction, error) {
	return _Data.Contract.DonateContract(&_Data.TransactOpts)
}

// DonateContract is a paid mutator transaction binding the contract method 0x569f9b77.
//
// Solidity: function donateContract() payable returns()
func (_Data *DataTransactorSession) DonateContract() (*types.Transaction, error) {
	return _Data.Contract.DonateContract(&_Data.TransactOpts)
}

// GetETH is a paid mutator transaction binding the contract method 0x14f6c3be.
//
// Solidity: function getETH() payable returns()
func (_Data *DataTransactor) GetETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Data.contract.Transact(opts, "getETH")
}

// GetETH is a paid mutator transaction binding the contract method 0x14f6c3be.
//
// Solidity: function getETH() payable returns()
func (_Data *DataSession) GetETH() (*types.Transaction, error) {
	return _Data.Contract.GetETH(&_Data.TransactOpts)
}

// GetETH is a paid mutator transaction binding the contract method 0x14f6c3be.
//
// Solidity: function getETH() payable returns()
func (_Data *DataTransactorSession) GetETH() (*types.Transaction, error) {
	return _Data.Contract.GetETH(&_Data.TransactOpts)
}
