package test

import (
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

func TestNewFieldError(t *testing.T) {
	err := core.NewFieldError("idTag", "must not be empty")
	expected := "invalid idTag: must not be empty"

	if err == nil || err.Error() != expected {
		t.Errorf("expected %q, got %v", expected, err)
	}
}
