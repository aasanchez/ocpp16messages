// Package sharedtypes_test contains unit tests for the shared data types
// defined in the `sharedtypes` package, specifically focusing on the `ciString`
// internal type and its validation logic.
//
// This test file exists to ensure that the core validation rules for OCPP 1.6J
// case-insensitive strings (CiString) are correctly implemented and enforced.
// These rules include maximum length constraints and restrictions on character sets
// (only printable ASCII characters are allowed).
//
// The tests verify the following real-world use cases and edge cases:
//   - **Maximum Length Enforcement**: Validates that strings exceeding their defined
//     maximum length (e.g., `CiString20Type` being longer than 20 characters)
//     are correctly rejected, aligning with OCPP 1.6J Part 2, Appendix 3: "Data Types".
//   - **Valid String Acceptance**: Confirms that well-formed strings within length
//     limits and containing only printable ASCII characters are accepted.
//   - **Empty String Handling**: Ensures that empty strings, which are often used
//     to represent optional fields in OCPP, are considered valid.
//   - **Non-Printable Character Rejection**: Verifies that strings containing
//     non-printable ASCII characters (e.g., control characters) are correctly
//     rejected, enforcing the character set constraint of OCPP CiStrings.
//   - **Value Retrieval**: Tests that the `value()` method of the internal `ciString`
//     correctly returns the original raw string.
//
// This testing strategy helps guarantee the robustness and compliance of the
// `sharedtypes` package with the OCPP 1.6J specification, providing a solid
// foundation for all messages that utilize these common string types.
package sharedtypes

import (
	"errors"
	"strings"
	"testing"
)

// TestCiString_TooLongFails verifies that `setCiString` returns an `ErrExceedsMaxLength`
// error when the input string's length exceeds the specified maximum length.
// This validates the OCPP 1.6J specification's length constraints for CiString types.
// For example, an `IdToken` (CiString20Type) must not exceed 20 characters.
func TestCiString_TooLongFails(t *testing.T) {
	t.Parallel()

	// Scenario: Attempt to create a CiString with a length greater than its defined maximum.
	// This simulates an invalid `IdToken` or other CiString field that is too long.
	_, err := setCiString(strings.Repeat("A", 21), 20)
	// Expectation: The function should return an `ErrExceedsMaxLength` error.
	if !errors.Is(err, ErrExceedsMaxLength) {
		t.Errorf("expected ErrExceedsMaxLength, got: %v", err)
	}
}

// TestCiString_ValidStringPasses verifies that `setCiString` successfully creates
// a `ciString` instance without error when provided with a valid string.
// A valid string is one that is within the specified maximum length and contains
// only printable ASCII characters.
func TestCiString_ValidStringPasses(t *testing.T) {
	t.Parallel()

	// Scenario: Provide a string that conforms to all CiString validation rules.
	// This represents a typical valid input for any CiString field in OCPP messages.
	_, err := setCiString("OCPP16-Test", 20)
	// Expectation: No error should be returned, indicating successful validation.
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

// TestCiString_EmptyStringPasses verifies that `setCiString` accepts an empty string
// as a valid input. Empty strings are often used in OCPP to signify optional fields
// that are not present.
func TestCiString_EmptyStringPasses(t *testing.T) {
	t.Parallel()

	// Scenario: Provide an empty string as input.
	// This tests the edge case of a zero-length string, which should be valid for optional fields.
	_, err := setCiString("", 20)
	// Expectation: No error should be returned, as empty strings are valid for CiString types.
	if err != nil {
		t.Errorf(ErrExpectedNoError, err)
	}
}

// TestCiString_ValueReturnsRaw verifies that the `value()` method of the `ciString`
// struct correctly returns the original raw string that was used to create the instance.
func TestCiString_ValueReturnsRaw(t *testing.T) {
	t.Parallel()

	// Scenario: Create a CiString with a known input and then retrieve its value.
	// This ensures the internal storage and retrieval mechanism is working correctly.
	input := "OCPP-Valid"

	ciStr, err := setCiString(input, 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Expectation: The retrieved value should be identical to the original input string.
	if ciStr.value() != input {
		t.Errorf(ErrExpectedValueMismatch, input, ciStr.value())
	}
}

// TestCiString_NonPrintableCharacterFails verifies that `setCiString` returns an
// `ErrNonPrintableASCII` error when the input string contains characters outside
// the printable ASCII range (ASCII 32-126).
// This enforces the OCPP 1.6J specification's requirement for CiString character sets.
func TestCiString_NonPrintableCharacterFails(t *testing.T) {
	t.Parallel()

	// Scenario: Provide a string containing a non-printable ASCII character (e.g., ASCII Bell character).
	// This tests an edge case where the string's content violates the character set constraint.
	_, err := setCiString("Hello\aWorld", 20)
	// Expectation: The function should return an `ErrNonPrintableASCII` error.
	if !errors.Is(err, ErrNonPrintableASCII) {
		t.Errorf("expected ErrNonPrintableASCII, got: %v", err)
	}
}
