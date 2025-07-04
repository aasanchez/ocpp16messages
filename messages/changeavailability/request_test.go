package changeavailability

import (
	"testing"

	cat "github.com/aasanchez/ocpp16messages/messages/changeavailability/types"
)

func TestRequest_Valid(t *testing.T) {
	t.Parallel()

	input := cat.RequestPayload{ConnectorId: "1", Type: cat.AvailabilityTypeOperative}
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

	if msg.Type.Value() != cat.AvailabilityTypeOperative {
		t.Errorf("expected Type %q, got %q", cat.AvailabilityTypeOperative, msg.Type.Value())
	}
}

func TestRequest_MissingFields(t *testing.T) {
	t.Parallel()

	input := cat.RequestPayload{ConnectorId: "", Type: ""}

	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for missing required fields, got nil")
	}
}

func TestRequest_InvalidConnectorId(t *testing.T) {
	t.Parallel()

	input := cat.RequestPayload{ConnectorId: "x", Type: cat.AvailabilityTypeOperative}
	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for invalid ConnectorId, got nil")
	}
}

func TestRequest_InvalidType(t *testing.T) {
	t.Parallel()

	input := cat.RequestPayload{ConnectorId: "1", Type: "Invalid"}
	_, err := Request(input)
	if err == nil {
		t.Fatal("expected error for invalid Type, got nil")
	}

}
