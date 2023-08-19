package exceptions

import (
	"fmt"
	"net/http"
)

func Unauthenticated() *RequestError {
	return &RequestError{
		Code:    http.StatusUnauthorized,
		Message: "unauthencated error",
	}
}

func InvalidToken() *RequestError {
	return &RequestError{
		Code:    http.StatusUnauthorized,
		Message: "invalid token",
	}
}

func InvalidLogin() *RequestError {
	return &RequestError{
		Code:    http.StatusUnauthorized,
		Message: "username or password is wrong",
	}
}

func OrganizationNotFound(id int) *RequestError {
	return &RequestError{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("organization %d is not found", id),
	}
}

func InternalServerError(message string, err error) *RequestError {
	return &RequestError{
		Code:    http.StatusInternalServerError,
		Message: fmt.Sprintf(message+" %v", err),
	}
}
