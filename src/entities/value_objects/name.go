package vo

import (
	shared_err "go_mongo_api/src/shared/errors"
)

type Name struct {
	Value string
}

func (n *Name) Validate() error {
	length := len(n.Value)

	if length <= 0 {
		return shared_err.NewInvalidOperationError(shared_err.INVALID_EMAIL)
	}

	return nil
}
