package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/meterValues/types"
)

// ExampleLocation demonstrates using the Location enumeration constants
// defined in OCPP 1.6.
func ExampleLocation() {
	location := types.Outlet
	fmt.Println("Location:", location.String())
	fmt.Println("IsValid:", location.IsValid())
	// Output:
	// Location: Outlet
	// IsValid: true
}

// ExampleLocation_allValues demonstrates all valid Location values
// as defined in OCPP 1.6.
func ExampleLocation_allValues() {
	locations := []types.Location{
		types.Body,
		types.Cable,
		types.EV,
		types.Inlet,
		types.Outlet,
	}

	for _, loc := range locations {
		fmt.Println(loc.String())
	}
	// Output:
	// Body
	// Cable
	// EV
	// Inlet
	// Outlet
}
