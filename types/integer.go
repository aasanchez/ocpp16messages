package types

import (
	"fmt"
	"math"
	"strconv"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*Integer)(nil)

const (
	// decimalBase is the base-10 radix for string formatting.
	decimalBase = 10
	// integerMin is the minimum allowed value for Integer.
	integerMin = 0
)

// Integer represents an OCPP 1.6 compliant integer value (uint16).
type Integer struct {
	value uint16
}

// NewInteger creates a new Integer from an int, validating range.
func NewInteger(value int) (Integer, error) {
	if value < integerMin || value > math.MaxUint16 {
		return Integer{}, fmt.Errorf(
			"NewInteger: "+ErrorFieldFormat,
			"value",
			fmt.Errorf("%w: %d out of range (0-65535)", ErrInvalidValue, value),
		)
	}

	return Integer{value: uint16(value)}, nil
}

// Value returns the underlying uint16 value of the Integer.
func (integer Integer) Value() uint16 {
	return integer.value
}

// String implements the fmt.Stringer interface, returning the decimal string
// representation of the integer value.
func (integer Integer) String() string {
	return strconv.FormatUint(uint64(integer.value), decimalBase)
}
