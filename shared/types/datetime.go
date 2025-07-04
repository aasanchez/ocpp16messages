package types

import (
	"errors"
	"fmt"
	"time"
)

type DateTime struct {
	value time.Time
}

var ErrNotFound = errors.New("not found")

func SetDateTime(input string) (*DateTime, error) {
	if input == "" {
		return nil, ErrNotFound
	}

	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return nil, fmt.Errorf("invalid datetime: %w", err)
	}

	return &DateTime{value: t}, nil
}

func (dt *DateTime) Value() time.Time {
	return dt.value
}

func (dt *DateTime) String() string {
	return dt.value.Format(time.RFC3339)
}
