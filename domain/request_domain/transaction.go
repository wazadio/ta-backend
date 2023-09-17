package requestdomain

type GetTransactionsRequest struct{}

type SendTransactionRequest struct {
	AskId      string `json:"ask_id"`
	PrivateKey string `json:"private_key"`
	To         string `json:"to"`
	Data       string `json:"data"`
}

type GetAskRequest struct {
	Address string `json:"address"`
	Status  int    `json:"status"`
}
