package responsedomain

import (
	"signature-app/domain"
	data "signature-app/service/smart_contract/contract"
)

type DeployContractResponse struct {
	Id              string `json:"id"`
	ContractAddress string `json:"contract_address"`
	From            string `json:"from"`
}

type GetDokumenResponse struct {
	Dokumen []string `json:"data"`
}

type GetIdentitasResponse struct {
	Identitas []data.DataIdentitas `json:"data"`
}

type AddDokumenResponse struct {
	IsSuccess bool `json:"is_success"`
}

type AddIdentitasResponse struct {
	IsSuccess bool `json:"is_success"`
}

type GetETHResponse struct {
	domain.BaseField
	ETHReceived float64 `json:"eth_received"`
}
