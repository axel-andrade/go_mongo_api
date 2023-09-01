package routes

import (
	"go_mongo_api/src/adapters/controllers"
	"go_mongo_api/src/infra/bootstrap"
	"go_mongo_api/src/infra/http/middlewares"

	"github.com/gin-gonic/gin"
)

func configureAuthRoutes(router *gin.RouterGroup, dependencies *bootstrap.Dependencies) {
	loginCtrl := new(controllers.LoginController)
	dependencies.Invoke(func(ctrl *controllers.LoginController) {
		loginCtrl = ctrl
	})

	logoutCtrl := new(controllers.LogoutController)
	dependencies.Invoke(func(ctrl *controllers.LogoutController) {
		logoutCtrl = ctrl
	})

	signupCtrl := new(controllers.SignUpController)
	dependencies.Invoke(func(ctrl *controllers.SignUpController) {
		signupCtrl = ctrl
	})

	auth := router.Group("auth")
	{
		auth.POST("/signup", middlewares.ValidateRequest("auth/signup"), signupCtrl.Handle)
		auth.POST("/login", middlewares.ValidateRequest("auth/login"), loginCtrl.Handle)
		auth.POST("/logout", logoutCtrl.Handle)
	}
}
