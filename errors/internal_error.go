package errors

import (
	"fmt"

	"github.com/ginanjar-template-golang/shared-pkg/constants"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
)

type InternalError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// ==========================
// RESOURCE ERRORS
// ==========================
func ResourceNotFound(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.NotFound,
		Message: fmt.Sprintf("%s: %s", t.T("notFoundResource"), key),
		Data:    data,
	}
}

func FindResourceError(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.FailedFindResource,
		Message: fmt.Sprintf("%s: %s", t.T("errorFindResource"), key),
		Data:    data,
	}
}

func CreateResourceError(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.FailedCreateResource,
		Message: fmt.Sprintf("%s: %s", t.T("errorCreateResource"), key),
		Data:    data,
	}
}

func UpdateResourceError(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.FailedUpdateResource,
		Message: fmt.Sprintf("%s: %s", t.T("errorUpdateResource"), key),
		Data:    data,
	}
}

func DeleteResourceError(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.FailedDeleteResource,
		Message: fmt.Sprintf("%s: %s", t.T("errorDeleteResource"), key),
		Data:    data,
	}
}

// ==========================
// VALIDATION ERRORS
// ==========================
func InvalidBody(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.BadRequest,
		Message: fmt.Sprintf("%s: %s", t.T("invalidBody"), key),
		Data:    data,
	}
}

func InvalidTypeError(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.BadRequestType,
		Message: fmt.Sprintf("%s: %s", t.T("invalidType"), key),
		Data:    data,
	}
}

func InvalidFormatError(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.BadRequestFormat,
		Message: fmt.Sprintf("%s: %s", t.T("invalidFormat"), key),
		Data:    data,
	}
}

func InCompleteKeyError(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.BadRequestKey,
		Message: fmt.Sprintf("%s: %s", t.T("incompleteKey"), key),
		Data:    data,
	}
}

func InCompleteValueError(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.BadRequestValue,
		Message: fmt.Sprintf("%s: %s", t.T("incompleteValue"), key),
		Data:    data,
	}
}

func AlreadyUsedError(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.DuplicateValue,
		Message: fmt.Sprintf("%s: %s", t.T("alreadyUsed"), key),
		Data:    data,
	}
}

func InvalidOptionError(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.OutsideOption,
		Message: fmt.Sprintf("%s: %s", t.T("invalidOption"), key),
		Data:    data,
	}
}

func ValueMissMatch(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.ValueMissMatch,
		Message: fmt.Sprintf("%s: %s", t.T("somethingWentWrong"), key),
		Data:    data,
	}
}

// ==========================
// AUTH ERRORS
// ==========================
func Unauthorized(t translator.Translator, data any) InternalError {
	return InternalError{
		Code:    constants.Unauthorized,
		Message: t.T("unauthorized"),
		Data:    data,
	}
}

func Forbidden(t translator.Translator, data any) InternalError {
	return InternalError{
		Code:    constants.Forbidden,
		Message: t.T("forbidden"),
		Data:    data,
	}
}

func LoginError(t translator.Translator, data any) InternalError {
	return InternalError{
		Code:    constants.InvalidAuth,
		Message: t.T("loginError"),
		Data:    data,
	}
}

func ExpiredError(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.ExpiredAuth,
		Message: fmt.Sprintf("%s: %s", t.T("expired"), key),
		Data:    data,
	}
}

// ==========================
// GENERAL ERRORS
// ==========================
func GeneralError(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.GeneralErrors,
		Message: fmt.Sprintf("%s: %s", t.T("generalRequestErrors"), key),
		Data:    data,
	}
}

func InvalidProcess(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.BadRequest,
		Message: fmt.Sprintf("%s: %s", t.T("invalidProcess"), key),
		Data:    data,
	}
}

func UnknownError(t translator.Translator, key string, data any) InternalError {
	return InternalError{
		Code:    constants.InternalServerError,
		Message: t.T("somethingWentWrong"),
		Data:    data,
	}
}
