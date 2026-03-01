package tests_json_test

import (
	"testing"

	fsn "github.com/aasanchez/ocpp16messages/firmwareStatusNotification"
)

func TestFirmwareStatusNotificationReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := fsn.Req(fsn.ReqInput{Status: "Idle"})
	if err != nil {
		t.Fatalf("firmwareStatusNotification.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestFirmwareStatusNotificationConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := fsn.Conf(fsn.ConfInput{})
	if err != nil {
		t.Fatalf("firmwareStatusNotification.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

