package app

import (
	"github.com/wailsapp/wails"
	"link-server/module/account"
)

type Account struct {
	log            *wails.CustomLogger
	runtime        *wails.Runtime
	accountService account.Service
}

func NewAccountApp(service account.Service) Account {
	return Account{
		accountService: service,
	}
}

func (s *Account) WailsInit(runtime *wails.Runtime) error {
	s.runtime = runtime
	s.log = runtime.Log.New("Account")
	return nil
}

// GetAccountInfo get user information
func (s *Account) GetAccountInfo() (account.Account, error) {
	info, err := s.accountService.GetAccount()
	return info, err
}

// IsAccount determine if the user exists
func (s *Account) IsAccount() bool {
	_, err := s.accountService.GetAccount()
	if err != nil {
		return false
	}
	return true
}

// IsAccountSetting determine whether to configure user information
func (s *Account) IsAccountSetting() bool {
	info, err := s.accountService.GetAccount()
	if err != nil {
		return false
	}
	return info.PublicKey != ""
}
