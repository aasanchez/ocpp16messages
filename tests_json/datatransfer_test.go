package testsjson_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/datatransfer"
)

func TestDataTransferReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := datatransfer.Req(datatransfer.ReqInput{
		VendorId:  "com.example.vendor",
		MessageId: nil,
		Data:      nil,
	})
	if err != nil {
		t.Fatalf("datatransfer.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestDataTransferConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := datatransfer.Conf(datatransfer.ConfInput{
		Status: "Accepted",
		Data:   nil,
	})
	if err != nil {
		t.Fatalf("datatransfer.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
