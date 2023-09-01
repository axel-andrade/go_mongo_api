package controllers

import (
	"go_mongo_api/src/adapters/presenters"
	ptr "go_mongo_api/src/adapters/presenters"
	"go_mongo_api/src/entities"
	"go_mongo_api/src/usecases/get_users"
	interactor "go_mongo_api/src/usecases/get_users"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUsersController struct {
	Interactor interactor.GetUsersInteractor
	Presenter  ptr.GetUsersPresenter
}

func BuildGetUsersController(i *get_users.GetUsersInteractor, ptr *presenters.GetUsersPresenter) *GetUsersController {
	return &GetUsersController{Interactor: *i, Presenter: *ptr}
}

// @Summary		Get users
// @Description	Returns a list of users from the database.
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			page	query		int	false	"Page number"
// @Param			limit	query		int	false	"Number of items per page"
// @Success		200		{object}	presenters.GetUsersOutputFormatted
//
// @Failure		400		{object}	shared_err.InvalidOperationError	"Bad Request"
// @Failure		500		{object}	shared_err.InternalError			"Internal Server Error"
//
// @Router			/api/v1/users [get]
func (ctrl *GetUsersController) Handle(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	paginationOptions := entities.PaginationOptions{
		Limit:  limit,
		Page:   page,
		Sort:   c.Query("sort"),
		Search: c.Query("search"),
	}

	input := get_users.GetUsersInputDTO{PaginationOptions: paginationOptions}

	result, err := ctrl.Interactor.Execute(input)
	output := ctrl.Presenter.Show(result, input.PaginationOptions, err)

	c.JSON(output.StatusCode, output.Data)
}
