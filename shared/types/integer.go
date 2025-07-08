// Package sharedtypes provides shared data types and utilities used throughout
// the OCPP 1.6J protocol implementation.
//
// This file defines the Integer type, a wrapper around Go's native integer types,
// intended to represent numeric fields used in various OCPP 1.6J messages
// (e.g., transactionId, meterValue, etc.).
//
// While the OCPP 1.6J specification references an 'integer' type, it does not
// explicitly define the expected size or signedness of these integers.
// This implementation assumes that a 32-bit unsigned integer (uint32) is
// sufficient for the protocol's current use cases.
//
// Wrapping the integer type in this abstraction provides future flexibility.
// If a different integer size or encoding is required (e.g., int64), the
// internal representation can be modified without impacting the external API.
//
// See OCPP 1.6J specification, Appendix 3: "Data Types", for detailed definition
// and constraints of the Integer type.
package sharedtypes

import (
	"fmt"
	"strconv"
)

// Integer represents the OCPP 1.6J Integer data type.
//
// This type is used for fields that require an unsigned 32-bit integer value,
// such as `transactionId` or `meterValue` in various OCPP messages.
//
// In the OCPP 1.6J specification, Integer values are typically constrained
// to be non-negative. The underlying Go type `uint32` enforces this.
//
// See OCPP 1.6J Part 2, Appendix 3: "Data Types" for the formal definition.
type Integer struct {
	value uint32 // value holds the unsigned 32-bit integer. (Required)
}

// SetInteger creates a new Integer type from a string representation.
//
// This function parses the input string as an unsigned 32-bit integer.
// It returns an error if the string cannot be parsed into a valid uint32.
//
// This is a constructor function for the Integer type.
//
// Parameters:
//   value: The string representation of the integer. (Required)
//
// Returns:
//   Integer: A new Integer object containing the parsed value.
//   error: An error if the input string is not a valid unsigned 32-bit integer.
func SetInteger(value string) (Integer, error) {
	parsedValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return Integer{}, fmt.Errorf("invalid Integer: %w", err)
	}

	return Integer{value: uint32(parsedValue)}, nil
}

// Value returns the underlying uint32 value of the Integer type.
//
// This is an accessor method for the Integer type.
//
// Returns:
//   uint32: The unsigned 32-bit integer value.
func (i Integer) Value() uint32 {
	return i.value
}
