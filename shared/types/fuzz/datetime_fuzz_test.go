package sharedtypes_fuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// FuzzSetDateTime ensures SetDateTime can handle a wide range of string inputs
// without crashing or producing unexpected behavior.
func FuzzSetDateTime(f *testing.F) {
	// Add seed values: typical valid cases, boundary dates, and malformed inputs.
	f.Add("2013-02-01T20:00:00.000Z")
	f.Add("2006-01-02T15:04:05Z")
	f.Add("2006-01-02T15:04:05.999Z")
	f.Add("2006-01-02T15:04:05.999999999Z")
	f.Add("2006-01-02T15:04:05+07:00")
	f.Add("2006-01-02T15:04:05.123+02:00")
	f.Add("1999-12-31T23:59:59Z")
	f.Add("0001-01-01T00:00:00Z")
	f.Add("")
	f.Add("invalid-datetime")
	f.Add("2020-01-32T25:61:61Z")

	// Run the fuzzer.
	f.Fuzz(func(t *testing.T, data string) {
		// Try to parse; if successful, perform a round-trip consistency check.
		dt, err := st.SetDateTime(data)
		if err != nil {
			// invalid input is fine; we only want to catch panics or crashes
			return
		}

		// Round-trip: format back to string and parse again
		s := dt.String()

		dt2, err := st.SetDateTime(s)
		if err != nil {
			t.Fatalf("round-trip parse failed: %q -> %q: %v", data, s, err)
		}

		if !dt.Value().Equal(dt2.Value()) {
			t.Fatalf("round-trip value mismatch: %q -> %q -> %q", data, dt.String(), dt2.String())
		}
	})
}
