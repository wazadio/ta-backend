package mocks

import (
	"mime/multipart"
	"signature-app/database/model"
	responsedomain "signature-app/domain/response_domain"

	"github.com/stretchr/testify/mock"
)

type MockTransactionService struct {
	mock.Mock
}

func (m *MockTransactionService) GetAllTransactions() (txs []responsedomain.TransactionResponse, err error) {
	ret := m.Called()

	if ret.Get(0) != nil {
		txs = ret.Get(0).([]responsedomain.TransactionResponse)
	}

	if ret.Get(1) != nil {
		err = ret.Get(1).(error)
	}

	return
}

func (m *MockTransactionService) SendTransaction(strPrivateKey string, body model.TransactionModel, pdf *multipart.FileHeader) (*responsedomain.TransactionResponse, error) {
	ret := m.Called(strPrivateKey, body, pdf)

	var r0 *responsedomain.TransactionResponse
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*responsedomain.TransactionResponse)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (m *MockTransactionService) AddAsk(payload model.TransactionModel) error {
	ret := m.Called(payload)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *MockTransactionService) AddAskFromServers(payload model.TransactionModel) error {
	ret := m.Called(payload)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *MockTransactionService) GetAsk(address string, status []int) (data []model.TransactionModel, err error) {
	ret := m.Called(address, status)

	if ret.Get(0) != nil {
		data = ret.Get(0).([]model.TransactionModel)
	}

	if ret.Get(1) != nil {
		err = ret.Get(1).(error)
	}

	return
}

func (m *MockTransactionService) GetOneAsk(id string) (data []model.TransactionModel, err error) {
	ret := m.Called(id)

	if ret.Get(0) != nil {
		data = ret.Get(0).([]model.TransactionModel)
	}

	if ret.Get(1) != nil {
		err = ret.Get(1).(error)
	}

	return
}

func (m *MockTransactionService) GetGive(address string, status []int) (data []model.TransactionModel, err error) {
	ret := m.Called(address, status)

	if ret.Get(0) != nil {
		data = ret.Get(0).([]model.TransactionModel)
	}

	if ret.Get(1) != nil {
		err = ret.Get(1).(error)
	}

	return
}

func (m *MockTransactionService) AcceptAskFromServer(data model.TransactionModel) error {
	ret := m.Called(data)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *MockTransactionService) GetASignedTransaction(txId string) (data model.TransactionModel, err error) {
	ret := m.Called(txId)

	if ret.Get(0) != nil {
		data = ret.Get(0).(model.TransactionModel)
	}

	if ret.Get(1) != nil {
		err = ret.Get(1).(error)
	}

	return
}

func (m *MockTransactionService) AddAskDirect(strPrivateKey string, data model.TransactionModel) error {
	ret := m.Called(data)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

func (m *MockTransactionService) SignDirect(strPrivateKey, id string) error {
	ret := m.Called(strPrivateKey, id)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
