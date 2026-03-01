package tests_json_test

import (
	"testing"

	cc "github.com/aasanchez/ocpp16messages/clearCache"
)

func TestClearCacheReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := cc.Req(cc.ReqInput{})
	if err != nil {
		t.Fatalf("clearCache.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestClearCacheConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := cc.Conf(cc.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Fatalf("clearCache.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

