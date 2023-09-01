package logout

import (
	shared_err "go_mongo_api/src/shared/errors"
)

type LogoutInteractor struct {
	Gateway LogoutGateway
}

func BuildLogoutInteractor(g LogoutGateway) *LogoutInteractor {
	return &LogoutInteractor{g}
}

func (bs *LogoutInteractor) Execute(encodedToken string) error {
	au, err := bs.Gateway.ExtractTokenMetadata(encodedToken)
	if err != nil {
		return shared_err.NewUnauthorizedError()
	}

	deleted, err := bs.Gateway.DeleteAuth(au.AccessUUID)
	if err != nil || deleted == 0 {
		return shared_err.NewUnauthorizedError()
	}

	return nil
}
