package main

import (
	"github/ariifysp/go-101/config"
	"github/ariifysp/go-101/pkg/cache"
	"github/ariifysp/go-101/pkg/database"
	"github/ariifysp/go-101/server"
)

func main() {
	config := config.GetConfig()

	db := database.NewPostgresDatabase(&config.Postgres)
	cache := cache.NewRedisClient(&config.Redis)

	server := server.NewServer(config, db, cache)
	server.Start()
}
