package p2p

import (
	"context"
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
	"hamster-client/config"
	//"hamster-client/module/account"
	"os/exec"
)

type ServiceImpl struct {
	ctx       context.Context
	db        *gorm.DB
	p2pClient *P2pClient
}

func NewServiceImpl(ctx context.Context, db *gorm.DB) ServiceImpl {
	return ServiceImpl{ctx: ctx, db: db}
}

// initialize p2p link
func (s *ServiceImpl) getP2pClient() (*P2pClient, error) {
	if s.p2pClient != nil {
		return s.p2pClient, nil
	}
	p2pConfig, err := s.GetSetting()
	if err != nil {
		runtime.LogError(s.ctx, "getP2pClient GetSetting is error %s"+err.Error())
		return nil, err
	}
	if p2pConfig.PrivateKey == "" {
		runtime.LogWarning(s.ctx, "getP2pClient p2p config is null")
		return nil, err
	}
	//perform p2p client initialization link
	return s.initP2pClient(p2pConfig.Port, p2pConfig.PrivateKey)
}

func (s *ServiceImpl) initP2pClient(port int, privateKey string) (*P2pClient, error) {
	var nodes []string
	api := CreateApi()
	meta, _ := api.RPC.State.GetMetadataLatest()
	key, err := types.CreateStorageKey(meta, "Gateway", "Gateways")
	api.RPC.State.GetStorageLatest(key, &nodes)
	host, dht, err := MakeRoutedHost(port, privateKey, nodes)
	if err != nil {
		return nil, err
	}
	p2p := MakeIpfsP2p(&host)
	s.p2pClient = &P2pClient{
		Host: host,
		P2P:  p2p,
		DHT:  dht,
	}
	return s.p2pClient, nil
}

// Link p2p linking
func (s *ServiceImpl) Link(port int, peerId string) error {
	client, err := s.getP2pClient()
	if err != nil {
		return err
	}
	err = client.Forward(port, peerId)
	if err != nil {
		return err
	}
	return nil
}

// Close Close Link
func (s *ServiceImpl) Close(target string) (int, error) {
	client, err := s.getP2pClient()
	if err != nil {
		return 0, err
	}
	return client.Close(target)
}

// Destroy p2plink destruction
func (s *ServiceImpl) Destroy() error {
	if s.p2pClient == nil {
		return nil
	}
	err := s.p2pClient.Destroy()
	if err != nil {
		return err
	}
	s.p2pClient = nil
	return nil
}

//GetLinks get a list of links
func (s *ServiceImpl) GetLinks() *[]LinkInfo {
	runtime.LogWarning(s.ctx, "GetLinks start")
	var links []LinkInfo
	client, err := s.getP2pClient()
	if err != nil {
		return &links
	}
	outPut := client.List()
	for _, value := range outPut.Listeners {
		linkInfo := LinkInfo{Protocol: value.Protocol, ListenAddress: value.ListenAddress, TargetAddress: value.TargetAddress}
		err := client.CheckForwardHealth(value.TargetAddress)
		linkInfo.Status = err == nil
		runtime.LogInfo(s.ctx, fmt.Sprintf("GetLinks %s\n", linkInfo.Status))
		links = append(links, linkInfo)
	}
	return &links
}

//InitSetting p2p parameter configuration
func (s *ServiceImpl) InitSetting() error {
	runtime.LogInfo(s.ctx, "InitSetting start")
	p2pConfig, err := s.GetSetting()
	port := config.Port
	if err != nil {
		runtime.LogError(s.ctx, fmt.Sprintf("InitSetting error :%s\n ", err))
		identity, err := CreateIdentity()
		if err != nil {
			runtime.LogError(s.ctx, fmt.Sprintf("InitSetting error :%s\n ", err))
			return err
		}
		for {
			err := portInUse(port)
			if err != nil {
				break
			} else {
				port = port + 1
			}
		}
		p2pConfig.Port = port
		p2pConfig.PrivateKey = identity.PrivKey
		p2pConfig.PeerId = identity.PeerID
		s.db.Save(&p2pConfig)
	}
	_, err = s.initP2pClient(port, p2pConfig.PrivateKey)
	if err != nil {
		runtime.LogError(s.ctx, fmt.Sprintf(""))
		runtime.LogError(s.ctx, fmt.Sprintf("InitSetting error :%s\n ", err))
		return err
	}
	runtime.LogInfo(s.ctx, "InitSetting success")
	return nil
}

//GetSetting query p2p parameter configuration information
func (s *ServiceImpl) GetSetting() (P2pConfig, error) {
	var p2pConfig P2pConfig
	result := s.db.First(&p2pConfig)
	if result.Error != nil {
		runtime.LogError(s.ctx, fmt.Sprintf("GetSetting error %s\n", result.Error))
	}
	return p2pConfig, result.Error
}

//determine if there is a 4001 locally
func portInUse(portNumber int) error {
	cmdStr := fmt.Sprintf("netstat -nlp | grep :%d", portNumber)
	cmd := exec.Command("bash", "-c", cmdStr)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
