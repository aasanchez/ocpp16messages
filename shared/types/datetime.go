package types

import (
	"errors"
	"fmt"
	"time"
)

type DateTimeType struct {
	value time.Time
}

var ErrZeroDateTime = errors.New("DateTime is zero and invalid")

func DateTime(input string) (DateTimeType, error) {
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return DateTimeType{}, fmt.Errorf("%w", err)
	}
	return DateTimeType{value: t}, nil
}

func (dt DateTimeType) Validate() error {
	if dt.value.IsZero() {
		return fmt.Errorf("%w", ErrZeroDateTime)
	}
	return nil
}

func (dt DateTimeType) Value() time.Time {
	return dt.value
}
