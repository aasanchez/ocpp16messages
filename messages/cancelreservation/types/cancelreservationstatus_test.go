package cancelreservationtypes

import (
	"errors"
	"strings"
	"testing"

	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestCancelReservationStatus_Accepted(t *testing.T) {
	t.Parallel()

	result, err := CancelReservationStatus(CancelReservationAccepted)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Value() != CancelReservationAccepted {
		t.Errorf("expected value %q, got %q", CancelReservationAccepted, result.Value())
	}
}

func TestCancelReservationStatus_Rejected(t *testing.T) {
	t.Parallel()

	result, err := CancelReservationStatus(CancelReservationRejected)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Value() != CancelReservationRejected {
		t.Errorf("expected value %q, got %q", CancelReservationRejected, result.Value())
	}
}

func TestCancelReservationStatus_Invalid(t *testing.T) {
	t.Parallel()

	_, err := CancelReservationStatus("InvalidValue")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !errors.Is(err, sharedtypes.ErrInvalidCancelReservationStatus) {
		t.Errorf("error is not ErrInvalidCancelReservationStatus: %v", err)
	}

	if !strings.Contains(err.Error(), "InvalidValue") {
		t.Errorf("error does not mention input: %v", err)
	}
}
