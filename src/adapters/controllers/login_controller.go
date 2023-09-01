package controllers

import (
	common_adapters "go_mongo_api/src/adapters/common"
	"go_mongo_api/src/adapters/presenters"
	"go_mongo_api/src/usecases/login"
	interactor "go_mongo_api/src/usecases/login"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	Interactor interactor.LoginInteractor
	Presenter  presenters.LoginPresenter
}

func (ctrl *LoginController) Run(input interactor.LoginInputDTO) common_adapters.OutputPort {
	result, err := ctrl.Interactor.Execute(input)
	return ctrl.Presenter.Show(result, err)
}

func BuildLoginController(i *login.LoginInteractor, ptr *presenters.LoginPresenter) *LoginController {
	return &LoginController{Interactor: *i, Presenter: *ptr}
}

// @Summary		Autenticação de usuário
// @Description	Autentica o usuário com email e senha.
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			body	body		login.LoginInputDTO	true	"Corpo da solicitação"
// @Success		200		{object}	presenters.LoginOutputFormatted
// @Failure		400		{object}	shared_err.InvalidOperationError	"Bad Request"
// @Failure		500		{object}	shared_err.InternalError			"Internal Server Error"
// @Router			/api/v1/auth/login [post]
func (ctrl *LoginController) Handle(c *gin.Context) {
	inputMap := c.MustGet("body").(map[string]any)
	loginInput := interactor.LoginInputDTO{
		Email:    inputMap["email"].(string),
		Password: inputMap["password"].(string),
	}

	result, err := ctrl.Interactor.Execute(loginInput)
	output := ctrl.Presenter.Show(result, err)

	c.JSON(output.StatusCode, output.Data)
}
