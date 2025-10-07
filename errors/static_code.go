package errors

// Global static codes (aplikasi)
const (
	CodeSuccess             = "SUCCESS"
	CodeBadRequest          = "BAD_REQUEST"
	CodeUnauthorized        = "UNAUTHORIZED"
	CodeForbidden           = "FORBIDDEN"
	CodeNotFound            = "NOT_FOUND"
	CodeConflict            = "CONFLICT"
	CodeUnprocessableEntity = "UNPROCESSABLE_ENTITY"
	CodeInternalError       = "INTERNAL_ERROR"
	CodeTimeout             = "TIMEOUT"

	// Operation related
	CodeResourceNotFound = "RESOURCE_NOT_FOUND"
	CodeFailedCreate     = "FAILED_CREATE"
	CodeFailedUpdate     = "FAILED_UPDATE"
	CodeFailedDelete     = "FAILED_DELETE"
	CodeInvalidInput     = "INVALID_INPUT"
	CodeDuplicateEntry   = "DUPLICATE_ENTRY"

	// DB / infra
	CodeDBConnection       = "DB_CONNECTION_ERROR"
	CodeMigrationError     = "MIGRATION_ERROR"
	CodeSerializationError = "SERIALIZATION_ERROR"
	CodeOptimisticLock     = "OPTIMISTIC_LOCK_FAILED"

	// External / payments / quota
	CodeExternalService = "EXTERNAL_SERVICE_ERROR"
	CodeRateLimit       = "RATE_LIMIT_EXCEEDED"
	CodeQuotaExceeded   = "QUOTA_EXCEEDED"
	CodePaymentFailed   = "PAYMENT_FAILED"

	// Permissions / features
	CodePermissionDenied = "PERMISSION_DENIED"
	CodeNotImplemented   = "NOT_IMPLEMENTED"
)
