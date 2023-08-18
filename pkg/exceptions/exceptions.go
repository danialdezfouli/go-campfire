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

func InvalidLogin() *RequestError {
	return &RequestError{
		Code:    http.StatusUnauthorized,
		Message: "username or password is wrong",
	}
}

func OrganizationNotFound(id int) *RequestError {
	return &RequestError{
		Code:    http.StatusUnauthorized,
		Message: fmt.Sprintf("organization %d is not found", id),
	}
}
