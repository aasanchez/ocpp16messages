package tests_json_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/startTransaction"
)

func TestStartTransactionReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := startTransaction.Req(startTransaction.ReqInput{
		ConnectorId:   1,
		IdTag:         "RFID-TAG-12345",
		MeterStart:    1000,
		Timestamp:     "2025-01-15T10:30:00Z",
		ReservationId: nil,
	})
	if err != nil {
		t.Fatalf("startTransaction.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestStartTransactionConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := startTransaction.Conf(startTransaction.ConfInput{
		TransactionId: 12345,
		Status:        "Accepted",
		ExpiryDate:    nil,
		ParentIdTag:   nil,
	})
	if err != nil {
		t.Fatalf("startTransaction.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

