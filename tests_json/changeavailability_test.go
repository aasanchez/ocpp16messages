package testsjson_test

import (
	"testing"

	ca "github.com/aasanchez/ocpp16messages/changeavailability"
)

func TestChangeAvailabilityReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := ca.Req(ca.ReqInput{
		ConnectorId: 0,
		Type:        "Inoperative",
	})
	if err != nil {
		t.Fatalf("changeavailability.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestChangeAvailabilityConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := ca.Conf(ca.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Fatalf("changeavailability.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
