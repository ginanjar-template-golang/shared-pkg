package constants

import "net/http"

// ======================================================
// Standard HTTP Status Codes
// ======================================================
const (
	// -----------------------------
	// 1xx Informational
	// -----------------------------
	Continue           = http.StatusContinue           // 100
	SwitchingProtocols = http.StatusSwitchingProtocols // 101
	Processing         = http.StatusProcessing         // 102
	EarlyHints         = http.StatusEarlyHints         // 103

	// -----------------------------
	// 2xx Success
	// -----------------------------
	SuccessOK                   = http.StatusOK                   // 200
	SuccessCreated              = http.StatusCreated              // 201
	SuccessAccepted             = http.StatusAccepted             // 202
	NonAuthoritativeInformation = http.StatusNonAuthoritativeInfo // 203
	NoContent                   = http.StatusNoContent            // 204
	ResetContent                = http.StatusResetContent         // 205
	PartialContent              = http.StatusPartialContent       // 206
	MultiStatus                 = http.StatusMultiStatus          // 207
	AlreadyReported             = http.StatusAlreadyReported      // 208
	IMUsed                      = http.StatusIMUsed               // 226

	// -----------------------------
	// 3xx Redirection
	// -----------------------------
	MultipleChoices   = http.StatusMultipleChoices   // 300
	MovedPermanently  = http.StatusMovedPermanently  // 301
	Found             = http.StatusFound             // 302
	SeeOther          = http.StatusSeeOther          // 303
	NotModified       = http.StatusNotModified       // 304
	UseProxy          = http.StatusUseProxy          // 305
	TemporaryRedirect = http.StatusTemporaryRedirect // 307
	PermanentRedirect = http.StatusPermanentRedirect // 308

	// -----------------------------
	// 4xx Client Error
	// -----------------------------
	BadRequest                   = http.StatusBadRequest                   // 400
	Unauthorized                 = http.StatusUnauthorized                 // 401
	PaymentRequired              = http.StatusPaymentRequired              // 402
	Forbidden                    = http.StatusForbidden                    // 403
	NotFound                     = http.StatusNotFound                     // 404
	MethodNotAllowed             = http.StatusMethodNotAllowed             // 405
	NotAcceptable                = http.StatusNotAcceptable                // 406
	ProxyAuthRequired            = http.StatusProxyAuthRequired            // 407
	RequestTimeout               = http.StatusRequestTimeout               // 408
	Conflict                     = http.StatusConflict                     // 409
	Gone                         = http.StatusGone                         // 410
	LengthRequired               = http.StatusLengthRequired               // 411
	PreconditionFailed           = http.StatusPreconditionFailed           // 412
	RequestEntityTooLarge        = http.StatusRequestEntityTooLarge        // 413
	RequestURITooLong            = http.StatusRequestURITooLong            // 414
	UnsupportedMediaType         = http.StatusUnsupportedMediaType         // 415
	RequestedRangeNotSatisfiable = http.StatusRequestedRangeNotSatisfiable // 416
	ExpectationFailed            = http.StatusExpectationFailed            // 417
	Teapot                       = http.StatusTeapot                       // 418
	MisdirectedRequest           = http.StatusMisdirectedRequest           // 421
	UnprocessableEntity          = http.StatusUnprocessableEntity          // 422
	Locked                       = http.StatusLocked                       // 423
	FailedDependency             = http.StatusFailedDependency             // 424
	TooEarly                     = http.StatusTooEarly                     // 425
	UpgradeRequired              = http.StatusUpgradeRequired              // 426
	PreconditionRequired         = http.StatusPreconditionRequired         // 428
	TooManyRequests              = http.StatusTooManyRequests              // 429
	RequestHeaderFieldsTooLarge  = http.StatusRequestHeaderFieldsTooLarge  // 431
	UnavailableForLegalReasons   = http.StatusUnavailableForLegalReasons   // 451

	// -----------------------------
	// 5xx Server Error
	// -----------------------------
	InternalServerError           = http.StatusInternalServerError           // 500
	NotImplemented                = http.StatusNotImplemented                // 501
	BadGateway                    = http.StatusBadGateway                    // 502
	ServiceUnavailable            = http.StatusServiceUnavailable            // 503
	GatewayTimeout                = http.StatusGatewayTimeout                // 504
	HTTPVersionNotSupported       = http.StatusHTTPVersionNotSupported       // 505
	VariantAlsoNegotiates         = http.StatusVariantAlsoNegotiates         // 506
	InsufficientStorage           = http.StatusInsufficientStorage           // 507
	LoopDetected                  = http.StatusLoopDetected                  // 508
	NotExtended                   = http.StatusNotExtended                   // 510
	NetworkAuthenticationRequired = http.StatusNetworkAuthenticationRequired // 511
)
