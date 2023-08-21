package domain

type BaseField struct {
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
}
