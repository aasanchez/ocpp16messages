package testsjson_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/setchargingprofile"
	types "github.com/aasanchez/ocpp16types"
)

func TestSetChargingProfileReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId: 0,
		CsChargingProfiles: types.ChargingProfileInput{
			ChargingProfileId:      1,
			TransactionId:          nil,
			StackLevel:             0,
			ChargingProfilePurpose: "TxDefaultProfile",
			ChargingProfileKind:    "Absolute",
			RecurrencyKind:         nil,
			ValidFrom:              nil,
			ValidTo:                nil,
			ChargingSchedule: types.ChargingScheduleInput{
				Duration:         nil,
				ChargingRateUnit: "W",
				ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
					{
						StartPeriod:  0,
						Limit:        30.0,
						NumberPhases: nil,
					},
				},
				MinChargingRate: nil,
				StartSchedule:   nil,
			},
		},
	})
	if err != nil {
		t.Fatalf("setchargingprofile.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestSetChargingProfileConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := setchargingprofile.Conf(setchargingprofile.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		t.Fatalf("setchargingprofile.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
