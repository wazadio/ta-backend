package model

import "time"

type TransactionModel struct {
	Id          string    `json:"id"`
	TxId        string    `json:"tx_id"`
	FromAddress string    `json:"from_address"`
	FromName    string    `json:"from_name"`
	ToAddress   string    `json:"to_address"`
	ToName      string    `json:"to_name"`
	Status      string    `json:"status"`
	Data        string    `json:"data"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
