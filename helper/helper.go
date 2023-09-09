package helper

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"signature-app/database/model"

	"github.com/ethereum/go-ethereum/common"
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
	_ = writer.WriteField("from_address", "address")
	_ = writer.WriteField("from_name", "Juni Dio Kasandra")
	_ = writer.WriteField("to_address", "toaddress")
	_ = writer.WriteField("to_name", "wazadio")
	file, err := os.Open("./" + data.Data)
	defer file.Close()
	part2, err := writer.CreateFormFile("file", filepath.Base("./"+data.Data))
	_, err = io.Copy(part2, file)
	if err != nil {
		fmt.Println(err)
	}

	err = writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	return nil
}
