package constants

// ======================================================
// Application-Level Error Codes (Internal Errors)
// ======================================================

// General Error
const (
	UnknownError         = 1000
	FailedFindResource   = 1001
	FailedCreateResource = 1002
	FailedUpdateResource = 1003
	FailedDeleteResource = 1004

	InvalidRequestData = 1010
	ValidationError    = 1011
	DuplicateValue     = 1012
	ResourceNotFound   = 1013
	OperationFailed    = 1014
	TimeoutError       = 1015
	UnauthorizedAction = 1016
	ForbiddenAccess    = 1017
)

// Database Layer
const (
	DBConnectionError  = 1101
	DBQueryError       = 1102
	DBTransactionError = 1103
	DBRecordNotFound   = 1104
	DBDuplicateKey     = 1105
)

// Repository / Cache
const (
	CacheConnectionError = 1201
	CacheReadError       = 1202
	CacheWriteError      = 1203
	CacheMissError       = 1204
)

// External / Integration
const (
	ExternalAPIError     = 1301
	ExternalTimeoutError = 1302
	ExternalParseError   = 1303
	ExternalAuthError    = 1304
	ExternalRateLimit    = 1305
)

// File / IO
const (
	FileNotFoundError = 1401
	FileReadError     = 1402
	FileWriteError    = 1403
	FileFormatError   = 1404
)

// Authentication & Authorization
const (
	AuthInvalidToken       = 1501
	AuthExpiredToken       = 1502
	AuthInvalidCredentials = 1503
	AuthPermissionDenied   = 1504
	AuthSessionExpired     = 1505
)

// Business Logic
const (
	InsufficientBalance = 1601
	QuotaExceeded       = 1602
	PaymentFailed       = 1603
	AlreadyProcessed    = 1604
	InvalidState        = 1605
	DependencyFailed    = 1606
)
