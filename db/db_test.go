package db

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestGetDB tests GetDB function
func TestGetDB(t *testing.T) {
	database := GetDB()
	require.NotEmpty(t, database)
}
