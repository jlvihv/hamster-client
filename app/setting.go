package app

import (
	"context"
	"hamster-client/module/account"
	"hamster-client/module/p2p"
)

type Config struct {
	PublicKey string
	Port      int
	PeerId    string
}

type Setting struct {
	ctx            context.Context
	p2pService     p2p.Service
	accountService account.Service
}

func NewSettingApp(service p2p.Service, accountService account.Service) Setting {
	return Setting{
		p2pService:     service,
		accountService: accountService,
	}
}

func (s *Setting) WailsInit(ctx context.Context) error {
	s.ctx = ctx
	return nil
}

// GetSetting view configuration information
func (s *Setting) GetSetting() (*Config, error) {
	config := &Config{}
	//query the public key in the user information
	info, err := s.accountService.GetAccount()
	if err != nil {
		return config, err
	}
	config.PublicKey = info.PublicKey
	//query the setting information in the p2p setting
	p2pConfig, err := s.p2pService.GetSetting()
	if err != nil {
		return config, nil
	}
	config.Port = p2pConfig.Port
	config.PeerId = p2pConfig.PeerId
	return config, nil
}

// Setting set public key information
func (s *Setting) Setting(publicKey string) (bool, error) {
	accountInfo, _ := s.accountService.GetAccount()

	//set user public key
	accountInfo.PublicKey = publicKey
	s.accountService.SaveAccount(&accountInfo)
	return true, nil
}

//InitP2pSetting initialize p2p settings
func (s *Setting) InitP2pSetting() (bool, error) {
	//initial configuration
	err := s.p2pService.InitSetting()
	if err != nil {
		return false, err
	}
	return true, nil
}
