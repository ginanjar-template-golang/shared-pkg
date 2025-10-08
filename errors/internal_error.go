package errors

import (
	"fmt"

	"github.com/ginanjar-template-golang/shared-pkg/constants"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
)

type InternalError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// ==========================
// helper for auto-log
// ==========================
func logError(code int, msg string, data any) {
	logger.Error(msg, map[string]any{
		"code": code,
		"data": data,
	})
}

// ==========================
// RESOURCE ERRORS
// ==========================
func ResourceNotFound(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("notFoundResource"), key)
	logError(constants.NotFound, msg, data)
	return InternalError{Code: constants.NotFound, Message: msg, Data: data}
}

func FindResourceError(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("errorFindResource"), key)
	logError(constants.FailedFindResource, msg, data)
	return InternalError{Code: constants.FailedFindResource, Message: msg, Data: data}
}

func CreateResourceError(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("errorCreateResource"), key)
	logError(constants.FailedCreateResource, msg, data)
	return InternalError{Code: constants.FailedCreateResource, Message: msg, Data: data}
}

func UpdateResourceError(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("errorUpdateResource"), key)
	logError(constants.FailedUpdateResource, msg, data)
	return InternalError{Code: constants.FailedUpdateResource, Message: msg, Data: data}
}

func DeleteResourceError(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("errorDeleteResource"), key)
	logError(constants.FailedDeleteResource, msg, data)
	return InternalError{Code: constants.FailedDeleteResource, Message: msg, Data: data}
}

// ==========================
// VALIDATION ERRORS
// ==========================
func InvalidBody(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("invalidBody"), key)
	logError(constants.BadRequest, msg, data)
	return InternalError{Code: constants.BadRequest, Message: msg, Data: data}
}

func InvalidTypeError(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("invalidType"), key)
	logError(constants.BadRequestType, msg, data)
	return InternalError{Code: constants.BadRequestType, Message: msg, Data: data}
}

func InvalidFormatError(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("invalidFormat"), key)
	logError(constants.BadRequestFormat, msg, data)
	return InternalError{Code: constants.BadRequestFormat, Message: msg, Data: data}
}

func InCompleteKeyError(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("incompleteKey"), key)
	logError(constants.BadRequestKey, msg, data)
	return InternalError{Code: constants.BadRequestKey, Message: msg, Data: data}
}

func InCompleteValueError(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("incompleteValue"), key)
	logError(constants.BadRequestValue, msg, data)
	return InternalError{Code: constants.BadRequestValue, Message: msg, Data: data}
}

func AlreadyUsedError(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("alreadyUsed"), key)
	logError(constants.DuplicateValue, msg, data)
	return InternalError{Code: constants.DuplicateValue, Message: msg, Data: data}
}

func InvalidOptionError(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("invalidOption"), key)
	logError(constants.OutsideOption, msg, data)
	return InternalError{Code: constants.OutsideOption, Message: msg, Data: data}
}

func ValueMissMatch(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("somethingWentWrong"), key)
	logError(constants.ValueMissMatch, msg, data)
	return InternalError{Code: constants.ValueMissMatch, Message: msg, Data: data}
}

// ==========================
// AUTH ERRORS
// ==========================
func Unauthorized(t translator.Translator, data any) InternalError {
	msg := t.T("unauthorized")
	logError(constants.Unauthorized, msg, data)
	return InternalError{Code: constants.Unauthorized, Message: msg, Data: data}
}

func Forbidden(t translator.Translator, data any) InternalError {
	msg := t.T("forbidden")
	logError(constants.Forbidden, msg, data)
	return InternalError{Code: constants.Forbidden, Message: msg, Data: data}
}

func LoginError(t translator.Translator, data any) InternalError {
	msg := t.T("loginError")
	logError(constants.InvalidAuth, msg, data)
	return InternalError{Code: constants.InvalidAuth, Message: msg, Data: data}
}

func ExpiredError(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("expired"), key)
	logError(constants.ExpiredAuth, msg, data)
	return InternalError{Code: constants.ExpiredAuth, Message: msg, Data: data}
}

// ==========================
// GENERAL ERRORS
// ==========================
func GeneralError(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("generalRequestErrors"), key)
	logError(constants.GeneralErrors, msg, data)
	return InternalError{Code: constants.GeneralErrors, Message: msg, Data: data}
}

func InvalidProcess(t translator.Translator, key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", t.T("invalidProcess"), key)
	logError(constants.BadRequest, msg, data)
	return InternalError{Code: constants.BadRequest, Message: msg, Data: data}
}

func UnknownError(t translator.Translator, key string, data any) InternalError {
	msg := t.T("somethingWentWrong")
	logError(constants.InternalServerError, msg, data)
	return InternalError{Code: constants.InternalServerError, Message: msg, Data: data}
}
