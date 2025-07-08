package sharedtypes

import (
	"fmt"
	"strconv"
)

// Integer is a custom type that represents an unsigned 32-bit integer.
// It is used to provide type safety and to encapsulate parsing logic for integer values
// that are part of OCPP 1.6 messages.
type Integer struct {
	value uint32 // The underlying unsigned 32-bit integer value.
}

// SetInteger creates a new Integer type from a string representation.
// It parses the input string as an unsigned 32-bit integer.
//
// Parameters:
//   value string - The string to be parsed into an Integer.
//
// Returns:
//   Integer - The created Integer type.
//   error   - An error if the string cannot be parsed into a valid unsigned 32-bit integer.
//
// This function is used to convert string-based integer values, often received from
// external sources (like JSON payloads), into a type-safe Integer for internal use.
// It ensures that the integer conforms to the expected format and range.
// In OCPP 1.6, integer values are typically represented as whole numbers.
// For example, transaction IDs, meter values, or durations are often integers.
// Refer to the OCPP 1.6 specification for specific message field definitions.
func SetInteger(value string) (Integer, error) {
	parsedValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return Integer{}, fmt.Errorf("invalid integer: %w", err)
	}

	return Integer{value: uint32(parsedValue)}, nil
}

// Value returns the underlying uint32 value of the Integer type.
//
// Returns:
//   uint32 - The unsigned 32-bit integer value.
//
// This method provides access to the raw integer value for calculations or
// when converting the Integer type back to a primitive type.
func (i Integer) Value() uint32 {
	return i.value
}
