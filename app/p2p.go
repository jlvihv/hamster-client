package app

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"hamster-client/module/p2p"
	"time"
)

type P2p struct {
	ctx       context.Context
	p2pServer p2p.Service
}

func NewP2pApp(service p2p.Service) P2p {
	return P2p{
		p2pServer: service,
	}
}

func (s *P2p) WailsInit(ctx context.Context) error {
	s.ctx = ctx
	go func() {
		for {
			runtime.EventsEmit(s.ctx, "Links", s.p2pServer.GetLinks())
			time.Sleep(5 * time.Second)
		}
	}()
	return nil
}

// IsP2PSetting determine whether p2p information is configured
func (s *P2p) IsP2PSetting() bool {
	config, err := s.p2pServer.GetSetting()
	if err != nil {
		return false
	}
	return config.PrivateKey != ""
}

// Link p2p link
func (s *P2p) Link(port int, peerId string) (bool, error) {
	//make a p2p link
	err := s.p2pServer.Link(port, peerId)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CloseLink close link
func (s *P2p) CloseLink(target string) (int, error) {
	//disconnect p2p
	return s.p2pServer.Close(target)
}

// GetLinkStatus query p2p link status
func (s *P2p) GetLinkStatus() *[]p2p.LinkInfo {
	return s.p2pServer.GetLinks()
}

// WailsShutdown close link
func (s *P2p) WailsShutdown() {
	_ = s.p2pServer.Destroy()
}

func (s *P2p) JudgeP2pReconnection() bool {
	return s.p2pServer.JudgeP2pReconnection()
}

func (s *P2p) ReconnectionProLink() (bool, error) {
	return s.p2pServer.ReconnectionProLink()
}
