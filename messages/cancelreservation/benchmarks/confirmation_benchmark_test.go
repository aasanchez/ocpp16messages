package cancelreservation_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/messages/cancelreservation"
	cancelreservationtypes "github.com/aasanchez/ocpp16messages/messages/cancelreservation/types"
)

func BenchmarkConfirmation_Accepted(b *testing.B) {
	input := cancelreservationtypes.ConfirmationPayload{Status: cancelreservationtypes.CancelReservationAccepted}
	for range b.N {
		_, _ = cancelreservation.Confirmation(input)
	}
}

func BenchmarkConfirmation_Rejected(b *testing.B) {
	input := cancelreservationtypes.ConfirmationPayload{Status: cancelreservationtypes.CancelReservationRejected}
	for range b.N {
		_, _ = cancelreservation.Confirmation(input)
	}
}

func BenchmarkConfirmation_EmptyStatus(b *testing.B) {
	input := cancelreservationtypes.ConfirmationPayload{Status: ""}
	for range b.N {
		_, _ = cancelreservation.Confirmation(input)
	}
}

func BenchmarkConfirmation_InvalidStatus(b *testing.B) {
	input := cancelreservationtypes.ConfirmationPayload{Status: "nonsense"}
	for range b.N {
		_, _ = cancelreservation.Confirmation(input)
	}
}
