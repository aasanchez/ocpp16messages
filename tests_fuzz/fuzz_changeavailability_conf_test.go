//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16messages/changeavailability"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzChangeAvailabilityConf(f *testing.F) {
	f.Add(types.AvailabilityStatusAccepted.String())
	f.Add(types.AvailabilityStatusRejected.String())
	f.Add(types.AvailabilityStatusScheduled.String())
	f.Add("bad-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := changeavailability.Conf(changeavailability.ConfInput{
			Status: status,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) {
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
