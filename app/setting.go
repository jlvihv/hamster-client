package app

import (
	"github.com/wailsapp/wails"
	"link-server/module/account"
	"link-server/module/p2p"
)

type Config struct {
	PublicKey string
	Port      int
	PeerId    string
}

type Setting struct {
	log            *wails.CustomLogger
	p2pService     p2p.Service
	accountService account.Service
}

func NewSettingApp(service p2p.Service, accountService account.Service) Setting {
	return Setting{
		p2pService:     service,
		accountService: accountService,
	}
}

func (s *Setting) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Setting")
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
	accountInfo, err := s.accountService.GetAccount()
	if err != nil {
		return false, err
	}
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
