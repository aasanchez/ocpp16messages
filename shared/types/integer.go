package sharedtypes

import (
	"fmt"
	"strconv"
)

const (
	decimalBase = 10
	bitSize16   = 16
)

// Integer represents an OCPP 1.6 compliant integer value (uint16).
type Integer struct {
	value uint16
}

// SetInteger creates a new Integer from a string value.
// Returns an error if the value cannot be parsed as a valid uint16.
func SetInteger(value string) (Integer, error) {
	parsedValue, err := strconv.ParseUint(value, decimalBase, bitSize16)
	if err != nil {
		return Integer{}, fmt.Errorf("invalid Integer: %w", err)
	}

	return Integer{value: uint16(parsedValue)}, nil
}

// Value returns the underlying uint16 value of the Integer.
func (integer Integer) Value() uint16 {
	return integer.value
}
