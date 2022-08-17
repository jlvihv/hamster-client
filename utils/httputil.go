package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type HttpUtil struct {
	client *resty.Client
}

func NewHttp() *HttpUtil {
	return &HttpUtil{client: resty.New()}
}

func (h *HttpUtil) NewRequest() *resty.Request {
	res := h.client.R().
		SetHeader("Accept", "application/json, text/plain, */*").
		SetHeader("Content-Type", "application/json;charset=UTF-8").
		SetContentLength(true)
	return res
}

func GetSS58AuthDataWithKeyringPair(keyringPair signature.KeyringPair) string {
	ss58Address := keyringPair.Address
	data := uuid.New().String()
	signDataHex := hex.EncodeToString(signWithKeyringPair(keyringPair, []byte(data)))
	return fmt.Sprintf("%s:%s:%s", ss58Address, data, signDataHex)
}

func signWithKeyringPair(keyringPair signature.KeyringPair, data []byte) []byte {
	signData, err := keyringPair.Sign(data)
	if err != nil {
		log.Errorf("signature.Sign error: %s", err)
		return nil
	}
	return signData
}
