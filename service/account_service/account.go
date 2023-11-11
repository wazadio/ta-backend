package accountservice

import (
	"context"
	"signature-app/database/repository"
	requestdomain "signature-app/domain/request_domain"
	responsedomain "signature-app/domain/response_domain"
	"signature-app/helper"
	kafkaclient "signature-app/kafka-client"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

type accountService struct {
	cl    *ethclient.Client
	ctx   context.Context
	Db    *repository.Database
	Kafka kafkaclient.KafkaClientInterface
}

func NewAccountService(ctx context.Context, cl *ethclient.Client, db *repository.Database) *accountService {
	return &accountService{
		ctx: ctx,
		cl:  cl,
		Db:  db,
	}
}

func (s *accountService) CreateNewAccount() (*responsedomain.AccountDetailResponse, error) {
	mnemonic, err := hdwallet.NewMnemonic(128)
	if err != nil {
		return nil, err
	}

	return s.ImportAccount(mnemonic)
}

func (s *accountService) ImportAccount(mnemonic string) (*responsedomain.AccountDetailResponse, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, err
	}

	help := helper.NewHelper(s.cl, s.ctx)
	err = help.SendIniatialETH(account.Address.Hex())
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

func (s *accountService) AddNewTokenFromServer(data requestdomain.NewTokenRequest) error {
	// msg, err := json.Marshal(data)
	// if err != nil {
	// 	return err
	// }
	// err = s.Kafka.ProduceMessageToken(msg)
	// log.Println("error sending message : ", err)
	// return err
	err := s.Db.AddNewToken(data.Address, data.Token)

	return err
}
