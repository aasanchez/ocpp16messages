//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"strconv"
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzNewInteger(f *testing.F) {
	f.Add(0)
	f.Add(1)
	f.Add(-1)
	f.Add(int(math.MaxUint16))
	f.Add(int(math.MaxUint16) + 1)
	f.Add(math.MaxInt)
	f.Add(math.MinInt)

	f.Fuzz(func(t *testing.T, value int) {
		integer, err := st.NewInteger(value)
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if value < 0 || value > math.MaxUint16 {
			t.Fatalf("NewInteger succeeded with value=%d", value)
		}

		if got := integer.Value(); got != uint16(value) {
			t.Fatalf("Value() = %d, want %d", got, value)
		}

		parsed, parseErr := strconv.ParseUint(integer.String(), 10, 16)
		if parseErr != nil {
			t.Fatalf("String() not parseable: %v", parseErr)
		}

		if uint16(parsed) != integer.Value() {
			t.Fatalf("String() parsed = %d, want %d", parsed, integer.Value())
		}
	})
}
