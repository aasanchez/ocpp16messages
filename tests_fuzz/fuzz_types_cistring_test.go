//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"unicode/utf8"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzNewCiString20Type(f *testing.F) {
	f.Add("RFID-ABC123")
	f.Add("")
	f.Add("toolongstringtoolong")
	f.Add("bad\x01")

	f.Fuzz(func(t *testing.T, input string) {
		value, err := types.NewCiString20Type(input)
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) && !errors.Is(err, types.ErrEmptyValue) {
				t.Fatalf("unexpected error sentinel for %q: %v", input, err)
			}
			return
		}

		if utf8.RuneCountInString(value.Value()) > types.CiString20Max {
			t.Fatalf("len(%q) exceeded max", value.Value())
		}

		for _, r := range value.Value() {
			if r < 32 || r > 126 {
				t.Fatalf("non-printable rune %q in %q", r, value.Value())
			}
		}
	})
}
