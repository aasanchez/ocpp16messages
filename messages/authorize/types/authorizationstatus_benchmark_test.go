package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

var statusSink authorizetypes.AuthorizationStatusType
var valueSink string

func BenchmarkAuthorizationStatus_Create(b *testing.B) {
	input := "Accepted"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		status, err := authorizetypes.AuthorizationStatus(input)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
		statusSink = status
	}
}

func BenchmarkAuthorizationStatus_Value(b *testing.B) {
	status, err := authorizetypes.AuthorizationStatus("Accepted")
	if err != nil {
		b.Fatalf("setup failed: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		valueSink = status.Value()
	}
}
