package main

import (
	"github.com/abhijitpattar/gin-rest-go/db"
	"github.com/abhijitpattar/gin-rest-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	// Create a Gin router with default middleware (logger and recovery)
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
