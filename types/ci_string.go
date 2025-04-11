// Package types contains shared Open Charge Point Protocol (OCPP) 1.6J data structures
// used across both request and response messages.
//
// This package provides strongly typed wrappers for commonly reused values with built-in
// validation logic to enforce constraints defined in the OCPP 1.6J specification.
// It ensures that strings conform to maximum length limits and consist exclusively of
// printable ASCII characters. These constraints are essential for reliable communication
// between EV charge points and Central Systems.
//
// The provided types serve as reusable building blocks for building OCPP-compliant
// charging station implementations, simulators, proxies, and testing frameworks.
package types

import (
	"errors"
	"fmt"
)

// Constants defining the maximum allowed lengths for different CiString variants.
// These values are derived directly from the OCPP 1.6J specification.
const (
	maxLenCiString20  = 20
	maxLenCiString25  = 25
	maxLenCiString50  = 50
	maxLenCiString255 = 255
	maxLenCiString500 = 500
)

// Predefined errors returned during validation of CiString values.
var (
	// errExceedsMaxLength indicates that a value exceeded the maximum allowed length
	// for its CiString type.
	errExceedsMaxLength = errors.New("value exceeds maximum allowed length")

	// errNonPrintableASCII indicates that a value contains characters outside
	// the printable ASCII range (decimal 32–126 inclusive).
	errNonPrintableASCII = errors.New("value contains non-printable ASCII characters")

	// errEmptyValueNotAllowed indicates that a CiString value was expected but
	// found to be empty.
	errEmptyValueNotAllowed = errors.New("value must not be empty")
)

// ciString is an internal utility type representing a case-insensitive string
// with a maximum length and a constraint to only use printable ASCII characters.
//
// This type is not exported and is used to implement all public CiStringXXType types.
// It encapsulates the common validation logic used across all length-restricted strings
// in OCPP 1.6J.
type ciString struct {
	Value  string
	MaxLen int
}

// CiString creates and validates a new ciString instance.
// It returns an error if the value exceeds the configured maximum length
// or includes non-printable ASCII characters.
//
// Printable ASCII characters are in the inclusive range [32, 126].
func CiString(value string, maxLen int) (ciString, error) {
	cs := ciString{Value: value, MaxLen: maxLen}
	if err := cs.validate(); err != nil {
		return ciString{}, err
	}

	return cs, nil
}

// validate performs internal validation of a ciString instance.
//
// It ensures that:
//   - The string is not empty.
//   - The string does not exceed the maximum length.
//   - All characters are printable ASCII.
//
// If validation fails, a wrapped error is returned to indicate the specific failure.
func (cs ciString) validate() error {
	if len(cs.Value) == 0 {
		return errEmptyValueNotAllowed
	}

	if len(cs.Value) > cs.MaxLen {
		return fmt.Errorf("%w: actual length %d, max %d", errExceedsMaxLength, len(cs.Value), cs.MaxLen)
	}

	for _, r := range cs.Value {
		if r < 32 || r > 126 {
			return errNonPrintableASCII
		}
	}

	return nil
}

// String returns the string representation of the ciString.
func (cs ciString) String() string {
	return cs.Value
}

// CiString20Type is a case-insensitive string with a maximum length of 20 characters,
// consisting only of printable ASCII characters.
//
// It is used in OCPP 1.6J messages where a CiString[20] type is expected.
type CiString20Type struct{ inner ciString }

// CiString20 returns a validated CiString20Type.
//
// Returns an error if the input exceeds 20 characters or contains invalid characters.
func CiString20(value string) (CiString20Type, error) {
	cs, err := CiString(value, maxLenCiString20)

	return CiString20Type{cs}, err
}

// String returns the string value of the CiString20Type.
func (c CiString20Type) String() string {
	return c.inner.String()
}

// Validate re-validates the internal value of CiString20Type.
func (c CiString20Type) Validate() error {
	return c.inner.validate()
}

// CiString25Type is a case-insensitive string with a maximum length of 25 characters,
// used frequently for idTags in messages like Authorize.req.
//
// This type enforces OCPP 1.6J constraints to ensure robust message validation.
type CiString25Type struct{ inner ciString }

// CiString25 returns a validated CiString25Type instance.
func CiString25(value string) (CiString25Type, error) {
	cs, err := CiString(value, maxLenCiString25)

	return CiString25Type{cs}, err
}

// String returns the string value of the CiString25Type.
func (c CiString25Type) String() string {
	return c.inner.String()
}

// Validate ensures the CiString25Type value remains within spec.
func (c CiString25Type) Validate() error {
	return c.inner.validate()
}

// CiString50Type is a case-insensitive string with a maximum length of 50 characters,
// used in OCPP fields requiring slightly longer descriptive strings.
type CiString50Type struct{ inner ciString }

// CiString50 returns a validated CiString50Type instance.
func CiString50(value string) (CiString50Type, error) {
	cs, err := CiString(value, maxLenCiString50)

	return CiString50Type{cs}, err
}

// String returns the string value of the CiString50Type.
func (c CiString50Type) String() string {
	return c.inner.String()
}

// Validate ensures the CiString50Type value remains valid.
func (c CiString50Type) Validate() error {
	return c.inner.validate()
}

// CiString255Type is a case-insensitive string with a maximum length of 255 characters.
//
// It supports medium-sized string fields such as vendor identifiers or descriptions
// within OCPP messages.
type CiString255Type struct{ inner ciString }

// CiString255 returns a validated CiString255Type instance.
func CiString255(value string) (CiString255Type, error) {
	cs, err := CiString(value, maxLenCiString255)

	return CiString255Type{cs}, err
}

// String returns the string value of the CiString255Type.
func (c CiString255Type) String() string {
	return c.inner.String()
}

// Validate ensures the CiString255Type value remains valid.
func (c CiString255Type) Validate() error {
	return c.inner.validate()
}

// CiString500Type is a case-insensitive string with a maximum length of 500 characters.
//
// It is suitable for longer free-text fields or extended metadata values allowed by the
// OCPP 1.6J schema.
type CiString500Type struct{ inner ciString }

// CiString500 returns a validated CiString500Type instance.
func CiString500(value string) (CiString500Type, error) {
	cs, err := CiString(value, maxLenCiString500)

	return CiString500Type{cs}, err
}

// String returns the string value of the CiString500Type.
func (c CiString500Type) String() string {
	return c.inner.String()
}

// Validate ensures the CiString500Type value remains within spec.
func (c CiString500Type) Validate() error {
	return c.inner.validate()
}
