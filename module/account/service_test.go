package account

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"link-server/utils"
	"testing"
)

func TestServiceImpl_GetAccount(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("link.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Account{})
	accountService := NewServiceImpl(db, nil)
	account, err := accountService.GetAccount()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(account)
}

func TestServiceImpl_getUserIdFromToken(t *testing.T) {
	accountService := NewServiceImpl(nil, utils.NewHttp())
	token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcnQiOiIxNjM2NDQ2MDc3Mjg5IiwidXNlcl9pZCI6IjcyMGRjYmJhLWRmODgtNDZlZC04MjkxLWQ3YzQ0NWNkNjhiYSIsInNjb3BlIjpbImFsbCJdLCJtb2JpbGUiOiIxODIzNDE1OTA0NSIsImV4cCI6MTYzNjQ3NDg3NywiZGV2aWNlIjoiUEMiLCJqdGkiOiI2Y2NlNzA0My0yZjFhLTRkYzctYWE5ZC04ODhjZmMyM2YxYzkiLCJjbGllbnRfaWQiOiJ0bnRsaW5raW5nIn0.FNuKGI6Li0Qwit7-vAG2zSLfTsPlz9HiNxnYkRmRKV2Wpi0IXjrxYASEK_V7YGm_9j7FszyxLJ2bvs5kfQm0oKw0UQRaYBP_VlDfqJYGGgyl14cTqdEZpqX1nqBhZ7qf7nTwrOivWfd0kOtEazzz_EpwovTxr5L-WfQRuww8A8mWsyhjhOF9mJzfvSdpt8mS_CSdnHsfxbC3BIfpML3V4kRYkefIXifZO4Jj9ItwfJ_2a3Eslkzuz52jDtjQbINyhOEqRPfRqpmHRFyXCSZxVtrg6BY0ykZJeA-K9D2sy0gOK33Fh6XKGyj2sybgIS6GfEmmCq2VIdjCOhUh4CrBYw"
	fromToken, err := accountService.getUserIdFromToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fromToken)
}
