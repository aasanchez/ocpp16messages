package types

import (
	"testing"
)

const valid_id_tag = "valid-id-tag"

func TestNewIdTag_Valid(t *testing.T) {
	idTag, err := NewIdTag(valid_id_tag)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if idTag.String() != valid_id_tag {
		t.Errorf("unexpected value: got %q, want %q", idTag.String(), valid_id_tag)
	}
}

func TestNewIdTag_TooLong(t *testing.T) {
	longValue := "this-id-tag-is-way-too-long-for-ocpp"
	_, err := NewIdTag(longValue)
	if err == nil {
		t.Fatal("expected error for too long idTag, got nil")
	}
}

func TestIdTag_String_NilReceiver(t *testing.T) {
	var idTag *IdTag
	if idTag.String() != "" {
		t.Errorf("expected empty string from nil receiver, got %q", idTag.String())
	}
}
