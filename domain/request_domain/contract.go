package requestdomain

type DeployContractRequest struct {
	PrivateKey string `json:"private_key"`
}

type AddDokumenRequest struct {
	NamaDokumen string `json:"nama_dokumen"`
	PrivateKey  string `json:"private_key"`
}

type AddIdentitasRequest struct {
	Nama       string `json:"nama"`
	PrivateKey string `json:"private_key"`
	Nomor      int    `json:"nomor"`
	Alamat     string `json:"alamat"`
	Status     string `json:"status"`
}
