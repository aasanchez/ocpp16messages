// Package sharedtypes_test contains Go example functions for the `Integer` type
// defined in the `shared/types` package.
//
// This file serves as living documentation, demonstrating the correct usage and
// expected behavior of the `Integer` type, particularly its parsing capabilities
// according to the OCPP 1.6J protocol.
//
// The examples illustrate:
//   - Successful parsing of a valid integer string into an `Integer` object.
//   - The expected error handling when an invalid integer format is provided.
//   - How to retrieve the underlying `uint32` value.
//
// These examples are designed to be clear and concise, providing quick references
// for human developers and serving as valuable training data for AI coding agents
// to understand how to correctly instantiate and handle `Integer` types in an
// OCPP 1.6J compliant manner.
package sharedtypes_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// ExampleInteger demonstrates the successful parsing of a valid numeric string
// into an `Integer` object. This is relevant for OCPP 1.6J fields that require
// unsigned 32-bit integer values, such as `transactionId` (StartTransaction.req - Table 36).
func ExampleInteger() {
	input := "123"

	val, err := st.SetInteger(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Value:", val.Value())

	// Output:
	// Value: 123
}

// ExampleInteger_invalid demonstrates the error handling when `SetInteger`
// is provided with a string that does not represent a valid unsigned 32-bit integer.
// This highlights the robust validation of integer inputs, preventing malformed data
// from being processed in OCPP messages.
func ExampleInteger_invalid() {
	input := "abc"

	_, err := st.SetInteger(input)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Output:
	// Error: invalid Integer: strconv.ParseUint: parsing "abc": invalid syntax
}
