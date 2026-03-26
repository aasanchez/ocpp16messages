//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	scp "github.com/aasanchez/ocpp16messages/setchargingprofile"
	spt "github.com/aasanchez/ocpp16messages/setchargingprofile/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzSetChargingProfileConf(f *testing.F) {
	f.Add(spt.ChargingProfileStatusAccepted.String())
	f.Add(spt.ChargingProfileStatusRejected.String())
	f.Add(spt.ChargingProfileStatusNotSupported.String())
	f.Add("")
	f.Add("invalid-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := scp.Conf(scp.ConfInput{Status: status})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if !conf.Status.IsValid() {
			t.Fatalf("Status = %q, want valid", conf.Status.String())
		}
	})
}
