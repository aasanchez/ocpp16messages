package testsjson_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/stoptransaction"
)

func TestStopTransactionReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionId:   12345,
		IdTag:           nil,
		MeterStop:       5000,
		Timestamp:       "2025-01-15T10:30:00Z",
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		t.Fatalf("stoptransaction.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestStopTransactionConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	status := "Accepted"

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIdTag: nil,
	})
	if err != nil {
		t.Fatalf("stoptransaction.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
