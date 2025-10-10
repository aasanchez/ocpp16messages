//go:build fuzz

// Package sharedtypes_test contains fuzz
// tests that verify RFC3339/RFC3339Nano
// parsing and formatting for OCPP 1.6
// timestamps via SetDateTime. It protects
// CSMS and charge points against malformed
// time payloads in messages like
// StartTransaction, MeterValues and
// StatusNotification.
//
// In OCPP 1.6, timestamps annotate
// transactions, connector status, and
// meter samples. Robust handling prevents
// time drift, overflow and injection that
// could corrupt billing or monitoring.
//
// Run with:
//
//	go test -tags fuzz -fuzz
//	FuzzSetDateTime ./...
package sharedtypes_test

import (
	"testing"
	"time"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// addSeedInputs seeds the corpus with
// representative valid and invalid
// RFC3339/RFC3339Nano forms seen in
// OCPP 1.6 traffic. This drives early
// coverage across zones, leap seconds,
// partial forms and unicode noise.
func addSeedInputs(f *testing.F) {
	seeds := []string{
		// Valid RFC3339
		"2025-09-19T12:34:56Z",
		"2025-09-19T12:34:56.789Z",
		"2025-09-19T12:34:56+02:00",
		"2025-09-19T12:34:56.789+02:00",
		"0001-01-01T00:00:00Z",      // min
		"9999-12-31T23:59:59Z",      // max
		"2016-12-31T23:59:60Z",      // leap second
		"1800-01-01T00:00:00Z",      // far past
		"3000-01-01T00:00:00Z",      // far future
		"2020-02-29T12:00:00Z",      // leap year
		"2021-03-14T02:30:00-07:00", // DST US
		// Invalid
		"not-a-date",
		"",
		"2025-13-01T00:00:00Z",           // invalid month
		"2025-00-01T00:00:00Z",           // invalid month
		"2025-12-32T00:00:00Z",           // invalid day
		"2025-12-00T00:00:00Z",           // invalid day
		"2025-12-01T24:00:00Z",           // invalid hour
		"2025-12-01T00:60:00Z",           // invalid minute
		"2025-12-01T00:00:60Z",           // invalid second
		"2025-12-01T00:00:00+99:99",      // invalid offset
		"2025-12-01T00:00:00.123456789Z", // too many decimals
		"2025-12-01T00:00:00.123456Z",    // high precision
		"2025-12-01T00:00:00.1Z",         // low precision
		"2025-12-01T00:00:00.000Z",       // zero ms
		"2025-12-01T00:00:00.000000000Z", // zero ns
		"2025-12-01T00:00:00.000000001Z", // 1ns
		"2025-12-01T00:00:00.999999999Z", // 999999999ns
		"2025-12-01T00:00:00.000000000+00:00",
		"2025-12-01T00:00:00.000000000-00:00",
		// Unicode, random, partial
		"日期:2025-09-19T12:34:56Z",
		"2025-09-19",
		"T12:34:56Z",
		"2025-09-19T12:34",
		"2025-09-19T12:34:56",
		"2025-09-19T12:34:56.789",
		"2025-09-19T12:34:56.789+",
		"2025-09-19T12:34:56.789-",
		"2025-09-19T12:34:56.789+00",
		"2025-09-19T12:34:56.789-00",
		"2025-09-19T12:34:56.789+00:",
		"2025-09-19T12:34:56.789-00:",
		"2025-09-19T12:34:56.789+00:0",
		"2025-09-19T12:34:56.789-00:0",
		"2025-09-19T12:34:56.789+00:00:00",
		"2025-09-19T12:34:56.789-00:00:00",
	}
	for _, s := range seeds {
		f.Add(s)
	}
}

// FuzzSetDateTime stresses st.SetDateTime
// with adversarial inputs to harden
// OCPP 1.6 timestamp handling. It checks:
// - no panics for any input
// - String() round-trips the value
// - output matches RFC3339Nano
// - nanosecond precision is kept
// - location normalizes to UTC
// These invariants protect transaction and
// meter time lines in CSMS.
func FuzzSetDateTime(f *testing.F) {
	addSeedInputs(f)
	f.Fuzz(func(t *testing.T, input string) {
		dt, err := st.SetDateTime(input)
		if err != nil {
			// For invalid input, ensure no panic and error is expected
			return
		}
		// If parsing succeeded, check round-trip
		dtStr := dt.String()
		dt2, err2 := st.SetDateTime(dtStr)
		if err2 != nil {
			t.Errorf("round-trip parse failed: input=%q, dtStr=%q, err2=%v", input, dtStr, err2)
		}
		if !dt.Value().Equal(dt2.Value()) {
			t.Errorf(
				"round-trip value mismatch: input=%q, dtStr=%q, got=%v, want=%v",
				input,
				dtStr,
				dt2.Value(),
				dt.Value(),
			)
		}
		// Check that String() always produces a valid RFC3339Nano string
		if _, err := time.Parse(time.RFC3339Nano, dtStr); err != nil {
			t.Errorf("String() output not RFC3339Nano: %q, err=%v", dtStr, err)
		}
		// Check nanosecond precision is preserved
		if dt.Value().Nanosecond() != dt2.Value().Nanosecond() {
			t.Errorf(
				"nanosecond precision lost: input=%q, got=%d, want=%d",
				input,
				dt2.Value().Nanosecond(),
				dt.Value().Nanosecond(),
			)
		}
		// Check time zone normalization: both should be UTC after round-trip
		if dt2.Value().Location().String() != "UTC" {
			t.Errorf("round-trip location not UTC: input=%q, got=%v", input, dt2.Value().Location())
		}
	})
}
