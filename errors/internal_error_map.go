package errors

import "net/http"

// Mapping recommended internal code -> HTTP status default
var InternalErrorToHTTP = map[string]int{
	// standard mapping
	CodeSuccess:             http.StatusOK,
	CodeBadRequest:          http.StatusBadRequest,
	CodeUnauthorized:        http.StatusUnauthorized,
	CodeForbidden:           http.StatusForbidden,
	CodeNotFound:            http.StatusNotFound,
	CodeConflict:            http.StatusConflict,
	CodeUnprocessableEntity: http.StatusUnprocessableEntity,
	CodeInternalError:       http.StatusInternalServerError,
	CodeTimeout:             http.StatusGatewayTimeout,

	// app-specific
	CodeResourceNotFound: http.StatusNotFound,
	CodeFailedCreate:     http.StatusInternalServerError,
	CodeFailedUpdate:     http.StatusInternalServerError,
	CodeFailedDelete:     http.StatusInternalServerError,
	CodeInvalidInput:     http.StatusBadRequest,
	CodeDuplicateEntry:   http.StatusConflict,

	// infra
	CodeDBConnection:       http.StatusInternalServerError,
	CodeMigrationError:     http.StatusInternalServerError,
	CodeSerializationError: http.StatusInternalServerError,
	CodeOptimisticLock:     http.StatusConflict,

	// externals
	CodeExternalService: http.StatusBadGateway, // 502
	CodeRateLimit:       http.StatusTooManyRequests,
	CodeQuotaExceeded:   http.StatusTooManyRequests,
	CodePaymentFailed:   http.StatusPaymentRequired, // 402

	CodePermissionDenied: http.StatusForbidden,
	CodeNotImplemented:   http.StatusNotImplemented,
}
