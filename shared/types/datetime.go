package sharedtypes

import (
	"fmt"
	"time"
)

type DateTime struct {
	value time.Time
}

func SetDateTime(input string) (DateTime, error) {
	time, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return DateTime{}, fmt.Errorf("invalid datetime: %w", err)
	}

	time = time.UTC()

	return DateTime{value: time}, nil
}

func (dt DateTime) Value() time.Time {
	return dt.value
}

func (dt DateTime) String() string {
	return dt.value.Format(time.RFC3339Nano)
}
