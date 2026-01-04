package types

import (
	"errors"
)

// Test error message format strings.
const (
	// ErrorMethodMismatch is used in tests to format method return value
	// comparison failures. The first parameter is the actual value, and
	// the second parameter is the expected value.
	ErrorMethodMismatch = "AuthorizationStatus.String() = %v, want %v"
	// ErrorStatusMismatch is used in tests to format IdTagInfo.Status
	// comparison failures. The first parameter is the actual value, and
	// the second parameter is the expected value.
	ErrorStatusMismatch = "IdTagInfo.Status = %v, want %v"
)

// ErrInvalidIDTag is returned when an ID tag value is invalid.
var ErrInvalidIDTag = errors.New("invalid ID tag")
