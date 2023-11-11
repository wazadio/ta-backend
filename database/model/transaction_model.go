package model

type TransactionModel struct {
	Id           string `json:"id" form:"id"`
	TxId         string `json:"tx_id" form:"tx_id"`
	FromAddress  string `json:"from_address" form:"from_address"`
	FromName     string `json:"from_name" form:"from_name"`
	ToAddress    string `json:"to_address" form:"to_address"`
	ToName       string `json:"to_name" form:"to_name"`
	Status       int8   `json:"status"`
	DocumentName string `json:"document_name" form:"document_name"`
	Description  string `json:"description" form:"description"`
	Data         string `json:"data" form:"data"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type DataContract struct {
	FromAddress  string `json:"from_address"`
	FromName     string `json:"from_name"`
	ToAddress    string `json:"to_address"`
	ToName       string `json:"to_name"`
	DocumentName string `json:"document_name"`
	TimeStamp    string `json:"time_stamp"`
}
