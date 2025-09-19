//go:build benchmark

package sharedtypes_test

import (
	st "github.com/aasanchez/ocpp16messages/shared/types"
	"testing"
)

func BenchmarkSetDateTime(b *testing.B) {
	cases := []struct {
		name  string
		input string
	}{
		{"valid_now", "2025-09-19T12:34:56Z"},
		{"valid_fractional", "2025-09-19T12:34:56.789Z"},
		{"invalid_format", "not-a-date"},
		{"edge_min", "0001-01-01T00:00:00Z"},
		{"edge_max", "9999-12-31T23:59:59Z"},
		{"leap_second", "2016-12-31T23:59:60Z"},
		{"far_past", "1800-01-01T00:00:00Z"},
		{"far_future", "3000-01-01T00:00:00Z"},
	}
	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				_, _ = st.SetDateTime(tc.input)
			}
		})
	}
}

func BenchmarkDateTimeString(b *testing.B) {
	now, _ := st.SetDateTime("2025-09-19T12:34:56Z")
	min, _ := st.SetDateTime("0001-01-01T00:00:00Z")
	max, _ := st.SetDateTime("9999-12-31T23:59:59Z")
	frac, _ := st.SetDateTime("2025-09-19T12:34:56.789Z")

	cases := []struct {
		name string
		dt   st.DateTime
	}{
		{"now", now},
		{"min", min},
		{"max", max},
		{"fractional", frac},
	}
	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				_ = tc.dt.String()
			}
		})
	}
}

func BenchmarkDateTimeValue(b *testing.B) {
	dt, _ := st.SetDateTime("2025-09-19T12:34:56Z")
	b.ReportAllocs()
	for b.Loop() {
		_ = dt.Value()
	}
}

func BenchmarkSetDateTimeParallel(b *testing.B) {
	input := "2025-09-19T12:34:56Z"
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = st.SetDateTime(input)
		}
	})
}
