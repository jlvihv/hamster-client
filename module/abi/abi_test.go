package abi

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"testing"
)

const (
	senderAddress = "0xa7Aa9A6D3D682b29D9C9e30Bb206c961d860861C"
	endpointUrl   = "https://rinkeby.infura.io/v3/bab2a1a435b04c07a488d847cf6788f7"
	privateKey    = "8d290072a3517f47e6c7c8db5ecd1876b6ef6576096797a7d6470022e7b3910a"
)

var (
	sender  = GetEthAddress(senderAddress)
	privKey = GetPrivateKey(privateKey)
	client  = getClient()
)

func getClient() *ethclient.Client {
	client, err := GetEthClient(endpointUrl)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func TestAll(t *testing.T) {
	var stakingAddress common.Address
	var err error
	var stakingAmount = big.NewInt(100000)

	// 获取质押合约地址
	stakingAddress, err = StakeProxyFactoryAbiGetStakingAddress(context.Background(), sender, client)
	if err != nil {
		// 如果获取质押合约地址失败，则创建质押合约
		err = StakeProxyFactoryAbiCreateStakingContract(sender, client, big.NewInt(4), privKey)
		if err != nil {
			t.Fatal(err)
		}
		// 再次获取质押合约地址
		stakingAddress, err = StakeProxyFactoryAbiGetStakingAddress(context.Background(), sender, client)
		if err != nil {
			t.Fatal(err)
		}
	}
	fmt.Println("get staking stakingAddress:", stakingAddress)

	// 批准
	err = Ecr20AbiApprove(stakingAddress, client, big.NewInt(4), stakingAmount, privKey)
	if err != nil {
		t.Fatal(err)
	}

	// 从质押地址获取质押金额
	amount, err := StakeDistributionProxyAbiGetStakingAmount(context.Background(), stakingAddress, client)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get staking amount:", amount)
	if amount.Cmp(big.NewInt(0)) == 0 {
		// 质押
		err = StakeDistributionProxyAbiStaking(stakingAddress, client, big.NewInt(4), stakingAmount, privKey)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestGetPrivateKeyWithMnemonicAndPassword(t *testing.T) {
	mnemonic := ""
	password := ""
	privateKeyString, err := GetPrivateKeyHexStringWithMnemonicAndPassword(mnemonic, password)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(privateKeyString)
}
