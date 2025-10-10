package sharedtypes

import (
	"fmt"
	"time"
)

// DateTime wraps a time.Time normalized to UTC
// for OCPP 1.6 timestamps.
//
// OCPP 1.6
//   - Timestamps are RFC3339 strings exchanged
//     between Charge Point (CP) and CSMS.
//   - Inputs may include an offset (e.g. +02:00) or Z.
//   - Internally we normalize to UTC for consistency.
//
// Behavior
//   - Immutable after construction.
//   - Concurrency-safe (read-only state).
//   - No panics; errors are returned.
//
// Interop
//   - Accepts RFC3339 (time.RFC3339) input.
//   - Value() returns a UTC time.Time.
//   - String() renders RFC3339Nano in UTC, which is
//     stable for logs, tests, and message signing.
//
// Errors
//   - SetDateTime wraps the parse error using
//     fmt.Errorf("invalid datetime: %w", err) so callers
//     can inspect the root cause with errors.As.
//
// Usage
//
//	dt, err := SetDateTime("2025-08-30T14:34:56+02:00")
//	if err != nil { /* handle */ }
//	_ = dt.Value()  // time.Time (UTC)
//	_ = dt.String() // "2025-08-30T12:34:56Z"
type DateTime struct {
	value time.Time
}

// SetDateTime parses an RFC3339 string and returns
// a UTC-normalized DateTime.
func SetDateTime(input string) (DateTime, error) {
	dateTime, err := time.Parse(time.RFC3339, input)
	if err != nil {
		return DateTime{}, fmt.Errorf("invalid datetime: %w", err)
	}

	dateTime = dateTime.UTC()

	return DateTime{value: dateTime}, nil
}

// Value returns the underlying UTC time.Time.
func (dt DateTime) Value() time.Time {
	return dt.value
}

// String returns the RFC3339Nano representation
// of the stored UTC time.
func (dt DateTime) String() string {
	return dt.value.Format(time.RFC3339Nano)
}
