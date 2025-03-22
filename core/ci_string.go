// Package core provides shared data types and validation utilities
// used throughout the OCPP 1.6J message structure.
//
// This file defines case-insensitive strings with length limits (CiString types)
// as specified in the OCPP 1.6J specification.
package core

import (
	"fmt"
	"strings"
)

// validateCiString validates that a string is non-empty and does not exceed
// the maximum allowed length. Used by all CiStringXXType types.
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
type CiString20Type string

// Validate checks that the CiString20Type value complies with its length constraint.
func (s CiString20Type) Validate() error {
	return validateCiString(string(s), 20, "CiString20Type")
}

// CiString25Type represents a case-insensitive string with a maximum length of 25 characters.
type CiString25Type string

// Validate checks that the CiString25Type value complies with its length constraint.
func (s CiString25Type) Validate() error {
	return validateCiString(string(s), 25, "CiString25Type")
}

// CiString50Type represents a case-insensitive string with a maximum length of 50 characters.
type CiString50Type string

// Validate checks that the CiString50Type value complies with its length constraint.
func (s CiString50Type) Validate() error {
	return validateCiString(string(s), 50, "CiString50Type")
}

// ToLower safely returns the lowercase value of any CiStringXXType.
func ToLower(s string) string {
	return strings.ToLower(s)
}
