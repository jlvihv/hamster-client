package account

import (
	"context"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
	"hamster-client/utils"
)

type ServiceImpl struct {
	ctx      context.Context
	db       *gorm.DB
	httpUtil *utils.HttpUtil
}

func NewServiceImpl(ctx context.Context, db *gorm.DB, httpUtil *utils.HttpUtil) ServiceImpl {
	return ServiceImpl{ctx, db, httpUtil}
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
	u, _ := a.GetAccount()
	//save or update account
	u.PublicKey = account.PublicKey
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
