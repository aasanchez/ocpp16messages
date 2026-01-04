package types

import (
	"errors"
	"fmt"
)

// Maximum length constants for OCPP 1.6 CiString types.
const (
	CiString20Max  = 20
	CiString25Max  = 25
	CiString50Max  = 50
	CiString255Max = 255
	CiString500Max = 500
)

// ASCII printable character range boundaries.
const (
	asciiPrintableMin = 32  // Space character
	asciiPrintableMax = 126 // Tilde character
)

var (
	errExceedsMaxLength = errors.New("exceeds maximum length")
	errInvalidASCII     = errors.New(
		"CiString: value contains non-printable ASCII characters",
	)
)

type ciString struct {
	value string
}

func newCiString(input string, maxLen int) (ciString, error) {
	cis := ciString{value: input}

	if len(cis.value) > maxLen {
		return ciString{}, fmt.Errorf(
			"CiString Error on Construct (len=%d, max=%d): %w",
			len(cis.value),
			maxLen,
			errExceedsMaxLength,
		)
	}

	for _, r := range input {
		if r < asciiPrintableMin || r > asciiPrintableMax {
			return ciString{}, errInvalidASCII
		}
	}

	return cis, nil
}

func (cis ciString) val() string {
	return cis.value
}

// CiString20Type represents an OCPP 1.6 case-insensitive string with a maximum
// length of 20 characters. The string must contain only printable ASCII
// characters (32-126).
type CiString20Type struct {
	value ciString
}

// NewCiString20Type creates a new CiString20Type from a string value. Returns
// an error if the value exceeds 20 characters or contains non-printable ASCII.
func NewCiString20Type(value string) (CiString20Type, error) {
	cs, err := newCiString(value, CiString20Max)

	return CiString20Type{value: cs}, err
}

// Value returns the underlying string value of the CiString20Type.
func (c CiString20Type) Value() string {
	return c.value.val()
}

// CiString25Type represents an OCPP 1.6 case-insensitive string with a maximum
// length of 25 characters. The string must contain only printable ASCII
// characters (32-126).
type CiString25Type struct {
	value ciString
}

// NewCiString25Type creates a new CiString25Type from a string value. Returns
// an error if the value exceeds 25 characters or contains non-printable ASCII.
func NewCiString25Type(value string) (CiString25Type, error) {
	cs, err := newCiString(value, CiString25Max)

	return CiString25Type{value: cs}, err
}

// Value returns the underlying string value of the CiString25Type.
func (c CiString25Type) Value() string {
	return c.value.val()
}

// CiString50Type represents an OCPP 1.6 case-insensitive string with a maximum
// length of 50 characters. The string must contain only printable ASCII
// characters (32-126).
type CiString50Type struct {
	value ciString
}

// NewCiString50Type creates a new CiString50Type from a string value. Returns
// an error if the value exceeds 50 characters or contains non-printable ASCII.
func NewCiString50Type(value string) (CiString50Type, error) {
	cs, err := newCiString(value, CiString50Max)

	return CiString50Type{value: cs}, err
}

// Value returns the underlying string value of the CiString50Type.
func (c CiString50Type) Value() string {
	return c.value.val()
}

// CiString255Type represents an OCPP 1.6 case-insensitive string with a maximum
// length of 255 characters. The string must contain only printable ASCII
// characters (32-126).
type CiString255Type struct {
	value ciString
}

// NewCiString255Type creates a new CiString255Type from a string value. Returns
// an error if the value exceeds 255 characters or contains non-printable ASCII.
func NewCiString255Type(value string) (CiString255Type, error) {
	cs, err := newCiString(value, CiString255Max)

	return CiString255Type{value: cs}, err
}

// Value returns the underlying string value of the CiString255Type.
func (c CiString255Type) Value() string {
	return c.value.val()
}

// CiString500Type represents an OCPP 1.6 case-insensitive string with a maximum
// length of 500 characters. The string must contain only printable ASCII
// characters (32-126).
type CiString500Type struct {
	value ciString
}

// NewCiString500Type creates a new CiString500Type from a string value. Returns
// an error if the value exceeds 500 characters or contains non-printable ASCII.
func NewCiString500Type(value string) (CiString500Type, error) {
	cis, err := newCiString(value, CiString500Max)

	return CiString500Type{value: cis}, err
}

// Value returns the underlying string value of the CiString500Type.
func (c CiString500Type) Value() string {
	return c.value.val()
}
