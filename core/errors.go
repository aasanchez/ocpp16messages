// Package core defines common error types used across message validation and plugin logic.
package core

import "fmt"

// FieldError represents a validation error for a specific field.
type FieldError struct {
	Field  string // The name of the field with the error.
	Reason string // The reason why the field is considered invalid.
}

// Error returns the formatted error message for FieldError.
func (e *FieldError) Error() string {
	return fmt.Sprintf("invalid field '%s': %s", e.Field, e.Reason)
}

// NewFieldError creates a new FieldError instance for structured validation errors.
func NewFieldError(field string, reason error) error {
	return &FieldError{
		Field:  field,
		Reason: reason.Error(),
	}
}
