//go:build fuzz

package fuzz

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/clearCache"
)

func FuzzClearCacheReq(f *testing.F) {
	f.Add(uint8(0))

	f.Fuzz(func(t *testing.T, _ uint8) {
		_, err := clearCache.Req(clearCache.ReqInput{})
		if err != nil {
			t.Fatalf("error = %v, want nil", err)
		}
	})
}
