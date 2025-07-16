// Package sharedtypes provides common types and utilities for the OCPP 1.6J protocol.
package sharedtypes

import (
	"fmt"
	"time"
)

// DateTime represents the OCPP 1.6J DateTime data type.
//
// This type is used for fields that require a date and time value, formatted
// according to ISO 8601 (e.g., "2013-02-01T20:00:00.000Z").
// The underlying Go type is `time.Time`.
//
// See OCPP 1.6J Part 2, Appendix 3: "Data Types" for the formal definition.
type DateTime struct {
	value time.Time // value holds the parsed time.Time object. (Required)
}

// SetDateTime creates a new DateTime type from an ISO 8601 formatted string.
//
// This function parses the input string using `time.RFC3339` layout.
// It returns an error if the string cannot be parsed into a valid `time.Time`.
//
// This is a constructor function for the DateTime type.
//
// Parameters:
//
//	input: The string representation of the datetime, e.g., "2013-02-01T20:00:00.000Z". (Required)
//
// Returns:
//
//	DateTime: A new DateTime object containing the parsed time.
//	error: An error if the input string is not a valid ISO 8601 datetime.
func SetDateTime(input string) (DateTime, error) {
	t, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return DateTime{}, fmt.Errorf("invalid datetime: %w", err)
	}

	return DateTime{value: t}, nil
}

// Value returns the underlying `time.Time` value of the DateTime type.
//
// This is an accessor method for the DateTime type.
//
// Returns:
//
//	time.Time: The parsed `time.Time` object.
func (dt DateTime) Value() time.Time {
	return dt.value
}

// String returns the ISO 8601 formatted string representation of the DateTime.
//
// This method formats the underlying `time.Time` value using `time.RFC3339` layout.
//
// Returns:
//
//	string: The ISO 8601 string representation of the datetime.
func (dt DateTime) String() string {
	return dt.value.Format(time.RFC3339)
}
