package errors

import (
	"fmt"

	"github.com/ginanjar-template-golang/shared-pkg/constants"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
)

type InternalError struct {
	Code       int    `json:"code"`
	MessageKey string `json:"message_key"`
	Data       any    `json:"data,omitempty"`
}

// implementasi interface error
func (e InternalError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.MessageKey)
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
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("notFoundResource"), key)
	logError(constants.NotFound, msg, data)
	return InternalError{Code: constants.NotFound, MessageKey: "notFoundResource", Data: data}
}

func FindResourceError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("errorFindResource"), key)
	logError(constants.FailedFindResource, msg, data)
	return InternalError{Code: constants.FailedFindResource, MessageKey: "errorFindResource", Data: data}
}

func CreateResourceError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("errorCreateResource"), key)
	logError(constants.FailedCreateResource, msg, data)
	return InternalError{Code: constants.FailedCreateResource, MessageKey: "errorCreateResource", Data: data}
}

func UpdateResourceError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("errorUpdateResource"), key)
	logError(constants.FailedUpdateResource, msg, data)
	return InternalError{Code: constants.FailedUpdateResource, MessageKey: "errorUpdateResource", Data: data}
}

func DeleteResourceError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("errorDeleteResource"), key)
	logError(constants.FailedDeleteResource, msg, data)
	return InternalError{Code: constants.FailedDeleteResource, MessageKey: "errorDeleteResource", Data: data}
}

// ==========================
// VALIDATION ERRORS
// ==========================
func InvalidBody(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidBody"), key)
	logError(constants.BadRequest, msg, data)
	return InternalError{Code: constants.BadRequest, MessageKey: "invalidBody", Data: data}
}

func InvalidTypeError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidType"), key)
	logError(constants.BadRequestType, msg, data)
	return InternalError{Code: constants.BadRequestType, MessageKey: "invalidType", Data: data}
}

func InvalidFormatError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidFormat"), key)
	logError(constants.BadRequestFormat, msg, data)
	return InternalError{Code: constants.BadRequestFormat, MessageKey: "invalidFormat", Data: data}
}

func InCompleteKeyError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("incompleteKey"), key)
	logError(constants.BadRequestKey, msg, data)
	return InternalError{Code: constants.BadRequestKey, MessageKey: "incompleteKey", Data: data}
}

func InCompleteValueError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("incompleteValue"), key)
	logError(constants.BadRequestValue, msg, data)
	return InternalError{Code: constants.BadRequestValue, MessageKey: "incompleteValue", Data: data}
}

func AlreadyUsedError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("alreadyUsed"), key)
	logError(constants.DuplicateValue, msg, data)
	return InternalError{Code: constants.DuplicateValue, MessageKey: "alreadyUsed", Data: data}
}

func InvalidOptionError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidOption"), key)
	logError(constants.OutsideOption, msg, data)
	return InternalError{Code: constants.OutsideOption, MessageKey: "invalidOption", Data: data}
}

func ValueMissMatch(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("somethingWentWrong"), key)
	logError(constants.ValueMissMatch, msg, data)
	return InternalError{Code: constants.ValueMissMatch, MessageKey: "somethingWentWrong", Data: data}
}

// ==========================
// AUTH ERRORS
// ==========================
func Unauthorized(data any) InternalError {
	msg := translator.GetMessageByLang("unauthorized")
	logError(constants.Unauthorized, msg, data)
	return InternalError{Code: constants.Unauthorized, MessageKey: "unauthorized", Data: data}
}

func Forbidden(data any) InternalError {
	msg := translator.GetMessageByLang("forbidden")
	logError(constants.Forbidden, msg, data)
	return InternalError{Code: constants.Forbidden, MessageKey: "forbidden", Data: data}
}

func LoginError(data any) InternalError {
	msg := translator.GetMessageByLang("loginError")
	logError(constants.InvalidAuth, msg, data)
	return InternalError{Code: constants.InvalidAuth, MessageKey: "loginError", Data: data}
}

func ExpiredError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("expired"), key)
	logError(constants.ExpiredAuth, msg, data)
	return InternalError{Code: constants.ExpiredAuth, MessageKey: "expired", Data: data}
}

func InvalidTokenError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidToken"), key)
	logError(constants.InvalidToken, msg, data)
	return InternalError{Code: constants.InvalidToken, MessageKey: "invalidToken", Data: data}
}

// ==========================
// GENERAL ERRORS
// ==========================
func GeneralError(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("generalRequestErrors"), key)
	logError(constants.GeneralErrors, msg, data)
	return InternalError{Code: constants.GeneralErrors, MessageKey: "generalRequestErrors", Data: data}
}

func InvalidProcess(key string, data any) InternalError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidProcess"), key)
	logError(constants.BadRequest, msg, data)
	return InternalError{Code: constants.BadRequest, MessageKey: "invalidProcess", Data: data}
}

func UnknownError(key string, data any) InternalError {
	msg := translator.GetMessageByLang("somethingWentWrong")
	logError(constants.InternalServerError, msg, data)
	return InternalError{Code: constants.InternalServerError, MessageKey: "somethingWentWrong", Data: data}
}
