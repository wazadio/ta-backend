package contractcontroller

import (
	"context"
	"net/http"
	requestdomain "signature-app/domain/request_domain"
	smartcontract "signature-app/service/smart_contract"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type controller struct {
	cl              *ethclient.Client
	ctx             context.Context
	contractAddress string
}

func NewController(cl *ethclient.Client, ctx context.Context, contractAddress string) *controller {
	return &controller{
		cl:              cl,
		ctx:             ctx,
		contractAddress: contractAddress,
	}
}

func (c *controller) Deploy(ctx *gin.Context) {
	body := requestdomain.DeployContractRequest{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	caller := smartcontract.NewCaller(c.cl, c.ctx, c.contractAddress)

	res, err := caller.Deploy(body.PrivateKey)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (c *controller) GetAdmin(ctx *gin.Context) {
	caller := smartcontract.NewCaller(c.cl, c.ctx, c.contractAddress)

	res, err := caller.GetAdmin()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (c *controller) GetDokumen(ctx *gin.Context) {
	caller := smartcontract.NewCaller(c.cl, c.ctx, c.contractAddress)

	res, err := caller.GetDokumen()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (c *controller) GetIdentitas(ctx *gin.Context) {
	caller := smartcontract.NewCaller(c.cl, c.ctx, c.contractAddress)

	res, err := caller.GetIdentitas()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (c *controller) AddDokumen(ctx *gin.Context) {
	body := requestdomain.AddDokumenRequest{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	caller := smartcontract.NewCaller(c.cl, c.ctx, c.contractAddress)

	res, err := caller.AddDokumen(body.NamaDokumen, body.PrivateKey)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (c *controller) AddIdentitas(ctx *gin.Context) {
	body := requestdomain.AddIdentitasRequest{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	caller := smartcontract.NewCaller(c.cl, c.ctx, c.contractAddress)

	res, err := caller.AddIdentitas(body)
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
	privateKey := ctx.Request.Header.Get("private-key")

	caller := smartcontract.NewCaller(c.cl, c.ctx, c.contractAddress)

	res, err := caller.GetETH(privateKey)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}
