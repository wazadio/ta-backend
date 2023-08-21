package requestdomain

type GetTransactionsRequest struct{}

type SendTransactionRequest struct {
	PrivateKey string `json:"private_key"`
	To         string `json:"to"`
	Data       string `json:"data"`
}
