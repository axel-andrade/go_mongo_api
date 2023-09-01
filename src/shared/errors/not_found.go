package shared_err

const (
	USER_NOT_FOUND = "user not found"
)

type NotFoundError struct {
	Message string `json:"message" example:"John Doe"`
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Message: message,
	}
}
