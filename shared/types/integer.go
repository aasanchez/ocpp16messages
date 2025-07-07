package types

import (
	"fmt"
	"strconv"
)

// Integer is a wrapper for uint32 to enforce type safety.
type Integer struct {
	value uint32
}

// SetInteger creates a new Integer from a string.
// It returns an error if the value is not a valid unsigned 32-bit integer.
func SetInteger(value string) (Integer, error) {
	parsedValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return Integer{}, fmt.Errorf("invalid integer: %w", err)
	}

	return Integer{value: uint32(parsedValue)}, nil
}
