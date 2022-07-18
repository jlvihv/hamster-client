package wallet

import (
	"context"
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
func (w *ServiceImpl) SaveWallet(address string, json string) (bool, error) {
	u, err := w.GetWallet()
	if err != nil {
		return false, err
	}
	//save or update account
	u.Address = address
	u.AddressJson = json
	var wallet Wallet
	wallet.Address = u.Address
	wallet.AddressJson = u.AddressJson
	w.db.Save(&wallet)
	return true, nil
}

// DeleteWallet delete wallet information
func (w *ServiceImpl) DeleteWallet() (bool, error) {
	err := w.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Wallet{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
