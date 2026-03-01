package tests_json_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/getConfiguration"
	gct "github.com/aasanchez/ocpp16messages/getConfiguration/types"
)

func TestGetConfigurationReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := getConfiguration.Req(getConfiguration.ReqInput{
		Key: []string{"HeartbeatInterval"},
	})
	if err != nil {
		t.Fatalf("getConfiguration.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	decoded := roundTripJSON(t, req)
	assertAllFieldsValid(t, decoded)
}

func TestGetConfigurationConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	value := "300"

	conf, err := getConfiguration.Conf(getConfiguration.ConfInput{
		ConfigurationKey: []gct.KeyValueInput{
			{
				Key:      "HeartbeatInterval",
				Readonly: false,
				Value:    &value,
			},
		},
		UnknownKey: nil,
	})
	if err != nil {
		t.Fatalf("getConfiguration.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	decoded := roundTripJSON(t, conf)
	assertAllFieldsValid(t, decoded)
}

