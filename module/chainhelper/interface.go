package chainhelper

import (
	"hamster-client/module/account"
	"hamster-client/module/application"
	"hamster-client/module/keystorage"
	"hamster-client/module/p2p"
	"hamster-client/module/queue"
	"hamster-client/module/wallet"

	"gorm.io/gorm"
)

type Helper interface {
	GetP2pForwardPort(appID int) (int, error)

	DB() *gorm.DB
	App() application.Service
	Account() account.Service
	P2p() p2p.Service
	Wallet() wallet.Service
	Queue() queue.Service
	KS() keystorage.Service
}
