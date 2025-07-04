package cancelreservationtypes_test

import (
	"testing"

	cancelreservationtypes "github.com/aasanchez/ocpp16messages/messages/cancelreservation/types"
)

func BenchmarkRequestPayload_Validate_Valid(b *testing.B) {
	payload := cancelreservationtypes.RequestPayload{ReservationId: "12345"}
	for range b.N {
		_ = payload.Validate()
	}
}

func BenchmarkRequestPayload_Validate_Empty(b *testing.B) {
	payload := cancelreservationtypes.RequestPayload{ReservationId: ""}
	for range b.N {
		_ = payload.Validate()
	}
}
