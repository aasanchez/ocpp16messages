package cancelreservation

import (
	"strings"
	"testing"

	cancelreservationtypes "github.com/aasanchez/ocpp16messages/messages/cancelreservation/types"
)

func TestRequest_ValidReservationId(t *testing.T) {
	t.Parallel()

	input := cancelreservationtypes.RequestPayload{ReservationId: "1234"}
	expected := RequestMessage{ReservationId: 1234}

	got, err := Request(input)
	if err != nil {
		t.Fatalf("did not expect error, got: %v", err)
	}

	if got != expected {
		t.Errorf("unexpected result, expected: %+v, got: %+v", expected, got)
	}
}

func TestRequest_InvalidReservationId_NotANumber(t *testing.T) {
	t.Parallel()

	input := cancelreservationtypes.RequestPayload{ReservationId: "abc"}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error but got nil")
	}

	if !strings.Contains(err.Error(), "failed to parse ReservationId") {
		t.Errorf("expected error to contain: 'failed to parse ReservationId', got: %v", err)
	}
}

func TestRequest_EmptyReservationId(t *testing.T) {
	t.Parallel()

	input := cancelreservationtypes.RequestPayload{ReservationId: ""}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error but got nil")
	}

	if !strings.Contains(err.Error(), "failed to parse ReservationId") {
		t.Errorf("expected error to contain: 'failed to parse ReservationId', got: %v", err)
	}
}
