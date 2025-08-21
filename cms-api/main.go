package main

import (
	"cms-api/database"
	"cms-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.Connect()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
