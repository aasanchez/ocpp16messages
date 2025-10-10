//go:build benchmark

// Package sharedtypes_test provides micro-
// benchmarks for DateTime used in OCPP 1.6
// payloads. It measures the cost to parse,
// format and access timestamps that appear
// in StartTransaction, MeterValues and
// StatusNotification between charge points
// and the CSMS.
//
// These measurements help keep message paths
// predictable under edge values (min/max,
// fractional seconds, leap seconds) and
// malformed inputs, supporting reliable
// billing, auditing and monitoring.
//
// Run with:
//
//	go test -tags benchmark -bench . -benchmem
package sharedtypes_test

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// BenchmarkSetDateTime measures parse and
// allocation cost across representative
// valid and invalid inputs (now, fractional,
// edges and malformed). It reflects inbound
// OCPP processing when decoding payloads.
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

// BenchmarkDateTimeString measures the cost to
// format RFC3339Nano strings from DateTime.
// It reflects outbound OCPP serialization for
// messages that carry timestamps.
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

// BenchmarkDateTimeValue measures accessor
// overhead to retrieve the underlying time
// value. Should be O(1) and allocation free.
func BenchmarkDateTimeValue(b *testing.B) {
	dt, _ := st.SetDateTime("2025-09-19T12:34:56Z")
	b.ReportAllocs()
	for b.Loop() {
		_ = dt.Value()
	}
}

// BenchmarkSetDateTimeParallel measures parse
// behavior under concurrency, mirroring CP or
// CSMS handlers that decode timestamps on many
// goroutines.
func BenchmarkSetDateTimeParallel(b *testing.B) {
	input := "2025-09-19T12:34:56Z"
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = st.SetDateTime(input)
		}
	})
}
