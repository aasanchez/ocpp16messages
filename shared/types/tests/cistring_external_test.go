// Package types_test contains blackbox (external) tests for the public `CiStringXXType`
// types defined in the `shared/types` package.
//
// This test file exists to verify the end-to-end functionality and adherence to OCPP 1.6J
// specifications for the various fixed-length case-insensitive string types (`CiString20Type`,
// `CiString25Type`, `CiString50Type`, `CiString255Type`, `CiString500Type`). Unlike the
// internal `cistring_test.go`, these tests interact with the public API (`SetCiStringXXType`
// and `Validate` methods) to ensure that the types behave correctly from a consumer's perspective.
//
// The tests cover the following real-world use cases and boundary conditions:
//   - **Successful Creation and Validation**: Confirms that each `CiStringXXType` can be
//     successfully created from a valid input string (within its specified maximum length)
//     and that its `Validate()` method returns no error. This simulates typical valid data
//     exchange in OCPP messages.
//   - **Maximum Length Enforcement**: Specifically tests that attempting to create a `CiString20Type`
//     with a string longer than 20 characters results in an `ErrExceedsMaxLength` error.
//     This validates the strict length constraints imposed by the OCPP 1.6J specification
//     for fields like `IdToken` (OCPP 1.6J Part 2, Appendix 3: "Data Types").
//   - **Value Retrieval**: Ensures that the `Value()` method for each `CiStringXXType`
//     correctly returns the original string content.
//
// This blackbox testing approach complements the whitebox tests by validating the public
// interface and ensuring that the combined internal logic (from `ciString` and `setCiString`)
// correctly translates into compliant external behavior.
package sharedtypes_test

import (
	"errors"
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// TestCiString20Type_CreateAndValidate verifies the successful creation and validation
// of a `CiString20Type` instance with a valid input string.
// This test ensures that the `SetCiString20Type` constructor and the `Validate` method
// correctly handle strings that comply with the 20-character maximum length for this type
// as specified in OCPP 1.6J Part 2, Appendix 3: "Data Types" (e.g., for `IdToken`).
func TestCiString20Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	// Scenario: Provide a valid 20-character string, which is the maximum allowed length.
	input := "ABCDEFGHIJKLMNOPQRST"
	ciStr, err := st.SetCiString20Type(input)

	// Expectation: No error during creation and validation, and the stored value matches the input.
	if err != nil {
		t.Fatalf("CiString20 creation failed: %v", err)
	}

	err = ciStr.Validate()
	if err != nil {
		t.Errorf("CiString20 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

// TestCiString20Type_CreateTooLongFails verifies that `SetCiString20Type` returns
// an `ErrExceedsMaxLength` error when the input string exceeds the 20-character limit.
// This directly validates the length constraint for `CiString20Type` as per OCPP 1.6J.
func TestCiString20Type_CreateTooLongFails(t *testing.T) {
	t.Parallel()

	// Scenario: Attempt to create a CiString20Type with a string longer than 20 characters.
	// This simulates an invalid `IdToken` that violates the OCPP specification.
	_, err := st.SetCiString20Type(strings.Repeat("X", 21))
	// Expectation: An `ErrExceedsMaxLength` error should be returned.
	if !errors.Is(err, st.ErrExceedsMaxLength) {
		t.Errorf("expected ErrExceedsMaxLength for >20 chars, got: %v", err)
	}
}

// TestCiString25Type_CreateAndValidate verifies the successful creation and validation
// of a `CiString25Type` instance with a valid input string.
// This test ensures compliance with the 25-character maximum length for this type
// as specified in OCPP 1.6J Part 2, Appendix 3: "Data Types" (e.g., for `chargePointVendor`).
func TestCiString25Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	// Scenario: Provide a valid 25-character string.
	input := strings.Repeat("B", 25)
	ciStr, err := st.SetCiString25Type(input)

	// Expectation: No error during creation and validation, and the stored value matches the input.
	if err != nil {
		t.Fatalf("CiString25 creation failed: %v", err)
	}

	err = ciStr.Validate()
	if err != nil {
		t.Errorf("CiString25 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

// TestCiString50Type_CreateAndValidate verifies the successful creation and validation
// of a `CiString50Type` instance with a valid input string.
// This test ensures compliance with the 50-character maximum length for this type
// as specified in OCPP 1.6J Part 2, Appendix 3: "Data Types" (e.g., for `chargePointModel`).
func TestCiString50Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	// Scenario: Provide a valid 50-character string.
	input := strings.Repeat("C", 50)
	ciStr, err := st.SetCiString50Type(input)

	// Expectation: No error during creation and validation, and the stored value matches the input.
	if err != nil {
		t.Fatalf("CiString50 creation failed: %v", err)
	}

	err = ciStr.Validate()
	if err != nil {
		t.Errorf("CiString50 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

// TestCiString255Type_CreateAndValidate verifies the successful creation and validation
// of a `CiString255Type` instance with a valid input string.
// This test ensures compliance with the 255-character maximum length for this type
// as specified in OCPP 1.6J Part 2, Appendix 3: "Data Types" (e.g., for `firmwareVersion`).
func TestCiString255Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	// Scenario: Provide a valid 255-character string.
	input := strings.Repeat("D", 255)
	ciStr, err := st.SetCiString255Type(input)

	// Expectation: No error during creation and validation, and the stored value matches the input.
	if err != nil {
		t.Fatalf("CiString255 creation failed: %v", err)
	}

	err = ciStr.Validate()
	if err != nil {
		t.Errorf("CiString255 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}

// TestCiString500Type_CreateAndValidate verifies the successful creation and validation
// of a `CiString500Type` instance with a valid input string.
// This test ensures compliance with the 500-character maximum length for this type
// as specified in OCPP 1.6J Part 2, Appendix 3: "Data Types" (e.g., for `diagnosticsStatus`).
func TestCiString500Type_CreateAndValidate(t *testing.T) {
	t.Parallel()

	// Scenario: Provide a valid 500-character string.
	input := strings.Repeat("E", 500)
	ciStr, err := st.SetCiString500Type(input)

	// Expectation: No error during creation and validation, and the stored value matches the input.
	if err != nil {
		t.Fatalf("CiString500 creation failed: %v", err)
	}

	err = ciStr.Validate()
	if err != nil {
		t.Errorf("CiString500 validation failed: %v", err)
	}

	if ciStr.Value() != input {
		t.Errorf(st.ErrExpectedValueMismatch, input, ciStr.Value())
	}
}
