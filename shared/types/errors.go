package types

import (
	"errors"
)

// Error message format strings used throughout the types package.
const (
	// ErrorFieldFormat is the standard format string for wrapping errors
	// with field name context. The first parameter is the field name
	// (string), and the second parameter is the wrapped error. This format
	// enables error chain compatibility using fmt.Errorf with the %w verb.
	ErrorExpectedError   = "Expected Error: %v"
	ErrorFieldFormat     = "%s: %w"
	ErrorMismatch        = "Expected %q, got %q"
	ErrorUnexpectedError = "Unexpected Error: %v"
	ErrorWrapping        = "%v, wants wrapping %v"
)

// ErrEmptyValue is a sentinel error for empty field validation.
var ErrEmptyValue = errors.New("value cannot be empty")

// ErrInvalidValue is a sentinel error for invalid field validation.
var ErrInvalidValue = errors.New("invalid value")
