package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/changeAvailability/types"
)

// ExampleAvailabilityStatus_IsValid demonstrates validating an
// AvailabilityStatus.
func ExampleAvailabilityStatus_IsValid() {
	accepted := types.AvailabilityStatus("Accepted")
	fmt.Println("Accepted valid:", accepted.IsValid())

	rejected := types.AvailabilityStatus("Rejected")
	fmt.Println("Rejected valid:", rejected.IsValid())

	scheduled := types.AvailabilityStatus("Scheduled")
	fmt.Println("Scheduled valid:", scheduled.IsValid())

	invalid := types.AvailabilityStatus("Unknown")
	fmt.Println("Unknown valid:", invalid.IsValid())
	// Output:
	// Accepted valid: true
	// Rejected valid: true
	// Scheduled valid: true
	// Unknown valid: false
}

// ExampleAvailabilityStatus_String demonstrates getting the string
// representation of an AvailabilityStatus.
func ExampleAvailabilityStatus_String() {
	scheduled := types.AvailabilityStatusScheduled
	fmt.Println(scheduled.String())
	// Output:
	// Scheduled
}
