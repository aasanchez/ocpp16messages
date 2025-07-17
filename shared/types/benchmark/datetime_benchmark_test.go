package sharedtypes_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// BenchmarkSetDateTime measures the cost of parsing an ISO8601 datetime string.
func BenchmarkSetDateTime(b *testing.B) {
	for range b.N {
		_, err := st.SetDateTime("2013-02-01T20:00:00.000Z")
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

// BenchmarkDateTime_String measures the cost of formatting a DateTime to string.
func BenchmarkDateTime_String(b *testing.B) {
	dt, err := st.SetDateTime("2013-02-01T20:00:00.000Z")
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()
	for range b.N {
		_ = dt.String()
	}
}

// BenchmarkDateTime_Value measures the cost of accessing the underlying time.Time value.
func BenchmarkDateTime_Value(b *testing.B) {
	dt, err := st.SetDateTime("2013-02-01T20:00:00.000Z")
	if err != nil {
		b.Fatalf("unexpected error: %v", err)
	}

	b.ResetTimer()
	for range b.N {
		_ = dt.Value()
	}
}
