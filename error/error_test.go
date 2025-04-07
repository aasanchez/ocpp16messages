package core

import (
	"testing"
)

func TestValidateCallErrorValid(t *testing.T) {
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

func TestValidateCallErrorInvalidType(t *testing.T) {
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

func TestValidateCallErrorMissingFields(t *testing.T) {
	msg := []any{4, "id", "Error"}

	_, err := ValidateCallError(msg)
	if err == nil {
		t.Error("expected error for missing elements")
	}
}

func TestValidateCallErrorInvalidJSONDetails(t *testing.T) {
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

func TestValidateCallErrorMissingUniqueID(t *testing.T) {
	msg := []any{4, "", "Code", "Description", map[string]any{}}
	_, err := ValidateCallError(msg)
	if err == nil {
		t.Error("expected error for missing UniqueID")
	}
}

func TestValidateCallErrorMissingErrorCode(t *testing.T) {
	msg := []any{4, "id", "", "Description", map[string]any{}}
	_, err := ValidateCallError(msg)
	if err == nil {
		t.Error("expected error for missing ErrorCode")
	}
}

func TestValidateCallErrorMissingErrorDescription(t *testing.T) {
	msg := []any{4, "id", "Code", "", map[string]any{}}
	_, err := ValidateCallError(msg)
	if err == nil {
		t.Error("expected error for missing ErrorDescription")
	}
}
