//go:build fuzz

package fuzz

import (
	"errors"
	"strings"
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzNewCiString25Type(f *testing.F) {
	f.Add("a")
	f.Add("")
	f.Add(" ")
	f.Add("\x00")
	f.Add(strings.Repeat("a", types.CiString25Max))
	f.Add(strings.Repeat("a", types.CiString25Max+1))

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzStringLen {
			t.Skip()
		}

		cis, err := types.NewCiString25Type(input)
		if err != nil {
			if !errors.Is(err, types.ErrEmptyValue) && !errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if input == "" {
			t.Fatal("NewCiString25Type succeeded with empty input")
		}
		if len(input) > types.CiString25Max {
			t.Fatalf("NewCiString25Type succeeded with len=%d", len(input))
		}

		for _, r := range input {
			if r < 32 || r > 126 {
				t.Fatalf("NewCiString25Type succeeded with non-printable ASCII rune=%U", r)
			}
		}

		if got := cis.Value(); got != input {
			t.Fatalf("Value = %q, want %q", got, input)
		}
	})
}
