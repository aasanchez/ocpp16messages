package cancelreservation_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/messages/cancelreservation"
	cancelreservationtypes "github.com/aasanchez/ocpp16messages/messages/cancelreservation/types"
)

func BenchmarkRequest_Valid(b *testing.B) {
	input := cancelreservationtypes.RequestPayload{ReservationId: "1000"}
	for range b.N {
		_, _ = cancelreservation.Request(input)
	}
}

func BenchmarkRequest_Invalid(b *testing.B) {
	input := cancelreservationtypes.RequestPayload{ReservationId: "not-a-number"}
	for range b.N {
		_, _ = cancelreservation.Request(input)
	}
}
