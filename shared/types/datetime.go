package types

import (
	"fmt"
	"time"
)

// DateTime represents an OCPP 1.6–compliant RFC3339 timestamp in UTC.
type DateTime struct {
	value time.Time
}

// NewDateTime creates a new DateTime from an RFC3339 formatted string. The
// resulting DateTime is automatically normalized to UTC. Returns an error if
// the input string is not a valid RFC3339 timestamp.
func NewDateTime(input string) (DateTime, error) {
	timestamp, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return DateTime{}, fmt.Errorf("invalid datetime: %w", err)
	}

	timestamp = timestamp.UTC()

	return DateTime{value: timestamp}, nil
}

// Value returns the underlying time.Time value of the DateTime.
func (dt DateTime) Value() time.Time {
	return dt.value
}

// String returns the DateTime formatted as an RFC3339Nano string.
func (dt DateTime) String() string {
	return dt.value.Format(time.RFC3339Nano)
}
