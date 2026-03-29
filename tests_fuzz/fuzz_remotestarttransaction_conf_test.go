//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	rst "github.com/aasanchez/ocpp16messages/remotestarttransaction"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzRemoteStartTransactionConf(f *testing.F) {
	f.Add(types.RemoteStartTransactionStatusAccepted.String())
	f.Add(types.RemoteStartTransactionStatusRejected.String())
	f.Add("")
	f.Add("invalid-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		conf, err := rst.Conf(rst.ConfInput{Status: status})
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
