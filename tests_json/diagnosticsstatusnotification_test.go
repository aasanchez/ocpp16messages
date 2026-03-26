package testsjson_test

import (
	"testing"

	dsn "github.com/aasanchez/ocpp16messages/diagnosticsstatusnotification"
)

func TestDiagnosticsStatusNotificationReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := dsn.Req(dsn.ReqInput{Status: "Idle"})
	if err != nil {
		t.Fatalf("diagnosticsstatusnotification.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestDiagnosticsStatusNotificationConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := dsn.Conf(dsn.ConfInput{})
	if err != nil {
		t.Fatalf("diagnosticsstatusnotification.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
