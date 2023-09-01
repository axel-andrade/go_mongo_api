package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func configureDefaultRoutes(router *gin.RouterGroup) {
	main := router.Group("/")
	{
		main.GET("/healthcheck", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "OK"})
		})

		main.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
