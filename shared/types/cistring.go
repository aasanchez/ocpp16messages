// Package sharedtypes provides common types and utilities for the OCPP 1.6J protocol.
//
// This package defines custom Go types for handling Case-Insensitive Strings (CiString)
// as required by the OCPP 1.6J specification. The primary goal is to enforce
// constraints such as maximum length and the use of printable ASCII characters
// for various string-based fields in OCPP messages.
//
// The OCPP 1.6J specification mandates that certain fields, like `IdToken`,
// must be treated as case-insensitive. While this library provides type validation
// for length and character set, the actual case-insensitive comparison logic
// (e.g., converting to lowercase before comparison) is the responsibility of the
// consuming application.
//
// This file implements several fixed-size CiString types, each corresponding
// to a specific maximum length defined in the OCPP 1.6J specification:
//   - CiString20Type (max 20 characters)
//   - CiString25Type (max 25 characters)
//   - CiString50Type (max 50 characters)
//   - CiString255Type (max 255 characters)
//   - CiString500Type (max 500 characters)
//
// Each of these exported types wraps an internal `ciString` struct, which handles
// the core validation logic. All string values are expected to be composed of
// printable ASCII characters (ASCII values 32-126).
//
// See OCPP 1.6J Part 2, Appendix 3: "Data Types" for detailed definitions
// and constraints of these string types.
package sharedtypes

import (
	"fmt"
)

// Constants defining the maximum allowed lengths for various CiString types
// as specified in the OCPP 1.6J protocol.
const (
	// maxLenCiString20Type is the maximum length for a CiString20Type.
	// Used for fields like `IdToken` (Authorize.req - Table 12, StartTransaction.req - Table 36).
	maxLenCiString20Type = 20
	// maxLenCiString25Type is the maximum length for a CiString25Type.
	// Used for fields like `chargePointVendor` (BootNotification.req - Table 1).
	maxLenCiString25Type = 25
	// maxLenCiString50Type is the maximum length for a CiString50Type.
	// Used for fields like `chargePointModel` (BootNotification.req - Table 1).
	maxLenCiString50Type = 50
	// maxLenCiString255Type is the maximum length for a CiString255Type.
	// Used for fields like `firmwareVersion` (BootNotification.req - Table 1).
	maxLenCiString255Type = 255
	// maxLenCiString500Type is the maximum length for a CiString500Type.
	// Used for fields like `diagnosticsStatus` (DiagnosticsStatusNotification.req - Table 10).
	maxLenCiString500Type = 500
)

// ciString is an internal, unexported struct that provides the core implementation
// for case-insensitive strings (CiString) as defined in the OCPP 1.6J specification.
// It encapsulates the raw string value and its maximum allowed length.
// This struct is not intended for direct use outside this package; instead, use
// the exported `CiStringXXType` types and their respective `Set` functions.
//
// See OCPP 1.6J Part 2, Appendix 3: "Data Types" for more information on
// case-insensitive strings and their constraints.
type ciString struct {
	raw    string // raw holds the actual string value. (Required)
	maxLen int    // maxLen specifies the maximum allowed length for this string. (Required)
}

// setCiString is an internal helper function used by the exported `SetCiStringXXType`
// functions to create and validate a `ciString` instance.
//
// It performs initial validation based on the provided maximum length.
//
// Parameters:
//   value: The string value to be encapsulated. (Required)
//   maxLen: The maximum allowed length for the string, as per OCPP 1.6J specification. (Required)
//
// Returns:
//   ciString: An initialized `ciString` struct.
//   error: An error if the `value` exceeds `maxLen` or contains non-printable ASCII characters.
func setCiString(value string, maxLen int) (ciString, error) {
	cs := ciString{raw: value, maxLen: maxLen}
	if err := cs.validate(); err != nil {
		return ciString{}, err
	}

	return cs, nil
}

// value returns the raw string value stored within the `ciString`.
//
// Returns:
//   string: The raw string content.
func (cs ciString) value() string {
	return cs.raw
}

// validate checks if the `ciString` instance conforms to the OCPP 1.6J constraints.
// These constraints include maximum length and the use of printable ASCII characters.
// An empty string is considered valid, representing an optional field that is not present.
//
// The validation performs the following checks:
// 1. Length Check: Ensures the string's length does not exceed `cs.maxLen`.
//    If it does, an `ErrExceedsMaxLength` error is returned.
// 2. Character Set Check: Verifies that all characters in the string are within
//    the printable ASCII range (ASCII values 32 to 126, inclusive).
//    If a non-printable character is found, an `ErrNonPrintableASCII` error is returned.
//
// Returns:
//   error: `nil` if the string is valid, otherwise an error detailing the validation failure.
func (cs ciString) validate() error {
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

// CiString20Type represents a case-insensitive string with a maximum length of 20 characters.
// This type is used for various fields in OCPP 1.6J messages that require a string
// of up to 20 characters, treated case-insensitively.
//
// Examples of usage include:
//   - `idTag` in Authorize.req (Table 12), StartTransaction.req (Table 36),
//     StopTransaction.req (Table 44), and DataTransfer.req (Table 9).
//   - `vendorId` in DataTransfer.req (Table 9).
//
// Conforms to the `CiString20Type` data type definition in OCPP 1.6J Part 2,
// Appendix 3: "Data Types".
type CiString20Type struct{
	value ciString // value holds the underlying ciString with max length 20. (Required)
}

// Value returns the raw string representation of the CiString20Type.
//
// Returns:
//   string: The string content of the CiString20Type.
func (c CiString20Type) Value() string {
	return c.value.value()
}

// Validate checks if the CiString20Type instance is valid according to OCPP 1.6J rules.
// This includes checking its length and character set.
//
// Returns:
//   error: `nil` if the string is valid, otherwise an error detailing the validation failure.
func (c CiString20Type) Validate() error {
	return c.value.validate()
}

// SetCiString20Type creates a new CiString20Type from a standard Go string.
// This is the recommended way to instantiate a CiString20Type.
//
// Parameters:
//   value: The string to be converted into a CiString20Type. (Required)
//
// Returns:
//   CiString20Type: A new CiString20Type instance.
//   error: An error if the input `value` is too long or contains invalid characters.
func SetCiString20Type(value string) (CiString20Type, error) {
	cs, err := setCiString(value, maxLenCiString20Type)

	return CiString20Type{value: cs}, err
}

// CiString25Type represents a case-insensitive string with a maximum length of 25 characters.
// This type is used for various fields in OCPP 1.6J messages that require a string
// of up to 25 characters, treated case-insensitively.
//
// Examples of usage include:
//   - `chargePointVendor` in BootNotification.req (Table 1).
//
// Conforms to the `CiString25Type` data type definition in OCPP 1.6J Part 2,
// Appendix 3: "Data Types".
type CiString25Type struct{
	value ciString // value holds the underlying ciString with max length 25. (Required)
}

// Value returns the raw string representation of the CiString25Type.
//
// Returns:
//   string: The string content of the CiString25Type.
func (c CiString25Type) Value() string {
	return c.value.value()
}

// Validate checks if the CiString25Type instance is valid according to OCPP 1.6J rules.
// This includes checking its length and character set.
//
// Returns:
//   error: `nil` if the string is valid, otherwise an error detailing the validation failure.
func (c CiString25Type) Validate() error {
	return c.value.validate()
}

// SetCiString25Type creates a new CiString25Type from a standard Go string.
// This is the recommended way to instantiate a CiString25Type.
//
// Parameters:
//   value: The string to be converted into a CiString25Type. (Required)
//
// Returns:
//   CiString25Type: A new CiString25Type instance.
//   error: An error if the input `value` is too long or contains invalid characters.
func SetCiString25Type(value string) (CiString25Type, error) {
	cs, err := setCiString(value, maxLenCiString25Type)

	return CiString25Type{value: cs}, err
}

// CiString50Type represents a case-insensitive string with a maximum length of 50 characters.
// This type is used for various fields in OCPP 1.6J messages that require a string
// of up to 50 characters, treated case-insensitively.
//
// Examples of usage include:
//   - `chargePointModel` in BootNotification.req (Table 1).
//   - `vendorErrorCode` in DataTransfer.conf (Table 11).
//
// Conforms to the `CiString50Type` data type definition in OCPP 1.6J Part 2,
// Appendix 3: "Data Types".
type CiString50Type struct{
	value ciString // value holds the underlying ciString with max length 50. (Required)
}

// Value returns the raw string representation of the CiString50Type.
//
// Returns:
//   string: The string content of the CiString50Type.
func (c CiString50Type) Value() string {
	return c.value.value()
}

// Validate checks if the CiString50Type instance is valid according to OCPP 1.6J rules.
// This includes checking its length and character set.
//
// Returns:
//   error: `nil` if the string is valid, otherwise an error detailing the validation failure.
func (c CiString50Type) Validate() error {
	return c.value.validate()
}

// SetCiString50Type creates a new CiString50Type from a standard Go string.
// This is the recommended way to instantiate a CiString50Type.
//
// Parameters:
//   value: The string to be converted into a CiString50Type. (Required)
//
// Returns:
//   CiString50Type: A new CiString50Type instance.
//   error: An error if the input `value` is too long or contains invalid characters.
func SetCiString50Type(value string) (CiString50Type, error) {
	cs, err := setCiString(value, maxLenCiString50Type)

	return CiString50Type{value: cs}, err
}

// CiString255Type represents a case-insensitive string with a maximum length of 255 characters.
// This type is used for various fields in OCPP 1.6J messages that require a string
// of up to 255 characters, treated case-insensitively.
//
// Examples of usage include:
//   - `firmwareVersion` in BootNotification.req (Table 1).
//   - `model` in GetCompositeSchedule.conf (Table 20).
//
// Conforms to the `CiString255Type` data type definition in OCPP 1.6J Part 2,
// Appendix 3: "Data Types".
type CiString255Type struct{
	value ciString // value holds the underlying ciString with max length 255. (Required)
}

// Value returns the raw string representation of the CiString255Type.
//
// Returns:
//   string: The string content of the CiString255Type.
func (c CiString255Type) Value() string {
	return c.value.value()
}

// Validate checks if the CiString255Type instance is valid according to OCPP 1.6J rules.
// This includes checking its length and character set.
//
// Returns:
//   error: `nil` if the string is valid, otherwise an error detailing the validation failure.
func (c CiString255Type) Validate() error {
	return c.value.validate()
}

// SetCiString255Type creates a new CiString255Type from a standard Go string.
// This is the recommended way to instantiate a CiString255Type.
//
// Parameters:
//   value: The string to be converted into a CiString255Type. (Required)
//
// Returns:
//   CiString255Type: A new CiString255Type instance.
//   error: An error if the input `value` is too long or contains invalid characters.
func SetCiString255Type(value string) (CiString255Type, error) {
	cs, err := setCiString(value, maxLenCiString255Type)

	return CiString255Type{value: cs}, err
}

// CiString500Type represents a case-insensitive string with a maximum length of 500 characters.
// This type is used for various fields in OCPP 1.6J messages that require a string
// of up to 500 characters, treated case-insensitively.
//
// Examples of usage include:
//   - `diagnosticsStatus` in DiagnosticsStatusNotification.req (Table 10).
//   - `status` in GetDiagnostics.conf (Table 19).
//
// Conforms to the `CiString500Type` data type definition in OCPP 1.6J Part 2,
// Appendix 3: "Data Types".
type CiString500Type struct{
	value ciString // value holds the underlying ciString with max length 500. (Required)
}

// Value returns the raw string representation of the CiString500Type.
//
// Returns:
//   string: The string content of the CiString500Type.
func (c CiString500Type) Value() string {
	return c.value.value()
}

// Validate checks if the CiString500Type instance is valid according to OCPP 1.6J rules.
// This includes checking its length and character set.
//
// Returns:
//   error: `nil` if the string is valid, otherwise an error detailing the validation failure.
func (c CiString500Type) Validate() error {
	return c.value.validate()
}

// SetCiString500Type creates a new CiString500Type from a standard Go string.
// This is the recommended way to instantiate a CiString500Type.
//
// Parameters:
//   value: The string to be converted into a CiString500Type. (Required)
//
// Returns:
//   CiString500Type: A new CiString500Type instance.
//   error: An error if the input `value` is too long or contains invalid characters.
func SetCiString500Type(value string) (CiString500Type, error) {
	cs, err := setCiString(value, maxLenCiString500Type)

	return CiString500Type{value: cs}, err
}
