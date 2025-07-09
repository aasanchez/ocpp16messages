package authorize_benchmark_test

import (
	"testing"

	ma "github.com/aasanchez/ocpp16messages/messages/authorize"
	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func BenchmarkAuthorizeConfirmation(b *testing.B) {
	input := mat.ConfirmationPayload{
		IdTagInfo: mat.IdTagInfoPayload{
			Status:      "Accepted",
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		_, err := ma.Confirmation(input)
		if err != nil {
			b.Fatalf("Confirmation() returned unexpected error: %v", err)
		}
	}
}
