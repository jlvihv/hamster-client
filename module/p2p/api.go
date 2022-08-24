package p2p

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"hamster-client/module/account"
	"os"
	"path/filepath"
)

func initDB() *gorm.DB {
	configPath := initConfigPath()
	db, err := gorm.Open(sqlite.Open(filepath.Join(configPath, "link.db")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(
		&account.Account{},
	)
	var user account.Account
	result := db.First(&user)
	if result.Error != nil {
		println("------------------------------------")
		user := account.Account{WsUrl: "ws://127.0.0.1:9944"}
		db.Create(&user)
	}

	if err != nil {
		panic("failed to AutoMigrate Account")
	}
	return db
}

func initConfigPath() string {
	// initialize the configuration file
	dir := "~/.link/"
	linkConfig, err := homedir.Expand(dir)
	if err != nil {
		panic("failed to homedir Expand")
	}
	_, err = os.Stat(linkConfig)
	if err == nil {
		return linkConfig
	}
	err = os.MkdirAll(linkConfig, os.ModePerm)
	if err != nil {
		fmt.Printf("failed to config Mkdir err%s\n", err)
		panic("failed to config Mkdir err")
	}
	return linkConfig
}
