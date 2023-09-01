package presenters

import (
	cad "go_mongo_api/src/adapters/common"
	common_ptr "go_mongo_api/src/adapters/presenters/common"
	"go_mongo_api/src/entities"
	shared_err "go_mongo_api/src/shared/errors"
	interactor "go_mongo_api/src/usecases/get_users"
	"net/http"
)

type GetUsersOutputFormatted struct {
	Users      []common_ptr.UserFormatted `json:"users"`
	Pagination common_ptr.PaginateResult  `json:"pagination"`
}

type GetUsersPresenter struct {
	userPtr       common_ptr.UserPresenter
	paginationPtr common_ptr.PaginationPresenter
}

func BuildGetUsersPresenter() *GetUsersPresenter {
	return &GetUsersPresenter{}
}

func (ptr *GetUsersPresenter) Show(result *interactor.GetUsersOutputDTO, paginationOptions entities.PaginationOptions, err error) cad.OutputPort {
	if err != nil {
		return cad.OutputPort{StatusCode: http.StatusBadRequest, Data: cad.ErrorMessage{Message: shared_err.INTERNAL_ERROR}}
	}

	u := ptr.userPtr.FormatList(result.Users)
	p := ptr.paginationPtr.Format(paginationOptions, result.TotalUsers)
	data := GetUsersOutputFormatted{Users: u, Pagination: p}

	return cad.OutputPort{StatusCode: http.StatusOK, Data: data}
}
