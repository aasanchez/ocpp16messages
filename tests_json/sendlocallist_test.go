package tests_json_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/sendLocalList"
)

func TestSendLocalListReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := sendLocalList.Req(sendLocalList.ReqInput{
		ListVersion:            1,
		LocalAuthorizationList: nil,
		UpdateType:             "Full",
	})
	if err != nil {
		t.Fatalf("sendLocalList.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestSendLocalListConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := sendLocalList.Conf(sendLocalList.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Fatalf("sendLocalList.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

