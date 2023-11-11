package transactioncontroller

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"signature-app/database/model"
	requestdomain "signature-app/domain/request_domain"
	"signature-app/service/interfaces"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid"
)

type controller struct {
	service interfaces.TransactionService
}

func NewController(transactionService interfaces.TransactionService) *controller {
	return &controller{
		service: transactionService,
	}
}

func (c *controller) GetAllTransactions(ctx *gin.Context) {
	res, err := c.service.GetAllTransactions()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (c *controller) SendData(ctx *gin.Context) {
	body := requestdomain.SendTransactionRequest{}
	err := ctx.Bind(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.service.SendTransaction(body.PrivateKey, body.TransactionModel, file)
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

	newUuuid := shortuuid.New()

	file, err := ctx.FormFile("file")
	if err == nil && file != nil {
		fmt.Println("Uploading")
		body.Data = filepath.Join("static", newUuuid+filepath.Ext(file.Filename))
		err = ctx.SaveUploadedFile(file, body.Data)
		if err != nil {
			log.Println("error saving file")
		}
	}

	toAddresses := ctx.PostFormArray("to_address")
	toNames := ctx.PostFormArray("to_name")

	for i, v := range toAddresses {
		data := body
		data.Id = newUuuid
		data.ToAddress = v
		data.ToName = toNames[i]

		err = c.service.AddAsk(data)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
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
		err = ctx.SaveUploadedFile(file, body.Data)
		if err != nil {
			log.Println("error saving file")
		}
	}

	err = c.service.AddAskFromServers(body)
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
	file, err := ctx.FormFile("file")
	if err == nil && file != nil {
		fmt.Println("Uploading")
		filePath := filepath.Join("static", file.Filename)
		err = ctx.SaveUploadedFile(file, filePath)
		if err != nil {
			log.Println("error saving file")
		}
	}

	ctx.JSON(
		http.StatusOK,
		err,
	)
}

func (c *controller) GetAsk(ctx *gin.Context) {
	address := ctx.Request.Header.Get("address")
	listStatus := strings.Split(ctx.Request.Header.Get("status"), ",")
	intStatus := []int{}

	for _, v := range listStatus {
		status, err := strconv.Atoi(v)
		if err != nil {
			intStatus = append(intStatus, 0)
			continue
		}

		intStatus = append(intStatus, status)
	}

	res, err := c.service.GetAsk(address, intStatus)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (c *controller) GetOneAsk(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.service.GetOneAsk(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (c *controller) GetGive(ctx *gin.Context) {
	address := ctx.Request.Header.Get("address")
	listStatus := strings.Split(ctx.Request.Header.Get("status"), ",")
	intStatus := []int{}

	for _, v := range listStatus {
		status, err := strconv.Atoi(v)
		if err != nil {
			intStatus = append(intStatus, 0)
			continue
		}

		intStatus = append(intStatus, status)
	}

	res, err := c.service.GetGive(address, intStatus)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

func (c *controller) GetASignedTransaction(ctx *gin.Context) {
	txId := ctx.Param("tx-id")

	res, err := c.service.GetASignedTransaction(txId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(
		http.StatusOK,
		res,
	)
}

// give sign form signer without asking
func (c *controller) AddAskDirect(ctx *gin.Context) {
	body := requestdomain.SendTransactionRequest{}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = c.service.AddAskDirect(body.PrivateKey, body.TransactionModel)
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

// give sign without embed sign image in pdf
func (c *controller) SignDirect(ctx *gin.Context) {
	id := ctx.Param("id")
	privateKey := ctx.Request.Header.Get("private-key")
	err := c.service.SignDirect(privateKey, id)
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
