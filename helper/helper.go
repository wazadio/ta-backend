package helper

import (
	"context"
	"math"
	"math/big"
	"net/http"

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

func (h *helper) PostRequestAsk(url string, payload any) (int, error) {
	resp, err := http.Post(url)
}
