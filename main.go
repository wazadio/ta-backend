package main

import (
	"context"
	"log"
	"os"
	accountcontroller "signature-app/controller/account_controller"
	contractcontroller "signature-app/controller/contract_controller"
	transactioncontroller "signature-app/controller/transaction_controller"
	"signature-app/database/repository"
	"signature-app/helper"
	accountservice "signature-app/service/account_service"
	"signature-app/service/interfaces"
	transactionservice "signature-app/service/transaction_service"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	wsUrl := os.Getenv("NODE_WEBSOCKET_ADDRESS")

	cl, err := ethclient.Dial(wsUrl)
	if err != nil {
		log.Fatal("error dialing eth client : ", err)
	}

	db, err := repository.NewDb()
	if err != nil {
		log.Fatal(err)
	}

	// kafkaTopics := kafkaclient.KafkaTopics{
	// 	Token:     os.Getenv("KAFKA_TOPIC_TOKEN"),
	// 	PendingTx: os.Getenv("KAFKA_TOPIC_PENDING_TX"),
	// 	NewBlock:  os.Getenv("KAFKA_TOPIC_NEW_BLOCK"),
	// }

	// log.Printf("%+v", kafkaTopics)

	// kc := kafkaclient.NewKafkaClient(ctx, os.Getenv("KAFKA_ADDRESS"), os.Getenv("KAFKA_GROUP"), db, kafkaTopics)
	subcription := helper.NewSubcription(cl, ctx, db)

	// go kc.ReadMessagesToken()
	// go kc.ReadMessagesPendingTx()
	go subcription.SubcribeNewBlock()
	go subcription.SubcribePendingTx()

	var accountService interfaces.AccountService = accountservice.NewAccountService(ctx, cl, db)
	accountController := accountcontroller.NewController(cl, ctx, accountService)

	var transactionService interfaces.TransactionService = transactionservice.NewTransaction(cl, ctx, db)
	transactioncontroller := transactioncontroller.NewController(transactionService)

	contractController := contractcontroller.NewController(cl, ctx, os.Getenv("CONTRACT_ADDRESS"), db)

	r := gin.Default()
	r.Static("/static", "./static")

	r.POST("/create-account", accountController.CreateAccount)
	r.POST("/import-account", accountController.ImportAccount)
	r.POST("/new-device-token", accountController.AddNewToken)

	r.GET("/transactions", transactioncontroller.GetAllTransactions)
	r.POST("/send", transactioncontroller.SendData)

	r.POST("/deploy-contract", contractController.Deploy)
	r.GET("/admin", contractController.GetAdmin)
	r.GET("/dokumen", contractController.GetDokumen)
	r.GET("/identitas", contractController.GetIdentitas)
	r.POST("/add-dokumen", contractController.AddDokumen)
	r.POST("/add-identitas", contractController.AddIdentitas)
	r.GET("/faucet", contractController.GetETH)
	r.GET("/balance", accountController.GetETH)
	r.DELETE("/delete-dokumen", contractController.DeleteDokumen)
	r.DELETE("/delete-identitas", contractController.DeleteIdentitas)

	r.POST("/Ask", transactioncontroller.AddAsk)
	r.POST("/ask-server", transactioncontroller.AddAskFromServers)
	r.GET("/get-ask", transactioncontroller.GetAsk)
	r.POST("/accept-ask", transactioncontroller.AcceptAsk)
	r.GET("/get-ask/:id", transactioncontroller.GetOneAsk)
	r.GET("/get-give", transactioncontroller.GetGive)
	r.GET("/get-signed-transaction/:tx-id", transactioncontroller.GetASignedTransaction)
	r.POST("/give-direct", transactioncontroller.AddAskDirect)
	r.POST("/sign-direct/:id", transactioncontroller.SignDirect)

	r.Run(os.Getenv("PORT_ADDRESS"))
}
