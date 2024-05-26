package db

import (
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"testing"
	"url-shortener/db/service"
	"url-shortener/util"
)

// saveTestUrl saves a random url record into the urls table
func saveTestUrl(t *testing.T) *service.SaveUrlResponse {
	db := GetDB()

	originalUrl := util.GenerateShortUrl() + util.GenerateShortUrl()
	shortUrl := util.GenerateShortUrl()

	req := service.NewSaveUrlRequest(originalUrl, shortUrl)

	res, err := db.SaveUrl(req)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, req.OriginalUrl(), res.OriginalUrl())
	require.Equal(t, req.ShortUrl(), res.ShortUrl())

	return res
}

// TestDatabase_SaveUrl tests GetOriginalUrl method of database type
func TestDatabase_SaveUrl(t *testing.T) {
	saveTestUrl(t)
}

// TestDatabase_GetOriginalUrl tests GetOriginalUrl method of database type
func TestDatabase_GetOriginalUrl(t *testing.T) {
	db := GetDB()

	t.Run("OK", func(t *testing.T) {
		url := saveTestUrl(t)

		res, err := db.GetOriginalUrl(service.NewGetOriginalUrlRequest(url.ShortUrl()))
		require.NoError(t, err)
		require.NotEmpty(t, res)

		require.Equal(t, url.OriginalUrl(), res.OriginalUrl())
	})
	t.Run("NotFound", func(t *testing.T) {
		res, err := db.GetOriginalUrl(service.NewGetOriginalUrlRequest("non existing short url"))
		require.Error(t, err)
		require.ErrorIs(t, err, gorm.ErrRecordNotFound)
		require.Nil(t, res)
	})
}
