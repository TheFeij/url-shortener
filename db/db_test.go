package db

import (
	"github.com/stretchr/testify/require"
	"testing"
	"url-shortener/config"
)

// TestGetDB tests GetDB function
func TestGetDB(t *testing.T) {
	conf := config.GetConfig("config", "../config", "json")
	Init(conf.DatabaseAddress())

	database := GetDB()
	require.NotEmpty(t, database)
}
