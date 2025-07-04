package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func newMinimalPayload() authorizetypes.IdTagInfoPayload {
	return authorizetypes.IdTagInfoPayload{
		Status:      authorizetypes.Accepted,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	}
}

func newFullPayload() authorizetypes.IdTagInfoPayload {
	expiry := "2027-04-12T14:03:04Z"
	parent := "ABC123456789"

	return authorizetypes.IdTagInfoPayload{
		Status:      authorizetypes.Accepted,
		ExpiryDate:  &expiry,
		ParentIdTag: &parent,
	}
}

func BenchmarkIdTagInfo_Create_Minimal(b *testing.B) {
	var result authorizetypes.IdTagInfoType

	payload := newMinimalPayload()

	for range b.N {
		info, err := authorizetypes.IdTagInfo(payload)
		if err != nil {
			b.Fatalf("unexpected error on create minimal: %v", err)
		}

		result = info
	}

	_ = result
}

func BenchmarkIdTagInfo_Create_Full(b *testing.B) {
	var result authorizetypes.IdTagInfoType

	payload := newFullPayload()

	for range b.N {
		info, err := authorizetypes.IdTagInfo(payload)
		if err != nil {
			b.Fatalf("unexpected error when create full: %v", err)
		}

		result = info
	}

	_ = result
}

func BenchmarkIdTagInfo_Value_Minimal(b *testing.B) {
	info, err := authorizetypes.IdTagInfo(newMinimalPayload())
	if err != nil {
		b.Fatalf("unexpected error with value minimal: %v", err)
	}

	var val authorizetypes.IdTagInfoPayload

	b.ResetTimer()

	for range b.N {
		val = info.Value()
	}

	_ = val
}

func BenchmarkIdTagInfo_Value_Full(b *testing.B) {
	info, err := authorizetypes.IdTagInfo(newFullPayload())
	if err != nil {
		b.Fatalf("unexpected error with value full: %v", err)
	}

	var val authorizetypes.IdTagInfoPayload

	b.ResetTimer()

	for range b.N {
		val = info.Value()
	}

	_ = val
}
