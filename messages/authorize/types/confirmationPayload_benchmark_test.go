package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func BenchmarkConfirmationPayload_Validate_minimal(b *testing.B) {
	var errResult error

	b.Cleanup(func() { _ = errResult })

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      authorizetypes.Accepted,
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	for i := 0; i < b.N; i++ {
		errResult = payload.Validate()
	}
}

func BenchmarkConfirmationPayload_Validate_full(b *testing.B) {
	var errResult error

	b.Cleanup(func() { _ = errResult })

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
	var resultConfirmation authorizetypes.ConfirmationPayload

	b.Cleanup(func() { _ = resultConfirmation })

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      authorizetypes.Accepted,
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	for i := 0; i < b.N; i++ {
		resultConfirmation = payload.Value()
	}
}

func BenchmarkConfirmationPayload_Value_full(b *testing.B) {
	var resultConfirmation authorizetypes.ConfirmationPayload

	b.Cleanup(func() { _ = resultConfirmation })

	expiry := "2027-04-12T10:03:04Z"
	parent := "ABC1234"

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
