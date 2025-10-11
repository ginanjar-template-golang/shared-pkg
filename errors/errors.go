package errors

import (
	"fmt"

	"github.com/ginanjar-template-golang/shared-pkg/constants"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
	"github.com/ginanjar-template-golang/shared-pkg/utils"
)

type AppError struct {
	Code       int    `json:"code"`
	MessageKey string `json:"message_key"`
	Data       any    `json:"data,omitempty"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.MessageKey)
}

// ===============================================
// LOG HELPER (LEVEL AWARE)
// ===============================================
func logWithLevel(level string, internalCode int, msg string, data any, err ...error) {
	fields := map[string]any{
		"request_id":    utils.GetRequestID(),
		"internal_code": internalCode,
		"data":          data,
	}

	switch level {
	case "debug":
		logger.Debug(msg, fields, err...)
	case "info":
		logger.Info(msg, fields)
	case "warn":
		logger.Warn(msg, fields, err...)
	default:
		logger.Error(msg, fields, err...)
	}
}

// ===============================================
// RESOURCE ERRORS
// ===============================================
func ResourceNotFound(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("notFoundResource"), key)
	logWithLevel("warn", constants.ResourceNotFound, msg, data)
	return AppError{Code: constants.NotFound, MessageKey: "notFoundResource", Data: data}
}

func FindResourceError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("errorFindResource"), key)
	logWithLevel("error", constants.FailedFindResource, msg, data)
	return AppError{Code: constants.InternalServerError, MessageKey: "errorFindResource", Data: data}
}

func CreateResourceError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("errorCreateResource"), key)
	logWithLevel("error", constants.FailedCreateResource, msg, data)
	return AppError{Code: constants.InternalServerError, MessageKey: "errorCreateResource", Data: data}
}

func UpdateResourceError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("errorUpdateResource"), key)
	logWithLevel("error", constants.FailedUpdateResource, msg, data)
	return AppError{Code: constants.InternalServerError, MessageKey: "errorUpdateResource", Data: data}
}

func DeleteResourceError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("errorDeleteResource"), key)
	logWithLevel("error", constants.FailedDeleteResource, msg, data)
	return AppError{Code: constants.InternalServerError, MessageKey: "errorDeleteResource", Data: data}
}

// ===============================================
// VALIDATION & REQUEST ERRORS
// ===============================================
func InvalidBody(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidBody"), key)
	logWithLevel("warn", constants.InvalidRequestData, msg, data)
	return AppError{Code: constants.BadRequest, MessageKey: "invalidBody", Data: data}
}

func InvalidTypeError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidType"), key)
	logWithLevel("warn", constants.ValidationError, msg, data)
	return AppError{Code: constants.BadRequest, MessageKey: "invalidType", Data: data}
}

func InvalidFormatError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidFormat"), key)
	logWithLevel("warn", constants.ValidationError, msg, data)
	return AppError{Code: constants.BadRequest, MessageKey: "invalidFormat", Data: data}
}

func AlreadyUsedError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("alreadyUsed"), key)
	logWithLevel("info", constants.DuplicateValue, msg, data)
	return AppError{Code: constants.Conflict, MessageKey: "alreadyUsed", Data: data}
}

func InvalidOptionError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidOption"), key)
	logWithLevel("warn", constants.OperationFailed, msg, data)
	return AppError{Code: constants.BadRequest, MessageKey: "invalidOption", Data: data}
}

func ValueMissMatch(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("valueMismatch"), key)
	logWithLevel("warn", constants.OperationFailed, msg, data)
	return AppError{Code: constants.UnprocessableEntity, MessageKey: "valueMismatch", Data: data}
}

// ===============================================
// AUTHENTICATION & AUTHORIZATION
// ===============================================
func Unauthorized(data any) AppError {
	msg := translator.GetMessageByLang("unauthorized")
	logWithLevel("warn", constants.AuthInvalidCredentials, msg, data)
	return AppError{Code: constants.Unauthorized, MessageKey: "unauthorized", Data: data}
}

func Forbidden(data any) AppError {
	msg := translator.GetMessageByLang("forbidden")
	logWithLevel("warn", constants.AuthPermissionDenied, msg, data)
	return AppError{Code: constants.Forbidden, MessageKey: "forbidden", Data: data}
}

func InvalidTokenError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidToken"), key)
	logWithLevel("warn", constants.AuthInvalidToken, msg, data)
	return AppError{Code: constants.Unauthorized, MessageKey: "invalidToken", Data: data}
}

func ExpiredError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("expired"), key)
	logWithLevel("info", constants.AuthExpiredToken, msg, data)
	return AppError{Code: constants.Unauthorized, MessageKey: "expired", Data: data}
}

// ===============================================
// DATABASE & CACHE ERRORS
// ===============================================
func DatabaseError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("databaseError"), key)
	logWithLevel("error", constants.DBQueryError, msg, data)
	return AppError{Code: constants.InternalServerError, MessageKey: "databaseError", Data: data}
}

func DuplicateKeyError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("duplicateKey"), key)
	logWithLevel("info", constants.DBDuplicateKey, msg, data)
	return AppError{Code: constants.Conflict, MessageKey: "duplicateKey", Data: data}
}

func CacheError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("cacheError"), key)
	logWithLevel("error", constants.CacheReadError, msg, data)
	return AppError{Code: constants.ServiceUnavailable, MessageKey: "cacheError", Data: data}
}

// ===============================================
// EXTERNAL SERVICE / INTEGRATION
// ===============================================
func ExternalAPIError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("externalAPIError"), key)
	logWithLevel("error", constants.ExternalAPIError, msg, data)
	return AppError{Code: constants.BadGateway, MessageKey: "externalAPIError", Data: data}
}

func ExternalTimeoutError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("externalTimeout"), key)
	logWithLevel("warn", constants.ExternalTimeoutError, msg, data)
	return AppError{Code: constants.GatewayTimeout, MessageKey: "externalTimeout", Data: data}
}

func ExternalAuthError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("externalAuthError"), key)
	logWithLevel("warn", constants.ExternalAuthError, msg, data)
	return AppError{Code: constants.Unauthorized, MessageKey: "externalAuthError", Data: data}
}

func ExternalRateLimitError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("externalRateLimit"), key)
	logWithLevel("warn", constants.ExternalRateLimit, msg, data)
	return AppError{Code: constants.ServiceUnavailable, MessageKey: "externalRateLimit", Data: data}
}

// ===============================================
// FILE & IO
// ===============================================
func FileNotFound(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("fileNotFound"), key)
	logWithLevel("warn", constants.FileNotFoundError, msg, data)
	return AppError{Code: constants.NotFound, MessageKey: "fileNotFound", Data: data}
}

func FileReadError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("fileReadError"), key)
	logWithLevel("error", constants.FileReadError, msg, data)
	return AppError{Code: constants.InternalServerError, MessageKey: "fileReadError", Data: data}
}

// ===============================================
// TIMEOUT / GENERAL
// ===============================================
func TimeoutError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("timeoutError"), key)
	logWithLevel("warn", constants.TimeoutError, msg, data)
	return AppError{Code: constants.GatewayTimeout, MessageKey: "timeoutError", Data: data}
}

func GeneralError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("generalRequestErrors"), key)
	logWithLevel("error", constants.OperationFailed, msg, data)
	return AppError{Code: constants.InternalServerError, MessageKey: "generalRequestErrors", Data: data}
}

func UnknownError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("somethingWentWrong"), key)
	logWithLevel("error", constants.UnknownError, msg, data)
	return AppError{Code: constants.InternalServerError, MessageKey: "somethingWentWrong", Data: data}
}
