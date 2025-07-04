package changeavailability_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/messages/changeavailability"
	changeavailabilitytypes "github.com/aasanchez/ocpp16messages/messages/changeavailability/types"
)

func BenchmarkRequest_Valid(b *testing.B) {
	input := changeavailabilitytypes.RequestPayload{ConnectorId: "1", Type: changeavailabilitytypes.AvailabilityTypeOperative}

	b.ReportAllocs()
	for range b.N {
		_, err := changeavailability.Request(input)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}
