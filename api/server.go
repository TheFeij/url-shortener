package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"url-shortener/db/service"
)

// server holds router
type server struct {
	router    *gin.Engine
	address   string
	dbService service.DBService
}

// singleton instance of server
var apiServer server

// sync.Once to ensure Init is only called once
var once sync.Once

// Init initializes singleton instance of server
func Init(dbService service.DBService, address string) {
	once.Do(func() {
		// instance a new gin router
		router := gin.New()

		// add middlewares
		router.Use(gin.Recovery())
		router.Use(gin.Logger())

		// initialize singleton server object
		apiServer = server{
			router:    router,
			dbService: dbService,
			address:   address,
		}

		// add route handlers
		apiServer.router.GET("/", func(context *gin.Context) {
			context.String(http.StatusOK, "Welcome!")
		})
		apiServer.router.POST("/shorten", apiServer.shortenUrl)
	})
}

// addRouteHandlers add route handlers to apiServer's router
func (s *server) addRouteHandlers() {
	s.router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Welcome!")
	})
	s.router.POST("/shorten", s.shortenUrl)
	s.router.GET("/redirect/:short_url", s.redirectShortUrl)
}

// GetServer returns singleton instance of server
func GetServer() *server {
	return &apiServer
}

// StartServer starts server on the specified address
func (s server) StartServer() error {
	return s.router.Run(s.address)
}
