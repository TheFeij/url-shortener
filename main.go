package main

import (
	"url-shortener/api"
	"url-shortener/config"
	"url-shortener/db"
)

func main() {
	// load configs
	configs := config.GetConfig("config", "./config", "json")

	// init database
	db.Init(configs.DatabaseAddress())

	// get the singleton server instance
	api.Init(db.GetDB(), configs.ServerAddress())
	server := api.GetServer()

	// start the server
	server.StartServer()
}
