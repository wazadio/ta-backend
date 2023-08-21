package accountservice

import (
	responsedomain "signature-app/domain/response_domain"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type AccountService struct{}

func NewAccountDetail() *AccountService {
	return &AccountService{}
}

func (s *AccountService) CreateNewAccount() (*responsedomain.AccountDetailResponse, error) {
	mnemonic, err := hdwallet.NewMnemonic(128)
	if err != nil {
		return nil, err
	}

	return s.ImportAccount(mnemonic)
}

func (s *AccountService) ImportAccount(mnemonic string) (*responsedomain.AccountDetailResponse, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, err
	}

	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		return nil, err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	return &responsedomain.AccountDetailResponse{
		Mnemonic:   mnemonic,
		PrivateKey: hexutil.Encode(privateKeyBytes)[2:],
		Address:    account.Address.Hex(),
	}, nil
}
