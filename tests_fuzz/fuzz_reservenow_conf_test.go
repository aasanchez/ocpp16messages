//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	rn "github.com/aasanchez/ocpp16messages/reservenow"
	rt "github.com/aasanchez/ocpp16messages/reservenow/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzReserveNowConf(f *testing.F) {
	f.Add(rt.ReservationStatusAccepted.String())
	f.Add(rt.ReservationStatusFaulted.String())
	f.Add(rt.ReservationStatusOccupied.String())
	f.Add(rt.ReservationStatusRejected.String())
	f.Add(rt.ReservationStatusUnavailable.String())
	f.Add("")
	f.Add("invalid-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := rn.Conf(rn.ConfInput{Status: status})
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
