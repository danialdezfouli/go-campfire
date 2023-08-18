package exceptions

import "fmt"

type CustomError interface {
	GetCode() int
	Error() string
}

type RequestError struct {
	Code    int    `json:"-"`
	Message string `json:"messages"`
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", e.Code, e.Message)
}

func (e *RequestError) GetCode() int {
	return e.Code
}
