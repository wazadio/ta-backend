package requestdomain

type CreateAccountRequest struct{}

type ImportAccountRequest struct {
	Mnemonic string `json:"mnemonic"`
}

type CreateAdminAccountRequest struct {
	Password string `json:"password"`
}

type NewTokenRequest struct {
	Address string `json:"address"`
	Token   string `json:"token"`
}
