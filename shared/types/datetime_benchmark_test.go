package types

import (
	"testing"
)

const rfc3339Input = "2025-06-19T12:34:56Z"

func BenchmarkNewDateTime_Valid(b *testing.B) {
	for range b.N {
		_, err := DateTime(rfc3339Input)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkNewDateTime_InvalidFormat(b *testing.B) {
	invalid := "not-a-valid-time"

	for range b.N {
		_, _ = DateTime(invalid) // we're measuring performance, not correctness
	}
}

func BenchmarkNewDateTime_ZeroTime(b *testing.B) {
	zero := "0001-01-01T00:00:00Z" // will parse to zero time

	for range b.N {
		_, _ = DateTime(zero)
	}
}

func BenchmarkDateTime_Value(b *testing.B) {
	time, err := DateTime(rfc3339Input)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()

	for range b.N {
		_ = time.Value()
	}
}
