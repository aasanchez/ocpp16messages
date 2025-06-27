package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func BenchmarkConfirmationPayload_Validate_minimal(b *testing.B) {
	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      authorizetypes.Accepted,
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	var errResult error

	b.Cleanup(func() { _ = errResult })

	for range b.N {
		errResult = payload.Validate()
	}
}

func BenchmarkConfirmationPayload_Validate_full(b *testing.B) {
	expiry := "2027-02-12T10:03:04Z"
	parent := "ABC123"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      authorizetypes.Accepted,
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	var errResult error

	b.Cleanup(func() { _ = errResult })

	for range b.N {
		errResult = payload.Validate()
	}
}

func BenchmarkConfirmationPayload_Value_minimal(b *testing.B) {
	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      authorizetypes.Accepted,
			ExpiryDate:  nil,
			ParentIdTag: nil,
		},
	}

	var result authorizetypes.ConfirmationPayload

	b.Cleanup(func() { _ = result })

	for range b.N {
		result = payload.Value()
	}
}

func BenchmarkConfirmationPayload_Value_full(b *testing.B) {
	expiry := "2029-04-12T10:03:04Z"
	parent := "ABC1234"

	payload := authorizetypes.ConfirmationPayload{
		IdTagInfo: authorizetypes.IdTagInfoPayload{
			Status:      authorizetypes.Accepted,
			ExpiryDate:  &expiry,
			ParentIdTag: &parent,
		},
	}

	var result authorizetypes.ConfirmationPayload

	b.Cleanup(func() { _ = result })

	for range b.N {
		result = payload.Value()
	}
}
