package controllers

import (
	"go_mongo_api/src/adapters/presenters"
	"go_mongo_api/src/usecases/signup"
	interactor "go_mongo_api/src/usecases/signup"

	"github.com/gin-gonic/gin"
)

type SignUpController struct {
	Interactor interactor.SignupInteractor
	Presenter  presenters.SignupPresenter
}

func BuildSignUpController(i *signup.SignupInteractor, ptr *presenters.SignupPresenter) *SignUpController {
	return &SignUpController{Interactor: *i, Presenter: *ptr}
}

// @Summary		Register user
// @Description	Register an user.
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			body	body		signup.SignupInputDTO	true	"Corpo da solicitação"
// @Success		200		{object}	common_ptr.UserFormatted
// @Failure		400		{object}	shared_err.InvalidOperationError	"Bad Request"
// @Failure		500		{object}	shared_err.InternalError			"Internal Server Error"
// @Router			/api/v1/auth/signup [post]
func (ctrl *SignUpController) Handle(c *gin.Context) {
	inputMap := c.MustGet("body").(map[string]any)
	signupInput := interactor.SignupInputDTO{
		Email:    inputMap["email"].(string),
		Name:     inputMap["name"].(string),
		Password: inputMap["password"].(string),
	}

	result, err := ctrl.Interactor.Execute(signupInput)
	output := ctrl.Presenter.Show(result, err)

	c.JSON(output.StatusCode, output.Data)

}
