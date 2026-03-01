package tests_json_test

import (
	"testing"

	gd "github.com/aasanchez/ocpp16messages/getDiagnostics"
)

func TestGetDiagnosticsReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := gd.Req(gd.ReqInput{
		Location:      "https://example.com/diagnostics",
		Retries:       nil,
		RetryInterval: nil,
		StartTime:     nil,
		StopTime:      nil,
	})
	if err != nil {
		t.Fatalf("getDiagnostics.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestGetDiagnosticsConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	fileName := "diagnostics_20250101.zip"

	conf, err := gd.Conf(gd.ConfInput{
		FileName: &fileName,
	})
	if err != nil {
		t.Fatalf("getDiagnostics.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
