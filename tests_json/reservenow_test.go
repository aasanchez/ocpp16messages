package tests_json_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/reserveNow"
)

func TestReserveNowReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := reserveNow.Req(reserveNow.ReqInput{
		ReservationId: 1,
		ConnectorId:   1,
		IdTag:         "RFID-TAG-12345",
		ExpiryDate:    "2025-01-15T10:00:00Z",
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Fatalf("reserveNow.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestReserveNowConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := reserveNow.Conf(reserveNow.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		t.Fatalf("reserveNow.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

