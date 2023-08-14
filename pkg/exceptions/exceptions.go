package exceptions

import "fmt"

type AuthenticationError struct{}

func (e *AuthenticationError) Error() string {
	return "Unauthencated error"
}

type InvalidLoginError struct{}

func (e *InvalidLoginError) Error() string {
	return "username or password is wrong"
}

type ValidationError struct {
	Err error
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error: %s", e.Err.Error())
}
