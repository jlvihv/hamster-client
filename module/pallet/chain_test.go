package pallet

import (
	"bufio"
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"hamster-client/utils"
	"math/big"
	"os"
	"testing"
	"time"
)

func TestWaitResource(t *testing.T) {
	substrateApi, err := gsrpc.NewSubstrateAPI("ws://183.66.65.207:49944")
	assert.NoError(t, err)
	meta, err := substrateApi.RPC.State.GetMetadataLatest()
	assert.NoError(t, err)

	print := func() {

		filePath := "/tmp/golang.txt"
		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("文件打开失败", err)
		}
		//及时关闭file句柄
		defer f.Close()

		mapData, _ := GetResourceList(meta, substrateApi, func(resource *ComputingResource) bool {
			//return resource.Status.IsUnused
			return true
		})

		now := time.Now()
		for _, val := range mapData {
			//写入文件时，使用带缓存的 *Writer
			write := bufio.NewWriter(f)
			write.WriteString(now.Format("2006-01-02 15:04:05") + "\t" + utils.AccountIdToAddress(val.AccountId) + "\t" + val.Status.toString() + "\r\n")
			//Flush将缓存的文件真正写入到文件中
			write.Flush()
		}
	}

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				print()
			}
		}
	}()

	time.Sleep(3 * 12 * time.Hour)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func TestGetEvent(t *testing.T) {
	const blockNumber = 1854
	substrateApi, err := gsrpc.NewSubstrateAPI("ws://59.80.40.149:9944")
	assert.NoError(t, err)
	meta, err := substrateApi.RPC.State.GetMetadataLatest()
	assert.NoError(t, err)

	event, err := GetEvent(substrateApi, meta, blockNumber)

	fmt.Println(err)
	fmt.Println(len(event.ResourceOrder_CreateOrderSuccess))
}

func TestSimpleTransfer(t *testing.T) {
	// This sample shows how to create a transaction to make a transfer from one an account to another.

	// Instantiate the API
	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}

	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		panic(err)
	}

	// Create a call, transferring 12345 units to Bob
	bob, err := types.NewMultiAddressFromHexAccountID("0x8eaf04151687736326c9fea17e25fc5287613693c912909cb226aa4794f26a48")
	if err != nil {
		panic(err)
	}

	// 1 unit of transfer
	bal, ok := new(big.Int).SetString("100000000000000", 10)
	if !ok {
		panic(fmt.Errorf("failed to convert balance"))
	}

	c, err := types.NewCall(meta, "Balances.transfer", bob, types.NewUCompact(bal))
	if err != nil {
		panic(err)
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(c)

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		panic(err)
	}

	rv, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		panic(err)
	}

	key, err := types.CreateStorageKey(meta, "System", "Account", signature.TestKeyringPairAlice.PublicKey)
	if err != nil {
		panic(err)
	}

	var accountInfo types.AccountInfo
	ok, err = api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil || !ok {
		panic(err)
	}

	nonce := uint32(accountInfo.Nonce)
	o := types.SignatureOptions{
		BlockHash:          genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
		SpecVersion:        rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(100),
		TransactionVersion: rv.TransactionVersion,
	}

	// Sign the transaction using Alice's default account
	err = ext.Sign(signature.TestKeyringPairAlice, o)
	if err != nil {
		panic(err)
	}

	// Send the extrinsic
	_, err = api.RPC.Author.SubmitExtrinsic(ext)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Balance transferred from Alice to Bob: %v\n", bal.String())
}

func TestGetOrder(t *testing.T) {

	substrateApi, err := gsrpc.NewSubstrateAPI("ws://59.80.40.149:9944")
	assert.NoError(t, err)
	meta, err := substrateApi.RPC.State.GetMetadataLatest()
	assert.NoError(t, err)

	resource, err := GetResource(3, meta, substrateApi)
	assert.NoError(t, err)
	log.Info("Resource: ", resource)

	order, err := GetOrder(types.NewU64(2), meta, substrateApi)
	assert.NoError(t, err)

	log.Info("order: ", order)

	agreement, err := GetRentalAgreement(types.NewU64(1), meta, substrateApi)
	assert.NoError(t, err)
	log.Info("rentalAgreement: ", agreement)
}
