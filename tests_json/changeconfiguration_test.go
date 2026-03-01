package tests_json_test

import (
	"testing"

	cc "github.com/aasanchez/ocpp16messages/changeConfiguration"
)

func TestChangeConfigurationReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := cc.Req(cc.ReqInput{
		Key:   "HeartbeatInterval",
		Value: "900",
	})
	if err != nil {
		t.Fatalf("changeConfiguration.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestChangeConfigurationConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := cc.Conf(cc.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Fatalf("changeConfiguration.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

