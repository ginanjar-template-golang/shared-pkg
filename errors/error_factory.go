package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code       string `json:"code"`
	Detail     any    `json:"detail,omitempty"`
	HTTPStatus int    `json:"-"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %v", e.Code, e.Detail)
}

func New(code string, detail any) *AppError {
	status := MapAppCodeToHTTPStatus(code)
	return &AppError{
		Code:       code,
		Detail:     detail,
		HTTPStatus: status,
	}
}

func NewWithStatus(code string, detail any, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		Detail:     detail,
		HTTPStatus: httpStatus,
	}
}

func ParseError(err error) *AppError {
	if err == nil {
		return &AppError{
			Code:       CodeSuccess,
			Detail:     nil,
			HTTPStatus: http.StatusOK,
		}
	}
	if ae, ok := err.(*AppError); ok {
		// already an AppError
		return ae
	}
	// fallback ke INTERNAL
	return New(CodeInternalError, err.Error())
}

func MapAppCodeToHTTPStatus(code string) int {
	if s, ok := InternalErrorToHTTP[code]; ok {
		return s
	}
	// fallback ke mapping HTTP status names
	if s, ok := HTTPStatusMap[code]; ok {
		return s
	}
	// ultimate fallback
	return http.StatusInternalServerError
}
