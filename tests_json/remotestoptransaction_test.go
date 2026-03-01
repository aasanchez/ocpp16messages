package tests_json_test

import (
	"testing"

	rst "github.com/aasanchez/ocpp16messages/remoteStopTransaction"
)

func TestRemoteStopTransactionReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := rst.Req(rst.ReqInput{TransactionId: 12345})
	if err != nil {
		t.Fatalf("remoteStopTransaction.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestRemoteStopTransactionConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := rst.Conf(rst.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Fatalf("remoteStopTransaction.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

