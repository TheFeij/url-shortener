package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// TestShortenUrl tests GenerateShortUrl
func TestGenerateShortUrl(t *testing.T) {
	shortUrl := GenerateShortUrl()
	require.NotEmpty(t, shortUrl)
	require.Len(t, shortUrl, 6)
}

// TestGenerateRandomString tests GenerateRandomString
func TestGenerateRandomString(t *testing.T) {
	randomString := generateRandomString(6, lowercase+uppercase+numbers)
	require.NotEmpty(t, randomString)
	require.Len(t, randomString, 6)
}
