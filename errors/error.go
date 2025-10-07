package errors

import (
	"fmt"
)

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Locale  string `json:"locale,omitempty"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func New(code, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func Wrap(code, message string, err error) *AppError {
	return &AppError{Code: code, Message: message, Err: err}
}
