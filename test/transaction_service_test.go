package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	transactioncontroller "signature-app/controller/transaction_controller"
	"signature-app/database/model"
	responsedomain "signature-app/domain/response_domain"
	"signature-app/test/mocks"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllTransactions(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockResp := []responsedomain.TransactionResponse{
			{
				Id:        "id",
				From:      "from",
				To:        "to",
				Data:      "data",
				TimeStamp: time.Now(),
			},
			{
				Id:        "id 2",
				From:      "from 2",
				To:        "to 2",
				Data:      "data 2",
				TimeStamp: time.Now().Add(time.Duration(time.Now().Hour())),
			},
		}

		mockService := new(mocks.MockTransactionService)
		mockService.On("GetAllTransactions").Return(mockResp, nil)

		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.GET("/transactions", transactioncontroller.GetAllTransactions)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, "/transactions", bytes.NewReader([]byte("{}")))
		assert.NoError(t, err)

		r.ServeHTTP(rr, req)

		resp, err := json.Marshal(mockResp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, resp, rr.Body.Bytes())
		mockService.AssertExpectations(t)
	})

	t.Run("not success", func(t *testing.T) {
		mockService := new(mocks.MockTransactionService)
		mockService.On("GetAllTransactions").Return(nil, errors.New("error"))

		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.GET("/transactions", transactioncontroller.GetAllTransactions)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, "/transactions", bytes.NewReader([]byte("{}")))
		assert.NoError(t, err)

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockService.AssertExpectations(t)
	})
}

func TestSend(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockResp := responsedomain.TransactionResponse{
			Id:        "id",
			From:      "from",
			To:        "to",
			Data:      "data",
			TimeStamp: time.Now(),
		}

		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)
		file, err := os.Open("./" + "mocks/test.pdf")
		assert.NoError(t, err)
		defer file.Close()

		part2, _ := writer.CreateFormFile("file", filepath.Base("./"+"mocks/test.pdf"))
		_, err = io.Copy(part2, file)
		assert.NoError(t, err)
		err = writer.Close()
		assert.NoError(t, err)

		mockService := new(mocks.MockTransactionService)
		mockService.On("SendTransaction", mock.Anything, mock.AnythingOfType("model.TransactionModel"), mock.AnythingOfType("*multipart.FileHeader")).Return(&mockResp, nil)

		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.POST("/send", transactioncontroller.SendData)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodPost, "/send", payload)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		r.ServeHTTP(rr, req)

		resp, err := json.Marshal(mockResp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, resp, rr.Body.Bytes())
		mockService.AssertExpectations(t)
	})

	t.Run("not success", func(t *testing.T) {
		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)
		file, err := os.Open("./" + "mocks/test.pdf")
		assert.NoError(t, err)
		defer file.Close()

		part2, _ := writer.CreateFormFile("file", filepath.Base("./"+"mocks/test.pdf"))
		_, err = io.Copy(part2, file)
		assert.NoError(t, err)
		err = writer.Close()
		assert.NoError(t, err)

		mockService := new(mocks.MockTransactionService)
		mockService.On("SendTransaction", mock.Anything, mock.AnythingOfType("model.TransactionModel"), mock.AnythingOfType("*multipart.FileHeader")).Return(nil, errors.New("error"))

		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.POST("/send", transactioncontroller.SendData)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodPost, "/send", payload)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockService.AssertExpectations(t)
	})
}

func TestAddAsk(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)
		_ = writer.WriteField("to_address", "")
		_ = writer.WriteField("to_name", "")
		err := writer.Close()
		assert.NoError(t, err)

		mockService := new(mocks.MockTransactionService)
		mockService.On("AddAsk", mock.AnythingOfType("model.TransactionModel")).Return(nil)
		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.POST("/Ask", transactioncontroller.AddAsk)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodPost, "/Ask", payload)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, `{"message":"oke"}`, rr.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("not success", func(t *testing.T) {
		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)
		_ = writer.WriteField("to_address", "")
		_ = writer.WriteField("to_name", "")
		err := writer.Close()
		assert.NoError(t, err)

		mockService := new(mocks.MockTransactionService)
		mockService.On("AddAsk", mock.AnythingOfType("model.TransactionModel")).Return(errors.New("error"))
		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.POST("/Ask", transactioncontroller.AddAsk)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodPost, "/Ask", payload)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockService.AssertExpectations(t)
	})
}

func TestAddAskFromServer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)
		err := writer.Close()
		assert.NoError(t, err)

		mockService := new(mocks.MockTransactionService)
		mockService.On("AddAskFromServers", mock.AnythingOfType("model.TransactionModel")).Return(nil)
		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.POST("/ask-server", transactioncontroller.AddAskFromServers)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodPost, "/ask-server", payload)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, `{"message":"oke"}`, rr.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("not success", func(t *testing.T) {
		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)
		err := writer.Close()
		assert.NoError(t, err)

		mockService := new(mocks.MockTransactionService)
		mockService.On("AddAskFromServers", mock.AnythingOfType("model.TransactionModel")).Return(errors.New("error"))
		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.POST("/ask-server", transactioncontroller.AddAskFromServers)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodPost, "/ask-server", payload)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockService.AssertExpectations(t)
	})
}

// func TestAcceptAsk(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	t.Run("success", func(t *testing.T) {
// 		payload := &bytes.Buffer{}
// 		writer := multipart.NewWriter(payload)
// 		err := writer.Close()
// 		assert.NoError(t, err)

// 		mockService := new(mocks.MockTransactionService)
// 		mockService.On("AcceptAskFromServer", mock.AnythingOfType("model.TransactionModel")).Return(nil)
// 		transactioncontroller := transactioncontroller.NewController(mockService)
// 		r := gin.Default()
// 		r.POST("/accept-ask", transactioncontroller.AcceptAsk)

// 		rr := httptest.NewRecorder()

// 		req, err := http.NewRequest(http.MethodPost, "/accept-ask", payload)
// 		assert.NoError(t, err)
// 		req.Header.Set("Content-Type", writer.FormDataContentType())

// 		r.ServeHTTP(rr, req)

// 		assert.Equal(t, http.StatusOK, rr.Code)
// 		assert.Equal(t, "null", rr.Body.String())
// 		mockService.AssertExpectations(t)
// 	})

// 	t.Run("not success", func(t *testing.T) {
// 		payload := &bytes.Buffer{}
// 		writer := multipart.NewWriter(payload)
// 		err := writer.Close()
// 		assert.NoError(t, err)

// 		mockService := new(mocks.MockTransactionService)
// 		mockService.On("AcceptAskFromServer", mock.AnythingOfType("model.TransactionModel")).Return(errors.New("error"))
// 		transactioncontroller := transactioncontroller.NewController(mockService)
// 		r := gin.Default()
// 		r.POST("/accept-ask", transactioncontroller.AcceptAsk)

// 		rr := httptest.NewRecorder()

// 		req, err := http.NewRequest(http.MethodPost, "/accept-ask", payload)
// 		assert.NoError(t, err)
// 		req.Header.Set("Content-Type", writer.FormDataContentType())

// 		r.ServeHTTP(rr, req)

// 		assert.Equal(t, http.StatusInternalServerError, rr.Code)
// 		mockService.AssertExpectations(t)
// 	})
// }

func TestGetAsk(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockResp := []model.TransactionModel{
			{
				Id: "",
			},
		}

		mockService := new(mocks.MockTransactionService)
		mockService.On("GetAsk", mock.Anything, mock.AnythingOfType("int")).Return(mockResp, nil)
		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.GET("/get-ask", transactioncontroller.GetAsk)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, "/get-ask", nil)
		assert.NoError(t, err)
		req.Header.Set("address", "")
		req.Header.Set("status", "0")

		r.ServeHTTP(rr, req)

		respBody, err := json.Marshal(mockResp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockService.AssertExpectations(t)
	})

	t.Run("not success", func(t *testing.T) {
		mockService := new(mocks.MockTransactionService)
		mockService.On("GetAsk", mock.Anything, mock.AnythingOfType("int")).Return(nil, errors.New("error"))
		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.GET("/get-ask", transactioncontroller.GetAsk)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, "/get-ask", nil)
		assert.NoError(t, err)
		req.Header.Set("address", "")
		req.Header.Set("status", "0")

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockService.AssertExpectations(t)
	})
}

func TestGetGive(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockResp := []model.TransactionModel{
			{
				Id: "",
			},
		}

		mockService := new(mocks.MockTransactionService)
		mockService.On("GetGive", mock.Anything, mock.AnythingOfType("int")).Return(mockResp, nil)
		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.GET("/get-give", transactioncontroller.GetGive)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, "/get-give", nil)
		assert.NoError(t, err)
		req.Header.Set("address", "")
		req.Header.Set("address", "0")

		r.ServeHTTP(rr, req)

		respBody, err := json.Marshal(mockResp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockService.AssertExpectations(t)
	})

	t.Run("not success", func(t *testing.T) {
		mockService := new(mocks.MockTransactionService)
		mockService.On("GetGive", mock.Anything, mock.AnythingOfType("int")).Return(nil, errors.New("error"))
		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.GET("/get-give", transactioncontroller.GetGive)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, "/get-give", nil)
		assert.NoError(t, err)
		req.Header.Set("address", "")
		req.Header.Set("address", "0")

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockService.AssertExpectations(t)
	})
}

func TestGetASignedTransaction(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockResp := model.TransactionModel{
			Id: "",
		}

		mockService := new(mocks.MockTransactionService)
		mockService.On("GetASignedTransaction", mock.Anything).Return(mockResp, nil)
		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.GET("/get-signed-transaction/:tx-id", transactioncontroller.GetASignedTransaction)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, "/get-signed-transaction/123", nil)
		assert.NoError(t, err)

		r.ServeHTTP(rr, req)

		respBody, err := json.Marshal(mockResp)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockService.AssertExpectations(t)
	})

	t.Run("not success", func(t *testing.T) {
		mockService := new(mocks.MockTransactionService)
		mockService.On("GetASignedTransaction", mock.Anything).Return(nil, errors.New("error"))
		transactioncontroller := transactioncontroller.NewController(mockService)
		r := gin.Default()
		r.GET("/get-signed-transaction/:tx-id", transactioncontroller.GetASignedTransaction)

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, "/get-signed-transaction/123", nil)
		assert.NoError(t, err)

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockService.AssertExpectations(t)
	})
}
