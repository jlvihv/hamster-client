package app

import (
	"context"
	"fmt"
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
func (w *Wallet) GetWalletInfo() (wallet.WalletVo, error) {
	info, err := w.walletService.GetWallet()
	return info, err
}

// SaveWallet save wallet information
func (w *Wallet) SaveWallet(address string, json string, passphrase string) (bool, error) {

	fmt.Println("SaveWallet: ", passphrase)
	return w.walletService.SaveWallet(address, json, passphrase)
}

// DeleteWallet delete wallet information
func (w *Wallet) DeleteWallet() (bool, error) {
	return w.walletService.DeleteWallet()
}
