package testsjson_test

import (
	"testing"

	cc "github.com/aasanchez/ocpp16messages/clearcache"
)

func TestClearCacheReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := cc.Req(cc.ReqInput{})
	if err != nil {
		t.Fatalf("clearcache.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestClearCacheConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := cc.Conf(cc.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Fatalf("clearcache.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
