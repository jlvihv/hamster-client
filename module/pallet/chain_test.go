package pallet

import (
	"bufio"
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/stretchr/testify/assert"
	"hamster-client/utils"
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
