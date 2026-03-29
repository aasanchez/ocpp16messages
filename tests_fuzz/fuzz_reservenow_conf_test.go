//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	rn "github.com/aasanchez/ocpp16messages/reservenow"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzReserveNowConf(f *testing.F) {
	f.Add(types.ReservationStatusAccepted.String())
	f.Add(types.ReservationStatusFaulted.String())
	f.Add(types.ReservationStatusOccupied.String())
	f.Add(types.ReservationStatusRejected.String())
	f.Add(types.ReservationStatusUnavailable.String())
	f.Add("")
	f.Add("invalid-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := rn.Conf(rn.ConfInput{Status: status})
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
