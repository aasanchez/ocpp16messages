//go:build fuzz

package fuzz

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/heartbeat"
)

func FuzzHeartbeatReq(f *testing.F) {
	f.Add(uint8(0))

	f.Fuzz(func(t *testing.T, _ uint8) {
		_, err := heartbeat.Req(heartbeat.ReqInput{})
		if err != nil {
			t.Fatalf("error = %v, want nil", err)
		}
	})
}
