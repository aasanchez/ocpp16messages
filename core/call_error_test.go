package core

import (
	"testing"
)

func TestValidateCallError_Valid(t *testing.T) {
	msg := []any{
		4,
		"12345",
		"ProtocolError",
		"Unsupported operation",
		map[string]any{"detail": "info"},
	}

	result, err := ValidateCallError(msg)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result == nil || result.UniqueID != "12345" {
		t.Errorf("unexpected result: %+v", result)
	}
}

func TestValidateCallError_InvalidType(t *testing.T) {
	msg := []any{
		3,
		"12345",
		"InternalError",
		"Something went wrong",
		map[string]any{},
	}

	_, err := ValidateCallError(msg)
	if err == nil {
		t.Error("expected error for wrong MessageTypeId")
	}
}

func TestValidateCallError_MissingFields(t *testing.T) {
	msg := []any{4, "id", "Error"}

	_, err := ValidateCallError(msg)
	if err == nil {
		t.Error("expected error for missing elements")
	}
}

func TestValidateCallError_InvalidJSONDetails(t *testing.T) {
	msg := []any{
		4,
		"987",
		"FormatViolation",
		"Description",
		func() {}, // Not JSON serializable
	}

	_, err := ValidateCallError(msg)
	if err == nil {
		t.Error("expected JSON marshal error")
	}
}

func TestValidateCallError_MissingUniqueID(t *testing.T) {
	msg := []any{4, "", "Code", "Description", map[string]any{}}
	_, err := ValidateCallError(msg)
	if err == nil {
		t.Error("expected error for missing UniqueID")
	}
}

func TestValidateCallError_MissingErrorCode(t *testing.T) {
	msg := []any{4, "id", "", "Description", map[string]any{}}
	_, err := ValidateCallError(msg)
	if err == nil {
		t.Error("expected error for missing ErrorCode")
	}
}

func TestValidateCallError_MissingErrorDescription(t *testing.T) {
	msg := []any{4, "id", "Code", "", map[string]any{}}
	_, err := ValidateCallError(msg)
	if err == nil {
		t.Error("expected error for missing ErrorDescription")
	}
}
