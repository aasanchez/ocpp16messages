package tests_json_test

import (
	"testing"

	cr "github.com/aasanchez/ocpp16messages/cancelReservation"
)

func TestCancelReservationReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := cr.Req(cr.ReqInput{ReservationId: 0})
	if err != nil {
		t.Fatalf("cancelReservation.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestCancelReservationConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := cr.Conf(cr.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Fatalf("cancelReservation.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
