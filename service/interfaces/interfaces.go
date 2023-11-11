package interfaces

import (
	"mime/multipart"
	"signature-app/database/model"
	requestdomain "signature-app/domain/request_domain"
	responsedomain "signature-app/domain/response_domain"
)

type AccountService interface {
	CreateNewAccount() (*responsedomain.AccountDetailResponse, error)
	ImportAccount(mnemonic string) (*responsedomain.AccountDetailResponse, error)
	AddNewTokenFromServer(data requestdomain.NewTokenRequest) error
}

type TransactionService interface {
	GetAllTransactions() (txs []responsedomain.TransactionResponse, err error)
	SendTransaction(strPrivateKey string, body model.TransactionModel, pdf *multipart.FileHeader) (*responsedomain.TransactionResponse, error)
	AddAsk(payload model.TransactionModel) error
	AddAskFromServers(payload model.TransactionModel) error
	GetAsk(address string, status []int) (data []model.TransactionModel, err error)
	GetOneAsk(id string) (data []model.TransactionModel, err error)
	GetGive(address string, status []int) (data []model.TransactionModel, err error)
	// AcceptAskFromServer(data model.TransactionModel) error
	GetASignedTransaction(txId string) (data model.TransactionModel, err error)
	AddAskDirect(strPrivateKey string, payload model.TransactionModel) error
	SignDirect(strPrivateKey, id string) error
}
