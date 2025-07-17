package sharedtypes_test

import (
   "testing"

   st "github.com/aasanchez/ocpp16messages/shared/types"
)

// BenchmarkSetInteger measures the performance of SetInteger parsing from string.
func BenchmarkSetInteger(b *testing.B) {
	var err error
	for range b.N {
		_, err = st.SetInteger("1234567890")
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

// BenchmarkInteger_Value measures the performance of retrieving the underlying value.
func BenchmarkInteger_Value(b *testing.B) {
	i, err := st.SetInteger("1234567890")
	if err != nil {
		b.Fatalf("failed to create Integer: %v", err)
	}

	b.ResetTimer()
	for range b.N {
		_ = i.Value()
	}
}
