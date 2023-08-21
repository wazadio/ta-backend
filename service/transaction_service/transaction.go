package transactionservice

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"signature-app/database/model"
	"signature-app/database/repository"
	responsedomain "signature-app/domain/response_domain"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionService struct {
	cl  *ethclient.Client
	ctx context.Context
	db  *repository.Database
}

func NewTransaction(cl *ethclient.Client, ctx context.Context, db *repository.Database) *TransactionService {
	return &TransactionService{
		cl:  cl,
		ctx: ctx,
		db:  db,
	}
}

func (s *TransactionService) GetAllTransactions() (txs []responsedomain.TransactionResponse, err error) {
	head, err := s.cl.HeaderByNumber(s.ctx, nil)
	if err != nil {
		panic(err)
	}

	for i := int64(0); i <= head.Number.Int64(); i++ {
		block, err := s.cl.BlockByNumber(s.ctx, big.NewInt(i))
		if err != nil {
			return txs, err
		}

		for _, tx := range block.Transactions() {
			from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
			if err != nil {
				return txs, err
			}

			var toAddress string
			if tx.To() == nil {
				toAddress = ""
			} else {
				toAddress = tx.To().Hex()
			}

			txs = append(txs,
				responsedomain.TransactionResponse{
					Id:        string(tx.Hash().String()),
					From:      from.Hex(),
					To:        toAddress,
					Data:      string(tx.Data()),
					TimeStamp: time.Unix(int64(block.Time()), 0),
				})
		}
	}

	return txs, nil
}

func (s *TransactionService) SendTransaction(strPrivateKey, hexToAddress, data string) (*responsedomain.TransactionResponse, error) {
	privateKey, err := crypto.HexToECDSA(strPrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.cl.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	value := big.NewInt(0)     // in wei (1 eth)
	gasLimit := uint64(210000) // in units
	gasPrice, err := s.cl.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	toAddress := common.HexToAddress(hexToAddress)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, []byte(data))

	chainID, err := s.cl.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	err = s.cl.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	return &responsedomain.TransactionResponse{
		Id: signedTx.Hash().Hex(),
	}, nil
}

func (s *TransactionService) Ask(payload model.TransactionModel) error {

}
