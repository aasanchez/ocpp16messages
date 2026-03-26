package testsjson_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/sendlocallist"
)

func TestSendLocalListReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            1,
		LocalAuthorizationList: nil,
		UpdateType:             "Full",
	})
	if err != nil {
		t.Fatalf("sendlocallist.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestSendLocalListConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := sendlocallist.Conf(sendlocallist.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Fatalf("sendlocallist.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
