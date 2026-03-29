//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	tm "github.com/aasanchez/ocpp16messages/triggermessage"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzTriggerMessageConf(f *testing.F) {
	f.Add(types.TriggerMessageStatusAccepted.String())
	f.Add(types.TriggerMessageStatusRejected.String())
	f.Add(types.TriggerMessageStatusNotImplemented.String())
	f.Add("")
	f.Add("invalid-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := tm.Conf(tm.ConfInput{Status: status})
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
