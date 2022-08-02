package pallet

import (
	ctx "context"
	"errors"
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/decred/base58"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hamster-client/config"
	"hamster-client/module/account"
	"hamster-client/module/application"
	"hamster-client/module/deploy"
	"hamster-client/module/keystorage"
	"hamster-client/module/p2p"
	"hamster-client/module/wallet"
	"strconv"
)

type ChainListener struct {
	cancel func()
	ctx2   ctx.Context
}

func NewChainListener() *ChainListener {
	return &ChainListener{}
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

func (c *ChainListener) watchEvent(ctx ctx.Context, db *gorm.DB, keyStorageService keystorage.Service, deployService deploy.Service) {
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
						var wallet wallet.Wallet
						walletResult := db.First(&wallet)
						if walletResult.Error == nil {
							publicKey, _ := AddressToPublicKey(wallet.Address)
							key, err := types.CreateStorageKey(meta, "ResourceOrder", "ApplyUsers", publicKey)
							if err != nil {
								log.Error(err)
							}
							var orderIndex types.U64
							ok, err := api.RPC.State.GetStorageLatest(key, &orderIndex)
							if err != nil {
								log.Error(err)
							}
							log.Info(ok)
							result := db.First(&user)
							if result.Error == nil {
								if e.OrderIndex == orderIndex {
									fmt.Println(user.OrderIndex)
									user.OrderIndex = int(orderIndex)
									user.PeerId = e.PeerId
									db.Save(&user)
									// Query whether there is an application waiting for resources
									var data application.Application
									result := db.Where("status = ? ", config.WAIT_RESOURCE).First(&data).Error
									if result == nil {
										jsonData := keyStorageService.Get("graph_" + strconv.Itoa(int(data.ID)))
										deployService.DeployTheGraph(int(data.ID), jsonData)
									}
								}
							}
						}
					}
				}
			}

		}
	}
}

func (c *ChainListener) StartListen(db *gorm.DB, keyStorageService keystorage.Service, deployService deploy.Service) error {
	if c.cancel != nil {
		c.cancel()
	}
	c.ctx2, c.cancel = ctx.WithCancel(ctx.Background())
	go c.watchEvent(c.ctx2, db, keyStorageService, deployService)
	return nil
}

func (c *ChainListener) CancelListen() {
	if c.cancel != nil {
		c.cancel()
		c.cancel = nil
	}
}

// AddressToPublicKey Convert address to public key
func AddressToPublicKey(address string) ([]byte, error) {
	if len(address) < 33 {
		return []byte{}, errors.New("帐号格式不合法")
	}
	return base58.Decode(address)[1:33], nil
}
