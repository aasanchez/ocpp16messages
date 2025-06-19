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

	for range b.N {
		_, err := authorize.Confirmation(payload)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkConfirmation_invalidStatus(b *testing.B) {
	expiry := "2027-01-01T00:00:00Z"
	parent := "VALID-TAG"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      "INVALID-STATUS",
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	b.ReportAllocs()

	for range b.N {
		_, _ = authorize.Confirmation(payload)
	}
}

func BenchmarkConfirmation_invalidExpiryDate(b *testing.B) {
	invalidDate := "invalid-date"
	parent := "VALID-TAG"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      "Accepted",
			ExpiryDate:  &invalidDate,
			ParentIdTag: &parent,
		},
	}

	b.ReportAllocs()

	for range b.N {
		_, _ = authorize.Confirmation(payload)
	}
}

func BenchmarkConfirmation_invalidParentIdTag(b *testing.B) {
	expiry := "2027-01-01T00:00:00Z"
	invalid := "THIS_IS_WAY_TOO_LONG_FOR_A_CISTRING20_INPUT"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      "Accepted",
			ExpiryDate:  &expiry,
			ParentIdTag: &invalid,
		},
	}

	b.ReportAllocs()

	for range b.N {
		_, _ = authorize.Confirmation(payload)
	}
}
