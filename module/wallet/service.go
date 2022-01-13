package wallet

import (
	"github.com/wailsapp/wails/lib/logger"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	log *logger.CustomLogger
	db  *gorm.DB
}

func NewServiceImpl(db *gorm.DB) ServiceImpl {
	log := logger.NewCustomLogger("Module_Account")
	return ServiceImpl{log, db}
}

// GetWallet  get wallet information
func (w *ServiceImpl) GetWallet() (Wallet, error) {
	var wallet Wallet
	result := w.db.First(&wallet)
	if result.Error != nil {
		w.log.Error("GetWallet error")
	}
	return wallet, result.Error
}

// SaveWallet save wallet information
func (w *ServiceImpl) SaveWallet(address string, json string) (*Wallet, error) {
	u, _ := w.GetWallet()
	//save or update account
	u.Address = address
	u.AddressJson = json
	w.db.Save(&u)
	return &u, nil
}

// DeleteWallet delete wallet information
func (w *ServiceImpl) DeleteWallet() {
	w.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Wallet{})
}
