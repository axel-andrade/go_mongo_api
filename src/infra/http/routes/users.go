package routes

import (
	"go_mongo_api/src/adapters/controllers"
	"go_mongo_api/src/infra/bootstrap"
	"go_mongo_api/src/infra/http/middlewares"

	"github.com/gin-gonic/gin"
)

func configureUsersRoutes(router *gin.RouterGroup, dependencies *bootstrap.Dependencies) {
	getUsersCtrl := new(controllers.GetUsersController)
	dependencies.Invoke(func(ctrl *controllers.GetUsersController) {
		getUsersCtrl = ctrl
	})

	users := router.Group("users")
	{
		users.GET("/", middlewares.Authorize(dependencies), middlewares.ValidateRequest("users/get_users"), getUsersCtrl.Handle)
	}
}
