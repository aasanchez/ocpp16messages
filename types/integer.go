package types

import (
	"fmt"
	"strconv"
)

const (
	// decimalBase is the base-10 radix used for parsing integer strings.
	decimalBase = 10

	// bitSize16 is the bit size for uint16 parsing validation.
	bitSize16 = 16
)

// Integer represents an OCPP 1.6 compliant integer value (uint16).
type Integer struct {
	value uint16
}

// NewInteger creates a new Integer from a string value. Returns an error if
// the value cannot be parsed as a valid uint16.
func NewInteger(value string) (Integer, error) {
	parsedValue, err := strconv.ParseUint(value, decimalBase, bitSize16)
	if err != nil {
		return Integer{}, fmt.Errorf("%w", err)
	}

	return Integer{value: uint16(parsedValue)}, nil
}

// Value returns the underlying uint16 value of the Integer.
func (integer Integer) Value() uint16 {
	return integer.value
}
