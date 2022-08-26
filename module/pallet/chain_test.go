package pallet

import (
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
