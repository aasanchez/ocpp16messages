package types_test

import (
	"fmt"

	"github.com/aasanchez/ocpp16messages/clearChargingProfile/types"
)

// ExampleClearChargingProfileStatus demonstrates using the
// ClearChargingProfileStatus enumeration constants defined in OCPP 1.6.
func ExampleClearChargingProfileStatus() {
	status := types.ClearChargingProfileStatusAccepted
	fmt.Println("Status:", status.String())
	fmt.Println("IsValid:", status.IsValid())
	// Output:
	// Status: Accepted
	// IsValid: true
}

// ExampleClearChargingProfileStatus_allValues demonstrates all valid
// ClearChargingProfileStatus values as defined in OCPP 1.6.
func ExampleClearChargingProfileStatus_allValues() {
	statuses := []types.ClearChargingProfileStatus{
		types.ClearChargingProfileStatusAccepted,
		types.ClearChargingProfileStatusUnknown,
	}

	for _, s := range statuses {
		fmt.Println(s.String())
	}
	// Output:
	// Accepted
	// Unknown
}

// ExampleClearChargingProfileStatus_IsValid demonstrates validation of
// ClearChargingProfileStatus values.
func ExampleClearChargingProfileStatus_IsValid() {
	valid := types.ClearChargingProfileStatusAccepted
	invalid := types.ClearChargingProfileStatus("Invalid")

	fmt.Println("Accepted IsValid:", valid.IsValid())
	fmt.Println("Invalid IsValid:", invalid.IsValid())
	// Output:
	// Accepted IsValid: true
	// Invalid IsValid: false
}
