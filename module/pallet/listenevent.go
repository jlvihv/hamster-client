package pallet

import (
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hamster-client/module/account"
	"hamster-client/module/p2p"
)

func WatchEvent(db *gorm.DB) {
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
		set := <-sub.Chan()
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
