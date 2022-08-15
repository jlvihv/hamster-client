package wallet

import (
	"context"
	"encoding/json"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	ctx context.Context
	db  *gorm.DB
}

func NewServiceImpl(ctx context.Context, db *gorm.DB) ServiceImpl {
	return ServiceImpl{ctx, db}
}

// GetWallet  get wallet information
func (w *ServiceImpl) GetWallet() (WalletVo, error) {
	var wallet Wallet
	var data WalletVo
	result := w.db.First(&wallet)
	if result.Error != nil {
		runtime.LogError(w.ctx, "GetWallet error")
	}
	data.Address = wallet.Address
	data.AddressJson = wallet.AddressJson
	return data, nil
}

// SaveWallet save wallet information
func (w *ServiceImpl) SaveWallet(address string, json string, passphrase string) (bool, error) {
	var wallet Wallet
	result := w.db.First(&wallet)
	if result.Error != nil {
		wallet = Wallet{}
	}
	//save or update account
	wallet.Address = address
	wallet.AddressJson = json
	wallet.Passphrase = passphrase
	result = w.db.Save(&wallet)
	return result.Error == nil, result.Error
}

// DeleteWallet delete wallet information
func (w *ServiceImpl) DeleteWallet() (bool, error) {
	err := w.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Wallet{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (w *ServiceImpl) GetWalletJson() (WalletJson, string, error) {
	var wallet Wallet
	result := w.db.First(&wallet)
	if result.Error != nil {
		runtime.LogError(w.ctx, "GetWallet error")
		return WalletJson{}, "", result.Error
	}

	var walletJson WalletJson
	err := json.Unmarshal([]byte(wallet.AddressJson), &walletJson)
	return walletJson, wallet.Passphrase, err
}
