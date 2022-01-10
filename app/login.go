package app

import (
	"github.com/wailsapp/wails"
	"link-server/module/account"
	"link-server/module/p2p"
	"link-server/module/wallet"
)

type Login struct {
	log            *wails.CustomLogger
	accountService account.Service
	p2pService     p2p.Service
	walletService  wallet.Service
}

func NewLoginApp(service account.Service, p2pService p2p.Service, walletService wallet.Service) Login {
	return Login{
		accountService: service,
		p2pService:     p2pService,
		walletService:  walletService,
	}
}

func (s *Login) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Login")
	return nil
}

// Login user login
func (s *Login) Login(mobile string, smsCode string) (*account.Account, error) {
	return s.accountService.Login(mobile, smsCode)
}

// Logout user logged out
func (s *Login) Logout() {
	//clear user information
	s.accountService.Logout()
	//break the p2p link
	_ = s.p2pService.Destroy()
	//clear import json
	s.walletService.DeleteWallet()
}

// GetCode send the verification code
func (s *Login) GetCode(mobile string) (bool, error) {
	err := s.accountService.GetCode(mobile)
	if err != nil {
		return false, err
	}
	return true, nil
}
