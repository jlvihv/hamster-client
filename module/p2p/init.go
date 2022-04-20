package p2p

import (
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
)

var api *gsrpc.SubstrateAPI

func init() {
	substrateApi, _ := gsrpc.NewSubstrateAPI(CONFIG_DEFAULT_CHAIN_API)
	api = substrateApi
}
