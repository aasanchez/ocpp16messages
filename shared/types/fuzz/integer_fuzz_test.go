//go:build fuzz

// Package sharedtypes_test provides fuzz tests for
// integer parsing used by OCPP 1.6 message types.
// It explains why strict integer handling matters
// to charge point–central system interoperability
// and validates those rules against wide inputs.
//
// Role of fuzzing in OCPP 1.6 validation.
// Fuzzing explores diverse, often unexpected text
// inputs that may arrive over the wire. It probes
// parser edges to confirm that numeric constraints
// are enforced consistently. In OCPP 1.6, this is
// essential for message integrity and robustness.
//
// Why integers are critical in OCPP 1.6.
// Many fields are small non‑negative integers:
// connector identifiers, transaction identifiers,
// meter counters, sequence numbers, retry counts,
// and sample sizes. These must be decimal, bounded,
// and unambiguous to avoid truncation or overflow.
//
// How these tests harden reliability.
// The fuzz tests compare SetInteger with
// strconv.ParseUint using base 10 and 16 bits.
// Valid decimal inputs must parse identically.
// Invalid forms—negatives, prefixes, floats,
// whitespace‑only, and overflow—must be rejected.
// This symmetry prevents malformed or hostile
// payloads from bypassing validation in either
// charge points or the central system.
//
// Running
//
//	go test -fuzz=Fuzz -run=^$ ./shared/types/fuzz
//
// Seeds emphasize boundaries and common mistakes to
// accelerate discovery while the fuzzer explores.
package sharedtypes_test

import (
	"strconv"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// FuzzSetInteger exercises SetInteger with
// generated strings and checks behavior against
// strconv.ParseUint(s, 10, 16).
//
// Expected valid behavior.
// Decimal values in the range 0..65535 must be
// accepted, and the returned value must equal the
// uint16 produced by strconv.
//
// Expected invalid behavior.
// Inputs that are negative, overflow the 16‑bit
// range, contain signs or prefixes, include decimal
// points, or are empty/whitespace must be rejected.
// The comparison requires strconv to reject them as
// well, ensuring aligned semantics.
//
// OCPP 1.6 interpretation.
// Accepting only canonical decimal digits within
// bounds ensures that identifiers and counters are
// stable across implementations. Peers cannot coerce
// acceptance via hexadecimal notation, leading/trailing
// whitespace, sign tricks, or overflow effects.
// This strengthens message validation and reduces the
// risk of malformed payloads passing unchecked.
func FuzzSetInteger(f *testing.F) {
	// Seeds cover valid edges and invalid forms.
	seeds := []string{
		"0", "1", "9", "10", "99", "100", "1024", "65535",
		"", " ", "-1", "65536", "999999999999999", "1.5", "abc", "+12", "0x10",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	// Compare SetInteger with strconv to ensure
	// aligned behavior for decimal uint16.
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

		// If our parser rejects but strconv accepts,
		// that indicates a mismatch in semantics.
		if perr == nil {
			t.Errorf("SetInteger rejected %q, but strconv.ParseUint succeeded (=%d)", s, parsed)
		}
	})
}
