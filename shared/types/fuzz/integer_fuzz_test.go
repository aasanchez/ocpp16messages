package sharedtypes_test

import (
	"errors"
	"strconv"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

// FuzzSetInteger exercises SetInteger for a wide range of inputs,
// verifying error contracts and round-trip/idempotence properties.
func FuzzSetInteger(f *testing.F) {
	// Seed values: valid boundaries and invalid cases.
	f.Add("0")
	f.Add(strconv.FormatUint(uint64(^uint32(0)), 10)) // max uint32
	f.Add("1")
	f.Add("42")
	f.Add("")
	f.Add("-1")
	f.Add("4294967296") // overflow
	f.Add("not-a-number")

	f.Fuzz(func(t *testing.T, text string) {
		i, err := st.SetInteger(text)
		if err != nil {
			// Contract: invalid inputs produce a *strconv.NumError
			var ne *strconv.NumError
			if !errors.As(err, &ne) {
				t.Fatalf("expected NumError for %q, got: %v", text, err)
			}

			return
		}
		// Round-Trip: format back and parse again
		out := strconv.FormatUint(uint64(i.Value()), 10)
		i2, err := st.SetInteger(out)
		if err != nil {
			t.Fatalf("round-trip parse failed: %q -> %q: %v", text, out, err)
		}
		if i2.Value() != i.Value() {
			t.Fatalf("round-trip value mismatch: %q -> %q -> %q", text, i.Value(), i2.Value())
		}
	})
}
