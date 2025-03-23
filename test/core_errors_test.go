package test

import (
	"errors"
	"strings"
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

func TestNewFieldError(t *testing.T) {
	err := errors.New("cannot be empty")

	fieldErr := core.NewFieldError("fieldName", err)

	if fieldErr == nil {
		t.Fatal("expected FieldError, got nil")
	}

	expectedMsg := "invalid field 'fieldName': cannot be empty"
	if fieldErr.Error() != expectedMsg {
		t.Errorf("unexpected error message.\nExpected: %s\nGot     : %s", expectedMsg, fieldErr.Error())
	}

	if !strings.Contains(fieldErr.Error(), "fieldName") {
		t.Errorf("error message does not contain field name")
	}
}
