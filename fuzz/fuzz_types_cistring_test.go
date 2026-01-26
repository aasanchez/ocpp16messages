//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"unicode/utf8"

	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzNewCiString20Type(f *testing.F) {
	f.Add("RFID-ABC123")
	f.Add("")
	f.Add("toolongstringtoolong")
	f.Add("bad\x01")

	f.Fuzz(func(t *testing.T, input string) {
		value, err := st.NewCiString20Type(input)
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) && !errors.Is(err, st.ErrEmptyValue) {
				t.Fatalf("unexpected error sentinel for %q: %v", input, err)
			}
			return
		}

		if utf8.RuneCountInString(value.Value()) > st.CiString20Max {
			t.Fatalf("len(%q) exceeded max", value.Value())
		}

		for _, r := range value.Value() {
			if r < 32 || r > 126 {
				t.Fatalf("non-printable rune %q in %q", r, value.Value())
			}
		}
	})
}
