package testsjson_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/starttransaction"
)

func TestStartTransactionReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorId:   1,
		IdTag:         "RFID-TAG-12345",
		MeterStart:    1000,
		Timestamp:     "2025-01-15T10:30:00Z",
		ReservationId: nil,
	})
	if err != nil {
		t.Fatalf("starttransaction.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestStartTransactionConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionId: 12345,
		Status:        "Accepted",
		ExpiryDate:    nil,
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Fatalf("starttransaction.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
