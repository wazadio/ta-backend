package smartcontract

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"signature-app/database/repository"
	"signature-app/domain"
	requestdomain "signature-app/domain/request_domain"
	responsedomain "signature-app/domain/response_domain"
	"signature-app/helper"
	data "signature-app/service/smart_contract/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type caller struct {
	cl              *ethclient.Client
	ctx             context.Context
	contractAddress string
	db              *repository.Database
}

func NewCaller(cl *ethclient.Client, ctx context.Context, contractAddress string) *caller {
	return &caller{
		cl:              cl,
		ctx:             ctx,
		contractAddress: contractAddress,
	}
}

func (c *caller) Deploy(privateKey string) (contractDetail responsedomain.DeployContractResponse, err error) {
	chainId, err := c.cl.ChainID(c.ctx)
	if err != nil {
		log.Println(err)
		return contractDetail, err
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Println(err)
		return contractDetail, err
	}

	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println(err)
		return contractDetail, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.cl.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return contractDetail, err
	}

	gasPrice, err := c.cl.SuggestGasPrice(c.ctx)
	if err != nil {
		log.Println(err)
		return contractDetail, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainId)
	if err != nil {
		log.Println(err)
		return contractDetail, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, _, err := data.DeployData(auth, c.cl)
	if err != nil {
		log.Println(err)
		return contractDetail, err
	}

	from, err := types.Sender(types.NewEIP155Signer(chainId), tx)
	if err != nil {
		log.Println(err)
		return contractDetail, err
	}

	fmt.Println("adress")
	fmt.Println(address.Hex())

	c.contractAddress = address.Hex()

	return responsedomain.DeployContractResponse{
		Id:              tx.Hash().Hex(),
		ContractAddress: address.Hex(),
		From:            from.Hex(),
	}, nil
}

func (c *caller) GetAdmin() (adminAddress string, err error) {
	address := common.HexToAddress(c.contractAddress)
	instance, err := data.NewData(address, c.cl)
	if err != nil {
		return adminAddress, err
	}

	fmt.Println("contract is loaded")

	res, err := instance.Admin(nil)
	if err != nil {
		return adminAddress, err
	}

	adminAddress = res.Hex()

	return
}

func (c *caller) GetDokumen() (dokumen responsedomain.GetDokumenResponse, err error) {
	address := common.HexToAddress(c.contractAddress)
	instance, err := data.NewData(address, c.cl)
	if err != nil {
		return dokumen, err
	}

	fmt.Println("contract is loaded")

	dokumen.Dokumen, err = instance.GetDokumens(nil)
	if err != nil {
		return dokumen, err
	}

	return
}

func (c *caller) GetIdentitas() (identitas responsedomain.GetIdentitasResponse, err error) {
	address := common.HexToAddress(c.contractAddress)
	instance, err := data.NewData(address, c.cl)
	if err != nil {
		return identitas, err
	}

	fmt.Println("contract is loaded")

	identitas.Identitas, err = instance.GetIdentitas(nil)
	if err != nil {
		return identitas, err
	}

	return
}

func (c *caller) AddDokumen(namaDokumen, privateKey string) (isSuccess responsedomain.AddDokumenResponse, err error) {
	address := common.HexToAddress(c.contractAddress)
	instance, err := data.NewData(address, c.cl)
	if err != nil {
		return isSuccess, err
	}

	fmt.Println("contract is loaded")

	chainId, err := c.cl.ChainID(c.ctx)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println(err)
		return isSuccess, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.cl.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	gasPrice, err := c.cl.SuggestGasPrice(c.ctx)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainId)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	_, err = instance.AddDokumen(auth, namaDokumen)
	if err != nil {
		return isSuccess, err
	}

	isSuccess.IsSuccess = true

	return
}

func (c *caller) AddIdentitas(payload requestdomain.AddIdentitasRequest) (isSuccess responsedomain.AddIdentitasResponse, err error) {
	address := common.HexToAddress(c.contractAddress)
	instance, err := data.NewData(address, c.cl)
	if err != nil {
		return isSuccess, err
	}

	chainId, err := c.cl.ChainID(c.ctx)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(payload.PrivateKey)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println(err)
		return isSuccess, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.cl.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	adminAddress, err := c.GetAdmin()
	if err != nil {
		return isSuccess, err
	}

	if adminAddress != fromAddress.Hex() {
		return isSuccess, errors.New("not admin")
	}

	gasPrice, err := c.cl.SuggestGasPrice(c.ctx)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainId)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	_, err = instance.AddIdentitas(auth, payload.Alamat, big.NewInt(int64(payload.Nomor)), payload.Nama, payload.Status)
	if err != nil {
		return isSuccess, err
	}

	if payload.Status == os.Getenv("PARTY") {
		updateErr := c.db.UpdateDeviceTokenStatus(payload.Alamat)
		if updateErr != nil {
			log.Println("error update token status : ", updateErr)
		}
	}

	isSuccess.IsSuccess = true

	return
}

func (c *caller) GetETH(privateKey string) (isSuccess domain.BaseField, err error) {
	address := common.HexToAddress(c.contractAddress)
	instance, err := data.NewData(address, c.cl)
	if err != nil {
		return isSuccess, err
	}

	chainId, err := c.cl.ChainID(c.ctx)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println(err)
		return isSuccess, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.cl.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	help := helper.NewHelper(c.cl, c.ctx)
	balance, err := help.CheckBalance(fromAddress.Hex())
	if err != nil {
		return isSuccess, err
	}

	b, _ := balance.Float64()
	if b > float64(100) {
		return isSuccess, errors.New("high balance")
	}

	gasPrice, err := c.cl.SuggestGasPrice(c.ctx)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainId)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	_, err = instance.GetETH(auth)
	if err != nil {
		return isSuccess, err
	}

	isSuccess.IsSuccess = true

	return
}

func (c *caller) DeleteDokumen(namaDokumen, privateKey string) (isSuccess responsedomain.AddDokumenResponse, err error) {
	address := common.HexToAddress(c.contractAddress)
	instance, err := data.NewData(address, c.cl)
	if err != nil {
		return isSuccess, err
	}

	chainId, err := c.cl.ChainID(c.ctx)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println(err)
		return isSuccess, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.cl.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	gasPrice, err := c.cl.SuggestGasPrice(c.ctx)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainId)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	_, err = instance.DeleteDokumen(auth, namaDokumen)
	if err != nil {
		return isSuccess, err
	}

	isSuccess.IsSuccess = true

	return
}

func (c *caller) DeleteIdentitas(alamat, privateKey string) (isSuccess responsedomain.AddDokumenResponse, err error) {
	address := common.HexToAddress(c.contractAddress)
	instance, err := data.NewData(address, c.cl)
	if err != nil {
		return isSuccess, err
	}

	chainId, err := c.cl.ChainID(c.ctx)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	publicKey := ecdsaPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println(err)
		return isSuccess, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.cl.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	gasPrice, err := c.cl.SuggestGasPrice(c.ctx)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainId)
	if err != nil {
		log.Println(err)
		return isSuccess, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	_, err = instance.DeleteIdentitas(auth, alamat)
	if err != nil {
		return isSuccess, err
	}

	isSuccess.IsSuccess = true

	return
}
