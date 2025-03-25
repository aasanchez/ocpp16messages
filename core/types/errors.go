package types

import "fmt"

// FieldError represents a validation error for a specific field.
type FieldError struct {
	Field  string
	Reason string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("invalid field '%s': %s", e.Field, e.Reason)
}

// NewFieldError creates a new FieldError with the given field and reason.
func NewFieldError(field string, reason string) *FieldError {
	return &FieldError{
		Field:  field,
		Reason: reason,
	}
}
