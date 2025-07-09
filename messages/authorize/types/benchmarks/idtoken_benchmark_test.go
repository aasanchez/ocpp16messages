package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
	sharedtypes "github.com/aasanchez/ocpp16messages/shared/types"
)

func BenchmarkIdToken(b *testing.B) {
	str, _ := sharedtypes.SetCiString20Type("ABC1234567890123456789")

	b.ResetTimer()

	for range b.N {
		_, err := authorizetypes.IdToken(str)
		if err != nil {
			b.Fatalf("IdToken returned an error: %v", err)
		}
	}
}

func BenchmarkIdTokenValue(b *testing.B) {
	str, _ := sharedtypes.SetCiString20Type("ABC1234567890123456789")
	idToken, _ := authorizetypes.IdToken(str)

	b.ResetTimer()

	for range b.N {
		_ = idToken.Value()
	}
}
