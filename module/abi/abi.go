package abi

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"hamster-client/module/abi/ecr20"
	"hamster-client/module/abi/stake_distribution_proxy"
	"hamster-client/module/abi/stake_proxy_factory"
	"math/big"
)

const (
	ecr20ContractAddress                  = "0x5c946740441C12510a167B447B7dE565C20b9E3C"
	stakeDistributionProxyContractAddress = "0x2d44C0e097F6cD0f514edAC633d82E01280B4A5c"
	stakeProxyFactoryContractAddress      = "0x1625649b8Fa14A17F93CfEFA6E9285b206a2243A"
)

func GetEthClient(url string) (*ethclient.Client, error) {
	return ethclient.Dial(url)
}

func GetEthAddress(hexStr string) common.Address {
	address := common.HexToAddress(hexStr)
	return address
}

func GetPrivateKey(privateStr string) *ecdsa.PrivateKey {
	var e ecdsa.PrivateKey
	e.D, _ = new(big.Int).SetString(privateStr, 16)
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	return &e
}

func GetPrivateKeyWithMnemonic(mnemonic string) (*ecdsa.PrivateKey, error) {
	var e ecdsa.PrivateKey
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		fmt.Println("get wallet err,err is: ", err)
		return &e, err
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		fmt.Println("get wallet account err,err is: ", err)
		return &e, err
	}

	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		fmt.Println("get private key failed, ", err)
		return &e, err
	}
	return privateKey, nil
}

func GetPrivateKeyHexStringWithMnemonic(mnemonic string) (string, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return "", err
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return "", err
	}
	return account.Address.Hex(), nil
}

// Ecr20AbiApprove call ecr20Abi.approve(address,  amount)
func Ecr20AbiApprove(stakingAddress common.Address, backend bind.ContractBackend, chainID *big.Int, amount *big.Int, privateKey *ecdsa.PrivateKey) error {
	if amount.Cmp(big.NewInt(100000)) == -1 {
		return errors.New("amount must be greater than or equal to 100000")
	}
	cli, err := ecr20.NewEcr20(GetEthAddress(ecr20ContractAddress), backend)
	if err != nil {
		return err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}
	_, err = cli.Approve(opts, stakingAddress, amount)
	if err != nil {
		return err
	}
	return nil
}

// StakeDistributionProxyAbiStaking call stakeDistributionProxyAbi.staking(amount)
func StakeDistributionProxyAbiStaking(stakingAddress common.Address, backend bind.ContractBackend, chainID *big.Int, amount *big.Int, privateKey *ecdsa.PrivateKey) error {
	cli, err := stake_distribution_proxy.NewStakeDistributionProxy(stakingAddress, backend)
	if err != nil {
		return err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}
	_, err = cli.Staking(opts, amount)
	if err != nil {
		return err
	}
	return nil
}

// StakeDistributionProxyAbiGetStakingAmount call stakeDistributionProxyAbi.getStakingAmount() // call
func StakeDistributionProxyAbiGetStakingAmount(ctx context.Context, contractAddress common.Address, caller bind.ContractCaller) (*big.Int, error) {
	cli, err := stake_distribution_proxy.NewStakeDistributionProxyCaller(contractAddress, caller)
	if err != nil {
		return nil, err
	}
	opts := &bind.CallOpts{
		//From:    from,
		Context: ctx,
	}
	return cli.GetStakingAmount(opts)
}

// StakeProxyFactoryAbiCreateStakingContract call stakeProxyFactoryAbi.createStakingContract(address)
func StakeProxyFactoryAbiCreateStakingContract(senderAddress common.Address, backend bind.ContractBackend, chainID *big.Int, privateKey *ecdsa.PrivateKey) error {
	cli, err := stake_proxy_factory.NewStakeProxyFactory(GetEthAddress(stakeProxyFactoryContractAddress), backend)
	if err != nil {
		return err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}
	_, err = cli.CreateStakingContract(opts, senderAddress)
	if err != nil {
		return err
	}
	return nil
}

// StakeProxyFactoryAbiGetStakingAddress call stakeProxyFactoryAbi.getStakingAddress(address)  //call
func StakeProxyFactoryAbiGetStakingAddress(ctx context.Context, senderAddress common.Address, caller bind.ContractCaller) (common.Address, error) {
	cli, err := stake_proxy_factory.NewStakeProxyFactoryCaller(GetEthAddress(stakeProxyFactoryContractAddress), caller)
	if err != nil {
		return common.Address{}, err
	}
	opts := &bind.CallOpts{
		Context: ctx,
	}
	return cli.GetStakingAddress(opts, senderAddress)
}
