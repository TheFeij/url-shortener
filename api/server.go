package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

		// CORS middleware configuration
		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"*"}
		config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
		config.AllowHeaders = []string{"Origin", "Content-Type"}

		// use the CORS middleware with the custom configuration
		router.Use(cors.New(config))

		// initialize singleton server object
		apiServer = server{
			router:    router,
			dbService: dbService,
			address:   address,
		}

		// add route handlers
		apiServer.addRouteHandlers()
	})
}

// addRouteHandlers add route handlers to apiServer's router
func (s *server) addRouteHandlers() {
	// Define API routes
	s.router.POST("/links", s.shortenUrl)
	s.router.GET("/links/:short_url", s.redirectShortUrl)
}

// GetServer returns singleton instance of server
func GetServer() *server {
	return &apiServer
}

// StartServer starts server on the specified address
func (s server) StartServer() error {
	return s.router.Run(s.address)
}
