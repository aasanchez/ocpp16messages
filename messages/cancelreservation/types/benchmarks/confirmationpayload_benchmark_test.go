package cancelreservationtypes_test

import (
	"testing"

	cancelreservationtypes "github.com/aasanchez/ocpp16messages/messages/cancelreservation/types"
)

func BenchmarkConfirmationPayload_Validate_Valid(b *testing.B) {
	payload := cancelreservationtypes.ConfirmationPayload{Status: "Accepted"}
	for range b.N {
		_ = payload.Validate()
	}
}

func BenchmarkConfirmationPayload_Validate_Empty(b *testing.B) {
	payload := cancelreservationtypes.ConfirmationPayload{Status: ""}
	for range b.N {
		_ = payload.Validate()
	}
}
