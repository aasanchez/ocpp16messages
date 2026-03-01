//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	rst "github.com/aasanchez/ocpp16messages/remoteStartTransaction"
	rtt "github.com/aasanchez/ocpp16messages/remoteStartTransaction/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzRemoteStartTransactionConf(f *testing.F) {
	f.Add(rtt.RemoteStartTransactionStatusAccepted.String())
	f.Add(rtt.RemoteStartTransactionStatusRejected.String())
	f.Add("")
	f.Add("invalid-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := rst.Conf(rst.ConfInput{Status: status})
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
