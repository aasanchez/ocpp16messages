package testsjson_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/metervalues"
	"github.com/aasanchez/ocpp16messages/types"
)

func TestMeterValuesReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := metervalues.Req(metervalues.ReqInput{
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
		t.Fatalf("metervalues.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestMeterValuesConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := metervalues.Conf(metervalues.ConfInput{})
	if err != nil {
		t.Fatalf("metervalues.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
