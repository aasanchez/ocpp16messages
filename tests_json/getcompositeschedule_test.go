package tests_json_test

import (
	"testing"

	gcs "github.com/aasanchez/ocpp16messages/getCompositeSchedule"
)

func TestGetCompositeScheduleReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := gcs.Req(gcs.ReqInput{
		ConnectorId:      1,
		Duration:         300,
		ChargingRateUnit: nil,
	})
	if err != nil {
		t.Fatalf("getCompositeSchedule.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestGetCompositeScheduleConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorId:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err != nil {
		t.Fatalf("getCompositeSchedule.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

