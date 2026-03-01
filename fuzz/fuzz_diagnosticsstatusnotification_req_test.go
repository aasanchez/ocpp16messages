//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16messages/diagnosticsStatusNotification"
	dt "github.com/aasanchez/ocpp16messages/diagnosticsStatusNotification/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzDiagnosticsStatusNotificationReq(f *testing.F) {
	f.Add(dt.DiagnosticsStatusIdle.String())
	f.Add(dt.DiagnosticsStatusUploaded.String())
	f.Add(dt.DiagnosticsStatusUploading.String())
	f.Add(dt.DiagnosticsStatusUploadFailed.String())
	f.Add("bad-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		req, err := diagnosticsStatusNotification.Req(
			diagnosticsStatusNotification.ReqInput{
				Status: status,
			},
		)
		if err != nil {
			if !errors.Is(err, st.ErrInvalidValue) {
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
