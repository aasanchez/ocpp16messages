package sharedtypes_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// Example: create and read an Integer.
//
// This example shows the typical flow:
//  1. parse a base-10 string with SetInteger
//  2. retrieve the stored number with Value
//
// Notes
//   - accepted range: 0..65535 (uint16)
//   - zero value is valid (0)
//   - errors wrap the parse failure
//   - no panics are raised
//   - immutable after creation
//
// Tips
//   - always check the returned error
//   - prefer explicit error messages
//   - keep inputs short and numeric
//   - use Value() for serialization
func ExampleSetInteger() {
	intVal, err := st.SetInteger("73")
	if err != nil {
		fmt.Printf(st.ErrorUnexpectedError, err)

		return
	}

	fmt.Println("Value is:", intVal.Value())

	// Output:
	// Value is: 73
}

// Example: invalid input handling.
//
// Demonstrates how SetInteger reports
// errors for non-decimal input. The
// returned error wraps strconv.ParseUint
// to preserve the root cause.
//
// Guidance
//   - show a short, friendly message
//   - print the wrapped error below
//   - do not ignore the error path
func ExampleSetInteger_invalid() {
	_, err := st.SetInteger("abc")
	if err != nil {
		fmt.Println("Failed to parse int from 'abc'")
		fmt.Print(err)
	}

	// Output:
	// Failed to parse int from 'abc'
	// invalid Integer: strconv.ParseUint: parsing "abc": invalid syntax
}
