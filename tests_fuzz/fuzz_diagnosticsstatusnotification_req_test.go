//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16messages/diagnosticsstatusnotification"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzDiagnosticsStatusNotificationReq(f *testing.F) {
	f.Add(types.DiagnosticsStatusIdle.String())
	f.Add(types.DiagnosticsStatusUploaded.String())
	f.Add(types.DiagnosticsStatusUploading.String())
	f.Add(types.DiagnosticsStatusUploadFailed.String())
	f.Add("bad-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		req, err := diagnosticsstatusnotification.Req(
			diagnosticsstatusnotification.ReqInput{
				Status: status,
			},
		)
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if !req.Status.IsValid() {
			t.Fatalf("Status = %q, want valid", req.Status.String())
		}
		if req.Status.String() != status {
			t.Fatalf("Status = %q, want %q", req.Status.String(), status)
		}
	})
}
