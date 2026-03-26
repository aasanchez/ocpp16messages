package testsjson_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/triggermessage"
)

func TestTriggerMessageReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorId:      nil,
	})
	if err != nil {
		t.Fatalf("triggermessage.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestTriggerMessageConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := triggermessage.Conf(triggermessage.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		t.Fatalf("triggermessage.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
