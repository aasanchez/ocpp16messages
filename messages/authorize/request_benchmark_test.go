package authorize

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func BenchmarkRequest_Valid(b *testing.B) {
	input := authorizetypes.RequestPayload{
		IdTag: "ABC1234567890XYZ78",
	}

	b.ResetTimer()

	for range b.N {
		_, err := Request(input)
		if err != nil {
			b.Fatalf("unexpected error in benchmark: %v", err)
		}
	}
}
