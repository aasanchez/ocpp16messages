package testsjson_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/getlocallistversion"
)

func TestGetLocalListVersionReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := getlocallistversion.Req(getlocallistversion.ReqInput{})
	if err != nil {
		t.Fatalf("getlocallistversion.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestGetLocalListVersionConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := getlocallistversion.Conf(getlocallistversion.ConfInput{
		ListVersion: 5,
	})
	if err != nil {
		t.Fatalf("getlocallistversion.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
