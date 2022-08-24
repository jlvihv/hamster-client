package account

import (
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	PublicKey string `json:"publicKey"`
	WsUrl     string `json:"wsUrl"`
}

type Service interface {
	GetAccount() (Account, error)
	SaveAccount(account *Account)
	GetSubstrateApi() (*gsrpc.SubstrateAPI, error)
}
