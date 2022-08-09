package v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"hamster-client/config"
	ethAbi "hamster-client/module/abi"
	"hamster-client/module/deploy"
	"hamster-client/module/keystorage"
	"hamster-client/module/account"
	"hamster-client/module/application"
	"hamster-client/module/p2p"
	"hamster-client/module/pallet"
	"hamster-client/module/queue"
	"hamster-client/utils"
	"strconv"
	"time"
)

type PullImageJob struct {
	err                error
	applicationId      int
	applicationService application.Service
}

func (j *PullImageJob) Run(sc chan queue.StatusCode) (queue.StatusCode, error) {
	sc <- queue.Running
	vo, err := j.applicationService.QueryApplicationById(j.applicationId)
	if err != nil {
		fmt.Println("query application fail, err is :", err)
		sc <- queue.Failed
		return queue.Failed, err
	}

	url := fmt.Sprintf("http://localhost:%d/api/v1/thegraph/pullImage", vo.P2pForwardPort)
	response, err := utils.NewHttp().NewRequest().Post(url)
	if err != nil {
		j.err = err
		return queue.Failed, err
	}

	if response.IsSuccess() {
		sc <- queue.Succeeded
		return queue.Succeeded, nil
	} else {
		sc <- queue.Failed
		return queue.Failed, errors.New("api response fail")
	}
}

func (j *PullImageJob) Name() string {
	return "Service Pull"
}

func (j *PullImageJob) Error() error {
	return j.err
}

func NewPullImageJob(service application.Service, applicationId int) PullImageJob {
	return PullImageJob{
		applicationId:      applicationId,
		applicationService: service,
	}
}

type WaitResourceJob struct {
	err                error
	api                *gsrpc.SubstrateAPI
	meta               *types.Metadata
	accountService     account.Service
	applicationService application.Service
	p2pService         p2p.Service
	applicationId      int
}

func (j *WaitResourceJob) Run(sc chan queue.StatusCode) (queue.StatusCode, error) {

	sc <- queue.Running

	for i := 0; i < 60; i++ {

		mapData, err := pallet.GetResourceList(j.meta, j.api, func(resource *pallet.ComputingResource) bool {
			return resource.Status.IsUnused
		})
		if err != nil {
			sc <- queue.Failed
			j.err = err
			return queue.Failed, err
		}

		fmt.Println("可用资源数：", len(mapData))

		for _, val := range mapData {

			if val.Status.IsUnused {
				fmt.Println("发现未使用资源，占用。资源号：", val.Index)
				c, err := types.NewCall(j.meta, "ResourceOrder.create_order_info", val.Index, types.NewU32(10), "")
				if err != nil {
					continue
				}
				err = pallet.CallAndWatch(j.api, c, j.meta, func(header *types.Header) error {
					fmt.Println("资源占用成功，资源号：", val.Index)
					return nil
				})
				if err != nil {
					continue
				}

				ac, _ := j.accountService.GetAccount()
				ac.PeerId = string(val.PeerId)
				j.accountService.SaveAccount(&ac)

				port := j.applicationService.QueryNextP2pPort()

				err = j.applicationService.UpdateApplicationP2pForwardPort(j.applicationId, port)
				if err != nil {
					fmt.Println("query p2p max port fail")
				}
				err = j.p2pService.LinkByProtocol("/x/provider", port, ac.PeerId)

				if err != nil {
					fmt.Println("create p2p network forward fail")
				}

				sc <- queue.Succeeded
				return queue.Succeeded, nil
			}
		}

		time.Sleep(time.Second * 30)

	}
	return queue.Failed, errors.New("NO_RESOURCE_TO_USE")
}

func (j *WaitResourceJob) Name() string {
	return "Resource Waiting"
}
func (j *WaitResourceJob) Error() error {
	return j.err
}

func NewWaitResourceJob(api *gsrpc.SubstrateAPI, accountService account.Service, applicationService application.Service, p2pService p2p.Service, applicationId int) (*WaitResourceJob, error) {

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, err
	}

	return &WaitResourceJob{
		api:                api,
		meta:               meta,
		accountService:     accountService,
		applicationService: applicationService,
		p2pService:         p2pService,
		applicationId:      applicationId,
	}, nil
}

type GraphStakingJob struct {
	err               error
	id                int
	keyStorageService keystorage.Service
}

func NewGraphStakingJob(id int, keyStorageService keystorage.Service) GraphStakingJob {
	return GraphStakingJob{id: id, keyStorageService: keyStorageService}
}

func (g *GraphStakingJob) Run(sc chan queue.StatusCode) (queue.StatusCode, error) {
	sc <- queue.Running
	//获取质押金额以及助记词
	var param deploy.DeployParameter
	jsonParam := g.keyStorageService.Get("graph_" + strconv.Itoa(g.id))
	if err := json.Unmarshal([]byte(jsonParam), &param); err != nil {
		g.err = err
		sc <- queue.Failed
		return queue.Failed, err
	}
	if param.Initialization.AccountMnemonic == "" {
		err := errors.New("Saved mnemonic is empty")
		g.err = err
		sc <- queue.Failed
		return queue.Failed, err
	}
	//根据助记词获取address
	addr, err := ethAbi.GetAccountAddress(param.Initialization.AccountMnemonic)
	if err != nil {
		g.err = err
		sc <- queue.Failed
		return queue.Failed, err
	}
	address := ethAbi.GetEthAddress(addr)
	client, err := ethAbi.GetEthClient(config.EndpointUrl)
	if err != nil {
		g.err = err
		sc <- queue.Failed
		return queue.Failed, err
	}
	stakingAddress, err := ethAbi.StakeProxyFactoryAbiGetStakingAddress(context.Background(), address, client)
	if err != nil {
		g.err = err
		sc <- queue.Failed
		return queue.Failed, err
	}
	if stakingAddress == ethAbi.GetEthAddress("0") {
	}
	return queue.Succeeded, nil
}

func (g *GraphStakingJob) Name() string {
	return "TheGraph Staking"
}

func (g *GraphStakingJob) Error() error {
	return g.err
}
