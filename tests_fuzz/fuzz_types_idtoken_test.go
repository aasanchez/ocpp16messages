//go:build fuzz

package fuzz

import (
	"testing"

	types "github.com/aasanchez/ocpp16types"
)

func FuzzNewIdToken(f *testing.F) {
	f.Add("RFID-ABC123")
	f.Add("")
	f.Add("a")
	f.Add("12345678901234567890") // exactly 20
	f.Add("123456789012345678901") // 21 = too long
	f.Add("bad\x01char")

	f.Fuzz(func(t *testing.T, input string) {
		if len(input) > maxFuzzLen {
			t.Skip()
		}

		ciStr, err := types.NewCiString20Type(input)
		if err != nil {
			return
		}

		token := types.NewIdToken(ciStr)

		if got := token.Value().String(); got != input {
			t.Fatalf(
				"Value().String() = %q, want %q",
				got, input,
			)
		}

		if got := token.String(); got != input {
			t.Fatalf("String() = %q, want %q", got, input)
		}

		first := token.String()
		second := token.String()

		if first != second {
			t.Fatalf(
				"String() not deterministic: %q vs %q",
				first, second,
			)
		}
	})
}
