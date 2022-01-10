package wallet

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Address     string `json:"address"`      //account address
	AddressJson string `json:"address_json"` //json file information
}

type Service interface {
	// GetWallet get wallet information
	GetWallet() (Wallet, error)
	// SaveWallet save wallet information
	SaveWallet(address string, json string) (*Wallet, error)
	// DeleteWallet delete wallet information
	DeleteWallet()
}
