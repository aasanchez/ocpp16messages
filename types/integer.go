package types

import (
	"fmt"
	"math"
	"strconv"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*Integer)(nil)

const (
	// decimalBase is the base-10 radix used for parsing integer strings.
	decimalBase = 10

	// bitSize16 is the bit size for uint16 parsing validation.
	bitSize16 = 16

	// integerMin is the minimum allowed value for Integer.
	integerMin = 0
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
		return Integer{}, fmt.Errorf(
			"NewInteger: "+ErrorFieldFormat,
			"value",
			fmt.Errorf("%w: %w", ErrInvalidValue, err),
		)
	}

	return NewIntegerFromInt(int(parsedValue))
}

// NewIntegerFromInt creates a new Integer from an int, validating range.
func NewIntegerFromInt(value int) (Integer, error) {
	if value < integerMin || value > math.MaxUint16 {
		return Integer{}, fmt.Errorf(
			"NewIntegerFromInt: "+ErrorFieldFormat,
			"value",
			fmt.Errorf("%w: %d out of range (0-65535)", ErrInvalidValue, value),
		)
	}

	return Integer{value: uint16(value)}, nil
}

// NewIntegerFromUint16 creates a new Integer from a uint16 without error.
func NewIntegerFromUint16(value uint16) Integer {
	return Integer{value: value}
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
