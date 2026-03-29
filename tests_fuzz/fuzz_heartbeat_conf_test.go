//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"time"

	"github.com/aasanchez/ocpp16messages/heartbeat"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzHeartbeatConf(f *testing.F) {
	f.Add("2025-01-15T10:30:00Z")
	f.Add("")
	f.Add("bad-time")

	f.Fuzz(func(t *testing.T, currentTime string) {
		if len(currentTime) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := heartbeat.Conf(heartbeat.ConfInput{
			CurrentTime: currentTime,
		})
		if err != nil {
			if !errors.Is(err, types.ErrEmptyValue) && !errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if conf.CurrentTime.Value().Location() != time.UTC {
			t.Fatalf(
				"CurrentTime location = %v, want UTC",
				conf.CurrentTime.Value().Location(),
			)
		}
	})
}
