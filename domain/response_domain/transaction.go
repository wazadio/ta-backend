package responsedomain

import "time"

type TransactionResponse struct {
	Id        string    `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Data      string    `json:"data"`
	TimeStamp time.Time `json:"time_stamp"`
}
