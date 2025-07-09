package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func BenchmarkConfirmationPayloadValidate(b *testing.B) {
	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status: authorizetypes.Accepted,
			ExpiryDate: nil,
			ParentIdTag: nil,
		},
	}

	b.ResetTimer()

	for range b.N {
		_ = payload.Validate()
	}
}

func BenchmarkConfirmationPayloadValue(b *testing.B) {
	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status: authorizetypes.Accepted,
			ExpiryDate: nil,
			ParentIdTag: nil,
		},
	}

	b.ResetTimer()

	for range b.N {
		_ = payload.Value()
	}
}
