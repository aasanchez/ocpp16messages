package tests_json_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/stopTransaction"
)

func TestStopTransactionReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := stopTransaction.Req(stopTransaction.ReqInput{
		TransactionId:   12345,
		IdTag:           nil,
		MeterStop:       5000,
		Timestamp:       "2025-01-15T10:30:00Z",
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		t.Fatalf("stopTransaction.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestStopTransactionConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	status := "Accepted"

	conf, err := stopTransaction.Conf(stopTransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	})
	if err != nil {
		t.Fatalf("stopTransaction.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
