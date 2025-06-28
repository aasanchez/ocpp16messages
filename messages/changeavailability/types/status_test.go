package changeavailabilitytypes

import (
	"errors"
	"strings"
	"testing"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestChangeAvailabilityStatus_Accepted(t *testing.T) {
	t.Parallel()
	result, err := ChangeAvailabilityStatus(ChangeAvailabilityStatusAccepted)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Value() != ChangeAvailabilityStatusAccepted {
		t.Errorf("expected value %q, got %q", ChangeAvailabilityStatusAccepted, result.Value())
	}
}

func TestChangeAvailabilityStatus_Rejected(t *testing.T) {
	t.Parallel()
	result, err := ChangeAvailabilityStatus(ChangeAvailabilityStatusRejected)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Value() != ChangeAvailabilityStatusRejected {
		t.Errorf("expected value %q, got %q", ChangeAvailabilityStatusRejected, result.Value())
	}
}

func TestChangeAvailabilityStatus_Invalid(t *testing.T) {
	t.Parallel()
	_, err := ChangeAvailabilityStatus("Invalid")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, sharedtypes.ErrInvalidChangeAvailabilityStatus) {
		t.Errorf("error is not ErrInvalidChangeAvailabilityStatus: %v", err)
	}
	if !strings.Contains(err.Error(), "Invalid") {
		t.Errorf("error does not mention input: %v", err)
	}
}
