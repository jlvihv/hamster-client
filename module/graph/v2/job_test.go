package v2

import (
	"context"
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/mitchellh/go-homedir"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hamster-client/module/account"
	"hamster-client/module/application"
	"hamster-client/module/graph"
	"hamster-client/module/p2p"
	queue2 "hamster-client/module/queue"
	"hamster-client/module/resource"
	"hamster-client/module/wallet"
	"hamster-client/utils"
	"path/filepath"
	"testing"
	"time"
)

func TestDeploy(t *testing.T) {
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

	applicationId := 11
	ctx := context.Background()
	httpUtil := utils.NewHttp()
	accountService := account.NewServiceImpl(ctx, db, httpUtil)
	applicationService := application.NewServiceImpl(ctx, db)
	p2pService := p2p.NewServiceImpl(ctx, db)
	pullJob := NewPullImageJob(&applicationService, applicationId)
	substrateApi, _ := gsrpc.NewSubstrateAPI("ws://183.66.65.207:49944")

	job2, _ := NewWaitResourceJob(substrateApi, &accountService, &applicationService, &p2pService, applicationId)
	queue := queue2.NewQueue("1", job2, &pullJob)

	channel := make(chan struct{})
	go queue.Start(channel)
	go func() {
		for {
			time.Sleep(time.Second)
			info, err := queue.GetStatus()
			if err != nil {
				t.Error(err)
			}
			for _, v := range info {
				fmt.Print(v, "; ")
			}
			fmt.Println()
		}
	}()
	// wait
	<-channel
	// view status
	info, _ := queue.GetStatus()

	fmt.Println(info)
}
