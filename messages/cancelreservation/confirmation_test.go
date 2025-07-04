package cancelreservation

import (
	"strings"
	"testing"

	cancelreservationtypes "github.com/aasanchez/ocpp16messages/messages/cancelreservation/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func TestConfirmation_Accepted(t *testing.T) {
	t.Parallel()

	input := cancelreservationtypes.ConfirmationPayload{Status: cancelreservationtypes.CancelReservationAccepted}

	msg, err := Confirmation(input)
	if err != nil {
		t.Fatalf(sharedtypes.ErrExpectedNoError, err)
	}

	if msg.Status.Value() != cancelreservationtypes.CancelReservationAccepted {
		t.Errorf("expected status %q, got %q", cancelreservationtypes.CancelReservationAccepted, msg.Status.Value())
	}
}

func TestConfirmation_Rejected(t *testing.T) {
	t.Parallel()

	input := cancelreservationtypes.ConfirmationPayload{Status: cancelreservationtypes.CancelReservationRejected}

	msg, err := Confirmation(input)
	if err != nil {
		t.Fatalf(sharedtypes.ErrExpectedNoError, err)
	}

	if msg.Status.Value() != cancelreservationtypes.CancelReservationRejected {
		t.Errorf("expected status %q, got %q", cancelreservationtypes.CancelReservationRejected, msg.Status.Value())
	}
}

func TestConfirmation_EmptyStatus(t *testing.T) {
	t.Parallel()

	input := cancelreservationtypes.ConfirmationPayload{Status: ""}

	_, err := Confirmation(input)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "missing required field: status") {
		t.Errorf(sharedtypes.ErrExpectedNoError, err)
	}
}

func TestConfirmation_InvalidStatus(t *testing.T) {
	t.Parallel()

	input := cancelreservationtypes.ConfirmationPayload{Status: "nonsense"}

	_, err := Confirmation(input)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "failed to parse Status") {
		t.Errorf(sharedtypes.ErrExpectedNoError, err)
	}
}
