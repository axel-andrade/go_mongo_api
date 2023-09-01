package vo

import (
	shared_err "go_mongo_api/src/shared/errors"
)

type Password struct {
	Value string
}

func (p *Password) Validate() error {
	if length := len(p.Value); length >= 6 {
		return nil
	}

	return shared_err.NewInvalidOperationError(shared_err.INVALID_PASSWORD)
}
