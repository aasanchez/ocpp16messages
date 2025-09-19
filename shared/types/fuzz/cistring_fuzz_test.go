//go:build fuzz

package sharedtypes_test

import (
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func FuzzCiStringTypes(f *testing.F) {
	seedPrintable := func(n int) string { return strings.Repeat("A", n) }
	seedNonPrintable := func(n int, r rune) string {
		if n <= 0 {
			n = 1
		}
		runes := make([]rune, n)
		for i := range runes {
			runes[i] = 'A'
		}
		runes[n/2] = r
		return string(runes)
	}

	seeds := []struct {
		max int
		mk  func(string) (any, error)
	}{
		{20, func(s string) (any, error) { return st.SetCiString20Type(s) }},
		{25, func(s string) (any, error) { return st.SetCiString25Type(s) }},
		{50, func(s string) (any, error) { return st.SetCiString50Type(s) }},
		{255, func(s string) (any, error) { return st.SetCiString255Type(s) }},
		{500, func(s string) (any, error) { return st.SetCiString500Type(s) }},
	}

	for _, s := range seeds {
		f.Add(0, s.max, seedPrintable(1))
		f.Add(1, s.max, seedPrintable(s.max))
		f.Add(2, s.max, seedPrintable(s.max+1))
		f.Add(3, s.max, seedNonPrintable(s.max, rune(0x1F)))
		f.Add(4, s.max, seedNonPrintable(s.max, rune(0x7F)))
		f.Add(5, s.max, seedNonPrintable(s.max, rune(0xE9)))
	}

	f.Fuzz(func(t *testing.T, selector int, max int, s string) {
		// normalize selector
		if selector < 0 {
			selector = -selector
		}
		selector = selector % 5

		var (
			mk    func(string) (any, error)
			limit int
		)
		switch selector {
		case 0:
			mk = func(in string) (any, error) { return st.SetCiString20Type(in) }
			limit = 20
		case 1:
			mk = func(in string) (any, error) { return st.SetCiString25Type(in) }
			limit = 25
		case 2:
			mk = func(in string) (any, error) { return st.SetCiString50Type(in) }
			limit = 50
		case 3:
			mk = func(in string) (any, error) { return st.SetCiString255Type(in) }
			limit = 255
		case 4:
			mk = func(in string) (any, error) { return st.SetCiString500Type(in) }
			limit = 500
		}

		_, err := mk(s)
		validASCII := true
		for _, r := range s {
			if r < 32 || r > 126 {
				validASCII = false
				break
			}
		}
		withinLimit := len(s) <= limit

		if validASCII && withinLimit {
			if err != nil {
				t.Fatalf(
					"expected success for len=%d limit=%d printable, got err=%v",
					len(s),
					limit,
					err,
				)
			}
		} else {
			if err == nil {
				t.Fatalf("expected error for len=%d limit=%d printable=%v, got nil", len(s), limit, validASCII)
			}
		}
	})
}
