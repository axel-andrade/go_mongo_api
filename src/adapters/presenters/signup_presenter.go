package presenters

import (
	cad "go_mongo_api/src/adapters/common"
	common_ptr "go_mongo_api/src/adapters/presenters/common"
	e "go_mongo_api/src/shared/errors"
	shared_err "go_mongo_api/src/shared/errors"
	interactor "go_mongo_api/src/usecases/signup"
	"net/http"
)

type SignupPresenter struct {
	UserPtr common_ptr.UserPresenter
}

func BuildSignupPresenter() *SignupPresenter {
	return &SignupPresenter{}
}

func (p *SignupPresenter) Show(result *interactor.SignupOutputDTO, err error) cad.OutputPort {
	if err != nil {
		return p.formatError(err)
	}

	return cad.OutputPort{StatusCode: http.StatusCreated, Data: p.UserPtr.Format(result.User)}
}

func (p *SignupPresenter) formatError(err error) cad.OutputPort {
	if cErr, ok := err.(*e.ConflictError); ok {
		return cad.OutputPort{StatusCode: http.StatusConflict, Data: cad.ErrorMessage{Message: cErr.Error()}}
	}

	return cad.OutputPort{StatusCode: http.StatusBadRequest, Data: cad.ErrorMessage{Message: shared_err.INTERNAL_ERROR}}
}
