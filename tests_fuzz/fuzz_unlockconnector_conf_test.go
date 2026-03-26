//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	st "github.com/aasanchez/ocpp16messages/types"
	"github.com/aasanchez/ocpp16messages/unlockconnector"
	ut "github.com/aasanchez/ocpp16messages/unlockconnector/types"
)

func FuzzUnlockConnectorConf(f *testing.F) {
	f.Add(ut.UnlockStatusUnlocked.String())
	f.Add(ut.UnlockStatusUnlockFailed.String())
	f.Add(ut.UnlockStatusNotSupported.String())
	f.Add("bad-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := unlockconnector.Conf(unlockconnector.ConfInput{
			Status: status,
		})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if !conf.Status.IsValid() {
			t.Fatalf("Status = %q, want valid", conf.Status.String())
		}
		if conf.Status.String() != status {
			t.Fatalf("Status = %q, want %q", conf.Status.String(), status)
		}
	})
}
