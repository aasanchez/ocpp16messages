package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/metervalues/types"
)

// ExampleNewSampledValue demonstrates creating a SampledValue with only
// the required Value field.
func ExampleNewSampledValue() {
	input := types.SampledValueInput{
		Value:     "12500",
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	sampledValue, err := types.NewSampledValue(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Value:", sampledValue.Value.Value())
	// Output:
	// Value: 12500
}

// ExampleNewSampledValue_withOptionalFields demonstrates creating a
// SampledValue with all optional fields populated.
func ExampleNewSampledValue_withOptionalFields() {
	context := "Sample.Periodic"
	format := "Raw"
	measurand := "Energy.Active.Import.Register"
	phase := "L1"
	location := "Outlet"
	unit := "Wh"

	input := types.SampledValueInput{
		Value:     "12500",
		Context:   &context,
		Format:    &format,
		Measurand: &measurand,
		Phase:     &phase,
		Location:  &location,
		Unit:      &unit,
	}

	sampledValue, err := types.NewSampledValue(input)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Value:", sampledValue.Value.Value())
	fmt.Println("Context:", sampledValue.Context.String())
	fmt.Println("Unit:", sampledValue.Unit.String())
	// Output:
	// Value: 12500
	// Context: Sample.Periodic
	// Unit: Wh
}

// ExampleNewSampledValue_emptyValue demonstrates the error when Value is empty.
func ExampleNewSampledValue_emptyValue() {
	input := types.SampledValueInput{
		Value:     "",
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := types.NewSampledValue(input)
	if err != nil {
		fmt.Println("Error: Value cannot be empty")
	}
	// Output:
	// Error: Value cannot be empty
}
