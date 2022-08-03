package wallet

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Address     string `json:"address"`     //account address
	AddressJson string `json:"addressJson"` //json file information
}

type WalletVo struct {
	Address     string `json:"address"`     //account address
	AddressJson string `json:"addressJson"` //json file information
}

type Service interface {
	// GetWallet get wallet information
	GetWallet() (WalletVo, error)
	// SaveWallet save wallet information
	SaveWallet(address string, json string) (bool, error)
	// DeleteWallet delete wallet information
	DeleteWallet() (bool, error)
}
