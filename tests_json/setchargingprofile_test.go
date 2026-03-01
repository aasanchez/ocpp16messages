package tests_json_test

import (
	"testing"

	"github.com/aasanchez/ocpp16messages/setChargingProfile"
	spt "github.com/aasanchez/ocpp16messages/setChargingProfile/types"
	"github.com/aasanchez/ocpp16messages/types"
)

func TestSetChargingProfileReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := setChargingProfile.Req(setChargingProfile.ReqInput{
		ConnectorId: 0,
		CsChargingProfiles: spt.ChargingProfileInput{
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
		t.Fatalf("setChargingProfile.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestSetChargingProfileConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := setChargingProfile.Conf(setChargingProfile.ConfInput{
		Status: "Accepted",
	})
	if err != nil {
		t.Fatalf("setChargingProfile.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
