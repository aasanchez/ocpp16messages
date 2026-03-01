package tests_json_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/triggerMessage"
)

func TestTriggerMessageReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := triggerMessage.Req(triggerMessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorId:      nil,
	})
	if err != nil {
		t.Fatalf("triggerMessage.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestTriggerMessageConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := triggerMessage.Conf(triggerMessage.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		t.Fatalf("triggerMessage.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
