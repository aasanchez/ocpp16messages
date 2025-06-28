package changeavailability

import (
	"strings"
	"testing"

	cat "github.com/aasanchez/ocpp16messages/messages/changeavailability/types"
)

func TestConfirmation_Valid(t *testing.T) {
	t.Parallel()

	input := cat.ConfirmationPayload{Status: cat.ChangeAvailabilityStatusAccepted}
	msg, err := Confirmation(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg.Status.Value() != cat.ChangeAvailabilityStatusAccepted {
		t.Errorf("expected Status %q, got %q", cat.ChangeAvailabilityStatusAccepted, msg.Status.Value())
	}
}

func TestConfirmation_EmptyStatus(t *testing.T) {
	t.Parallel()

	input := cat.ConfirmationPayload{Status: ""}
	_, err := Confirmation(input)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestConfirmation_InvalidStatus(t *testing.T) {
	t.Parallel()

	input := cat.ConfirmationPayload{Status: "Invalid"}
	_, err := Confirmation(input)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "failed to parse Status") {
		t.Errorf("expected parse error, got: %v", err)
	}
}
