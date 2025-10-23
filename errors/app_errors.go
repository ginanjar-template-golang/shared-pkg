package errors

import (
	"fmt"

	grpcCode "github.com/ginanjar-template-golang/shared-pkg/constants/grpc_code"
	httpCode "github.com/ginanjar-template-golang/shared-pkg/constants/http_code"
	internalCode "github.com/ginanjar-template-golang/shared-pkg/constants/internal_code"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
	"google.golang.org/grpc/codes"
)

// =======================================================
// STRUCT
// =======================================================
type AppError struct {
	HttpCode   int        `json:"http_code"`
	GrpcCode   codes.Code `json:"grpc_code"`
	MessageKey string     `json:"message_key"`
	Data       any        `json:"data,omitempty"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("[HTTP %d | GRPC %d] %s", e.HttpCode, e.GrpcCode, e.MessageKey)
}

// =======================================================
// RESOURCE ERRORS
// =======================================================
func ResourceNotFound(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("notFoundResource"), key)
	logger.LogMapLevel("warn", internalCode.ResourceNotFound, msg, data)
	return AppError{HttpCode: httpCode.NotFound, GrpcCode: grpcCode.NotFound, MessageKey: "notFoundResource", Data: data}
}

func FindResourceError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("errorFindResource"), key)
	logger.LogMapLevel("error", internalCode.FailedFindResource, msg, data)
	return AppError{HttpCode: httpCode.InternalServerError, GrpcCode: grpcCode.Internal, MessageKey: "errorFindResource", Data: data}
}

func CreateResourceError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("errorCreateResource"), key)
	logger.LogMapLevel("error", internalCode.FailedCreateResource, msg, data)
	return AppError{HttpCode: httpCode.InternalServerError, GrpcCode: grpcCode.Internal, MessageKey: "errorCreateResource", Data: data}
}

func UpdateResourceError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("errorUpdateResource"), key)
	logger.LogMapLevel("error", internalCode.FailedUpdateResource, msg, data)
	return AppError{HttpCode: httpCode.InternalServerError, GrpcCode: grpcCode.Internal, MessageKey: "errorUpdateResource", Data: data}
}

func DeleteResourceError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("errorDeleteResource"), key)
	logger.LogMapLevel("error", internalCode.FailedDeleteResource, msg, data)
	return AppError{HttpCode: httpCode.InternalServerError, GrpcCode: grpcCode.Internal, MessageKey: "errorDeleteResource", Data: data}
}

// =======================================================
// VALIDATION & REQUEST ERRORS
// =======================================================
func InvalidBody(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidBody"), key)
	logger.LogMapLevel("warn", internalCode.InvalidRequestData, msg, data)
	return AppError{HttpCode: httpCode.BadRequest, GrpcCode: grpcCode.InvalidArgument, MessageKey: "invalidBody", Data: data}
}

func InvalidTypeError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidType"), key)
	logger.LogMapLevel("warn", internalCode.ValidationError, msg, data)
	return AppError{HttpCode: httpCode.BadRequest, GrpcCode: grpcCode.InvalidArgument, MessageKey: "invalidType", Data: data}
}

func InvalidFormatError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidFormat"), key)
	logger.LogMapLevel("warn", internalCode.ValidationError, msg, data)
	return AppError{HttpCode: httpCode.BadRequest, GrpcCode: grpcCode.InvalidArgument, MessageKey: "invalidFormat", Data: data}
}

func AlreadyUsedError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("alreadyUsed"), key)
	logger.LogMapLevel("info", internalCode.DuplicateValue, msg, data)
	return AppError{HttpCode: httpCode.Conflict, GrpcCode: grpcCode.AlreadyExists, MessageKey: "alreadyUsed", Data: data}
}

func InvalidOptionError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidOption"), key)
	logger.LogMapLevel("warn", internalCode.OperationFailed, msg, data)
	return AppError{HttpCode: httpCode.BadRequest, GrpcCode: grpcCode.InvalidArgument, MessageKey: "invalidOption", Data: data}
}

func ValueMissMatch(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("valueMismatch"), key)
	logger.LogMapLevel("warn", internalCode.OperationFailed, msg, data)
	return AppError{HttpCode: httpCode.UnprocessableEntity, GrpcCode: grpcCode.FailedPrecondition, MessageKey: "valueMismatch", Data: data}
}

func ValidationFailed(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("validationFailed"), key)
	logger.LogMapLevel("info", internalCode.ValidationError, msg, data)
	return AppError{HttpCode: httpCode.BadRequest, GrpcCode: grpcCode.InvalidArgument, MessageKey: "validationFailed", Data: data}
}

// =======================================================
// AUTHENTICATION & AUTHORIZATION
// =======================================================
func Unauthorized(data any) AppError {
	msg := translator.GetMessageByLang("unauthorized")
	logger.LogMapLevel("warn", internalCode.AuthInvalidCredentials, msg, data)
	return AppError{HttpCode: httpCode.Unauthorized, GrpcCode: grpcCode.Unauthenticated, MessageKey: "unauthorized", Data: data}
}

func Forbidden(data any) AppError {
	msg := translator.GetMessageByLang("forbidden")
	logger.LogMapLevel("warn", internalCode.AuthPermissionDenied, msg, data)
	return AppError{HttpCode: httpCode.Forbidden, GrpcCode: grpcCode.PermissionDenied, MessageKey: "forbidden", Data: data}
}

func InvalidTokenError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("invalidToken"), key)
	logger.LogMapLevel("warn", internalCode.AuthInvalidToken, msg, data)
	return AppError{HttpCode: httpCode.Unauthorized, GrpcCode: grpcCode.Unauthenticated, MessageKey: "invalidToken", Data: data}
}

func ExpiredError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("expired"), key)
	logger.LogMapLevel("info", internalCode.AuthExpiredToken, msg, data)
	return AppError{HttpCode: httpCode.Unauthorized, GrpcCode: grpcCode.Unauthenticated, MessageKey: "expired", Data: data}
}

// =======================================================
// DATABASE & CACHE ERRORS
// =======================================================
func DatabaseError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("databaseError"), key)
	logger.LogMapLevel("error", internalCode.DBQueryError, msg, data)
	return AppError{HttpCode: httpCode.InternalServerError, GrpcCode: grpcCode.Internal, MessageKey: "databaseError", Data: data}
}

func DuplicateKeyError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("duplicateKey"), key)
	logger.LogMapLevel("info", internalCode.DBDuplicateKey, msg, data)
	return AppError{HttpCode: httpCode.Conflict, GrpcCode: grpcCode.AlreadyExists, MessageKey: "duplicateKey", Data: data}
}

func CacheError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("cacheError"), key)
	logger.LogMapLevel("error", internalCode.CacheReadError, msg, data)
	return AppError{HttpCode: httpCode.ServiceUnavailable, GrpcCode: grpcCode.Unavailable, MessageKey: "cacheError", Data: data}
}

// =======================================================
// EXTERNAL SERVICE / INTEGRATION
// =======================================================
func ExternalAPIError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("externalAPIError"), key)
	logger.LogMapLevel("error", internalCode.ExternalAPIError, msg, data)
	return AppError{HttpCode: httpCode.BadGateway, GrpcCode: grpcCode.Unavailable, MessageKey: "externalAPIError", Data: data}
}

func ExternalTimeoutError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("externalTimeout"), key)
	logger.LogMapLevel("warn", internalCode.ExternalTimeoutError, msg, data)
	return AppError{HttpCode: httpCode.GatewayTimeout, GrpcCode: grpcCode.DeadlineExceeded, MessageKey: "externalTimeout", Data: data}
}

func ExternalAuthError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("externalAuthError"), key)
	logger.LogMapLevel("warn", internalCode.ExternalAuthError, msg, data)
	return AppError{HttpCode: httpCode.Unauthorized, GrpcCode: grpcCode.Unauthenticated, MessageKey: "externalAuthError", Data: data}
}

func ExternalRateLimitError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("externalRateLimit"), key)
	logger.LogMapLevel("warn", internalCode.ExternalRateLimit, msg, data)
	return AppError{HttpCode: httpCode.ServiceUnavailable, GrpcCode: grpcCode.ResourceExhausted, MessageKey: "externalRateLimit", Data: data}
}

// =======================================================
// FILE & IO
// =======================================================
func FileNotFound(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("fileNotFound"), key)
	logger.LogMapLevel("warn", internalCode.FileNotFoundError, msg, data)
	return AppError{HttpCode: httpCode.NotFound, GrpcCode: grpcCode.NotFound, MessageKey: "fileNotFound", Data: data}
}

func FileReadError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("fileReadError"), key)
	logger.LogMapLevel("error", internalCode.FileReadError, msg, data)
	return AppError{HttpCode: httpCode.InternalServerError, GrpcCode: grpcCode.Internal, MessageKey: "fileReadError", Data: data}
}

// =======================================================
// TIMEOUT / GENERAL
// =======================================================
func TimeoutError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("timeoutError"), key)
	logger.LogMapLevel("warn", internalCode.TimeoutError, msg, data)
	return AppError{HttpCode: httpCode.GatewayTimeout, GrpcCode: grpcCode.DeadlineExceeded, MessageKey: "timeoutError", Data: data}
}

func GeneralError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("generalRequestErrors"), key)
	logger.LogMapLevel("error", internalCode.OperationFailed, msg, data)
	return AppError{HttpCode: httpCode.InternalServerError, GrpcCode: grpcCode.Internal, MessageKey: "generalRequestErrors", Data: data}
}

func UnknownError(key string, data any) AppError {
	msg := fmt.Sprintf("%s: %s", translator.GetMessageByLang("somethingWentWrong"), key)
	logger.LogMapLevel("error", internalCode.UnknownError, msg, data)
	return AppError{HttpCode: httpCode.InternalServerError, GrpcCode: grpcCode.Unknown, MessageKey: "somethingWentWrong", Data: data}
}
