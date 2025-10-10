package sharedtypes_test

import (
	"fmt"
	"time"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// Example: OCPP 1.6 timestamp parsing.
//
// Context
//   - OCPP 1.6 uses RFC3339 for timestamps.
//   - SetDateTime validates and normalizes to UTC.
//
// What this shows
//  1. Parse a timestamp string.
//  2. Read the time.Time value (UTC).
//  3. Render the canonical string form.
func ExampleSetDateTime() {
	datetime, err := st.SetDateTime("2025-08-30T14:34:56Z")
	if err != nil {
		fmt.Printf(st.ErrorUnexpectedError, err)

		return
	}

	fmt.Println("Time value:", datetime.Value().UTC())
	fmt.Println("String:", datetime.String())

	// Output:
	// Time value: 2025-08-30 14:34:56 +0000 UTC
	// String: 2025-08-30T14:34:56Z
}

// Example: view in a local timezone.
//
// Context
//   - Charge points send UTC; operators may view
//     in local time. The internal value is UTC,
//     but can be converted for display.
//   - String() remains canonical (UTC RFC3339).
func ExampleSetDateTime_changeTimezone() {
	datetime, err := st.SetDateTime("2025-08-30T14:34:56Z")
	if err != nil {
		fmt.Printf(st.ErrorUnexpectedError, err)

		return
	}

	loc, _ := time.LoadLocation("Europe/Madrid")

	fmt.Println("Time value:", datetime.Value().In(loc))
	fmt.Println("String:", datetime.String())

	// Output:
	// Time value: 2025-08-30 16:34:56 +0200 CEST
	// String: 2025-08-30T14:34:56Z
}

// Example: invalid timestamp handling.
//
// Context
//   - Inputs can be malformed. SetDateTime returns
//     a wrapped error that preserves the parse failure.
//
// Guidance
//   - Show a short, friendly message if needed.
//   - Log the wrapped error to keep root cause.
func ExampleSetDateTime_invalid() {
	_, err := st.SetDateTime("not-a-time")
	if err != nil {
		fmt.Printf("%.70v...\n", err)
	}

	// Output:
	// invalid datetime: parsing time "not-a-time" as "2006-01-02T15:04:05Z07...
}
