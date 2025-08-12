package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rest-api/db"
	"github.com/go-rest-api/routes"
)

func main() {
	db.Init()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
