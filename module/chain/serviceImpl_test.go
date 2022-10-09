package chain

import (
	"context"
	"github.com/stretchr/testify/assert"
	"hamster-client/module/account"
	"hamster-client/module/application"
	"hamster-client/module/graph"
	"hamster-client/module/p2p"
	"hamster-client/module/resource"
	"hamster-client/module/wallet"
	"hamster-client/utils"
	"path/filepath"
	"testing"
	"time"

	"github.com/mitchellh/go-homedir"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestStartQueue(t *testing.T) {
	configPath, _ := homedir.Expand("~/.link/")
	db, err := gorm.Open(sqlite.Open(filepath.Join(configPath, "link.db")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(
		&account.Account{},
		&p2p.P2pConfig{},
		&resource.Resource{},
		&wallet.Wallet{},
		&application.Application{},
		&graph.GraphParameter{},
	)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	httpUtil := utils.NewHttp()
	accountService := account.NewServiceImpl(ctx, db, httpUtil)
	applicationService := application.NewServiceImpl(ctx, db)
	p2pService := p2p.NewServiceImpl(ctx, db, &accountService)
	si := NewServiceImpl(db, &applicationService, &p2pService)
	err = si.StartQueue(123)
	if err != nil {
		panic(err)
	}
	time.Sleep(60 * time.Second)
}

func TestDb(t *testing.T) {
	configPath, _ := homedir.Expand("~/.link/")
	db, err := gorm.Open(sqlite.Open(filepath.Join(configPath, "link.db")), &gorm.Config{})
	assert.NoError(t, err)
	d
}
