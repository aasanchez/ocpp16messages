package cancelreservationtypes_test

import (
	"testing"

	cancelreservationtypes "github.com/aasanchez/ocpp16messages/messages/cancelreservation/types"
)

func BenchmarkCancelReservationStatus_Accepted(b *testing.B) {
	for range b.N {
		_, _ = cancelreservationtypes.CancelReservationStatus(cancelreservationtypes.CancelReservationAccepted)
	}
}

func BenchmarkCancelReservationStatus_Rejected(b *testing.B) {
	for range b.N {
		_, _ = cancelreservationtypes.CancelReservationStatus(cancelreservationtypes.CancelReservationRejected)
	}
}

func BenchmarkCancelReservationStatus_Invalid(b *testing.B) {
	for range b.N {
		_, _ = cancelreservationtypes.CancelReservationStatus("NotAStatus")
	}
}
