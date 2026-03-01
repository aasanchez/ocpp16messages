package tests_json_test

import (
	"testing"

	rst "github.com/aasanchez/ocpp16messages/remoteStartTransaction"
)

func TestRemoteStartTransactionReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := rst.Req(rst.ReqInput{
		IdTag:       "RFID-TAG-12345",
		ConnectorId: nil,
	})
	if err != nil {
		t.Fatalf("remoteStartTransaction.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestRemoteStartTransactionConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := rst.Conf(rst.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Fatalf("remoteStartTransaction.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

