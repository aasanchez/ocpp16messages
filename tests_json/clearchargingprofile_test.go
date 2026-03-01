package tests_json_test

import (
	"testing"

	ccp "github.com/aasanchez/ocpp16messages/clearChargingProfile"
)

func TestClearChargingProfileReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorId:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		t.Fatalf("clearChargingProfile.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestClearChargingProfileConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := ccp.Conf(ccp.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Fatalf("clearChargingProfile.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
