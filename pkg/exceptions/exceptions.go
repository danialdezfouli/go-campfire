package exceptions

import (
	"errors"
	"fmt"
	"net/http"
)

type RequestError struct {
	Code int
	Err  error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.Code, r.Err)
}

func AuthenticationError() *RequestError {
	return &RequestError{
		Code: http.StatusUnauthorized,
		Err:  errors.New("unauthencated error"),
	}
}

func InvalidLoginError() *RequestError {
	return &RequestError{
		Code: http.StatusUnauthorized,
		Err:  errors.New("username or password is wrong"),
	}
}

func ValidationError(err error) *RequestError {
	return &RequestError{
		Code: http.StatusBadRequest,
		Err:  err,
	}
}
