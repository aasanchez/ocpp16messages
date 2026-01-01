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
			"CiString Error on Construc (len=%d, max=%d): %w",
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

// CiString20 represents an OCPP 1.6 case-insensitive string with a maximum
// length of 20 characters. The string must contain only printable ASCII
// characters (32-126).
type CiString20 struct {
	value ciString
}

// NewCiString20 creates a new CiString20 from a string value. Returns
// an error if the value exceeds 20 characters or contains non-printable ASCII.
func NewCiString20(value string) (CiString20, error) {
	cs, err := newCiString(value, CiString20Max)

	return CiString20{value: cs}, err
}

// Value returns the underlying string value of the CiString20.
func (c CiString20) Value() string {
	return c.value.val()
}

// CiString25 represents an OCPP 1.6 case-insensitive string with a maximum
// length of 25 characters. The string must contain only printable ASCII
// characters (32-126).
type CiString25 struct {
	value ciString
}

// NewCiString25 creates a new CiString25 from a string value. Returns
// an error if the value exceeds 25 characters or contains non-printable ASCII.
func NewCiString25(value string) (CiString25, error) {
	cs, err := newCiString(value, CiString25Max)

	return CiString25{value: cs}, err
}

// Value returns the underlying string value of the CiString25.
func (c CiString25) Value() string {
	return c.value.val()
}

// CiString50 represents an OCPP 1.6 case-insensitive string with a maximum
// length of 50 characters. The string must contain only printable ASCII
// characters (32-126).
type CiString50 struct {
	value ciString
}

// NewCiString50 creates a new CiString50 from a string value. Returns
// an error if the value exceeds 50 characters or contains non-printable ASCII.
func NewCiString50(value string) (CiString50, error) {
	cs, err := newCiString(value, CiString50Max)

	return CiString50{value: cs}, err
}

// Value returns the underlying string value of the CiString50.
func (c CiString50) Value() string {
	return c.value.val()
}

// CiString255 represents an OCPP 1.6 case-insensitive string with a maximum
// length of 255 characters. The string must contain only printable ASCII
// characters (32-126).
type CiString255 struct {
	value ciString
}

// NewCiString255 creates a new CiString255 from a string value. Returns
// an error if the value exceeds 255 characters or contains non-printable ASCII.
func NewCiString255(value string) (CiString255, error) {
	cs, err := newCiString(value, CiString255Max)

	return CiString255{value: cs}, err
}

// Value returns the underlying string value of the CiString255.
func (c CiString255) Value() string {
	return c.value.val()
}

// CiString500 represents an OCPP 1.6 case-insensitive string with a maximum
// length of 500 characters. The string must contain only printable ASCII
// characters (32-126).
type CiString500 struct {
	value ciString
}

// NewCiString500 creates a new CiString500 from a string value. Returns
// an error if the value exceeds 500 characters or contains non-printable ASCII.
func NewCiString500(value string) (CiString500, error) {
	cis, err := newCiString(value, CiString500Max)

	return CiString500{value: cis}, err
}

// Value returns the underlying string value of the CiString500.
func (c CiString500) Value() string {
	return c.value.val()
}
