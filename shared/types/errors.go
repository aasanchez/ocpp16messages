package types

import (
	"errors"
	"fmt"
)

// Error message format strings used throughout the types package.
const (
	// ErrorExpectedError is used in tests to format error assertion
	// messages.
	ErrorExpectedError = "Expected Error: %v"
	// ErrorMismatch is used in tests to format expected vs actual value
	// comparison failures.
	ErrorMismatch = "expected %q, got %q"
)

// ErrEmptyValue is a sentinel error for empty field validation.
var ErrEmptyValue = errors.New("value cannot be empty")

// ErrInvalidValue is a sentinel error for invalid field validation.
var ErrInvalidValue = errors.New("invalid value")

// ErrEmpty returns an error indicating that the specified field cannot
// be empty. It wraps ErrEmptyValue for error chain compatibility.
func ErrEmpty(fieldName string) error {
	return fmt.Errorf("%s: %w", fieldName, ErrEmptyValue)
}

// ErrInvalid returns an error indicating that the specified field has an
// invalid value. It wraps ErrInvalidValue for error chain compatibility.
func ErrInvalid(fieldName string) error {
	return fmt.Errorf("%s: %w", fieldName, ErrInvalidValue)
}
