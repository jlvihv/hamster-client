package abi

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"hamster-client/module/abi/ecr20"
	"hamster-client/module/abi/stake_distribution_proxy"
	"hamster-client/module/abi/stake_proxy_factory"
	"math/big"
)

const (
	ecr20ContractAddress                  = "0x54Fe55d5d255b8460fB3Bc52D5D676F9AE5697CD"
	stakeDistributionProxyContractAddress = "0x2d44C0e097F6cD0f514edAC633d82E01280B4A5c"
	stakeProxyFactoryContractAddress      = "0xeFF0ed9Fc8276Fcf4eda2e012dD065A3DC18591D"
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

func GetPrivateKeyWithMnemonicAndPassword(mnemonic, password string) (*ecdsa.PrivateKey, error) {
	privateStr, err := GetPrivateKeyHexStringWithMnemonicAndPassword(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return GetPrivateKey(privateStr), nil
}

func GetPrivateKeyHexStringWithMnemonicAndPassword(mnemonic, password string) (string, error) {
	seed := bip39.NewSeed(mnemonic, password)
	seedHexStr := hex.EncodeToString(seed)
	return PrivateKeyBySeed(seedHexStr)
}

// PrivateKeyBySeed returns private key from seed.
func PrivateKeyBySeed(seed string) (string, error) {
	bytes, err := hex.DecodeString(seed)
	if err != nil {
		return "", err
	}

	// Generate a new master node using the seed.
	masterKey, err := bip32.NewMasterKey(bytes)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H
	acc44H, err := masterKey.NewChildKey(bip32.FirstHardenedChild + 44)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H/60H
	acc44H60H, err := acc44H.NewChildKey(bip32.FirstHardenedChild + 60)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H/60H/0H
	acc44H60H0H, err := acc44H60H.NewChildKey(bip32.FirstHardenedChild + 0)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H/60H/0H/0
	acc44H60H0H0, err := acc44H60H0H.NewChildKey(0)
	if err != nil {
		return "", err
	}

	// This gives the path: m/44H/60H/0H/0/0
	//acc44H60H0H00, err := acc44H60H0H0.NewChildKey(0)
	//if err != nil {
	//	return "", err
	//}

	return hex.EncodeToString(acc44H60H0H0.Key), nil
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
