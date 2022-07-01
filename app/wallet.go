package app

import (
	"context"
	"hamster-client/module/wallet"
)

type Wallet struct {
	ctx           context.Context
	walletService wallet.Service
}

func NewWalletApp(service wallet.Service) Wallet {
	return Wallet{
		walletService: service,
	}
}

func (s *Wallet) WailsInit(ctx context.Context) error {
	s.ctx = ctx
	return nil
}

// GetWalletInfo get user information
func (w *Wallet) GetWalletInfo() (wallet.Wallet, error) {
	info, err := w.walletService.GetWallet()
	return info, err
}

// SaveWallet save wallet information
func (w *Wallet) SaveWallet(address string, json string) error {
	return w.walletService.SaveWallet(address, json)
}

// DeleteWallet delete wallet information
func (w *Wallet) DeleteWallet() {
	w.walletService.DeleteWallet()
}
