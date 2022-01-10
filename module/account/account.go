package account

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username  string `json:"username"`
	UserId    string `json:"userId"`
	PublicKey string `json:"publicKey"`
	Phone     string `json:"phone"`
}

type Service interface {
	GetCode(phone string) error
	GetAccount() (Account, error)
	SaveAccount(account *Account)
	DeleteAccount()
	Login(mobile string, smsCode string) (*Account, error)
	Logout()
}
