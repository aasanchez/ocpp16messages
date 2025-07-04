package authorize_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func BenchmarkAuthorizeRequest(b *testing.B) {
	input := authorizetypes.RequestPayload{
		IdTag: "ABC1234567890XYZ78",
	}

	b.ResetTimer()

	for range b.N {
		_, err := authorize.Request(input)
		if err != nil {
			b.Fatalf("unexpected error in benchmark: %v", err)
		}
	}
}
