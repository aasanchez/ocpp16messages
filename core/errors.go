// Package core provides foundational types and utilities shared across the OCPP 1.6 message handling.
package core

import (
	"fmt"
)

// FieldError represents a validation error for a specific field in a message.
type FieldError struct {
	Field string // The name of the field that failed validation.
	Err   error  // The original error that occurred during validation.
}

// Error returns the string representation of the FieldError.
func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %v", e.Field, e.Err)
}

// Unwrap allows errors.Unwrap to retrieve the underlying error.
func (e *FieldError) Unwrap() error {
	return e.Err
}

// NewFieldError creates a new FieldError for the given field and error.
func NewFieldError(field string, err error) error {
	if err == nil {
		return nil
	}
	return &FieldError{
		Field: field,
		Err:   err,
	}
}
