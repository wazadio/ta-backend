package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"signature-app/database/model"
	"signature-app/database/repository"
	requestdomain "signature-app/domain/request_domain"
	"signature-app/notification"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type subcription struct {
	Cl  *ethclient.Client
	Gcl *gethclient.Client
	Ctx context.Context
	Db  *repository.Database
}

type NewPendingTransaction struct {
	TxId      string
	TimeStamp string
}

func NewSubcription(cl *ethclient.Client, ctx context.Context, db *repository.Database) *subcription {
	rpcClient, err := rpc.Dial(os.Getenv("NODE_WEBSOCKET_ADDRESS"))
	if err != nil {
		log.Fatalln("error creating new subcription : ", err)
	}

	gcl := gethclient.New(rpcClient)
	log.Printf("gcl value = %+v", gcl)
	return &subcription{
		Cl:  cl,
		Ctx: ctx,
		Gcl: gcl,
		Db:  db,
	}
}

func (s *subcription) SubcribePendingTx() {
	pendingTxHash := make(chan common.Hash)
	sub, err := s.Gcl.SubscribePendingTransactions(s.Ctx, pendingTxHash)
	for err != nil {
		log.Println("error subcribe to pending transaction pool : ", err)
		log.Println("retrying connection")
		sub, err = s.Gcl.SubscribePendingTransactions(s.Ctx, pendingTxHash)
	}

	log.Println("subcription to pending tx pool successful")

	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			for err != nil {
				log.Println("error subcribe to pending transaction pool : ", err)
				log.Println("retrying connection")
				sub, err = s.Gcl.SubscribePendingTransactions(s.Ctx, pendingTxHash)
			}

			log.Println("subcription to pending tx pool successful")
		case pendingTx := <-pendingTxHash:
			tx, isPending, err := s.Cl.TransactionByHash(s.Ctx, pendingTx)
			if err != nil {
				log.Printf("error getting pending transaction with hash %s : %s\n", pendingTx, err.Error())
				continue
			}

			if isPending {
				log.Println("new pending transaction arrived")
			} else {
				log.Printf("so weird transaction status is not pending %s\n", pendingTx)
				continue
			}
			// newPendingTx := NewPendingTransaction{
			// 	TxId:      tx.Hash().Hex(),
			// 	TimeStamp: tx.Time().Format(time.RFC3339),
			// }
			// data, err := json.Marshal(newPendingTx)
			// if err != nil {
			// 	log.Println("error marshal new pending transaction data")
			// }

			// err = s.Kc.ProduceMessagePendingTx(data)
			// if err != nil {
			// 	log.Println("error produce kafka message for new pending transaction")
			// }

			var data model.TransactionModel
			err = json.Unmarshal(tx.Data(), &data)
			if err != nil {
				log.Println("Error parsing new pending block transaction data : ", err)
			}

			id, err := s.Db.GetSingleTx(tx.Hash().String(), data.Id)
			if err != nil {
				continue
			}

			err = s.Db.UpdatePendingAsk(id, tx.Time().Format(time.RFC3339))
			if err != nil {
				log.Println("Error updating pending ask")
				continue
			}
		}
	}
}

func (s *subcription) SubcribeNewBlock() {
	headers := make(chan *types.Header)
	sub, err := s.Cl.SubscribeNewHead(s.Ctx, headers)
	for err != nil {
		log.Println("error subcribe to new block : ", err)
		log.Println("Retrying connection")
		sub, err = s.Cl.SubscribeNewHead(s.Ctx, headers)
	}

	log.Println("subcription to new block successful")

	for {
		select {
		case err := <-sub.Err():
			for err != nil {
				log.Println("error subcribe to new block : ", err)
				log.Println("Retrying connection")
				sub, err = s.Cl.SubscribeNewHead(s.Ctx, headers)
			}

			log.Println("subcription to new block successful")
		case header := <-headers:
			log.Println("New Block Arrived")
			block, err := s.Cl.BlockByHash(s.Ctx, header.Hash())
			if err != nil {
				log.Println("error getting block by header : ", err)
			}

			for _, tx := range block.Transactions() {
				txId := tx.Hash().String()
				if tx.To().Hex() == "" || tx.To().Hex() == os.Getenv("CONTRACT_ADDRESS") {
					continue
				} else {
					var data model.TransactionModel
					err := json.Unmarshal(tx.Data(), &data)
					if err != nil {
						log.Println("Error parsing new block transaction data : ", err)
					}

					log.Printf("Data : %+v\n", data)

					if data.Id == "" {
						continue
					}

					id, err := s.Db.GetSingleTx(txId, data.Id)
					if err != nil {
						log.Println("error get single tx : ", err)
					}

					log.Printf("id = %s\n", id)

					if id == "" {
						data.Status = 1
						_, err = s.Db.AddAsk(data)
						if err != nil {
							log.Println("error add new ask direct : ", err)
						}
						continue
					}

					log.Println("new id : ", id)

					err = s.Db.AcceptAsk(id, txId, tx.Time().Format(time.RFC3339))
					if err != nil {
						log.Println("error accept ask : ", err)
					}
					log.Println("new transaction accepted")

					log.Println("sending notif")
					token, err := s.Db.GetDeviceToken(data.FromAddress, 1)
					if err != nil || token == "" {
						log.Println("skip notif")
					} else {
						notif, err := notification.NewNotification(s.Ctx)
						if err != nil {
							continue
						}

						notif.SendNotification(requestdomain.NotificationData{
							Data: map[string]string{
								"title": "Tanda Tangan Baru",
								"body":  fmt.Sprintf("%s memberikan anda tangan tangan", data.ToName),
							},
							Token: token,
						})
					}
				}
			}
		}
	}
}
