package tests_json_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/meterValues"
	"github.com/aasanchez/ocpp16messages/types"
)

func TestMeterValuesReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := meterValues.Req(meterValues.ReqInput{
		ConnectorId:   1,
		TransactionId: nil,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: "2025-01-02T15:00:00Z",
				SampledValue: []types.SampledValueInput{
					{
						Value:     "100",
						Context:   nil,
						Format:    nil,
						Measurand: nil,
						Phase:     nil,
						Location:  nil,
						Unit:      nil,
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("meterValues.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestMeterValuesConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := meterValues.Conf(meterValues.ConfInput{})
	if err != nil {
		t.Fatalf("meterValues.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
