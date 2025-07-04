package changeavailabilitytypes

import (
	"errors"
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestAvailabilityType_Operative(t *testing.T) {
	t.Parallel()

	result, err := AvailabilityType(AvailabilityTypeOperative)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Value() != AvailabilityTypeOperative {
		t.Errorf("expected value %q, got %q", AvailabilityTypeOperative, result.Value())
	}
}

func TestAvailabilityType_Inoperative(t *testing.T) {
	t.Parallel()
	result, err := AvailabilityType(AvailabilityTypeInoperative)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Value() != AvailabilityTypeInoperative {
		t.Errorf("expected value %q, got %q", AvailabilityTypeInoperative, result.Value())
	}
}

func TestAvailabilityType_Invalid(t *testing.T) {
	t.Parallel()
	_, err := AvailabilityType("InvalidValue")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, st.ErrInvalidChangeAvailabilityType) {
		t.Errorf("error is not ErrInvalidChangeAvailabilityType: %v", err)
	}
	if !strings.Contains(err.Error(), "InvalidValue") {
		t.Errorf("error does not mention input: %v", err)
	}
}
