package accountcontroller

import (
	"context"
	"net/http"
	requestdomain "signature-app/domain/request_domain"
	"signature-app/helper"
	accountservice "signature-app/service/account_service"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type controller struct {
	cl      *ethclient.Client
	ctx     context.Context
	Service *accountservice.AccountService
}

func NewController(cl *ethclient.Client, ctx context.Context, accountService *accountservice.AccountService) *controller {
	return &controller{
		cl:      cl,
		ctx:     ctx,
		Service: accountService,
	}
}

func (c *controller) CreateAccount(ctx *gin.Context) {
	body := requestdomain.CreateAccountRequest{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Service.CreateNewAccount()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (c *controller) ImportAccount(ctx *gin.Context) {
	body := requestdomain.ImportAccountRequest{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Service.ImportAccount(body.Mnemonic)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (c *controller) GetETH(ctx *gin.Context) {
	address := ctx.Request.Header.Get("address")

	service := helper.NewHelper(c.cl, c.ctx)

	res, err := service.CheckBalance(address)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	balance, _ := res.Float64()

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"balance": balance,
		},
	)
}
