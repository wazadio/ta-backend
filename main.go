package main

import (
	"context"
	"log"
	"os"
	accountcontroller "signature-app/controller/account_controller"
	contractcontroller "signature-app/controller/contract_controller"
	transactioncontroller "signature-app/controller/transaction_controller"
	"signature-app/database/repository"
	accountservice "signature-app/service/account_service"
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
	ipcUrl := "../fakultas_home/database/geth.ipc"
	httpUrl := "http://127.0.0.1:9001"
	_ = ipcUrl

	cl, err := ethclient.Dial(httpUrl)
	if err != nil {
		panic(err)
	}

	db, err := repository.NewDb()
	if err != nil {
		panic(err)
	}

	accountService := accountservice.NewAccountDetail()
	accountController := accountcontroller.NewController(cl, ctx, accountService)

	transactionService := transactionservice.NewTransaction(cl, ctx, db)
	transactioncontroller := transactioncontroller.NewController(transactionService)

	contractController := contractcontroller.NewController(cl, ctx, os.Getenv("CONTRACT_ADDRESS"))

	r := gin.Default()
	r.Static("/static", "./static")

	r.POST("/create-account", accountController.CreateAccount)
	r.POST("/import-account", accountController.ImportAccount)

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

	r.POST("/Ask", transactioncontroller.AddAsk)
	r.POST("/ask-server", transactioncontroller.AddAskFromServers)

	r.Run()
}
