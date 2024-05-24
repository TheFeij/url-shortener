package db

import (
	"testing"
	"url-shortener/config"
)

// TestMain initializes test database and performs cleanup after tests are done
func TestMain(m *testing.M) {
	// load configurations
	conf := config.GetConfig("config", "../config", "json")

	// init database
	Init(conf.TestDatabaseAddress())
	db := GetDB()

	// cleanup database
	defer db.db.Exec("DELETE FROM urls")

	// migrate database schema
	db.migrate()

	// run tests
	m.Run()
}
