package helper

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"signature-app/database/model"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type helper struct {
	cl  *ethclient.Client
	ctx context.Context
}

func NewHelper(cl *ethclient.Client, ctx context.Context) *helper {
	return &helper{
		cl:  cl,
		ctx: ctx,
	}
}

func (h *helper) CheckBalance(address string) (*big.Float, error) {
	account := common.HexToAddress(address)
	balance, err := h.cl.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return ethValue, err
}

func (h *helper) PostRequestAsk(url string, data model.TransactionModel) error {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("id", data.Id)
	_ = writer.WriteField("from_address", data.FromAddress)
	_ = writer.WriteField("from_name", data.FromName)
	_ = writer.WriteField("to_address", data.ToAddress)
	_ = writer.WriteField("to_name", data.ToName)
	_ = writer.WriteField("document_name", data.DocumentName)
	_ = writer.WriteField("description", data.Description)
	_ = writer.WriteField("data", data.Data)
	file, err := os.Open("./" + data.Data)
	if err != nil {
		return err
	}

	defer file.Close()
	part2, _ := writer.CreateFormFile("file", filepath.Base("./"+data.Data))
	_, err = io.Copy(part2, file)
	if err != nil {
		fmt.Println(err)
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return err
}

func (h *helper) PostRequestAcceptAsk(url, fileName string, file *multipart.FileHeader) error {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("file_name", fileName)
	pdf, err := file.Open()
	if err != nil {
		log.Println("error : ", err)
	}

	part2, _ := writer.CreateFormFile("file", file.Filename)

	_, err = io.Copy(part2, pdf)
	if err != nil {
		fmt.Println(err)
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))

	return err
}

func (h *helper) SendIniatialETH(address string) error {
	privateKey, err := crypto.HexToECDSA(os.Getenv("NODE_ADMIN_PRIVATE_KEY"))
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := h.cl.PendingNonceAt(h.ctx, fromAddress)
	if err != nil {
		return err
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(210000)               // in units
	gasPrice, err := h.cl.SuggestGasPrice(h.ctx)
	if err != nil {
		return err
	}

	toAddress := common.HexToAddress(address)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, []byte("initial ETH"))

	chainID, err := h.cl.NetworkID(h.ctx)
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return err
	}

	err = h.cl.SendTransaction(h.ctx, signedTx)
	if err != nil {
		return err
	}

	return err
}
