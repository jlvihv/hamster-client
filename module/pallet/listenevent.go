package pallet

import (
	ctx "context"
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hamster-client/module/account"
	"hamster-client/module/p2p"
)

type ChainListener struct {
	db     *gorm.DB
	cancel func()
	ctx2   ctx.Context
}

func NewChainListener(db *gorm.DB) *ChainListener {
	return &ChainListener{
		db: db,
	}
}

func (c *ChainListener) WatchEvent(db *gorm.DB, ctx ctx.Context) {
	api := p2p.CreateApi()
	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		panic(err)
	}
	// Subscribe to system events via storage
	key, err := types.CreateStorageKey(meta, "System", "Events", nil)
	if err != nil {
		panic(err)
	}

	sub, err := api.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关闭旧协程")
			return
		case set := <-sub.Chan():
			fmt.Println("listen block number：", set.Block.Hex())
			for _, chng := range set.Changes {
				if !types.Eq(chng.StorageKey, key) || !chng.HasStorageData {
					// skip, we are only interested in events with content
					continue
				}
				// Decode the event records
				evt := MyEventRecords{}
				storageData := chng.StorageData
				meta, err := api.RPC.State.GetMetadataLatest()
				err = types.EventRecordsRaw(storageData).DecodeEventRecords(meta, &evt)
				if err != nil {
					fmt.Println(err)
					log.Error(err)
					continue
				}

				for _, e := range evt.ResourceOrder_FreeResourceProcessed {
					// order successfully created
					var user account.Account
					result := db.First(&user)
					if result.Error == nil {
						if int(e.OrderIndex) == user.OrderIndex {
							fmt.Println(user.OrderIndex)
							user.PeerId = e.PeerId
							db.Save(&user)
						}
					}
				}
			}
		}

	}
}

func (c *ChainListener) watchEvent(ctx ctx.Context) {
	api := p2p.CreateApi()
	if api != nil {
		meta, err := api.RPC.State.GetMetadataLatest()
		if err != nil {
			panic(err)
		}
		// Subscribe to system events via storage
		key, err := types.CreateStorageKey(meta, "System", "Events", nil)
		if err != nil {
			panic(err)
		}

		sub, err := api.RPC.State.SubscribeStorageRaw([]types.StorageKey{key})
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()
		for {
			select {
			case <-ctx.Done():
				return
			case set := <-sub.Chan():
				fmt.Println("listen block number：", set.Block.Hex())
				for _, chng := range set.Changes {
					if !types.Eq(chng.StorageKey, key) || !chng.HasStorageData {
						// skip, we are only interested in events with content
						continue
					}
					// Decode the event records
					evt := MyEventRecords{}
					storageData := chng.StorageData
					meta, err := api.RPC.State.GetMetadataLatest()
					err = types.EventRecordsRaw(storageData).DecodeEventRecords(meta, &evt)
					if err != nil {
						fmt.Println(err)
						log.Error(err)
						continue
					}

					for _, e := range evt.ResourceOrder_FreeResourceProcessed {
						// order successfully created
						var user account.Account
						result := c.db.First(&user)
						if result.Error == nil {
							if int(e.OrderIndex) == user.OrderIndex {
								fmt.Println(user.OrderIndex)
								user.PeerId = e.PeerId
								c.db.Save(&user)
							}
						}
					}
				}
			}

		}
	}
}

func (c *ChainListener) StartListen() error {
	if c.cancel != nil {
		c.cancel()
	}
	c.ctx2, c.cancel = ctx.WithCancel(ctx.Background())
	go c.watchEvent(c.ctx2)
	return nil
}

func (c *ChainListener) CancelListen() {
	if c.cancel != nil {
		c.cancel()
		c.cancel = nil
	}
}
