package sharedtypes

import (
	"fmt"
	"strconv"
)

type Integer struct {
	value uint16
}

func SetInteger(value string) (Integer, error) {
	parsedValue, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return Integer{}, fmt.Errorf("invalid Integer: %w", err)
	}

	return Integer{value: uint16(parsedValue)}, nil
}

func (integer Integer) Value() uint16 {
	return integer.value
}
