//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"strconv"
	"testing"

	lt "github.com/aasanchez/ocpp16messages/getlocallistversion/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzNewListVersionNumber(f *testing.F) {
	f.Add(lt.ListVersionUnsupported)
	f.Add(lt.ListVersionEmpty)
	f.Add(int(math.MinInt32))
	f.Add(int(math.MaxInt32))
	f.Add(int(math.MinInt32) - 1)
	f.Add(int(math.MaxInt32) + 1)

	f.Fuzz(func(t *testing.T, value int) {
		listVersion, err := lt.NewListVersionNumber(value)
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if value < math.MinInt32 || value > math.MaxInt32 {
			t.Fatalf("NewListVersionNumber succeeded with value=%d", value)
		}

		if got := listVersion.Value(); got != int32(value) {
			t.Fatalf("Value() = %d, want %d", got, value)
		}

		if listVersion.IsUnsupported() != (listVersion.Value() == lt.ListVersionUnsupported) {
			t.Fatal("IsUnsupported() mismatch")
		}

		if listVersion.IsEmpty() != (listVersion.Value() == lt.ListVersionEmpty) {
			t.Fatal("IsEmpty() mismatch")
		}

		parsed, parseErr := strconv.ParseInt(listVersion.String(), 10, 32)
		if parseErr != nil {
			t.Fatalf("String() not parseable: %v", parseErr)
		}

		if int32(parsed) != listVersion.Value() {
			t.Fatalf("String() parsed = %d, want %d", parsed, listVersion.Value())
		}
	})
}
