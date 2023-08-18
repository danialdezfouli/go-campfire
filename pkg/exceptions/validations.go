package exceptions

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Code    int      `json:"-"`
	Message string   `json:"messages"`
	Errors  []string `json:"errors"`
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("status %d: err %v", e.Code, e.Message)
}

func (e *ValidationError) GetCode() int {
	return e.Code
}

func NewValidationError(err error) *ValidationError {
	validationErrors := err.(validator.ValidationErrors)
	errors := []string{}

	for _, validationError := range validationErrors {
		errors = append(errors, validationError.Error())
	}

	return &ValidationError{
		Code:    http.StatusBadRequest,
		Message: "input is invalid",
		Errors:  errors,
	}
}
