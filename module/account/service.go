package account

import (
	"context"
	"errors"
	"fmt"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
	"hamster-client/utils"
	"sync"
)

type ServiceImpl struct {
	ctx          context.Context
	db           *gorm.DB
	httpUtil     *utils.HttpUtil
	substrateApi *gsrpc.SubstrateAPI
	m            sync.Mutex
}

func NewServiceImpl(ctx context.Context, db *gorm.DB, httpUtil *utils.HttpUtil) ServiceImpl {
	return ServiceImpl{ctx: ctx, db: db, httpUtil: httpUtil}
}

// GetAccount get account information
func (a *ServiceImpl) GetAccount() (Account, error) {
	var account Account
	result := a.db.First(&account)
	if result.Error != nil {
		runtime.LogError(a.ctx, "GetAccount error")
	}

	return account, result.Error
}

// SaveAccount save account information
func (a *ServiceImpl) SaveAccount(account *Account) {
	a.m.Lock()
	defer a.m.Unlock()
	u, _ := a.GetAccount()
	//save or update account
	u.PublicKey = account.PublicKey
	if u.WsUrl != account.WsUrl {
		u.WsUrl = account.WsUrl
		if a.substrateApi != nil {
			a.substrateApi = nil
		}
	}

	a.db.Save(&u)
}

func (a *ServiceImpl) getUserIdFromToken(accessToken string) (string, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return fmt.Sprint(claims["user_id"]), nil
	}
	return "", errors.New("cannot parse accessToken")
}

func (a *ServiceImpl) GetSubstrateApi() (*gsrpc.SubstrateAPI, error) {
	a.m.Lock()
	defer a.m.Unlock()
	if a.substrateApi != nil {
		return a.substrateApi, nil
	}
	account, err := a.GetAccount()
	if err != nil {
		return nil, err
	}
	api, err := gsrpc.NewSubstrateAPI(account.WsUrl)
	if api == nil {
		return nil, err
	}
	a.substrateApi = api
	return a.substrateApi, nil
}
