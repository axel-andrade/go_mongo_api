package handlers

import (
	common_ptr "go_mongo_api/src/adapters/presenters/common"

	"github.com/go-playground/validator"
)

var validate = validator.New()

type JsonHandler struct{}

func BuildJsonHandler() *JsonHandler {
	return &JsonHandler{}
}

func (e *JsonHandler) ValidateStruct(input any) []common_ptr.ValidateDetail {
	if err := validate.Struct(input); err != nil {
		var details []common_ptr.ValidateDetail
		for _, err := range err.(validator.ValidationErrors) {
			details = append(details, common_ptr.ValidateDetail{
				Namespace: err.Namespace(),
				Tag:       err.Tag(),
				Param:     err.Param(),
			})
		}

		return details
	}

	return nil
}
