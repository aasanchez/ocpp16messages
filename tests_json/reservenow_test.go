package testsjson_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/reservenow"
)

func TestReserveNowReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := reservenow.Req(reservenow.ReqInput{
		ReservationId: 1,
		ConnectorId:   1,
		IdTag:         "RFID-TAG-12345",
		ExpiryDate:    "2025-01-15T10:00:00Z",
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Fatalf("reservenow.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestReserveNowConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := reservenow.Conf(reservenow.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		t.Fatalf("reservenow.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
