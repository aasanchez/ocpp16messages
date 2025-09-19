//go:build fuzz

package sharedtypes_test

import (
	"strconv"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func FuzzSetInteger(f *testing.F) {
	seeds := []string{
		"0", "1", "9", "10", "99", "100", "1024", "65535",
		"", " ", "-1", "65536", "999999999999999", "1.5", "abc", "+12", "0x10",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, s string) {
		got, err := st.SetInteger(s)
		parsed, perr := strconv.ParseUint(s, 10, 16)

		if err == nil {
			if perr != nil {
				t.Fatalf("SetInteger accepted input %q but strconv.ParseUint failed: %v", s, perr)
			}
			if uint16(parsed) != got.Value() {
				t.Fatalf("value mismatch for %q: got=%d want=%d", s, got.Value(), uint16(parsed))
			}
			return
		}

		// err != nil: ensure strconv indicates an error too
		if perr == nil {
			t.Fatalf("SetInteger rejected %q, but strconv.ParseUint succeeded (=%d)", s, parsed)
		}
	})
}
