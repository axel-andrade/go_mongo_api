package shared_err

const (
	INCORRECT_PASSWORD = "incorrect password"
	INVALID_EMAIL      = "invalid email"
	INVALID_PASSWORD   = "invalid password"
	NAME_IS_EMPTY      = "name is empty"
)

type InvalidOperationError struct {
	Message string `json:"message" example:"incorrect password"`
}

func (e *InvalidOperationError) Error() string {
	return e.Message
}

func NewInvalidOperationError(message string) *InvalidOperationError {
	return &InvalidOperationError{
		Message: message,
	}
}
