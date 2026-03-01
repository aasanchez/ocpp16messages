package tests_json_test

import (
	"testing"

	ca "github.com/aasanchez/ocpp16messages/changeAvailability"
)

func TestChangeAvailabilityReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := ca.Req(ca.ReqInput{
		ConnectorId: 0,
		Type:        "Inoperative",
	})
	if err != nil {
		t.Fatalf("changeAvailability.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestChangeAvailabilityConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := ca.Conf(ca.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Fatalf("changeAvailability.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

