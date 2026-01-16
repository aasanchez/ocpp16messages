package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
)

const (
	exampleStartPeriodZero   = 0
	exampleStartPeriodThirty = 30
	exampleLimitTen          = 10.0
	exampleLimitTwenty       = 20.0
	exampleNumberPhasesThree = 3
)

// ExampleNewChargingSchedulePeriod demonstrates creating a
// ChargingSchedulePeriod with required fields only.
func ExampleNewChargingSchedulePeriod() {
	period, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod: exampleStartPeriodZero,
		Limit:       exampleLimitTen,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("StartPeriod:", period.StartPeriod().Value())
	fmt.Println("Limit:", period.Limit())
	// Output:
	// StartPeriod: 0
	// Limit: 10
}

// ExampleNewChargingSchedulePeriod_withNumberPhases demonstrates creating a
// ChargingSchedulePeriod with the optional NumberPhases field.
func ExampleNewChargingSchedulePeriod_withNumberPhases() {
	phases := exampleNumberPhasesThree

	period, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod:  exampleStartPeriodThirty,
		Limit:        exampleLimitTwenty,
		NumberPhases: &phases,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("StartPeriod:", period.StartPeriod().Value())
	fmt.Println("Limit:", period.Limit())
	fmt.Println("NumberPhases:", period.NumberPhases().Value())
	// Output:
	// StartPeriod: 30
	// Limit: 20
	// NumberPhases: 3
}

// ExampleNewChargingSchedulePeriod_invalidLimit demonstrates the error
// returned when a negative limit is provided.
func ExampleNewChargingSchedulePeriod_invalidLimit() {
	_, err := types.NewChargingSchedulePeriod(types.ChargingSchedulePeriodInput{
		StartPeriod: exampleStartPeriodZero,
		Limit:       -1.0,
	})
	if err != nil {
		fmt.Println("Error: invalid limit")
	}
	// Output:
	// Error: invalid limit
}
