package requestdomain

import "signature-app/database/model"

type SendTransactionRequest struct {
	PrivateKey string `json:"private_key" form:"private_key"`
	model.TransactionModel
}

type GetAskRequest struct {
	Address string `json:"address"`
	Status  int    `json:"status"`
}
