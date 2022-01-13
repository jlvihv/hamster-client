package app

import (
	"github.com/wailsapp/wails"
	"link-server/module/wallet"
)

type Wallet struct {
	log           *wails.CustomLogger
	runtime       *wails.Runtime
	walletService wallet.Service
}

func NewWalletApp(service wallet.Service) Wallet {
	return Wallet{
		walletService: service,
	}
}

func (w *Wallet) WailsInit(runtime *wails.Runtime) error {
	w.runtime = runtime
	w.log = runtime.Log.New("Wallet")
	return nil
}

// GetWalletInfo get user information
func (w *Wallet) GetWalletInfo() (wallet.Wallet, error) {
	info, err := w.walletService.GetWallet()
	return info, err
}

// SaveWallet save wallet information
func (w *Wallet) SaveWallet(address string, json string) (*wallet.Wallet, error) {
	return w.walletService.SaveWallet(address, json)
}

// DeleteWallet delete wallet information
func (w *Wallet) DeleteWallet() {
	w.walletService.DeleteWallet()
}
