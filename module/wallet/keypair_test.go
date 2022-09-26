package wallet

import (
	"fmt"
	"github.com/ChainSafe/go-schnorrkel"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/gtank/merlin"
	"github.com/stretchr/testify/assert"
	"hamster-client/utils"
	"testing"
)

func TestKeypair(t *testing.T) {

	address := "5C7b1P6eyY6mhtF9eBs99PYP26vw6z5nXaZh8KycgP1ybmmj"
	fmt.Println(len([]byte(address)))
	keypair, _ := signature.KeyringPairFromSecret("0x972e92880c4538e8cbbc1539c53f96cdc576e0e721c0dd3d9357bf6f0cb41c3b", 42)

	msg := []byte("hello")

	signature, err := keypair.Sign(msg)

	assert.NoError(t, err)

	bool, err := keypair.Verify(msg, signature)

	assert.NoError(t, err)
	assert.True(t, bool)

	var publicKeyData [32]byte
	aToP, _ := utils.AddressToPublicKey(address)
	copy(publicKeyData[:], aToP)

	fmt.Println(publicKeyData)

	keypairFromPublic, _ := schnorrkel.NewPublicKey(publicKeyData)

	var sigs [64]byte
	copy(sigs[:], signature)
	sig := new(schnorrkel.Signature)
	if err := sig.Decode(sigs); err != nil {
		panic(err)
	}
	result, err := keypairFromPublic.Verify(sig, signingContext(msg))
	fmt.Println(result)

}

func signingContext(msg []byte) *merlin.Transcript {
	return schnorrkel.NewSigningContext([]byte("substrate"), msg)
}
