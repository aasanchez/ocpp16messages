package core_test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

func TestValidateCallError_Valid(t *testing.T) {
	msg := []any{
		4,
		"12345",
		"ProtocolError",
		"Unsupported operation",
		map[string]any{"detail": "info"},
	}

	result, err := core.ValidateCallError(msg)
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

	_, err := core.ValidateCallError(msg)
	if err == nil {
		t.Error("expected error for wrong MessageTypeId")
	}
}

func TestValidateCallError_MissingFields(t *testing.T) {
	msg := []any{4, "id", "Error"}

	_, err := core.ValidateCallError(msg)
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

	_, err := core.ValidateCallError(msg)
	if err == nil {
		t.Error("expected JSON marshal error")
	}
}
