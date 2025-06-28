package changeavailabilitytypes_test

import (
	"testing"

	cat "github.com/aasanchez/ocpp16messages/messages/changeavailability/types"
)

func BenchmarkAvailabilityType_Create(b *testing.B) {
	input := cat.AvailabilityTypeOperative
	var result cat.AvailabilityTypeType
	b.ResetTimer()
	for range b.N {
		r, err := cat.AvailabilityType(input)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
		result = r
	}
	_ = result
}

func BenchmarkAvailabilityType_Value(b *testing.B) {
	status, _ := cat.AvailabilityType(cat.AvailabilityTypeOperative)
	var result string
	b.ResetTimer()
	for range b.N {
		result = status.Value()
	}
	_ = result
}
