package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
)

const (
	exampleDurationSixHundred  = 600
	exampleMinChargingRateFive = 5.0
)

// ExampleNewChargingSchedule demonstrates creating a ChargingSchedule
// with required fields only.
func ExampleNewChargingSchedule() {
	schedule, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		ChargingRateUnit: "W",
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{StartPeriod: 0, Limit: 10.0},
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ChargingRateUnit:", schedule.ChargingRateUnit().String())
	fmt.Println("Periods:", len(schedule.ChargingSchedulePeriod()))
	// Output:
	// ChargingRateUnit: W
	// Periods: 1
}

// ExampleNewChargingSchedule_withAllFields demonstrates creating a
// ChargingSchedule with all optional fields.
func ExampleNewChargingSchedule_withAllFields() {
	duration := exampleDurationSixHundred
	minRate := exampleMinChargingRateFive

	schedule, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		Duration:         &duration,
		ChargingRateUnit: "A",
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{StartPeriod: 0, Limit: 16.0},
			{StartPeriod: 300, Limit: 32.0},
		},
		MinChargingRate: &minRate,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Duration:", schedule.Duration().Value())
	fmt.Println("ChargingRateUnit:", schedule.ChargingRateUnit().String())
	fmt.Println("Periods:", len(schedule.ChargingSchedulePeriod()))
	fmt.Println("MinChargingRate:", *schedule.MinChargingRate())
	// Output:
	// Duration: 600
	// ChargingRateUnit: A
	// Periods: 2
	// MinChargingRate: 5
}

// ExampleNewChargingSchedule_invalidChargingRateUnit demonstrates the error
// returned when an invalid charging rate unit is provided.
func ExampleNewChargingSchedule_invalidChargingRateUnit() {
	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		ChargingRateUnit: "X",
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
			{StartPeriod: 0, Limit: 10.0},
		},
	})
	if err != nil {
		fmt.Println("Error: invalid charging rate unit")
	}
	// Output:
	// Error: invalid charging rate unit
}

// ExampleNewChargingSchedule_emptyPeriods demonstrates the error returned
// when no charging schedule periods are provided.
func ExampleNewChargingSchedule_emptyPeriods() {
	_, err := types.NewChargingSchedule(types.ChargingScheduleInput{
		ChargingRateUnit:       "W",
		ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{},
	})
	if err != nil {
		fmt.Println("Error: empty charging schedule periods")
	}
	// Output:
	// Error: empty charging schedule periods
}
