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

// implementasi interface error
func (e InternalError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
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
// INTERNAL ERRORS
// ==========================
func ResourceNotFound(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("notFoundResource"), key)
	logError(constants.NotFound, msg, data)
	return InternalError{Code: constants.NotFound, Message: msg, Data: data}
}

func FindResourceError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("errorFindResource"), key)
	logError(constants.FailedFindResource, msg, data)
	return InternalError{Code: constants.FailedFindResource, Message: msg, Data: data}
}

func CreateResourceError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("errorCreateResource"), key)
	logError(constants.FailedCreateResource, msg, data)
	return InternalError{Code: constants.FailedCreateResource, Message: msg, Data: data}
}

func UpdateResourceError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("errorUpdateResource"), key)
	logError(constants.FailedUpdateResource, msg, data)
	return InternalError{Code: constants.FailedUpdateResource, Message: msg, Data: data}
}

func DeleteResourceError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("errorDeleteResource"), key)
	logError(constants.FailedDeleteResource, msg, data)
	return InternalError{Code: constants.FailedDeleteResource, Message: msg, Data: data}
}

// ==========================
// VALIDATION ERRORS
// ==========================
func InvalidBody(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("invalidBody"), key)
	logError(constants.BadRequest, msg, data)
	return InternalError{Code: constants.BadRequest, Message: msg, Data: data}
}

func InvalidTypeError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("invalidType"), key)
	logError(constants.BadRequestType, msg, data)
	return InternalError{Code: constants.BadRequestType, Message: msg, Data: data}
}

func InvalidFormatError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("invalidFormat"), key)
	logError(constants.BadRequestFormat, msg, data)
	return InternalError{Code: constants.BadRequestFormat, Message: msg, Data: data}
}

func InCompleteKeyError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("incompleteKey"), key)
	logError(constants.BadRequestKey, msg, data)
	return InternalError{Code: constants.BadRequestKey, Message: msg, Data: data}
}

func InCompleteValueError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("incompleteValue"), key)
	logError(constants.BadRequestValue, msg, data)
	return InternalError{Code: constants.BadRequestValue, Message: msg, Data: data}
}

func AlreadyUsedError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("alreadyUsed"), key)
	logError(constants.DuplicateValue, msg, data)
	return InternalError{Code: constants.DuplicateValue, Message: msg, Data: data}
}

func InvalidOptionError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("invalidOption"), key)
	logError(constants.OutsideOption, msg, data)
	return InternalError{Code: constants.OutsideOption, Message: msg, Data: data}
}

func ValueMissMatch(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("somethingWentWrong"), key)
	logError(constants.ValueMissMatch, msg, data)
	return InternalError{Code: constants.ValueMissMatch, Message: msg, Data: data}
}

// ==========================
// AUTH ERRORS
// ==========================
func Unauthorized(data any) InternalError {
	msg := translator.GetGlobalTranslator().T("unauthorized")
	logError(constants.Unauthorized, msg, data)
	return InternalError{Code: constants.Unauthorized, Message: msg, Data: data}
}

func Forbidden(data any) InternalError {
	msg := translator.GetGlobalTranslator().T("forbidden")
	logError(constants.Forbidden, msg, data)
	return InternalError{Code: constants.Forbidden, Message: msg, Data: data}
}

func LoginError(data any) InternalError {
	msg := translator.GetGlobalTranslator().T("loginError")
	logError(constants.InvalidAuth, msg, data)
	return InternalError{Code: constants.InvalidAuth, Message: msg, Data: data}
}

func ExpiredError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("expired"), key)
	logError(constants.ExpiredAuth, msg, data)
	return InternalError{Code: constants.ExpiredAuth, Message: msg, Data: data}
}

// ==========================
// GENERAL ERRORS
// ==========================
func GeneralError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("generalRequestErrors"), key)
	logError(constants.GeneralErrors, msg, data)
	return InternalError{Code: constants.GeneralErrors, Message: msg, Data: data}
}

func InvalidProcess(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetGlobalTranslator().T("invalidProcess"), key)
	logError(constants.BadRequest, msg, data)
	return InternalError{Code: constants.BadRequest, Message: msg, Data: data}
}

func UnknownError(key string, data any) InternalError {
	msg := translator.GetGlobalTranslator().T("somethingWentWrong")
	logError(constants.InternalServerError, msg, data)
	return InternalError{Code: constants.InternalServerError, Message: msg, Data: data}
}
