package wallet

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Address     string `json:"address"`     //account address
	AddressJson string `json:"addressJson"` //json file information
	Passphrase  string `json:"passphrase"`
}

type WalletVo struct {
	Address     string `json:"address"`     //account address
	AddressJson string `json:"addressJson"` //json file information
}

type WalletJson struct {
	Encoded  string             `json:"encoded"`
	Encoding WalletJsonEncoding `json:"encoding"`
	Address  string             `json:"address"`
}

type WalletJsonEncoding struct {
	Content []string `json:"content"`
	Type    []string `json:"type"`
	Version string   `json:"version"`
}
type WalletJsonMeta struct {
	IsHardware  bool     `json:"isHardware"`
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	WhenCreated uint     `json:"whenCreated"`
}

type Service interface {
	// GetWallet get wallet information
	GetWallet() (WalletVo, error)
	// SaveWallet save wallet information
	SaveWallet(address string, json string, passphrase string) (bool, error)
	// DeleteWallet delete wallet information
	DeleteWallet() (bool, error)
	// GetWalletKeypair get wallet keypair struct
	GetWalletKeypair() (signature.KeyringPair, error)
}
