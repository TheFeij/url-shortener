package db

import (
	"github.com/stretchr/testify/require"
	"testing"
	"url-shortener/db/store"
)

func TestDatabase_SaveUrl(t *testing.T) {
	db := GetDB()
	req := store.NewSaveUrlRequest("original url", "short url")

	res, err := db.SaveUrl(req)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, req.OriginalUrl(), res.OriginalUrl())
	require.Equal(t, req.ShortUrl(), res.ShortUrl())
}
