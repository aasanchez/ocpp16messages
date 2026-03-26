package testsjson_test

import (
	"testing"

	uf "github.com/aasanchez/ocpp16messages/updatefirmware"
)

func TestUpdateFirmwareReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := uf.Req(uf.ReqInput{
		Location:      "https://example.com/firmware/v1.2.3.bin",
		RetrieveDate:  "2025-01-15T10:00:00Z",
		Retries:       nil,
		RetryInterval: nil,
	})
	if err != nil {
		t.Fatalf("updatefirmware.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestUpdateFirmwareConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := uf.Conf(uf.ConfInput{})
	if err != nil {
		t.Fatalf("updatefirmware.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
