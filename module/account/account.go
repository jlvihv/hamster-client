package account

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	PublicKey string `json:"publicKey"`
	Nodes     string `json:"nodes"`
}

type Service interface {
	GetAccount() (Account, error)
	SaveAccount(account *Account)
}
