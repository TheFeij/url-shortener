package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"url-shortener/config"
)

type database struct {
	db *gorm.DB
}

var db database

func init() {
	// load configs
	conf := config.GetConfig("config", "../config", "json")

	var err error
	db.db, err = gorm.Open(postgres.Open(conf.DatabaseAddress()), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("unable to start a connection to dataabse: %w", err))
	}
}
