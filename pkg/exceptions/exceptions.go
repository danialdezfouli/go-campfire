package exceptions

import (
	"campfire/internal/domain"
	"fmt"
	"net/http"
)

func NewOrganizationNotFound(id domain.OrganizationId) *RequestError {
	return &RequestError{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("organization %d is not found", id),
	}
}

func NewInternalServerError(message string, err error) *RequestError {
	return &RequestError{
		Code:    http.StatusInternalServerError,
		Message: fmt.Sprintf(message+" %v", err),
	}
}

var (
	Unauthenticated = &RequestError{
		Code:    http.StatusUnauthorized,
		Message: "unauthencated error",
	}

	InvalidLogin = &RequestError{
		Code:    http.StatusUnauthorized,
		Message: "username or password is wrong",
	}

	InvalidToken = &RequestError{
		Code:    http.StatusUnauthorized,
		Message: "invalid token",
	}

	InternalServerError = &RequestError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
)
