package types

import (
	"testing"
)

const rfc3339Input = "2025-06-19T12:34:56Z"

func BenchmarkNewDateTime_Valid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := DateTime(rfc3339Input)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkNewDateTime_InvalidFormat(b *testing.B) {
	invalid := "not-a-valid-time"
	for i := 0; i < b.N; i++ {
		_, _ = DateTime(invalid) // we're measuring performance, not correctness
	}
}

func BenchmarkNewDateTime_ZeroTime(b *testing.B) {
	zero := "0001-01-01T00:00:00Z" // will parse to zero time
	for i := 0; i < b.N; i++ {
		_, _ = DateTime(zero)
	}
}

func BenchmarkDateTime_Value(b *testing.B) {
	dt, err := DateTime(rfc3339Input)
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dt.Value()
	}
}
