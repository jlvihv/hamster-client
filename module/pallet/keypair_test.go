package pallet

import (
	"encoding/json"
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/stretchr/testify/assert"
	"hamster-client/module/wallet"
	"math/big"
	"testing"
)

const ()

func TestKeyPair(t *testing.T) {
	encoded := "cIQvvGijrfWCEF0pXvOsclvuDq+KHdIrpKsnFA0t4mwAgAAAAQAAAAgAAAAcszXmqxjls8lM/BRmBYfmhy/1niUWez0xnR7WbAqlpHETbGdxmvKUsYLlKUZTQ6pPpxVN76CZiGLON0icRbqMcesFj9GmIUtabT/Exyxg/VGzmwYT8gbdcDlyJ3fb63+6COYo1F4ObHYxDKJkSML11n1NUo6nKKH/sSBKSwXnShjUx5NWtwXzR7YVIZLFICixoK2gBLftusdYBM+V"

	passphrase := "123456"

	keypair, err := KeyringPairFromEncoded(encoded, passphrase, 42)
	assert.NoError(t, err)
	fmt.Println(keypair)
	sampleTransform(keypair)
}

func TestFromMonomeric(t *testing.T) {
	const monomeric = "roast double swamp expand around element conduct table prize stomach brief meat"
	keypair, err := signature.KeyringPairFromSecret(monomeric, 42)
	assert.NoError(t, err)
	fmt.Println(keypair)
	//sampleTransform(keypair)
}

func TestKeypairAddress(t *testing.T) {
	const addressJson = "{\"encoded\":\"cIQvvGijrfWCEF0pXvOsclvuDq+KHdIrpKsnFA0t4mwAgAAAAQAAAAgAAAAcszXmqxjls8lM/BRmBYfmhy/1niUWez0xnR7WbAqlpHETbGdxmvKUsYLlKUZTQ6pPpxVN76CZiGLON0icRbqMcesFj9GmIUtabT/Exyxg/VGzmwYT8gbdcDlyJ3fb63+6COYo1F4ObHYxDKJkSML11n1NUo6nKKH/sSBKSwXnShjUx5NWtwXzR7YVIZLFICixoK2gBLftusdYBM+V\",\"encoding\":{\"content\":[\"pkcs8\",\"sr25519\"],\"type\":[\"scrypt\",\"xsalsa20-poly1305\"],\"version\":\"3\"},\"address\":\"5DUQrTMDdCukEptpUpHxWobT3ZsrhwQqSLhnX8p9Xp1nSNay\",\"meta\":{\"isHardware\":false,\"name\":\"krp\",\"tags\":[],\"whenCreated\":1660186771936}}"
	const password = "123456"
	var walletJson wallet.WalletJson
	err := json.Unmarshal([]byte(addressJson), &walletJson)
	assert.NoError(t, err)

	assert.Equal(t, "5DUQrTMDdCukEptpUpHxWobT3ZsrhwQqSLhnX8p9Xp1nSNay", walletJson.Address)

	keyringPair, err := KeyringPairFromEncoded(walletJson.Encoded, password, 42)
	assert.NoError(t, err)

	assert.Equal(t, "5DUQrTMDdCukEptpUpHxWobT3ZsrhwQqSLhnX8p9Xp1nSNay", keyringPair.Address)

	data, err := keyringPair.Sign([]byte("hello"))
	assert.NoError(t, err)
	//
	//data := []byte{44, 155, 144,   4,   2, 219, 167, 166, 152,  80, 209,
	//	147, 209, 104, 191, 212, 210,  73, 150, 166,  57,  98,
	//	167,  22,  67,  11, 238, 239, 195,  87,  75,  33, 149,
	//	187, 188, 189, 190, 162, 111, 101, 119,  21, 134, 165,
	//	170, 166, 219, 161, 237,  43, 127,  76, 100, 117,  73,
	//	108,  71,  89,  96,   1,  50, 179, 145, 136,
	//}
	bool, err := keyringPair.Verify([]byte("hello"), data)
	assert.NoError(t, err)

	fmt.Println("signed data:", data)
	assert.True(t, bool)

}

func sampleTransform(pair signature.KeyringPair) {
	// This sample shows how to create a transaction to make a transfer from one an account to another.

	// Instantiate the API
	api, err := gsrpc.NewSubstrateAPI("ws://183.66.65.207:49944")
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

	key, err := types.CreateStorageKey(meta, "System", "Account", pair.PublicKey)
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
	err = ext.Sign(pair, o)
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
