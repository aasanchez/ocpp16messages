// Package sharedtypes_test contains unit tests for the shared data types
// defined in the `sharedtypes` package, specifically focusing on the `DateTime` type.
//
// This test file exists to ensure that the `DateTime` type, which represents
// date and time values in OCPP 1.6J messages, correctly handles parsing from
// and formatting to the ISO 8601 standard (RFC3339).
//
// The tests verify the following real-world use cases and edge cases:
//   - **Valid Parsing**: Confirms that `SetDateTime` can successfully parse
//     well-formed ISO 8601 strings, as expected for incoming OCPP messages.
//   - **Invalid Format Handling**: Ensures that `SetDateTime` correctly identifies
//     and rejects malformed date/time strings, preventing invalid data from
//     being processed.
//   - **Value Retrieval**: Validates that the `Value()` method accurately returns
//     the underlying `time.Time` object, allowing Go applications to work with
//     the parsed date/time.
//   - **String Representation**: Verifies that the `String()` method correctly
//     formats the `DateTime` object back into an ISO 8601 string, ensuring
//     proper serialization for outgoing OCPP messages.
//
// This testing strategy helps guarantee the robustness and compliance of the
// `DateTime` type with the OCPP 1.6J specification, particularly concerning
// date and time data as described in OCPP 1.6J Part 2, Appendix 3: "Data Types".
package sharedtypes

import (
	"testing"
)

// TestDateTime_parsesValidRFC3339 verifies that the `SetDateTime` function
// successfully parses a valid RFC3339 (ISO 8601) formatted string into a `DateTime` object.
// This ensures compliance with the OCPP 1.6J specification for DateTime data types.
func TestDateTime_parsesValidRFC3339(t *testing.T) {
	t.Parallel()

	// Scenario: Provide a valid ISO 8601 formatted date-time string.
	// This simulates a typical date-time value received in an OCPP message.
	input := "2026-04-12T10:03:04-04:00"

	// Expectation: The function should parse the string without returning an error.
	_, err := SetDateTime(input)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

// TestDateTime_failsOnInvalidFormat verifies that the `SetDateTime` function
// returns an error when provided with a string that does not conform to the
// RFC3339 (ISO 8601) format.
// This tests the error handling for malformed date-time inputs, ensuring data integrity.
func TestDateTime_failsOnInvalidFormat(t *testing.T) {
	t.Parallel()

	// Scenario: Provide a string that is clearly not a valid date-time format.
	// This simulates receiving corrupted or incorrectly formatted date-time data.
	input := "not-a-date"
	_, err := SetDateTime(input)

	// Expectation: An error should be returned, indicating parsing failure.
	if err == nil {
		t.Error("expected error for invalid RFC3339 string, got nil")
	}
}

// TestDateTime_valueReturnsCorrectTime verifies that the `Value()` method of the
// `DateTime` object correctly returns the underlying `time.Time` representation.
// This ensures that the parsed date-time can be accessed and used by Go's standard time utilities.
func TestDateTime_valueReturnsCorrectTime(t *testing.T) {
	t.Parallel()

	// Scenario: Parse a valid date-time string and then retrieve its `time.Time` value.
	// This confirms that the internal storage and retrieval mechanism is working as expected.
	input := "2027-04-12T09:03:04-04:00"
	dt, _ := SetDateTime(input)

	// Expectation: The returned `time.Time` object should not be its zero value,
	// indicating successful internal representation.
	if dt.Value().IsZero() {
		t.Error("expected non-zero time from Value(), got zero")
	}
}

// TestDateTime_String_returnsRFC3339 verifies that the `String()` method of the
// `DateTime` object correctly formats the date-time back into an RFC3339 (ISO 8601) string.
// This is crucial for serializing `DateTime` objects into outgoing OCPP messages.
func TestDateTime_String_returnsRFC3339(t *testing.T) {
	t.Parallel()

	// Scenario: Create a `DateTime` object from a known ISO 8601 string and then
	// convert it back to a string using the `String()` method.
	input := "2025-12-25T15:00:00Z"

	datetime, err := SetDateTime(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Expectation: The formatted string should exactly match the original input string.
	got := datetime.String()
	if got != input {
		t.Errorf("expected %q, got %q", input, got)
	}
}
