package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"url-shortener/db/models"
)

// database holds gorm.database and methods to interact with the database
type database struct {
	db *gorm.DB
}

// singleton instance of database
var db database

// Init initializes singleton instance of database. connects to the database with the input address
func Init(address string) {
	var err error
	db.db, err = gorm.Open(postgres.Open(address), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("unable to start a connection to databse: %w", err))
	}

	db.migrate()
}

// GetDB returns singleton instance of the database
func GetDB() *database {
	return &db
}

// migrate automatically migrates the database schemas for the registered models.
//
// It uses GORM's AutoMigrate function, which will create tables, missing columns, and missing indexes
// without deleting any existing data.
func (d database) migrate() {
	d.db.AutoMigrate(&models.Url{})
}
