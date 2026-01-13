package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/changeAvailability/types"
)

// ExampleAvailabilityType_IsValid demonstrates validating an AvailabilityType.
func ExampleAvailabilityType_IsValid() {
	operative := types.AvailabilityType("Operative")
	fmt.Println("Operative valid:", operative.IsValid())

	inoperative := types.AvailabilityType("Inoperative")
	fmt.Println("Inoperative valid:", inoperative.IsValid())

	invalid := types.AvailabilityType("Unknown")
	fmt.Println("Unknown valid:", invalid.IsValid())
	// Output:
	// Operative valid: true
	// Inoperative valid: true
	// Unknown valid: false
}

// ExampleAvailabilityType_String demonstrates getting the string
// representation of an AvailabilityType.
func ExampleAvailabilityType_String() {
	operative := types.AvailabilityTypeOperative
	fmt.Println(operative.String())
	// Output:
	// Operative
}
