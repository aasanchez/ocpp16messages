package changeavailability

import (
	"testing"

	changeavailabilitytypes "github.com/aasanchez/ocpp16messages/messages/changeavailability/types"
)

func TestRequest_Valid(t *testing.T) {
	t.Parallel()

	input := changeavailabilitytypes.RequestPayload{ConnectorId: "1", Type: changeavailabilitytypes.AvailabilityTypeOperative}
	if err := input.Validate(); err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}

	msg, err := Request(input)
	if err != nil {
		t.Fatalf("Request() returned unexpected error: %v", err)
	}

	if msg.ConnectorId != 1 {
		t.Errorf("expected ConnectorId %d, got %d", 1, msg.ConnectorId)
	}
	if msg.Type.Value() != changeavailabilitytypes.AvailabilityTypeOperative {
		t.Errorf("expected Type %q, got %q", changeavailabilitytypes.AvailabilityTypeOperative, msg.Type.Value())
	}
}

func TestRequest_MissingFields(t *testing.T) {
	t.Parallel()

	input := changeavailabilitytypes.RequestPayload{ConnectorId: "", Type: ""}
	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for missing required fields, got nil")
	}
}

func TestRequest_InvalidConnectorId(t *testing.T) {
	t.Parallel()

	input := changeavailabilitytypes.RequestPayload{ConnectorId: "x", Type: changeavailabilitytypes.AvailabilityTypeOperative}
	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for invalid ConnectorId, got nil")
	}
}

func TestRequest_InvalidType(t *testing.T) {
	t.Parallel()

	input := changeavailabilitytypes.RequestPayload{ConnectorId: "1", Type: "Invalid"}
	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for invalid Type, got nil")
	}
	if _, ok := err.(error); !ok {
		t.Errorf("expected error type, got %T", err)
	}
}
