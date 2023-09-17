package model

type TransactionModel struct {
	Id          string `json:"id" form:"id"`
	TxId        string `json:"tx_id"`
	FromAddress string `json:"from_address" form:"from_address"`
	FromName    string `json:"from_name" form:"from_name"`
	ToAddress   string `json:"to_address" form:"to_address"`
	ToName      string `json:"to_name" form:"to_name"`
	Status      int8   `json:"status"`
	Data        string `json:"data"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
