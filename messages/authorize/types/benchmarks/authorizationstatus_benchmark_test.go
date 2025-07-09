package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func BenchmarkAuthorizationStatus(b *testing.B) {
	for range b.N {
		_, err := authorizetypes.AuthorizationStatus(authorizetypes.Accepted)
		if err != nil {
			b.Fatalf("AuthorizationStatus returned an error: %v", err)
		}
	}
}

func BenchmarkAuthorizationStatusValue(b *testing.B) {
	status, _ := authorizetypes.AuthorizationStatus(authorizetypes.Accepted)

	b.ResetTimer()

	for range b.N {
		_ = status.Value()
	}
}
