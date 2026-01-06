package types

import (
	"errors"
)

// Error message format strings used throughout the types package.
const (
	// ErrorExpectedError is used in tests to format error assertion messages.
	// Parameter: error value to be displayed.
	ErrorExpectedError = "Expected Error: %v"

	// ErrorFieldFormat is the standard format string for wrapping errors
	// with field name context. The first parameter is the field name
	// (string), and the second parameter is the wrapped error. This format
	// enables error chain compatibility using fmt.Errorf with the %w verb.
	ErrorFieldFormat = "%s: %w"

	// ErrorMismatch is used in tests to format expected vs actual value
	// comparison failures. The first parameter is the expected value,
	// and the second parameter is the actual value.
	ErrorMismatch = "Expected %q, got %q"

	// ErrorUnexpectedError is used in tests to format unexpected error
	// messages. Parameter: error value that was not expected.
	ErrorUnexpectedError = "Unexpected Error: %v"

	// ErrorWrapping is used in tests to verify error wrapping behavior.
	// First parameter is the outer error, second parameter is the expected
	// wrapped error.
	ErrorWrapping = "%v, wants wrapping %v"
)

var (
	// ErrEmptyValue is a sentinel error for empty field validation.
	ErrEmptyValue = errors.New("value cannot be empty")

	// ErrInvalidValue is a sentinel error for invalid field validation.
	ErrInvalidValue = errors.New("invalid value")
)
