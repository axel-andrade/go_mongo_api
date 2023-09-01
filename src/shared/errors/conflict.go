package shared_err

const (
	EMAIL_ALREADY_EXISTS = "email already registered"
)

type ConflictError struct {
	Message string `json:"message" example:"incorrect password"`
}

func (e *ConflictError) Error() string {
	return e.Message
}

func NewConflictError(message string) *ConflictError {
	return &ConflictError{
		Message: message,
	}
}
