package testsjson_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/getconfiguration"
	types "github.com/aasanchez/ocpp16types"
)

func TestGetConfigurationReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := getconfiguration.Req(getconfiguration.ReqInput{
		Key: []string{"HeartbeatInterval"},
	})
	if err != nil {
		t.Fatalf("getconfiguration.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestGetConfigurationConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	value := "300"

	conf, err := getconfiguration.Conf(getconfiguration.ConfInput{
		ConfigurationKey: []types.KeyValueInput{
			{
				Key:      "HeartbeatInterval",
				Readonly: false,
				Value:    &value,
			},
		},
		UnknownKey: nil,
	})
	if err != nil {
		t.Fatalf("getconfiguration.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
