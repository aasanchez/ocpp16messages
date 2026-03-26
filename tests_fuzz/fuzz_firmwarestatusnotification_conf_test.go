//go:build fuzz

package fuzz

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/firmwarestatusnotification"
)

func FuzzFirmwareStatusNotificationConf(f *testing.F) {
	f.Add(uint8(0))

	f.Fuzz(func(t *testing.T, _ uint8) {
		_, err := firmwarestatusnotification.Conf(
			firmwarestatusnotification.ConfInput{},
		)
		if err != nil {
			t.Fatalf("error = %v, want nil", err)
		}
	})
}
