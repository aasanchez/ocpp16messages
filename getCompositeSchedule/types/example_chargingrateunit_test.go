package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/getCompositeSchedule/types"
)

// ExampleChargingRateUnit demonstrates using the ChargingRateUnit
// enumeration constants defined in OCPP 1.6.
func ExampleChargingRateUnit() {
	unit := types.ChargingRateUnitWatts
	fmt.Println("Unit:", unit.String())
	fmt.Println("IsValid:", unit.IsValid())
	// Output:
	// Unit: W
	// IsValid: true
}

// ExampleChargingRateUnit_allValues demonstrates all valid ChargingRateUnit
// values as defined in OCPP 1.6.
func ExampleChargingRateUnit_allValues() {
	units := []types.ChargingRateUnit{
		types.ChargingRateUnitWatts,
		types.ChargingRateUnitAmperes,
	}

	for _, u := range units {
		fmt.Println(u.String())
	}
	// Output:
	// W
	// A
}

// ExampleChargingRateUnit_IsValid demonstrates validation of ChargingRateUnit
// values.
func ExampleChargingRateUnit_IsValid() {
	valid := types.ChargingRateUnitAmperes
	invalid := types.ChargingRateUnit("X")

	fmt.Println("Amperes IsValid:", valid.IsValid())
	fmt.Println("Invalid IsValid:", invalid.IsValid())
	// Output:
	// Amperes IsValid: true
	// Invalid IsValid: false
}
