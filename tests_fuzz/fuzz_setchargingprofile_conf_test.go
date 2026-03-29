//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	scp "github.com/aasanchez/ocpp16messages/setchargingprofile"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzSetChargingProfileConf(f *testing.F) {
	f.Add(types.ChargingProfileStatusAccepted.String())
	f.Add(types.ChargingProfileStatusRejected.String())
	f.Add(types.ChargingProfileStatusNotSupported.String())
	f.Add("")
	f.Add("invalid-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := scp.Conf(scp.ConfInput{Status: status})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if !conf.Status.IsValid() {
			t.Fatalf("Status = %q, want valid", conf.Status.String())
		}
	})
}
