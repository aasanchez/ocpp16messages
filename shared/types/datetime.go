package types

import (
	"fmt"
	"time"
)

type DateTimeType struct {
	value time.Time
}

func DateTime(input string) (DateTimeType, error) {
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return DateTimeType{}, fmt.Errorf("invalid datetime: %w", err)
	}

	return DateTimeType{value: t}, nil
}

func (dt DateTimeType) Value() time.Time {
	return dt.value
}

func (dt DateTimeType) String() string {
	return dt.value.Format(time.RFC3339)
}
