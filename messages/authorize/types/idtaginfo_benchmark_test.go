package authorizetypes_test

import (
	"testing"

	authorizetypes "github.com/aasanchez/ocpp16messages/messages/authorize/types"
)

var (
	benchmarkStatus       = authorizetypes.Accepted
	benchmarkExpiryString = "2027-04-12T14:03:04Z"
	benchmarkParentIdTag  = "ABC123456789"
)

var (
	benchmarkPayloadFull = authorizetypes.IdTagInfoPayload{
		Status:      benchmarkStatus,
		ExpiryDate:  &benchmarkExpiryString,
		ParentIdTag: &benchmarkParentIdTag,
	}

	benchmarkPayloadMinimal = authorizetypes.IdTagInfoPayload{
		Status: benchmarkStatus,
	}

	sinkIdTagInfo authorizetypes.IdTagInfoType
	sinkValue     authorizetypes.IdTagInfoValue
)

func BenchmarkIdTagInfo_Create_Minimal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info, err := authorizetypes.IdTagInfo(benchmarkPayloadMinimal)
		if err != nil {
			b.Fatalf("unexpected error on create minamal: %v", err)
		}
		sinkIdTagInfo = info
	}
}

func BenchmarkIdTagInfo_Create_Full(b *testing.B) {
	for i := 0; i < b.N; i++ {
		info, err := authorizetypes.IdTagInfo(benchmarkPayloadFull)
		if err != nil {
			b.Fatalf("unexpected error when create full: %v", err)
		}
		sinkIdTagInfo = info
	}
}

func BenchmarkIdTagInfo_Value_Minimal(b *testing.B) {
	info, err := authorizetypes.IdTagInfo(benchmarkPayloadMinimal)
	if err != nil {
		b.Fatalf("unexpected error with value minimal: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkValue = info.Value()
	}
}

func BenchmarkIdTagInfo_Value_Full(b *testing.B) {
	info, err := authorizetypes.IdTagInfo(benchmarkPayloadFull)
	if err != nil {
		b.Fatalf("unexpected error with value full: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkValue = info.Value()
	}
}
