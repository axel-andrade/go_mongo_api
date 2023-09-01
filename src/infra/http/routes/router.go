package routes

import (
	"go_mongo_api/docs"
	"go_mongo_api/src/infra/bootstrap"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, dependencies *bootstrap.Dependencies) *gin.Engine {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Swapp API"
	docs.SwaggerInfo.Description = "This is a sample server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "swagg.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	main := router.Group("/")
	configureDefaultRoutes(main)

	v1 := router.Group("api/v1")
	configureAuthRoutes(v1, dependencies)
	configureUsersRoutes(v1, dependencies)

	return router
}
