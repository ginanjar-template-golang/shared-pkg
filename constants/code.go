package constants

import "net/http"

// Standard HTTP Status
const (
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

	BadRequest          = http.StatusBadRequest          // 400
	Unauthorized        = http.StatusUnauthorized        // 401
	Forbidden           = http.StatusForbidden           // 403
	NotFound            = http.StatusNotFound            // 404
	InternalServerError = http.StatusInternalServerError // 500
)
