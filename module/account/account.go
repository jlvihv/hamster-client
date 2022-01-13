package account

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	PublicKey string `json:"publicKey"`
}

type Service interface {
	GetAccount() (Account, error)
	SaveAccount(account *Account)
}
