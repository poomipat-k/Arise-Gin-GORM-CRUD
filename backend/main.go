package main

import (
	"github.com/gin-gonic/gin"
	"github.com/poomipat-k/crud-arise/internal/config"
	"github.com/poomipat-k/crud-arise/internal/database"
	"github.com/poomipat-k/crud-arise/internal/routes"
)

func main() {
	// load config
	config.LoadConfig()

	// Init database
	database.InitDB()

	// Auto-migrate the database schema
	database.AutoMigrate()

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
