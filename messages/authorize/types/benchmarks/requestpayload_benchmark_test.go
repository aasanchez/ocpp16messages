package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func BenchmarkRequestPayloadValidate(b *testing.B) {
	payload := authorizetypes.RequestPayload{IdTag: "ABC1234567890123456789"}

	b.ResetTimer()

	for range b.N {
		_ = payload.Validate()
	}
}

func BenchmarkRequestPayloadValue(b *testing.B) {
	payload := authorizetypes.RequestPayload{IdTag: "ABC1234567890123456789"}

	b.ResetTimer()

	for range b.N {
		_ = payload.Value()
	}
}
