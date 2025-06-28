package changeavailabilitytypes_test

import (
	"testing"

	changeavailabilitytypes "github.com/aasanchez/ocpp16messages/messages/changeavailability/types"
)

func BenchmarkAvailabilityType_Create(b *testing.B) {
	input := changeavailabilitytypes.AvailabilityTypeOperative
	var result changeavailabilitytypes.AvailabilityTypeType
	b.ResetTimer()
	for range b.N {
		r, err := changeavailabilitytypes.AvailabilityType(input)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
		result = r
	}
	_ = result
}

func BenchmarkAvailabilityType_Value(b *testing.B) {
	status, _ := changeavailabilitytypes.AvailabilityType(changeavailabilitytypes.AvailabilityTypeOperative)
	var result string
	b.ResetTimer()
	for range b.N {
		result = status.Value()
	}
	_ = result
}
