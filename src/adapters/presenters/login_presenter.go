package presenters

import (
	common_adapters "go_mongo_api/src/adapters/common"
	common_ptr "go_mongo_api/src/adapters/presenters/common"
	e "go_mongo_api/src/shared/errors"
	shared_err "go_mongo_api/src/shared/errors"
	interactor "go_mongo_api/src/usecases/login"
	"net/http"
)

type LoginOutputFormatted struct {
	AccessToken  string                   `json:"access_token" description:"Access token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
	RefreshToken string                   `json:"refresh" description:"Access token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
	User         common_ptr.UserFormatted `json:"user"`
}

type LoginPresenter struct {
	userPtr common_ptr.UserPresenter
}

func BuildLoginPresenter() *LoginPresenter {
	return &LoginPresenter{}
}

func (p *LoginPresenter) Show(result *interactor.LoginOutputDTO, err error) common_adapters.OutputPort {
	if err != nil {
		return p.formatError(err)
	}

	return p.formatSuccessOutput(result)
}

func (p *LoginPresenter) formatSuccessOutput(result *interactor.LoginOutputDTO) common_adapters.OutputPort {
	data := LoginOutputFormatted{
		AccessToken:  result.AccessToken,
		RefreshToken: result.RefreshToken,
		User:         p.userPtr.Format(result.User),
	}

	return common_adapters.OutputPort{
		StatusCode: http.StatusOK,
		Data:       data,
	}
}

func (p *LoginPresenter) formatError(err error) common_adapters.OutputPort {
	if ipErr, ok := err.(*e.InvalidOperationError); ok {
		return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: common_adapters.ErrorMessage{Message: ipErr.Error()}}
	}

	if nfErr, ok := err.(*e.NotFoundError); ok {
		return common_adapters.OutputPort{StatusCode: http.StatusNotFound, Data: common_adapters.ErrorMessage{Message: nfErr.Error()}}

	}

	return common_adapters.OutputPort{StatusCode: http.StatusBadRequest, Data: common_adapters.ErrorMessage{Message: shared_err.INTERNAL_ERROR}}
}
