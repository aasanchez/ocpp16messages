package authorizetypes

import (
	"reflect"
	"testing"
)

func TestRequestPayload_Validate_IdTagPresent(t *testing.T) {
	t.Parallel()

	payload := RequestPayload{IdTag: "SOME-ID"}
	err := payload.Validate()

	if err != nil {
		t.Errorf("expected validation to pass, got error: %v", err)
	}
}

func TestRequestPayload_Validate_IdTagMissing(t *testing.T) {
	t.Parallel()

	payload := RequestPayload{IdTag: ""}
	err := payload.Validate()

	if err == nil {
		t.Error("expected validation to fail for empty IdTag")
	}
}

func TestRequestPayload_Value_ReturnsCorrectString(t *testing.T) {
	t.Parallel()

	payload := RequestPayload{IdTag: "VALID-123"}
	got := payload.Value()

	if got != "VALID-123" {
		t.Errorf("expected Value() to return 'VALID-123', got '%s'", got)
	}
}

func TestRequestPayload_Structure_HasOnlyOneField(t *testing.T) {
	t.Parallel()

	typ := reflect.TypeOf(RequestPayload{})
	numFields := typ.NumField()

	if numFields != 1 {
		t.Errorf("expected RequestPayload to have 1 field, got %d", numFields)
	}
}

func TestRequestPayload_FieldName_IsIdTag(t *testing.T) {
	t.Parallel()

	typ := reflect.TypeOf(RequestPayload{})
	field := typ.Field(0)

	if field.Name != "IdTag" {
		t.Errorf("expected first field to be named 'IdTag', got '%s'", field.Name)
	}
}
