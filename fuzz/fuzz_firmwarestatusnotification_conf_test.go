//go:build fuzz

package fuzz

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/firmwareStatusNotification"
)

func FuzzFirmwareStatusNotificationConf(f *testing.F) {
	f.Add(uint8(0))

	f.Fuzz(func(t *testing.T, _ uint8) {
		_, err := firmwareStatusNotification.Conf(
			firmwareStatusNotification.ConfInput{},
		)
		if err != nil {
			t.Fatalf("error = %v, want nil", err)
		}
	})
}
