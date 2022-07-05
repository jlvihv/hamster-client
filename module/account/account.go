package account

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	PublicKey  string `json:"publicKey"`
	WsUrl      string `json:"wsUrl"`
	OrderIndex int    `json:"orderIndex"`
	PeerId     string `json:"peerId"`
}

type Service interface {
	GetAccount() (Account, error)
	SaveAccount(account *Account)
	SaveOrderIndex(orderIndex int)
}
