package testsjson_test

import (
	"testing"

	gcs "github.com/aasanchez/ocpp16messages/getcompositeschedule"
)

func TestGetCompositeScheduleReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := gcs.Req(gcs.ReqInput{
		ConnectorId:      1,
		Duration:         300,
		ChargingRateUnit: nil,
	})
	if err != nil {
		t.Fatalf("getcompositeschedule.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
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
		t.Fatalf("getcompositeschedule.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
