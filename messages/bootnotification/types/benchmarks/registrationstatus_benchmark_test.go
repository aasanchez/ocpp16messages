package bootnotification_test

import (
	"testing"

	bootnotificationtypes "github.com/aasanchez/ocpp16messages/messages/bootnotification/types"
)

func BenchmarkRegistrationStatus_Accepted(b *testing.B) {
	for range b.N {
		_, err := bootnotificationtypes.RegistrationStatus("Accepted")
		if err != nil {
			b.Fatalf("unexpected error for Accepted: %v", err)
		}
	}
}

func BenchmarkRegistrationStatus_Pending(b *testing.B) {
	for range b.N {
		_, err := bootnotificationtypes.RegistrationStatus("Pending")
		if err != nil {
			b.Fatalf("unexpected error for Pending: %v", err)
		}
	}
}

func BenchmarkRegistrationStatus_Rejected(b *testing.B) {
	for range b.N {
		_, err := bootnotificationtypes.RegistrationStatus("Rejected")
		if err != nil {
			b.Fatalf("unexpected error for Rejected: %v", err)
		}
	}
}
