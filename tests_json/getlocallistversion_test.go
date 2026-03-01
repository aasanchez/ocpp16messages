package tests_json_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/getLocalListVersion"
)

func TestGetLocalListVersionReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := getLocalListVersion.Req(getLocalListVersion.ReqInput{})
	if err != nil {
		t.Fatalf("getLocalListVersion.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestGetLocalListVersionConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := getLocalListVersion.Conf(getLocalListVersion.ConfInput{
		ListVersion: 5,
	})
	if err != nil {
		t.Fatalf("getLocalListVersion.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

