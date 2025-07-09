package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func BenchmarkIdTagInfo(b *testing.B) {
	input := authorizetypes.IdTagInfoPayload{
		Status:      authorizetypes.Accepted,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}

	b.ResetTimer()

	for range b.N {
		_, err := authorizetypes.IdTagInfo(input)
		if err != nil {
			b.Fatalf("IdTagInfo returned an error: %v", err)
		}
	}
}

func BenchmarkIdTagInfoValue(b *testing.B) {
	input := authorizetypes.IdTagInfoPayload{
		Status:      authorizetypes.Accepted,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}
	info, _ := authorizetypes.IdTagInfo(input)

	b.ResetTimer()

	for range b.N {
		_ = info.Value()
	}
}
