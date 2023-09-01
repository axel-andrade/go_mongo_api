package shared_err

const (
	INTERNAL_ERROR = "internal error"
)

type InternalError struct {
	Message string `json:"message" example:"internal error"`
}

func (e *InternalError) Error() string {
	return e.Message
}

func NewInternalError() *InternalError {
	return &InternalError{
		Message: INTERNAL_ERROR,
	}
}
