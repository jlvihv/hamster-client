package chainjob

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"hamster-client/config"
	"hamster-client/module/chainhelper"
	"hamster-client/module/pallet"
	"hamster-client/module/queue"
	"hamster-client/utils"
	"time"

	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type VersionVo struct {
	Version string `json:"version"`
}

type WaitResourceJob struct {
	appID      int
	si         queue.StatusInfo
	helper     chainhelper.Helper
	bond       bool
	meta       *types.Metadata
	deployType int
}

func NewWaitResourceJob(appID int, helper chainhelper.Helper, deployType int) queue.Job {
	return &WaitResourceJob{
		appID:      appID,
		helper:     helper,
		deployType: deployType,
	}
}

func (j *WaitResourceJob) InitStatus() {
	j.si.Name = "Wait Resource"
	j.si.Status = queue.None
}

func (j *WaitResourceJob) Run(si chan queue.StatusInfo) (queue.StatusInfo, error) {
	log.Info("WaitResourceJob running...")
	j.si.Status = queue.Running
	si <- j.si

	api, err := j.helper.Account().GetSubstrateApi()
	if err != nil {
		j.si.Status = queue.Failed
		j.si.Error = err.Error()
		si <- j.si
		return j.si, err
	}
	log.Info("get substrate api succeed")

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		log.Errorf("get meta data latest error: %v", err)
		j.si.Status = queue.Failed
		j.si.Error = err.Error()
		si <- j.si
		return j.si, err
	}
	j.meta = meta
	log.Infof("get meta data latest succeed")

	pair, err := j.helper.Wallet().GetWalletKeypair()
	if err != nil {
		j.si.Status = queue.Failed
		j.si.Error = "WALLET_LOAD_ERROR"
		si <- j.si
		return j.si, err
	}
	if pair.Address == "" {
		j.si.Status = queue.Failed
		j.si.Error = errors.New("get Keypair failed").Error()
		si <- j.si
		return j.si, err
	}
	log.Info("get wallet keypair succeed")

	if !j.bond {
		// 100 unit
		err = pallet.Bond(api, j.meta, 100000000000000, pair)
		if err != nil {
			log.Errorf("bond error: %s", err.Error())
			j.si.Status = queue.Failed
			j.si.Error = "WALLET_LOAD_ERROR"
			si <- j.si
			return j.si, err
		}
		log.Infof("bond succeed")
	}

	for i := 0; i < 60; i++ {

		mapData, err := pallet.GetResourceList(
			j.meta,
			api,
			func(resource *pallet.ComputingResource) bool {
				return resource.Status.IsUnused
			},
		)
		if err != nil {
			j.si.Status = queue.Failed
			j.si.Error = err.Error()
			si <- j.si
			return j.si, err
		}
		log.Infof("get resource list succeed")
		log.Infof("resource number: %d", len(mapData))

		failSet := make(map[int]string)

		for _, val := range mapData {

			if _, isMapContainsKey := failSet[int(val.Index)]; isMapContainsKey {
				continue
			}

			if val.Status.IsUnused {
				resourceIndex := val.Index
				log.Infof("use resource: %d", resourceIndex)
				data, err := j.helper.App().QueryApplicationById(j.appID)
				if err != nil {
					log.Errorf("get application failed: %v", err)
					continue
				}

				// check p2p connection can be connected
				port := j.helper.App().QueryNextP2pPort()
				err = j.helper.P2p().
					LinkByProtocol(config.ProviderProtocol, port, string(val.PeerId))

				if err != nil {
					log.Errorf("create p2p network forward failed: %v", err)
					failSet[int(val.Index)] = "fail"
					_, _ = j.helper.P2p().Close(fmt.Sprintf("/p2p/%s", string(val.PeerId)))
					continue
				}

				// check http is Ok
				url := fmt.Sprintf("http://localhost:%d/version", port)
				log.Infof("check http url connect: %s", url)
				req := utils.NewHttp().NewRequest()
				resp, httperr := req.Get(url)
				if httperr != nil {
					log.Errorf("check http url connect failed: %v", httperr)
					_, _ = j.helper.P2p().Close(fmt.Sprintf("/p2p/%s", string(val.PeerId)))
					failSet[int(val.Index)] = "fail"
					continue
				}
				log.Infof("connect succeed")

				var version VersionVo
				bandErr := json.Unmarshal(resp.Body(), &version)
				if bandErr != nil {
					_, _ = j.helper.P2p().Close(fmt.Sprintf("/p2p/%s", string(val.PeerId)))
					failSet[int(val.Index)] = "fail"
					continue
				}
				log.Infof("provider version: %s", version.Version)

				c, err := types.NewCall(
					j.meta,
					"ResourceOrder.create_order_info",
					resourceIndex,
					types.NewU32(uint32(data.LeaseTerm)),
					"",
					types.NewU32(uint32(j.deployType)),
				)
				if err != nil {
					fmt.Println(err.Error())
					_, _ = j.helper.P2p().Close(fmt.Sprintf("/p2p/%s", string(val.PeerId)))
					failSet[int(val.Index)] = "fail"
					continue
				}
				var events pallet.MyEventRecords
				err = pallet.CallAndWatch(api, c, j.meta, func(header *types.Header) error {
					log.Infof("use resource succeed, resource index: %d, block number: %d", resourceIndex, header.Number)
					// get order index
					e, err := pallet.GetEvent(api, j.meta, uint64(header.Number))
					events = *e
					return err
				}, pair)

				if err != nil {
					log.Errorf("watch event error: %v", err)
					_, _ = j.helper.P2p().Close(fmt.Sprintf("/p2p/%s", string(val.PeerId)))
					failSet[int(val.Index)] = "fail"
					continue
				}

				var orderIndex types.U64
				for _, e := range events.ResourceOrder_CreateOrderSuccess {
					if e.ResourceIndex == resourceIndex {
						orderIndex = e.OrderIndex
					}
				}

				fmt.Println("query orderIndex: ", orderIndex)

				// check provider exec order
				number, err := pallet.GetBlockNumber(api)
				if err != nil {
					_ = pallet.CancelOrder(api, j.meta, pair, int(orderIndex))
					_, _ = j.helper.P2p().Close(fmt.Sprintf("/p2p/%s", string(val.PeerId)))
					continue
				}

				accept := false

				for i := number; i < number+10; i++ {
					events, err := pallet.GetEvent(api, j.meta, i)
					if err != nil {
						continue
					}
					if len(events.ResourceOrder_OrderExecSuccess) > 0 {
						for _, e := range events.ResourceOrder_OrderExecSuccess {
							if e.OrderIndex == orderIndex {
								accept = true
								fmt.Println("confirm order : ", orderIndex)
								break
							}
						}
						if accept {
							break
						}
					}
					time.Sleep(time.Second * 6)
				}

				if !accept {
					_ = pallet.CancelOrder(api, j.meta, pair, int(orderIndex))
					_, _ = j.helper.P2p().Close(fmt.Sprintf("/p2p/%s", string(val.PeerId)))
					failSet[int(val.Index)] = "fail"
					continue
				}

				_ = j.helper.App().UpdatePeerIdAndOrderIndex(
					j.appID,
					int(orderIndex),
					int(resourceIndex),
					string(val.PeerId),
				)

				err = j.helper.App().UpdateApplicationP2pForwardPort(j.appID, port)
				if err != nil {
					log.Errorf("query p2p max port fail: %v", err)
					continue
				}

				j.si.Status = queue.Succeeded
				j.si.Error = ""
				si <- j.si
				return j.si, nil
			}
		}

		log.Infof("current number of retries: %d, retry after 30s...", i)
		time.Sleep(time.Second * 30)

	}
	j.si.Status = queue.Failed
	j.si.Error = "NO_RESOURCE_TO_USE"
	si <- j.si
	return j.si, errors.New("NO_RESOURCE_TO_USE")
}

func (j *WaitResourceJob) Status() queue.StatusInfo {
	return j.si
}
