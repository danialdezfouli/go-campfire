package exceptions

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CustomError interface {
	Code() int
	Error() string
}

type RequestError struct {
	code    int    `json:"-"`
	Message string `json:"messages"`
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", e.code, e.Message)
}

func (e *RequestError) Code() int {
	return e.code
}

func Unauthenticated() *RequestError {
	return &RequestError{
		code:    http.StatusUnauthorized,
		Message: "unauthencated error",
	}
}

func InvalidLogin() *RequestError {
	return &RequestError{
		code:    http.StatusUnauthorized,
		Message: "username or password is wrong",
	}
}

//
//
//

type ValidationError struct {
	code    int      `json:"-"`
	Message string   `json:"messages"`
	Errors  []string `json:"errors"`
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("status %d: err %v", e.code, e.Message)
}

func (e *ValidationError) Code() int {
	return e.code
}

func NewValidationError(err error) *ValidationError {
	validationErrors := err.(validator.ValidationErrors)
	errors := []string{}

	for _, validationError := range validationErrors {
		errors = append(errors, validationError.Error())
	}

	return &ValidationError{
		code:    http.StatusBadRequest,
		Message: "inputs is invalid",
		Errors:  errors,
	}
}
