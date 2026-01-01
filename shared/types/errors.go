package types

import (
	"errors"
	"fmt"
)

// Error message format strings used throughout the shared types package.
const (
	// ErrorExpectedError is used in tests to format error assertion messages.
	ErrorExpectedError  = "Expected Error: %v"
	ErrorStringMismatch = "Expected String: %s, got: %s"
	ErrorValueMismatch  = "expected Value = %q, got %q"
)

// ErrEmptyValue is a sentinel error for empty field validation.
var ErrEmptyValue = errors.New("value cannot be empty")

// ErrEmpty returns an error indicating that the specified field cannot
// be empty. It wraps ErrEmptyValue for error chain compatibility.
func ErrEmpty(fieldName string) error {
	return fmt.Errorf("%s: %w", fieldName, ErrEmptyValue)
}
