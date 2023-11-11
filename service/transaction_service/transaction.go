package transactionservice

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"mime/multipart"
	"os"
	"path/filepath"
	"signature-app/database/model"
	"signature-app/database/repository"
	requestdomain "signature-app/domain/request_domain"
	responsedomain "signature-app/domain/response_domain"
	"signature-app/helper"
	"signature-app/notification"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lithammer/shortuuid"
)

type transactionService struct {
	cl  *ethclient.Client
	ctx context.Context
	db  *repository.Database
}

func NewTransaction(cl *ethclient.Client, ctx context.Context, db *repository.Database) *transactionService {
	return &transactionService{
		cl:  cl,
		ctx: ctx,
		db:  db,
	}
}

func (s *transactionService) GetAllTransactions() (txs []responsedomain.TransactionResponse, err error) {
	head, err := s.cl.HeaderByNumber(s.ctx, nil)
	if err != nil {
		return
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
					TimeStamp: tx.Time(),
				})
		}
	}

	return txs, nil
}

func (s *transactionService) SendTransaction(strPrivateKey string, body model.TransactionModel, pdf *multipart.FileHeader) (*responsedomain.TransactionResponse, error) {
	filePath := filepath.Join("static", body.Id+filepath.Ext(pdf.Filename))
	body.Data = filePath

	byteData, err := json.Marshal(body)
	if err != nil {
		return nil, err
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
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, []byte(byteData))

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

	if pdf != nil {
		fmt.Println("Uploading")

		var wg sync.WaitGroup
		wg.Add(2)

		help := helper.NewHelper(s.cl, s.ctx)
		go func() {
			defer wg.Done()

			help.PostRequestAcceptAsk(fmt.Sprintf("http://localhost%s/accept-ask", os.Getenv("PORT_ADDRESS_2")), filePath, pdf)
		}()

		go func() {
			defer wg.Done()

			help.PostRequestAcceptAsk(fmt.Sprintf("http://localhost%s/accept-ask", os.Getenv("PORT_ADDRESS_3")), filePath, pdf)
		}()

		dst, err := os.Create(filePath)
		if err != nil {
			log.Println("error saving file")
		}
		defer dst.Close()
		file, err := pdf.Open()
		if err != nil {
			fmt.Println("error opening pdf file")
		}

		defer file.Close()
		_, err = io.Copy(dst, file)
		if err != nil {
			fmt.Println("Error saving signed pdf")
		}

		wg.Wait()
	}

	return &responsedomain.TransactionResponse{
		Id: signedTx.Hash().Hex(),
	}, nil
}

func (s *transactionService) AddAsk(payload model.TransactionModel) error {
	var wg sync.WaitGroup
	help := helper.NewHelper(s.cl, s.ctx)
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
	if err != nil {
		return err
	}

	wg.Wait()

	token, err := s.db.GetDeviceToken(payload.ToAddress, 1)
	if err == nil && token != "" {
		log.Println("sending notif")
		notif, err := notification.NewNotification(s.ctx)
		if err != nil {
			return nil
		}

		notif.SendNotification(requestdomain.NotificationData{
			Data: map[string]string{
				"title": "permintaan tanda tangan baru",
				"body":  fmt.Sprintf("%s meminta tanda tangan anda", payload.FromName),
			},
			Token: token,
		})
	} else {
		log.Println("Skip notif")
	}

	return err
}

func (s *transactionService) AddAskFromServers(payload model.TransactionModel) error {
	token, err := s.db.GetDeviceToken(payload.ToAddress, 1)
	if err == nil && token != "" {
		log.Println("sending notif")
		notif, err := notification.NewNotification(s.ctx)
		if err != nil {
			return nil
		}

		notif.SendNotification(requestdomain.NotificationData{
			Data: map[string]string{
				"title": "permintaan tanda tangan baru",
				"body":  fmt.Sprintf("%s meminta tanda tangan anda", payload.FromName),
			},
			Token: token,
		})
	} else {
		log.Println("Skip notif")
	}

	_, err = s.db.AddAsk(payload)

	return err
}

func (s *transactionService) GetAsk(address string, status []int) (data []model.TransactionModel, err error) {
	data, err = s.db.GetAsk(address, status)
	return
}

func (s *transactionService) GetOneAsk(id string) (data []model.TransactionModel, err error) {
	data, err = s.db.GetOneAsk(id)
	return
}

func (s *transactionService) GetGive(address string, status []int) (data []model.TransactionModel, err error) {
	data, err = s.db.GetGive(address, status)
	return
}

// func (s *transactionService) AcceptAskFromServer(data model.TransactionModel) error {
// 	err := s.db.AcceptAsk(data.Id, data.TxId, data.UpdatedAt)
// 	return err
// }

func (s *transactionService) GetASignedTransaction(txId string) (data model.TransactionModel, err error) {
	tx, isPending, err := s.cl.TransactionByHash(s.ctx, common.HexToHash(txId))
	if err != nil {
		return
	}

	fmt.Println("isPeding", isPending)
	err = json.Unmarshal(tx.Data(), &data)
	if err != nil {
		return
	}

	data.TxId = string(tx.Hash().String())
	from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		return
	}

	data.FromAddress = string(from.Hex())
	data.ToAddress = tx.To().Hex()
	data.UpdatedAt = tx.Time().Format(time.RFC3339)

	return
}

func (s *transactionService) AddAskDirect(strPrivateKey string, payload model.TransactionModel) error {
	payload.Id = shortuuid.New()
	payload.CreatedAt = time.Now().Format(time.RFC3339)
	payload.UpdatedAt = time.Now().Format(time.RFC3339)

	privateKey, err := crypto.HexToECDSA(strPrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.cl.PendingNonceAt(s.ctx, fromAddress)
	if err != nil {
		return err
	}

	value := big.NewInt(0)     // in wei (1 eth)
	gasLimit := uint64(210000) // in units
	gasPrice, err := s.cl.SuggestGasPrice(s.ctx)
	if err != nil {
		return err
	}

	toAddress := common.HexToAddress(payload.ToAddress)
	data, err := json.Marshal(payload)
	if err != nil {
		data = make([]byte, 0)
	}
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := s.cl.NetworkID(s.ctx)
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return err
	}

	err = s.cl.SendTransaction(s.ctx, signedTx)
	if err != nil {
		return err
	}

	return err
}

func (s *transactionService) SignDirect(strPrivateKey, id string) error {
	res, err := s.db.GetOneAskWithId(id)
	if err != nil || len(res) != 1 {
		return err
	}

	ask := res[0]

	privateKey, err := crypto.HexToECDSA(strPrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := s.cl.PendingNonceAt(s.ctx, fromAddress)
	if err != nil {
		return err
	}

	value := big.NewInt(0)     // in wei (1 eth)
	gasLimit := uint64(210000) // in units
	gasPrice, err := s.cl.SuggestGasPrice(s.ctx)
	if err != nil {
		return err
	}

	toAddress := common.HexToAddress(ask.ToAddress)
	data, err := json.Marshal(ask)
	if err != nil {
		data = make([]byte, 0)
	}
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := s.cl.NetworkID(s.ctx)
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return err
	}

	err = s.cl.SendTransaction(s.ctx, signedTx)
	if err != nil {
		return err
	}

	return err
}
