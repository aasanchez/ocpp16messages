//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16messages/cancelReservation"
	ct "github.com/aasanchez/ocpp16messages/cancelReservation/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzCancelReservationConf(f *testing.F) {
	f.Add(ct.CancelReservationStatusAccepted.String())
	f.Add(ct.CancelReservationStatusRejected.String())
	f.Add("bad-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := cancelReservation.Conf(cancelReservation.ConfInput{
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
