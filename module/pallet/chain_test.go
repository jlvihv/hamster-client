package pallet

import (
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryResource(t *testing.T) {
	substrateApi, err := gsrpc.NewSubstrateAPI("ws://183.66.65.207:49944")
	//assert.NoError(t, err)
	meta, err := substrateApi.RPC.State.GetMetadataLatest()
	//assert.NoError(t, err)

	mapData, err := GetResourceList(meta, substrateApi, nil)

	assert.NoError(t, err)

	for _, val := range mapData {

		if val.Status.IsUnused {
			fmt.Println("发现未使用资源，占用。资源号：", val.Index)
			c, err := types.NewCall(meta, "ResourceOrder.create_order_info", val.Index, types.NewU32(10), "")
			if err != nil {
				panic(err)
				return
			}
			err = callAndWatch(substrateApi, c, meta, func(header *types.Header) error {
				fmt.Println("资源占用成功，资源号：", val.Index)
				return nil
			})
			if err == nil {
				return
			}
		}
	}
}

func TestWaitResource(t *testing.T) {
	substrateApi, err := gsrpc.NewSubstrateAPI("ws://183.66.65.207:49944")
	assert.NoError(t, err)
	meta, err := substrateApi.RPC.State.GetMetadataLatest()
	assert.NoError(t, err)

	mapData, err := GetResourceList(meta, substrateApi, func(resource *ComputingResource) bool {
		return resource.Status.IsUnused
	})

	fmt.Println("可用资源数：", len(mapData))

}
