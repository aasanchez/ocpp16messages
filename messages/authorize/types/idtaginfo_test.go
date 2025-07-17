// Package authorizetypes_test contains unit tests for the `IdTagInfo` type
// and its associated `IdTagInfoPayload` within the `messages/authorize/types` package.
//
// This test file exists to ensure that the `IdTagInfo` type, which encapsulates
// authorization-related information for an IdTag in OCPP 1.6J, correctly handles
// various combinations of its fields and validates them according to the protocol's rules.
//
// The tests cover the following real-world use cases and edge cases:
//   - **Status-Only Payload**: Verifies that a payload with only the required `Status` field
//     is correctly processed and its value is accessible.
//   - **With ExpiryDate**: Confirms that the optional `ExpiryDate` field is correctly parsed
//     and stored when provided.
//   - **With ParentIdTag**: Ensures that the optional `ParentIdTag` field is correctly parsed
//     and stored, including its validation as a `CiString20Type`.
//   - **All Fields Present**: Validates the correct handling of a payload where all optional
//     fields (`ExpiryDate` and `ParentIdTag`) are provided along with the required `Status`.
//   - **Invalid Status**: Tests that an invalid `Status` value (not part of the predefined
//     `AuthorizationStatusType` enum) results in an error.
//   - **Invalid ExpiryDate Format**: Verifies that a malformed `ExpiryDate` string (not ISO 8601)
//     causes an error during parsing.
//   - **Invalid ParentIdTag Length**: Ensures that a `ParentIdTag` exceeding its maximum
//     allowed length (20 characters for `CiString20Type`) results in an error.
//
// These tests collectively ensure the robustness and compliance of the `IdTagInfo` type
// with the OCPP 1.6J specification, particularly concerning the structure and validation
// of IdTag information as detailed in OCPP 1.6J Part 2, Section 4.1.2.2, Table 14.
package authorizetypes

import (
	"testing"
)

// TestIdTagInfo_statusOnly verifies that `IdTagInfo` can be created with only the required `Status` field.
// This simulates a basic authorization response where only the status is relevant.
func TestIdTagInfo_statusOnly(t *testing.T) {
	t.Parallel()

	// Scenario: Create an `IdTagInfoPayload` with only the `Status` field set.
	input := IdTagInfoPayload{
		Status:      Accepted,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}

	// Expectation: No error during creation, and the `Status` value is correctly reflected.
	info, err := IdTagInfo(input)
	if err != nil {
		t.Fatalf("unexpected error statusOnly: %v", err)
	}

	val := info.Value()

	if val.Status != Accepted {
		t.Errorf("expected status %q, got %q", Accepted, val.Status)
	}

	if val.ExpiryDate != nil {
		t.Errorf("expected expiryDate nil, got %v", *val.ExpiryDate)
	}

	if val.ParentIdTag != nil {
		t.Errorf("expected parentIdTag nil, got %v", *val.ParentIdTag)
	}
}

// TestIdTagInfo_withExpiryDateOnly verifies that `IdTagInfo` correctly handles an optional `ExpiryDate`.
// This tests the parsing and storage of the `ExpiryDate` field.
func TestIdTagInfo_withExpiryDateOnly(t *testing.T) {
	t.Parallel()

	// Scenario: Create an `IdTagInfoPayload` with `Status` and a valid `ExpiryDate`.
	expiryDate := "2023-04-12T10:03:04Z"

	input := IdTagInfoPayload{
		Status:      Accepted,
		ExpiryDate:  &expiryDate,
		ParentIdTag: nil,
	}

	// Expectation: No error during creation, and the `ExpiryDate` is correctly stored.
	info, err := IdTagInfo(input)
	if err != nil {
		t.Fatalf("unexpected error withExpiryDateOnly: %v", err)
	}

	val := info.Value()
	if val.ExpiryDate == nil || *val.ExpiryDate != expiryDate {
		t.Errorf("expiryDate mismatch: want %s, got %v", expiryDate, val.ExpiryDate)
	}
}

// TestIdTagInfo_withParentIdTagOnly verifies that `IdTagInfo` correctly handles an optional `ParentIdTag`.
// This tests the parsing and storage of the `ParentIdTag` field, including its `CiString20Type` validation.
func TestIdTagInfo_withParentIdTagOnly(t *testing.T) {
	t.Parallel()

	// Scenario: Create an `IdTagInfoPayload` with `Status` and a valid `ParentIdTag`.
	parent := "ABC0101"

	input := IdTagInfoPayload{
		Status:      Accepted,
		ExpiryDate:  nil,
		ParentIdTag: &parent,
	}

	// Expectation: No error during creation, and the `ParentIdTag` is correctly stored.
	info, err := IdTagInfo(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	val := info.Value()
	if val.ParentIdTag == nil || *val.ParentIdTag != parent {
		t.Errorf("parentIdTag mismatch: want %s, got %v", parent, val.ParentIdTag)
	}
}

// TestIdTagInfo_allFieldsPresent verifies that `IdTagInfo` correctly handles a payload
// where all fields (`Status`, `ExpiryDate`, and `ParentIdTag`) are present and valid.
// This simulates a complete authorization response from the Central System.
func TestIdTagInfo_allFieldsPresent(t *testing.T) {
	t.Parallel()

	// Scenario: Create an `IdTagInfoPayload` with all fields populated with valid data.
	expiryDate := "2027-04-12T10:01:04Z"
	parent := "XYZ-987"

	input := IdTagInfoPayload{
		Status:      Accepted,
		ExpiryDate:  &expiryDate,
		ParentIdTag: &parent,
	}

	// Expectation: No error during creation, and all fields are correctly stored and accessible.
	info, err := IdTagInfo(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	val := info.Value()

	if val.Status != Accepted {
		t.Errorf("status mismatch: got %s", val.Status)
	}

	if val.ExpiryDate == nil || *val.ExpiryDate != expiryDate {
		t.Errorf("expiryDate mismatch: want %s, got %v", expiryDate, val.ExpiryDate)
	}

	if val.ParentIdTag == nil || *val.ParentIdTag != parent {
		t.Errorf("parentIdTag mismatch: want %s, got %v", parent, val.ParentIdTag)
	}
}

// TestIdTagInfo_invalidStatus verifies that `IdTagInfo` returns an error when an invalid
// `Status` value is provided in the payload.
// This enforces the `AuthorizationStatusType` enum validation as per OCPP 1.6J.
func TestIdTagInfo_invalidStatus(t *testing.T) {
	t.Parallel()

	// Scenario: Provide an `IdTagInfoPayload` with a `Status` that is not a valid `AuthorizationStatusType`.
	input := IdTagInfoPayload{
		Status:      "InvalidStatus",
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}

	// Expectation: An error should be returned, indicating an invalid status.
	_, err := IdTagInfo(input)
	if err == nil {
		t.Fatal("expected error for invalid status, got nil")
	}
}

// TestIdTagInfo_invalidExpiryDate verifies that `IdTagInfo` returns an error when the
// `ExpiryDate` field is not in a valid ISO 8601 (RFC3339) format.
// This ensures data integrity for date-time fields in OCPP messages.
func TestIdTagInfo_invalidExpiryDate(t *testing.T) {
	t.Parallel()

	// Scenario: Provide an `IdTagInfoPayload` with a malformed `ExpiryDate` string.
	bad := "not-a-date"

	input := IdTagInfoPayload{
		Status:      Accepted,
		ExpiryDate:  &bad,
		ParentIdTag: nil,
	}

	// Expectation: An error should be returned, indicating a parsing failure for the date.
	_, err := IdTagInfo(input)
	if err == nil {
		t.Fatal("expected error for invalid expiryDate, got nil")
	}
}

// TestIdTagInfo_invalidParentIdTag_tooLong verifies that `IdTagInfo` returns an error
// when the `ParentIdTag` exceeds its maximum allowed length (20 characters).
// This enforces the `CiString20Type` validation for the `ParentIdTag` field.
func TestIdTagInfo_invalidParentIdTag_tooLong(t *testing.T) {
	t.Parallel()

	// Scenario: Provide an `IdTagInfoPayload` with a `ParentIdTag` string longer than 20 characters.
	tooLong := "TOO-LONG-Id-TAG-FOR-CISTRING20-XXXXX"

	input := IdTagInfoPayload{
		Status:      Accepted,
		ExpiryDate:  nil,
		ParentIdTag: &tooLong,
	}

	// Expectation: An error should be returned, indicating that the `ParentIdTag` is too long.
	_, err := IdTagInfo(input)
	if err == nil {
		t.Fatal("expected error for invalid parentIdTag, got nil")
	}
}
