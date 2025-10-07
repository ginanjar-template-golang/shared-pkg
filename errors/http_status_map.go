package errors

import "net/http"

// HTTPStatusMap: nama -> kode HTTP (lengkap umum 1xx - 5xx)
var HTTPStatusMap = map[string]int{
	// 1xx
	"CONTINUE":            http.StatusContinue,           //100
	"SWITCHING_PROTOCOLS": http.StatusSwitchingProtocols, //101
	"PROCESSING":          102,                           //102 WebDAV
	"EARLY_HINTS":         103,

	// 2xx
	"OK":                http.StatusOK,                   //200
	"CREATED":           http.StatusCreated,              //201
	"ACCEPTED":          http.StatusAccepted,             //202
	"NON_AUTHORITATIVE": http.StatusNonAuthoritativeInfo, //203
	"NO_CONTENT":        http.StatusNoContent,            //204
	"RESET_CONTENT":     http.StatusResetContent,         //205
	"PARTIAL_CONTENT":   http.StatusPartialContent,       //206
	"MULTI_STATUS":      207,
	"ALREADY_REPORTED":  208,
	"IM_USED":           226,

	// 3xx
	"MULTIPLE_CHOICES":   http.StatusMultipleChoices,  //300
	"MOVED_PERMANENTLY":  http.StatusMovedPermanently, //301
	"FOUND":              http.StatusFound,            //302
	"SEE_OTHER":          http.StatusSeeOther,         //303
	"NOT_MODIFIED":       http.StatusNotModified,      //304
	"USE_PROXY":          305,
	"TEMPORARY_REDIRECT": http.StatusTemporaryRedirect, //307
	"PERMANENT_REDIRECT": 308,

	// 4xx
	"BAD_REQUEST":                     http.StatusBadRequest,                   //400
	"UNAUTHORIZED":                    http.StatusUnauthorized,                 //401
	"PAYMENT_REQUIRED":                http.StatusPaymentRequired,              //402
	"FORBIDDEN":                       http.StatusForbidden,                    //403
	"NOT_FOUND":                       http.StatusNotFound,                     //404
	"METHOD_NOT_ALLOWED":              http.StatusMethodNotAllowed,             //405
	"NOT_ACCEPTABLE":                  http.StatusNotAcceptable,                //406
	"PROXY_AUTH_REQUIRED":             http.StatusProxyAuthRequired,            //407
	"REQUEST_TIMEOUT":                 http.StatusRequestTimeout,               //408
	"CONFLICT":                        http.StatusConflict,                     //409
	"GONE":                            http.StatusGone,                         //410
	"LENGTH_REQUIRED":                 http.StatusLengthRequired,               //411
	"PRECONDITION_FAILED":             http.StatusPreconditionFailed,           //412
	"PAYLOAD_TOO_LARGE":               http.StatusRequestEntityTooLarge,        //413
	"URI_TOO_LONG":                    http.StatusRequestURITooLong,            //414
	"UNSUPPORTED_MEDIA_TYPE":          http.StatusUnsupportedMediaType,         //415
	"RANGE_NOT_SATISFIABLE":           http.StatusRequestedRangeNotSatisfiable, //416
	"EXPECTATION_FAILED":              http.StatusExpectationFailed,            //417
	"IM_A_TEAPOT":                     418,
	"MISDIRECTED_REQUEST":             421,
	"UNPROCESSABLE_ENTITY":            http.StatusUnprocessableEntity, //422
	"LOCKED":                          423,
	"FAILED_DEPENDENCY":               424,
	"TOO_EARLY":                       425,
	"UPGRADE_REQUIRED":                426,
	"PRECONDITION_REQUIRED":           428,
	"TOO_MANY_REQUESTS":               http.StatusTooManyRequests, //429
	"REQUEST_HEADER_FIELDS_TOO_LARGE": 431,
	"UNAVAILABLE_FOR_LEGAL_REASONS":   451,

	// 5xx
	"INTERNAL_SERVER_ERROR":      http.StatusInternalServerError,     //500
	"NOT_IMPLEMENTED":            http.StatusNotImplemented,          //501
	"BAD_GATEWAY":                http.StatusBadGateway,              //502
	"SERVICE_UNAVAILABLE":        http.StatusServiceUnavailable,      //503
	"GATEWAY_TIMEOUT":            http.StatusGatewayTimeout,          //504
	"HTTP_VERSION_NOT_SUPPORTED": http.StatusHTTPVersionNotSupported, //505
	"VARIANT_ALSO_NEGOTIATES":    506,
	"INSUFFICIENT_STORAGE":       507,
	"LOOP_DETECTED":              508,
	"NOT_EXTENDED":               510,
	"NETWORK_AUTH_REQUIRED":      511,
}
