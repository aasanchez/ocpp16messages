package core

import "fmt"

// NewFieldError returns a descriptive error for a specific field validation failure.
// This is useful when validating OCPP message structures and data types.
//
// Example:
//
//	return NewFieldError("idTag", "must not be empty")
func NewFieldError(field string, reason string) error {
	return fmt.Errorf("invalid %s: %s", field, reason)
}
