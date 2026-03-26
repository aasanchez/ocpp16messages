package testsjson_test

import (
	"testing"

	fsn "github.com/aasanchez/ocpp16messages/firmwarestatusnotification"
)

func TestFirmwareStatusNotificationReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := fsn.Req(fsn.ReqInput{Status: "Idle"})
	if err != nil {
		t.Fatalf("firmwarestatusnotification.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestFirmwareStatusNotificationConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := fsn.Conf(fsn.ConfInput{})
	if err != nil {
		t.Fatalf("firmwarestatusnotification.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
