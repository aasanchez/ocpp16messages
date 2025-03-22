// Package core defines reusable types used across multiple OCPP 1.6 messages,
// including length-restricted strings as specified by the protocol.
package core

import (
	"fmt"
)

// validateCiString checks the constraints of a CiString (length-limited string).
func validateCiString(value string, max int, field string) error {
	if len(value) == 0 {
		return fmt.Errorf("%s: cannot be empty", field)
	}
	if len(value) > max {
		return fmt.Errorf("%s: exceeds %d characters", field, max)
	}
	return nil
}

// CiString20Type represents a string limited to 20 characters, per OCPP 1.6 spec.
type CiString20Type string

// Validate returns an error if the string is empty or exceeds 20 characters.
func (s CiString20Type) Validate() error {
	return validateCiString(string(s), 20, "CiString20Type")
}

// CiString25Type represents a string limited to 25 characters, per OCPP 1.6 spec.
type CiString25Type string

// Validate returns an error if the string is empty or exceeds 25 characters.
func (s CiString25Type) Validate() error {
	return validateCiString(string(s), 25, "CiString25Type")
}

// CiString50Type represents a string limited to 50 characters, per OCPP 1.6 spec.
type CiString50Type string

// Validate returns an error if the string is empty or exceeds 50 characters.
func (s CiString50Type) Validate() error {
	return validateCiString(string(s), 50, "CiString50Type")
}
