package types_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/shared/types"
)

const rfc3339Input = "2025-06-19T12:34:56Z"

func BenchmarkNewDateTime_Valid(b *testing.B) {
	for range b.N {
		_, err := types.DateTime(rfc3339Input)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkNewDateTime_InvalidFormat(b *testing.B) {
	invalid := "not-a-valid-time"

	for range b.N {
		_, _ = types.DateTime(invalid)
	}
}

func BenchmarkNewDateTime_ZeroTime(b *testing.B) {
	zero := "0001-01-01T00:00:00Z"

	for range b.N {
		_, _ = types.DateTime(zero)
	}
}

func BenchmarkDateTime_Value(b *testing.B) {
	time, err := types.DateTime(rfc3339Input)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()

	for range b.N {
		_ = time.Value()
	}
}
