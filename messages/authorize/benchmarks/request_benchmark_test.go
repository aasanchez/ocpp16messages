package authorize_benchmark_test

import (
	"testing"

	ma "github.com/aasanchez/ocpp16messages/messages/authorize"
	mat "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func BenchmarkAuthorizeRequest(b *testing.B) {
	input := mat.RequestPayload{IdTag: "ABC1234567890XYZ78"}

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		_, err := ma.Request(input)
		if err != nil {
			b.Fatalf("Request() returned unexpected error: %v", err)
		}
	}
}
