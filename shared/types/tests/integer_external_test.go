// Package types_test contains blackbox (external) tests for the `Integer` type
// defined in the `shared/types` package.
//
// This test file exists to verify the end-to-end functionality and adherence to OCPP 1.6J
// specifications for the `Integer` data type. The `Integer` type is designed to represent
// unsigned 32-bit integer values, which are commonly used in various OCPP messages
// (e.g., `transactionId`, `meterValue`).
//
// The tests cover the following real-world use cases and boundary conditions:
//   - **Valid Integer Parsing**: Confirms that `SetInteger` can successfully parse
//     numeric strings that represent valid unsigned 32-bit integers, including
//     zero and the maximum `uint32` value.
//   - **Invalid Input Handling**: Ensures that `SetInteger` correctly rejects
//     strings that are not valid unsigned 32-bit integers, such as negative numbers,
//     non-numeric strings, empty strings, or values exceeding the `uint32` range.
//   - **Value Retrieval**: Validates that the `Value()` method accurately returns
//     the underlying `uint32` value, allowing Go applications to work with the
//     parsed integer.
//
// This blackbox testing approach ensures that the `Integer` type behaves correctly
// from a consumer's perspective and complies with the OCPP 1.6J specification,
// particularly concerning integer data as described in OCPP 1.6J Part 2, Appendix 3:
// "Data Types".
package types_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// TestSetInteger_ValidInteger verifies that `SetInteger` successfully parses
// a standard valid numeric string into an `Integer` type without returning an error.
// This simulates a typical valid integer value received in an OCPP message.
func TestSetInteger_ValidInteger(t *testing.T) {
	t.Parallel()

	// Scenario: Provide a simple, positive numeric string.
	// Expectation: No error should be returned.
	if _, err := st.SetInteger("123"); err != nil {
		t.Fatalf("expected no error for valid integer, got %v", err)
	}
}

// TestSetInteger_Zero verifies that `SetInteger` correctly accepts "0" as a valid
// unsigned 32-bit integer. Zero is a valid value for OCPP Integer types.
func TestSetInteger_Zero(t *testing.T) {
	t.Parallel()

	// Scenario: Provide the string "0".
	// Expectation: No error should be returned.
	if _, err := st.SetInteger("0"); err != nil {
		t.Fatalf("expected no error for zero, got %v", err)
	}
}

// TestSetInteger_MaxUint32 verifies that `SetInteger` successfully parses the string
// representation of the maximum `uint32` value without error.
// This tests a boundary condition for the `Integer` type, ensuring it can handle
// the largest possible valid unsigned integer as per Go's `uint32` and OCPP's `Integer`.
func TestSetInteger_MaxUint32(t *testing.T) {
	t.Parallel()

	// Scenario: Provide the string representation of `2^32 - 1`.
	// Expectation: No error should be returned.
	if _, err := st.SetInteger("4294967295"); err != nil {
		t.Fatalf("expected no error for max uint32, got %v", err)
	}
}

// TestSetInteger_NegativeInteger verifies that `SetInteger` rejects strings
// representing negative integers, as OCPP `Integer` types are unsigned.
func TestSetInteger_NegativeInteger(t *testing.T) {
	t.Parallel()

	// Scenario: Provide a negative numeric string.
	// Expectation: An error should be returned, as negative values are invalid for `uint32`.
	if _, err := st.SetInteger("-1"); err == nil {
		t.Fatal("expected error for negative integer, got none")
	}
}

// TestSetInteger_NotANumber verifies that `SetInteger` rejects non-numeric strings.
// This ensures robust parsing and adherence to the numeric nature of OCPP `Integer`.
func TestSetInteger_NotANumber(t *testing.T) {
	t.Parallel()

	// Scenario: Provide a string that is not a number.
	// Expectation: An error should be returned.
	if _, err := st.SetInteger("abc"); err == nil {
		t.Fatal("expected error for non-numeric input, got none")
	}
}

// TestSetInteger_EmptyString verifies that `SetInteger` rejects an empty string.
// While some OCPP fields can be optional, an empty string is not a valid representation
// for an `Integer` value itself.
func TestSetInteger_EmptyString(t *testing.T) {
	t.Parallel()

	// Scenario: Provide an empty string.
	// Expectation: An error should be returned.
	if _, err := st.SetInteger(""); err == nil {
		t.Fatal("expected error for empty string, got none")
	}
}

// TestSetInteger_TooLarge verifies that `SetInteger` rejects numeric strings
// that represent values exceeding the maximum range of a `uint32`.
// This tests the upper boundary condition for the `Integer` type.
func TestSetInteger_TooLarge(t *testing.T) {
	t.Parallel()

	// Scenario: Provide a numeric string larger than `2^32 - 1`.
	// Expectation: An error should be returned.
	if _, err := st.SetInteger("4294967296"); err == nil {
		t.Fatal("expected error for value exceeding uint32 range, got none")
	}
}

// TestInteger_Value verifies that the `Value()` accessor method of the `Integer` type
// correctly returns the underlying `uint32` value that was set.
func TestInteger_Value(t *testing.T) {
	t.Parallel()

	// Scenario: Create an `Integer` from a valid string and then retrieve its value.
	// Expectation: The retrieved `uint32` should match the original numeric input.
	i, err := st.SetInteger("7")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := i.Value(); got != 7 {
		t.Errorf("Value() = %d; want %d", got, 7)
	}
}
