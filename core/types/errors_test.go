package types

import (
	"testing"
)

func TestFieldError_ErrorOutput(t *testing.T) {
	err := NewFieldError("idTag", "must not be empty")
	expected := "invalid field 'idTag': must not be empty"
	if err.Error() != expected {
		t.Errorf("unexpected error message: got %q, want %q", err.Error(), expected)
	}
}

func TestFieldError_ImplementsErrorInterface(t *testing.T) {
	err := NewFieldError("status", "invalid value")
	if err.Error() == "" {
		t.Fatal("FieldError does not implement the error interface")
	}
}
