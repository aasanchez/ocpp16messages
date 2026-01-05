package types

import (
	"errors"
)

// Error message format strings used throughout the types package.
const (
	// ErrorExpectedError is used in tests to format error assertion
	// messages.
	ErrorExpectedError = "Expected Error: %v"
	// ErrorMismatch is used in tests to format expected vs actual value
	// comparison failures.
	ErrorMismatch = "expected %q, got %q"
	// ErrorFieldFormat is the standard format string for wrapping errors
	// with field name context. The first parameter is the field name
	// (string), and the second parameter is the wrapped error. This format
	// enables error chain compatibility using fmt.Errorf with the %w verb.
	ErrorFieldFormat = "%s: %w"
)

// ErrEmptyValue is a sentinel error for empty field validation.
var ErrEmptyValue = errors.New("value cannot be empty")

// ErrInvalidValue is a sentinel error for invalid field validation.
var ErrInvalidValue = errors.New("invalid value")
