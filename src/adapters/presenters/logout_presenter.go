package presenters

import (
	cad "go_mongo_api/src/adapters/common"
	e "go_mongo_api/src/shared/errors"
	shared_err "go_mongo_api/src/shared/errors"
	"net/http"
)

type LogoutPresenter struct{}

func BuildLogoutPresenter() *LogoutPresenter {
	return &LogoutPresenter{}
}

func (p *LogoutPresenter) Show(err error) cad.OutputPort {
	if err != nil {
		return p.formatError(err)
	}

	return cad.OutputPort{StatusCode: http.StatusNoContent}
}

func (p *LogoutPresenter) formatError(err error) cad.OutputPort {
	if uErr, ok := err.(*e.UnauthorizedError); ok {
		return cad.OutputPort{StatusCode: http.StatusConflict, Data: cad.ErrorMessage{Message: uErr.Error()}}
	}

	return cad.OutputPort{StatusCode: http.StatusBadRequest, Data: cad.ErrorMessage{Message: shared_err.INTERNAL_ERROR}}
}
