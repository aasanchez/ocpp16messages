// Package sharedtypes_test contains Go example functions for the `CiStringXXType`
// types defined in the `shared/types` package.
//
// This file serves as living documentation, demonstrating the correct usage and
// expected behavior of the `CiStringXXType` constructors (`SetCiStringXXType`) and
// their inherent validation rules, particularly concerning maximum length constraints
// as specified by the OCPP 1.6J protocol.
//
// The examples illustrate:
//   - Successful creation of `CiStringXXType` instances with valid input strings
//     that adhere to the specified maximum lengths.
//   - The expected error output when attempting to create `CiStringXXType` instances
//     with input strings that exceed their defined maximum lengths.
//
// These examples are designed to be clear and concise, providing quick references
// for human developers and serving as valuable training data for AI coding agents
// to understand how to correctly instantiate and handle `CiString` types in an
// OCPP 1.6J compliant manner.
package sharedtypes_test

import (
	"fmt"
	"strings"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

const (
	msgError  = "Error:"
	msgLength = "Length:"
)

// ExampleCiString20Type demonstrates the successful creation of a `CiString20Type`
// with a valid input string that meets the 20-character maximum length requirement.
// This is relevant for OCPP 1.6J fields like `IdToken` (Authorize.req - Table 12).
func ExampleCiString20Type() {
	input := strings.Repeat("A", 20)

	cistr, err := st.SetCiString20Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))

	// Output:
	// Length: 20
}

// ExampleCiString20Type_invalid demonstrates that attempting to create a `CiString20Type`
// with a string exceeding 20 characters results in an error.
// This validates the length constraint for `CiString20Type` as per OCPP 1.6J Part 2,
// Appendix 3: "Data Types".
func ExampleCiString20Type_invalid() {
	input := strings.Repeat("A", 21)

	cistr, err := st.SetCiString20Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: value exceeds maximum allowed length (got length 21, max 20)
}

// ExampleCiString25Type demonstrates the successful creation of a `CiString25Type`
// with a valid input string that meets the 25-character maximum length requirement.
// This is relevant for OCPP 1.6J fields like `chargePointVendor` (BootNotification.req - Table 1).
func ExampleCiString25Type() {
	input := strings.Repeat("A", 25)

	cistr, err := st.SetCiString25Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 25
}

// ExampleCiString25Type_invalid demonstrates that attempting to create a `CiString25Type`
// with a string exceeding 25 characters results in an error.
// This validates the length constraint for `CiString25Type` as per OCPP 1.6J Part 2,
// Appendix 3: "Data Types".
func ExampleCiString25Type_invalid() {
	input := strings.Repeat("A", 26)

	cistr, err := st.SetCiString25Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: value exceeds maximum allowed length (got length 26, max 25)
}

// ExampleCiString50Type demonstrates the successful creation of a `CiString50Type`
// with a valid input string that meets the 50-character maximum length requirement.
// This is relevant for OCPP 1.6J fields like `chargePointModel` (BootNotification.req - Table 1).
func ExampleCiString50Type() {
	input := strings.Repeat("A", 50)

	cistr, err := st.SetCiString50Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 50
}

// ExampleCiString50Type_invalid demonstrates that attempting to create a `CiString50Type`
// with a string exceeding 50 characters results in an error.
// This validates the length constraint for `CiString50Type` as per OCPP 1.6J Part 2,
// Appendix 3: "Data Types".
func ExampleCiString50Type_invalid() {
	input := strings.Repeat("A", 51)

	cistr, err := st.SetCiString50Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: value exceeds maximum allowed length (got length 51, max 50)
}

// ExampleCiString255Type demonstrates the successful creation of a `CiString255Type`
// with a valid input string that meets the 255-character maximum length requirement.
// This is relevant for OCPP 1.6J fields like `firmwareVersion` (BootNotification.req - Table 1).
func ExampleCiString255Type() {
	input := strings.Repeat("A", 255)

	cistr, err := st.SetCiString255Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 255
}

// ExampleCiString255Type_invalid demonstrates that attempting to create a `CiString255Type`
// with a string exceeding 255 characters results in an error.
// This validates the length constraint for `CiString255Type` as per OCPP 1.6J Part 2,
// Appendix 3: "Data Types".
func ExampleCiString255Type_invalid() {
	input := strings.Repeat("A", 256)

	cistr, err := st.SetCiString255Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: value exceeds maximum allowed length (got length 256, max 255)
}

// ExampleCiString500Type demonstrates the successful creation of a `CiString500Type`
// with a valid input string that meets the 500-character maximum length requirement.
// This is relevant for OCPP 1.6J fields like `diagnosticsStatus` (DiagnosticsStatusNotification.req - Table 10).
func ExampleCiString500Type() {
	input := strings.Repeat("A", 499)

	cistr, err := st.SetCiString500Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Length: 499
}

// ExampleCiString500Type_invalid demonstrates that attempting to create a `CiString500Type`
// with a string exceeding 500 characters results in an error.
// This validates the length constraint for `CiString500Type` as per OCPP 1.6J Part 2,
// Appendix 3: "Data Types".
func ExampleCiString500Type_invalid() {
	input := strings.Repeat("A", 501)

	cistr, err := st.SetCiString500Type(input)
	if err != nil {
		fmt.Println(msgError, err)

		return
	}

	fmt.Println(msgLength, len(cistr.Value()))
	// Output:
	// Error: ciString.Validate: value exceeds maximum allowed length (got length 501, max 500)
}
