package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

func BenchmarkRequestPayload_Validate_valid(b *testing.B) {
	payload := authorizetypes.RequestPayload{IdTag: "B85A-50CBE9678EC6"}
	for range b.N {
		if err := payload.Validate(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRequestPayload_Validate_empty(b *testing.B) {
	payload := authorizetypes.RequestPayload{IdTag: ""}
	for range b.N {
		_ = payload.Validate()
	}
}

func BenchmarkRequestPayload_Value(b *testing.B) {
	payload := authorizetypes.RequestPayload{IdTag: "B85A-50CBE9678EC6"}

	var s string
	for range b.N {
		s = payload.Value()
	}

	_ = s
}
