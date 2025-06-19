package authorizetypes_test

import (
	"errors"
	"testing"

	types "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func TestRequestPayloadType_Validate_ValidIdTag(t *testing.T) {
	t.Parallel()

	payload := types.RequestPayload{
		IdTag: "A123456789B987654321", // exactly 20 chars
	}

	if err := payload.Validate(); err != nil {
		t.Errorf("expected valid idTag, got error: %v", err)
	}
}

func TestRequestPayloadType_Validate_EmptyIdTag(t *testing.T) {
	t.Parallel()

	payload := types.RequestPayload{IdTag: ""}

	err := payload.Validate()
	if err == nil {
		t.Fatal("expected error for empty idTag, got nil")
	}

	if !errors.Is(err, types.ErrInvalidIdTag) {
		t.Errorf("expected ErrInvalidIdTag, got: %v", err)
	}
}

func TestRequestPayloadType_Validate_TooLongIdTag(t *testing.T) {
	t.Parallel()

	input := types.RequestPayload{
		IdTag: "THIS-ID-TAG-IS-TOO-LONG-FOR-CISTRING20",
	}

	err := input.Validate()
	if err == nil {
		t.Fatal("expected error for too long idTag, got nil")
	}
}

func TestRequestPayloadType_Value(t *testing.T) {
	t.Parallel()

	expected := "SampleIdTag"
	payload := types.RequestPayload{IdTag: expected}

	if payload.Value() != expected {
		t.Errorf("Value() returned %q, expected %q", payload.Value(), expected)
	}
}
