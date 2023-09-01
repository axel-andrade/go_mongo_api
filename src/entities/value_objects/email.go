package vo

import (
	shared_err "go_mongo_api/src/shared/errors"
	"regexp"
)

type Email struct {
	Value string
}

func (e *Email) Validate() error {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if regex.MatchString(e.Value) {
		return nil
	}

	return shared_err.NewInvalidOperationError(shared_err.INVALID_EMAIL)
}
