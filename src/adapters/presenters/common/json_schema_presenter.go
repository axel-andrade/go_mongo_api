package common_ptr

type Detail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidateDetail struct {
	Namespace string
	Tag       string
	Param     string
}

type JsonSchemaError struct {
	Message string   `json:"message"`
	Details []Detail `json:"details"`
}

type JsonSchemaPresenter struct{}

func BuildJsonSchemaPresenter() *JsonSchemaPresenter {
	return &JsonSchemaPresenter{}
}

func (ptr *JsonSchemaPresenter) Format(validateDetails []ValidateDetail) *JsonSchemaError {
	var d []Detail
	for _, vd := range validateDetails {
		d = append(d, Detail{
			Field:   vd.Namespace,
			Message: ptr.formatMessage(vd.Namespace, vd.Param),
		})
	}

	return &JsonSchemaError{Message: "Invalid JSON Schema", Details: d}
}

func (ptr *JsonSchemaPresenter) formatMessage(field string, defaultMsg string) string {
	errorMessages := map[string]string{
		"query.limit": "must be a positive number between 1 and 100.",
		"query.page":  "must be a positive number between 1 and 10.",
	}

	if msg, ok := errorMessages[field]; ok {
		return msg
	}

	return defaultMsg
}
