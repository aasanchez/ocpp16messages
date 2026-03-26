//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16messages/firmwarestatusnotification"
	ft "github.com/aasanchez/ocpp16messages/firmwarestatusnotification/types"
	st "github.com/aasanchez/ocpp16messages/types"
)

func FuzzFirmwareStatusNotificationReq(f *testing.F) {
	f.Add(ft.FirmwareStatusDownloaded.String())
	f.Add(ft.FirmwareStatusDownloadFailed.String())
	f.Add(ft.FirmwareStatusDownloading.String())
	f.Add(ft.FirmwareStatusInstalled.String())
	f.Add(ft.FirmwareStatusInstallationFailed.String())
	f.Add(ft.FirmwareStatusInstalling.String())
	f.Add("bad-status")

	f.Fuzz(func(t *testing.T, status string) {
		if len(status) > maxFuzzStringLen {
			t.Skip()
		}

		req, err := firmwarestatusnotification.Req(
			firmwarestatusnotification.ReqInput{
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
