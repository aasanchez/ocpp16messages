package cancelreservationtypes

import (
	"strings"
	"testing"
)

func TestRequestPayload_Validate_OK(t *testing.T) {
	t.Parallel()

	payload := RequestPayload{ReservationId: "42"}

	err := payload.Validate()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestRequestPayload_Validate_EmptyReservationId(t *testing.T) {
	t.Parallel()
	payload := RequestPayload{ReservationId: ""}

	err := payload.Validate()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "missing required field: ReservationId") {
		t.Errorf("unexpected error: %v", err)
	}
}
