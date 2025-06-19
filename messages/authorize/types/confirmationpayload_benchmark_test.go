package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

var resultConfirmation authorizetypes.ConfirmationPayload
var errResult error

func BenchmarkConfirmationPayload_Validate_minimal(b *testing.B) {
	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status: authorizetypes.Accepted,
		},
	}

	for i := 0; i < b.N; i++ {
		errResult = payload.Validate()
	}
}

func BenchmarkConfirmationPayload_Validate_full(b *testing.B) {
	expiry := "2027-04-12T10:03:04Z"
	parent := "ABC123"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      authorizetypes.Accepted,
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	for i := 0; i < b.N; i++ {
		errResult = payload.Validate()
	}
}

func BenchmarkConfirmationPayload_Value_minimal(b *testing.B) {
	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status: authorizetypes.Accepted,
		},
	}

	for i := 0; i < b.N; i++ {
		resultConfirmation = payload.Value()
	}
}

func BenchmarkConfirmationPayload_Value_full(b *testing.B) {
	expiry := "2027-04-12T10:03:04Z"
	parent := "ABC123"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      authorizetypes.Accepted,
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	for i := 0; i < b.N; i++ {
		resultConfirmation = payload.Value()
	}
}
