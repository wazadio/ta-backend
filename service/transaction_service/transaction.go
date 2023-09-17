package transactionservice

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"signature-app/database/model"
	"signature-app/database/repository"
	responsedomain "signature-app/domain/response_domain"
	"signature-app/helper"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lithammer/shortuuid"
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

func (s *TransactionService) SendTransaction(strPrivateKey, to, data, askId string) (*responsedomain.TransactionResponse, error) {
	var wg sync.WaitGroup
	help := helper.NewHelper(s.cl, s.ctx)
	body := model.TransactionModel{
		Id:        askId,
		ToAddress: to,
		Data:      data,
	}

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

	toAddress := common.HexToAddress(body.ToAddress)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, []byte(body.Data))

	chainID, err := s.cl.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	body.TxId = string(signedTx.Hash().Hex())
	body.UpdatedAt = signedTx.Time().Format(time.RFC3339)

	err = s.cl.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	wg.Add(2)
	go func() {
		defer wg.Done()

		help.PostRequestAcceptAsk(fmt.Sprintf("http://localhost%s/accept-ask", os.Getenv("PORT_ADDRESS_2")), body)
	}()

	go func() {
		defer wg.Done()

		help.PostRequestAcceptAsk(fmt.Sprintf("http://localhost%s/accept-ask", os.Getenv("PORT_ADDRESS_3")), body)
	}()
	err = s.db.AcceptAsk(body.Id, body.TxId, body.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &responsedomain.TransactionResponse{
		Id: signedTx.Hash().Hex(),
	}, nil
}

func (s *TransactionService) AddAsk(payload model.TransactionModel) error {
	var wg sync.WaitGroup
	help := helper.NewHelper(s.cl, s.ctx)
	payload.Id = shortuuid.New()
	payload.Status = 0
	payload.CreatedAt = time.Now().Format(time.RFC3339)

	wg.Add(2)
	go func() {
		defer wg.Done()

		help.PostRequestAsk(fmt.Sprintf("http://localhost%s/ask-server", os.Getenv("PORT_ADDRESS_2")), payload)
	}()

	go func() {
		defer wg.Done()

		help.PostRequestAsk(fmt.Sprintf("http://localhost%s/ask-server", os.Getenv("PORT_ADDRESS_3")), payload)
	}()
	_, err := s.db.AddAsk(payload)

	wg.Wait()

	return err
}

func (s *TransactionService) AddAskFromServers(payload model.TransactionModel) error {
	_, err := s.db.AddAsk(payload)
	return err
}

func (s *TransactionService) UpdateAsk(data model.TransactionModel) error {
	err := s.db.AcceptAsk(data.Id, data.TxId, data.UpdatedAt)
	return err
}

func (s *TransactionService) GetAsk(address string, status int) (data []model.TransactionModel, err error) {
	data, err = s.db.GetAsk(address, status)
	return
}
