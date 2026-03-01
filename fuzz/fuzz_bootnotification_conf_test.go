//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	bn "github.com/aasanchez/ocpp16messages/bootNotification"
	bnt "github.com/aasanchez/ocpp16messages/bootNotification/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzBootNotificationConf(f *testing.F) {
	f.Add(bnt.RegistrationStatusAccepted.String(), "2025-01-15T10:30:00Z", 60)
	f.Add("invalid-status", "2025-01-15T10:30:00Z", 60)
	f.Add(bnt.RegistrationStatusAccepted.String(), "bad-time", 60)
	f.Add(bnt.RegistrationStatusAccepted.String(), "2025-01-15T10:30:00Z", -1)

	f.Fuzz(func(t *testing.T, status string, currentTime string, interval int) {
		if len(status) > maxFuzzStringLen || len(currentTime) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := bn.Conf(bn.ConfInput{
			Status:      status,
			CurrentTime: currentTime,
			Interval:    interval,
		})
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) && !errors.Is(err, st.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if !conf.Status.IsValid() {
			t.Fatalf("Status = %q, want valid", conf.Status.String())
		}

		if conf.CurrentTime.Value().Location() != time.UTC {
			t.Fatalf(
				"CurrentTime location = %v, want UTC",
				conf.CurrentTime.Value().Location(),
			)
		}

		if interval < 0 || interval > math.MaxUint16 {
			t.Fatalf("Conf succeeded with interval=%d", interval)
		}

		if got := conf.Interval.Value(); got != uint16(interval) {
			t.Fatalf("Interval = %d, want %d", got, interval)
		}
	})
}
