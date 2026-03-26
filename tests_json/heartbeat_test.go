package testsjson_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/heartbeat"
)

func TestHeartbeatReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := heartbeat.Req(heartbeat.ReqInput{})
	if err != nil {
		t.Fatalf("heartbeat.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestHeartbeatConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := heartbeat.Conf(heartbeat.ConfInput{
		CurrentTime: "2025-01-15T10:30:00Z",
	})
	if err != nil {
		t.Fatalf("heartbeat.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
