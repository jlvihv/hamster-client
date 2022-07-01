package wallet

import (
	"context"
	"fmt"
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
func (w *ServiceImpl) GetWallet() (Wallet, error) {
	var wallet Wallet
	fmt.Println(w.db)
	result := w.db.First(&wallet)
	if result.Error != nil {
		runtime.LogError(w.ctx, "GetWallet error")
	}
	return wallet, result.Error
}

// SaveWallet save wallet information
func (w *ServiceImpl) SaveWallet(address string, json string) error {
	u, err := w.GetWallet()
	if err != nil {
		return err
	}
	//save or update account
	u.Address = address
	u.AddressJson = json
	w.db.Save(&u)
	return nil
}

// DeleteWallet delete wallet information
func (w *ServiceImpl) DeleteWallet() {
	w.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Wallet{})
}
