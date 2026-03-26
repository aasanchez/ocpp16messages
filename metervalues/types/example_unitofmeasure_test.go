package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/metervalues/types"
)

// ExampleUnitOfMeasure demonstrates using the UnitOfMeasure enumeration
// constants defined in OCPP 1.6.
func ExampleUnitOfMeasure() {
	unit := types.Wh
	fmt.Println("Unit:", unit.String())
	fmt.Println("IsValid:", unit.IsValid())
	// Output:
	// Unit: Wh
	// IsValid: true
}

// ExampleUnitOfMeasure_commonValues demonstrates commonly used UnitOfMeasure
// values.
func ExampleUnitOfMeasure_commonValues() {
	units := []types.UnitOfMeasure{
		types.Wh,
		types.KWh,
		types.W,
		types.A,
		types.V,
		types.Percent,
	}

	for _, u := range units {
		fmt.Println(u.String())
	}
	// Output:
	// Wh
	// kWh
	// W
	// A
	// V
	// Percent
}
