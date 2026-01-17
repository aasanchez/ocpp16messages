package heartbeat_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/heartbeat"
)

func TestReq_Valid(t *testing.T) {
	t.Parallel()

	_, err := heartbeat.Req(heartbeat.ReqInput{})
	if err != nil {
		t.Errorf("Req() unexpected error = %v", err)
	}
}

func TestReq_AlwaysSucceeds(t *testing.T) {
	t.Parallel()

	req, err := heartbeat.Req(heartbeat.ReqInput{})
	if err != nil {
		t.Errorf("Req() unexpected error = %v", err)
	}

	// ReqMessage is an empty struct, so just verify it's the zero value
	if req != (heartbeat.ReqMessage{}) {
		t.Error("Req() returned non-empty ReqMessage, want empty struct")
	}
}
