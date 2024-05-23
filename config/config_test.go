package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestGetConfig tests GetConfig function
func TestGetConfig(t *testing.T) {
	configs := GetConfig("config_test", "./config", "json")
	require.NotEmpty(t, configs)

	require.Equal(
		t,
		"postgresql://****:****@localhost:5432/url_shortener?sslmode=disable",
		configs.DatabaseAddress(),
	)
	require.Equal(
		t,
		"postgresql://****:****@localhost:5432/url_shortener_test?sslmode=disable",
		configs.TestDatabaseAddress(),
	)
}
