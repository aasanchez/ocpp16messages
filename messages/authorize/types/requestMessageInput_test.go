package authorizetypes

import (
	"reflect"
	"testing"
)

func TestRequestMessageInput_Validate_IdTagPresent(t *testing.T) {
	t.Parallel()

	input := RequestMessageInput{IdTag: "SOME-ID"}
	err := input.Validate()

	if err != nil {
		t.Errorf("expected validation to pass, got error: %v", err)
	}
}

func TestRequestMessageInput_Validate_IdTagMissing(t *testing.T) {
	t.Parallel()

	input := RequestMessageInput{}
	err := input.Validate()

	if err == nil {
		t.Error("expected validation to fail when IdTag is empty")
	}
}

func TestRequestMessageInput_Field_IdTagValueIsCorrect(t *testing.T) {
	t.Parallel()

	input := RequestMessageInput{IdTag: "ONLY"}
	if input.IdTag != "ONLY" {
		t.Errorf("expected IdTag to be 'ONLY', got '%s'", input.IdTag)
	}
}

func TestRequestMessageInput_HasOnlyOneField(t *testing.T) {
	t.Parallel()

	typ := reflect.TypeOf(RequestMessageInput{})
	numFields := typ.NumField()

	if numFields != 1 {
		t.Errorf("expected RequestMessageInput to have 1 field, got %d", numFields)
	}
}

func TestRequestMessageInput_FieldName_IsIdTag(t *testing.T) {
	t.Parallel()

	typ := reflect.TypeOf(RequestMessageInput{})
	field := typ.Field(0)

	if field.Name != "IdTag" {
		t.Errorf("expected first field to be named 'IdTag', got '%s'", field.Name)
	}
}
