package tests_json_test

import (
	"testing"

	uc "github.com/aasanchez/ocpp16messages/unlockConnector"
)

func TestUnlockConnectorReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := uc.Req(uc.ReqInput{ConnectorId: 1})
	if err != nil {
		t.Fatalf("unlockConnector.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestUnlockConnectorConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := uc.Conf(uc.ConfInput{
		Status: "Unlocked",
	})
	if err != nil {
		t.Fatalf("unlockConnector.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
