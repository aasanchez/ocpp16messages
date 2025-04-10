// Package types contain shared Open Charge Point Protocol (OCPP) 1.6J data structures used across both request and
// response messages. It includes reusable types that enforce constraints such as maximum string length and printable
// ASCII encoding.
package types

import (
	"errors"
	"fmt"
)

// Constants defining the maximum allowed lengths for different CiString types.
const (
	maxLenCiString20  = 20
	maxLenCiString25  = 25
	maxLenCiString50  = 50
	maxLenCiString255 = 255
	maxLenCiString500 = 500
)

// Predefined errors were returned during the validation of CiString values.
var (
	ErrExceedsMaxLength  = errors.New("value exceeds maximum allowed length")
	ErrNonPrintableASCII = errors.New("value contains non-printable ASCII characters")
)

// ciString is an internal representation of a case-insensitive string constrained by a maximum length and limited to
// printable ASCII characters. It is not exported and is used internally by all CiStringXXType types.
type ciString struct {
	Value  string
	MaxLen int
}

// CiString creates a new ciString instance with the specified value and maximum length. It returns an error if the
// value exceeds the max length or contains non-printable characters.
func CiString(value string, maxLen int) (ciString, error) {
	cs := ciString{Value: value, MaxLen: maxLen}
	if err := cs.validate(); err != nil {
		return ciString{}, err
	}

	return cs, nil
}

// validate checks that the ciString value:
//   - Does not exceed the specified maximum length.
//   - Contains only printable ASCII characters (decimal 32 to 126).
//
// Returns a wrapped static error in case of validation failure.
func (cs ciString) validate() error {
	if len(cs.Value) > cs.MaxLen {
		return fmt.Errorf("%w: actual length %d, max %d", ErrExceedsMaxLength, len(cs.Value), cs.MaxLen)
	}

	for _, r := range cs.Value {
		if r < 32 || r > 126 {
			return ErrNonPrintableASCII
		}
	}

	return nil
}

// String returns the underlying string value held in the ciString.
func (cs ciString) String() string { return cs.Value }

// CiString20Type

// CiString20Type is an OCPP 1.6J data type representing a case-insensitive string limited to a maximum of 20 printable
// ASCII characters. This type is commonly used for identifiers such as `idTag` in OCPP messages like Authorize.req.
type CiString20Type struct{ inner ciString }

// CiString20 creates a new instance of CiString20Type, validating that the provided value does not exceed
// 20 charactersand contains only printable ASCII characters. It returns a validated CiString20Type or an error if the
// constraints are violated.
func CiString20(value string) (CiString20Type, error) {
	cs, err := CiString(value, maxLenCiString20)

	return CiString20Type{cs}, err
}

// String returns the string value of the CiString20Type instance.
func (c CiString20Type) String() string { return c.inner.String() }

// Validate runs the internal validation logic on the CiString20Type instance. It ensures that the stored value still
// conforms to the OCPP 1.6J constraints.
func (c CiString20Type) Validate() error { return c.inner.validate() }

// CiString25Type

// CiString25Type is an OCPP 1.6J data type representing a case-insensitive string limited to a maximum of 25 printable
// ASCII characters. This type is commonly used for identifiers such as `idTag` in OCPP messages like Authorize.req.
type CiString25Type struct{ inner ciString }

// CiString25 creates a new instance of CiString25Type, validating that the provided value does not exceed
// 25 charactersand contains only printable ASCII characters. It returns a validated CiString25Type or an error if the
// constraints are violated.
func CiString25(value string) (CiString25Type, error) {
	cs, err := CiString(value, maxLenCiString25)

	return CiString25Type{cs}, err
}

// String returns the string value of the CiString25Type instance.
func (c CiString25Type) String() string { return c.inner.String() }

// Validate runs the internal validation logic on the CiString25Type instance. It ensures that the stored value still
// conforms to the OCPP 1.6J constraints.
func (c CiString25Type) Validate() error { return c.inner.validate() }

// CiString50Type

// CiString50Type is an OCPP 1.6J data type representing a case-insensitive string limited to a maximum of 50 printable
// ASCII characters. This type is commonly used for identifiers such as `idTag` in OCPP messages like Authorize.req.
type CiString50Type struct{ inner ciString }

// CiString50 creates a new instance of CiString50Type, validating that the provided value does not exceed
// 50 charactersand contains only printable ASCII characters. It returns a validated CiString50Type or an error if the
// constraints are violated.
func CiString50(value string) (CiString50Type, error) {
	cs, err := CiString(value, maxLenCiString50)

	return CiString50Type{cs}, err
}

// String returns the string value of the CiString50Type instance.
func (c CiString50Type) String() string { return c.inner.String() }

// Validate runs the internal validation logic on the CiString50Type instance. It ensures that the stored value still
// conforms to the OCPP 1.6J constraints.
func (c CiString50Type) Validate() error { return c.inner.validate() }

// CiString255Type

// CiString255Type is an OCPP 1.6J data type representing a case-insensitive string limited to a maximum of 255
// printable ASCII characters. This type is commonly used for identifiers such as `idTag` in OCPP messages like
// Authorize.req.
type CiString255Type struct{ inner ciString }

// CiString255 creates a new instance of CiString255Type, validating that the provided value does not exceed
// 255 charactersand contains only printable ASCII characters. It returns a validated CiString255Type or an error if the
// constraints are violated.
func CiString255(value string) (CiString255Type, error) {
	cs, err := CiString(value, maxLenCiString255)

	return CiString255Type{cs}, err
}

// String returns the string value of the CiString255Type instance.
func (c CiString255Type) String() string { return c.inner.String() }

// Validate runs the internal validation logic on the CiString255Type instance. It ensures that the stored value still
// conforms to the OCPP 1.6J constraints.
func (c CiString255Type) Validate() error { return c.inner.validate() }

// CiString500Type

// CiString500Type is an OCPP 1.6J data type representing a case-insensitive string limited to a maximum of 500
// printable ASCII characters. This type is commonly used for identifiers such as `idTag` in OCPP messages like
// Authorize.req.
type CiString500Type struct{ inner ciString }

// CiString500 creates a new instance of CiString500Type, validating that the provided value does not exceed
// 500 charactersand contains only printable ASCII characters. It returns a validated CiString500Type or an error if the
// constraints are violated.
func CiString500(value string) (CiString500Type, error) {
	cs, err := CiString(value, maxLenCiString500)

	return CiString500Type{cs}, err
}

// String returns the string value of the CiString500Type instance.
func (c CiString500Type) String() string { return c.inner.String() }

// Validate runs the internal validation logic on the CiString500Type instance. It ensures that the stored value still
// conforms to the OCPP 1.6J constraints.
func (c CiString500Type) Validate() error { return c.inner.validate() }
