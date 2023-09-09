package transactioncontroller

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"signature-app/database/model"
	requestdomain "signature-app/domain/request_domain"
	transactionservice "signature-app/service/transaction_service"
	"time"

	"github.com/gin-gonic/gin"
)

type controller struct {
	Service *transactionservice.TransactionService
}

func NewController(transactionService *transactionservice.TransactionService) *controller {
	return &controller{
		Service: transactionService,
	}
}

func (c *controller) GetAllTransactions(ctx *gin.Context) {
	body := requestdomain.GetTransactionsRequest{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	res, err := c.Service.GetAllTransactions()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (c *controller) SendData(ctx *gin.Context) {
	body := requestdomain.SendTransactionRequest{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Service.SendTransaction(body.PrivateKey, body.To, body.Data)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (c *controller) AddAsk(ctx *gin.Context) {
	body := model.TransactionModel{}
	err := ctx.Bind(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	file, err := ctx.FormFile("file")
	if err == nil && file != nil {
		fmt.Println("Uploading")
		filePath := filepath.Join("static", fmt.Sprint(time.Now().Format(time.DateOnly)+body.FromAddress)+filepath.Ext(file.Filename))
		err = ctx.SaveUploadedFile(file, filePath)
		if err != nil {
			log.Println("error saving file")
		}
		body.Data = filePath
	}

	err = c.Service.AddAsk(body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "oke",
		},
	)
}

func (c *controller) AddAskFromServers(ctx *gin.Context) {
	body := model.TransactionModel{}
	err := ctx.Bind(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	file, err := ctx.FormFile("file")
	if err == nil && file != nil {
		fmt.Println("Uploading")
		filePath := filepath.Join("static", fmt.Sprint(time.Now().Format(time.DateOnly)+body.FromAddress)+filepath.Ext(file.Filename))
		err = ctx.SaveUploadedFile(file, filePath)
		if err != nil {
			log.Println("error saving file")
		}
		body.Data = filePath
	}

	err = c.Service.AddAskFromServers(body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": "oke",
		},
	)
}

func (c *controller) AcceptAsk(ctx *gin.Context) {
	body := model.TransactionModel{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.Service.UpdateAsk(body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		err,
	)
}
