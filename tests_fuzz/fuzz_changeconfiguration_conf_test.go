//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16messages/changeConfiguration"
	ct "github.com/aasanchez/ocpp16messages/changeConfiguration/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzChangeConfigurationConf(f *testing.F) {
	f.Add(ct.ConfigurationStatusAccepted.String())
	f.Add(ct.ConfigurationStatusRejected.String())
	f.Add(ct.ConfigurationStatusRebootRequired.String())
	f.Add(ct.ConfigurationStatusNotSupported.String())
	f.Add("bad-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := changeConfiguration.Conf(changeConfiguration.ConfInput{
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
