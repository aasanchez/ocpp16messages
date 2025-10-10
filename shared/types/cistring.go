package sharedtypes

import (
	"errors"
	"fmt"
)

// CiString types: printable ASCII strings for OCPP 1.6.
//
// Context
//   - OCPP 1.6 uses compact, bounded strings for many fields.
//   - Only printable ASCII (codes 32..126) is allowed to ensure safe
//     transport, logging, and predictable interop.
//   - Empty string is valid unless the specific message schema forbids it.
//
// Families (max lengths)
//   - CiString20Type  (20)
//   - CiString25Type  (25)
//   - CiString50Type  (50)
//   - CiString255Type (255)
//   - CiString500Type (500)
//
// Construction
//   - Use SetCiStringXXType to validate and build the value.
//   - Returns (value, error); never panics.
//   - Errors on non-ASCII or length overflow.
//
// Access
//   - Call Value() to retrieve the stored string.
//   - Values are immutable and safe to share across goroutines.
const (
	ci20  = 20
	ci25  = 25
	ci50  = 50
	ci255 = 255
	ci500 = 500
)

var (
	errExceedsMaxLength = errors.New("exceeds maximum length")
	errInvalidASCII     = errors.New("CiString has non-printable ASCII")
)

type ciString struct {
	value string
}

// setCiString validates length and ASCII range, then returns an internal value.
//
// Rules
//   - len(input) <= maxLen
//   - Only runes in [32, 126]
func setCiString(input string, maxLen int) (ciString, error) {
	cis := ciString{value: input}

	if len(cis.value) > maxLen {
		return ciString{}, fmt.Errorf(
			"CiString len=%d > max=%d (error): %w",
			len(cis.value),
			maxLen,
			errExceedsMaxLength,
		)
	}

	for _, r := range input {
		if r < 32 || r > 126 {
			return ciString{}, errInvalidASCII
		}
	}

	return cis, nil
}

func (cis ciString) val() string {
	return cis.value
}

// CiString20Type holds up to 20 printable ASCII characters.
//
// OCPP 1.6 usage examples
//   - IdTag and short identifiers/tokens
//   - Compact codes and status keys
//   - Short connector labels or names
//
// Notes
//   - Empty is allowed unless the payload forbids it.
//   - Prefer for small, frequently transported fields.
type CiString20Type struct {
	value ciString
}

// SetCiString20Type validates input for CiString20Type.
// Accepts: printable ASCII, length <= 20, empty ok.
// Rejects: non-ASCII, overflow. Returns error.
func SetCiString20Type(value string) (CiString20Type, error) {
	cs, err := setCiString(value, ci20)

	return CiString20Type{value: cs}, err
}

// Value returns the stored string.
func (c CiString20Type) Value() string {
	return c.value.val()
}

// CiString25Type holds up to 25 printable ASCII characters.
//
// OCPP 1.6 usage examples
//   - Vendor/Model-like descriptors that are short
//   - Firmware build tags or concise device labels
//   - Short site or group codes
//
// Notes
//   - Empty is allowed unless the payload forbids it.
//   - Use when 20 chars may be too tight but 50 is excessive.
type CiString25Type struct {
	value ciString
}

// SetCiString25Type validates input for CiString25Type.
// Accepts: printable ASCII, length <= 25, empty ok.
// Rejects: non-ASCII, overflow. Returns error.
func SetCiString25Type(value string) (CiString25Type, error) {
	cs, err := setCiString(value, ci25)

	return CiString25Type{value: cs}, err
}

// Value returns the stored string.
func (c CiString25Type) Value() string {
	return c.value.val()
}

// CiString50Type holds up to 50 printable ASCII characters.
//
// OCPP 1.6 usage examples
//   - FirmwareVersion, device serials, human labels
//   - Vendor error codes, info strings, model variants
//   - URLs or paths that must remain concise
//
// Notes
//   - Empty is allowed unless the payload forbids it.
//   - Good balance for medium-length descriptors.
type CiString50Type struct {
	value ciString
}

// SetCiString50Type validates input for CiString50Type.
// Accepts: printable ASCII, length <= 50, empty ok.
// Rejects: non-ASCII, overflow. Returns error.
func SetCiString50Type(value string) (CiString50Type, error) {
	cs, err := setCiString(value, ci50)

	return CiString50Type{value: cs}, err
}

// Value returns the stored string.
func (c CiString50Type) Value() string {
	return c.value.val()
}

// CiString255Type holds up to 255 printable ASCII characters.
//
// OCPP 1.6 usage examples
//   - Descriptions, diagnostics info, vendor messages
//   - Longer URLs, identifiers, or composite labels
//   - Free-form fields that stay within 255 chars
//
// Notes
//   - Empty is allowed unless the payload forbids it.
//   - Use when human-readable context is needed.
type CiString255Type struct {
	value ciString
}

// SetCiString255Type validates input for CiString255Type.
// Accepts: printable ASCII, length <= 255, empty ok.
// Rejects: non-ASCII, overflow. Returns error.
func SetCiString255Type(value string) (CiString255Type, error) {
	cs, err := setCiString(value, ci255)

	return CiString255Type{value: cs}, err
}

// Value returns the stored string.
func (c CiString255Type) Value() string {
	return c.value.val()
}

// CiString500Type holds up to 500 printable ASCII characters.
//
// OCPP 1.6 usage examples
//   - Extended diagnostics, troubleshooting notes
//   - Bulk vendor context or long free-form text
//   - Large URLs or serialized parameters (ASCII)
//
// Notes
//   - Empty is allowed unless the payload forbids it.
//   - Choose only when shorter bounds are insufficient.
type CiString500Type struct {
	value ciString
}

// SetCiString500Type validates input for CiString500Type.
// Accepts: printable ASCII, length <= 500, empty ok.
// Rejects: non-ASCII, overflow. Returns error.
func SetCiString500Type(value string) (CiString500Type, error) {
	cis, err := setCiString(value, ci500)

	return CiString500Type{value: cis}, err
}

// Value returns the stored string.
func (c CiString500Type) Value() string {
	return c.value.val()
}
