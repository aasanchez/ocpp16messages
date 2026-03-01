//go:build fuzz

package fuzz

import (
	"errors"
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzNewCiString500Type(f *testing.F) {
	f.Add("a")
	f.Add("")
	f.Add(" ")
	f.Add("\x00")
	f.Add(strings.Repeat("a", st.CiString500Max))
	f.Add(strings.Repeat("a", st.CiString500Max+1))

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		cis, err := st.NewCiString500Type(input)
		if err != nil {
			if !errors.Is(err, st.ErrEmptyValue) && !errors.Is(err, st.ErrInvalidValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if input == "" {
			t.Fatal("NewCiString500Type succeeded with empty input")
		}
		if len(input) > st.CiString500Max {
			t.Fatalf("NewCiString500Type succeeded with len=%d", len(input))
		}

		for _, r := range input {
			if r < 32 || r > 126 {
				t.Fatalf(
					"NewCiString500Type succeeded with non-printable ASCII rune=%U",
					r,
				)
			}
		}

		if got := cis.Value(); got != input {
			t.Fatalf("Value = %q, want %q", got, input)
		}
	})
}
