package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"url-shortener/db/service"
)

// TestMain initializes test database and api server before running tests
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

// newTestServer creates and returns a server instance to be used for testing
func newTestServer(dbService service.DBService) *server {
	// instance a new gin router
	router := gin.New()

	// add middlewares
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// initialize the test server object
	testServer := server{
		router:    router,
		dbService: dbService,
	}

	// add route handlers
	testServer.router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Welcome!")
	})
	testServer.router.POST("/shorten", testServer.shortenUrl)

	return &testServer
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
