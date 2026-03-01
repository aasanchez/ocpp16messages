//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	tm "github.com/aasanchez/ocpp16messages/triggerMessage"
	tt "github.com/aasanchez/ocpp16messages/triggerMessage/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzTriggerMessageConf(f *testing.F) {
	f.Add(tt.TriggerMessageStatusAccepted.String())
	f.Add(tt.TriggerMessageStatusRejected.String())
	f.Add(tt.TriggerMessageStatusNotImplemented.String())
	f.Add("")
	f.Add("invalid-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := tm.Conf(tm.ConfInput{Status: status})
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
