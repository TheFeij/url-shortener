package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestGetConfig tests GetConfig function
func TestGetConfig(t *testing.T) {
	configs, err := GetConfig("config_test", "./", "json")
	require.NoError(t, err)
	require.NotEmpty(t, configs)

	require.Equal(
		t,
		"postgresql://****:****@localhost:5432/url_shortener?sslmode=disable",
		configs.DatabaseAddress(),
	)
}
