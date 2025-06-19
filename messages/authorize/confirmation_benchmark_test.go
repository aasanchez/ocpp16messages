package authorize_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/messages/authorize"
	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func BenchmarkConfirmation_valid(b *testing.B) {
	expiry := "2027-04-12T10:03:04Z"
	parent := "A632-E2BB0231072C"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      "Accepted",
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := authorize.Confirmation(payload)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkConfirmation_invalidStatus(b *testing.B) {
	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status: "INVALID-STATUS",
		},
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = authorize.Confirmation(payload)
	}
}

func BenchmarkConfirmation_invalidExpiryDate(b *testing.B) {
	invalidDate := "invalid-date"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:     "Accepted",
			ExpiryDate: &invalidDate,
		},
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = authorize.Confirmation(payload)
	}
}

func BenchmarkConfirmation_invalidParentIdTag(b *testing.B) {
	invalid := "THIS_IS_WAY_TOO_LONG_FOR_A_CISTRING20_INPUT"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      "Accepted",
			ParentIdTag: &invalid,
		},
	}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = authorize.Confirmation(payload)
	}
}
