//go:build fuzz

// Package sharedtypes_test holds fuzz
// tests for Integer parsing in the
// OCPP 1.6 message domain.
// It validates SetInteger against
// strconv rules and uint16 limits.
// This helps keep numeric fields
// in OCPP 1.6 safe and strict.
// Usage (Go 1.24+ fuzzing):
//
//	go test -fuzz=Fuzz -run=^$ \
//	./shared/types/fuzz
//
// Seeds are chosen to hit edges,
// invalid forms, and large values.
package sharedtypes_test

import (
	"strconv"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// FuzzSetInteger checks that SetInteger
// accepts only valid uint16 decimal
// inputs and matches strconv limits.
// This guards OCPP 1.6 fields that
// use small non-negative integers,
// like meter counters or seq ids.
func FuzzSetInteger(f *testing.F) {
	// Seeds cover valid edges and
	// diverse invalid forms.
	seeds := []string{
		"0", "1", "9", "10", "99", "100", "1024", "65535",
		"", " ", "-1", "65536", "999999999999999", "1.5", "abc", "+12", "0x10",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	// Compare SetInteger with strconv
	// to ensure aligned behavior.
	f.Fuzz(func(t *testing.T, s string) {
		got, err := st.SetInteger(s)
		parsed, perr := strconv.ParseUint(s, 10, 16)

		if err == nil {
			if perr != nil {
				t.Errorf("SetInteger accepted input %q but strconv.ParseUint failed: %v", s, perr)
			}
			if uint16(parsed) != got.Value() {
				t.Errorf("value mismatch for %q: got=%d want=%d", s, got.Value(), uint16(parsed))
			}
			return
		}

		// If our parser rejects but strconv
		// accepts, that is a mismatch.
		if perr == nil {
			t.Errorf("SetInteger rejected %q, but strconv.ParseUint succeeded (=%d)", s, parsed)
		}
	})
}
