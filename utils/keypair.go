package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/ChainSafe/go-schnorrkel"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gtank/merlin"
	"github.com/vedhavyas/go-subkey"
	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/scrypt"
	"math"
	"strings"
)

const (
	miniSecretKeyLength = 32

	SecLength = 64

	signatureLength = 64

	NonceLength = 24

	publicKeyLength = 32

	publicKeyIndex = 85
)

func KeyringPairFromEncoded(encoded string, passphrase string, network uint8) (signature.KeyringPair, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return signature.KeyringPair{}, err
	}
	encType := []string{"scrypt", "xsalsa20-poly1305"}
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("panic error is: ", err)
		}
	}()
	secretKey, publicKey := decodePair(passphrase, decoded, encType)

	return KeyringPairFromSecretKey(secretKey, publicKey, network)
}

func KeyringPairFromSecretKey(secretKey [64]byte, publicKey []byte, network uint8) (signature.KeyringPair, error) {
	kp, err := NewKeyring(secretKey)

	ss58Address, err := subkey.SS58Address(publicKey, network)
	if err != nil {
		return signature.KeyringPair{}, err
	}
	fmt.Println("ss58Address:", ss58Address)

	return signature.KeyringPair{
		URI:       "",
		Address:   ss58Address,
		PublicKey: publicKey,
		KeyPair:   kp,
	}, nil

}

type keyRing struct {
	seed   []byte
	secret *schnorrkel.SecretKey
	pub    *schnorrkel.PublicKey
}

func NewKeyring(secretKey [64]byte) (*keyRing, error) {
	var key [32]byte
	var nonce [32]byte
	copy(key[:], secretKey[0:0+32])
	copy(nonce[:], secretKey[32:64])

	/*
		pub(crate) fn divide_scalar_bytes_by_cofactor(scalar: &mut [u8; 32]) {
		    let mut low = 0u8;
		    for i in scalar.iter_mut().rev() {
		        let r = *i & 0b00000111; // save remainder
		        *i >>= 3; // divide by 8
		        *i += low;
		        low = r << 5;
		    }
		}
	*/

	key = divideScalarBytesByCofactor(key)
	fmt.Println("key")
	secret := schnorrkel.NewSecretKey(key, nonce)
	pub, err := secret.Public()
	if err != nil {
		return nil, err
	}
	return &keyRing{
		secret: secret,
		pub:    pub,
	}, nil

}

func (kr keyRing) Sign(msg []byte) (signature []byte, err error) {
	sig, err := kr.secret.Sign(signingContext(msg))
	if err != nil {
		return signature, err
	}

	s := sig.Encode()
	return s[:], nil
}

func (kr keyRing) Verify(msg []byte, signature []byte) bool {
	var sigs [signatureLength]byte
	copy(sigs[:], signature)
	sig := new(schnorrkel.Signature)
	if err := sig.Decode(sigs); err != nil {
		return false
	}
	result, err := kr.pub.Verify(sig, signingContext(msg))
	if err != nil {
		return false
	}
	return result
}

func signingContext(msg []byte) *merlin.Transcript {
	return schnorrkel.NewSigningContext([]byte("substrate"), msg)
}

// Public returns the public key in bytes
func (kr keyRing) Public() []byte {
	pub := kr.pub.Encode()
	return pub[:]
}

func (kr keyRing) Seed() []byte {
	return kr.seed
}

func (kr keyRing) AccountID() []byte {
	return kr.Public()
}

func (kr keyRing) SS58Address(network uint8) (string, error) {
	return subkey.SS58Address(kr.AccountID(), network)
}

func (kr keyRing) SS58AddressWithAccountIDChecksum(network uint8) (string, error) {
	return subkey.SS58AddressWithAccountIDChecksum(kr.AccountID(), network)
}

type KeyringPair struct {
	signature.KeyringPair
}

func decodePair(passphrase string, encrypted []byte, encType []string) ([64]byte, []byte) {
	PKCS8_HEADER := []byte{48, 83, 2, 1, 1, 48, 5, 6, 3, 43, 101, 112, 4, 34, 4, 32}
	SEED_OFFSET := len(PKCS8_HEADER)

	decrypted := jsonDecryptData(encrypted, passphrase, encType)

	publicByte := decrypted[publicKeyIndex : publicKeyIndex+publicKeyLength]

	header := decrypted[0:len(PKCS8_HEADER)]
	if !bytes.Equal(header, PKCS8_HEADER) {
		panic("PKCS8_HEADER is invalid")
	}
	secretKey := decrypted[SEED_OFFSET : SEED_OFFSET+SecLength]
	var result [64]byte
	copy(result[:], secretKey)
	return result, publicByte
}

func jsonDecryptData(encrypted []byte, passphrase string, encType []string) []byte {

	params, slat := scryptFromU8a(encrypted)

	keyLen := 64
	password, _ := scrypt.Key([]byte(passphrase), slat, params.N, params.r, params.p, keyLen)
	encrypted = encrypted[44:]
	var nonce [24]byte
	copy(nonce[:], encrypted[0:NonceLength])
	return naclDecrypt(encrypted[NonceLength:], nonce, u8aFixLength(password, 256, true))
}

func u8aFixLength(value []byte, bitLength int, atStart bool) []byte {
	byteLength := int(math.Ceil(float64((bitLength) / 8)))

	if bitLength == -1 || len(value) == byteLength {
		return value
	} else if len(value) > byteLength {
		return value[0:byteLength]
	}

	result := make([]byte, byteLength)
	if atStart {
		copy(result, value)
	} else {
		// TODO ... result[byteLength - len(value):] = value
	}

	return result
}

func naclDecrypt(encrypted []byte, nonce [24]byte, secret []byte) []byte {

	// Load your secret key from a safe place and reuse it across multiple
	// Seal calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	var secretKey [32]byte
	copy(secretKey[:], secret)

	decrypted, ok := secretbox.Open(nil, encrypted, &nonce, &secretKey)
	if !ok {
		panic("secretbox.Open fail")
	}

	return decrypted
}

func scryptFromU8a(data []byte) (Params, []byte) {
	BN_LE_OPTS := true
	N := u8aToBn(data[32+0:32+4], BN_LE_OPTS) // 32768
	p := u8aToBn(data[32+4:32+8], BN_LE_OPTS)
	r := u8aToBn(data[32+8:32+12], BN_LE_OPTS)
	salt := data[0:32]
	return Params{int(N), int(p), int(r)}, salt
}

func u8aToBn(data []byte, isLe bool) uint {
	str := hexutil.Encode(data)
	if strings.HasPrefix(str, "0x") {
		str = strings.TrimPrefix(str, "0x")
	}

	return hexToBn(str).toNumber()
}

type Params struct {
	N int
	p int
	r int
}

func hexToBn(str string) *Bn {

	start := 0
	val := float64(len(str)-start) / float64(6)
	length := math.Ceil(val)

	words := make([]uint, int(length))

	for i := 0; i < int(length); i++ {
		words[i] = 0
	}

	off := 0
	j := 0

	parseLength := len(str) - start

	var x int
	if parseLength%2 == 0 {
		x = start + 1
	} else {
		x = start
	}

	var w uint
	for i := x; i < len(str); i += 2 {
		w = parseHexByte(str, start, i) << off
		words[j] |= w & 0x3ffffff
		if off >= 18 {
			off -= 18
			j += 1
			words[j] |= w >> 26
		} else {
			off += 8
		}
	}
	return &Bn{
		words:  words,
		length: int(length),
	}

}

func parseHexByte(string string, lowerBound, index int) uint {
	r := parseHex4Bits(string, index)
	if index-1 >= lowerBound {
		r |= parseHex4Bits(string, index-1) << 4
	}
	return uint(r)
}

func parseHex4Bits(string string, index int) uint8 {
	c := (string)[index]
	// '0' - '9'
	if c >= 48 && c <= 57 {
		return c - 48
		// 'A' - 'F'
	} else if c >= 65 && c <= 70 {
		return c - 55
		// 'a' - 'f'
	} else if c >= 97 && c <= 102 {
		return c - 87
	} else {
		panic("Invalid character in " + string)
	}
}

type Bn struct {
	words    []uint
	length   int
	negative int
}

func (bn *Bn) toNumber() uint {
	var ret = bn.words[0]
	if bn.length == 2 {
		ret += bn.words[1] * 0x4000000
	} else if bn.length == 3 && bn.words[2] == 0x01 {
		// NOTE: at this stage it is known that the top bit is set
		ret += 0x10000000000000 + (bn.words[1] * 0x4000000)
	} else if bn.length > 2 {
		panic("Number can only safely store up to 53 bits")
	}
	if bn.negative != 0 {
		return -ret
	} else {
		return ret
	}
}

func divideScalarBytesByCofactor(scalar [32]byte) [32]byte {
	// copy from rust
	// curve25519_dalek::scalar::Scalar::divide_scalar_bytes_by_cofactor
	/*
		pub(crate) fn divide_scalar_bytes_by_cofactor(scalar: &mut [u8; 32]) {
		    let mut low = 0u8;
		    for i in scalar.iter_mut().rev() {
		        let r = *i & 0b00000111; // save remainder
		        *i >>= 3; // divide by 8
		        *i += low;
		        low = r << 5;
		    }
		}
	*/

	var data = make([]byte, 32)

	var low byte = 0
	for index, i := range reverse(scalar[:]) {
		r := i & 0b00000111
		i >>= 3
		i += low
		low = r << 5
		data[index] = i
	}

	var result [32]byte
	copy(result[:], reverse(data))
	return result
}

func reverse(a []byte) []byte {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}
