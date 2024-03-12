package common

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func ValidateInput(data interface{}) validator.ValidationErrors {
	err := validate.Struct(data)
	if err != nil {
		errors := err.(validator.ValidationErrors)

		return errors
	}

	return nil
}
