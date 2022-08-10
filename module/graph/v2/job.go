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
	"hamster-client/module/account"
	"hamster-client/module/application"
	"hamster-client/module/deploy"
	"hamster-client/module/keystorage"
	"hamster-client/module/p2p"
	"hamster-client/module/pallet"
	"hamster-client/module/queue"
	"hamster-client/utils"
	"math/big"
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
				data, err := j.applicationService.QueryApplicationById(j.applicationId)
				if err != nil {
					fmt.Println("get application failed,err is: ", err)
					continue
				}
				c, err := types.NewCall(j.meta, "ResourceOrder.create_order_info", val.Index, types.NewU32(uint32(data.LeaseTerm)), "")
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
	sc <- queue.Failed
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

func NewGraphStakingJob(service keystorage.Service, applicationId int) GraphStakingJob {
	return GraphStakingJob{
		id:                applicationId,
		keyStorageService: service,
	}
}

func (g *GraphStakingJob) Run(sc chan queue.StatusCode) (queue.StatusCode, error) {
	sc <- queue.Running
	//Obtain pledge amount and mnemonic words
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
	//Get address from mnemonic
	addr, err := ethAbi.GetPrivateKeyHexStringWithMnemonic(param.Initialization.AccountMnemonic)
	if err != nil {
		fmt.Println("Get address from mnemonic failed, err is :", err)
		g.err = err
		sc <- queue.Failed
		return queue.Failed, err
	}
	address := ethAbi.GetEthAddress(addr)
	//Get eth client
	client, err := ethAbi.GetEthClient(config.EndpointUrl)
	if err != nil {
		fmt.Println("Get eth client failed, err is :", err)
		g.err = err
		sc <- queue.Failed
		return queue.Failed, err
	}
	// Obtain the agent pledge address
	stakingAddress, err := ethAbi.StakeProxyFactoryAbiGetStakingAddress(context.Background(), address, client)
	if err != nil {
		fmt.Println("Get agent proxy address failed, err is :", err)
		g.err = err
		sc <- queue.Failed
		return queue.Failed, err
	}
	//Get private key from mnemonic
	privateKey, err := ethAbi.GetPrivateKeyWithMnemonic(param.Initialization.AccountMnemonic)
	if err != nil {
		fmt.Println("Get private key from mnemonic failed, err is :", err)
		g.err = err
		sc <- queue.Failed
		return queue.Failed, err
	}
	if stakingAddress == ethAbi.GetEthAddress("0") {
		//Create agent pledge address
		err = ethAbi.StakeProxyFactoryAbiCreateStakingContract(address, client, big.NewInt(4), privateKey)
		if err != nil {
			fmt.Println("Create agent pledge address failed, err is :", err)
			g.err = err
			sc <- queue.Failed
			return queue.Failed, err
		}
		// Get the agent pledge address again
		stakingAddress, err = ethAbi.StakeProxyFactoryAbiGetStakingAddress(context.Background(), address, client)
		if err != nil {
			fmt.Println("Get the agent pledge address again failed, err is :", err)
			g.err = err
			sc <- queue.Failed
			return queue.Failed, err
		}
		//Convert the pledged amount into Wei
		stakingAmount := utils.ToWei18(int64(param.Staking.PledgeAmount))
		// Authorize the agency pledge address
		err = ethAbi.Ecr20AbiApprove(stakingAddress, client, big.NewInt(4), stakingAmount, privateKey)
		if err != nil {
			fmt.Println("approve failed, err is :", err)
			g.err = err
			sc <- queue.Failed
			return queue.Failed, err
		}
		//GRT pledge
		err = ethAbi.StakeDistributionProxyAbiStaking(stakingAddress, client, big.NewInt(4), stakingAmount, privateKey)
		if err != nil {
			fmt.Println("staking failed, err is :", err)
			g.err = err
			sc <- queue.Failed
			return queue.Failed, err
		}
	}
	param.Deployment.IndexerAddress = addr
	param.Staking.AgentAddress = stakingAddress.Hex()
	jsonData, err := json.Marshal(param)
	if err != nil {
		g.err = err
		sc <- queue.Failed
		return queue.Failed, err
	}
	g.keyStorageService.Set("graph_"+strconv.Itoa(int(g.id)), string(jsonData))
	return queue.Succeeded, nil
}

func (g *GraphStakingJob) Name() string {
	return "TheGraph Staking"
}

func (g *GraphStakingJob) Error() error {
	return g.err
}

type ServiceDeployJob struct {
	err               error
	id                int
	deployService     deploy.Service
	keyStorageService keystorage.Service
}

func NewServiceDeployJob(service keystorage.Service, deployService deploy.Service, applicationId int) ServiceDeployJob {
	return ServiceDeployJob{
		id:                applicationId,
		keyStorageService: service,
		deployService:     deployService,
	}
}

func (s *ServiceDeployJob) Run(sc chan queue.StatusCode) (queue.StatusCode, error) {
	sc <- queue.Running
	var param deploy.DeployParameter
	var sendData deploy.DeployParams
	jsonParam := s.keyStorageService.Get("graph_" + strconv.Itoa(s.id))
	if err := json.Unmarshal([]byte(jsonParam), &param); err != nil {
		s.err = err
		sc <- queue.Failed
		return queue.Failed, err
	}
	sendData.EthereumUrl = param.Deployment.EthereumUrl
	sendData.EthereumNetwork = param.Deployment.EthereumNetwork
	sendData.NodeEthereumUrl = param.Deployment.NodeEthereumUrl
	sendData.IndexerAddress = param.Staking.AgentAddress
	sendData.Mnemonic = param.Initialization.AccountMnemonic
	_, err := s.deployService.DeployGraph(s.id, sendData)
	if err != nil {
		fmt.Println("deploy service failed, err is :", err)
		s.err = err
		sc <- queue.Failed
		return queue.Failed, err
	}
	return queue.Succeeded, nil
}

func (s *ServiceDeployJob) Name() string {
	return "Service Deploy"
}

func (s *ServiceDeployJob) Error() error {
	return s.err
}
