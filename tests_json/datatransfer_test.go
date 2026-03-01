package tests_json_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/dataTransfer"
)

func TestDataTransferReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := dataTransfer.Req(dataTransfer.ReqInput{
		VendorId:  "com.example.vendor",
		MessageId: nil,
		Data:      nil,
	})
	if err != nil {
		t.Fatalf("dataTransfer.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestDataTransferConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := dataTransfer.Conf(dataTransfer.ConfInput{
		Status: "Accepted",
		Data:   nil,
	})
	if err != nil {
		t.Fatalf("dataTransfer.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

