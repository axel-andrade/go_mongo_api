package controllers

import (
	"go_mongo_api/src/adapters/presenters"
	"go_mongo_api/src/usecases/logout"
	interactor "go_mongo_api/src/usecases/logout"

	"github.com/gin-gonic/gin"
)

type LogoutController struct {
	Interactor interactor.LogoutInteractor
	Presenter  presenters.LogoutPresenter
}

func BuildLogoutController(logoutInteractor *logout.LogoutInteractor, logoutPtr *presenters.LogoutPresenter) *LogoutController {
	return &LogoutController{Interactor: *logoutInteractor, Presenter: *logoutPtr}
}

func (ctrl *LogoutController) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	encodedToken := authHeader[len("Bearer "):]

	err := ctrl.Interactor.Execute(encodedToken)
	output := ctrl.Presenter.Show(err)

	c.JSON(output.StatusCode, output.Data)
}
