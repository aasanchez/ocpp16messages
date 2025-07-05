// Package types contains shared data types for OCPP 1.6 messages.
// These types are used across different OCPP messages.
package types

import (
	"fmt"
)

// This file defines custom Go types for handling Case-Insensitive Strings (ciString)
// as required by the OCPP 1.6 specification. The primary goal is to enforce
// constraints such as maximum length and the use of printable ASCII characters.
//
// The OCPP 1.6 specification mandates that certain fields, like IdToken, must be
// treated as case-insensitive. While this library provides the type validation,
// the actual case-insensitive comparison is the responsibility of the consuming
// application logic. For example, when comparing two `CiString20Type` values, the
// application should convert both to a common case (e.g., lowercase) before comparing.
//
// This file implements several fixed-size ciString types:
// - CiString20Type
// - CiString25Type
// - CiString50Type
// - CiString255Type
// - CiString500Type
//
// Each of these types wraps a common `ciString` struct, which contains the raw
// string and its maximum allowed length. The `validate` method on `ciString`
// performs all the necessary checks.

// Constants for maximum lengths of ciString types.
// These are based on the OCPP 1.6 specification for various fields.
const (
	maxLenCiString20Type  = 20  // Maximum length for a CiString20Type. Used for fields like IdToken.
	maxLenCiString25Type  = 25  // Maximum length for a CiString25Type.
	maxLenCiString50Type  = 50  // Maximum length for a CiString50Type.
	maxLenCiString255Type = 255 // Maximum length for a CiString255Type.
	maxLenCiString500Type = 500 // Maximum length for a CiString500Type.
)

// ciString is an internal, unexported struct that provides the core implementation
// for case-insensitive strings. It is not intended to be used directly from outside
// this package.
//
// See OCPP 1.6 specification, Appendix 3: "Data Types", for more information on
// case-insensitive strings.
type ciString struct {
	raw    string // The raw string value.
	maxLen int    // The maximum allowed length for the string.
}

// setCiString is an internal helper function to create and validate a ciString.
// It is used by the exported factory functions like `SetCiString20`.
func setCiString(value string, maxLen int) (ciString, error) {
	cs := ciString{raw: value, maxLen: maxLen}
	if err := cs.validate(); err != nil {
		return ciString{}, err
	}

	return cs, nil
}

// value returns the raw string value of the ciString.
func (cs ciString) value() string {
	return cs.raw
}

// validate checks if the ciString instance conforms to the OCPP 1.6 constraints.
//
// It performs the following checks:
// 1. Ensures the string is not empty.
// 2. Validates that the string length does not exceed the specified maximum.
// 3. Verifies that all characters are within the printable ASCII range (32-126).
func (cs ciString) validate() error {
	if len(cs.raw) == 0 {
		return fmt.Errorf("ciString.Validate: %w", ErrEmptyValueNotAllowed)
	}

	if len(cs.raw) > cs.maxLen {
		return fmt.Errorf("ciString.Validate: %w (got length %d, max %d)", ErrExceedsMaxLength, len(cs.raw), cs.maxLen)
	}

	for _, r := range cs.raw {
		if r < 32 || r > 126 {
			return fmt.Errorf("ciString.Validate: %w", ErrNonPrintableASCII)
		}
	}

	return nil
}

// --- Exported CiString Types ---

// CiString20 represents a case-insensitive string with a maximum length of 20 characters.
// It is used for fields like `IdToken` in OCPP `Authorize` and `StartTransaction` messages.
// Conforms to the `CiString20Type` in the OCPP 1.6 specification.
type CiString20Type struct{ value ciString }

// Value returns the string representation of the CiString20.
func (c CiString20Type) Value() string { return c.value.value() }

// Validate checks if the CiString20Type is valid according to OCPP 1.6 rules.
func (c CiString20Type) Validate() error { return c.value.validate() }

// SetCiString20Type creates a new CiString20Type from a standard string.
// It returns an error if the value is invalid.
func SetCiString20Type(value string) (CiString20Type, error) {
	cs, err := setCiString(value, maxLenCiString20Type)

	return CiString20Type{value: cs}, err
}

// CiString25Type represents a case-insensitive string with a maximum length of 25 characters.
// Conforms to the `CiString25Type` in the OCPP 1.6 specification.
type CiString25Type struct{ value ciString }

// Value returns the string representation of the CiString25Type.
func (c CiString25Type) Value() string { return c.value.value() }

// Validate checks if the CiString25Type is valid according to OCPP 1.6 rules.
func (c CiString25Type) Validate() error { return c.value.validate() }

// SetCiString25Type creates a new CiString25Type from a standard string.
// It returns an error if the value is invalid.
func SetCiString25Type(value string) (CiString25Type, error) {
	cs, err := setCiString(value, maxLenCiString25Type)

	return CiString25Type{value: cs}, err
}

// CiString50Type represents a case-insensitive string with a maximum length of 50 characters.
// Conforms to the `CiString50Type` in the OCPP 1.6 specification.
type CiString50Type struct{ value ciString }

// Value returns the string representation of the CiString50.
func (c CiString50Type) Value() string { return c.value.value() }

// Validate checks if the CiString50 is valid according to OCPP 1.6 rules.
func (c CiString50Type) Validate() error { return c.value.validate() }

// SetCiString50 creates a new CiString50 from a standard string.
// It returns an error if the value is invalid.
func SetCiString50Type(value string) (CiString50Type, error) {
	cs, err := setCiString(value, maxLenCiString50Type)

	return CiString50Type{value: cs}, err
}

// CiString255Type represents a case-insensitive string with a maximum length of 255 characters.
// Conforms to the `CiString255Type` in the OCPP 1.6 specification.
type CiString255Type struct{ value ciString }

// Value returns the string representation of the CiString255Type.
func (c CiString255Type) Value() string { return c.value.value() }

// Validate checks if the CiString255Type is valid according to OCPP 1.6 rules.
func (c CiString255Type) Validate() error { return c.value.validate() }

// SetCiString255Type creates a new CiString255Type from a standard string.
// It returns an error if the value is invalid.
func SetCiString255Type(value string) (CiString255Type, error) {
	cs, err := setCiString(value, maxLenCiString255Type)

	return CiString255Type{value: cs}, err
}

// CiString500Type represents a case-insensitive string with a maximum length of 500 characters.
// Conforms to the `CiString500Type` in the OCPP 1.6 specification.
type CiString500Type struct{ value ciString }

// Value returns the string representation of the CiString500Type.
func (c CiString500Type) Value() string { return c.value.value() }

// Validate checks if the CiString500Type is valid according to OCPP 1.6 rules.
func (c CiString500Type) Validate() error { return c.value.validate() }

// SetCiString500Type creates a new CiString500Type from a standard string.
// It returns an error if the value is invalid.
func SetCiString500Type(value string) (CiString500Type, error) {
	cs, err := setCiString(value, maxLenCiString500Type)

	return CiString500Type{value: cs}, err
}
