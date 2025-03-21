// Package models provides shared types used across OCPP 1.6 message definitions,
// including CiString types with strict length constraints and validation logic.
package models

import (
	"fmt"
)

// validateCiString checks whether a given string value is non-empty
// and does not exceed the specified maximum length.
//
// It returns an error if the validation fails.
// This helper is used internally by the CiStringXType types.
func validateCiString(value string, max int, field string) error {
	if len(value) == 0 {
		return fmt.Errorf("%s: cannot be empty", field)
	}
	if len(value) > max {
		return fmt.Errorf("%s: exceeds %d characters", field, max)
	}
	return nil
}

// CiString20Type represents a case-insensitive string with a maximum length of 20 characters.
// It is used in fields like ChargePointModel and ChargePointVendor.
type CiString20Type string

// Validate checks whether the CiString20Type value is non-empty and within the 20-character limit.
func (s CiString20Type) Validate() error {
	return validateCiString(string(s), 20, "CiString20Type")
}

// CiString25Type represents a case-insensitive string with a maximum length of 25 characters.
// It is commonly used for serial numbers, ICCID, IMSI, etc.
type CiString25Type string

// Validate checks whether the CiString25Type value is non-empty and within the 25-character limit.
func (s CiString25Type) Validate() error {
	return validateCiString(string(s), 25, "CiString25Type")
}

// CiString50Type represents a case-insensitive string with a maximum length of 50 characters.
// It is typically used for firmware version fields and other longer strings.
type CiString50Type string

// Validate checks whether the CiString50Type value is non-empty and within the 50-character limit.
func (s CiString50Type) Validate() error {
	return validateCiString(string(s), 50, "CiString50Type")
}
