package sharedtypes_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func BenchmarkSetInteger_ValidInput(b *testing.B) {
	for range b.N {
		_, err := st.SetInteger("42")
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func BenchmarkInteger_Value(b *testing.B) {
	integer, err := st.SetInteger("42")
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()

	for range b.N {
		_ = integer.Value()
	}
}
