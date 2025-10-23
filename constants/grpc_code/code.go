package grpc_code

import "google.golang.org/grpc/codes"

// ======================================================
// Standard gRPC Status Codes
// ======================================================
//
// Referensi resmi:
// https://pkg.go.dev/google.golang.org/grpc/codes
const (
	// -----------------------------
	// OK
	// -----------------------------
	OK = codes.OK // 0

	// -----------------------------
	// Canceled / Unknown
	// -----------------------------
	Canceled = codes.Canceled // 1: client canceled request
	Unknown  = codes.Unknown  // 2: unknown error occurred

	// -----------------------------
	// Invalid arguments / validation
	// -----------------------------
	InvalidArgument  = codes.InvalidArgument  // 3: invalid parameter or validation error
	DeadlineExceeded = codes.DeadlineExceeded // 4: deadline exceeded

	// -----------------------------
	// Resource and Auth errors
	// -----------------------------
	NotFound          = codes.NotFound          // 5: resource not found
	AlreadyExists     = codes.AlreadyExists     // 6: resource already exists
	PermissionDenied  = codes.PermissionDenied  // 7: access denied
	ResourceExhausted = codes.ResourceExhausted // 8: quota exceeded or rate limited

	// -----------------------------
	// Precondition / State errors
	// -----------------------------
	FailedPrecondition = codes.FailedPrecondition // 9: precondition not met
	Aborted            = codes.Aborted            // 10: operation aborted
	OutOfRange         = codes.OutOfRange         // 11: numeric value out of range

	// -----------------------------
	// Not implemented / Internal errors
	// -----------------------------
	Unimplemented = codes.Unimplemented // 12: not implemented
	Internal      = codes.Internal      // 13: internal server error
	Unavailable   = codes.Unavailable   // 14: service unavailable
	DataLoss      = codes.DataLoss      // 15: unrecoverable data loss

	// -----------------------------
	// Auth errors
	// -----------------------------
	Unauthenticated = codes.Unauthenticated // 16: missing or invalid credentials
)
