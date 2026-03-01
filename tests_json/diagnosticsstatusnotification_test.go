package tests_json_test

import (
	"testing"

	dsn "github.com/aasanchez/ocpp16messages/diagnosticsStatusNotification"
)

func TestDiagnosticsStatusNotificationReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := dsn.Req(dsn.ReqInput{Status: "Idle"})
	if err != nil {
		t.Fatalf("diagnosticsStatusNotification.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestDiagnosticsStatusNotificationConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := dsn.Conf(dsn.ConfInput{})
	if err != nil {
		t.Fatalf("diagnosticsStatusNotification.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

