package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func BenchmarkAuthorizationStatus_Create(b *testing.B) {
	input := "Accepted"

	var result authorizetypes.AuthorizationStatusType

	b.ResetTimer()

	for range b.N {
		status, err := authorizetypes.AuthorizationStatus(input)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}

		result = status
	}

	_ = result
}

func BenchmarkAuthorizationStatus_Value(b *testing.B) {
	status, err := authorizetypes.AuthorizationStatus("Accepted")

	if err != nil {
		b.Fatalf("setup failed: %v", err)
	}

	var result string

	b.ResetTimer()

	for range b.N {
		result = status.Value()
	}

	_ = result
}
