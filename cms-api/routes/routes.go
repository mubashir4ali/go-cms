package routes

import (
	"cms-api/controllers"
	"cms-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		protected := api.Group("/")
		protected.Use(middleware.AuthRequired())
		{
			protected.GET("/me", func(c *gin.Context) {
				user, _ := c.Get("user")
				c.JSON(200, user)
			})
		}
	}
}
