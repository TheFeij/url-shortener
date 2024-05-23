package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"url-shortener/config"
	"url-shortener/db"
)

// TestMain initializes test database and api server before running tests
func TestMain(m *testing.M) {
	// load configs
	configs := config.GetConfig("config", "../config", "json")

	// init database
	db.Init(configs.TestDatabaseAddress())

	// set gin mode to test mode
	gin.SetMode(gin.TestMode)

	// initialize the server
	Init(configs.ServerAddress())

	os.Exit(m.Run())
}

// TestHomePage tests "/" route of the api server
func TestHomePage(t *testing.T) {
	server := GetServer()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	server.router.ServeHTTP(recorder, req)
	require.Equal(t, http.StatusOK, recorder.Code)

	body, err := io.ReadAll(recorder.Body)
	require.NoError(t, err)
	require.Equal(t, "Welcome!", string(body))
}
