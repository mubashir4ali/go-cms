package main

import (
    "github.com/gin-gonic/gin"
    "cms-api/database"
    "cms-api/routes"
)

func main() {
    router := gin.Default()
    database.Connect()
    routes.SetupRoutes(router)
    router.Run(":8080")
}
