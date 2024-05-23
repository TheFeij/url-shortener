package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

// server holds router
type server struct {
	router  *gin.Engine
	address string
}

// singleton instance of server
var apiServer server

// sync.Once to ensure Init is only called once
var once sync.Once

// Init initializes singleton instance of server
func Init(address string) {
	once.Do(func() {
		router := gin.New()

		router.Use(gin.Recovery())
		router.Use(gin.Logger())

		router.GET("/", func(context *gin.Context) {
			context.String(http.StatusOK, "Welcome!")
		})

		apiServer = server{
			router:  router,
			address: address,
		}
	})
}

// GetServer returns singleton instance of server
func GetServer() *server {
	return &apiServer
}

// StartServer starts server on the specified address
func (s server) StartServer() error {
	return s.router.Run(s.address)
}
