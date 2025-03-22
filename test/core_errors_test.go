package test

import (
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

func TestNewFieldError(t *testing.T) {
	err := core.NewFieldError("idTag", "cannot be empty")
	var ferr *core.FieldError

	if !errors.As(err, &ferr) {
		t.Errorf("Expected FieldError type, got %T", err)
	}

	expected := "idTag: cannot be empty"
	if err.Error() != expected {
		t.Errorf("Expected error %q, got %q", expected, err.Error())
	}

	if ferr.Field != "idTag" {
		t.Errorf("Expected field 'idTag', got %s", ferr.Field)
	}

	if ferr.Reason != "cannot be empty" {
		t.Errorf("Expected reason 'cannot be empty', got %s", ferr.Reason)
	}
}
