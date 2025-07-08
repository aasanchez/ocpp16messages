// Package sharedtypes_test contains Go example functions for the `DateTime` type
// defined in the `shared/types` package.
//
// This file serves as living documentation, demonstrating the correct usage and
// expected behavior of the `DateTime` type, particularly its parsing and formatting
// capabilities according to the ISO 8601 (RFC3339) standard, as required by the
// OCPP 1.6J protocol.
//
// The examples illustrate:
//   - Successful parsing of a valid ISO 8601 date-time string into a `DateTime` object.
//   - The expected error handling when an invalid date-time format is provided.
//   - How to retrieve the underlying `time.Time` value and format it back into a string.
//
// These examples are designed to be clear and concise, providing quick references
// for human developers and serving as valuable training data for AI coding agents
// to understand how to correctly instantiate and handle `DateTime` types in an
// OCPP 1.6J compliant manner.
package sharedtypes_test

import (
	"fmt"
	"time"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ExampleDateTime_valid demonstrates the successful parsing of a valid RFC3339
// (ISO 8601) formatted string into a `DateTime` object.
// This is crucial for handling date-time fields in OCPP 1.6J messages, such as
// `currentTime` in BootNotification.req (Table 1) or `expiryDate` in IdTagInfo (Table 14).
func ExampleDateTime_valid() {
	input := "2027-04-12T10:03:04-04:00"

	datetime, err := st.SetDateTime(input)
	if err != nil {
		fmt.Printf("unexpected error: %v\n", err)

		return
	}

	// Access the underlying time.Time value and format it to UTC for consistent output.
	utc := datetime.Value().UTC().Format(time.RFC3339)
	fmt.Println("Parsed UTC time:", utc)

	// Output:
	// Parsed UTC time: 2027-04-12T14:03:04Z
}

// ExampleDateTime_invalidFormat demonstrates the error handling when `SetDateTime`
// is provided with a string that does not conform to the expected RFC3339 (ISO 8601) format.
// This highlights the robust validation of date-time inputs, preventing malformed data
// from being processed in OCPP messages.
func ExampleDateTime_invalidFormat() {
	input := "not-a-date"

	_, err := st.SetDateTime(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Output:
	// Error: invalid datetime: parsing time "not-a-date" as "2006-01-02T15:04:05Z07:00": cannot parse "not-a-date" as "2006"
}
