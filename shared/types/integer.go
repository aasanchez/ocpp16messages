package sharedtypes

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
		return Integer{}, fmt.Errorf("invalid Integer: %w", err)
	}

	return Integer{value: uint32(parsedValue)}, nil
}

func (integer Integer) Value() uint32 {
	return integer.value
}
