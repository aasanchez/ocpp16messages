package types

import (
	"fmt"
	"strconv"
)

type Integer struct {
	value uint32
}

func SetInteger(value string) (Integer, error) {
	parsedValue, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return Integer{}, fmt.Errorf("invalid integer: %w", err)
	}

	return Integer{value: uint32(parsedValue)}, nil
}

// Value returns the underlying uint32 value of the Integer.
func (i Integer) Value() uint32 {
	return i.value
}
