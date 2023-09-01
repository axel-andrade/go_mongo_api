package shared_err

const (
	UNAUTHORIZED = "unauthorized"
)

type UnauthorizedError struct {
	Message string `json:"message" example:"unauthorized"`
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

func NewUnauthorizedError() *UnauthorizedError {
	return &UnauthorizedError{
		Message: UNAUTHORIZED,
	}
}
