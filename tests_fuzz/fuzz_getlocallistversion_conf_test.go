//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	"github.com/aasanchez/ocpp16messages/getLocalListVersion"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzGetLocalListVersionConf(f *testing.F) {
	f.Add(-1)
	f.Add(0)
	f.Add(1)
	f.Add(math.MinInt32)
	f.Add(math.MaxInt32)
	f.Add(math.MinInt32 - 1)
	f.Add(math.MaxInt32 + 1)

	f.Fuzz(func(t *testing.T, listVersion int) {
		conf, err := getLocalListVersion.Conf(getLocalListVersion.ConfInput{
			ListVersion: listVersion,
		})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if listVersion < math.MinInt32 || listVersion > math.MaxInt32 {
			t.Fatalf("Conf succeeded with listVersion=%d", listVersion)
		}

		if got := conf.ListVersion.Value(); got != int32(listVersion) {
			t.Fatalf("ListVersion = %d, want %d", got, listVersion)
		}
	})
}
