package errors

import "errors"

var (
	ErrInternalServer = errors.New("internal server error")
)

type AppError struct {
	Status  int
	Code    string
	Message string
}

func New(status int, code, message string) AppError {
	return AppError{Status: status, Code: code, Message: message}
}
