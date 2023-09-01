package common_adapters

type OutputPort struct {
	StatusCode int
	Data       any
}

type ErrorMessage struct {
	Message string `json:"message"`
}
