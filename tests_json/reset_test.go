package tests_json_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/reset"
)

func TestResetReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := reset.Req(reset.ReqInput{Type: "Hard"})
	if err != nil {
		t.Fatalf("reset.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestResetConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := reset.Conf(reset.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Fatalf("reset.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
