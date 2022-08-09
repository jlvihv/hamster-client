package graph

import (
	"errors"
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"hamster-client/module/pallet"
	"hamster-client/module/queue"
	"hamster-client/utils"
)

type PullImageJob struct {
	err         error
	ProviderApi string
}

func (j *PullImageJob) Run(sc chan queue.StatusCode) (queue.StatusCode, error) {
	sc <- queue.Running
	url := fmt.Sprintf("%s/api/v1/thegraph/pullImage", j.ProviderApi)
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

type WaitResourceJob struct {
	err  error
	api  *gsrpc.SubstrateAPI
	meta *types.Metadata
}

func (j *WaitResourceJob) Run(sc chan queue.StatusCode) (queue.StatusCode, error) {
	sc <- queue.Running
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
				panic(err)
			}
			err = pallet.CallAndWatch(j.api, c, j.meta, func(header *types.Header) error {
				fmt.Println("资源占用成功，资源号：", val.Index)
				return nil
			})
			if err == nil {
				continue
			}
		}
	}

	return queue.Succeeded, nil
}

func (j *WaitResourceJob) Name() string {
	return "Resource Waiting"
}
func (j *WaitResourceJob) Error() error {
	return j.err
}

func NewWaitResourceJob(api *gsrpc.SubstrateAPI) (*WaitResourceJob, error) {

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		return nil, err
	}

	return &WaitResourceJob{
		api:  api,
		meta: meta,
	}, nil
}
