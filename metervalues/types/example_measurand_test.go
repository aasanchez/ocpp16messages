package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/metervalues/types"
)

// ExampleMeasurand demonstrates using the Measurand enumeration constants
// defined in OCPP 1.6.
func ExampleMeasurand() {
	measurand := types.EnergyActiveImportRegister
	fmt.Println("Measurand:", measurand.String())
	fmt.Println("IsValid:", measurand.IsValid())
	// Output:
	// Measurand: Energy.Active.Import.Register
	// IsValid: true
}

// ExampleMeasurand_commonValues demonstrates commonly used Measurand values.
func ExampleMeasurand_commonValues() {
	measurands := []types.Measurand{
		types.EnergyActiveImportRegister,
		types.CurrentImport,
		types.Voltage,
		types.PowerActiveImport,
		types.SoC,
	}

	for _, m := range measurands {
		fmt.Println(m.String())
	}
	// Output:
	// Energy.Active.Import.Register
	// Current.Import
	// Voltage
	// Power.Active.Import
	// SoC
}
