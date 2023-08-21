package transactioncontroller

import (
	"net/http"
	requestdomain "signature-app/domain/request_domain"
	transactionservice "signature-app/service/transaction_service"

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
