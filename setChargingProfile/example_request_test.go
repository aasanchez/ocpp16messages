package setchargingprofile_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/setchargingprofile"
	spt "github.com/aasanchez/ocpp16messages/setchargingprofile/types"
	"github.com/aasanchez/ocpp16messages/types"
)

const (
	exampleConnectorId     = 1
	exampleProfileId       = 100
	exampleStackLevel      = 0
	exampleStartPeriod     = 0
	exampleLimit           = 32.0
	fmtConnectorId         = "ConnectorId: %d\n"
	fmtChargingProfileId   = "ChargingProfileId: %d\n"
	fmtChargingProfileKind = "ChargingProfileKind: %s\n"
)

// ExampleReq demonstrates creating a valid SetChargingProfile.req message
// with a basic charging profile.
func ExampleReq() {
	req, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId: exampleConnectorId,
		CsChargingProfiles: spt.ChargingProfileInput{
			ChargingProfileId:      exampleProfileId,
			TransactionId:          nil,
			StackLevel:             exampleStackLevel,
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
						StartPeriod:  exampleStartPeriod,
						Limit:        exampleLimit,
						NumberPhases: nil,
					},
				},
				MinChargingRate: nil,
				StartSchedule:   nil,
			},
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf(fmtConnectorId, req.ConnectorId.Value())

	profileId := req.CsChargingProfiles.ChargingProfileId().Value()
	fmt.Printf(fmtChargingProfileId, profileId)

	profileKind := req.CsChargingProfiles.ChargingProfileKind().String()
	fmt.Printf(fmtChargingProfileKind, profileKind)
	// Output:
	// ConnectorId: 1
	// ChargingProfileId: 100
	// ChargingProfileKind: Absolute
}

// ExampleReq_chargePointMaxProfile demonstrates creating a
// SetChargingProfile.req message with a ChargePointMaxProfile purpose
// for limiting overall charge point power.
func ExampleReq_chargePointMaxProfile() {
	req, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId: 0,
		CsChargingProfiles: spt.ChargingProfileInput{
			ChargingProfileId:      exampleProfileId,
			TransactionId:          nil,
			StackLevel:             exampleStackLevel,
			ChargingProfilePurpose: "ChargePointMaxProfile",
			ChargingProfileKind:    "Absolute",
			RecurrencyKind:         nil,
			ValidFrom:              nil,
			ValidTo:                nil,
			ChargingSchedule: types.ChargingScheduleInput{
				Duration:         nil,
				ChargingRateUnit: "A",
				ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
					{
						StartPeriod:  exampleStartPeriod,
						Limit:        exampleLimit,
						NumberPhases: nil,
					},
				},
				MinChargingRate: nil,
				StartSchedule:   nil,
			},
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	purpose := req.CsChargingProfiles.ChargingProfilePurpose().String()
	fmt.Println("Purpose:", purpose)
	// Output:
	// Purpose: ChargePointMaxProfile
}

// ExampleReq_invalidConnectorId demonstrates the error returned when
// a negative connector ID is provided.
func ExampleReq_invalidConnectorId() {
	_, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorId: -1,
		CsChargingProfiles: spt.ChargingProfileInput{
			ChargingProfileId:      exampleProfileId,
			TransactionId:          nil,
			StackLevel:             exampleStackLevel,
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
						StartPeriod:  exampleStartPeriod,
						Limit:        exampleLimit,
						NumberPhases: nil,
					},
				},
				MinChargingRate: nil,
				StartSchedule:   nil,
			},
		},
	})
	if err != nil {
		fmt.Println("Error: invalid connector ID")
	}
	// Output:
	// Error: invalid connector ID
}
