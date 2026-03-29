//go:build fuzz

package fuzz

import (
	"testing"
	"time"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzNewDateTime(f *testing.F) {
	f.Add("2025-01-15T10:30:00Z")
	f.Add("not-a-time")
	f.Add("")
	f.Add("2025-01-15T10:30:00+02:00")

	f.Fuzz(func(t *testing.T, input string) {
		dt, err := types.NewDateTime(input)
		if err != nil {
			return
		}

		if dt.Value().Location() != time.UTC {
			t.Fatalf("location = %v, want UTC", dt.Value().Location())
		}

		roundTrip := dt.String()
		if roundTrip == "" {
			t.Fatalf("String() returned empty for input %q", input)
		}

		if _, parseErr := time.Parse(time.RFC3339Nano, roundTrip); parseErr != nil {
			t.Fatalf("String() not RFC3339Nano: %v", parseErr)
		}

		// String() determinism
		if dt.String() != roundTrip {
			t.Fatalf(
				"String() not deterministic: %q vs %q",
				dt.String(), roundTrip,
			)
		}

		// Round-trip closure
		dt2, rtErr := types.NewDateTime(roundTrip)
		if rtErr != nil {
			t.Fatalf(
				"round-trip NewDateTime(%q) = %v",
				roundTrip, rtErr,
			)
		}

		if !dt2.Value().Equal(dt.Value()) {
			t.Fatalf(
				"round-trip mismatch: %v vs %v",
				dt2.Value(), dt.Value(),
			)
		}
	})
}
