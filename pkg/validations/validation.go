package validations

import "github.com/go-playground/validator/v10"

func Validate(input any) error {
	validate := validator.New()
	return validate.Struct(input)
}
