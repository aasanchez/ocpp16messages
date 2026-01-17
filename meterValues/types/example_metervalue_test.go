package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/meterValues/types"
)

const exampleTimestamp = "2025-01-02T15:00:00Z"

// ExampleNewMeterValue demonstrates creating a MeterValue with a single
// sampled value.
func ExampleNewMeterValue() {
	input := types.MeterValueInput{
		Timestamp: exampleTimestamp,
		SampledValue: []types.SampledValueInput{
			{
				Value:     "12500",
				Context:   nil,
				Format:    nil,
				Measurand: nil,
				Phase:     nil,
				Location:  nil,
				Unit:      nil,
			},
		},
	}

	meterValue, err := types.NewMeterValue(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("SampledValue count:", len(meterValue.SampledValue))
	// Output:
	// SampledValue count: 1
}

// ExampleNewMeterValue_multipleSampledValues demonstrates creating a
// MeterValue with multiple sampled values.
func ExampleNewMeterValue_multipleSampledValues() {
	measurandEnergy := "Energy.Active.Import.Register"
	measurandVoltage := "Voltage"
	unitWh := "Wh"
	unitV := "V"

	input := types.MeterValueInput{
		Timestamp: exampleTimestamp,
		SampledValue: []types.SampledValueInput{
			{
				Value:     "12500",
				Context:   nil,
				Format:    nil,
				Measurand: &measurandEnergy,
				Phase:     nil,
				Location:  nil,
				Unit:      &unitWh,
			},
			{
				Value:     "230",
				Context:   nil,
				Format:    nil,
				Measurand: &measurandVoltage,
				Phase:     nil,
				Location:  nil,
				Unit:      &unitV,
			},
		},
	}

	meterValue, err := types.NewMeterValue(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("SampledValue count:", len(meterValue.SampledValue))
	// Output:
	// SampledValue count: 2
}

// ExampleNewMeterValue_emptySampledValue demonstrates the error when
// SampledValue is empty.
func ExampleNewMeterValue_emptySampledValue() {
	input := types.MeterValueInput{
		Timestamp:    exampleTimestamp,
		SampledValue: []types.SampledValueInput{},
	}

	_, err := types.NewMeterValue(input)
	if err != nil {
		fmt.Println("Error: SampledValue cannot be empty")
	}
	// Output:
	// Error: SampledValue cannot be empty
}
