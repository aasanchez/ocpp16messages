//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	"github.com/aasanchez/ocpp16messages/firmwarestatusnotification"
	types "github.com/aasanchez/ocpp16types"
)

func FuzzFirmwareStatusNotificationReq(f *testing.F) {
	f.Add(types.FirmwareStatusDownloaded.String())
	f.Add(types.FirmwareStatusDownloadFailed.String())
	f.Add(types.FirmwareStatusDownloading.String())
	f.Add(types.FirmwareStatusInstalled.String())
	f.Add(types.FirmwareStatusInstallationFailed.String())
	f.Add(types.FirmwareStatusInstalling.String())
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
