package types

import (
	"fmt"
	"time"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*DateTime)(nil)

const (
	utcOffsetSeconds = 0
	dateTimeField    = "value"
)

// DateTime represents an OCPP 1.6–compliant RFC3339 timestamp in UTC.
// Inputs must already be expressed in UTC; non-UTC offsets are rejected.
type DateTime struct {
	value time.Time
}

// NewDateTime creates a new DateTime from an RFC3339 formatted string.
// The input must be UTC (offset 0). Returns an error if parsing fails or
// the input is not UTC.
func NewDateTime(input string) (DateTime, error) {
	if input == "" {
		return DateTime{}, fmt.Errorf(
			"NewDateTime: "+ErrorFieldFormat,
			dateTimeField,
			ErrEmptyValue,
		)
	}

	timestamp, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return DateTime{}, fmt.Errorf(
			"NewDateTime: "+ErrorFieldFormat,
			dateTimeField,
			fmt.Errorf("%w: %w", ErrInvalidValue, err),
		)
	}

	_, offset := timestamp.Zone()
	if offset != utcOffsetSeconds {
		return DateTime{}, fmt.Errorf(
			"NewDateTime: "+ErrorFieldFormat,
			dateTimeField,
			fmt.Errorf(
				"%w: expected UTC offset, got %s",
				ErrInvalidValue,
				timestamp.Format("Z07:00"),
			),
		)
	}

	return DateTime{value: timestamp.UTC()}, nil
}

// Value returns the underlying time.Time value of the DateTime.
func (dt DateTime) Value() time.Time {
	return dt.value
}

// String implements the fmt.Stringer interface, returning the RFC3339Nano
// formatted string representation of the DateTime in UTC.
func (dt DateTime) String() string {
	return dt.value.Format(time.RFC3339Nano)
}
