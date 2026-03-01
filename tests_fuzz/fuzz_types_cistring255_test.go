//go:build fuzz

package fuzz

import (
	"errors"
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzNewCiString255Type(f *testing.F) {
	f.Add("a")
	f.Add("")
	f.Add(" ")
	f.Add("\x00")
	f.Add(strings.Repeat("a", st.CiString255Max))
	f.Add(strings.Repeat("a", st.CiString255Max+1))

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		cis, err := st.NewCiString255Type(input)
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
			t.Fatal("NewCiString255Type succeeded with empty input")
		}
		if len(input) > st.CiString255Max {
			t.Fatalf("NewCiString255Type succeeded with len=%d", len(input))
		}

		for _, r := range input {
			if r < 32 || r > 126 {
				t.Fatalf(
					"NewCiString255Type succeeded with non-printable ASCII rune=%U",
					r,
				)
			}
		}

		if got := cis.Value(); got != input {
			t.Fatalf("Value = %q, want %q", got, input)
		}
	})
}
