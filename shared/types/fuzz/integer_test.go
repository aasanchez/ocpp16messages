package sharedtypes_fuzz

import (
	"math"
	"testing"

	st "github.com/aasanchez/ocpp16messages/shared/types"
)

func FuzzSetInteger(f *testing.F) {
	f.Add("42")
	f.Add("0")
	f.Add("4294967294")
	f.Add("-1")
	f.Add("abc")
	f.Add("")

	f.Fuzz(func(t *testing.T, input string) {
		integer, err := st.SetInteger(input)

		if err == nil {
			v := integer.Value()
			if v >= math.MaxUint32 -1 {
				t.Errorf("Value() returned %d, which exceeds uint32 max", v)
			}
		}
	})
}
