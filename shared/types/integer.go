package sharedtypes

import (
	"fmt"
	"strconv"
)

// Integer is an immutable unsigned 16-bit value.
//
// Zero value is valid and equals 0.
type Integer struct {
	value uint16
}

// SetInteger parses a base-10 string into an Integer.
//
// Accepts 0..65535. Returns an error for empty,
// non-decimal, negative, or overflow input.
// The error wraps the parse failure for context.
func SetInteger(value string) (Integer, error) {
	parsedValue, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return Integer{}, fmt.Errorf("invalid Integer: %w", err)
	}

	return Integer{value: uint16(parsedValue)}, nil
}

// Value returns the stored uint16.
func (integer Integer) Value() uint16 {
	return integer.value
}
