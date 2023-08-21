package requestdomain

type CreateAccountRequest struct{}

type ImportAccountRequest struct {
	Mnemonic string `json:"mnemonic"`
}

type CreateAdminAccountRequest struct {
	Password string `json:"password"`
}
